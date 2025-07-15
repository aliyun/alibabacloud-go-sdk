// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupStatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillGroupStatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillGroupStatesResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillGroupStatesResponseBody) *ListSkillGroupStatesResponse
	GetBody() *ListSkillGroupStatesResponseBody
}

type ListSkillGroupStatesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillGroupStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillGroupStatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupStatesResponse) GoString() string {
	return s.String()
}

func (s *ListSkillGroupStatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillGroupStatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillGroupStatesResponse) GetBody() *ListSkillGroupStatesResponseBody {
	return s.Body
}

func (s *ListSkillGroupStatesResponse) SetHeaders(v map[string]*string) *ListSkillGroupStatesResponse {
	s.Headers = v
	return s
}

func (s *ListSkillGroupStatesResponse) SetStatusCode(v int32) *ListSkillGroupStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillGroupStatesResponse) SetBody(v *ListSkillGroupStatesResponseBody) *ListSkillGroupStatesResponse {
	s.Body = v
	return s
}

func (s *ListSkillGroupStatesResponse) Validate() error {
	return dara.Validate(s)
}
