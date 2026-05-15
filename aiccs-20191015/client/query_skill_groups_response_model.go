// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySkillGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySkillGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySkillGroupsResponse
	GetStatusCode() *int32
	SetBody(v *QuerySkillGroupsResponseBody) *QuerySkillGroupsResponse
	GetBody() *QuerySkillGroupsResponseBody
}

type QuerySkillGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySkillGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySkillGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySkillGroupsResponse) GetBody() *QuerySkillGroupsResponseBody {
	return s.Body
}

func (s *QuerySkillGroupsResponse) SetHeaders(v map[string]*string) *QuerySkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *QuerySkillGroupsResponse) SetStatusCode(v int32) *QuerySkillGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySkillGroupsResponse) SetBody(v *QuerySkillGroupsResponseBody) *QuerySkillGroupsResponse {
	s.Body = v
	return s
}

func (s *QuerySkillGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
