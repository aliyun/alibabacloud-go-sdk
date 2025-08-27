// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingFlightListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketChangingFlightListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketChangingFlightListResponse
	GetStatusCode() *int32
	SetBody(v *TicketChangingFlightListResponseBody) *TicketChangingFlightListResponse
	GetBody() *TicketChangingFlightListResponseBody
}

type TicketChangingFlightListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketChangingFlightListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketChangingFlightListResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListResponse) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketChangingFlightListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketChangingFlightListResponse) GetBody() *TicketChangingFlightListResponseBody {
	return s.Body
}

func (s *TicketChangingFlightListResponse) SetHeaders(v map[string]*string) *TicketChangingFlightListResponse {
	s.Headers = v
	return s
}

func (s *TicketChangingFlightListResponse) SetStatusCode(v int32) *TicketChangingFlightListResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketChangingFlightListResponse) SetBody(v *TicketChangingFlightListResponseBody) *TicketChangingFlightListResponse {
	s.Body = v
	return s
}

func (s *TicketChangingFlightListResponse) Validate() error {
	return dara.Validate(s)
}
