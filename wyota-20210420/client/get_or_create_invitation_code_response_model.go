// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrCreateInvitationCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrCreateInvitationCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrCreateInvitationCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetOrCreateInvitationCodeResponseBody) *GetOrCreateInvitationCodeResponse
	GetBody() *GetOrCreateInvitationCodeResponseBody
}

type GetOrCreateInvitationCodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrCreateInvitationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrCreateInvitationCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrCreateInvitationCodeResponse) GoString() string {
	return s.String()
}

func (s *GetOrCreateInvitationCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrCreateInvitationCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrCreateInvitationCodeResponse) GetBody() *GetOrCreateInvitationCodeResponseBody {
	return s.Body
}

func (s *GetOrCreateInvitationCodeResponse) SetHeaders(v map[string]*string) *GetOrCreateInvitationCodeResponse {
	s.Headers = v
	return s
}

func (s *GetOrCreateInvitationCodeResponse) SetStatusCode(v int32) *GetOrCreateInvitationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrCreateInvitationCodeResponse) SetBody(v *GetOrCreateInvitationCodeResponseBody) *GetOrCreateInvitationCodeResponse {
	s.Body = v
	return s
}

func (s *GetOrCreateInvitationCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
