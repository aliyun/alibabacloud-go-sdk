// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDeviceResponse
	GetStatusCode() *int32
	SetBody(v *DisableDeviceResponseBody) *DisableDeviceResponse
	GetBody() *DisableDeviceResponseBody
}

type DisableDeviceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDeviceResponse) GoString() string {
	return s.String()
}

func (s *DisableDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDeviceResponse) GetBody() *DisableDeviceResponseBody {
	return s.Body
}

func (s *DisableDeviceResponse) SetHeaders(v map[string]*string) *DisableDeviceResponse {
	s.Headers = v
	return s
}

func (s *DisableDeviceResponse) SetStatusCode(v int32) *DisableDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDeviceResponse) SetBody(v *DisableDeviceResponseBody) *DisableDeviceResponse {
	s.Body = v
	return s
}

func (s *DisableDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
