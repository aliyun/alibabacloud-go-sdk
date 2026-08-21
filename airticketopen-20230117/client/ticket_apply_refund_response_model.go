// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketApplyRefundResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketApplyRefundResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketApplyRefundResponse
	GetStatusCode() *int32
	SetBody(v *TicketApplyRefundResponseBody) *TicketApplyRefundResponse
	GetBody() *TicketApplyRefundResponseBody
}

type TicketApplyRefundResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketApplyRefundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketApplyRefundResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketApplyRefundResponse) GoString() string {
	return s.String()
}

func (s *TicketApplyRefundResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketApplyRefundResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketApplyRefundResponse) GetBody() *TicketApplyRefundResponseBody {
	return s.Body
}

func (s *TicketApplyRefundResponse) SetHeaders(v map[string]*string) *TicketApplyRefundResponse {
	s.Headers = v
	return s
}

func (s *TicketApplyRefundResponse) SetStatusCode(v int32) *TicketApplyRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketApplyRefundResponse) SetBody(v *TicketApplyRefundResponseBody) *TicketApplyRefundResponse {
	s.Body = v
	return s
}

func (s *TicketApplyRefundResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
