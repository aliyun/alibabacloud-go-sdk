// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterVpcRequest
	GetDBClusterId() *string
	SetExistedEndpointSwitchType(v string) *ModifyDBClusterVpcRequest
	GetExistedEndpointSwitchType() *string
	SetOwnerAccount(v string) *ModifyDBClusterVpcRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterVpcRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterVpcRequest
	GetResourceOwnerId() *int64
	SetVPCId(v string) *ModifyDBClusterVpcRequest
	GetVPCId() *string
	SetVSwitchId(v string) *ModifyDBClusterVpcRequest
	GetVSwitchId() *string
}

type ModifyDBClusterVpcRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NONE
	ExistedEndpointSwitchType *string `json:"ExistedEndpointSwitchType,omitempty" xml:"ExistedEndpointSwitchType,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-**********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ModifyDBClusterVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterVpcRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterVpcRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterVpcRequest) GetExistedEndpointSwitchType() *string {
	return s.ExistedEndpointSwitchType
}

func (s *ModifyDBClusterVpcRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterVpcRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *ModifyDBClusterVpcRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBClusterVpcRequest) SetDBClusterId(v string) *ModifyDBClusterVpcRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterVpcRequest) SetExistedEndpointSwitchType(v string) *ModifyDBClusterVpcRequest {
	s.ExistedEndpointSwitchType = &v
	return s
}

func (s *ModifyDBClusterVpcRequest) SetOwnerAccount(v string) *ModifyDBClusterVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterVpcRequest) SetOwnerId(v int64) *ModifyDBClusterVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterVpcRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterVpcRequest) SetResourceOwnerId(v int64) *ModifyDBClusterVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterVpcRequest) SetVPCId(v string) *ModifyDBClusterVpcRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBClusterVpcRequest) SetVSwitchId(v string) *ModifyDBClusterVpcRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBClusterVpcRequest) Validate() error {
	return dara.Validate(s)
}
