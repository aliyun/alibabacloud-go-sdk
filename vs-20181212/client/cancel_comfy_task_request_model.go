// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelComfyTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *CancelComfyTaskRequest
	GetTaskId() *string
}

type CancelComfyTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6c8234f4-d1e1-4cea-b08b-7926fbdea144
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelComfyTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelComfyTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelComfyTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelComfyTaskRequest) SetTaskId(v string) *CancelComfyTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelComfyTaskRequest) Validate() error {
	return dara.Validate(s)
}
