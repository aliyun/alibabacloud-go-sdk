// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallSkillsResponse
	GetStatusCode() *int32
	SetBody(v *InstallSkillsResponseBody) *InstallSkillsResponse
	GetBody() *InstallSkillsResponseBody
}

type InstallSkillsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallSkillsResponse) GoString() string {
	return s.String()
}

func (s *InstallSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallSkillsResponse) GetBody() *InstallSkillsResponseBody {
	return s.Body
}

func (s *InstallSkillsResponse) SetHeaders(v map[string]*string) *InstallSkillsResponse {
	s.Headers = v
	return s
}

func (s *InstallSkillsResponse) SetStatusCode(v int32) *InstallSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallSkillsResponse) SetBody(v *InstallSkillsResponseBody) *InstallSkillsResponse {
	s.Body = v
	return s
}

func (s *InstallSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
