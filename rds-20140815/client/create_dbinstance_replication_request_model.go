// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceReplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *CreateDBInstanceReplicationRequest
	GetChannelName() *string
	SetDbInstanceId(v string) *CreateDBInstanceReplicationRequest
	GetDbInstanceId() *string
	SetMasterHost(v string) *CreateDBInstanceReplicationRequest
	GetMasterHost() *string
	SetMasterPassword(v string) *CreateDBInstanceReplicationRequest
	GetMasterPassword() *string
	SetMasterPort(v int32) *CreateDBInstanceReplicationRequest
	GetMasterPort() *int32
	SetMasterUser(v string) *CreateDBInstanceReplicationRequest
	GetMasterUser() *string
	SetOwnerId(v int64) *CreateDBInstanceReplicationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateDBInstanceReplicationRequest
	GetRegionId() *string
}

type CreateDBInstanceReplicationRequest struct {
	// 复制通道名称，用于标识复制链路
	//
	// example:
	//
	// replication-channel-001
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// 目标RDS实例ID，复制链路将在此实例上创建
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1234567890abcdef
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// 主数据库主机地址，支持IP或域名
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.100
	MasterHost *string `json:"MasterHost,omitempty" xml:"MasterHost,omitempty"`
	// 主数据库密码，用于验证复制用户，需要提前经过Base64编码
	//
	// This parameter is required.
	//
	// example:
	//
	// U2VjdXJlUGFzczEyMyE=
	MasterPassword *string `json:"MasterPassword,omitempty" xml:"MasterPassword,omitempty"`
	// 主数据库端口号，通常为3306（MySQL）或5432（PostgreSQL）
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	MasterPort *int32 `json:"MasterPort,omitempty" xml:"MasterPort,omitempty"`
	// 主数据库用户名，用于建立复制连接
	//
	// This parameter is required.
	//
	// example:
	//
	// repl_user
	MasterUser *string `json:"MasterUser,omitempty" xml:"MasterUser,omitempty"`
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

func (s CreateDBInstanceReplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceReplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceReplicationRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *CreateDBInstanceReplicationRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *CreateDBInstanceReplicationRequest) GetMasterHost() *string {
	return s.MasterHost
}

func (s *CreateDBInstanceReplicationRequest) GetMasterPassword() *string {
	return s.MasterPassword
}

func (s *CreateDBInstanceReplicationRequest) GetMasterPort() *int32 {
	return s.MasterPort
}

func (s *CreateDBInstanceReplicationRequest) GetMasterUser() *string {
	return s.MasterUser
}

func (s *CreateDBInstanceReplicationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBInstanceReplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceReplicationRequest) SetChannelName(v string) *CreateDBInstanceReplicationRequest {
	s.ChannelName = &v
	return s
}

func (s *CreateDBInstanceReplicationRequest) SetDbInstanceId(v string) *CreateDBInstanceReplicationRequest {
	s.DbInstanceId = &v
	return s
}

func (s *CreateDBInstanceReplicationRequest) SetMasterHost(v string) *CreateDBInstanceReplicationRequest {
	s.MasterHost = &v
	return s
}

func (s *CreateDBInstanceReplicationRequest) SetMasterPassword(v string) *CreateDBInstanceReplicationRequest {
	s.MasterPassword = &v
	return s
}

func (s *CreateDBInstanceReplicationRequest) SetMasterPort(v int32) *CreateDBInstanceReplicationRequest {
	s.MasterPort = &v
	return s
}

func (s *CreateDBInstanceReplicationRequest) SetMasterUser(v string) *CreateDBInstanceReplicationRequest {
	s.MasterUser = &v
	return s
}

func (s *CreateDBInstanceReplicationRequest) SetOwnerId(v int64) *CreateDBInstanceReplicationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceReplicationRequest) SetRegionId(v string) *CreateDBInstanceReplicationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceReplicationRequest) Validate() error {
	return dara.Validate(s)
}
