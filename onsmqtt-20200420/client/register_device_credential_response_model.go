// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDeviceCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterDeviceCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterDeviceCredentialResponse
	GetStatusCode() *int32
	SetBody(v *RegisterDeviceCredentialResponseBody) *RegisterDeviceCredentialResponse
	GetBody() *RegisterDeviceCredentialResponseBody
}

type RegisterDeviceCredentialResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterDeviceCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterDeviceCredentialResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterDeviceCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterDeviceCredentialResponse) GetBody() *RegisterDeviceCredentialResponseBody {
	return s.Body
}

func (s *RegisterDeviceCredentialResponse) SetHeaders(v map[string]*string) *RegisterDeviceCredentialResponse {
	s.Headers = v
	return s
}

func (s *RegisterDeviceCredentialResponse) SetStatusCode(v int32) *RegisterDeviceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDeviceCredentialResponse) SetBody(v *RegisterDeviceCredentialResponseBody) *RegisterDeviceCredentialResponse {
	s.Body = v
	return s
}

func (s *RegisterDeviceCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
