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
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// example:
	//
	// cp-fe416457e0d90022****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The ID of the rule. Separate multiple rule IDs with commas (,).
	//
	// For more information about how to obtain the ID of a rule, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// example:
	//
	// cr-19a56457e0d90058****
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
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
