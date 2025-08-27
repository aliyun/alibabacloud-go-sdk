// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundDetailV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightRefundDetailV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightRefundDetailV2Response
	GetStatusCode() *int32
	SetBody(v *FlightRefundDetailV2ResponseBody) *FlightRefundDetailV2Response
	GetBody() *FlightRefundDetailV2ResponseBody
}

type FlightRefundDetailV2Response struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightRefundDetailV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightRefundDetailV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailV2Response) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightRefundDetailV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightRefundDetailV2Response) GetBody() *FlightRefundDetailV2ResponseBody {
	return s.Body
}

func (s *FlightRefundDetailV2Response) SetHeaders(v map[string]*string) *FlightRefundDetailV2Response {
	s.Headers = v
	return s
}

func (s *FlightRefundDetailV2Response) SetStatusCode(v int32) *FlightRefundDetailV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightRefundDetailV2Response) SetBody(v *FlightRefundDetailV2ResponseBody) *FlightRefundDetailV2Response {
	s.Body = v
	return s
}

func (s *FlightRefundDetailV2Response) Validate() error {
	return dara.Validate(s)
}
