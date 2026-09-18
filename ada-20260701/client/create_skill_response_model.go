// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSkillResponse
	GetStatusCode() *int32
	SetBody(v *CreateSkillResponseBody) *CreateSkillResponse
	GetBody() *CreateSkillResponseBody
}

type CreateSkillResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSkillResponse) GetBody() *CreateSkillResponseBody {
	return s.Body
}

func (s *CreateSkillResponse) SetHeaders(v map[string]*string) *CreateSkillResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillResponse) SetStatusCode(v int32) *CreateSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSkillResponse) SetBody(v *CreateSkillResponseBody) *CreateSkillResponse {
	s.Body = v
	return s
}

func (s *CreateSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
