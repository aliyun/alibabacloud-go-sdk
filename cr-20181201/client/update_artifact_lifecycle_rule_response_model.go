// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateArtifactLifecycleRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateArtifactLifecycleRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateArtifactLifecycleRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateArtifactLifecycleRuleResponseBody) *UpdateArtifactLifecycleRuleResponse
	GetBody() *UpdateArtifactLifecycleRuleResponseBody
}

type UpdateArtifactLifecycleRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateArtifactLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateArtifactLifecycleRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateArtifactLifecycleRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateArtifactLifecycleRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateArtifactLifecycleRuleResponse) GetBody() *UpdateArtifactLifecycleRuleResponseBody {
	return s.Body
}

func (s *UpdateArtifactLifecycleRuleResponse) SetHeaders(v map[string]*string) *UpdateArtifactLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateArtifactLifecycleRuleResponse) SetStatusCode(v int32) *UpdateArtifactLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleResponse) SetBody(v *UpdateArtifactLifecycleRuleResponseBody) *UpdateArtifactLifecycleRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateArtifactLifecycleRuleResponse) Validate() error {
	return dara.Validate(s)
}
