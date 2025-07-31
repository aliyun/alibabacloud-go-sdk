// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConnectionStringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest
	GetCurrentConnectionString() *string
	SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest
	GetDBInstanceId() *string
	SetNewConnectionString(v string) *ModifyDBInstanceConnectionStringRequest
	GetNewConnectionString() *string
	SetNewPort(v int32) *ModifyDBInstanceConnectionStringRequest
	GetNewPort() *int32
	SetNodeId(v string) *ModifyDBInstanceConnectionStringRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceConnectionStringRequest struct {
	// The current endpoint that is to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bpxxxxxxxx.mongodb.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The instance ID.
	//
	// > If you set this parameter to the ID of a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The new endpoint. It must be 8 to 64 characters in length and can contain letters and digits. It must start with a lowercase letter.
	//
	// > You need only to specify the prefix of the endpoint. The content other than the prefix cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyuntest111
	NewConnectionString *string `json:"NewConnectionString,omitempty" xml:"NewConnectionString,omitempty"`
	// The new port number of the instance. The port number must be within the range from 1000 to 65535.
	//
	// >  This parameter is available only when you set the **DBInstanceId*	- parameter to the ID of an instance that uses cloud disks.
	//
	// example:
	//
	// 3310
	NewPort *int32 `json:"NewPort,omitempty" xml:"NewPort,omitempty"`
	// The ID of the mongos in the specified sharded cluster instance. Only one mongos ID can be specified in each call.
	//
	// > This parameter is valid only when you specify the **DBInstanceId*	- parameter to the ID of a sharded cluster instance.
	//
	// example:
	//
	// s-bpxxxxxxxx
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) GetCurrentConnectionString() *string {
	return s.CurrentConnectionString
}

func (s *ModifyDBInstanceConnectionStringRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceConnectionStringRequest) GetNewConnectionString() *string {
	return s.NewConnectionString
}

func (s *ModifyDBInstanceConnectionStringRequest) GetNewPort() *int32 {
	return s.NewPort
}

func (s *ModifyDBInstanceConnectionStringRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ModifyDBInstanceConnectionStringRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceConnectionStringRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceConnectionStringRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceConnectionStringRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetNewConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.NewConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetNewPort(v int32) *ModifyDBInstanceConnectionStringRequest {
	s.NewPort = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetNodeId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) Validate() error {
	return dara.Validate(s)
}
