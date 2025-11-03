// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactSubscriptionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArtifactSubscriptionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArtifactSubscriptionRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetArtifactSubscriptionRuleResponseBody) *GetArtifactSubscriptionRuleResponse
	GetBody() *GetArtifactSubscriptionRuleResponseBody
}

type GetArtifactSubscriptionRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactSubscriptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactSubscriptionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactSubscriptionRuleResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArtifactSubscriptionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArtifactSubscriptionRuleResponse) GetBody() *GetArtifactSubscriptionRuleResponseBody {
	return s.Body
}

func (s *GetArtifactSubscriptionRuleResponse) SetHeaders(v map[string]*string) *GetArtifactSubscriptionRuleResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactSubscriptionRuleResponse) SetStatusCode(v int32) *GetArtifactSubscriptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponse) SetBody(v *GetArtifactSubscriptionRuleResponseBody) *GetArtifactSubscriptionRuleResponse {
	s.Body = v
	return s
}

func (s *GetArtifactSubscriptionRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
