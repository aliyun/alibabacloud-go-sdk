// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioAsrTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetAudioAsrTaskStatusRequest
	GetTaskId() *string
}

type GetAudioAsrTaskStatusRequest struct {
	// This parameter is required.
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetAudioAsrTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAudioAsrTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAudioAsrTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAudioAsrTaskStatusRequest) SetTaskId(v string) *GetAudioAsrTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetAudioAsrTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
