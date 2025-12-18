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
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-04b3fd170e340007****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
	//
	// example:
	//
	// cp-fe416457e0d90022****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The rule ID. Separate multiple rule IDs with commas (,).
	//
	// For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// example:
	//
	// cr-5772ba41209e007b****
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
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
