// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVpcAuthModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyInstanceVpcAuthModeRequest
	GetDBInstanceId() *string
	SetNodeId(v string) *ModifyInstanceVpcAuthModeRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *ModifyInstanceVpcAuthModeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceVpcAuthModeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceVpcAuthModeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceVpcAuthModeRequest
	GetResourceOwnerId() *int64
	SetVpcAuthMode(v string) *ModifyInstanceVpcAuthModeRequest
	GetVpcAuthMode() *string
}

type ModifyInstanceVpcAuthModeRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the mongos node in the sharded cluster instance.
	//
	// >  This parameter can be used only when the instance type is sharded cluster.
	//
	// example:
	//
	// s-bpxxxxxxxx
	NodeId               *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specify whether to disable password-free access over VPC. Set the value to **Close**.
	//
	// example:
	//
	// Open
	VpcAuthMode *string `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
}

func (s ModifyInstanceVpcAuthModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVpcAuthModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAuthModeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyInstanceVpcAuthModeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ModifyInstanceVpcAuthModeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceVpcAuthModeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceVpcAuthModeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceVpcAuthModeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceVpcAuthModeRequest) GetVpcAuthMode() *string {
	return s.VpcAuthMode
}

func (s *ModifyInstanceVpcAuthModeRequest) SetDBInstanceId(v string) *ModifyInstanceVpcAuthModeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetNodeId(v string) *ModifyInstanceVpcAuthModeRequest {
	s.NodeId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetOwnerAccount(v string) *ModifyInstanceVpcAuthModeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetOwnerId(v int64) *ModifyInstanceVpcAuthModeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceVpcAuthModeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetResourceOwnerId(v int64) *ModifyInstanceVpcAuthModeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetVpcAuthMode(v string) *ModifyInstanceVpcAuthModeRequest {
	s.VpcAuthMode = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) Validate() error {
	return dara.Validate(s)
}
