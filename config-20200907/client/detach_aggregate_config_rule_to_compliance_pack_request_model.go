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
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-75b4626622af00c3****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-5bb1626622af00bd****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The rule IDs. Separate multiple rule IDs with commas (,).
	//
	// For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-a124626622af00e7****
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
