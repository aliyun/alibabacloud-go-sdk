// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTicketTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferTicketTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferTicketTaskResponse
	GetStatusCode() *int32
	SetBody(v *TransferTicketTaskResponseBody) *TransferTicketTaskResponse
	GetBody() *TransferTicketTaskResponseBody
}

type TransferTicketTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferTicketTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferTicketTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketTaskResponse) GoString() string {
	return s.String()
}

func (s *TransferTicketTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferTicketTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferTicketTaskResponse) GetBody() *TransferTicketTaskResponseBody {
	return s.Body
}

func (s *TransferTicketTaskResponse) SetHeaders(v map[string]*string) *TransferTicketTaskResponse {
	s.Headers = v
	return s
}

func (s *TransferTicketTaskResponse) SetStatusCode(v int32) *TransferTicketTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferTicketTaskResponse) SetBody(v *TransferTicketTaskResponseBody) *TransferTicketTaskResponse {
	s.Body = v
	return s
}

func (s *TransferTicketTaskResponse) Validate() error {
	return dara.Validate(s)
}
