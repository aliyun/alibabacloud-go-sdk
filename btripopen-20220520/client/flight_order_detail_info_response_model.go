// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderDetailInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightOrderDetailInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightOrderDetailInfoResponse
	GetStatusCode() *int32
	SetBody(v *FlightOrderDetailInfoResponseBody) *FlightOrderDetailInfoResponse
	GetBody() *FlightOrderDetailInfoResponseBody
}

type FlightOrderDetailInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightOrderDetailInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightOrderDetailInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderDetailInfoResponse) GoString() string {
	return s.String()
}

func (s *FlightOrderDetailInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightOrderDetailInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightOrderDetailInfoResponse) GetBody() *FlightOrderDetailInfoResponseBody {
	return s.Body
}

func (s *FlightOrderDetailInfoResponse) SetHeaders(v map[string]*string) *FlightOrderDetailInfoResponse {
	s.Headers = v
	return s
}

func (s *FlightOrderDetailInfoResponse) SetStatusCode(v int32) *FlightOrderDetailInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightOrderDetailInfoResponse) SetBody(v *FlightOrderDetailInfoResponseBody) *FlightOrderDetailInfoResponse {
	s.Body = v
	return s
}

func (s *FlightOrderDetailInfoResponse) Validate() error {
	return dara.Validate(s)
}
