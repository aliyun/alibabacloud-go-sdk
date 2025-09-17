// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessTokensResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccessTokensResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccessTokensResponse
	GetStatusCode() *int32
	SetBody(v *ListAccessTokensResponseBody) *ListAccessTokensResponse
	GetBody() *ListAccessTokensResponseBody
}

type ListAccessTokensResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessTokensResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessTokensResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccessTokensResponse) GoString() string {
	return s.String()
}

func (s *ListAccessTokensResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccessTokensResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccessTokensResponse) GetBody() *ListAccessTokensResponseBody {
	return s.Body
}

func (s *ListAccessTokensResponse) SetHeaders(v map[string]*string) *ListAccessTokensResponse {
	s.Headers = v
	return s
}

func (s *ListAccessTokensResponse) SetStatusCode(v int32) *ListAccessTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessTokensResponse) SetBody(v *ListAccessTokensResponseBody) *ListAccessTokensResponse {
	s.Body = v
	return s
}

func (s *ListAccessTokensResponse) Validate() error {
	return dara.Validate(s)
}
