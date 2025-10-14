// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightChangeOfOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightChangeOfOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightChangeOfOrderResponse
	GetStatusCode() *int32
	SetBody(v *FlightChangeOfOrderResponseBody) *FlightChangeOfOrderResponse
	GetBody() *FlightChangeOfOrderResponseBody
}

type FlightChangeOfOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightChangeOfOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightChangeOfOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightChangeOfOrderResponse) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightChangeOfOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightChangeOfOrderResponse) GetBody() *FlightChangeOfOrderResponseBody {
	return s.Body
}

func (s *FlightChangeOfOrderResponse) SetHeaders(v map[string]*string) *FlightChangeOfOrderResponse {
	s.Headers = v
	return s
}

func (s *FlightChangeOfOrderResponse) SetStatusCode(v int32) *FlightChangeOfOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightChangeOfOrderResponse) SetBody(v *FlightChangeOfOrderResponseBody) *FlightChangeOfOrderResponse {
	s.Body = v
	return s
}

func (s *FlightChangeOfOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
