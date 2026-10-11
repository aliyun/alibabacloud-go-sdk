// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEgressProxyRuleGroupsFromConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachEgressProxyRuleGroupsFromConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachEgressProxyRuleGroupsFromConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DetachEgressProxyRuleGroupsFromConfigurationResponseBody) *DetachEgressProxyRuleGroupsFromConfigurationResponse
	GetBody() *DetachEgressProxyRuleGroupsFromConfigurationResponseBody
}

type DetachEgressProxyRuleGroupsFromConfigurationResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachEgressProxyRuleGroupsFromConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachEgressProxyRuleGroupsFromConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachEgressProxyRuleGroupsFromConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationResponse) GetBody() *DetachEgressProxyRuleGroupsFromConfigurationResponseBody {
	return s.Body
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationResponse) SetHeaders(v map[string]*string) *DetachEgressProxyRuleGroupsFromConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationResponse) SetStatusCode(v int32) *DetachEgressProxyRuleGroupsFromConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationResponse) SetBody(v *DetachEgressProxyRuleGroupsFromConfigurationResponseBody) *DetachEgressProxyRuleGroupsFromConfigurationResponse {
	s.Body = v
	return s
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
