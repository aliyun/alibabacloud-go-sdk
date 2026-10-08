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
	// The name of the replication channel, which is used to identify the replication channel.
	//
	// example:
	//
	// replication-channel-001
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1234567890abcdef
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The host address of the primary database. IP addresses and domain names are supported.
	//
	// example:
	//
	// 192.168.1.100
	MasterHost *string `json:"MasterHost,omitempty" xml:"MasterHost,omitempty"`
	// The password of the primary database, which is used to authenticate the replication user. The password must be Base64-encoded in advance.
	//
	// example:
	//
	// U2VjdXJlUGFzczEyMyE=
	MasterPassword *string `json:"MasterPassword,omitempty" xml:"MasterPassword,omitempty"`
	// The port number of the primary database. The default port is 3306 for MySQL.
	//
	// example:
	//
	// 3306
	MasterPort *int32 `json:"MasterPort,omitempty" xml:"MasterPort,omitempty"`
	// The username of the primary database, which is used to establish a replication connection. Specify this parameter only when an update is required.
	//
	// example:
	//
	// repl_user
	MasterUser *string `json:"MasterUser,omitempty" xml:"MasterUser,omitempty"`
	// The operation type. Specifies the operation to perform on the replication channel.
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
	// The region ID.
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
