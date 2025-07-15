// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WithdrawTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WithdrawTicketResponse
	GetStatusCode() *int32
	SetBody(v *WithdrawTicketResponseBody) *WithdrawTicketResponse
	GetBody() *WithdrawTicketResponseBody
}

type WithdrawTicketResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WithdrawTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WithdrawTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s WithdrawTicketResponse) GoString() string {
	return s.String()
}

func (s *WithdrawTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WithdrawTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WithdrawTicketResponse) GetBody() *WithdrawTicketResponseBody {
	return s.Body
}

func (s *WithdrawTicketResponse) SetHeaders(v map[string]*string) *WithdrawTicketResponse {
	s.Headers = v
	return s
}

func (s *WithdrawTicketResponse) SetStatusCode(v int32) *WithdrawTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *WithdrawTicketResponse) SetBody(v *WithdrawTicketResponseBody) *WithdrawTicketResponse {
	s.Body = v
	return s
}

func (s *WithdrawTicketResponse) Validate() error {
	return dara.Validate(s)
}
