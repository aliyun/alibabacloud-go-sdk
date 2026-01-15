// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppInstanceTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppInstanceTicketResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppInstanceTicketResponseBody) *CreateAppInstanceTicketResponse
	GetBody() *CreateAppInstanceTicketResponseBody
}

type CreateAppInstanceTicketResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppInstanceTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppInstanceTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppInstanceTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppInstanceTicketResponse) GetBody() *CreateAppInstanceTicketResponseBody {
	return s.Body
}

func (s *CreateAppInstanceTicketResponse) SetHeaders(v map[string]*string) *CreateAppInstanceTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateAppInstanceTicketResponse) SetStatusCode(v int32) *CreateAppInstanceTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppInstanceTicketResponse) SetBody(v *CreateAppInstanceTicketResponseBody) *CreateAppInstanceTicketResponse {
	s.Body = v
	return s
}

func (s *CreateAppInstanceTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
