// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDeviceResponse
	GetStatusCode() *int32
	SetBody(v *StopDeviceResponseBody) *StopDeviceResponse
	GetBody() *StopDeviceResponseBody
}

type StopDeviceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDeviceResponse) GoString() string {
	return s.String()
}

func (s *StopDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDeviceResponse) GetBody() *StopDeviceResponseBody {
	return s.Body
}

func (s *StopDeviceResponse) SetHeaders(v map[string]*string) *StopDeviceResponse {
	s.Headers = v
	return s
}

func (s *StopDeviceResponse) SetStatusCode(v int32) *StopDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDeviceResponse) SetBody(v *StopDeviceResponseBody) *StopDeviceResponse {
	s.Body = v
	return s
}

func (s *StopDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
