// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightOrderQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightOrderQueryResponse
	GetStatusCode() *int32
	SetBody(v *FlightOrderQueryResponseBody) *FlightOrderQueryResponse
	GetBody() *FlightOrderQueryResponseBody
}

type FlightOrderQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightOrderQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightOrderQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *FlightOrderQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightOrderQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightOrderQueryResponse) GetBody() *FlightOrderQueryResponseBody {
	return s.Body
}

func (s *FlightOrderQueryResponse) SetHeaders(v map[string]*string) *FlightOrderQueryResponse {
	s.Headers = v
	return s
}

func (s *FlightOrderQueryResponse) SetStatusCode(v int32) *FlightOrderQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightOrderQueryResponse) SetBody(v *FlightOrderQueryResponseBody) *FlightOrderQueryResponse {
	s.Body = v
	return s
}

func (s *FlightOrderQueryResponse) Validate() error {
	return dara.Validate(s)
}
