// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateCustomDomainInput) *CreateCustomDomainRequest
	GetBody() *CreateCustomDomainInput
}

type CreateCustomDomainRequest struct {
	Body *CreateCustomDomainInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainRequest) GetBody() *CreateCustomDomainInput {
	return s.Body
}

func (s *CreateCustomDomainRequest) SetBody(v *CreateCustomDomainInput) *CreateCustomDomainRequest {
	s.Body = v
	return s
}

func (s *CreateCustomDomainRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
