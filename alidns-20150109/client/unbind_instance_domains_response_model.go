// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindInstanceDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindInstanceDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindInstanceDomainsResponse
	GetStatusCode() *int32
	SetBody(v *UnbindInstanceDomainsResponseBody) *UnbindInstanceDomainsResponse
	GetBody() *UnbindInstanceDomainsResponseBody
}

type UnbindInstanceDomainsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindInstanceDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindInstanceDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindInstanceDomainsResponse) GoString() string {
	return s.String()
}

func (s *UnbindInstanceDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindInstanceDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindInstanceDomainsResponse) GetBody() *UnbindInstanceDomainsResponseBody {
	return s.Body
}

func (s *UnbindInstanceDomainsResponse) SetHeaders(v map[string]*string) *UnbindInstanceDomainsResponse {
	s.Headers = v
	return s
}

func (s *UnbindInstanceDomainsResponse) SetStatusCode(v int32) *UnbindInstanceDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindInstanceDomainsResponse) SetBody(v *UnbindInstanceDomainsResponseBody) *UnbindInstanceDomainsResponse {
	s.Body = v
	return s
}

func (s *UnbindInstanceDomainsResponse) Validate() error {
	return dara.Validate(s)
}
