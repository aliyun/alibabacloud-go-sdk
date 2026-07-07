// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentitySkillAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetIdentitySkillAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetIdentitySkillAuthResponse
	GetStatusCode() *int32
	SetBody(v *SetIdentitySkillAuthResponseBody) *SetIdentitySkillAuthResponse
	GetBody() *SetIdentitySkillAuthResponseBody
}

type SetIdentitySkillAuthResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetIdentitySkillAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetIdentitySkillAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s SetIdentitySkillAuthResponse) GoString() string {
	return s.String()
}

func (s *SetIdentitySkillAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetIdentitySkillAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetIdentitySkillAuthResponse) GetBody() *SetIdentitySkillAuthResponseBody {
	return s.Body
}

func (s *SetIdentitySkillAuthResponse) SetHeaders(v map[string]*string) *SetIdentitySkillAuthResponse {
	s.Headers = v
	return s
}

func (s *SetIdentitySkillAuthResponse) SetStatusCode(v int32) *SetIdentitySkillAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *SetIdentitySkillAuthResponse) SetBody(v *SetIdentitySkillAuthResponseBody) *SetIdentitySkillAuthResponse {
	s.Body = v
	return s
}

func (s *SetIdentitySkillAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
