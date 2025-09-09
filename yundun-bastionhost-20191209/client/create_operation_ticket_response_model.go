// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOperationTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOperationTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOperationTicketResponse
	GetStatusCode() *int32
	SetBody(v *CreateOperationTicketResponseBody) *CreateOperationTicketResponse
	GetBody() *CreateOperationTicketResponseBody
}

type CreateOperationTicketResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOperationTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOperationTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOperationTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateOperationTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOperationTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOperationTicketResponse) GetBody() *CreateOperationTicketResponseBody {
	return s.Body
}

func (s *CreateOperationTicketResponse) SetHeaders(v map[string]*string) *CreateOperationTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateOperationTicketResponse) SetStatusCode(v int32) *CreateOperationTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOperationTicketResponse) SetBody(v *CreateOperationTicketResponseBody) *CreateOperationTicketResponse {
	s.Body = v
	return s
}

func (s *CreateOperationTicketResponse) Validate() error {
	return dara.Validate(s)
}
