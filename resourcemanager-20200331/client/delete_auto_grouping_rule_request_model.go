// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoGroupingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleId(v string) *DeleteAutoGroupingRuleRequest
	GetRuleId() *string
}

type DeleteAutoGroupingRuleRequest struct {
	// The ID of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// gr-acfo******hy6a
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteAutoGroupingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoGroupingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoGroupingRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteAutoGroupingRuleRequest) SetRuleId(v string) *DeleteAutoGroupingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteAutoGroupingRuleRequest) Validate() error {
	return dara.Validate(s)
}
