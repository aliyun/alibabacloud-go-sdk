// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReviewRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetReviewRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetReviewRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetReviewRuleResponseBody) *GetReviewRuleResponse
	GetBody() *GetReviewRuleResponseBody
}

type GetReviewRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReviewRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReviewRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetReviewRuleResponse) GoString() string {
	return s.String()
}

func (s *GetReviewRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetReviewRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetReviewRuleResponse) GetBody() *GetReviewRuleResponseBody {
	return s.Body
}

func (s *GetReviewRuleResponse) SetHeaders(v map[string]*string) *GetReviewRuleResponse {
	s.Headers = v
	return s
}

func (s *GetReviewRuleResponse) SetStatusCode(v int32) *GetReviewRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReviewRuleResponse) SetBody(v *GetReviewRuleResponseBody) *GetReviewRuleResponse {
	s.Body = v
	return s
}

func (s *GetReviewRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
