// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderListQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelOrderListQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelOrderListQueryResponse
	GetStatusCode() *int32
	SetBody(v *HotelOrderListQueryResponseBody) *HotelOrderListQueryResponse
	GetBody() *HotelOrderListQueryResponseBody
}

type HotelOrderListQueryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelOrderListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelOrderListQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderListQueryResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelOrderListQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelOrderListQueryResponse) GetBody() *HotelOrderListQueryResponseBody {
	return s.Body
}

func (s *HotelOrderListQueryResponse) SetHeaders(v map[string]*string) *HotelOrderListQueryResponse {
	s.Headers = v
	return s
}

func (s *HotelOrderListQueryResponse) SetStatusCode(v int32) *HotelOrderListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelOrderListQueryResponse) SetBody(v *HotelOrderListQueryResponseBody) *HotelOrderListQueryResponse {
	s.Body = v
	return s
}

func (s *HotelOrderListQueryResponse) Validate() error {
	return dara.Validate(s)
}
