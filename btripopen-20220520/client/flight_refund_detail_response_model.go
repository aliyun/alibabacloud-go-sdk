// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightRefundDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightRefundDetailResponse
	GetStatusCode() *int32
	SetBody(v *FlightRefundDetailResponseBody) *FlightRefundDetailResponse
	GetBody() *FlightRefundDetailResponseBody
}

type FlightRefundDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightRefundDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightRefundDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailResponse) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightRefundDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightRefundDetailResponse) GetBody() *FlightRefundDetailResponseBody {
	return s.Body
}

func (s *FlightRefundDetailResponse) SetHeaders(v map[string]*string) *FlightRefundDetailResponse {
	s.Headers = v
	return s
}

func (s *FlightRefundDetailResponse) SetStatusCode(v int32) *FlightRefundDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightRefundDetailResponse) SetBody(v *FlightRefundDetailResponseBody) *FlightRefundDetailResponse {
	s.Body = v
	return s
}

func (s *FlightRefundDetailResponse) Validate() error {
	return dara.Validate(s)
}
