// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminateTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminateTicketResponse
	GetStatusCode() *int32
	SetBody(v *TerminateTicketResponseBody) *TerminateTicketResponse
	GetBody() *TerminateTicketResponseBody
}

type TerminateTicketResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminateTicketResponse) GoString() string {
	return s.String()
}

func (s *TerminateTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminateTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminateTicketResponse) GetBody() *TerminateTicketResponseBody {
	return s.Body
}

func (s *TerminateTicketResponse) SetHeaders(v map[string]*string) *TerminateTicketResponse {
	s.Headers = v
	return s
}

func (s *TerminateTicketResponse) SetStatusCode(v int32) *TerminateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateTicketResponse) SetBody(v *TerminateTicketResponseBody) *TerminateTicketResponse {
	s.Body = v
	return s
}

func (s *TerminateTicketResponse) Validate() error {
	return dara.Validate(s)
}
