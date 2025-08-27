// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundApplyV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightRefundApplyV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightRefundApplyV2Response
	GetStatusCode() *int32
	SetBody(v *FlightRefundApplyV2ResponseBody) *FlightRefundApplyV2Response
	GetBody() *FlightRefundApplyV2ResponseBody
}

type FlightRefundApplyV2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightRefundApplyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightRefundApplyV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyV2Response) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightRefundApplyV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightRefundApplyV2Response) GetBody() *FlightRefundApplyV2ResponseBody {
	return s.Body
}

func (s *FlightRefundApplyV2Response) SetHeaders(v map[string]*string) *FlightRefundApplyV2Response {
	s.Headers = v
	return s
}

func (s *FlightRefundApplyV2Response) SetStatusCode(v int32) *FlightRefundApplyV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightRefundApplyV2Response) SetBody(v *FlightRefundApplyV2ResponseBody) *FlightRefundApplyV2Response {
	s.Body = v
	return s
}

func (s *FlightRefundApplyV2Response) Validate() error {
	return dara.Validate(s)
}
