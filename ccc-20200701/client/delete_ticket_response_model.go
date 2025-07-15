// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTicketResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTicketResponseBody) *DeleteTicketResponse
	GetBody() *DeleteTicketResponseBody
}

type DeleteTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTicketResponse) GoString() string {
	return s.String()
}

func (s *DeleteTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTicketResponse) GetBody() *DeleteTicketResponseBody {
	return s.Body
}

func (s *DeleteTicketResponse) SetHeaders(v map[string]*string) *DeleteTicketResponse {
	s.Headers = v
	return s
}

func (s *DeleteTicketResponse) SetStatusCode(v int32) *DeleteTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTicketResponse) SetBody(v *DeleteTicketResponseBody) *DeleteTicketResponse {
	s.Body = v
	return s
}

func (s *DeleteTicketResponse) Validate() error {
	return dara.Validate(s)
}
