// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluatorSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEvaluatorSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEvaluatorSkillsResponse
	GetStatusCode() *int32
	SetBody(v *ListEvaluatorSkillsResponseBody) *ListEvaluatorSkillsResponse
	GetBody() *ListEvaluatorSkillsResponseBody
}

type ListEvaluatorSkillsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluatorSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluatorSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluatorSkillsResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluatorSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEvaluatorSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEvaluatorSkillsResponse) GetBody() *ListEvaluatorSkillsResponseBody {
	return s.Body
}

func (s *ListEvaluatorSkillsResponse) SetHeaders(v map[string]*string) *ListEvaluatorSkillsResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluatorSkillsResponse) SetStatusCode(v int32) *ListEvaluatorSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluatorSkillsResponse) SetBody(v *ListEvaluatorSkillsResponseBody) *ListEvaluatorSkillsResponse {
	s.Body = v
	return s
}

func (s *ListEvaluatorSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
