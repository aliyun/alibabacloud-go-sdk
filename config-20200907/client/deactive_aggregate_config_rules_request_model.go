// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactiveAggregateConfigRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *DeactiveAggregateConfigRulesRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *DeactiveAggregateConfigRulesRequest
	GetCompliancePackId() *string
	SetConfigRuleIds(v string) *DeactiveAggregateConfigRulesRequest
	GetConfigRuleIds() *string
}

type DeactiveAggregateConfigRulesRequest struct {
	// This parameter is required.
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleIds    *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s DeactiveAggregateConfigRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeactiveAggregateConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *DeactiveAggregateConfigRulesRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *DeactiveAggregateConfigRulesRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *DeactiveAggregateConfigRulesRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *DeactiveAggregateConfigRulesRequest) SetAggregatorId(v string) *DeactiveAggregateConfigRulesRequest {
	s.AggregatorId = &v
	return s
}

func (s *DeactiveAggregateConfigRulesRequest) SetCompliancePackId(v string) *DeactiveAggregateConfigRulesRequest {
	s.CompliancePackId = &v
	return s
}

func (s *DeactiveAggregateConfigRulesRequest) SetConfigRuleIds(v string) *DeactiveAggregateConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *DeactiveAggregateConfigRulesRequest) Validate() error {
	return dara.Validate(s)
}
