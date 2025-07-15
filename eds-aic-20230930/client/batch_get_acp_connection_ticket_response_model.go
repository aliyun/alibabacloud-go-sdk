// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetAcpConnectionTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetAcpConnectionTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetAcpConnectionTicketResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetAcpConnectionTicketResponseBody) *BatchGetAcpConnectionTicketResponse
	GetBody() *BatchGetAcpConnectionTicketResponseBody
}

type BatchGetAcpConnectionTicketResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetAcpConnectionTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetAcpConnectionTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetAcpConnectionTicketResponse) GoString() string {
	return s.String()
}

func (s *BatchGetAcpConnectionTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetAcpConnectionTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetAcpConnectionTicketResponse) GetBody() *BatchGetAcpConnectionTicketResponseBody {
	return s.Body
}

func (s *BatchGetAcpConnectionTicketResponse) SetHeaders(v map[string]*string) *BatchGetAcpConnectionTicketResponse {
	s.Headers = v
	return s
}

func (s *BatchGetAcpConnectionTicketResponse) SetStatusCode(v int32) *BatchGetAcpConnectionTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetAcpConnectionTicketResponse) SetBody(v *BatchGetAcpConnectionTicketResponseBody) *BatchGetAcpConnectionTicketResponse {
	s.Body = v
	return s
}

func (s *BatchGetAcpConnectionTicketResponse) Validate() error {
	return dara.Validate(s)
}
