// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachProjectsToReviewRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachProjectsToReviewRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachProjectsToReviewRuleResponse
	GetStatusCode() *int32
	SetBody(v *AttachProjectsToReviewRuleResponseBody) *AttachProjectsToReviewRuleResponse
	GetBody() *AttachProjectsToReviewRuleResponseBody
}

type AttachProjectsToReviewRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachProjectsToReviewRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachProjectsToReviewRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachProjectsToReviewRuleResponse) GoString() string {
	return s.String()
}

func (s *AttachProjectsToReviewRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachProjectsToReviewRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachProjectsToReviewRuleResponse) GetBody() *AttachProjectsToReviewRuleResponseBody {
	return s.Body
}

func (s *AttachProjectsToReviewRuleResponse) SetHeaders(v map[string]*string) *AttachProjectsToReviewRuleResponse {
	s.Headers = v
	return s
}

func (s *AttachProjectsToReviewRuleResponse) SetStatusCode(v int32) *AttachProjectsToReviewRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachProjectsToReviewRuleResponse) SetBody(v *AttachProjectsToReviewRuleResponseBody) *AttachProjectsToReviewRuleResponse {
	s.Body = v
	return s
}

func (s *AttachProjectsToReviewRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
