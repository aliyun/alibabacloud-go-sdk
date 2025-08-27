// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyListingSearchV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightModifyListingSearchV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightModifyListingSearchV2Response
	GetStatusCode() *int32
	SetBody(v *FlightModifyListingSearchV2ResponseBody) *FlightModifyListingSearchV2Response
	GetBody() *FlightModifyListingSearchV2ResponseBody
}

type FlightModifyListingSearchV2Response struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightModifyListingSearchV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightModifyListingSearchV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2Response) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightModifyListingSearchV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightModifyListingSearchV2Response) GetBody() *FlightModifyListingSearchV2ResponseBody {
	return s.Body
}

func (s *FlightModifyListingSearchV2Response) SetHeaders(v map[string]*string) *FlightModifyListingSearchV2Response {
	s.Headers = v
	return s
}

func (s *FlightModifyListingSearchV2Response) SetStatusCode(v int32) *FlightModifyListingSearchV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightModifyListingSearchV2Response) SetBody(v *FlightModifyListingSearchV2ResponseBody) *FlightModifyListingSearchV2Response {
	s.Body = v
	return s
}

func (s *FlightModifyListingSearchV2Response) Validate() error {
	return dara.Validate(s)
}
