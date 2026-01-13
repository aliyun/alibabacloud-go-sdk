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
	SetBody(v *CustomDomainResult) *CreateCustomDomainResponse
	GetBody() *CustomDomainResult
}

type CreateCustomDomainResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomDomainResult `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateCustomDomainResponse) GetBody() *CustomDomainResult {
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

func (s *CreateCustomDomainResponse) SetBody(v *CustomDomainResult) *CreateCustomDomainResponse {
	s.Body = v
	return s
}

func (s *CreateCustomDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
