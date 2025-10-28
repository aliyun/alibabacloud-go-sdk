// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *ListSwimmingLaneGroupRequest
	GetGroupId() *int64
	SetLogicalRegionId(v string) *ListSwimmingLaneGroupRequest
	GetLogicalRegionId() *string
}

type ListSwimmingLaneGroupRequest struct {
	// The ID of the lane group.
	//
	// example:
	//
	// 0
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the namespace.
	//
	// The ID of a custom namespace is in the region ID:namespace identifier format. Example: cn-beijing:test.\\
	//
	// The ID of the default namespace is in the region ID format. Example: cn-beijing.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen:publish
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
}

func (s ListSwimmingLaneGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListSwimmingLaneGroupRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *ListSwimmingLaneGroupRequest) SetGroupId(v int64) *ListSwimmingLaneGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ListSwimmingLaneGroupRequest) SetLogicalRegionId(v string) *ListSwimmingLaneGroupRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *ListSwimmingLaneGroupRequest) Validate() error {
	return dara.Validate(s)
}
