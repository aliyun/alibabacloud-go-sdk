// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCoordinationTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCoordinationTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCoordinationTicketResponse
	GetStatusCode() *int32
	SetBody(v *GetCoordinationTicketResponseBody) *GetCoordinationTicketResponse
	GetBody() *GetCoordinationTicketResponseBody
}

type GetCoordinationTicketResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCoordinationTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCoordinationTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCoordinationTicketResponse) GoString() string {
	return s.String()
}

func (s *GetCoordinationTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCoordinationTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCoordinationTicketResponse) GetBody() *GetCoordinationTicketResponseBody {
	return s.Body
}

func (s *GetCoordinationTicketResponse) SetHeaders(v map[string]*string) *GetCoordinationTicketResponse {
	s.Headers = v
	return s
}

func (s *GetCoordinationTicketResponse) SetStatusCode(v int32) *GetCoordinationTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCoordinationTicketResponse) SetBody(v *GetCoordinationTicketResponseBody) *GetCoordinationTicketResponse {
	s.Body = v
	return s
}

func (s *GetCoordinationTicketResponse) Validate() error {
	return dara.Validate(s)
}
