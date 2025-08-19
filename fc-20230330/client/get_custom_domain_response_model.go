// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomDomainResponse
	GetStatusCode() *int32
	SetBody(v *CustomDomain) *GetCustomDomainResponse
	GetBody() *CustomDomain
}

type GetCustomDomainResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomDomain      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *GetCustomDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomDomainResponse) GetBody() *CustomDomain {
	return s.Body
}

func (s *GetCustomDomainResponse) SetHeaders(v map[string]*string) *GetCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *GetCustomDomainResponse) SetStatusCode(v int32) *GetCustomDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomDomainResponse) SetBody(v *CustomDomain) *GetCustomDomainResponse {
	s.Body = v
	return s
}

func (s *GetCustomDomainResponse) Validate() error {
	return dara.Validate(s)
}
