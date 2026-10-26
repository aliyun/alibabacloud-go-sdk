// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReviewRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReviewRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReviewRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateReviewRuleResponseBody) *CreateReviewRuleResponse
	GetBody() *CreateReviewRuleResponseBody
}

type CreateReviewRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReviewRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReviewRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReviewRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateReviewRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReviewRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReviewRuleResponse) GetBody() *CreateReviewRuleResponseBody {
	return s.Body
}

func (s *CreateReviewRuleResponse) SetHeaders(v map[string]*string) *CreateReviewRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateReviewRuleResponse) SetStatusCode(v int32) *CreateReviewRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReviewRuleResponse) SetBody(v *CreateReviewRuleResponseBody) *CreateReviewRuleResponse {
	s.Body = v
	return s
}

func (s *CreateReviewRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
