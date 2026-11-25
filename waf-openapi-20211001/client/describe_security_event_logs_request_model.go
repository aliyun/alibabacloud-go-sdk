// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *DescribeSecurityEventLogsRequestFilter) *DescribeSecurityEventLogsRequest
	GetFilter() *DescribeSecurityEventLogsRequestFilter
	SetInstanceId(v string) *DescribeSecurityEventLogsRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeSecurityEventLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSecurityEventLogsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeSecurityEventLogsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventLogsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeSecurityEventLogsRequest struct {
	// The query filter conditions. Multiple filter conditions are connected by the logical AND operator.
	//
	// This parameter is required.
	Filter *DescribeSecurityEventLogsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number to return during paging in a paged query. Default value: **1**, which indicates the first page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paging in a paged query. Maximum value: **100**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeSecurityEventLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventLogsRequest) GetFilter() *DescribeSecurityEventLogsRequestFilter {
	return s.Filter
}

func (s *DescribeSecurityEventLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityEventLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSecurityEventLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSecurityEventLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityEventLogsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSecurityEventLogsRequest) SetFilter(v *DescribeSecurityEventLogsRequestFilter) *DescribeSecurityEventLogsRequest {
	s.Filter = v
	return s
}

func (s *DescribeSecurityEventLogsRequest) SetInstanceId(v string) *DescribeSecurityEventLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityEventLogsRequest) SetPageNumber(v int64) *DescribeSecurityEventLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecurityEventLogsRequest) SetPageSize(v int64) *DescribeSecurityEventLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityEventLogsRequest) SetRegionId(v string) *DescribeSecurityEventLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityEventLogsRequest) SetResourceManagerResourceGroupId(v string) *DescribeSecurityEventLogsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSecurityEventLogsRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityEventLogsRequestFilter struct {
	// The list of filter conditions. Each node describes a filter condition.
	Conditions []*DescribeSecurityEventLogsRequestFilterConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The time range for the query.
	//
	// This parameter is required.
	DateRange *DescribeSecurityEventLogsRequestFilterDateRange `json:"DateRange,omitempty" xml:"DateRange,omitempty" type:"Struct"`
}

func (s DescribeSecurityEventLogsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventLogsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventLogsRequestFilter) GetConditions() []*DescribeSecurityEventLogsRequestFilterConditions {
	return s.Conditions
}

func (s *DescribeSecurityEventLogsRequestFilter) GetDateRange() *DescribeSecurityEventLogsRequestFilterDateRange {
	return s.DateRange
}

func (s *DescribeSecurityEventLogsRequestFilter) SetConditions(v []*DescribeSecurityEventLogsRequestFilterConditions) *DescribeSecurityEventLogsRequestFilter {
	s.Conditions = v
	return s
}

func (s *DescribeSecurityEventLogsRequestFilter) SetDateRange(v *DescribeSecurityEventLogsRequestFilterDateRange) *DescribeSecurityEventLogsRequestFilter {
	s.DateRange = v
	return s
}

func (s *DescribeSecurityEventLogsRequestFilter) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DateRange != nil {
		if err := s.DateRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityEventLogsRequestFilterConditions struct {
	// The name of the field on which the filter operation is performed. This operation supports all fields.
	//
	// example:
	//
	// matched_host
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The operator.
	//
	// example:
	//
	// eq
	OpValue *string `json:"OpValue,omitempty" xml:"OpValue,omitempty"`
	// The filter values.
	//
	// example:
	//
	// test.waf-top
	Values interface{} `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeSecurityEventLogsRequestFilterConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventLogsRequestFilterConditions) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventLogsRequestFilterConditions) GetKey() *string {
	return s.Key
}

func (s *DescribeSecurityEventLogsRequestFilterConditions) GetOpValue() *string {
	return s.OpValue
}

func (s *DescribeSecurityEventLogsRequestFilterConditions) GetValues() interface{} {
	return s.Values
}

func (s *DescribeSecurityEventLogsRequestFilterConditions) SetKey(v string) *DescribeSecurityEventLogsRequestFilterConditions {
	s.Key = &v
	return s
}

func (s *DescribeSecurityEventLogsRequestFilterConditions) SetOpValue(v string) *DescribeSecurityEventLogsRequestFilterConditions {
	s.OpValue = &v
	return s
}

func (s *DescribeSecurityEventLogsRequestFilterConditions) SetValues(v interface{}) *DescribeSecurityEventLogsRequestFilterConditions {
	s.Values = v
	return s
}

func (s *DescribeSecurityEventLogsRequestFilterConditions) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventLogsRequestFilterDateRange struct {
	// The end time of the query. The end time is represented by a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888600
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The start time of the query. The data query range cannot exceed the past 30 days. The start time is represented by a UNIX timestamp in seconds.
	//
	//
	// > The start time must be later than the current time minus 30 days.
	//
	// > -
	//
	// This parameter is required.
	//
	// example:
	//
	// 1713888000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeSecurityEventLogsRequestFilterDateRange) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventLogsRequestFilterDateRange) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventLogsRequestFilterDateRange) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeSecurityEventLogsRequestFilterDateRange) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeSecurityEventLogsRequestFilterDateRange) SetEndDate(v int64) *DescribeSecurityEventLogsRequestFilterDateRange {
	s.EndDate = &v
	return s
}

func (s *DescribeSecurityEventLogsRequestFilterDateRange) SetStartDate(v int64) *DescribeSecurityEventLogsRequestFilterDateRange {
	s.StartDate = &v
	return s
}

func (s *DescribeSecurityEventLogsRequestFilterDateRange) Validate() error {
	return dara.Validate(s)
}
