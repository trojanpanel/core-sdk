package core

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestAddNode(t *testing.T) {
	dto := NodeAddDto{
		NodeTypeId: 4,
		Port:       443,
		Domain:     "example.com",
	}
	if err := AddNode("your_token",
		"127.0.0.1", 8100, &dto); err != nil {
		logrus.Errorf("AddNode err: %v", err)
		return
	}
	logrus.Errorf("AddNode success")
}

func TestRemoveNode(t *testing.T) {
	removeDto := NodeRemoveDto{NodeTypeId: 4, Port: 443}
	if err := RemoveNode("your_token",
		"127.0.0.1", 8100, &removeDto); err != nil {
		logrus.Errorf("RemoveNode err: %v", err)
		return
	}
	logrus.Errorf("RemoveNode success")
}

func TestRemoveAccount(t *testing.T) {
	accountRemoveDto := AccountRemoveDto{}
	if err := RemoveAccount("your_token",
		"127.0.0.1", 8100, &accountRemoveDto); err != nil {
		logrus.Errorf("RemoveAccount err: %v", err)
		return
	}
	logrus.Errorf("RemoveAccount success")
}

func TestNodeServerState(t *testing.T) {
	nodeState, err := GetNodeState("your_token",
		"127.0.0.1", 8100, 4, 443)
	if err != nil {
		logrus.Errorf("GetNodeState err: %v", err)
		return
	}
	logrus.Errorf("GetNodeState success result: %v", nodeState)
}

func TestGetNodeServerState(t *testing.T) {
	nodeServerState, err := GetNodeServerState("your_token",
		"127.0.0.1", 8100)
	if err != nil {
		logrus.Errorf("GetNodeServerState err: %v", err)
		return
	}
	logrus.Errorf("GetNodeServerState success result: %v", nodeServerState)
}

func TestGetNodeServerInfo(t *testing.T) {
	nodeServerInfo, err := GetNodeServerInfo("your_token",
		"127.0.0.1", 8100)
	if err != nil {
		logrus.Errorf("GetNodeServerInfo err: %v", err)
		return
	}
	logrus.Errorf("GetNodeServerInfo success result: %v", nodeServerInfo)
}
