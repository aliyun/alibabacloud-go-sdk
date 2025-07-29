// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEssayCorrectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetEssayCorrectionTaskRequest
	GetTaskId() *string
}

type GetEssayCorrectionTaskRequest struct {
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetEssayCorrectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEssayCorrectionTaskRequest) GoString() string {
	return s.String()
}

func (s *GetEssayCorrectionTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetEssayCorrectionTaskRequest) SetTaskId(v string) *GetEssayCorrectionTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetEssayCorrectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
