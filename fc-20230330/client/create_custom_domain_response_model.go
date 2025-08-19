// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomDomainResponse
	GetStatusCode() *int32
	SetBody(v *CustomDomain) *CreateCustomDomainResponse
	GetBody() *CustomDomain
}

type CreateCustomDomainResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomDomain      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomDomainResponse) GetBody() *CustomDomain {
	return s.Body
}

func (s *CreateCustomDomainResponse) SetHeaders(v map[string]*string) *CreateCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomDomainResponse) SetStatusCode(v int32) *CreateCustomDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomDomainResponse) SetBody(v *CustomDomain) *CreateCustomDomainResponse {
	s.Body = v
	return s
}

func (s *CreateCustomDomainResponse) Validate() error {
	return dara.Validate(s)
}
