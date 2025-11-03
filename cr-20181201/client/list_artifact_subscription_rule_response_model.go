// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactSubscriptionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListArtifactSubscriptionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListArtifactSubscriptionRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListArtifactSubscriptionRuleResponseBody) *ListArtifactSubscriptionRuleResponse
	GetBody() *ListArtifactSubscriptionRuleResponseBody
}

type ListArtifactSubscriptionRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactSubscriptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactSubscriptionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactSubscriptionRuleResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListArtifactSubscriptionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListArtifactSubscriptionRuleResponse) GetBody() *ListArtifactSubscriptionRuleResponseBody {
	return s.Body
}

func (s *ListArtifactSubscriptionRuleResponse) SetHeaders(v map[string]*string) *ListArtifactSubscriptionRuleResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactSubscriptionRuleResponse) SetStatusCode(v int32) *ListArtifactSubscriptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponse) SetBody(v *ListArtifactSubscriptionRuleResponseBody) *ListArtifactSubscriptionRuleResponse {
	s.Body = v
	return s
}

func (s *ListArtifactSubscriptionRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
