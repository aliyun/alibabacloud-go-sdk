// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelControlDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotelControlDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotelControlDeviceResponse
	GetStatusCode() *int32
	SetBody(v *ListHotelControlDeviceResponseBody) *ListHotelControlDeviceResponse
	GetBody() *ListHotelControlDeviceResponseBody
}

type ListHotelControlDeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelControlDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelControlDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotelControlDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotelControlDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelControlDeviceResponse) GetBody() *ListHotelControlDeviceResponseBody {
	return s.Body
}

func (s *ListHotelControlDeviceResponse) SetHeaders(v map[string]*string) *ListHotelControlDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListHotelControlDeviceResponse) SetStatusCode(v int32) *ListHotelControlDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelControlDeviceResponse) SetBody(v *ListHotelControlDeviceResponseBody) *ListHotelControlDeviceResponse {
	s.Body = v
	return s
}

func (s *ListHotelControlDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
