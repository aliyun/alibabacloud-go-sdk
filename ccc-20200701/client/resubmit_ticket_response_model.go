// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResubmitTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResubmitTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResubmitTicketResponse
	GetStatusCode() *int32
	SetBody(v *ResubmitTicketResponseBody) *ResubmitTicketResponse
	GetBody() *ResubmitTicketResponseBody
}

type ResubmitTicketResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResubmitTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResubmitTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s ResubmitTicketResponse) GoString() string {
	return s.String()
}

func (s *ResubmitTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResubmitTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResubmitTicketResponse) GetBody() *ResubmitTicketResponseBody {
	return s.Body
}

func (s *ResubmitTicketResponse) SetHeaders(v map[string]*string) *ResubmitTicketResponse {
	s.Headers = v
	return s
}

func (s *ResubmitTicketResponse) SetStatusCode(v int32) *ResubmitTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *ResubmitTicketResponse) SetBody(v *ResubmitTicketResponseBody) *ResubmitTicketResponse {
	s.Body = v
	return s
}

func (s *ResubmitTicketResponse) Validate() error {
	return dara.Validate(s)
}
