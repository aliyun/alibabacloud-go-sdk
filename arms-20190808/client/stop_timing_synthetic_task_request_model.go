// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTimingSyntheticTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *StopTimingSyntheticTaskRequest
	GetRegionId() *string
	SetTaskIds(v []*string) *StopTimingSyntheticTaskRequest
	GetTaskIds() []*string
}

type StopTimingSyntheticTaskRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task IDs.
	//
	// This parameter is required.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s StopTimingSyntheticTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTimingSyntheticTaskRequest) GoString() string {
	return s.String()
}

func (s *StopTimingSyntheticTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopTimingSyntheticTaskRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *StopTimingSyntheticTaskRequest) SetRegionId(v string) *StopTimingSyntheticTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StopTimingSyntheticTaskRequest) SetTaskIds(v []*string) *StopTimingSyntheticTaskRequest {
	s.TaskIds = v
	return s
}

func (s *StopTimingSyntheticTaskRequest) Validate() error {
	return dara.Validate(s)
}
