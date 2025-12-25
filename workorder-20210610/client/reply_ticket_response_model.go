// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplyTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplyTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplyTicketResponse
	GetStatusCode() *int32
	SetBody(v *ReplyTicketResponseBody) *ReplyTicketResponse
	GetBody() *ReplyTicketResponseBody
}

type ReplyTicketResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplyTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplyTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplyTicketResponse) GoString() string {
	return s.String()
}

func (s *ReplyTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplyTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplyTicketResponse) GetBody() *ReplyTicketResponseBody {
	return s.Body
}

func (s *ReplyTicketResponse) SetHeaders(v map[string]*string) *ReplyTicketResponse {
	s.Headers = v
	return s
}

func (s *ReplyTicketResponse) SetStatusCode(v int32) *ReplyTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplyTicketResponse) SetBody(v *ReplyTicketResponseBody) *ReplyTicketResponse {
	s.Body = v
	return s
}

func (s *ReplyTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
