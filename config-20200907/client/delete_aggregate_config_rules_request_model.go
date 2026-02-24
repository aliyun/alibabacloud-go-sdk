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
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
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
