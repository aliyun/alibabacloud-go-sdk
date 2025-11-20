// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateCredentialInput) *CreateCredentialRequest
	GetBody() *CreateCredentialInput
}

type CreateCredentialRequest struct {
	Body *CreateCredentialInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequest) GetBody() *CreateCredentialInput {
	return s.Body
}

func (s *CreateCredentialRequest) SetBody(v *CreateCredentialInput) *CreateCredentialRequest {
	s.Body = v
	return s
}

func (s *CreateCredentialRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
