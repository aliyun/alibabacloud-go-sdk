// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyCancelV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightModifyCancelV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightModifyCancelV2Response
	GetStatusCode() *int32
	SetBody(v *FlightModifyCancelV2ResponseBody) *FlightModifyCancelV2Response
	GetBody() *FlightModifyCancelV2ResponseBody
}

type FlightModifyCancelV2Response struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightModifyCancelV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightModifyCancelV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyCancelV2Response) GoString() string {
	return s.String()
}

func (s *FlightModifyCancelV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightModifyCancelV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightModifyCancelV2Response) GetBody() *FlightModifyCancelV2ResponseBody {
	return s.Body
}

func (s *FlightModifyCancelV2Response) SetHeaders(v map[string]*string) *FlightModifyCancelV2Response {
	s.Headers = v
	return s
}

func (s *FlightModifyCancelV2Response) SetStatusCode(v int32) *FlightModifyCancelV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightModifyCancelV2Response) SetBody(v *FlightModifyCancelV2ResponseBody) *FlightModifyCancelV2Response {
	s.Body = v
	return s
}

func (s *FlightModifyCancelV2Response) Validate() error {
	return dara.Validate(s)
}
