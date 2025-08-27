// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderInfoQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelOrderInfoQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelOrderInfoQueryResponse
	GetStatusCode() *int32
	SetBody(v *HotelOrderInfoQueryResponseBody) *HotelOrderInfoQueryResponse
	GetBody() *HotelOrderInfoQueryResponseBody
}

type HotelOrderInfoQueryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelOrderInfoQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelOrderInfoQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelOrderInfoQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelOrderInfoQueryResponse) GetBody() *HotelOrderInfoQueryResponseBody {
	return s.Body
}

func (s *HotelOrderInfoQueryResponse) SetHeaders(v map[string]*string) *HotelOrderInfoQueryResponse {
	s.Headers = v
	return s
}

func (s *HotelOrderInfoQueryResponse) SetStatusCode(v int32) *HotelOrderInfoQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelOrderInfoQueryResponse) SetBody(v *HotelOrderInfoQueryResponseBody) *HotelOrderInfoQueryResponse {
	s.Body = v
	return s
}

func (s *HotelOrderInfoQueryResponse) Validate() error {
	return dara.Validate(s)
}
