// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelQrBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelQrBindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelQrBindResponse
	GetStatusCode() *int32
	SetBody(v *HotelQrBindResponseBody) *HotelQrBindResponse
	GetBody() *HotelQrBindResponseBody
}

type HotelQrBindResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelQrBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelQrBindResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelQrBindResponse) GoString() string {
	return s.String()
}

func (s *HotelQrBindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelQrBindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelQrBindResponse) GetBody() *HotelQrBindResponseBody {
	return s.Body
}

func (s *HotelQrBindResponse) SetHeaders(v map[string]*string) *HotelQrBindResponse {
	s.Headers = v
	return s
}

func (s *HotelQrBindResponse) SetStatusCode(v int32) *HotelQrBindResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelQrBindResponse) SetBody(v *HotelQrBindResponseBody) *HotelQrBindResponse {
	s.Body = v
	return s
}

func (s *HotelQrBindResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
