// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferTicketResponse
	GetStatusCode() *int32
	SetBody(v *TransferTicketResponseBody) *TransferTicketResponse
	GetBody() *TransferTicketResponseBody
}

type TransferTicketResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketResponse) GoString() string {
	return s.String()
}

func (s *TransferTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferTicketResponse) GetBody() *TransferTicketResponseBody {
	return s.Body
}

func (s *TransferTicketResponse) SetHeaders(v map[string]*string) *TransferTicketResponse {
	s.Headers = v
	return s
}

func (s *TransferTicketResponse) SetStatusCode(v int32) *TransferTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferTicketResponse) SetBody(v *TransferTicketResponseBody) *TransferTicketResponse {
	s.Body = v
	return s
}

func (s *TransferTicketResponse) Validate() error {
	return dara.Validate(s)
}
