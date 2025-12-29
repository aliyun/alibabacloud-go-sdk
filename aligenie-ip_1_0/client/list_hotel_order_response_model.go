// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotelOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotelOrderResponse
	GetStatusCode() *int32
	SetBody(v *ListHotelOrderResponseBody) *ListHotelOrderResponse
	GetBody() *ListHotelOrderResponseBody
}

type ListHotelOrderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotelOrderResponse) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotelOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelOrderResponse) GetBody() *ListHotelOrderResponseBody {
	return s.Body
}

func (s *ListHotelOrderResponse) SetHeaders(v map[string]*string) *ListHotelOrderResponse {
	s.Headers = v
	return s
}

func (s *ListHotelOrderResponse) SetStatusCode(v int32) *ListHotelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelOrderResponse) SetBody(v *ListHotelOrderResponseBody) *ListHotelOrderResponse {
	s.Body = v
	return s
}

func (s *ListHotelOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
