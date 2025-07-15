// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterDevicesResponse
	GetStatusCode() *int32
	SetBody(v *RegisterDevicesResponseBody) *RegisterDevicesResponse
	GetBody() *RegisterDevicesResponseBody
}

type RegisterDevicesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterDevicesResponse) GoString() string {
	return s.String()
}

func (s *RegisterDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterDevicesResponse) GetBody() *RegisterDevicesResponseBody {
	return s.Body
}

func (s *RegisterDevicesResponse) SetHeaders(v map[string]*string) *RegisterDevicesResponse {
	s.Headers = v
	return s
}

func (s *RegisterDevicesResponse) SetStatusCode(v int32) *RegisterDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDevicesResponse) SetBody(v *RegisterDevicesResponseBody) *RegisterDevicesResponse {
	s.Body = v
	return s
}

func (s *RegisterDevicesResponse) Validate() error {
	return dara.Validate(s)
}
