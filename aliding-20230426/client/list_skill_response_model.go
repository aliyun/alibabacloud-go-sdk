// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillResponseBody) *ListSkillResponse
	GetBody() *ListSkillResponseBody
}

type ListSkillResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillResponse) GoString() string {
	return s.String()
}

func (s *ListSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillResponse) GetBody() *ListSkillResponseBody {
	return s.Body
}

func (s *ListSkillResponse) SetHeaders(v map[string]*string) *ListSkillResponse {
	s.Headers = v
	return s
}

func (s *ListSkillResponse) SetStatusCode(v int32) *ListSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillResponse) SetBody(v *ListSkillResponseBody) *ListSkillResponse {
	s.Body = v
	return s
}

func (s *ListSkillResponse) Validate() error {
	return dara.Validate(s)
}
