// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCancelOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelCancelOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelCancelOrderResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelCancelOrderResponseBody) *GlobalHotelCancelOrderResponse
	GetBody() *GlobalHotelCancelOrderResponseBody
}

type GlobalHotelCancelOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelCancelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelCancelOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCancelOrderResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelCancelOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelCancelOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelCancelOrderResponse) GetBody() *GlobalHotelCancelOrderResponseBody {
	return s.Body
}

func (s *GlobalHotelCancelOrderResponse) SetHeaders(v map[string]*string) *GlobalHotelCancelOrderResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelCancelOrderResponse) SetStatusCode(v int32) *GlobalHotelCancelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelCancelOrderResponse) SetBody(v *GlobalHotelCancelOrderResponseBody) *GlobalHotelCancelOrderResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelCancelOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
