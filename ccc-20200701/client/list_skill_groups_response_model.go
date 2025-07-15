// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillGroupsResponseBody) *ListSkillGroupsResponse
	GetBody() *ListSkillGroupsResponseBody
}

type ListSkillGroupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListSkillGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillGroupsResponse) GetBody() *ListSkillGroupsResponseBody {
	return s.Body
}

func (s *ListSkillGroupsResponse) SetHeaders(v map[string]*string) *ListSkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListSkillGroupsResponse) SetStatusCode(v int32) *ListSkillGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillGroupsResponse) SetBody(v *ListSkillGroupsResponseBody) *ListSkillGroupsResponse {
	s.Body = v
	return s
}

func (s *ListSkillGroupsResponse) Validate() error {
	return dara.Validate(s)
}
