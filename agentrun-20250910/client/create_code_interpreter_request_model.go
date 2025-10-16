// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeInterpreterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateCodeInterpreterInput) *CreateCodeInterpreterRequest
	GetBody() *CreateCodeInterpreterInput
}

type CreateCodeInterpreterRequest struct {
	// This parameter is required.
	Body *CreateCodeInterpreterInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCodeInterpreterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeInterpreterRequest) GoString() string {
	return s.String()
}

func (s *CreateCodeInterpreterRequest) GetBody() *CreateCodeInterpreterInput {
	return s.Body
}

func (s *CreateCodeInterpreterRequest) SetBody(v *CreateCodeInterpreterInput) *CreateCodeInterpreterRequest {
	s.Body = v
	return s
}

func (s *CreateCodeInterpreterRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
