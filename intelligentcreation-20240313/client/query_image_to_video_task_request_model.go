// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryImageToVideoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *QueryImageToVideoTaskRequest
	GetTaskId() *string
}

type QueryImageToVideoTaskRequest struct {
	// example:
	//
	// 868125994191405056
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryImageToVideoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryImageToVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryImageToVideoTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryImageToVideoTaskRequest) SetTaskId(v string) *QueryImageToVideoTaskRequest {
	s.TaskId = &v
	return s
}

func (s *QueryImageToVideoTaskRequest) Validate() error {
	return dara.Validate(s)
}
