// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimmingLaneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLaneId(v int64) *DeleteSwimmingLaneRequest
	GetLaneId() *int64
}

type DeleteSwimmingLaneRequest struct {
	// The ID of the lane.
	//
	// This parameter is required.
	//
	// example:
	//
	// 241
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
}

func (s DeleteSwimmingLaneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimmingLaneRequest) GoString() string {
	return s.String()
}

func (s *DeleteSwimmingLaneRequest) GetLaneId() *int64 {
	return s.LaneId
}

func (s *DeleteSwimmingLaneRequest) SetLaneId(v int64) *DeleteSwimmingLaneRequest {
	s.LaneId = &v
	return s
}

func (s *DeleteSwimmingLaneRequest) Validate() error {
	return dara.Validate(s)
}
