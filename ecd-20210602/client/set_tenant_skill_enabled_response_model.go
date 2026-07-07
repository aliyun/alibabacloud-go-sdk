// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTenantSkillEnabledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetTenantSkillEnabledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetTenantSkillEnabledResponse
	GetStatusCode() *int32
	SetBody(v *SetTenantSkillEnabledResponseBody) *SetTenantSkillEnabledResponse
	GetBody() *SetTenantSkillEnabledResponseBody
}

type SetTenantSkillEnabledResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetTenantSkillEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetTenantSkillEnabledResponse) String() string {
	return dara.Prettify(s)
}

func (s SetTenantSkillEnabledResponse) GoString() string {
	return s.String()
}

func (s *SetTenantSkillEnabledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetTenantSkillEnabledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetTenantSkillEnabledResponse) GetBody() *SetTenantSkillEnabledResponseBody {
	return s.Body
}

func (s *SetTenantSkillEnabledResponse) SetHeaders(v map[string]*string) *SetTenantSkillEnabledResponse {
	s.Headers = v
	return s
}

func (s *SetTenantSkillEnabledResponse) SetStatusCode(v int32) *SetTenantSkillEnabledResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTenantSkillEnabledResponse) SetBody(v *SetTenantSkillEnabledResponseBody) *SetTenantSkillEnabledResponse {
	s.Body = v
	return s
}

func (s *SetTenantSkillEnabledResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
