// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthLoginTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthLoginTicketResponse
	GetStatusCode() *int32
	SetBody(v *AuthLoginTicketResponseBody) *AuthLoginTicketResponse
	GetBody() *AuthLoginTicketResponseBody
}

type AuthLoginTicketResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthLoginTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthLoginTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginTicketResponse) GoString() string {
	return s.String()
}

func (s *AuthLoginTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthLoginTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthLoginTicketResponse) GetBody() *AuthLoginTicketResponseBody {
	return s.Body
}

func (s *AuthLoginTicketResponse) SetHeaders(v map[string]*string) *AuthLoginTicketResponse {
	s.Headers = v
	return s
}

func (s *AuthLoginTicketResponse) SetStatusCode(v int32) *AuthLoginTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthLoginTicketResponse) SetBody(v *AuthLoginTicketResponseBody) *AuthLoginTicketResponse {
	s.Body = v
	return s
}

func (s *AuthLoginTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
