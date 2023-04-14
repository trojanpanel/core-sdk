package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"time"
)

// 初始化gRPC
func newGrpcInstance(token string, ip string, grpcPort uint, timeout time.Duration) (conn *grpc.ClientConn, ctx context.Context, clo func(), err error) {
	tokenParam := TokenValidateParam{
		Token: token,
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(&tokenParam),
	}
	conn, err = grpc.Dial(fmt.Sprintf("%s:%d", ip, grpcPort),
		opts...)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	clo = func() {
		cancel()
		conn.Close()
	}
	if err != nil {
		logrus.Errorf("gRPC初始化失败 err: %v", err)
		err = errors.New(GrpcError)
	}
	return
}

// AddNode 添加节点
func AddNode(token string, ip string, grpcPort uint, nodeAddDto *NodeAddDto) error {
	conn, ctx, clo, err := newGrpcInstance(token, ip, grpcPort, 4*time.Second)
	defer clo()
	if err != nil {
		return err
	}
	client := NewApiNodeServiceClient(conn)
	send, err := client.AddNode(ctx, nodeAddDto)
	if err != nil {
		logrus.Errorf("gRPC添加节点异常 ip: %s grpc port: %d err: %v", ip, grpcPort, err)
		return errors.New(GrpcError)
	}
	if send.Success {
		return nil
	}
	return errors.New(send.Msg)
}

// RemoveNode 删除节点
func RemoveNode(token string, ip string, grpcPort uint, nodeRemoveDto *NodeRemoveDto) error {
	conn, ctx, clo, err := newGrpcInstance(token, ip, grpcPort, 4*time.Second)
	defer clo()
	if err != nil {
		return err
	}
	client := NewApiNodeServiceClient(conn)
	send, err := client.RemoveNode(ctx, nodeRemoveDto)
	if err != nil {
		logrus.Errorf("gRPC删除节点异常 ip: %s grpc port: %d err: %v", ip, grpcPort, err)
		return errors.New(GrpcError)
	}
	if send.Success {
		return nil
	}
	return errors.New(send.Msg)
}

// RemoveAccount 删除账户
func RemoveAccount(token string, ip string, grpcPort uint, accountRemoveDto *AccountRemoveDto) error {
	conn, ctx, clo, err := newGrpcInstance(token, ip, grpcPort, 4*time.Second)
	defer clo()
	if err != nil {
		return err
	}
	client := NewApiAccountServiceClient(conn)
	send, err := client.RemoveAccount(ctx, accountRemoveDto)
	if err != nil {
		logrus.Errorf("gRPC删除用户异常 ip: %s grpc porr: %d err: %v", ip, grpcPort, err)
		return errors.New(GrpcError)
	}
	if send.Success {
		return nil
	}
	return errors.New(send.Msg)
}

// GetNodeState 查询节点状态
func GetNodeState(token string, ip string, grpcPort uint, nodeTypeId uint, port uint) (*NodeStateVo, error) {
	conn, ctx, clo, err := newGrpcInstance(token, ip, grpcPort, 4*time.Second)
	defer clo()
	if err != nil {
		return nil, err
	}
	client := NewApiStateServiceClient(conn)
	nodeStateDto := NodeStateDto{
		NodeTypeId: uint64(nodeTypeId),
		Port:       uint64(port),
	}
	send, err := client.GetNodeState(ctx, &nodeStateDto)
	if err != nil {
		logrus.Errorf("gRPC GetNodeState 异常 ip: %s grpc port: %d err: %v", ip, grpcPort, err)
		return nil, errors.New(GrpcError)
	}
	if send.Success {
		var nodeStateVo NodeStateVo
		if err = anypb.UnmarshalTo(send.Data, &nodeStateVo, proto.UnmarshalOptions{}); err != nil {
			logrus.Errorf("gRPC GetNodeState 返序列化异常 ip: %s grpc port: %d err: %v", ip, grpcPort, err)
			return nil, errors.New(GrpcError)
		}
		return &nodeStateVo, nil
	}
	logrus.Errorf("gRPC GetNodeState 失败 ip: %s grpc port: %d err: %v", ip, grpcPort, err)
	return nil, errors.New(send.Msg)
}

// GetNodeServerState 查询服务器状态
func GetNodeServerState(token string, ip string, grpcPort uint) (*NodeServerStateVo, error) {
	conn, ctx, clo, err := newGrpcInstance(token, ip, grpcPort, 4*time.Second)
	defer clo()
	if err != nil {
		return nil, err
	}
	client := NewApiStateServiceClient(conn)
	nodeServerStateDto := NodeServerStateDto{}
	send, err := client.GetNodeServerState(ctx, &nodeServerStateDto)
	if err != nil {
		logrus.Errorf("gRPC GetNodeServerState 异常 ip: %s grpc port: %d err: %v", ip, grpcPort, err)
		return nil, errors.New(GrpcError)
	}
	if send.Success {
		var nodeServerStateVo NodeServerStateVo
		if err = anypb.UnmarshalTo(send.Data, &nodeServerStateVo, proto.UnmarshalOptions{}); err != nil {
			logrus.Errorf("gRPC GetNodeServerState 返序列化异常 ip: %s grpc port: %d err: %v", ip, grpcPort, err)
			return nil, errors.New(GrpcError)
		}
		return &nodeServerStateVo, nil
	}
	logrus.Errorf("gRPC GetNodeServerState 失败 ip: %s grpc port: %d err: %v", ip, grpcPort, err)
	return nil, errors.New(send.Msg)
}

// GetNodeServerInfo 查询服务器信息
func GetNodeServerInfo(token string, ip string, grpcPort uint) (*NodeServerInfoVo, error) {
	conn, ctx, clo, err := newGrpcInstance(token, ip, grpcPort, 4*time.Second)
	defer clo()
	if err != nil {
		return nil, err
	}
	client := NewApiNodeServerServiceClient(conn)
	nodeServerInfoDto := NodeServerInfoDto{}
	send, err := client.GetNodeServerInfo(ctx, &nodeServerInfoDto)
	if err != nil {
		logrus.Errorf("gRPC 查询服务器状态 异常 ip: %s grpc port: %d err: %v", ip, grpcPort, err)
		return nil, errors.New(GrpcError)
	}
	if send.Success {
		var nodeServerInfoVo NodeServerInfoVo
		if err = anypb.UnmarshalTo(send.Data, &nodeServerInfoVo, proto.UnmarshalOptions{}); err != nil {
			logrus.Errorf("查询服务器状态返序列化异常 ip: %s grpc port: %d err: %v", ip, grpcPort, err)
			return nil, errors.New(GrpcError)
		}
		return &nodeServerInfoVo, nil
	}
	return nil, errors.New(send.Msg)
}
