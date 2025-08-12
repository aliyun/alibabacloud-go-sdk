// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupDynamicRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DescribeMonitorGroupDynamicRulesRequest
	GetGroupId() *int64
	SetRegionId(v string) *DescribeMonitorGroupDynamicRulesRequest
	GetRegionId() *string
}

type DescribeMonitorGroupDynamicRulesRequest struct {
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

func (s DescribeMonitorGroupDynamicRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeMonitorGroupDynamicRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitorGroupDynamicRulesRequest) SetGroupId(v int64) *DescribeMonitorGroupDynamicRulesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesRequest) SetRegionId(v string) *DescribeMonitorGroupDynamicRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesRequest) Validate() error {
	return dara.Validate(s)
}
