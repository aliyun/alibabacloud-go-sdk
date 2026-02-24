// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactiveConfigRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *DeactiveConfigRulesRequest
	GetCompliancePackId() *string
	SetConfigRuleIds(v string) *DeactiveConfigRulesRequest
	GetConfigRuleIds() *string
}

type DeactiveConfigRulesRequest struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleIds    *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s DeactiveConfigRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeactiveConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *DeactiveConfigRulesRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *DeactiveConfigRulesRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *DeactiveConfigRulesRequest) SetCompliancePackId(v string) *DeactiveConfigRulesRequest {
	s.CompliancePackId = &v
	return s
}

func (s *DeactiveConfigRulesRequest) SetConfigRuleIds(v string) *DeactiveConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *DeactiveConfigRulesRequest) Validate() error {
	return dara.Validate(s)
}
