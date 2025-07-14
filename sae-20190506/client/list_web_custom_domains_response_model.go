// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebCustomDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWebCustomDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWebCustomDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ListWebCustomDomainBody) *ListWebCustomDomainsResponse
	GetBody() *ListWebCustomDomainBody
}

type ListWebCustomDomainsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWebCustomDomainBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWebCustomDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWebCustomDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListWebCustomDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWebCustomDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWebCustomDomainsResponse) GetBody() *ListWebCustomDomainBody {
	return s.Body
}

func (s *ListWebCustomDomainsResponse) SetHeaders(v map[string]*string) *ListWebCustomDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListWebCustomDomainsResponse) SetStatusCode(v int32) *ListWebCustomDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWebCustomDomainsResponse) SetBody(v *ListWebCustomDomainBody) *ListWebCustomDomainsResponse {
	s.Body = v
	return s
}

func (s *ListWebCustomDomainsResponse) Validate() error {
	return dara.Validate(s)
}
