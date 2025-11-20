// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignTicketResponse
	GetStatusCode() *int32
	SetBody(v *AssignTicketResponseBody) *AssignTicketResponse
	GetBody() *AssignTicketResponseBody
}

type AssignTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignTicketResponse) GoString() string {
	return s.String()
}

func (s *AssignTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignTicketResponse) GetBody() *AssignTicketResponseBody {
	return s.Body
}

func (s *AssignTicketResponse) SetHeaders(v map[string]*string) *AssignTicketResponse {
	s.Headers = v
	return s
}

func (s *AssignTicketResponse) SetStatusCode(v int32) *AssignTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignTicketResponse) SetBody(v *AssignTicketResponseBody) *AssignTicketResponse {
	s.Body = v
	return s
}

func (s *AssignTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
