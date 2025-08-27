// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaSearchV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightOtaSearchV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightOtaSearchV2Response
	GetStatusCode() *int32
	SetBody(v *FlightOtaSearchV2ResponseBody) *FlightOtaSearchV2Response
	GetBody() *FlightOtaSearchV2ResponseBody
}

type FlightOtaSearchV2Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightOtaSearchV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightOtaSearchV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2Response) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightOtaSearchV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightOtaSearchV2Response) GetBody() *FlightOtaSearchV2ResponseBody {
	return s.Body
}

func (s *FlightOtaSearchV2Response) SetHeaders(v map[string]*string) *FlightOtaSearchV2Response {
	s.Headers = v
	return s
}

func (s *FlightOtaSearchV2Response) SetStatusCode(v int32) *FlightOtaSearchV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightOtaSearchV2Response) SetBody(v *FlightOtaSearchV2ResponseBody) *FlightOtaSearchV2Response {
	s.Body = v
	return s
}

func (s *FlightOtaSearchV2Response) Validate() error {
	return dara.Validate(s)
}
