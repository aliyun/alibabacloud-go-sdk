// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAuthenticationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyAuthenticationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyAuthenticationResponse
	GetStatusCode() *int32
	SetBody(v *VerifyAuthenticationResponseBody) *VerifyAuthenticationResponse
	GetBody() *VerifyAuthenticationResponseBody
}

type VerifyAuthenticationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyAuthenticationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyAuthenticationResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyAuthenticationResponse) GoString() string {
	return s.String()
}

func (s *VerifyAuthenticationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyAuthenticationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyAuthenticationResponse) GetBody() *VerifyAuthenticationResponseBody {
	return s.Body
}

func (s *VerifyAuthenticationResponse) SetHeaders(v map[string]*string) *VerifyAuthenticationResponse {
	s.Headers = v
	return s
}

func (s *VerifyAuthenticationResponse) SetStatusCode(v int32) *VerifyAuthenticationResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyAuthenticationResponse) SetBody(v *VerifyAuthenticationResponseBody) *VerifyAuthenticationResponse {
	s.Body = v
	return s
}

func (s *VerifyAuthenticationResponse) Validate() error {
	return dara.Validate(s)
}
