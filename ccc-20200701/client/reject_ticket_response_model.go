// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RejectTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RejectTicketResponse
	GetStatusCode() *int32
	SetBody(v *RejectTicketResponseBody) *RejectTicketResponse
	GetBody() *RejectTicketResponseBody
}

type RejectTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s RejectTicketResponse) GoString() string {
	return s.String()
}

func (s *RejectTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RejectTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RejectTicketResponse) GetBody() *RejectTicketResponseBody {
	return s.Body
}

func (s *RejectTicketResponse) SetHeaders(v map[string]*string) *RejectTicketResponse {
	s.Headers = v
	return s
}

func (s *RejectTicketResponse) SetStatusCode(v int32) *RejectTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectTicketResponse) SetBody(v *RejectTicketResponseBody) *RejectTicketResponse {
	s.Body = v
	return s
}

func (s *RejectTicketResponse) Validate() error {
	return dara.Validate(s)
}
