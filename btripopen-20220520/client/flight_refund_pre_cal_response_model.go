// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundPreCalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightRefundPreCalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightRefundPreCalResponse
	GetStatusCode() *int32
	SetBody(v *FlightRefundPreCalResponseBody) *FlightRefundPreCalResponse
	GetBody() *FlightRefundPreCalResponseBody
}

type FlightRefundPreCalResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightRefundPreCalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightRefundPreCalResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalResponse) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightRefundPreCalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightRefundPreCalResponse) GetBody() *FlightRefundPreCalResponseBody {
	return s.Body
}

func (s *FlightRefundPreCalResponse) SetHeaders(v map[string]*string) *FlightRefundPreCalResponse {
	s.Headers = v
	return s
}

func (s *FlightRefundPreCalResponse) SetStatusCode(v int32) *FlightRefundPreCalResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightRefundPreCalResponse) SetBody(v *FlightRefundPreCalResponseBody) *FlightRefundPreCalResponse {
	s.Body = v
	return s
}

func (s *FlightRefundPreCalResponse) Validate() error {
	return dara.Validate(s)
}
