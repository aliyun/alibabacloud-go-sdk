// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSkillGroupResponseBody) *DeleteSkillGroupResponse
	GetBody() *DeleteSkillGroupResponseBody
}

type DeleteSkillGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSkillGroupResponse) GetBody() *DeleteSkillGroupResponseBody {
	return s.Body
}

func (s *DeleteSkillGroupResponse) SetHeaders(v map[string]*string) *DeleteSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSkillGroupResponse) SetStatusCode(v int32) *DeleteSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSkillGroupResponse) SetBody(v *DeleteSkillGroupResponseBody) *DeleteSkillGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteSkillGroupResponse) Validate() error {
	return dara.Validate(s)
}
