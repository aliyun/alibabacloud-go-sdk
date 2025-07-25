// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInstanceDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindInstanceDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindInstanceDomainsResponse
	GetStatusCode() *int32
	SetBody(v *BindInstanceDomainsResponseBody) *BindInstanceDomainsResponse
	GetBody() *BindInstanceDomainsResponseBody
}

type BindInstanceDomainsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindInstanceDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindInstanceDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s BindInstanceDomainsResponse) GoString() string {
	return s.String()
}

func (s *BindInstanceDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindInstanceDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindInstanceDomainsResponse) GetBody() *BindInstanceDomainsResponseBody {
	return s.Body
}

func (s *BindInstanceDomainsResponse) SetHeaders(v map[string]*string) *BindInstanceDomainsResponse {
	s.Headers = v
	return s
}

func (s *BindInstanceDomainsResponse) SetStatusCode(v int32) *BindInstanceDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *BindInstanceDomainsResponse) SetBody(v *BindInstanceDomainsResponseBody) *BindInstanceDomainsResponse {
	s.Body = v
	return s
}

func (s *BindInstanceDomainsResponse) Validate() error {
	return dara.Validate(s)
}
