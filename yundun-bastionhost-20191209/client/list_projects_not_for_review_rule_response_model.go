// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsNotForReviewRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectsNotForReviewRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectsNotForReviewRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectsNotForReviewRuleResponseBody) *ListProjectsNotForReviewRuleResponse
	GetBody() *ListProjectsNotForReviewRuleResponseBody
}

type ListProjectsNotForReviewRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsNotForReviewRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsNotForReviewRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsNotForReviewRuleResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsNotForReviewRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectsNotForReviewRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectsNotForReviewRuleResponse) GetBody() *ListProjectsNotForReviewRuleResponseBody {
	return s.Body
}

func (s *ListProjectsNotForReviewRuleResponse) SetHeaders(v map[string]*string) *ListProjectsNotForReviewRuleResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsNotForReviewRuleResponse) SetStatusCode(v int32) *ListProjectsNotForReviewRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsNotForReviewRuleResponse) SetBody(v *ListProjectsNotForReviewRuleResponseBody) *ListProjectsNotForReviewRuleResponse {
	s.Body = v
	return s
}

func (s *ListProjectsNotForReviewRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
