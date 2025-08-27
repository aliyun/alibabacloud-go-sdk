// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelCityCodeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelCityCodeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelCityCodeListResponse
	GetStatusCode() *int32
	SetBody(v *HotelCityCodeListResponseBody) *HotelCityCodeListResponse
	GetBody() *HotelCityCodeListResponseBody
}

type HotelCityCodeListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelCityCodeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelCityCodeListResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelCityCodeListResponse) GoString() string {
	return s.String()
}

func (s *HotelCityCodeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelCityCodeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelCityCodeListResponse) GetBody() *HotelCityCodeListResponseBody {
	return s.Body
}

func (s *HotelCityCodeListResponse) SetHeaders(v map[string]*string) *HotelCityCodeListResponse {
	s.Headers = v
	return s
}

func (s *HotelCityCodeListResponse) SetStatusCode(v int32) *HotelCityCodeListResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelCityCodeListResponse) SetBody(v *HotelCityCodeListResponseBody) *HotelCityCodeListResponse {
	s.Body = v
	return s
}

func (s *HotelCityCodeListResponse) Validate() error {
	return dara.Validate(s)
}
