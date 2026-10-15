// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserInvitationCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserInvitationCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserInvitationCodeResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserInvitationCodeResponseBody) *CreateUserInvitationCodeResponse
	GetBody() *CreateUserInvitationCodeResponseBody
}

type CreateUserInvitationCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserInvitationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserInvitationCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserInvitationCodeResponse) GoString() string {
	return s.String()
}

func (s *CreateUserInvitationCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserInvitationCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserInvitationCodeResponse) GetBody() *CreateUserInvitationCodeResponseBody {
	return s.Body
}

func (s *CreateUserInvitationCodeResponse) SetHeaders(v map[string]*string) *CreateUserInvitationCodeResponse {
	s.Headers = v
	return s
}

func (s *CreateUserInvitationCodeResponse) SetStatusCode(v int32) *CreateUserInvitationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserInvitationCodeResponse) SetBody(v *CreateUserInvitationCodeResponseBody) *CreateUserInvitationCodeResponse {
	s.Body = v
	return s
}

func (s *CreateUserInvitationCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
