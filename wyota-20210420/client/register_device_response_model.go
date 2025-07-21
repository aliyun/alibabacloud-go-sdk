// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterDeviceResponse
	GetStatusCode() *int32
	SetBody(v *RegisterDeviceResponseBody) *RegisterDeviceResponse
	GetBody() *RegisterDeviceResponseBody
}

type RegisterDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterDeviceResponse) GetBody() *RegisterDeviceResponseBody {
	return s.Body
}

func (s *RegisterDeviceResponse) SetHeaders(v map[string]*string) *RegisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *RegisterDeviceResponse) SetStatusCode(v int32) *RegisterDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDeviceResponse) SetBody(v *RegisterDeviceResponseBody) *RegisterDeviceResponse {
	s.Body = v
	return s
}

func (s *RegisterDeviceResponse) Validate() error {
	return dara.Validate(s)
}
