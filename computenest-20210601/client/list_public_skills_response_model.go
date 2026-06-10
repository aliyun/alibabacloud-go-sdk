// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublicSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublicSkillsResponse
	GetStatusCode() *int32
	SetBody(v *ListPublicSkillsResponseBody) *ListPublicSkillsResponse
	GetBody() *ListPublicSkillsResponseBody
}

type ListPublicSkillsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublicSkillsResponse) GoString() string {
	return s.String()
}

func (s *ListPublicSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublicSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublicSkillsResponse) GetBody() *ListPublicSkillsResponseBody {
	return s.Body
}

func (s *ListPublicSkillsResponse) SetHeaders(v map[string]*string) *ListPublicSkillsResponse {
	s.Headers = v
	return s
}

func (s *ListPublicSkillsResponse) SetStatusCode(v int32) *ListPublicSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicSkillsResponse) SetBody(v *ListPublicSkillsResponseBody) *ListPublicSkillsResponse {
	s.Body = v
	return s
}

func (s *ListPublicSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
