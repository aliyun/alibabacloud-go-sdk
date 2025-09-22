// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketOrCredentialsSignInPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketOrCredentialsSignInPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketOrCredentialsSignInPopResponse
	GetStatusCode() *int32
	SetBody(v *TicketOrCredentialsSignInPopResponseBody) *TicketOrCredentialsSignInPopResponse
	GetBody() *TicketOrCredentialsSignInPopResponseBody
}

type TicketOrCredentialsSignInPopResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketOrCredentialsSignInPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketOrCredentialsSignInPopResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketOrCredentialsSignInPopResponse) GoString() string {
	return s.String()
}

func (s *TicketOrCredentialsSignInPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketOrCredentialsSignInPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketOrCredentialsSignInPopResponse) GetBody() *TicketOrCredentialsSignInPopResponseBody {
	return s.Body
}

func (s *TicketOrCredentialsSignInPopResponse) SetHeaders(v map[string]*string) *TicketOrCredentialsSignInPopResponse {
	s.Headers = v
	return s
}

func (s *TicketOrCredentialsSignInPopResponse) SetStatusCode(v int32) *TicketOrCredentialsSignInPopResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketOrCredentialsSignInPopResponse) SetBody(v *TicketOrCredentialsSignInPopResponseBody) *TicketOrCredentialsSignInPopResponse {
	s.Body = v
	return s
}

func (s *TicketOrCredentialsSignInPopResponse) Validate() error {
	return dara.Validate(s)
}
