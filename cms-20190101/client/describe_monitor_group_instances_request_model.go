// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeMonitorGroupInstancesRequest
	GetCategory() *string
	SetGroupId(v int64) *DescribeMonitorGroupInstancesRequest
	GetGroupId() *int64
	SetInstanceIds(v string) *DescribeMonitorGroupInstancesRequest
	GetInstanceIds() *string
	SetKeyword(v string) *DescribeMonitorGroupInstancesRequest
	GetKeyword() *string
	SetPageNumber(v int32) *DescribeMonitorGroupInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMonitorGroupInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMonitorGroupInstancesRequest
	GetRegionId() *string
}

type DescribeMonitorGroupInstancesRequest struct {
	// The abbreviation of the cloud service name. Valid values of N: 1 to 200.
	//
	// >  For more information about how to obtain the abbreviation of a cloud service name, see `metricCategory` in the response parameter `Labels` of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
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
	// 12345
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID. You can query multiple instances by specifying multiple IDs.
	//
	// example:
	//
	// i-x1234568
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The keyword used to search for instances. Fuzzy search based on instance names is supported.
	//
	// example:
	//
	// s1
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitorGroupInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstancesRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeMonitorGroupInstancesRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeMonitorGroupInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeMonitorGroupInstancesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeMonitorGroupInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMonitorGroupInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMonitorGroupInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitorGroupInstancesRequest) SetCategory(v string) *DescribeMonitorGroupInstancesRequest {
	s.Category = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) SetGroupId(v int64) *DescribeMonitorGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) SetInstanceIds(v string) *DescribeMonitorGroupInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) SetKeyword(v string) *DescribeMonitorGroupInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) SetPageNumber(v int32) *DescribeMonitorGroupInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) SetPageSize(v int32) *DescribeMonitorGroupInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) SetRegionId(v string) *DescribeMonitorGroupInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) Validate() error {
	return dara.Validate(s)
}
