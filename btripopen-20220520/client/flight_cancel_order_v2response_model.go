// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCancelOrderV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightCancelOrderV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightCancelOrderV2Response
	GetStatusCode() *int32
	SetBody(v *FlightCancelOrderV2ResponseBody) *FlightCancelOrderV2Response
	GetBody() *FlightCancelOrderV2ResponseBody
}

type FlightCancelOrderV2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightCancelOrderV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightCancelOrderV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightCancelOrderV2Response) GoString() string {
	return s.String()
}

func (s *FlightCancelOrderV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightCancelOrderV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightCancelOrderV2Response) GetBody() *FlightCancelOrderV2ResponseBody {
	return s.Body
}

func (s *FlightCancelOrderV2Response) SetHeaders(v map[string]*string) *FlightCancelOrderV2Response {
	s.Headers = v
	return s
}

func (s *FlightCancelOrderV2Response) SetStatusCode(v int32) *FlightCancelOrderV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightCancelOrderV2Response) SetBody(v *FlightCancelOrderV2ResponseBody) *FlightCancelOrderV2Response {
	s.Body = v
	return s
}

func (s *FlightCancelOrderV2Response) Validate() error {
	return dara.Validate(s)
}
