// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDeviceResponse
	GetStatusCode() *int32
	SetBody(v *CheckDeviceResponseBody) *CheckDeviceResponse
	GetBody() *CheckDeviceResponseBody
}

type CheckDeviceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDeviceResponse) GoString() string {
	return s.String()
}

func (s *CheckDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDeviceResponse) GetBody() *CheckDeviceResponseBody {
	return s.Body
}

func (s *CheckDeviceResponse) SetHeaders(v map[string]*string) *CheckDeviceResponse {
	s.Headers = v
	return s
}

func (s *CheckDeviceResponse) SetStatusCode(v int32) *CheckDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDeviceResponse) SetBody(v *CheckDeviceResponseBody) *CheckDeviceResponse {
	s.Body = v
	return s
}

func (s *CheckDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
