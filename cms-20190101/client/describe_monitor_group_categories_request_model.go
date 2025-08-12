// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupCategoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DescribeMonitorGroupCategoriesRequest
	GetGroupId() *int64
	SetRegionId(v string) *DescribeMonitorGroupCategoriesRequest
	GetRegionId() *string
}

type DescribeMonitorGroupCategoriesRequest struct {
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId  *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitorGroupCategoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeMonitorGroupCategoriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitorGroupCategoriesRequest) SetGroupId(v int64) *DescribeMonitorGroupCategoriesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesRequest) SetRegionId(v string) *DescribeMonitorGroupCategoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesRequest) Validate() error {
	return dara.Validate(s)
}
