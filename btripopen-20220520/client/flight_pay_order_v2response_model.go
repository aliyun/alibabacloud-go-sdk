// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightPayOrderV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightPayOrderV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightPayOrderV2Response
	GetStatusCode() *int32
	SetBody(v *FlightPayOrderV2ResponseBody) *FlightPayOrderV2Response
	GetBody() *FlightPayOrderV2ResponseBody
}

type FlightPayOrderV2Response struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightPayOrderV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightPayOrderV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightPayOrderV2Response) GoString() string {
	return s.String()
}

func (s *FlightPayOrderV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightPayOrderV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightPayOrderV2Response) GetBody() *FlightPayOrderV2ResponseBody {
	return s.Body
}

func (s *FlightPayOrderV2Response) SetHeaders(v map[string]*string) *FlightPayOrderV2Response {
	s.Headers = v
	return s
}

func (s *FlightPayOrderV2Response) SetStatusCode(v int32) *FlightPayOrderV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightPayOrderV2Response) SetBody(v *FlightPayOrderV2ResponseBody) *FlightPayOrderV2Response {
	s.Body = v
	return s
}

func (s *FlightPayOrderV2Response) Validate() error {
	return dara.Validate(s)
}
