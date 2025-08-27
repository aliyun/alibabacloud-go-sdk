// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCancelOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightCancelOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightCancelOrderResponse
	GetStatusCode() *int32
	SetBody(v *FlightCancelOrderResponseBody) *FlightCancelOrderResponse
	GetBody() *FlightCancelOrderResponseBody
}

type FlightCancelOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightCancelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightCancelOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightCancelOrderResponse) GoString() string {
	return s.String()
}

func (s *FlightCancelOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightCancelOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightCancelOrderResponse) GetBody() *FlightCancelOrderResponseBody {
	return s.Body
}

func (s *FlightCancelOrderResponse) SetHeaders(v map[string]*string) *FlightCancelOrderResponse {
	s.Headers = v
	return s
}

func (s *FlightCancelOrderResponse) SetStatusCode(v int32) *FlightCancelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightCancelOrderResponse) SetBody(v *FlightCancelOrderResponseBody) *FlightCancelOrderResponse {
	s.Body = v
	return s
}

func (s *FlightCancelOrderResponse) Validate() error {
	return dara.Validate(s)
}
