// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetTaskStatusByIdRequest
	GetTaskId() *string
}

type GetTaskStatusByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 95906135
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskStatusByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusByIdRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusByIdRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskStatusByIdRequest) SetTaskId(v string) *GetTaskStatusByIdRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskStatusByIdRequest) Validate() error {
	return dara.Validate(s)
}
