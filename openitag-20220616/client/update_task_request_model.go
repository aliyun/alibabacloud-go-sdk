// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateTaskDTO) *UpdateTaskRequest
	GetBody() *UpdateTaskDTO
}

type UpdateTaskRequest struct {
	// Task Status
	//
	// This parameter is required.
	Body *UpdateTaskDTO `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequest) GetBody() *UpdateTaskDTO {
	return s.Body
}

func (s *UpdateTaskRequest) SetBody(v *UpdateTaskDTO) *UpdateTaskRequest {
	s.Body = v
	return s
}

func (s *UpdateTaskRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
