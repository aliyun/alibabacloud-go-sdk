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
	// The paging Settings for the current page number in a paged query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp in milliseconds. Operation logs generated before this time are queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1640880000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Anti-DDoS Origin instance to query.
	//
	// > You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation type. Valid values:
	//
	// - **3**: attach an IP address
	//
	// - **4**: unbind an IP address
	//
	// - **5**: decrease the quota
	//
	// - **6**: deactivate blackhole filtering
	//
	// - **7**: Reset quota on blackhole filtering deactivation
	//
	// - **8**: recover burstable protection
	//
	// - **9**: add an Asset group
	//
	// - **10**: delete an Asset group
	//
	// - **11**: enable daily 95th percentile bandwidth
	//
	// - **12**: enable monthly 95th percentile bandwidth
	//
	// - **13**: periodically switch between daily and monthly 95th percentile bandwidth patterns
	//
	// - **14**: shutdown daily 95th percentile bandwidth
	//
	// - **15**: shutdown monthly 95th percentile bandwidth
	//
	// - **16**: shutdown burstable clean bandwidth due to overdue payments
	//
	// - **17**: shutdown burstable clean bandwidth when they expire
	//
	// example:
	//
	// 3
	OpAction *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// The field by which to sort the operation logs. Set the value to **opdate**, which indicates that the operation logs are sorted by operation time.
	//
	// example:
	//
	// opdate
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The sort order of the operation logs. Valid values:
	//
	// - **ASC**: ascending order
	//
	// - **DESC**: descending order
	//
	// Default value: **DESC**.
	//
	// example:
	//
	// ASC
	OrderDir *string `json:"OrderDir,omitempty" xml:"OrderDir,omitempty"`
	// The paging Settings for the number of Operation logs per page in a paged query. Maximum value: 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Anti-DDoS Origin instance.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the IDs of all regions supported by Anti-DDoS Origin.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the default resource group is used.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp in milliseconds. Operation logs generated after this time are queried.
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
