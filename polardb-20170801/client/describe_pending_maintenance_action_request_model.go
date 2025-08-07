// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePendingMaintenanceActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsHistory(v int32) *DescribePendingMaintenanceActionRequest
	GetIsHistory() *int32
	SetOwnerAccount(v string) *DescribePendingMaintenanceActionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePendingMaintenanceActionRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribePendingMaintenanceActionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePendingMaintenanceActionRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribePendingMaintenanceActionRequest
	GetRegion() *string
	SetResourceGroupId(v string) *DescribePendingMaintenanceActionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribePendingMaintenanceActionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePendingMaintenanceActionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribePendingMaintenanceActionRequest
	GetSecurityToken() *string
	SetTaskType(v string) *DescribePendingMaintenanceActionRequest
	GetTaskType() *string
}

type DescribePendingMaintenanceActionRequest struct {
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
	// 0
	IsHistory    *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Specify the parameter to a positive integer that does not exceed the maximum value of the INTEGER data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30**, **50**, and **100**.
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the pending event. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the regions and zones that are supported by PolarDB.
	//
	// >- You can set this parameter to **all*	- to view all pending events within your account.
	//
	// >- If you set `Region` to **all**, you must set `TaskType` to **all**.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The task type of pending events. Valid values:
	//
	// 	- **DatabaseSoftwareUpgrading**: database software upgrades
	//
	// 	- **DatabaseHardwareMaintenance**: hardware maintenance and upgrades
	//
	// 	- **DatabaseStorageUpgrading**: database storage upgrades
	//
	// 	- **DatabaseProxyUpgrading**: minor version upgrades of the proxy
	//
	// 	- **all**: queries the details of the pending events of all preceding types.
	//
	// > If the `Region` parameter is set to **all**, the `TaskType` parameter must be set to **all**.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribePendingMaintenanceActionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePendingMaintenanceActionRequest) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionRequest) GetIsHistory() *int32 {
	return s.IsHistory
}

func (s *DescribePendingMaintenanceActionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePendingMaintenanceActionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePendingMaintenanceActionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePendingMaintenanceActionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePendingMaintenanceActionRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribePendingMaintenanceActionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePendingMaintenanceActionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePendingMaintenanceActionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePendingMaintenanceActionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePendingMaintenanceActionRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribePendingMaintenanceActionRequest) SetIsHistory(v int32) *DescribePendingMaintenanceActionRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetOwnerAccount(v string) *DescribePendingMaintenanceActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetOwnerId(v int64) *DescribePendingMaintenanceActionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetPageNumber(v int32) *DescribePendingMaintenanceActionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetPageSize(v int32) *DescribePendingMaintenanceActionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetRegion(v string) *DescribePendingMaintenanceActionRequest {
	s.Region = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetResourceGroupId(v string) *DescribePendingMaintenanceActionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetResourceOwnerAccount(v string) *DescribePendingMaintenanceActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetResourceOwnerId(v int64) *DescribePendingMaintenanceActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetSecurityToken(v string) *DescribePendingMaintenanceActionRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetTaskType(v string) *DescribePendingMaintenanceActionRequest {
	s.TaskType = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) Validate() error {
	return dara.Validate(s)
}
