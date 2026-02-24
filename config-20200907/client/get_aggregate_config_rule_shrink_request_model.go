// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateConfigRuleShrinkRequest
	GetAggregatorId() *string
	SetConfigRuleId(v string) *GetAggregateConfigRuleShrinkRequest
	GetConfigRuleId() *string
	SetTagShrink(v string) *GetAggregateConfigRuleShrinkRequest
	GetTagShrink() *string
}

type GetAggregateConfigRuleShrinkRequest struct {
	// The ID of the account group.
	//
	// For more information, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-7f00626622af0041****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The rule ID.
	//
	// For more information, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-7f7d626622af0041****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// Deprecated
	//
	// The tags.
	//
	// This parameter is deprecated. If you specify this parameter, the value does not take effect.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetAggregateConfigRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateConfigRuleShrinkRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetAggregateConfigRuleShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *GetAggregateConfigRuleShrinkRequest) SetAggregatorId(v string) *GetAggregateConfigRuleShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRuleShrinkRequest) SetConfigRuleId(v string) *GetAggregateConfigRuleShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *GetAggregateConfigRuleShrinkRequest) SetTagShrink(v string) *GetAggregateConfigRuleShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *GetAggregateConfigRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
