// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTicketResponse
	GetStatusCode() *int32
	SetBody(v *GetTicketResponseBody) *GetTicketResponse
	GetBody() *GetTicketResponseBody
}

type GetTicketResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponse) GoString() string {
	return s.String()
}

func (s *GetTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTicketResponse) GetBody() *GetTicketResponseBody {
	return s.Body
}

func (s *GetTicketResponse) SetHeaders(v map[string]*string) *GetTicketResponse {
	s.Headers = v
	return s
}

func (s *GetTicketResponse) SetStatusCode(v int32) *GetTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTicketResponse) SetBody(v *GetTicketResponseBody) *GetTicketResponse {
	s.Body = v
	return s
}

func (s *GetTicketResponse) Validate() error {
	return dara.Validate(s)
}
