// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunScriptRefineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *RunScriptRefineRequest
	GetTaskId() *string
}

type RunScriptRefineRequest struct {
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RunScriptRefineRequest) String() string {
	return dara.Prettify(s)
}

func (s RunScriptRefineRequest) GoString() string {
	return s.String()
}

func (s *RunScriptRefineRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunScriptRefineRequest) SetTaskId(v string) *RunScriptRefineRequest {
	s.TaskId = &v
	return s
}

func (s *RunScriptRefineRequest) Validate() error {
	return dara.Validate(s)
}
