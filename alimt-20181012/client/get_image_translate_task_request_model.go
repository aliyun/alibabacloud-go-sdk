// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetImageTranslateTaskRequest
	GetTaskId() *string
}

type GetImageTranslateTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetImageTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageTranslateTaskRequest) SetTaskId(v string) *GetImageTranslateTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetImageTranslateTaskRequest) Validate() error {
	return dara.Validate(s)
}
