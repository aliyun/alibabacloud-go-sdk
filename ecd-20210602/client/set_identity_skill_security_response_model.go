// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentitySkillSecurityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetIdentitySkillSecurityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetIdentitySkillSecurityResponse
	GetStatusCode() *int32
	SetBody(v *SetIdentitySkillSecurityResponseBody) *SetIdentitySkillSecurityResponse
	GetBody() *SetIdentitySkillSecurityResponseBody
}

type SetIdentitySkillSecurityResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetIdentitySkillSecurityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetIdentitySkillSecurityResponse) String() string {
	return dara.Prettify(s)
}

func (s SetIdentitySkillSecurityResponse) GoString() string {
	return s.String()
}

func (s *SetIdentitySkillSecurityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetIdentitySkillSecurityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetIdentitySkillSecurityResponse) GetBody() *SetIdentitySkillSecurityResponseBody {
	return s.Body
}

func (s *SetIdentitySkillSecurityResponse) SetHeaders(v map[string]*string) *SetIdentitySkillSecurityResponse {
	s.Headers = v
	return s
}

func (s *SetIdentitySkillSecurityResponse) SetStatusCode(v int32) *SetIdentitySkillSecurityResponse {
	s.StatusCode = &v
	return s
}

func (s *SetIdentitySkillSecurityResponse) SetBody(v *SetIdentitySkillSecurityResponseBody) *SetIdentitySkillSecurityResponse {
	s.Body = v
	return s
}

func (s *SetIdentitySkillSecurityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
