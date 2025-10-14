// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRescaleDeviceServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RescaleDeviceServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RescaleDeviceServiceResponse
	GetStatusCode() *int32
	SetBody(v *RescaleDeviceServiceResponseBody) *RescaleDeviceServiceResponse
	GetBody() *RescaleDeviceServiceResponseBody
}

type RescaleDeviceServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RescaleDeviceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RescaleDeviceServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s RescaleDeviceServiceResponse) GoString() string {
	return s.String()
}

func (s *RescaleDeviceServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RescaleDeviceServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RescaleDeviceServiceResponse) GetBody() *RescaleDeviceServiceResponseBody {
	return s.Body
}

func (s *RescaleDeviceServiceResponse) SetHeaders(v map[string]*string) *RescaleDeviceServiceResponse {
	s.Headers = v
	return s
}

func (s *RescaleDeviceServiceResponse) SetStatusCode(v int32) *RescaleDeviceServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RescaleDeviceServiceResponse) SetBody(v *RescaleDeviceServiceResponseBody) *RescaleDeviceServiceResponse {
	s.Body = v
	return s
}

func (s *RescaleDeviceServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
