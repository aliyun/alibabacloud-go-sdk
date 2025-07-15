// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApiProductsAuthoritiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetApiProductsAuthoritiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetApiProductsAuthoritiesResponse
	GetStatusCode() *int32
	SetBody(v *SetApiProductsAuthoritiesResponseBody) *SetApiProductsAuthoritiesResponse
	GetBody() *SetApiProductsAuthoritiesResponseBody
}

type SetApiProductsAuthoritiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApiProductsAuthoritiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApiProductsAuthoritiesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetApiProductsAuthoritiesResponse) GoString() string {
	return s.String()
}

func (s *SetApiProductsAuthoritiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetApiProductsAuthoritiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetApiProductsAuthoritiesResponse) GetBody() *SetApiProductsAuthoritiesResponseBody {
	return s.Body
}

func (s *SetApiProductsAuthoritiesResponse) SetHeaders(v map[string]*string) *SetApiProductsAuthoritiesResponse {
	s.Headers = v
	return s
}

func (s *SetApiProductsAuthoritiesResponse) SetStatusCode(v int32) *SetApiProductsAuthoritiesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApiProductsAuthoritiesResponse) SetBody(v *SetApiProductsAuthoritiesResponseBody) *SetApiProductsAuthoritiesResponse {
	s.Body = v
	return s
}

func (s *SetApiProductsAuthoritiesResponse) Validate() error {
	return dara.Validate(s)
}
