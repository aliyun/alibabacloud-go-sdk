// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterResourceGroupRequest
	GetDBClusterId() *string
	SetNewResourceGroupId(v string) *ModifyDBClusterResourceGroupRequest
	GetNewResourceGroupId() *string
	SetOwnerAccount(v string) *ModifyDBClusterResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterResourceGroupRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyDBClusterResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterResourceGroupRequest
	GetResourceOwnerId() *int64
}

type ModifyDBClusterResourceGroupRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the new resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-**********
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the original resource group.
	//
	// example:
	//
	// rg-**********
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterResourceGroupRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ModifyDBClusterResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBClusterResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterResourceGroupRequest) SetDBClusterId(v string) *ModifyDBClusterResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetNewResourceGroupId(v string) *ModifyDBClusterResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetOwnerAccount(v string) *ModifyDBClusterResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetOwnerId(v int64) *ModifyDBClusterResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetResourceGroupId(v string) *ModifyDBClusterResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) SetResourceOwnerId(v int64) *ModifyDBClusterResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
