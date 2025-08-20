// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactSubscriptionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateArtifactSubscriptionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateArtifactSubscriptionRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateArtifactSubscriptionRuleResponseBody) *CreateArtifactSubscriptionRuleResponse
	GetBody() *CreateArtifactSubscriptionRuleResponseBody
}

type CreateArtifactSubscriptionRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArtifactSubscriptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactSubscriptionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactSubscriptionRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateArtifactSubscriptionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateArtifactSubscriptionRuleResponse) GetBody() *CreateArtifactSubscriptionRuleResponseBody {
	return s.Body
}

func (s *CreateArtifactSubscriptionRuleResponse) SetHeaders(v map[string]*string) *CreateArtifactSubscriptionRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponse) SetStatusCode(v int32) *CreateArtifactSubscriptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponse) SetBody(v *CreateArtifactSubscriptionRuleResponseBody) *CreateArtifactSubscriptionRuleResponse {
	s.Body = v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponse) Validate() error {
	return dara.Validate(s)
}
