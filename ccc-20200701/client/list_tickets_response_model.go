// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTicketsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTicketsResponse
	GetStatusCode() *int32
	SetBody(v *ListTicketsResponseBody) *ListTicketsResponse
	GetBody() *ListTicketsResponseBody
}

type ListTicketsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTicketsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponse) GoString() string {
	return s.String()
}

func (s *ListTicketsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTicketsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTicketsResponse) GetBody() *ListTicketsResponseBody {
	return s.Body
}

func (s *ListTicketsResponse) SetHeaders(v map[string]*string) *ListTicketsResponse {
	s.Headers = v
	return s
}

func (s *ListTicketsResponse) SetStatusCode(v int32) *ListTicketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTicketsResponse) SetBody(v *ListTicketsResponseBody) *ListTicketsResponse {
	s.Body = v
	return s
}

func (s *ListTicketsResponse) Validate() error {
	return dara.Validate(s)
}
