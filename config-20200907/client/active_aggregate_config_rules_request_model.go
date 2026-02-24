// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveAggregateConfigRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ActiveAggregateConfigRulesRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *ActiveAggregateConfigRulesRequest
	GetCompliancePackId() *string
	SetConfigRuleIds(v string) *ActiveAggregateConfigRulesRequest
	GetConfigRuleIds() *string
}

type ActiveAggregateConfigRulesRequest struct {
	// This parameter is required.
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleIds    *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s ActiveAggregateConfigRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ActiveAggregateConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *ActiveAggregateConfigRulesRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ActiveAggregateConfigRulesRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ActiveAggregateConfigRulesRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *ActiveAggregateConfigRulesRequest) SetAggregatorId(v string) *ActiveAggregateConfigRulesRequest {
	s.AggregatorId = &v
	return s
}

func (s *ActiveAggregateConfigRulesRequest) SetCompliancePackId(v string) *ActiveAggregateConfigRulesRequest {
	s.CompliancePackId = &v
	return s
}

func (s *ActiveAggregateConfigRulesRequest) SetConfigRuleIds(v string) *ActiveAggregateConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *ActiveAggregateConfigRulesRequest) Validate() error {
	return dara.Validate(s)
}
