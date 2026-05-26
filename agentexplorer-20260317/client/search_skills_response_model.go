// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchSkillsResponse
	GetStatusCode() *int32
	SetBody(v *SearchSkillsResponseBody) *SearchSkillsResponse
	GetBody() *SearchSkillsResponseBody
}

type SearchSkillsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchSkillsResponse) GoString() string {
	return s.String()
}

func (s *SearchSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchSkillsResponse) GetBody() *SearchSkillsResponseBody {
	return s.Body
}

func (s *SearchSkillsResponse) SetHeaders(v map[string]*string) *SearchSkillsResponse {
	s.Headers = v
	return s
}

func (s *SearchSkillsResponse) SetStatusCode(v int32) *SearchSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchSkillsResponse) SetBody(v *SearchSkillsResponseBody) *SearchSkillsResponse {
	s.Body = v
	return s
}

func (s *SearchSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
