// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelGoodsQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelGoodsQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelGoodsQueryResponse
	GetStatusCode() *int32
	SetBody(v *HotelGoodsQueryResponseBody) *HotelGoodsQueryResponse
	GetBody() *HotelGoodsQueryResponseBody
}

type HotelGoodsQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelGoodsQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelGoodsQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponse) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelGoodsQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelGoodsQueryResponse) GetBody() *HotelGoodsQueryResponseBody {
	return s.Body
}

func (s *HotelGoodsQueryResponse) SetHeaders(v map[string]*string) *HotelGoodsQueryResponse {
	s.Headers = v
	return s
}

func (s *HotelGoodsQueryResponse) SetStatusCode(v int32) *HotelGoodsQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelGoodsQueryResponse) SetBody(v *HotelGoodsQueryResponseBody) *HotelGoodsQueryResponse {
	s.Body = v
	return s
}

func (s *HotelGoodsQueryResponse) Validate() error {
	return dara.Validate(s)
}
