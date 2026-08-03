// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelQueryOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelQueryOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelQueryOrderResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelQueryOrderResponseBody) *GlobalHotelQueryOrderResponse
	GetBody() *GlobalHotelQueryOrderResponseBody
}

type GlobalHotelQueryOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelQueryOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelQueryOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelQueryOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelQueryOrderResponse) GetBody() *GlobalHotelQueryOrderResponseBody {
	return s.Body
}

func (s *GlobalHotelQueryOrderResponse) SetHeaders(v map[string]*string) *GlobalHotelQueryOrderResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelQueryOrderResponse) SetStatusCode(v int32) *GlobalHotelQueryOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelQueryOrderResponse) SetBody(v *GlobalHotelQueryOrderResponseBody) *GlobalHotelQueryOrderResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelQueryOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
