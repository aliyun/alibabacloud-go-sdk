// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightSearchListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightSearchListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightSearchListResponse
	GetStatusCode() *int32
	SetBody(v *FlightSearchListResponseBody) *FlightSearchListResponse
	GetBody() *FlightSearchListResponseBody
}

type FlightSearchListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightSearchListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightSearchListResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListResponse) GoString() string {
	return s.String()
}

func (s *FlightSearchListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightSearchListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightSearchListResponse) GetBody() *FlightSearchListResponseBody {
	return s.Body
}

func (s *FlightSearchListResponse) SetHeaders(v map[string]*string) *FlightSearchListResponse {
	s.Headers = v
	return s
}

func (s *FlightSearchListResponse) SetStatusCode(v int32) *FlightSearchListResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightSearchListResponse) SetBody(v *FlightSearchListResponseBody) *FlightSearchListResponse {
	s.Body = v
	return s
}

func (s *FlightSearchListResponse) Validate() error {
	return dara.Validate(s)
}
