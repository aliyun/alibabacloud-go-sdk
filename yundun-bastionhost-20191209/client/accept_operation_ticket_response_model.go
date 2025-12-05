// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptOperationTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcceptOperationTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcceptOperationTicketResponse
	GetStatusCode() *int32
	SetBody(v *AcceptOperationTicketResponseBody) *AcceptOperationTicketResponse
	GetBody() *AcceptOperationTicketResponseBody
}

type AcceptOperationTicketResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptOperationTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptOperationTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s AcceptOperationTicketResponse) GoString() string {
	return s.String()
}

func (s *AcceptOperationTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcceptOperationTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcceptOperationTicketResponse) GetBody() *AcceptOperationTicketResponseBody {
	return s.Body
}

func (s *AcceptOperationTicketResponse) SetHeaders(v map[string]*string) *AcceptOperationTicketResponse {
	s.Headers = v
	return s
}

func (s *AcceptOperationTicketResponse) SetStatusCode(v int32) *AcceptOperationTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptOperationTicketResponse) SetBody(v *AcceptOperationTicketResponseBody) *AcceptOperationTicketResponse {
	s.Body = v
	return s
}

func (s *AcceptOperationTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
