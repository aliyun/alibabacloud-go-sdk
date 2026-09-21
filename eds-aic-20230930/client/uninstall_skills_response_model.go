// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallSkillsResponse
	GetStatusCode() *int32
	SetBody(v *UninstallSkillsResponseBody) *UninstallSkillsResponse
	GetBody() *UninstallSkillsResponseBody
}

type UninstallSkillsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallSkillsResponse) GoString() string {
	return s.String()
}

func (s *UninstallSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallSkillsResponse) GetBody() *UninstallSkillsResponseBody {
	return s.Body
}

func (s *UninstallSkillsResponse) SetHeaders(v map[string]*string) *UninstallSkillsResponse {
	s.Headers = v
	return s
}

func (s *UninstallSkillsResponse) SetStatusCode(v int32) *UninstallSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallSkillsResponse) SetBody(v *UninstallSkillsResponseBody) *UninstallSkillsResponse {
	s.Body = v
	return s
}

func (s *UninstallSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
