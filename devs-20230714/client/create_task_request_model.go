// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Task) *CreateTaskRequest
	GetBody() *Task
}

type CreateTaskRequest struct {
	Body *Task `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) GetBody() *Task {
	return s.Body
}

func (s *CreateTaskRequest) SetBody(v *Task) *CreateTaskRequest {
	s.Body = v
	return s
}

func (s *CreateTaskRequest) Validate() error {
	return dara.Validate(s)
}
