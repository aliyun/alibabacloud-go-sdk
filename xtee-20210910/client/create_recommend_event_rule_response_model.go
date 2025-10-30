// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecommendEventRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecommendEventRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecommendEventRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecommendEventRuleResponseBody) *CreateRecommendEventRuleResponse
	GetBody() *CreateRecommendEventRuleResponseBody
}

type CreateRecommendEventRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecommendEventRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecommendEventRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecommendEventRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRecommendEventRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecommendEventRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecommendEventRuleResponse) GetBody() *CreateRecommendEventRuleResponseBody {
	return s.Body
}

func (s *CreateRecommendEventRuleResponse) SetHeaders(v map[string]*string) *CreateRecommendEventRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRecommendEventRuleResponse) SetStatusCode(v int32) *CreateRecommendEventRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecommendEventRuleResponse) SetBody(v *CreateRecommendEventRuleResponseBody) *CreateRecommendEventRuleResponse {
	s.Body = v
	return s
}

func (s *CreateRecommendEventRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
