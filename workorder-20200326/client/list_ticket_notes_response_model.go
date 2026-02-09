// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketNotesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTicketNotesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTicketNotesResponse
	GetStatusCode() *int32
	SetBody(v *ListTicketNotesResponseBody) *ListTicketNotesResponse
	GetBody() *ListTicketNotesResponseBody
}

type ListTicketNotesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTicketNotesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTicketNotesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTicketNotesResponse) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTicketNotesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTicketNotesResponse) GetBody() *ListTicketNotesResponseBody {
	return s.Body
}

func (s *ListTicketNotesResponse) SetHeaders(v map[string]*string) *ListTicketNotesResponse {
	s.Headers = v
	return s
}

func (s *ListTicketNotesResponse) SetStatusCode(v int32) *ListTicketNotesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTicketNotesResponse) SetBody(v *ListTicketNotesResponseBody) *ListTicketNotesResponse {
	s.Body = v
	return s
}

func (s *ListTicketNotesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
