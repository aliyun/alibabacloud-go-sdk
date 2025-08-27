// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightItineraryScanQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightItineraryScanQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightItineraryScanQueryResponse
	GetStatusCode() *int32
	SetBody(v *FlightItineraryScanQueryResponseBody) *FlightItineraryScanQueryResponse
	GetBody() *FlightItineraryScanQueryResponseBody
}

type FlightItineraryScanQueryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightItineraryScanQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightItineraryScanQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightItineraryScanQueryResponse) GoString() string {
	return s.String()
}

func (s *FlightItineraryScanQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightItineraryScanQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightItineraryScanQueryResponse) GetBody() *FlightItineraryScanQueryResponseBody {
	return s.Body
}

func (s *FlightItineraryScanQueryResponse) SetHeaders(v map[string]*string) *FlightItineraryScanQueryResponse {
	s.Headers = v
	return s
}

func (s *FlightItineraryScanQueryResponse) SetStatusCode(v int32) *FlightItineraryScanQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightItineraryScanQueryResponse) SetBody(v *FlightItineraryScanQueryResponseBody) *FlightItineraryScanQueryResponse {
	s.Body = v
	return s
}

func (s *FlightItineraryScanQueryResponse) Validate() error {
	return dara.Validate(s)
}
