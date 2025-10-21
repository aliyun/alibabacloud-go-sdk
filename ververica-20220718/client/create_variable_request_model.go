// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Variable) *CreateVariableRequest
	GetBody() *Variable
}

type CreateVariableRequest struct {
	// The parameter that is used to create the variable.
	//
	// This parameter is required.
	Body *Variable `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVariableRequest) GoString() string {
	return s.String()
}

func (s *CreateVariableRequest) GetBody() *Variable {
	return s.Body
}

func (s *CreateVariableRequest) SetBody(v *Variable) *CreateVariableRequest {
	s.Body = v
	return s
}

func (s *CreateVariableRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
