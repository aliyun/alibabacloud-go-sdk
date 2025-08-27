// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingEnquiryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketChangingEnquiryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketChangingEnquiryResponse
	GetStatusCode() *int32
	SetBody(v *TicketChangingEnquiryResponseBody) *TicketChangingEnquiryResponse
	GetBody() *TicketChangingEnquiryResponseBody
}

type TicketChangingEnquiryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketChangingEnquiryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketChangingEnquiryResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryResponse) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketChangingEnquiryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketChangingEnquiryResponse) GetBody() *TicketChangingEnquiryResponseBody {
	return s.Body
}

func (s *TicketChangingEnquiryResponse) SetHeaders(v map[string]*string) *TicketChangingEnquiryResponse {
	s.Headers = v
	return s
}

func (s *TicketChangingEnquiryResponse) SetStatusCode(v int32) *TicketChangingEnquiryResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketChangingEnquiryResponse) SetBody(v *TicketChangingEnquiryResponseBody) *TicketChangingEnquiryResponse {
	s.Body = v
	return s
}

func (s *TicketChangingEnquiryResponse) Validate() error {
	return dara.Validate(s)
}
