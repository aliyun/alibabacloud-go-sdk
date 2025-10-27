// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMaintenanceActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsHistory(v int32) *DescribeMaintenanceActionRequest
	GetIsHistory() *int32
	SetOwnerAccount(v string) *DescribeMaintenanceActionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeMaintenanceActionRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeMaintenanceActionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMaintenanceActionRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribeMaintenanceActionRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeMaintenanceActionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeMaintenanceActionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMaintenanceActionRequest
	GetResourceOwnerId() *int64
	SetTaskType(v string) *DescribeMaintenanceActionRequest
	GetTaskType() *string
}

type DescribeMaintenanceActionRequest struct {
	// Specifies whether to return the information about pending or historical O\\&M events. Valid values:
	//
	// 	- **0**: returns the information about pending O\\&M event.
	//
	// 	- **1**: returns the information about historical O\\&M event.
	//
	// If you do not specify this parameter, the information about pending O\\&M event are returned.
	//
	// example:
	//
	// 1
	IsHistory    *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30**, **50**, and **100**. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. Valid values:
	//
	// 	- The ID of the region where the O\\&M event occurs. Example: `cn-hangzhou`. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// 	- You can also set Region to `all` to query the O\\&M events in all regions. If you set `Region` to `all`, you must set `TaskType` to `all`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the region where the O\\&M event occurs.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the O\\&M event. Valid values:
	//
	// 	- **rds_apsaradb_upgrade**: database software upgrades.
	//
	// 	- **all**: all the O\\&M events in all regions within the current account. If you set `Region` to `all`, you must set `TaskType` to `all`.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds_apsaradb_upgrade
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeMaintenanceActionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaintenanceActionRequest) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionRequest) GetIsHistory() *int32 {
	return s.IsHistory
}

func (s *DescribeMaintenanceActionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeMaintenanceActionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMaintenanceActionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMaintenanceActionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMaintenanceActionRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeMaintenanceActionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMaintenanceActionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMaintenanceActionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMaintenanceActionRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeMaintenanceActionRequest) SetIsHistory(v int32) *DescribeMaintenanceActionRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetOwnerAccount(v string) *DescribeMaintenanceActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetOwnerId(v int64) *DescribeMaintenanceActionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetPageNumber(v int32) *DescribeMaintenanceActionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetPageSize(v int32) *DescribeMaintenanceActionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetRegion(v string) *DescribeMaintenanceActionRequest {
	s.Region = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetRegionId(v string) *DescribeMaintenanceActionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetResourceOwnerAccount(v string) *DescribeMaintenanceActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetResourceOwnerId(v int64) *DescribeMaintenanceActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) SetTaskType(v string) *DescribeMaintenanceActionRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeMaintenanceActionRequest) Validate() error {
	return dara.Validate(s)
}
