// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTimingSyntheticTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *StartTimingSyntheticTaskShrinkRequest
	GetRegionId() *string
	SetTaskIdsShrink(v string) *StartTimingSyntheticTaskShrinkRequest
	GetTaskIdsShrink() *string
}

type StartTimingSyntheticTaskShrinkRequest struct {
	// The region ID. Default value: cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task IDs.
	TaskIdsShrink *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s StartTimingSyntheticTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTimingSyntheticTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartTimingSyntheticTaskShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartTimingSyntheticTaskShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *StartTimingSyntheticTaskShrinkRequest) SetRegionId(v string) *StartTimingSyntheticTaskShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *StartTimingSyntheticTaskShrinkRequest) SetTaskIdsShrink(v string) *StartTimingSyntheticTaskShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *StartTimingSyntheticTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
