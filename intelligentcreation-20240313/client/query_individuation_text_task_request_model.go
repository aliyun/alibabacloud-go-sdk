// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIndividuationTextTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *QueryIndividuationTextTaskRequest
	GetTaskId() *string
}

type QueryIndividuationTextTaskRequest struct {
	// example:
	//
	// 829682927337963520
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryIndividuationTextTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryIndividuationTextTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryIndividuationTextTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryIndividuationTextTaskRequest) SetTaskId(v string) *QueryIndividuationTextTaskRequest {
	s.TaskId = &v
	return s
}

func (s *QueryIndividuationTextTaskRequest) Validate() error {
	return dara.Validate(s)
}
