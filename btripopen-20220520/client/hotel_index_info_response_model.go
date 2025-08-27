// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelIndexInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelIndexInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelIndexInfoResponse
	GetStatusCode() *int32
	SetBody(v *HotelIndexInfoResponseBody) *HotelIndexInfoResponse
	GetBody() *HotelIndexInfoResponseBody
}

type HotelIndexInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelIndexInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelIndexInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelIndexInfoResponse) GoString() string {
	return s.String()
}

func (s *HotelIndexInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelIndexInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelIndexInfoResponse) GetBody() *HotelIndexInfoResponseBody {
	return s.Body
}

func (s *HotelIndexInfoResponse) SetHeaders(v map[string]*string) *HotelIndexInfoResponse {
	s.Headers = v
	return s
}

func (s *HotelIndexInfoResponse) SetStatusCode(v int32) *HotelIndexInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelIndexInfoResponse) SetBody(v *HotelIndexInfoResponseBody) *HotelIndexInfoResponse {
	s.Body = v
	return s
}

func (s *HotelIndexInfoResponse) Validate() error {
	return dara.Validate(s)
}
