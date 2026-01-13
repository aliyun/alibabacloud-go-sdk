// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomDomainsResult) *ListCustomDomainsResponse
	GetBody() *ListCustomDomainsResult
}

type ListCustomDomainsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomDomainsResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomDomainsResponse) GetBody() *ListCustomDomainsResult {
	return s.Body
}

func (s *ListCustomDomainsResponse) SetHeaders(v map[string]*string) *ListCustomDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomDomainsResponse) SetStatusCode(v int32) *ListCustomDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomDomainsResponse) SetBody(v *ListCustomDomainsResult) *ListCustomDomainsResponse {
	s.Body = v
	return s
}

func (s *ListCustomDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
