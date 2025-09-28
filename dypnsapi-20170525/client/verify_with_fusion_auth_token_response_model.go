// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWithFusionAuthTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyWithFusionAuthTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyWithFusionAuthTokenResponse
	GetStatusCode() *int32
	SetBody(v *VerifyWithFusionAuthTokenResponseBody) *VerifyWithFusionAuthTokenResponse
	GetBody() *VerifyWithFusionAuthTokenResponseBody
}

type VerifyWithFusionAuthTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyWithFusionAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyWithFusionAuthTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyWithFusionAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *VerifyWithFusionAuthTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyWithFusionAuthTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyWithFusionAuthTokenResponse) GetBody() *VerifyWithFusionAuthTokenResponseBody {
	return s.Body
}

func (s *VerifyWithFusionAuthTokenResponse) SetHeaders(v map[string]*string) *VerifyWithFusionAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *VerifyWithFusionAuthTokenResponse) SetStatusCode(v int32) *VerifyWithFusionAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponse) SetBody(v *VerifyWithFusionAuthTokenResponseBody) *VerifyWithFusionAuthTokenResponse {
	s.Body = v
	return s
}

func (s *VerifyWithFusionAuthTokenResponse) Validate() error {
	return dara.Validate(s)
}
