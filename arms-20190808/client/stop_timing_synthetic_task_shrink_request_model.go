// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTimingSyntheticTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *StopTimingSyntheticTaskShrinkRequest
	GetRegionId() *string
	SetTaskIdsShrink(v string) *StopTimingSyntheticTaskShrinkRequest
	GetTaskIdsShrink() *string
}

type StopTimingSyntheticTaskShrinkRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task IDs.
	//
	// This parameter is required.
	TaskIdsShrink *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s StopTimingSyntheticTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTimingSyntheticTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopTimingSyntheticTaskShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopTimingSyntheticTaskShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *StopTimingSyntheticTaskShrinkRequest) SetRegionId(v string) *StopTimingSyntheticTaskShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *StopTimingSyntheticTaskShrinkRequest) SetTaskIdsShrink(v string) *StopTimingSyntheticTaskShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *StopTimingSyntheticTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
