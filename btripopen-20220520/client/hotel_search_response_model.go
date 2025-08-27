// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelSearchResponse
	GetStatusCode() *int32
	SetBody(v *HotelSearchResponseBody) *HotelSearchResponse
	GetBody() *HotelSearchResponseBody
}

type HotelSearchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelSearchResponse) GoString() string {
	return s.String()
}

func (s *HotelSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelSearchResponse) GetBody() *HotelSearchResponseBody {
	return s.Body
}

func (s *HotelSearchResponse) SetHeaders(v map[string]*string) *HotelSearchResponse {
	s.Headers = v
	return s
}

func (s *HotelSearchResponse) SetStatusCode(v int32) *HotelSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelSearchResponse) SetBody(v *HotelSearchResponseBody) *HotelSearchResponse {
	s.Body = v
	return s
}

func (s *HotelSearchResponse) Validate() error {
	return dara.Validate(s)
}
