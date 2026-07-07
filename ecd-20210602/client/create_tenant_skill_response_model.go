// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTenantSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTenantSkillResponse
	GetStatusCode() *int32
	SetBody(v *CreateTenantSkillResponseBody) *CreateTenantSkillResponse
	GetBody() *CreateTenantSkillResponseBody
}

type CreateTenantSkillResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTenantSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTenantSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantSkillResponse) GoString() string {
	return s.String()
}

func (s *CreateTenantSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTenantSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTenantSkillResponse) GetBody() *CreateTenantSkillResponseBody {
	return s.Body
}

func (s *CreateTenantSkillResponse) SetHeaders(v map[string]*string) *CreateTenantSkillResponse {
	s.Headers = v
	return s
}

func (s *CreateTenantSkillResponse) SetStatusCode(v int32) *CreateTenantSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTenantSkillResponse) SetBody(v *CreateTenantSkillResponseBody) *CreateTenantSkillResponse {
	s.Body = v
	return s
}

func (s *CreateTenantSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
