// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateModelServiceInput) *UpdateModelServiceRequest
	GetBody() *UpdateModelServiceInput
}

type UpdateModelServiceRequest struct {
	Body *UpdateModelServiceInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelServiceRequest) GetBody() *UpdateModelServiceInput {
	return s.Body
}

func (s *UpdateModelServiceRequest) SetBody(v *UpdateModelServiceInput) *UpdateModelServiceRequest {
	s.Body = v
	return s
}

func (s *UpdateModelServiceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
