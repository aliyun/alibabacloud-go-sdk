// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallPolarClawSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallPolarClawSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallPolarClawSkillsResponse
	GetStatusCode() *int32
	SetBody(v *UninstallPolarClawSkillsResponseBody) *UninstallPolarClawSkillsResponse
	GetBody() *UninstallPolarClawSkillsResponseBody
}

type UninstallPolarClawSkillsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallPolarClawSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallPolarClawSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallPolarClawSkillsResponse) GoString() string {
	return s.String()
}

func (s *UninstallPolarClawSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallPolarClawSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallPolarClawSkillsResponse) GetBody() *UninstallPolarClawSkillsResponseBody {
	return s.Body
}

func (s *UninstallPolarClawSkillsResponse) SetHeaders(v map[string]*string) *UninstallPolarClawSkillsResponse {
	s.Headers = v
	return s
}

func (s *UninstallPolarClawSkillsResponse) SetStatusCode(v int32) *UninstallPolarClawSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallPolarClawSkillsResponse) SetBody(v *UninstallPolarClawSkillsResponseBody) *UninstallPolarClawSkillsResponse {
	s.Body = v
	return s
}

func (s *UninstallPolarClawSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
