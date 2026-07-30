// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOAuthAuthorizationSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOAuthAuthorizationSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOAuthAuthorizationSessionResponse
	GetStatusCode() *int32
	SetBody(v *GetOAuthAuthorizationSessionResponseBody) *GetOAuthAuthorizationSessionResponse
	GetBody() *GetOAuthAuthorizationSessionResponseBody
}

type GetOAuthAuthorizationSessionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOAuthAuthorizationSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOAuthAuthorizationSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOAuthAuthorizationSessionResponse) GoString() string {
	return s.String()
}

func (s *GetOAuthAuthorizationSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOAuthAuthorizationSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOAuthAuthorizationSessionResponse) GetBody() *GetOAuthAuthorizationSessionResponseBody {
	return s.Body
}

func (s *GetOAuthAuthorizationSessionResponse) SetHeaders(v map[string]*string) *GetOAuthAuthorizationSessionResponse {
	s.Headers = v
	return s
}

func (s *GetOAuthAuthorizationSessionResponse) SetStatusCode(v int32) *GetOAuthAuthorizationSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOAuthAuthorizationSessionResponse) SetBody(v *GetOAuthAuthorizationSessionResponseBody) *GetOAuthAuthorizationSessionResponse {
	s.Body = v
	return s
}

func (s *GetOAuthAuthorizationSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
