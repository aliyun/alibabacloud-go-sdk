// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderDetailInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelOrderDetailInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelOrderDetailInfoResponse
	GetStatusCode() *int32
	SetBody(v *HotelOrderDetailInfoResponseBody) *HotelOrderDetailInfoResponse
	GetBody() *HotelOrderDetailInfoResponseBody
}

type HotelOrderDetailInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelOrderDetailInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelOrderDetailInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelOrderDetailInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelOrderDetailInfoResponse) GetBody() *HotelOrderDetailInfoResponseBody {
	return s.Body
}

func (s *HotelOrderDetailInfoResponse) SetHeaders(v map[string]*string) *HotelOrderDetailInfoResponse {
	s.Headers = v
	return s
}

func (s *HotelOrderDetailInfoResponse) SetStatusCode(v int32) *HotelOrderDetailInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelOrderDetailInfoResponse) SetBody(v *HotelOrderDetailInfoResponseBody) *HotelOrderDetailInfoResponse {
	s.Body = v
	return s
}

func (s *HotelOrderDetailInfoResponse) Validate() error {
	return dara.Validate(s)
}
