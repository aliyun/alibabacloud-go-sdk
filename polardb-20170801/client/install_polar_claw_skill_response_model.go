// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPolarClawSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallPolarClawSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallPolarClawSkillResponse
	GetStatusCode() *int32
	SetBody(v *InstallPolarClawSkillResponseBody) *InstallPolarClawSkillResponse
	GetBody() *InstallPolarClawSkillResponseBody
}

type InstallPolarClawSkillResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallPolarClawSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallPolarClawSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallPolarClawSkillResponse) GoString() string {
	return s.String()
}

func (s *InstallPolarClawSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallPolarClawSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallPolarClawSkillResponse) GetBody() *InstallPolarClawSkillResponseBody {
	return s.Body
}

func (s *InstallPolarClawSkillResponse) SetHeaders(v map[string]*string) *InstallPolarClawSkillResponse {
	s.Headers = v
	return s
}

func (s *InstallPolarClawSkillResponse) SetStatusCode(v int32) *InstallPolarClawSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallPolarClawSkillResponse) SetBody(v *InstallPolarClawSkillResponseBody) *InstallPolarClawSkillResponse {
	s.Body = v
	return s
}

func (s *InstallPolarClawSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
