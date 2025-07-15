// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcAsrTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *StopRtcAsrTaskRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StopRtcAsrTaskRequest
	GetRegionId() *string
	SetTaskId(v string) *StopRtcAsrTaskRequest
	GetTaskId() *string
}

type StopRtcAsrTaskRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// asr-d794cc89-a63e-4d08-8b44-242a6597****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopRtcAsrTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRtcAsrTaskRequest) GoString() string {
	return s.String()
}

func (s *StopRtcAsrTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopRtcAsrTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopRtcAsrTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopRtcAsrTaskRequest) SetOwnerId(v int64) *StopRtcAsrTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopRtcAsrTaskRequest) SetRegionId(v string) *StopRtcAsrTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StopRtcAsrTaskRequest) SetTaskId(v string) *StopRtcAsrTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopRtcAsrTaskRequest) Validate() error {
	return dara.Validate(s)
}
