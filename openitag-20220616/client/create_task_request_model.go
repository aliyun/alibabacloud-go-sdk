// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateTaskDetail) *CreateTaskRequest
	GetBody() *CreateTaskDetail
}

type CreateTaskRequest struct {
	// Job details
	//
	// This parameter is required.
	Body *CreateTaskDetail `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) GetBody() *CreateTaskDetail {
	return s.Body
}

func (s *CreateTaskRequest) SetBody(v *CreateTaskDetail) *CreateTaskRequest {
	s.Body = v
	return s
}

func (s *CreateTaskRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
