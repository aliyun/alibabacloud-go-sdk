// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishTicketTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FinishTicketTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FinishTicketTaskResponse
	GetStatusCode() *int32
	SetBody(v *FinishTicketTaskResponseBody) *FinishTicketTaskResponse
	GetBody() *FinishTicketTaskResponseBody
}

type FinishTicketTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishTicketTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishTicketTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketTaskResponse) GoString() string {
	return s.String()
}

func (s *FinishTicketTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FinishTicketTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FinishTicketTaskResponse) GetBody() *FinishTicketTaskResponseBody {
	return s.Body
}

func (s *FinishTicketTaskResponse) SetHeaders(v map[string]*string) *FinishTicketTaskResponse {
	s.Headers = v
	return s
}

func (s *FinishTicketTaskResponse) SetStatusCode(v int32) *FinishTicketTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishTicketTaskResponse) SetBody(v *FinishTicketTaskResponseBody) *FinishTicketTaskResponse {
	s.Body = v
	return s
}

func (s *FinishTicketTaskResponse) Validate() error {
	return dara.Validate(s)
}
