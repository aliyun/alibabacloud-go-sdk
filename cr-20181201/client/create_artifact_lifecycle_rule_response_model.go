// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactLifecycleRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateArtifactLifecycleRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateArtifactLifecycleRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateArtifactLifecycleRuleResponseBody) *CreateArtifactLifecycleRuleResponse
	GetBody() *CreateArtifactLifecycleRuleResponseBody
}

type CreateArtifactLifecycleRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArtifactLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactLifecycleRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactLifecycleRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateArtifactLifecycleRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateArtifactLifecycleRuleResponse) GetBody() *CreateArtifactLifecycleRuleResponseBody {
	return s.Body
}

func (s *CreateArtifactLifecycleRuleResponse) SetHeaders(v map[string]*string) *CreateArtifactLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactLifecycleRuleResponse) SetStatusCode(v int32) *CreateArtifactLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactLifecycleRuleResponse) SetBody(v *CreateArtifactLifecycleRuleResponseBody) *CreateArtifactLifecycleRuleResponse {
	s.Body = v
	return s
}

func (s *CreateArtifactLifecycleRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
