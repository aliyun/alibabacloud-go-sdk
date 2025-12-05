// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectOperationTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RejectOperationTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RejectOperationTicketResponse
	GetStatusCode() *int32
	SetBody(v *RejectOperationTicketResponseBody) *RejectOperationTicketResponse
	GetBody() *RejectOperationTicketResponseBody
}

type RejectOperationTicketResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectOperationTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectOperationTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s RejectOperationTicketResponse) GoString() string {
	return s.String()
}

func (s *RejectOperationTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RejectOperationTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RejectOperationTicketResponse) GetBody() *RejectOperationTicketResponseBody {
	return s.Body
}

func (s *RejectOperationTicketResponse) SetHeaders(v map[string]*string) *RejectOperationTicketResponse {
	s.Headers = v
	return s
}

func (s *RejectOperationTicketResponse) SetStatusCode(v int32) *RejectOperationTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectOperationTicketResponse) SetBody(v *RejectOperationTicketResponseBody) *RejectOperationTicketResponse {
	s.Body = v
	return s
}

func (s *RejectOperationTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
