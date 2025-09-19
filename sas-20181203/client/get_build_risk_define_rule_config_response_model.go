// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBuildRiskDefineRuleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBuildRiskDefineRuleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBuildRiskDefineRuleConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetBuildRiskDefineRuleConfigResponseBody) *GetBuildRiskDefineRuleConfigResponse
	GetBody() *GetBuildRiskDefineRuleConfigResponseBody
}

type GetBuildRiskDefineRuleConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBuildRiskDefineRuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBuildRiskDefineRuleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBuildRiskDefineRuleConfigResponse) GoString() string {
	return s.String()
}

func (s *GetBuildRiskDefineRuleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBuildRiskDefineRuleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBuildRiskDefineRuleConfigResponse) GetBody() *GetBuildRiskDefineRuleConfigResponseBody {
	return s.Body
}

func (s *GetBuildRiskDefineRuleConfigResponse) SetHeaders(v map[string]*string) *GetBuildRiskDefineRuleConfigResponse {
	s.Headers = v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponse) SetStatusCode(v int32) *GetBuildRiskDefineRuleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponse) SetBody(v *GetBuildRiskDefineRuleConfigResponseBody) *GetBuildRiskDefineRuleConfigResponse {
	s.Body = v
	return s
}

func (s *GetBuildRiskDefineRuleConfigResponse) Validate() error {
	return dara.Validate(s)
}
