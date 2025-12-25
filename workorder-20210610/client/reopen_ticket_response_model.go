// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReopenTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReopenTicketResponse
	GetStatusCode() *int32
	SetBody(v *ReopenTicketResponseBody) *ReopenTicketResponse
	GetBody() *ReopenTicketResponseBody
}

type ReopenTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReopenTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReopenTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s ReopenTicketResponse) GoString() string {
	return s.String()
}

func (s *ReopenTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReopenTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReopenTicketResponse) GetBody() *ReopenTicketResponseBody {
	return s.Body
}

func (s *ReopenTicketResponse) SetHeaders(v map[string]*string) *ReopenTicketResponse {
	s.Headers = v
	return s
}

func (s *ReopenTicketResponse) SetStatusCode(v int32) *ReopenTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *ReopenTicketResponse) SetBody(v *ReopenTicketResponseBody) *ReopenTicketResponse {
	s.Body = v
	return s
}

func (s *ReopenTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
