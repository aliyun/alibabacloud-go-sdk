// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateModelServiceInput) *CreateModelServiceRequest
	GetBody() *CreateModelServiceInput
}

type CreateModelServiceRequest struct {
	Body *CreateModelServiceInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateModelServiceRequest) GetBody() *CreateModelServiceInput {
	return s.Body
}

func (s *CreateModelServiceRequest) SetBody(v *CreateModelServiceInput) *CreateModelServiceRequest {
	s.Body = v
	return s
}

func (s *CreateModelServiceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
