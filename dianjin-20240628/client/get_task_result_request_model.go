// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetTaskResultRequest
	GetTaskId() *string
}

type GetTaskResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 17071319
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskResultRequest) SetTaskId(v string) *GetTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
