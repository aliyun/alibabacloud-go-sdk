// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderDetailV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightOrderDetailV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightOrderDetailV2Response
	GetStatusCode() *int32
	SetBody(v *FlightOrderDetailV2ResponseBody) *FlightOrderDetailV2Response
	GetBody() *FlightOrderDetailV2ResponseBody
}

type FlightOrderDetailV2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightOrderDetailV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightOrderDetailV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailV2Response) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightOrderDetailV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightOrderDetailV2Response) GetBody() *FlightOrderDetailV2ResponseBody {
	return s.Body
}

func (s *FlightOrderDetailV2Response) SetHeaders(v map[string]*string) *FlightOrderDetailV2Response {
	s.Headers = v
	return s
}

func (s *FlightOrderDetailV2Response) SetStatusCode(v int32) *FlightOrderDetailV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightOrderDetailV2Response) SetBody(v *FlightOrderDetailV2ResponseBody) *FlightOrderDetailV2Response {
	s.Body = v
	return s
}

func (s *FlightOrderDetailV2Response) Validate() error {
	return dara.Validate(s)
}
