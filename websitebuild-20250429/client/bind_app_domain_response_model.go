// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAppDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAppDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAppDomainResponse
	GetStatusCode() *int32
	SetBody(v *BindAppDomainResponseBody) *BindAppDomainResponse
	GetBody() *BindAppDomainResponseBody
}

type BindAppDomainResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAppDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAppDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAppDomainResponse) GoString() string {
	return s.String()
}

func (s *BindAppDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAppDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAppDomainResponse) GetBody() *BindAppDomainResponseBody {
	return s.Body
}

func (s *BindAppDomainResponse) SetHeaders(v map[string]*string) *BindAppDomainResponse {
	s.Headers = v
	return s
}

func (s *BindAppDomainResponse) SetStatusCode(v int32) *BindAppDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAppDomainResponse) SetBody(v *BindAppDomainResponseBody) *BindAppDomainResponse {
	s.Body = v
	return s
}

func (s *BindAppDomainResponse) Validate() error {
	return dara.Validate(s)
}
