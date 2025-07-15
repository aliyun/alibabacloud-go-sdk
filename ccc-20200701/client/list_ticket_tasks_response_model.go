// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTicketTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTicketTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListTicketTasksResponseBody) *ListTicketTasksResponse
	GetBody() *ListTicketTasksResponseBody
}

type ListTicketTasksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTicketTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTicketTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTicketTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTicketTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTicketTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTicketTasksResponse) GetBody() *ListTicketTasksResponseBody {
	return s.Body
}

func (s *ListTicketTasksResponse) SetHeaders(v map[string]*string) *ListTicketTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTicketTasksResponse) SetStatusCode(v int32) *ListTicketTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTicketTasksResponse) SetBody(v *ListTicketTasksResponseBody) *ListTicketTasksResponse {
	s.Body = v
	return s
}

func (s *ListTicketTasksResponse) Validate() error {
	return dara.Validate(s)
}
