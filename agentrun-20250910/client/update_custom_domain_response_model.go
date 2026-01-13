// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomDomainResponse
	GetStatusCode() *int32
	SetBody(v *CustomDomainResult) *UpdateCustomDomainResponse
	GetBody() *CustomDomainResult
}

type UpdateCustomDomainResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomDomainResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomDomainResponse) GetBody() *CustomDomainResult {
	return s.Body
}

func (s *UpdateCustomDomainResponse) SetHeaders(v map[string]*string) *UpdateCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomDomainResponse) SetStatusCode(v int32) *UpdateCustomDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomDomainResponse) SetBody(v *CustomDomainResult) *UpdateCustomDomainResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
