// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckPolarClawSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckPolarClawSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckPolarClawSkillResponse
	GetStatusCode() *int32
	SetBody(v *CheckPolarClawSkillResponseBody) *CheckPolarClawSkillResponse
	GetBody() *CheckPolarClawSkillResponseBody
}

type CheckPolarClawSkillResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckPolarClawSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckPolarClawSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckPolarClawSkillResponse) GoString() string {
	return s.String()
}

func (s *CheckPolarClawSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckPolarClawSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckPolarClawSkillResponse) GetBody() *CheckPolarClawSkillResponseBody {
	return s.Body
}

func (s *CheckPolarClawSkillResponse) SetHeaders(v map[string]*string) *CheckPolarClawSkillResponse {
	s.Headers = v
	return s
}

func (s *CheckPolarClawSkillResponse) SetStatusCode(v int32) *CheckPolarClawSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckPolarClawSkillResponse) SetBody(v *CheckPolarClawSkillResponseBody) *CheckPolarClawSkillResponse {
	s.Body = v
	return s
}

func (s *CheckPolarClawSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
