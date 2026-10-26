// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachReviewersFromReviewRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachReviewersFromReviewRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachReviewersFromReviewRuleResponse
	GetStatusCode() *int32
	SetBody(v *DetachReviewersFromReviewRuleResponseBody) *DetachReviewersFromReviewRuleResponse
	GetBody() *DetachReviewersFromReviewRuleResponseBody
}

type DetachReviewersFromReviewRuleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachReviewersFromReviewRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachReviewersFromReviewRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachReviewersFromReviewRuleResponse) GoString() string {
	return s.String()
}

func (s *DetachReviewersFromReviewRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachReviewersFromReviewRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachReviewersFromReviewRuleResponse) GetBody() *DetachReviewersFromReviewRuleResponseBody {
	return s.Body
}

func (s *DetachReviewersFromReviewRuleResponse) SetHeaders(v map[string]*string) *DetachReviewersFromReviewRuleResponse {
	s.Headers = v
	return s
}

func (s *DetachReviewersFromReviewRuleResponse) SetStatusCode(v int32) *DetachReviewersFromReviewRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachReviewersFromReviewRuleResponse) SetBody(v *DetachReviewersFromReviewRuleResponseBody) *DetachReviewersFromReviewRuleResponse {
	s.Body = v
	return s
}

func (s *DetachReviewersFromReviewRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
