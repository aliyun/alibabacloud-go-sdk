// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunSkillResponse
	GetStatusCode() *int32
	SetBody(v *RunSkillResponseBody) *RunSkillResponse
	GetBody() *RunSkillResponseBody
}

type RunSkillResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s RunSkillResponse) GoString() string {
	return s.String()
}

func (s *RunSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunSkillResponse) GetBody() *RunSkillResponseBody {
	return s.Body
}

func (s *RunSkillResponse) SetHeaders(v map[string]*string) *RunSkillResponse {
	s.Headers = v
	return s
}

func (s *RunSkillResponse) SetStatusCode(v int32) *RunSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSkillResponse) SetBody(v *RunSkillResponseBody) *RunSkillResponse {
	s.Body = v
	return s
}

func (s *RunSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
