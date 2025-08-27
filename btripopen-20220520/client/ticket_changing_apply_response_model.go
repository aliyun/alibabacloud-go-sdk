// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingApplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketChangingApplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketChangingApplyResponse
	GetStatusCode() *int32
	SetBody(v *TicketChangingApplyResponseBody) *TicketChangingApplyResponse
	GetBody() *TicketChangingApplyResponseBody
}

type TicketChangingApplyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketChangingApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketChangingApplyResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingApplyResponse) GoString() string {
	return s.String()
}

func (s *TicketChangingApplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketChangingApplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketChangingApplyResponse) GetBody() *TicketChangingApplyResponseBody {
	return s.Body
}

func (s *TicketChangingApplyResponse) SetHeaders(v map[string]*string) *TicketChangingApplyResponse {
	s.Headers = v
	return s
}

func (s *TicketChangingApplyResponse) SetStatusCode(v int32) *TicketChangingApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketChangingApplyResponse) SetBody(v *TicketChangingApplyResponseBody) *TicketChangingApplyResponse {
	s.Body = v
	return s
}

func (s *TicketChangingApplyResponse) Validate() error {
	return dara.Validate(s)
}
