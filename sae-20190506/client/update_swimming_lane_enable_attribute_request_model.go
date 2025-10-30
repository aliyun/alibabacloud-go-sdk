// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimmingLaneEnableAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *UpdateSwimmingLaneEnableAttributeRequest
	GetEnable() *bool
	SetGroupId(v int64) *UpdateSwimmingLaneEnableAttributeRequest
	GetGroupId() *int64
	SetLaneId(v int64) *UpdateSwimmingLaneEnableAttributeRequest
	GetLaneId() *int64
	SetNamespaceId(v string) *UpdateSwimmingLaneEnableAttributeRequest
	GetNamespaceId() *string
}

type UpdateSwimmingLaneEnableAttributeRequest struct {
	// Lane status:
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
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

func (s UpdateSwimmingLaneEnableAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneEnableAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneEnableAttributeRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateSwimmingLaneEnableAttributeRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateSwimmingLaneEnableAttributeRequest) GetLaneId() *int64 {
	return s.LaneId
}

func (s *UpdateSwimmingLaneEnableAttributeRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateSwimmingLaneEnableAttributeRequest) SetEnable(v bool) *UpdateSwimmingLaneEnableAttributeRequest {
	s.Enable = &v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeRequest) SetGroupId(v int64) *UpdateSwimmingLaneEnableAttributeRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeRequest) SetLaneId(v int64) *UpdateSwimmingLaneEnableAttributeRequest {
	s.LaneId = &v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeRequest) SetNamespaceId(v string) *UpdateSwimmingLaneEnableAttributeRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateSwimmingLaneEnableAttributeRequest) Validate() error {
	return dara.Validate(s)
}
