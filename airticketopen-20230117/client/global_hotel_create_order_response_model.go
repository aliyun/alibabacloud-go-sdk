// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCreateOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelCreateOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelCreateOrderResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelCreateOrderResponseBody) *GlobalHotelCreateOrderResponse
	GetBody() *GlobalHotelCreateOrderResponseBody
}

type GlobalHotelCreateOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelCreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelCreateOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateOrderResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelCreateOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelCreateOrderResponse) GetBody() *GlobalHotelCreateOrderResponseBody {
	return s.Body
}

func (s *GlobalHotelCreateOrderResponse) SetHeaders(v map[string]*string) *GlobalHotelCreateOrderResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelCreateOrderResponse) SetStatusCode(v int32) *GlobalHotelCreateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelCreateOrderResponse) SetBody(v *GlobalHotelCreateOrderResponseBody) *GlobalHotelCreateOrderResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelCreateOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
