// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightListingSearchV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightListingSearchV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightListingSearchV2Response
	GetStatusCode() *int32
	SetBody(v *FlightListingSearchV2ResponseBody) *FlightListingSearchV2Response
	GetBody() *FlightListingSearchV2ResponseBody
}

type FlightListingSearchV2Response struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightListingSearchV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightListingSearchV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2Response) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightListingSearchV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightListingSearchV2Response) GetBody() *FlightListingSearchV2ResponseBody {
	return s.Body
}

func (s *FlightListingSearchV2Response) SetHeaders(v map[string]*string) *FlightListingSearchV2Response {
	s.Headers = v
	return s
}

func (s *FlightListingSearchV2Response) SetStatusCode(v int32) *FlightListingSearchV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightListingSearchV2Response) SetBody(v *FlightListingSearchV2ResponseBody) *FlightListingSearchV2Response {
	s.Body = v
	return s
}

func (s *FlightListingSearchV2Response) Validate() error {
	return dara.Validate(s)
}
