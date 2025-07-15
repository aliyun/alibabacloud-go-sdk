// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchAuthenticationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LaunchAuthenticationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LaunchAuthenticationResponse
	GetStatusCode() *int32
	SetBody(v *LaunchAuthenticationResponseBody) *LaunchAuthenticationResponse
	GetBody() *LaunchAuthenticationResponseBody
}

type LaunchAuthenticationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LaunchAuthenticationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LaunchAuthenticationResponse) String() string {
	return dara.Prettify(s)
}

func (s LaunchAuthenticationResponse) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LaunchAuthenticationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LaunchAuthenticationResponse) GetBody() *LaunchAuthenticationResponseBody {
	return s.Body
}

func (s *LaunchAuthenticationResponse) SetHeaders(v map[string]*string) *LaunchAuthenticationResponse {
	s.Headers = v
	return s
}

func (s *LaunchAuthenticationResponse) SetStatusCode(v int32) *LaunchAuthenticationResponse {
	s.StatusCode = &v
	return s
}

func (s *LaunchAuthenticationResponse) SetBody(v *LaunchAuthenticationResponseBody) *LaunchAuthenticationResponse {
	s.Body = v
	return s
}

func (s *LaunchAuthenticationResponse) Validate() error {
	return dara.Validate(s)
}
