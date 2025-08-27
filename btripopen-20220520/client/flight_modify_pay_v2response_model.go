// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyPayV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightModifyPayV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightModifyPayV2Response
	GetStatusCode() *int32
	SetBody(v *FlightModifyPayV2ResponseBody) *FlightModifyPayV2Response
	GetBody() *FlightModifyPayV2ResponseBody
}

type FlightModifyPayV2Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightModifyPayV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightModifyPayV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyPayV2Response) GoString() string {
	return s.String()
}

func (s *FlightModifyPayV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightModifyPayV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightModifyPayV2Response) GetBody() *FlightModifyPayV2ResponseBody {
	return s.Body
}

func (s *FlightModifyPayV2Response) SetHeaders(v map[string]*string) *FlightModifyPayV2Response {
	s.Headers = v
	return s
}

func (s *FlightModifyPayV2Response) SetStatusCode(v int32) *FlightModifyPayV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightModifyPayV2Response) SetBody(v *FlightModifyPayV2ResponseBody) *FlightModifyPayV2Response {
	s.Body = v
	return s
}

func (s *FlightModifyPayV2Response) Validate() error {
	return dara.Validate(s)
}
