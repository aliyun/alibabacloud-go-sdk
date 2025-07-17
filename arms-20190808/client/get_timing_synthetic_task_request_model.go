// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTimingSyntheticTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetTimingSyntheticTaskRequest
	GetRegionId() *string
	SetTaskId(v string) *GetTimingSyntheticTaskRequest
	GetTaskId() *string
}

type GetTimingSyntheticTaskRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the synthetic monitoring task.
	//
	// example:
	//
	// 856566a9cb2a4cafa05aa95ed0ec8f21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTimingSyntheticTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTimingSyntheticTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTimingSyntheticTaskRequest) SetRegionId(v string) *GetTimingSyntheticTaskRequest {
	s.RegionId = &v
	return s
}

func (s *GetTimingSyntheticTaskRequest) SetTaskId(v string) *GetTimingSyntheticTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetTimingSyntheticTaskRequest) Validate() error {
	return dara.Validate(s)
}
