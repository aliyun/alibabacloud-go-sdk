// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBriefSkillGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBriefSkillGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBriefSkillGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListBriefSkillGroupsResponseBody) *ListBriefSkillGroupsResponse
	GetBody() *ListBriefSkillGroupsResponseBody
}

type ListBriefSkillGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBriefSkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBriefSkillGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBriefSkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListBriefSkillGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBriefSkillGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBriefSkillGroupsResponse) GetBody() *ListBriefSkillGroupsResponseBody {
	return s.Body
}

func (s *ListBriefSkillGroupsResponse) SetHeaders(v map[string]*string) *ListBriefSkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListBriefSkillGroupsResponse) SetStatusCode(v int32) *ListBriefSkillGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBriefSkillGroupsResponse) SetBody(v *ListBriefSkillGroupsResponseBody) *ListBriefSkillGroupsResponse {
	s.Body = v
	return s
}

func (s *ListBriefSkillGroupsResponse) Validate() error {
	return dara.Validate(s)
}
