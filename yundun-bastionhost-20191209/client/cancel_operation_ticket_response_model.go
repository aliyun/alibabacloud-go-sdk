// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOperationTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelOperationTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelOperationTicketResponse
	GetStatusCode() *int32
	SetBody(v *CancelOperationTicketResponseBody) *CancelOperationTicketResponse
	GetBody() *CancelOperationTicketResponseBody
}

type CancelOperationTicketResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOperationTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelOperationTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelOperationTicketResponse) GoString() string {
	return s.String()
}

func (s *CancelOperationTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelOperationTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelOperationTicketResponse) GetBody() *CancelOperationTicketResponseBody {
	return s.Body
}

func (s *CancelOperationTicketResponse) SetHeaders(v map[string]*string) *CancelOperationTicketResponse {
	s.Headers = v
	return s
}

func (s *CancelOperationTicketResponse) SetStatusCode(v int32) *CancelOperationTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOperationTicketResponse) SetBody(v *CancelOperationTicketResponseBody) *CancelOperationTicketResponse {
	s.Body = v
	return s
}

func (s *CancelOperationTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
