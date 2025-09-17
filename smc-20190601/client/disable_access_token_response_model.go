// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *DisableAccessTokenResponseBody) *DisableAccessTokenResponse
	GetBody() *DisableAccessTokenResponseBody
}

type DisableAccessTokenResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *DisableAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAccessTokenResponse) GetBody() *DisableAccessTokenResponseBody {
	return s.Body
}

func (s *DisableAccessTokenResponse) SetHeaders(v map[string]*string) *DisableAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *DisableAccessTokenResponse) SetStatusCode(v int32) *DisableAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAccessTokenResponse) SetBody(v *DisableAccessTokenResponseBody) *DisableAccessTokenResponse {
	s.Body = v
	return s
}

func (s *DisableAccessTokenResponse) Validate() error {
	return dara.Validate(s)
}
