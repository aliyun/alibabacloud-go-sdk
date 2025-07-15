// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTicketResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTicketResponseBody) *UpdateTicketResponse
	GetBody() *UpdateTicketResponseBody
}

type UpdateTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketResponse) GoString() string {
	return s.String()
}

func (s *UpdateTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTicketResponse) GetBody() *UpdateTicketResponseBody {
	return s.Body
}

func (s *UpdateTicketResponse) SetHeaders(v map[string]*string) *UpdateTicketResponse {
	s.Headers = v
	return s
}

func (s *UpdateTicketResponse) SetStatusCode(v int32) *UpdateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTicketResponse) SetBody(v *UpdateTicketResponseBody) *UpdateTicketResponse {
	s.Body = v
	return s
}

func (s *UpdateTicketResponse) Validate() error {
	return dara.Validate(s)
}
