// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FinishTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FinishTicketResponse
	GetStatusCode() *int32
	SetBody(v *FinishTicketResponseBody) *FinishTicketResponse
	GetBody() *FinishTicketResponseBody
}

type FinishTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketResponse) GoString() string {
	return s.String()
}

func (s *FinishTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FinishTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FinishTicketResponse) GetBody() *FinishTicketResponseBody {
	return s.Body
}

func (s *FinishTicketResponse) SetHeaders(v map[string]*string) *FinishTicketResponse {
	s.Headers = v
	return s
}

func (s *FinishTicketResponse) SetStatusCode(v int32) *FinishTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishTicketResponse) SetBody(v *FinishTicketResponseBody) *FinishTicketResponse {
	s.Body = v
	return s
}

func (s *FinishTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
