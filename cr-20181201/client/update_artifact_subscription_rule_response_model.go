// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateArtifactSubscriptionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateArtifactSubscriptionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateArtifactSubscriptionRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateArtifactSubscriptionRuleResponseBody) *UpdateArtifactSubscriptionRuleResponse
	GetBody() *UpdateArtifactSubscriptionRuleResponseBody
}

type UpdateArtifactSubscriptionRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateArtifactSubscriptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateArtifactSubscriptionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactSubscriptionRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateArtifactSubscriptionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateArtifactSubscriptionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateArtifactSubscriptionRuleResponse) GetBody() *UpdateArtifactSubscriptionRuleResponseBody {
	return s.Body
}

func (s *UpdateArtifactSubscriptionRuleResponse) SetHeaders(v map[string]*string) *UpdateArtifactSubscriptionRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateArtifactSubscriptionRuleResponse) SetStatusCode(v int32) *UpdateArtifactSubscriptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleResponse) SetBody(v *UpdateArtifactSubscriptionRuleResponseBody) *UpdateArtifactSubscriptionRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateArtifactSubscriptionRuleResponse) Validate() error {
	return dara.Validate(s)
}
