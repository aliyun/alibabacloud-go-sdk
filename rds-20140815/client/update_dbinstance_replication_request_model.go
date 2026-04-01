// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstanceReplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *UpdateDBInstanceReplicationRequest
	GetChannelName() *string
	SetDbInstanceId(v string) *UpdateDBInstanceReplicationRequest
	GetDbInstanceId() *string
	SetMasterHost(v string) *UpdateDBInstanceReplicationRequest
	GetMasterHost() *string
	SetMasterPassword(v string) *UpdateDBInstanceReplicationRequest
	GetMasterPassword() *string
	SetMasterPort(v int32) *UpdateDBInstanceReplicationRequest
	GetMasterPort() *int32
	SetMasterUser(v string) *UpdateDBInstanceReplicationRequest
	GetMasterUser() *string
	SetOperation(v string) *UpdateDBInstanceReplicationRequest
	GetOperation() *string
	SetOwnerId(v int64) *UpdateDBInstanceReplicationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateDBInstanceReplicationRequest
	GetRegionId() *string
}

type UpdateDBInstanceReplicationRequest struct {
	// 复制通道名称，用于标识复制链路
	//
	// example:
	//
	// replication-channel-001
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// 目标RDS实例ID，复制链路将在此实例上更新
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1234567890abcdef
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// 主数据库主机地址，支持IP或域名，仅在需要更新时提供
	//
	// example:
	//
	// 192.168.1.100
	MasterHost *string `json:"MasterHost,omitempty" xml:"MasterHost,omitempty"`
	// 主数据库密码，用于验证复制用户，需要提前经过Base64编码，仅在需要更新时提供
	//
	// example:
	//
	// U2VjdXJlUGFzczEyMyE=
	MasterPassword *string `json:"MasterPassword,omitempty" xml:"MasterPassword,omitempty"`
	// 主数据库端口号，通常为3306（MySQL）或5432（PostgreSQL），仅在需要更新时提供
	//
	// example:
	//
	// 3306
	MasterPort *int32 `json:"MasterPort,omitempty" xml:"MasterPort,omitempty"`
	// 主数据库用户名，用于建立复制连接，仅在需要更新时提供
	//
	// example:
	//
	// repl_user
	MasterUser *string `json:"MasterUser,omitempty" xml:"MasterUser,omitempty"`
	// 操作类型，指定对复制链路执行的操作
	//
	// This parameter is required.
	//
	// example:
	//
	// Start
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// 阿里云账号ID，用于指定资源的所有者
	//
	// example:
	//
	// 1234567890123456
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 地域ID，表示RDS实例所在的地域
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateDBInstanceReplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstanceReplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceReplicationRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *UpdateDBInstanceReplicationRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *UpdateDBInstanceReplicationRequest) GetMasterHost() *string {
	return s.MasterHost
}

func (s *UpdateDBInstanceReplicationRequest) GetMasterPassword() *string {
	return s.MasterPassword
}

func (s *UpdateDBInstanceReplicationRequest) GetMasterPort() *int32 {
	return s.MasterPort
}

func (s *UpdateDBInstanceReplicationRequest) GetMasterUser() *string {
	return s.MasterUser
}

func (s *UpdateDBInstanceReplicationRequest) GetOperation() *string {
	return s.Operation
}

func (s *UpdateDBInstanceReplicationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateDBInstanceReplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDBInstanceReplicationRequest) SetChannelName(v string) *UpdateDBInstanceReplicationRequest {
	s.ChannelName = &v
	return s
}

func (s *UpdateDBInstanceReplicationRequest) SetDbInstanceId(v string) *UpdateDBInstanceReplicationRequest {
	s.DbInstanceId = &v
	return s
}

func (s *UpdateDBInstanceReplicationRequest) SetMasterHost(v string) *UpdateDBInstanceReplicationRequest {
	s.MasterHost = &v
	return s
}

func (s *UpdateDBInstanceReplicationRequest) SetMasterPassword(v string) *UpdateDBInstanceReplicationRequest {
	s.MasterPassword = &v
	return s
}

func (s *UpdateDBInstanceReplicationRequest) SetMasterPort(v int32) *UpdateDBInstanceReplicationRequest {
	s.MasterPort = &v
	return s
}

func (s *UpdateDBInstanceReplicationRequest) SetMasterUser(v string) *UpdateDBInstanceReplicationRequest {
	s.MasterUser = &v
	return s
}

func (s *UpdateDBInstanceReplicationRequest) SetOperation(v string) *UpdateDBInstanceReplicationRequest {
	s.Operation = &v
	return s
}

func (s *UpdateDBInstanceReplicationRequest) SetOwnerId(v int64) *UpdateDBInstanceReplicationRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDBInstanceReplicationRequest) SetRegionId(v string) *UpdateDBInstanceReplicationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDBInstanceReplicationRequest) Validate() error {
	return dara.Validate(s)
}
