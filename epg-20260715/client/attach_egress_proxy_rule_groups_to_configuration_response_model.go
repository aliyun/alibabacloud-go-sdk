// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEgressProxyRuleGroupsToConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachEgressProxyRuleGroupsToConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachEgressProxyRuleGroupsToConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *AttachEgressProxyRuleGroupsToConfigurationResponseBody) *AttachEgressProxyRuleGroupsToConfigurationResponse
	GetBody() *AttachEgressProxyRuleGroupsToConfigurationResponseBody
}

type AttachEgressProxyRuleGroupsToConfigurationResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachEgressProxyRuleGroupsToConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachEgressProxyRuleGroupsToConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachEgressProxyRuleGroupsToConfigurationResponse) GoString() string {
	return s.String()
}

func (s *AttachEgressProxyRuleGroupsToConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachEgressProxyRuleGroupsToConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachEgressProxyRuleGroupsToConfigurationResponse) GetBody() *AttachEgressProxyRuleGroupsToConfigurationResponseBody {
	return s.Body
}

func (s *AttachEgressProxyRuleGroupsToConfigurationResponse) SetHeaders(v map[string]*string) *AttachEgressProxyRuleGroupsToConfigurationResponse {
	s.Headers = v
	return s
}

func (s *AttachEgressProxyRuleGroupsToConfigurationResponse) SetStatusCode(v int32) *AttachEgressProxyRuleGroupsToConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachEgressProxyRuleGroupsToConfigurationResponse) SetBody(v *AttachEgressProxyRuleGroupsToConfigurationResponseBody) *AttachEgressProxyRuleGroupsToConfigurationResponse {
	s.Body = v
	return s
}

func (s *AttachEgressProxyRuleGroupsToConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
