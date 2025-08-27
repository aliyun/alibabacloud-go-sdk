// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderListQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightOrderListQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightOrderListQueryResponse
	GetStatusCode() *int32
	SetBody(v *FlightOrderListQueryResponseBody) *FlightOrderListQueryResponse
	GetBody() *FlightOrderListQueryResponseBody
}

type FlightOrderListQueryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightOrderListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightOrderListQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryResponse) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightOrderListQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightOrderListQueryResponse) GetBody() *FlightOrderListQueryResponseBody {
	return s.Body
}

func (s *FlightOrderListQueryResponse) SetHeaders(v map[string]*string) *FlightOrderListQueryResponse {
	s.Headers = v
	return s
}

func (s *FlightOrderListQueryResponse) SetStatusCode(v int32) *FlightOrderListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightOrderListQueryResponse) SetBody(v *FlightOrderListQueryResponseBody) *FlightOrderListQueryResponse {
	s.Body = v
	return s
}

func (s *FlightOrderListQueryResponse) Validate() error {
	return dara.Validate(s)
}
