// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyApplyV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightModifyApplyV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightModifyApplyV2Response
	GetStatusCode() *int32
	SetBody(v *FlightModifyApplyV2ResponseBody) *FlightModifyApplyV2Response
	GetBody() *FlightModifyApplyV2ResponseBody
}

type FlightModifyApplyV2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightModifyApplyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightModifyApplyV2Response) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyApplyV2Response) GoString() string {
	return s.String()
}

func (s *FlightModifyApplyV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightModifyApplyV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightModifyApplyV2Response) GetBody() *FlightModifyApplyV2ResponseBody {
	return s.Body
}

func (s *FlightModifyApplyV2Response) SetHeaders(v map[string]*string) *FlightModifyApplyV2Response {
	s.Headers = v
	return s
}

func (s *FlightModifyApplyV2Response) SetStatusCode(v int32) *FlightModifyApplyV2Response {
	s.StatusCode = &v
	return s
}

func (s *FlightModifyApplyV2Response) SetBody(v *FlightModifyApplyV2ResponseBody) *FlightModifyApplyV2Response {
	s.Body = v
	return s
}

func (s *FlightModifyApplyV2Response) Validate() error {
	return dara.Validate(s)
}
