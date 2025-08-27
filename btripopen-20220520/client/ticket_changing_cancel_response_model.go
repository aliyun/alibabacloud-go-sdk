// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingCancelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketChangingCancelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketChangingCancelResponse
	GetStatusCode() *int32
	SetBody(v *TicketChangingCancelResponseBody) *TicketChangingCancelResponse
	GetBody() *TicketChangingCancelResponseBody
}

type TicketChangingCancelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketChangingCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketChangingCancelResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingCancelResponse) GoString() string {
	return s.String()
}

func (s *TicketChangingCancelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketChangingCancelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketChangingCancelResponse) GetBody() *TicketChangingCancelResponseBody {
	return s.Body
}

func (s *TicketChangingCancelResponse) SetHeaders(v map[string]*string) *TicketChangingCancelResponse {
	s.Headers = v
	return s
}

func (s *TicketChangingCancelResponse) SetStatusCode(v int32) *TicketChangingCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketChangingCancelResponse) SetBody(v *TicketChangingCancelResponseBody) *TicketChangingCancelResponse {
	s.Body = v
	return s
}

func (s *TicketChangingCancelResponse) Validate() error {
	return dara.Validate(s)
}
