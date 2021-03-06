package battleManager

import (
	"cmgameserver/face"
	"time"
	"sync"
	"cmgameserver/message"
)

type BattleManager struct {
	GameServer face.IGameServer
	battles map[int32]face.IBattle
	sync.RWMutex
}
var genBattleId = int32(0)

func NewBattleManager(server face.IGameServer) *BattleManager{
	return &BattleManager{
		battles:make(map[int32]face.IBattle),
		GameServer:server,
	}
}
func (battleManager *BattleManager) CreateBattle(roles [4]face.IOnlineRole) face.IBattle{
	genBattleId++
	//设置战斗Id
	for _,v := range roles {
		if v != nil{
			v.SetBattleId(genBattleId)
		}
	}
	battle := &Battle{
		BattleId: genBattleId,
		done:     make(chan bool),
		timer:    time.NewTicker(80 * time.Millisecond),
		players:  roles,
		msg:      new(message.M2C_BattleFrame),
		GameServer:battleManager.GameServer,
	}
	battle.msg.Cmd = make([]*message.Command,0)
	battleManager.battles[battle.BattleId] = battle
	return battle
}
//取得空闲的战斗
func (battleManager *BattleManager)GetBattleInFree() face.IBattle{
	battleManager.Lock()
	defer battleManager.Unlock()
	for _,v := range battleManager.battles{
		if v.GetBattleState() == face.Free{
			return v
		}
	}
	return nil
}
func (battleManager *BattleManager)GetBattle(battleId int32)face.IBattle{
	return battleManager.battles[battleId]
}