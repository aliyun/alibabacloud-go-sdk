// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSkillResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSkillResponseBody) *UpdateSkillResponse
	GetBody() *UpdateSkillResponseBody
}

type UpdateSkillResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSkillResponse) GetBody() *UpdateSkillResponseBody {
	return s.Body
}

func (s *UpdateSkillResponse) SetHeaders(v map[string]*string) *UpdateSkillResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillResponse) SetStatusCode(v int32) *UpdateSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSkillResponse) SetBody(v *UpdateSkillResponseBody) *UpdateSkillResponse {
	s.Body = v
	return s
}

func (s *UpdateSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
