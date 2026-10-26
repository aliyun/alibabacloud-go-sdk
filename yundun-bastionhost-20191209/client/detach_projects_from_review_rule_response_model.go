// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachProjectsFromReviewRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachProjectsFromReviewRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachProjectsFromReviewRuleResponse
	GetStatusCode() *int32
	SetBody(v *DetachProjectsFromReviewRuleResponseBody) *DetachProjectsFromReviewRuleResponse
	GetBody() *DetachProjectsFromReviewRuleResponseBody
}

type DetachProjectsFromReviewRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachProjectsFromReviewRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachProjectsFromReviewRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachProjectsFromReviewRuleResponse) GoString() string {
	return s.String()
}

func (s *DetachProjectsFromReviewRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachProjectsFromReviewRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachProjectsFromReviewRuleResponse) GetBody() *DetachProjectsFromReviewRuleResponseBody {
	return s.Body
}

func (s *DetachProjectsFromReviewRuleResponse) SetHeaders(v map[string]*string) *DetachProjectsFromReviewRuleResponse {
	s.Headers = v
	return s
}

func (s *DetachProjectsFromReviewRuleResponse) SetStatusCode(v int32) *DetachProjectsFromReviewRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachProjectsFromReviewRuleResponse) SetBody(v *DetachProjectsFromReviewRuleResponseBody) *DetachProjectsFromReviewRuleResponse {
	s.Body = v
	return s
}

func (s *DetachProjectsFromReviewRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
