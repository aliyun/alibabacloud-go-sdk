// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteArtifactLifecycleRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteArtifactLifecycleRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteArtifactLifecycleRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteArtifactLifecycleRuleResponseBody) *DeleteArtifactLifecycleRuleResponse
	GetBody() *DeleteArtifactLifecycleRuleResponseBody
}

type DeleteArtifactLifecycleRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteArtifactLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteArtifactLifecycleRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteArtifactLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteArtifactLifecycleRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteArtifactLifecycleRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteArtifactLifecycleRuleResponse) GetBody() *DeleteArtifactLifecycleRuleResponseBody {
	return s.Body
}

func (s *DeleteArtifactLifecycleRuleResponse) SetHeaders(v map[string]*string) *DeleteArtifactLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteArtifactLifecycleRuleResponse) SetStatusCode(v int32) *DeleteArtifactLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteArtifactLifecycleRuleResponse) SetBody(v *DeleteArtifactLifecycleRuleResponseBody) *DeleteArtifactLifecycleRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteArtifactLifecycleRuleResponse) Validate() error {
	return dara.Validate(s)
}
