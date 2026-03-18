// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeOpEntitiesRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *DescribeOpEntitiesRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeOpEntitiesRequest
	GetInstanceId() *string
	SetOpAction(v int32) *DescribeOpEntitiesRequest
	GetOpAction() *int32
	SetOrderBy(v string) *DescribeOpEntitiesRequest
	GetOrderBy() *string
	SetOrderDir(v string) *DescribeOpEntitiesRequest
	GetOrderDir() *string
	SetPageSize(v int32) *DescribeOpEntitiesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeOpEntitiesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeOpEntitiesRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeOpEntitiesRequest
	GetStartTime() *int64
}

type DescribeOpEntitiesRequest struct {
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time. Operation logs that were generated before this time are queried.***	- The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1640880000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance to query.
	//
	// > You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all instances.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the operation. Valid values:
	//
	// 	- **3**: Add an IP address to the instance.
	//
	// 	- **4**: Remove an IP address from the instance.
	//
	// 	- **5**: Downgrade the instance.
	//
	// 	- **6**: Deactivate blackhole filtering.
	//
	// 	- **7**: Reset the number of times that you can deactivate blackhole filtering.
	//
	// 	- **8**: Restore the mitigation capability.
	//
	// 	- **9**: Add an asset group.
	//
	// 	- **10**: Remove an asset group.
	//
	// 	- **11**: Enable the metering method of daily 95th percentile for the burstable clean bandwidth feature.
	//
	// 	- **12**: Enable the metering method of monthly 95th percentile for the burstable clean bandwidth feature.
	//
	// 	- **13**: Periodically switch between the metering methods of daily 95th percentile and monthly 95th percentile for the burstable clean bandwidth feature.
	//
	// 	- **14**: Disable the metering method of daily 95th percentile for the burstable clean bandwidth feature.
	//
	// 	- **15**: Disable the metering method of monthly 95th percentile for the burstable clean bandwidth feature.
	//
	// 	- **16**: Disable burstable clean bandwidth due to overdue payments.
	//
	// 	- **17**: Disable burstable clean bandwidth due to instance expiration.
	//
	// example:
	//
	// 3
	OpAction *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// The sorting method of operation logs. Set the value to **opdate**, which indicates sorting based on the operation time.
	//
	// example:
	//
	// opdate
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The sort order of operation logs. Valid values:
	//
	// 	- **ASC**: the ascending order.
	//
	// 	- **DESC**: the descending order.
	//
	// Default value: **DESC**.
	//
	// example:
	//
	// ASC
	OrderDir *string `json:"OrderDir,omitempty" xml:"OrderDir,omitempty"`
	// The number of entries per page. Maximum value: 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the instance resides.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time. Operation logs that were generated after this time are queried.***	- The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1609430400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOpEntitiesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeOpEntitiesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeOpEntitiesRequest) GetOpAction() *int32 {
	return s.OpAction
}

func (s *DescribeOpEntitiesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeOpEntitiesRequest) GetOrderDir() *string {
	return s.OrderDir
}

func (s *DescribeOpEntitiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOpEntitiesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpEntitiesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeOpEntitiesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeOpEntitiesRequest) SetCurrentPage(v int32) *DescribeOpEntitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetInstanceId(v string) *DescribeOpEntitiesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOpAction(v int32) *DescribeOpEntitiesRequest {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOrderBy(v string) *DescribeOpEntitiesRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOrderDir(v string) *DescribeOpEntitiesRequest {
	s.OrderDir = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int32) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetRegionId(v string) *DescribeOpEntitiesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetResourceGroupId(v string) *DescribeOpEntitiesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) Validate() error {
	return dara.Validate(s)
}
