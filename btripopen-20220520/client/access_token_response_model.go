// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *AccessTokenResponseBody) *AccessTokenResponse
	GetBody() *AccessTokenResponseBody
}

type AccessTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s AccessTokenResponse) GoString() string {
	return s.String()
}

func (s *AccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AccessTokenResponse) GetBody() *AccessTokenResponseBody {
	return s.Body
}

func (s *AccessTokenResponse) SetHeaders(v map[string]*string) *AccessTokenResponse {
	s.Headers = v
	return s
}

func (s *AccessTokenResponse) SetStatusCode(v int32) *AccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *AccessTokenResponse) SetBody(v *AccessTokenResponseBody) *AccessTokenResponse {
	s.Body = v
	return s
}

func (s *AccessTokenResponse) Validate() error {
	return dara.Validate(s)
}
