// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstanceDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppInstanceDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppInstanceDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppInstanceDomainsResponseBody) *ListAppInstanceDomainsResponse
	GetBody() *ListAppInstanceDomainsResponseBody
}

type ListAppInstanceDomainsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppInstanceDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppInstanceDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListAppInstanceDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppInstanceDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppInstanceDomainsResponse) GetBody() *ListAppInstanceDomainsResponseBody {
	return s.Body
}

func (s *ListAppInstanceDomainsResponse) SetHeaders(v map[string]*string) *ListAppInstanceDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListAppInstanceDomainsResponse) SetStatusCode(v int32) *ListAppInstanceDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppInstanceDomainsResponse) SetBody(v *ListAppInstanceDomainsResponseBody) *ListAppInstanceDomainsResponse {
	s.Body = v
	return s
}

func (s *ListAppInstanceDomainsResponse) Validate() error {
	return dara.Validate(s)
}
