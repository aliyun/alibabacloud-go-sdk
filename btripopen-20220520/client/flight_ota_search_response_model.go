// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightOtaSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightOtaSearchResponse
	GetStatusCode() *int32
	SetBody(v *FlightOtaSearchResponseBody) *FlightOtaSearchResponse
	GetBody() *FlightOtaSearchResponseBody
}

type FlightOtaSearchResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightOtaSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightOtaSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchResponse) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightOtaSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightOtaSearchResponse) GetBody() *FlightOtaSearchResponseBody {
	return s.Body
}

func (s *FlightOtaSearchResponse) SetHeaders(v map[string]*string) *FlightOtaSearchResponse {
	s.Headers = v
	return s
}

func (s *FlightOtaSearchResponse) SetStatusCode(v int32) *FlightOtaSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightOtaSearchResponse) SetBody(v *FlightOtaSearchResponseBody) *FlightOtaSearchResponse {
	s.Body = v
	return s
}

func (s *FlightOtaSearchResponse) Validate() error {
	return dara.Validate(s)
}
