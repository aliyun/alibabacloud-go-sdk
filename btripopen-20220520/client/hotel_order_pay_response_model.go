// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderPayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelOrderPayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelOrderPayResponse
	GetStatusCode() *int32
	SetBody(v *HotelOrderPayResponseBody) *HotelOrderPayResponse
	GetBody() *HotelOrderPayResponseBody
}

type HotelOrderPayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelOrderPayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelOrderPayResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPayResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderPayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelOrderPayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelOrderPayResponse) GetBody() *HotelOrderPayResponseBody {
	return s.Body
}

func (s *HotelOrderPayResponse) SetHeaders(v map[string]*string) *HotelOrderPayResponse {
	s.Headers = v
	return s
}

func (s *HotelOrderPayResponse) SetStatusCode(v int32) *HotelOrderPayResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelOrderPayResponse) SetBody(v *HotelOrderPayResponseBody) *HotelOrderPayResponse {
	s.Body = v
	return s
}

func (s *HotelOrderPayResponse) Validate() error {
	return dara.Validate(s)
}
