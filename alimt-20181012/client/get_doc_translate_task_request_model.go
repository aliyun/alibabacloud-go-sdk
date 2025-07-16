// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetDocTranslateTaskRequest
	GetTaskId() *string
}

type GetDocTranslateTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0586df512c8b4bb382d7d9a6a01b5854
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDocTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *GetDocTranslateTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDocTranslateTaskRequest) SetTaskId(v string) *GetDocTranslateTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetDocTranslateTaskRequest) Validate() error {
	return dara.Validate(s)
}
