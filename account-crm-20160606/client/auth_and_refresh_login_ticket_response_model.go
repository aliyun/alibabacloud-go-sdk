// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthAndRefreshLoginTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthAndRefreshLoginTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthAndRefreshLoginTicketResponse
	GetStatusCode() *int32
	SetBody(v *AuthAndRefreshLoginTicketResponseBody) *AuthAndRefreshLoginTicketResponse
	GetBody() *AuthAndRefreshLoginTicketResponseBody
}

type AuthAndRefreshLoginTicketResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthAndRefreshLoginTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthAndRefreshLoginTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthAndRefreshLoginTicketResponse) GoString() string {
	return s.String()
}

func (s *AuthAndRefreshLoginTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthAndRefreshLoginTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthAndRefreshLoginTicketResponse) GetBody() *AuthAndRefreshLoginTicketResponseBody {
	return s.Body
}

func (s *AuthAndRefreshLoginTicketResponse) SetHeaders(v map[string]*string) *AuthAndRefreshLoginTicketResponse {
	s.Headers = v
	return s
}

func (s *AuthAndRefreshLoginTicketResponse) SetStatusCode(v int32) *AuthAndRefreshLoginTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthAndRefreshLoginTicketResponse) SetBody(v *AuthAndRefreshLoginTicketResponseBody) *AuthAndRefreshLoginTicketResponse {
	s.Body = v
	return s
}

func (s *AuthAndRefreshLoginTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
