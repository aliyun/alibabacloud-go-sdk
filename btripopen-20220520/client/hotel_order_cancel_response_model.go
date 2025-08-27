// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderCancelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelOrderCancelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelOrderCancelResponse
	GetStatusCode() *int32
	SetBody(v *HotelOrderCancelResponseBody) *HotelOrderCancelResponse
	GetBody() *HotelOrderCancelResponseBody
}

type HotelOrderCancelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelOrderCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelOrderCancelResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCancelResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderCancelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelOrderCancelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelOrderCancelResponse) GetBody() *HotelOrderCancelResponseBody {
	return s.Body
}

func (s *HotelOrderCancelResponse) SetHeaders(v map[string]*string) *HotelOrderCancelResponse {
	s.Headers = v
	return s
}

func (s *HotelOrderCancelResponse) SetStatusCode(v int32) *HotelOrderCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelOrderCancelResponse) SetBody(v *HotelOrderCancelResponseBody) *HotelOrderCancelResponse {
	s.Body = v
	return s
}

func (s *HotelOrderCancelResponse) Validate() error {
	return dara.Validate(s)
}
