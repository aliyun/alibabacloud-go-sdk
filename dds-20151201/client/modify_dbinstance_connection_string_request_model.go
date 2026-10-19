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
	// The current endpoint, which is the endpoint to be modified.
	//
	// example:
	//
	// s-bpxxxxxxxx.mongodb.rds.aliyuncs.com
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty"`
	// The instance ID.
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
	// The network type of the connection string when only the port number is modified.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The new endpoint. The endpoint must meet the following requirements:
	//
	// - It must start with a lowercase letter.
	//
	// - It must end with a lowercase letter or a digit.
	//
	// - It can contain only lowercase letters, digits, and hyphens (-).
	//
	// - It must be 8 to 63 characters in length.
	//
	// > Specify only the prefix of the endpoint. The remaining part cannot be modified.
	//
	// example:
	//
	// aliyuntest111
	NewConnectionString *string `json:"NewConnectionString,omitempty" xml:"NewConnectionString,omitempty"`
	// The new port. Valid values: 1000 to 65535.
	//
	// > This parameter is active only when the **DBInstanceId*	- parameter is set to a cloud disk instance ID.
	//
	// example:
	//
	// 3310
	NewPort *int32 `json:"NewPort,omitempty" xml:"NewPort,omitempty"`
	// Instance ID of the Mongos node in the sharded cluster instance. You can specify only one Mongos node ID in each call.
	//
	// > This parameter is valid only when the **DBInstanceId*	- parameter is set to instance ID of a sharded cluster instance.
	//
	// example:
	//
	// s-bpxxxxxxxx
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to modify only port numbers. If you modify only port numbers, batch modification is supported.
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
