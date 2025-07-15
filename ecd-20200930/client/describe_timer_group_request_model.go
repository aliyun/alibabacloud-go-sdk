// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTimerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeTimerGroupRequest
	GetGroupId() *string
	SetRegionId(v string) *DescribeTimerGroupRequest
	GetRegionId() *string
}

type DescribeTimerGroupRequest struct {
	// The ID of the configuration group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg-hs3i1w39o68ma****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the region. Set the value to `cn-shanghai`.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTimerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimerGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeTimerGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeTimerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTimerGroupRequest) SetGroupId(v string) *DescribeTimerGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeTimerGroupRequest) SetRegionId(v string) *DescribeTimerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTimerGroupRequest) Validate() error {
	return dara.Validate(s)
}
