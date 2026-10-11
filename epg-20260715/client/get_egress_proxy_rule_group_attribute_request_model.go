// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyRuleGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyRuleGroupId(v string) *GetEgressProxyRuleGroupAttributeRequest
	GetEgressProxyRuleGroupId() *string
}

type GetEgressProxyRuleGroupAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eprg-uf60966a9633667c****
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
}

func (s GetEgressProxyRuleGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyRuleGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetEgressProxyRuleGroupAttributeRequest) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *GetEgressProxyRuleGroupAttributeRequest) SetEgressProxyRuleGroupId(v string) *GetEgressProxyRuleGroupAttributeRequest {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
