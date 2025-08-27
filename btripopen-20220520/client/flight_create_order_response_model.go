// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCreateOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightCreateOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightCreateOrderResponse
	GetStatusCode() *int32
	SetBody(v *FlightCreateOrderResponseBody) *FlightCreateOrderResponse
	GetBody() *FlightCreateOrderResponseBody
}

type FlightCreateOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightCreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightCreateOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderResponse) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightCreateOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightCreateOrderResponse) GetBody() *FlightCreateOrderResponseBody {
	return s.Body
}

func (s *FlightCreateOrderResponse) SetHeaders(v map[string]*string) *FlightCreateOrderResponse {
	s.Headers = v
	return s
}

func (s *FlightCreateOrderResponse) SetStatusCode(v int32) *FlightCreateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightCreateOrderResponse) SetBody(v *FlightCreateOrderResponseBody) *FlightCreateOrderResponse {
	s.Body = v
	return s
}

func (s *FlightCreateOrderResponse) Validate() error {
	return dara.Validate(s)
}
