// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyOrderDetailV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightModifyOrderDetailV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightModifyOrderDetailV2Response
	GetStatusCode() *int32
	SetBody(v *FlightModifyOrderDetailV2ResponseBody) *FlightModifyOrderDetailV2Response
	GetBody() *FlightModifyOrderDetailV2ResponseBody
}

type FlightModifyOrderDetailV2Response struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightModifyOrderDetailV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightModifyOrderDetailV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOrderDetailV2Response) GoString() string {
	return s.String()
}

func (s *FlightModifyOrderDetailV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightModifyOrderDetailV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightModifyOrderDetailV2Response) GetBody() *FlightModifyOrderDetailV2ResponseBody {
	return s.Body
}

func (s *FlightModifyOrderDetailV2Response) SetHeaders(v map[string]*string) *FlightModifyOrderDetailV2Response {
	s.Headers = v
	return s
}

func (s *FlightModifyOrderDetailV2Response) SetStatusCode(v int32) *FlightModifyOrderDetailV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightModifyOrderDetailV2Response) SetBody(v *FlightModifyOrderDetailV2ResponseBody) *FlightModifyOrderDetailV2Response {
	s.Body = v
	return s
}

func (s *FlightModifyOrderDetailV2Response) Validate() error {
	return dara.Validate(s)
}
