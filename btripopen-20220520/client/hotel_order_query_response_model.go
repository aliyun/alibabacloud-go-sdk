// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelOrderQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelOrderQueryResponse
	GetStatusCode() *int32
	SetBody(v *HotelOrderQueryResponseBody) *HotelOrderQueryResponse
	GetBody() *HotelOrderQueryResponseBody
}

type HotelOrderQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelOrderQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelOrderQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelOrderQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelOrderQueryResponse) GetBody() *HotelOrderQueryResponseBody {
	return s.Body
}

func (s *HotelOrderQueryResponse) SetHeaders(v map[string]*string) *HotelOrderQueryResponse {
	s.Headers = v
	return s
}

func (s *HotelOrderQueryResponse) SetStatusCode(v int32) *HotelOrderQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelOrderQueryResponse) SetBody(v *HotelOrderQueryResponseBody) *HotelOrderQueryResponse {
	s.Body = v
	return s
}

func (s *HotelOrderQueryResponse) Validate() error {
	return dara.Validate(s)
}
