// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateCustomDomainInput) *UpdateCustomDomainRequest
	GetBody() *UpdateCustomDomainInput
}

type UpdateCustomDomainRequest struct {
	Body *UpdateCustomDomainInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainRequest) GetBody() *UpdateCustomDomainInput {
	return s.Body
}

func (s *UpdateCustomDomainRequest) SetBody(v *UpdateCustomDomainInput) *UpdateCustomDomainRequest {
	s.Body = v
	return s
}

func (s *UpdateCustomDomainRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
