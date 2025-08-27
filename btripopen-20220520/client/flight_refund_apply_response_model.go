// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundApplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightRefundApplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightRefundApplyResponse
	GetStatusCode() *int32
	SetBody(v *FlightRefundApplyResponseBody) *FlightRefundApplyResponse
	GetBody() *FlightRefundApplyResponseBody
}

type FlightRefundApplyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightRefundApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightRefundApplyResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyResponse) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightRefundApplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightRefundApplyResponse) GetBody() *FlightRefundApplyResponseBody {
	return s.Body
}

func (s *FlightRefundApplyResponse) SetHeaders(v map[string]*string) *FlightRefundApplyResponse {
	s.Headers = v
	return s
}

func (s *FlightRefundApplyResponse) SetStatusCode(v int32) *FlightRefundApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightRefundApplyResponse) SetBody(v *FlightRefundApplyResponseBody) *FlightRefundApplyResponse {
	s.Body = v
	return s
}

func (s *FlightRefundApplyResponse) Validate() error {
	return dara.Validate(s)
}
