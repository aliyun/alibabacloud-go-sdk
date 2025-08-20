// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthorizationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthorizationTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthorizationTokenResponseBody) *GetAuthorizationTokenResponse
	GetBody() *GetAuthorizationTokenResponseBody
}

type GetAuthorizationTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthorizationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthorizationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorizationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthorizationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthorizationTokenResponse) GetBody() *GetAuthorizationTokenResponseBody {
	return s.Body
}

func (s *GetAuthorizationTokenResponse) SetHeaders(v map[string]*string) *GetAuthorizationTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorizationTokenResponse) SetStatusCode(v int32) *GetAuthorizationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthorizationTokenResponse) SetBody(v *GetAuthorizationTokenResponseBody) *GetAuthorizationTokenResponse {
	s.Body = v
	return s
}

func (s *GetAuthorizationTokenResponse) Validate() error {
	return dara.Validate(s)
}
