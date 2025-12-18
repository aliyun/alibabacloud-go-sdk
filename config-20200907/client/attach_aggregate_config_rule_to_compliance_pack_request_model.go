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
	// cp-0453626622af0020****
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
