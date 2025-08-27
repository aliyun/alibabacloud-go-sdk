// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderChangeDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelOrderChangeDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelOrderChangeDetailResponse
	GetStatusCode() *int32
	SetBody(v *HotelOrderChangeDetailResponseBody) *HotelOrderChangeDetailResponse
	GetBody() *HotelOrderChangeDetailResponseBody
}

type HotelOrderChangeDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelOrderChangeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelOrderChangeDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeDetailResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelOrderChangeDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelOrderChangeDetailResponse) GetBody() *HotelOrderChangeDetailResponseBody {
	return s.Body
}

func (s *HotelOrderChangeDetailResponse) SetHeaders(v map[string]*string) *HotelOrderChangeDetailResponse {
	s.Headers = v
	return s
}

func (s *HotelOrderChangeDetailResponse) SetStatusCode(v int32) *HotelOrderChangeDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelOrderChangeDetailResponse) SetBody(v *HotelOrderChangeDetailResponseBody) *HotelOrderChangeDetailResponse {
	s.Body = v
	return s
}

func (s *HotelOrderChangeDetailResponse) Validate() error {
	return dara.Validate(s)
}
