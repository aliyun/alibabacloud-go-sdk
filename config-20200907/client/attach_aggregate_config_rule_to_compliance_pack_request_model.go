// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAggregateConfigRuleToCompliancePackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *AttachAggregateConfigRuleToCompliancePackRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *AttachAggregateConfigRuleToCompliancePackRequest
	GetCompliancePackId() *string
	SetConfigRuleIds(v string) *AttachAggregateConfigRuleToCompliancePackRequest
	GetConfigRuleIds() *string
}

type AttachAggregateConfigRuleToCompliancePackRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// This parameter is required.
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s AttachAggregateConfigRuleToCompliancePackRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachAggregateConfigRuleToCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *AttachAggregateConfigRuleToCompliancePackRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *AttachAggregateConfigRuleToCompliancePackRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *AttachAggregateConfigRuleToCompliancePackRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *AttachAggregateConfigRuleToCompliancePackRequest) SetAggregatorId(v string) *AttachAggregateConfigRuleToCompliancePackRequest {
	s.AggregatorId = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackRequest) SetCompliancePackId(v string) *AttachAggregateConfigRuleToCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackRequest) SetConfigRuleIds(v string) *AttachAggregateConfigRuleToCompliancePackRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackRequest) Validate() error {
	return dara.Validate(s)
}
