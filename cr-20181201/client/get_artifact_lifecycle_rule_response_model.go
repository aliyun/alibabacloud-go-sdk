// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactLifecycleRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArtifactLifecycleRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArtifactLifecycleRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetArtifactLifecycleRuleResponseBody) *GetArtifactLifecycleRuleResponse
	GetBody() *GetArtifactLifecycleRuleResponseBody
}

type GetArtifactLifecycleRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactLifecycleRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactLifecycleRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArtifactLifecycleRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArtifactLifecycleRuleResponse) GetBody() *GetArtifactLifecycleRuleResponseBody {
	return s.Body
}

func (s *GetArtifactLifecycleRuleResponse) SetHeaders(v map[string]*string) *GetArtifactLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactLifecycleRuleResponse) SetStatusCode(v int32) *GetArtifactLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponse) SetBody(v *GetArtifactLifecycleRuleResponseBody) *GetArtifactLifecycleRuleResponse {
	s.Body = v
	return s
}

func (s *GetArtifactLifecycleRuleResponse) Validate() error {
	return dara.Validate(s)
}
