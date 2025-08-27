// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelPricePullResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelPricePullResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelPricePullResponse
	GetStatusCode() *int32
	SetBody(v *HotelPricePullResponseBody) *HotelPricePullResponse
	GetBody() *HotelPricePullResponseBody
}

type HotelPricePullResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelPricePullResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelPricePullResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullResponse) GoString() string {
	return s.String()
}

func (s *HotelPricePullResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelPricePullResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelPricePullResponse) GetBody() *HotelPricePullResponseBody {
	return s.Body
}

func (s *HotelPricePullResponse) SetHeaders(v map[string]*string) *HotelPricePullResponse {
	s.Headers = v
	return s
}

func (s *HotelPricePullResponse) SetStatusCode(v int32) *HotelPricePullResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelPricePullResponse) SetBody(v *HotelPricePullResponseBody) *HotelPricePullResponse {
	s.Body = v
	return s
}

func (s *HotelPricePullResponse) Validate() error {
	return dara.Validate(s)
}
