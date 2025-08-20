// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactLifecycleRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListArtifactLifecycleRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListArtifactLifecycleRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListArtifactLifecycleRuleResponseBody) *ListArtifactLifecycleRuleResponse
	GetBody() *ListArtifactLifecycleRuleResponseBody
}

type ListArtifactLifecycleRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactLifecycleRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactLifecycleRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListArtifactLifecycleRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListArtifactLifecycleRuleResponse) GetBody() *ListArtifactLifecycleRuleResponseBody {
	return s.Body
}

func (s *ListArtifactLifecycleRuleResponse) SetHeaders(v map[string]*string) *ListArtifactLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactLifecycleRuleResponse) SetStatusCode(v int32) *ListArtifactLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponse) SetBody(v *ListArtifactLifecycleRuleResponseBody) *ListArtifactLifecycleRuleResponse {
	s.Body = v
	return s
}

func (s *ListArtifactLifecycleRuleResponse) Validate() error {
	return dara.Validate(s)
}
