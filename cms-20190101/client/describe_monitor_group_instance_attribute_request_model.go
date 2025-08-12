// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeMonitorGroupInstanceAttributeRequest
	GetCategory() *string
	SetGroupId(v int64) *DescribeMonitorGroupInstanceAttributeRequest
	GetGroupId() *int64
	SetInstanceIds(v string) *DescribeMonitorGroupInstanceAttributeRequest
	GetInstanceIds() *string
	SetKeyword(v string) *DescribeMonitorGroupInstanceAttributeRequest
	GetKeyword() *string
	SetPageNumber(v int32) *DescribeMonitorGroupInstanceAttributeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMonitorGroupInstanceAttributeRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMonitorGroupInstanceAttributeRequest
	GetRegionId() *string
	SetTotal(v bool) *DescribeMonitorGroupInstanceAttributeRequest
	GetTotal() *bool
}

type DescribeMonitorGroupInstanceAttributeRequest struct {
	// The abbreviation of the cloud service name.
	//
	// For more information about how to obtain the abbreviation of a cloud service name, see `metricCategory` in the response parameter `Labels` of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/2513265.html) operation.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The resource ID. Separate multiple resource IDs with commas (,). You can query the details about a maximum of 20 resources at a time.
	//
	// example:
	//
	// i-m5e0k0bexac8tykr****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The keyword that is used to search for resources.
	//
	// example:
	//
	// portal
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// Valid values: 1 to 1000000000.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 1000000000.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to return the total number of resources in the specified application group. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	Total *bool `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) GetTotal() *bool {
	return s.Total
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetCategory(v string) *DescribeMonitorGroupInstanceAttributeRequest {
	s.Category = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetGroupId(v int64) *DescribeMonitorGroupInstanceAttributeRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetInstanceIds(v string) *DescribeMonitorGroupInstanceAttributeRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetKeyword(v string) *DescribeMonitorGroupInstanceAttributeRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetPageNumber(v int32) *DescribeMonitorGroupInstanceAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetPageSize(v int32) *DescribeMonitorGroupInstanceAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetRegionId(v string) *DescribeMonitorGroupInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetTotal(v bool) *DescribeMonitorGroupInstanceAttributeRequest {
	s.Total = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
