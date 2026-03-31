// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateConfigRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *DeleteAggregateConfigRulesRequest
	GetAggregatorId() *string
	SetConfigRuleIds(v string) *DeleteAggregateConfigRulesRequest
	GetConfigRuleIds() *string
}

type DeleteAggregateConfigRulesRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a4e5626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The rule ID. Separate multiple rule IDs with commas (,).
	//
	// For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-4e3d626622af0080****
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s DeleteAggregateConfigRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigRulesRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *DeleteAggregateConfigRulesRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *DeleteAggregateConfigRulesRequest) SetAggregatorId(v string) *DeleteAggregateConfigRulesRequest {
	s.AggregatorId = &v
	return s
}

func (s *DeleteAggregateConfigRulesRequest) SetConfigRuleIds(v string) *DeleteAggregateConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *DeleteAggregateConfigRulesRequest) Validate() error {
	return dara.Validate(s)
}
