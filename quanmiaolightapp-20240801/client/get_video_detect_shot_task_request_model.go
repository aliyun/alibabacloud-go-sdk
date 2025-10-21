// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoDetectShotTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetVideoDetectShotTaskRequest
	GetTaskId() *string
}

type GetVideoDetectShotTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetVideoDetectShotTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotTaskRequest) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoDetectShotTaskRequest) SetTaskId(v string) *GetVideoDetectShotTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetVideoDetectShotTaskRequest) Validate() error {
	return dara.Validate(s)
}
