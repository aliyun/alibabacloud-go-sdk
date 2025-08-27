// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaItemDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightOtaItemDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightOtaItemDetailResponse
	GetStatusCode() *int32
	SetBody(v *FlightOtaItemDetailResponseBody) *FlightOtaItemDetailResponse
	GetBody() *FlightOtaItemDetailResponseBody
}

type FlightOtaItemDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightOtaItemDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightOtaItemDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaItemDetailResponse) GoString() string {
	return s.String()
}

func (s *FlightOtaItemDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightOtaItemDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightOtaItemDetailResponse) GetBody() *FlightOtaItemDetailResponseBody {
	return s.Body
}

func (s *FlightOtaItemDetailResponse) SetHeaders(v map[string]*string) *FlightOtaItemDetailResponse {
	s.Headers = v
	return s
}

func (s *FlightOtaItemDetailResponse) SetStatusCode(v int32) *FlightOtaItemDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightOtaItemDetailResponse) SetBody(v *FlightOtaItemDetailResponseBody) *FlightOtaItemDetailResponse {
	s.Body = v
	return s
}

func (s *FlightOtaItemDetailResponse) Validate() error {
	return dara.Validate(s)
}
