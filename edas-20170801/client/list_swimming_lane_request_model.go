// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *ListSwimmingLaneRequest
	GetGroupId() *int64
}

type ListSwimmingLaneRequest struct {
	// The ID of the lane group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ListSwimmingLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneRequest) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListSwimmingLaneRequest) SetGroupId(v int64) *ListSwimmingLaneRequest {
	s.GroupId = &v
	return s
}

func (s *ListSwimmingLaneRequest) Validate() error {
	return dara.Validate(s)
}
