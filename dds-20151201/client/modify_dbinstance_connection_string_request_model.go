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
	SetForceModifySuffix(v bool) *ModifyDBInstanceConnectionStringRequest
	GetForceModifySuffix() *bool
	SetNetworkType(v string) *ModifyDBInstanceConnectionStringRequest
	GetNetworkType() *string
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
	SetPortModifyOnly(v bool) *ModifyDBInstanceConnectionStringRequest
	GetPortModifyOnly() *bool
	SetResourceOwnerAccount(v string) *ModifyDBInstanceConnectionStringRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceConnectionStringRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceConnectionStringRequest struct {
	// The current connection address—the address to modify.
	//
	// example:
	//
	// s-bpxxxxxxxx.mongodb.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The ID of the instance.
	//
	// > If you specify the ID of a sharded cluster instance, you must also specify the **NodeId*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ForceModifySuffix *bool   `json:"ForceModifySuffix,omitempty" xml:"ForceModifySuffix,omitempty"`
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The new connection address. It must meet these requirements:
	//
	// - Start with a lowercase letter.
	//
	// - End with a lowercase letter or digit.
	//
	// - Contain only lowercase letters, digits, and hyphens (-).
	//
	// - Be 8 to 63 characters long.
	//
	// > Specify only the prefix of the connection address. You cannot change any part beyond the prefix.
	//
	// example:
	//
	// aliyuntest111
	NewConnectionString *string `json:"NewConnectionString,omitempty" xml:"NewConnectionString,omitempty"`
	// The new port number. Valid values are from 1000 to 65535.
	//
	// > This parameter is valid only when **DBInstanceId*	- specifies the ID of a cloud disk instance.
	//
	// example:
	//
	// 3310
	NewPort *int32 `json:"NewPort,omitempty" xml:"NewPort,omitempty"`
	// The ID of a Mongos node in a sharded cluster instance. You can specify only one Mongos node ID per call.
	//
	// > This parameter is valid only when **DBInstanceId*	- specifies the ID of a sharded cluster instance.
	//
	// example:
	//
	// s-bpxxxxxxxx
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PortModifyOnly       *bool   `json:"PortModifyOnly,omitempty" xml:"PortModifyOnly,omitempty"`
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

func (s *ModifyDBInstanceConnectionStringRequest) GetForceModifySuffix() *bool {
	return s.ForceModifySuffix
}

func (s *ModifyDBInstanceConnectionStringRequest) GetNetworkType() *string {
	return s.NetworkType
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

func (s *ModifyDBInstanceConnectionStringRequest) GetPortModifyOnly() *bool {
	return s.PortModifyOnly
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

func (s *ModifyDBInstanceConnectionStringRequest) SetForceModifySuffix(v bool) *ModifyDBInstanceConnectionStringRequest {
	s.ForceModifySuffix = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetNetworkType(v string) *ModifyDBInstanceConnectionStringRequest {
	s.NetworkType = &v
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

func (s *ModifyDBInstanceConnectionStringRequest) SetPortModifyOnly(v bool) *ModifyDBInstanceConnectionStringRequest {
	s.PortModifyOnly = &v
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
