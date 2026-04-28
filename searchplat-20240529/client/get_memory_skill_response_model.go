// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemorySkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemorySkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemorySkillResponse
	GetStatusCode() *int32
	SetBody(v *GetMemorySkillResponseBody) *GetMemorySkillResponse
	GetBody() *GetMemorySkillResponseBody
}

type GetMemorySkillResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemorySkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemorySkillResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemorySkillResponse) GoString() string {
	return s.String()
}

func (s *GetMemorySkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemorySkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemorySkillResponse) GetBody() *GetMemorySkillResponseBody {
	return s.Body
}

func (s *GetMemorySkillResponse) SetHeaders(v map[string]*string) *GetMemorySkillResponse {
	s.Headers = v
	return s
}

func (s *GetMemorySkillResponse) SetStatusCode(v int32) *GetMemorySkillResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemorySkillResponse) SetBody(v *GetMemorySkillResponseBody) *GetMemorySkillResponse {
	s.Body = v
	return s
}

func (s *GetMemorySkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
