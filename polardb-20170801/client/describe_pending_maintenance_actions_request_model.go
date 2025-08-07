// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePendingMaintenanceActionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsHistory(v int32) *DescribePendingMaintenanceActionsRequest
	GetIsHistory() *int32
	SetOwnerAccount(v string) *DescribePendingMaintenanceActionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePendingMaintenanceActionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribePendingMaintenanceActionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribePendingMaintenanceActionsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribePendingMaintenanceActionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePendingMaintenanceActionsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribePendingMaintenanceActionsRequest
	GetSecurityToken() *string
}

type DescribePendingMaintenanceActionsRequest struct {
	// Specifies whether to return the historical tasks. Valid values:
	//
	// 	- **0**: returns the current task.
	//
	// 	- **1**: returns the historical tasks.
	//
	// Default value: **0**.
	//
	// example:
	//
	// 1
	IsHistory    *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePendingMaintenanceActionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePendingMaintenanceActionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionsRequest) GetIsHistory() *int32 {
	return s.IsHistory
}

func (s *DescribePendingMaintenanceActionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePendingMaintenanceActionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePendingMaintenanceActionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePendingMaintenanceActionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePendingMaintenanceActionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePendingMaintenanceActionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePendingMaintenanceActionsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePendingMaintenanceActionsRequest) SetIsHistory(v int32) *DescribePendingMaintenanceActionsRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetOwnerAccount(v string) *DescribePendingMaintenanceActionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetOwnerId(v int64) *DescribePendingMaintenanceActionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetRegionId(v string) *DescribePendingMaintenanceActionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetResourceGroupId(v string) *DescribePendingMaintenanceActionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetResourceOwnerAccount(v string) *DescribePendingMaintenanceActionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetResourceOwnerId(v int64) *DescribePendingMaintenanceActionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetSecurityToken(v string) *DescribePendingMaintenanceActionsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) Validate() error {
	return dara.Validate(s)
}
