// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelStaticInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelStaticInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelStaticInfoResponse
	GetStatusCode() *int32
	SetBody(v *HotelStaticInfoResponseBody) *HotelStaticInfoResponse
	GetBody() *HotelStaticInfoResponseBody
}

type HotelStaticInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelStaticInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelStaticInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoResponse) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelStaticInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelStaticInfoResponse) GetBody() *HotelStaticInfoResponseBody {
	return s.Body
}

func (s *HotelStaticInfoResponse) SetHeaders(v map[string]*string) *HotelStaticInfoResponse {
	s.Headers = v
	return s
}

func (s *HotelStaticInfoResponse) SetStatusCode(v int32) *HotelStaticInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelStaticInfoResponse) SetBody(v *HotelStaticInfoResponseBody) *HotelStaticInfoResponse {
	s.Body = v
	return s
}

func (s *HotelStaticInfoResponse) Validate() error {
	return dara.Validate(s)
}
