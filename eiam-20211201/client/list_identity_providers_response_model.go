// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIdentityProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIdentityProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListIdentityProvidersResponseBody) *ListIdentityProvidersResponse
	GetBody() *ListIdentityProvidersResponseBody
}

type ListIdentityProvidersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIdentityProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdentityProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIdentityProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIdentityProvidersResponse) GetBody() *ListIdentityProvidersResponseBody {
	return s.Body
}

func (s *ListIdentityProvidersResponse) SetHeaders(v map[string]*string) *ListIdentityProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListIdentityProvidersResponse) SetStatusCode(v int32) *ListIdentityProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdentityProvidersResponse) SetBody(v *ListIdentityProvidersResponseBody) *ListIdentityProvidersResponse {
	s.Body = v
	return s
}

func (s *ListIdentityProvidersResponse) Validate() error {
	return dara.Validate(s)
}
