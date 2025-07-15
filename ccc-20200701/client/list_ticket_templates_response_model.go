// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTicketTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTicketTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListTicketTemplatesResponseBody) *ListTicketTemplatesResponse
	GetBody() *ListTicketTemplatesResponseBody
}

type ListTicketTemplatesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTicketTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTicketTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTicketTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTicketTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTicketTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTicketTemplatesResponse) GetBody() *ListTicketTemplatesResponseBody {
	return s.Body
}

func (s *ListTicketTemplatesResponse) SetHeaders(v map[string]*string) *ListTicketTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTicketTemplatesResponse) SetStatusCode(v int32) *ListTicketTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTicketTemplatesResponse) SetBody(v *ListTicketTemplatesResponseBody) *ListTicketTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListTicketTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
