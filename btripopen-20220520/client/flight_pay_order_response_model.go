// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightPayOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightPayOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightPayOrderResponse
	GetStatusCode() *int32
	SetBody(v *FlightPayOrderResponseBody) *FlightPayOrderResponse
	GetBody() *FlightPayOrderResponseBody
}

type FlightPayOrderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightPayOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightPayOrderResponse) GoString() string {
	return s.String()
}

func (s *FlightPayOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightPayOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightPayOrderResponse) GetBody() *FlightPayOrderResponseBody {
	return s.Body
}

func (s *FlightPayOrderResponse) SetHeaders(v map[string]*string) *FlightPayOrderResponse {
	s.Headers = v
	return s
}

func (s *FlightPayOrderResponse) SetStatusCode(v int32) *FlightPayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightPayOrderResponse) SetBody(v *FlightPayOrderResponseBody) *FlightPayOrderResponse {
	s.Body = v
	return s
}

func (s *FlightPayOrderResponse) Validate() error {
	return dara.Validate(s)
}
