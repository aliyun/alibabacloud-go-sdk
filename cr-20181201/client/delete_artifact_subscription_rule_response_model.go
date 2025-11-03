// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteArtifactSubscriptionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteArtifactSubscriptionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteArtifactSubscriptionRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteArtifactSubscriptionRuleResponseBody) *DeleteArtifactSubscriptionRuleResponse
	GetBody() *DeleteArtifactSubscriptionRuleResponseBody
}

type DeleteArtifactSubscriptionRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteArtifactSubscriptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteArtifactSubscriptionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteArtifactSubscriptionRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteArtifactSubscriptionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteArtifactSubscriptionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteArtifactSubscriptionRuleResponse) GetBody() *DeleteArtifactSubscriptionRuleResponseBody {
	return s.Body
}

func (s *DeleteArtifactSubscriptionRuleResponse) SetHeaders(v map[string]*string) *DeleteArtifactSubscriptionRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteArtifactSubscriptionRuleResponse) SetStatusCode(v int32) *DeleteArtifactSubscriptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteArtifactSubscriptionRuleResponse) SetBody(v *DeleteArtifactSubscriptionRuleResponseBody) *DeleteArtifactSubscriptionRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteArtifactSubscriptionRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
