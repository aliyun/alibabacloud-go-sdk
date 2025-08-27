// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingPayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketChangingPayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketChangingPayResponse
	GetStatusCode() *int32
	SetBody(v *TicketChangingPayResponseBody) *TicketChangingPayResponse
	GetBody() *TicketChangingPayResponseBody
}

type TicketChangingPayResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketChangingPayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketChangingPayResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingPayResponse) GoString() string {
	return s.String()
}

func (s *TicketChangingPayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketChangingPayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketChangingPayResponse) GetBody() *TicketChangingPayResponseBody {
	return s.Body
}

func (s *TicketChangingPayResponse) SetHeaders(v map[string]*string) *TicketChangingPayResponse {
	s.Headers = v
	return s
}

func (s *TicketChangingPayResponse) SetStatusCode(v int32) *TicketChangingPayResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketChangingPayResponse) SetBody(v *TicketChangingPayResponseBody) *TicketChangingPayResponse {
	s.Body = v
	return s
}

func (s *TicketChangingPayResponse) Validate() error {
	return dara.Validate(s)
}
