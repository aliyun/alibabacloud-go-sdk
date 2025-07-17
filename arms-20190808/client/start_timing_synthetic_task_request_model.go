// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTimingSyntheticTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *StartTimingSyntheticTaskRequest
	GetRegionId() *string
	SetTaskIds(v []*string) *StartTimingSyntheticTaskRequest
	GetTaskIds() []*string
}

type StartTimingSyntheticTaskRequest struct {
	// The region ID. Default value: cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task IDs.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s StartTimingSyntheticTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTimingSyntheticTaskRequest) GoString() string {
	return s.String()
}

func (s *StartTimingSyntheticTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartTimingSyntheticTaskRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *StartTimingSyntheticTaskRequest) SetRegionId(v string) *StartTimingSyntheticTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StartTimingSyntheticTaskRequest) SetTaskIds(v []*string) *StartTimingSyntheticTaskRequest {
	s.TaskIds = v
	return s
}

func (s *StartTimingSyntheticTaskRequest) Validate() error {
	return dara.Validate(s)
}
