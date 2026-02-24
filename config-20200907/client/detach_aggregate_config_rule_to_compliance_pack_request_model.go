// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAggregateConfigRuleToCompliancePackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *DetachAggregateConfigRuleToCompliancePackRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *DetachAggregateConfigRuleToCompliancePackRequest
	GetCompliancePackId() *string
	SetConfigRuleIds(v string) *DetachAggregateConfigRuleToCompliancePackRequest
	GetConfigRuleIds() *string
}

type DetachAggregateConfigRuleToCompliancePackRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// This parameter is required.
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s DetachAggregateConfigRuleToCompliancePackRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachAggregateConfigRuleToCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *DetachAggregateConfigRuleToCompliancePackRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *DetachAggregateConfigRuleToCompliancePackRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *DetachAggregateConfigRuleToCompliancePackRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *DetachAggregateConfigRuleToCompliancePackRequest) SetAggregatorId(v string) *DetachAggregateConfigRuleToCompliancePackRequest {
	s.AggregatorId = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackRequest) SetCompliancePackId(v string) *DetachAggregateConfigRuleToCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackRequest) SetConfigRuleIds(v string) *DetachAggregateConfigRuleToCompliancePackRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackRequest) Validate() error {
	return dara.Validate(s)
}
