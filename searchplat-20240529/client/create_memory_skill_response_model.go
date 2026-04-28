// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemorySkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMemorySkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMemorySkillResponse
	GetStatusCode() *int32
	SetBody(v *CreateMemorySkillResponseBody) *CreateMemorySkillResponse
	GetBody() *CreateMemorySkillResponseBody
}

type CreateMemorySkillResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemorySkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemorySkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMemorySkillResponse) GoString() string {
	return s.String()
}

func (s *CreateMemorySkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMemorySkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMemorySkillResponse) GetBody() *CreateMemorySkillResponseBody {
	return s.Body
}

func (s *CreateMemorySkillResponse) SetHeaders(v map[string]*string) *CreateMemorySkillResponse {
	s.Headers = v
	return s
}

func (s *CreateMemorySkillResponse) SetStatusCode(v int32) *CreateMemorySkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemorySkillResponse) SetBody(v *CreateMemorySkillResponseBody) *CreateMemorySkillResponse {
	s.Body = v
	return s
}

func (s *CreateMemorySkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
