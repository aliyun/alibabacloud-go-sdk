// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundPreCalV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightRefundPreCalV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightRefundPreCalV2Response
	GetStatusCode() *int32
	SetBody(v *FlightRefundPreCalV2ResponseBody) *FlightRefundPreCalV2Response
	GetBody() *FlightRefundPreCalV2ResponseBody
}

type FlightRefundPreCalV2Response struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightRefundPreCalV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightRefundPreCalV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalV2Response) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightRefundPreCalV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightRefundPreCalV2Response) GetBody() *FlightRefundPreCalV2ResponseBody {
	return s.Body
}

func (s *FlightRefundPreCalV2Response) SetHeaders(v map[string]*string) *FlightRefundPreCalV2Response {
	s.Headers = v
	return s
}

func (s *FlightRefundPreCalV2Response) SetStatusCode(v int32) *FlightRefundPreCalV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightRefundPreCalV2Response) SetBody(v *FlightRefundPreCalV2ResponseBody) *FlightRefundPreCalV2Response {
	s.Body = v
	return s
}

func (s *FlightRefundPreCalV2Response) Validate() error {
	return dara.Validate(s)
}
