// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBuildRiskDefineRuleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetBuildRiskDefineRuleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetBuildRiskDefineRuleConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetBuildRiskDefineRuleConfigResponseBody) *SetBuildRiskDefineRuleConfigResponse
	GetBody() *SetBuildRiskDefineRuleConfigResponseBody
}

type SetBuildRiskDefineRuleConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetBuildRiskDefineRuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetBuildRiskDefineRuleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetBuildRiskDefineRuleConfigResponse) GoString() string {
	return s.String()
}

func (s *SetBuildRiskDefineRuleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetBuildRiskDefineRuleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetBuildRiskDefineRuleConfigResponse) GetBody() *SetBuildRiskDefineRuleConfigResponseBody {
	return s.Body
}

func (s *SetBuildRiskDefineRuleConfigResponse) SetHeaders(v map[string]*string) *SetBuildRiskDefineRuleConfigResponse {
	s.Headers = v
	return s
}

func (s *SetBuildRiskDefineRuleConfigResponse) SetStatusCode(v int32) *SetBuildRiskDefineRuleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetBuildRiskDefineRuleConfigResponse) SetBody(v *SetBuildRiskDefineRuleConfigResponseBody) *SetBuildRiskDefineRuleConfigResponse {
	s.Body = v
	return s
}

func (s *SetBuildRiskDefineRuleConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
