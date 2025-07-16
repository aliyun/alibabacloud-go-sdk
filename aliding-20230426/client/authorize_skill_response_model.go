// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeSkillResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeSkillResponseBody) *AuthorizeSkillResponse
	GetBody() *AuthorizeSkillResponseBody
}

type AuthorizeSkillResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSkillResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeSkillResponse) GetBody() *AuthorizeSkillResponseBody {
	return s.Body
}

func (s *AuthorizeSkillResponse) SetHeaders(v map[string]*string) *AuthorizeSkillResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeSkillResponse) SetStatusCode(v int32) *AuthorizeSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeSkillResponse) SetBody(v *AuthorizeSkillResponseBody) *AuthorizeSkillResponse {
	s.Body = v
	return s
}

func (s *AuthorizeSkillResponse) Validate() error {
	return dara.Validate(s)
}
