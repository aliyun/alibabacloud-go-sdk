// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderListQueryV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightOrderListQueryV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightOrderListQueryV2Response
	GetStatusCode() *int32
	SetBody(v *FlightOrderListQueryV2ResponseBody) *FlightOrderListQueryV2Response
	GetBody() *FlightOrderListQueryV2ResponseBody
}

type FlightOrderListQueryV2Response struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightOrderListQueryV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightOrderListQueryV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2Response) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightOrderListQueryV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightOrderListQueryV2Response) GetBody() *FlightOrderListQueryV2ResponseBody {
	return s.Body
}

func (s *FlightOrderListQueryV2Response) SetHeaders(v map[string]*string) *FlightOrderListQueryV2Response {
	s.Headers = v
	return s
}

func (s *FlightOrderListQueryV2Response) SetStatusCode(v int32) *FlightOrderListQueryV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightOrderListQueryV2Response) SetBody(v *FlightOrderListQueryV2ResponseBody) *FlightOrderListQueryV2Response {
	s.Body = v
	return s
}

func (s *FlightOrderListQueryV2Response) Validate() error {
	return dara.Validate(s)
}
