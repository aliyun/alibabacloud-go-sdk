// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPendingMaintenanceActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *ModifyPendingMaintenanceActionRequest
	GetIds() *string
	SetOwnerAccount(v string) *ModifyPendingMaintenanceActionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyPendingMaintenanceActionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyPendingMaintenanceActionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyPendingMaintenanceActionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyPendingMaintenanceActionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPendingMaintenanceActionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyPendingMaintenanceActionRequest
	GetSecurityToken() *string
	SetSwitchTime(v string) *ModifyPendingMaintenanceActionRequest
	GetSwitchTime() *string
}

type ModifyPendingMaintenanceActionRequest struct {
	// The ID of the task. You can specify multiple task IDs at a time to modify the switching time of the tasks in a unified manner. The task IDs must be separated with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 111111
	Ids          *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the region ID details.
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
	// The time that you specify for the background to perform the action that corresponds to the pending event. Specify the time in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-06-09T22:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s ModifyPendingMaintenanceActionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPendingMaintenanceActionRequest) GoString() string {
	return s.String()
}

func (s *ModifyPendingMaintenanceActionRequest) GetIds() *string {
	return s.Ids
}

func (s *ModifyPendingMaintenanceActionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyPendingMaintenanceActionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPendingMaintenanceActionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPendingMaintenanceActionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyPendingMaintenanceActionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPendingMaintenanceActionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPendingMaintenanceActionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyPendingMaintenanceActionRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyPendingMaintenanceActionRequest) SetIds(v string) *ModifyPendingMaintenanceActionRequest {
	s.Ids = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetOwnerAccount(v string) *ModifyPendingMaintenanceActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetOwnerId(v int64) *ModifyPendingMaintenanceActionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetRegionId(v string) *ModifyPendingMaintenanceActionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetResourceGroupId(v string) *ModifyPendingMaintenanceActionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetResourceOwnerAccount(v string) *ModifyPendingMaintenanceActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetResourceOwnerId(v int64) *ModifyPendingMaintenanceActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetSecurityToken(v string) *ModifyPendingMaintenanceActionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetSwitchTime(v string) *ModifyPendingMaintenanceActionRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) Validate() error {
	return dara.Validate(s)
}
