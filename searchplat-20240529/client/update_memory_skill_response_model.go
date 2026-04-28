// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemorySkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMemorySkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMemorySkillResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMemorySkillResponseBody) *UpdateMemorySkillResponse
	GetBody() *UpdateMemorySkillResponseBody
}

type UpdateMemorySkillResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMemorySkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemorySkillResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemorySkillResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemorySkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMemorySkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMemorySkillResponse) GetBody() *UpdateMemorySkillResponseBody {
	return s.Body
}

func (s *UpdateMemorySkillResponse) SetHeaders(v map[string]*string) *UpdateMemorySkillResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemorySkillResponse) SetStatusCode(v int32) *UpdateMemorySkillResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemorySkillResponse) SetBody(v *UpdateMemorySkillResponseBody) *UpdateMemorySkillResponse {
	s.Body = v
	return s
}

func (s *UpdateMemorySkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
