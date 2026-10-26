// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachReviewersToReviewRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachReviewersToReviewRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachReviewersToReviewRuleResponse
	GetStatusCode() *int32
	SetBody(v *AttachReviewersToReviewRuleResponseBody) *AttachReviewersToReviewRuleResponse
	GetBody() *AttachReviewersToReviewRuleResponseBody
}

type AttachReviewersToReviewRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachReviewersToReviewRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachReviewersToReviewRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachReviewersToReviewRuleResponse) GoString() string {
	return s.String()
}

func (s *AttachReviewersToReviewRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachReviewersToReviewRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachReviewersToReviewRuleResponse) GetBody() *AttachReviewersToReviewRuleResponseBody {
	return s.Body
}

func (s *AttachReviewersToReviewRuleResponse) SetHeaders(v map[string]*string) *AttachReviewersToReviewRuleResponse {
	s.Headers = v
	return s
}

func (s *AttachReviewersToReviewRuleResponse) SetStatusCode(v int32) *AttachReviewersToReviewRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachReviewersToReviewRuleResponse) SetBody(v *AttachReviewersToReviewRuleResponseBody) *AttachReviewersToReviewRuleResponse {
	s.Body = v
	return s
}

func (s *AttachReviewersToReviewRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
