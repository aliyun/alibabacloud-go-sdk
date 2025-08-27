// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightExceedApplyQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlightExceedApplyQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlightExceedApplyQueryResponse
	GetStatusCode() *int32
	SetBody(v *FlightExceedApplyQueryResponseBody) *FlightExceedApplyQueryResponse
	GetBody() *FlightExceedApplyQueryResponseBody
}

type FlightExceedApplyQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightExceedApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightExceedApplyQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s FlightExceedApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlightExceedApplyQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlightExceedApplyQueryResponse) GetBody() *FlightExceedApplyQueryResponseBody {
	return s.Body
}

func (s *FlightExceedApplyQueryResponse) SetHeaders(v map[string]*string) *FlightExceedApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *FlightExceedApplyQueryResponse) SetStatusCode(v int32) *FlightExceedApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightExceedApplyQueryResponse) SetBody(v *FlightExceedApplyQueryResponseBody) *FlightExceedApplyQueryResponse {
	s.Body = v
	return s
}

func (s *FlightExceedApplyQueryResponse) Validate() error {
	return dara.Validate(s)
}
