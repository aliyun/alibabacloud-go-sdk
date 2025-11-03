// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactBuildRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArtifactBuildRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArtifactBuildRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetArtifactBuildRuleResponseBody) *GetArtifactBuildRuleResponse
	GetBody() *GetArtifactBuildRuleResponseBody
}

type GetArtifactBuildRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactBuildRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArtifactBuildRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArtifactBuildRuleResponse) GetBody() *GetArtifactBuildRuleResponseBody {
	return s.Body
}

func (s *GetArtifactBuildRuleResponse) SetHeaders(v map[string]*string) *GetArtifactBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactBuildRuleResponse) SetStatusCode(v int32) *GetArtifactBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactBuildRuleResponse) SetBody(v *GetArtifactBuildRuleResponseBody) *GetArtifactBuildRuleResponse {
	s.Body = v
	return s
}

func (s *GetArtifactBuildRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
