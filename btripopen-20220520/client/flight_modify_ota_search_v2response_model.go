// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyOtaSearchV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightModifyOtaSearchV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightModifyOtaSearchV2Response
	GetStatusCode() *int32
	SetBody(v *FlightModifyOtaSearchV2ResponseBody) *FlightModifyOtaSearchV2Response
	GetBody() *FlightModifyOtaSearchV2ResponseBody
}

type FlightModifyOtaSearchV2Response struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightModifyOtaSearchV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightModifyOtaSearchV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2Response) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightModifyOtaSearchV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightModifyOtaSearchV2Response) GetBody() *FlightModifyOtaSearchV2ResponseBody {
	return s.Body
}

func (s *FlightModifyOtaSearchV2Response) SetHeaders(v map[string]*string) *FlightModifyOtaSearchV2Response {
	s.Headers = v
	return s
}

func (s *FlightModifyOtaSearchV2Response) SetStatusCode(v int32) *FlightModifyOtaSearchV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2Response) SetBody(v *FlightModifyOtaSearchV2ResponseBody) *FlightModifyOtaSearchV2Response {
	s.Body = v
	return s
}

func (s *FlightModifyOtaSearchV2Response) Validate() error {
	return dara.Validate(s)
}
