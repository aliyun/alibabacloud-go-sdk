// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Variable) *UpdateVariableRequest
	GetBody() *Variable
}

type UpdateVariableRequest struct {
	// This parameter is required.
	Body *Variable `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVariableRequest) GoString() string {
	return s.String()
}

func (s *UpdateVariableRequest) GetBody() *Variable {
	return s.Body
}

func (s *UpdateVariableRequest) SetBody(v *Variable) *UpdateVariableRequest {
	s.Body = v
	return s
}

func (s *UpdateVariableRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
