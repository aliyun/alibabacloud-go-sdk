// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightListingSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightListingSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightListingSearchResponse
	GetStatusCode() *int32
	SetBody(v *FlightListingSearchResponseBody) *FlightListingSearchResponse
	GetBody() *FlightListingSearchResponseBody
}

type FlightListingSearchResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightListingSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightListingSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchResponse) GoString() string {
	return s.String()
}

func (s *FlightListingSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightListingSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightListingSearchResponse) GetBody() *FlightListingSearchResponseBody {
	return s.Body
}

func (s *FlightListingSearchResponse) SetHeaders(v map[string]*string) *FlightListingSearchResponse {
	s.Headers = v
	return s
}

func (s *FlightListingSearchResponse) SetStatusCode(v int32) *FlightListingSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightListingSearchResponse) SetBody(v *FlightListingSearchResponseBody) *FlightListingSearchResponse {
	s.Body = v
	return s
}

func (s *FlightListingSearchResponse) Validate() error {
	return dara.Validate(s)
}
