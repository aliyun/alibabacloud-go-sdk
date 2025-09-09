// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTestTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetServiceTestTaskRequest
	GetRegionId() *string
	SetTaskId(v string) *GetServiceTestTaskRequest
	GetTaskId() *string
}

type GetServiceTestTaskRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// stt-xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetServiceTestTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTestTaskRequest) GoString() string {
	return s.String()
}

func (s *GetServiceTestTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceTestTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetServiceTestTaskRequest) SetRegionId(v string) *GetServiceTestTaskRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceTestTaskRequest) SetTaskId(v string) *GetServiceTestTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetServiceTestTaskRequest) Validate() error {
	return dara.Validate(s)
}
