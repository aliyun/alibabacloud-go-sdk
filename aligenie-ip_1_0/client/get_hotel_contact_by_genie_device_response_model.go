// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactByGenieDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelContactByGenieDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelContactByGenieDeviceResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelContactByGenieDeviceResponseBody) *GetHotelContactByGenieDeviceResponse
	GetBody() *GetHotelContactByGenieDeviceResponseBody
}

type GetHotelContactByGenieDeviceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelContactByGenieDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelContactByGenieDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByGenieDeviceResponse) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelContactByGenieDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelContactByGenieDeviceResponse) GetBody() *GetHotelContactByGenieDeviceResponseBody {
	return s.Body
}

func (s *GetHotelContactByGenieDeviceResponse) SetHeaders(v map[string]*string) *GetHotelContactByGenieDeviceResponse {
	s.Headers = v
	return s
}

func (s *GetHotelContactByGenieDeviceResponse) SetStatusCode(v int32) *GetHotelContactByGenieDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponse) SetBody(v *GetHotelContactByGenieDeviceResponseBody) *GetHotelContactByGenieDeviceResponse {
	s.Body = v
	return s
}

func (s *GetHotelContactByGenieDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
