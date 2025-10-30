// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSwimmingLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *DescribeSwimmingLaneRequest
	GetGroupId() *int64
	SetLaneId(v int64) *DescribeSwimmingLaneRequest
	GetLaneId() *int64
	SetNamespaceId(v string) *DescribeSwimmingLaneRequest
	GetNamespaceId() *string
}

type DescribeSwimmingLaneRequest struct {
	// The ID of the lane group.
	//
	// example:
	//
	// 2074
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the lane.
	//
	// example:
	//
	// 9637
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// The ID of a namespace.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DescribeSwimmingLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSwimmingLaneRequest) GoString() string {
	return s.String()
}

func (s *DescribeSwimmingLaneRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeSwimmingLaneRequest) GetLaneId() *int64 {
	return s.LaneId
}

func (s *DescribeSwimmingLaneRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeSwimmingLaneRequest) SetGroupId(v int64) *DescribeSwimmingLaneRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSwimmingLaneRequest) SetLaneId(v int64) *DescribeSwimmingLaneRequest {
	s.LaneId = &v
	return s
}

func (s *DescribeSwimmingLaneRequest) SetNamespaceId(v string) *DescribeSwimmingLaneRequest {
	s.NamespaceId = &v
	return s
}

func (s *DescribeSwimmingLaneRequest) Validate() error {
	return dara.Validate(s)
}
