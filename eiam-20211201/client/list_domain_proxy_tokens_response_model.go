// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainProxyTokensResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDomainProxyTokensResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDomainProxyTokensResponse
	GetStatusCode() *int32
	SetBody(v *ListDomainProxyTokensResponseBody) *ListDomainProxyTokensResponse
	GetBody() *ListDomainProxyTokensResponseBody
}

type ListDomainProxyTokensResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainProxyTokensResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainProxyTokensResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDomainProxyTokensResponse) GoString() string {
	return s.String()
}

func (s *ListDomainProxyTokensResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDomainProxyTokensResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDomainProxyTokensResponse) GetBody() *ListDomainProxyTokensResponseBody {
	return s.Body
}

func (s *ListDomainProxyTokensResponse) SetHeaders(v map[string]*string) *ListDomainProxyTokensResponse {
	s.Headers = v
	return s
}

func (s *ListDomainProxyTokensResponse) SetStatusCode(v int32) *ListDomainProxyTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainProxyTokensResponse) SetBody(v *ListDomainProxyTokensResponseBody) *ListDomainProxyTokensResponse {
	s.Body = v
	return s
}

func (s *ListDomainProxyTokensResponse) Validate() error {
	return dara.Validate(s)
}
