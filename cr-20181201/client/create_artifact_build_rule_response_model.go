// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactBuildRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateArtifactBuildRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateArtifactBuildRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateArtifactBuildRuleResponseBody) *CreateArtifactBuildRuleResponse
	GetBody() *CreateArtifactBuildRuleResponseBody
}

type CreateArtifactBuildRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArtifactBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactBuildRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactBuildRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateArtifactBuildRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateArtifactBuildRuleResponse) GetBody() *CreateArtifactBuildRuleResponseBody {
	return s.Body
}

func (s *CreateArtifactBuildRuleResponse) SetHeaders(v map[string]*string) *CreateArtifactBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactBuildRuleResponse) SetStatusCode(v int32) *CreateArtifactBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactBuildRuleResponse) SetBody(v *CreateArtifactBuildRuleResponseBody) *CreateArtifactBuildRuleResponse {
	s.Body = v
	return s
}

func (s *CreateArtifactBuildRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
