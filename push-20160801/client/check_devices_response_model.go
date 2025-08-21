// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDevicesResponse
	GetStatusCode() *int32
	SetBody(v *CheckDevicesResponseBody) *CheckDevicesResponse
	GetBody() *CheckDevicesResponseBody
}

type CheckDevicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDevicesResponse) GoString() string {
	return s.String()
}

func (s *CheckDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDevicesResponse) GetBody() *CheckDevicesResponseBody {
	return s.Body
}

func (s *CheckDevicesResponse) SetHeaders(v map[string]*string) *CheckDevicesResponse {
	s.Headers = v
	return s
}

func (s *CheckDevicesResponse) SetStatusCode(v int32) *CheckDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDevicesResponse) SetBody(v *CheckDevicesResponseBody) *CheckDevicesResponse {
	s.Body = v
	return s
}

func (s *CheckDevicesResponse) Validate() error {
	return dara.Validate(s)
}
