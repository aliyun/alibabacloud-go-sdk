// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReviewRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyReviewRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyReviewRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyReviewRuleResponseBody) *ModifyReviewRuleResponse
	GetBody() *ModifyReviewRuleResponseBody
}

type ModifyReviewRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyReviewRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyReviewRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyReviewRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyReviewRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyReviewRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyReviewRuleResponse) GetBody() *ModifyReviewRuleResponseBody {
	return s.Body
}

func (s *ModifyReviewRuleResponse) SetHeaders(v map[string]*string) *ModifyReviewRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyReviewRuleResponse) SetStatusCode(v int32) *ModifyReviewRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyReviewRuleResponse) SetBody(v *ModifyReviewRuleResponseBody) *ModifyReviewRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyReviewRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
