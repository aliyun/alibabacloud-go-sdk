// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSkillGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveSkillGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveSkillGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveSkillGroupResponseBody) *RemoveSkillGroupResponse
	GetBody() *RemoveSkillGroupResponseBody
}

type RemoveSkillGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSkillGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveSkillGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveSkillGroupResponse) GetBody() *RemoveSkillGroupResponseBody {
	return s.Body
}

func (s *RemoveSkillGroupResponse) SetHeaders(v map[string]*string) *RemoveSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveSkillGroupResponse) SetStatusCode(v int32) *RemoveSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSkillGroupResponse) SetBody(v *RemoveSkillGroupResponseBody) *RemoveSkillGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveSkillGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
