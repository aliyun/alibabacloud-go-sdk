// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateCredentialInput) *UpdateCredentialRequest
	GetBody() *UpdateCredentialInput
}

type UpdateCredentialRequest struct {
	Body *UpdateCredentialInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialRequest) GetBody() *UpdateCredentialInput {
	return s.Body
}

func (s *UpdateCredentialRequest) SetBody(v *UpdateCredentialInput) *UpdateCredentialRequest {
	s.Body = v
	return s
}

func (s *UpdateCredentialRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
