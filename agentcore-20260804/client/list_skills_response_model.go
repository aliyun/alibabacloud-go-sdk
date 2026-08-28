// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillsResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillsResponseBody) *ListSkillsResponse
	GetBody() *ListSkillsResponseBody
}

type ListSkillsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponse) GoString() string {
	return s.String()
}

func (s *ListSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillsResponse) GetBody() *ListSkillsResponseBody {
	return s.Body
}

func (s *ListSkillsResponse) SetHeaders(v map[string]*string) *ListSkillsResponse {
	s.Headers = v
	return s
}

func (s *ListSkillsResponse) SetStatusCode(v int32) *ListSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillsResponse) SetBody(v *ListSkillsResponseBody) *ListSkillsResponse {
	s.Body = v
	return s
}

func (s *ListSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
