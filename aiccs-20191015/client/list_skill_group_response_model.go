// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillGroupResponseBody) *ListSkillGroupResponse
	GetBody() *ListSkillGroupResponseBody
}

type ListSkillGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillGroupResponse) GetBody() *ListSkillGroupResponseBody {
	return s.Body
}

func (s *ListSkillGroupResponse) SetHeaders(v map[string]*string) *ListSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ListSkillGroupResponse) SetStatusCode(v int32) *ListSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillGroupResponse) SetBody(v *ListSkillGroupResponseBody) *ListSkillGroupResponse {
	s.Body = v
	return s
}

func (s *ListSkillGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
