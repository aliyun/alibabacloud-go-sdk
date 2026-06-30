// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPolarClawSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchPolarClawSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchPolarClawSkillsResponse
	GetStatusCode() *int32
	SetBody(v *SearchPolarClawSkillsResponseBody) *SearchPolarClawSkillsResponse
	GetBody() *SearchPolarClawSkillsResponseBody
}

type SearchPolarClawSkillsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchPolarClawSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchPolarClawSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchPolarClawSkillsResponse) GoString() string {
	return s.String()
}

func (s *SearchPolarClawSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchPolarClawSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchPolarClawSkillsResponse) GetBody() *SearchPolarClawSkillsResponseBody {
	return s.Body
}

func (s *SearchPolarClawSkillsResponse) SetHeaders(v map[string]*string) *SearchPolarClawSkillsResponse {
	s.Headers = v
	return s
}

func (s *SearchPolarClawSkillsResponse) SetStatusCode(v int32) *SearchPolarClawSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchPolarClawSkillsResponse) SetBody(v *SearchPolarClawSkillsResponseBody) *SearchPolarClawSkillsResponse {
	s.Body = v
	return s
}

func (s *SearchPolarClawSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
