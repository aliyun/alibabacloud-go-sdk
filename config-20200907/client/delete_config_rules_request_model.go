// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleIds(v string) *DeleteConfigRulesRequest
	GetConfigRuleIds() *string
}

type DeleteConfigRulesRequest struct {
	// The rule IDs. Separate multiple rule IDs with commas (,).
	//
	// For more information about how to obtain the ID of a rule, see [ListConfigRules](https://help.aliyun.com/document_detail/609222.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-9908626622af0035****
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s DeleteConfigRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesRequest) GetConfigRuleIds() *string {
	return s.ConfigRuleIds
}

func (s *DeleteConfigRulesRequest) SetConfigRuleIds(v string) *DeleteConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

func (s *DeleteConfigRulesRequest) Validate() error {
	return dara.Validate(s)
}
