// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Task) *PutTaskStatusRequest
	GetBody() *Task
	SetForce(v bool) *PutTaskStatusRequest
	GetForce() *bool
}

type PutTaskStatusRequest struct {
	Body *Task `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s PutTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s PutTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *PutTaskStatusRequest) GetBody() *Task {
	return s.Body
}

func (s *PutTaskStatusRequest) GetForce() *bool {
	return s.Force
}

func (s *PutTaskStatusRequest) SetBody(v *Task) *PutTaskStatusRequest {
	s.Body = v
	return s
}

func (s *PutTaskStatusRequest) SetForce(v bool) *PutTaskStatusRequest {
	s.Force = &v
	return s
}

func (s *PutTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
