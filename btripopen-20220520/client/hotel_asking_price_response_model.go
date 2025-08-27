// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelAskingPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelAskingPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelAskingPriceResponse
	GetStatusCode() *int32
	SetBody(v *HotelAskingPriceResponseBody) *HotelAskingPriceResponse
	GetBody() *HotelAskingPriceResponseBody
}

type HotelAskingPriceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelAskingPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelAskingPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelAskingPriceResponse) GoString() string {
	return s.String()
}

func (s *HotelAskingPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelAskingPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelAskingPriceResponse) GetBody() *HotelAskingPriceResponseBody {
	return s.Body
}

func (s *HotelAskingPriceResponse) SetHeaders(v map[string]*string) *HotelAskingPriceResponse {
	s.Headers = v
	return s
}

func (s *HotelAskingPriceResponse) SetStatusCode(v int32) *HotelAskingPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelAskingPriceResponse) SetBody(v *HotelAskingPriceResponseBody) *HotelAskingPriceResponse {
	s.Body = v
	return s
}

func (s *HotelAskingPriceResponse) Validate() error {
	return dara.Validate(s)
}
