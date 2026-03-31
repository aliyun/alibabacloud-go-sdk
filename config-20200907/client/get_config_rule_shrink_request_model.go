// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *GetConfigRuleShrinkRequest
	GetConfigRuleId() *string
	SetTagShrink(v string) *GetConfigRuleShrinkRequest
	GetTagShrink() *string
}

type GetConfigRuleShrinkRequest struct {
	// The rule ID.
	//
	// For more information about how to obtain the ID of a rule, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-7f7d626622af0041****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// Deprecated
	//
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetConfigRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetConfigRuleShrinkRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetConfigRuleShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *GetConfigRuleShrinkRequest) SetConfigRuleId(v string) *GetConfigRuleShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *GetConfigRuleShrinkRequest) SetTagShrink(v string) *GetConfigRuleShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *GetConfigRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
