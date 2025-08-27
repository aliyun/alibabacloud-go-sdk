// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCreateOrderV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightCreateOrderV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightCreateOrderV2Response
	GetStatusCode() *int32
	SetBody(v *FlightCreateOrderV2ResponseBody) *FlightCreateOrderV2Response
	GetBody() *FlightCreateOrderV2ResponseBody
}

type FlightCreateOrderV2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightCreateOrderV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightCreateOrderV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderV2Response) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightCreateOrderV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightCreateOrderV2Response) GetBody() *FlightCreateOrderV2ResponseBody {
	return s.Body
}

func (s *FlightCreateOrderV2Response) SetHeaders(v map[string]*string) *FlightCreateOrderV2Response {
	s.Headers = v
	return s
}

func (s *FlightCreateOrderV2Response) SetStatusCode(v int32) *FlightCreateOrderV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightCreateOrderV2Response) SetBody(v *FlightCreateOrderV2ResponseBody) *FlightCreateOrderV2Response {
	s.Body = v
	return s
}

func (s *FlightCreateOrderV2Response) Validate() error {
	return dara.Validate(s)
}
