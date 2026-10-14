// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotfixTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetHotfixTaskRequest
	GetTaskId() *string
}

type GetHotfixTaskRequest struct {
	// The ID of the asynchronous task.
	//
	// This parameter is required.
	//
	// example:
	//
	// hotfix-da3dcc0b-5019-4b59-b12f-2e500da9083c
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetHotfixTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotfixTaskRequest) GoString() string {
	return s.String()
}

func (s *GetHotfixTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetHotfixTaskRequest) SetTaskId(v string) *GetHotfixTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetHotfixTaskRequest) Validate() error {
	return dara.Validate(s)
}
