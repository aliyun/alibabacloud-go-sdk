// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoTagRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleIdList(v string) *DeleteAutoTagRulesRequest
	GetRuleIdList() *string
}

type DeleteAutoTagRulesRequest struct {
	// The ID of the asset auto-tagging rule. Separate multiple IDs with commas (,).
	//
	// >  You can call the [ListAutoTagRules](~~ListAutoTagRules~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2331,56,5644
	RuleIdList *string `json:"RuleIdList,omitempty" xml:"RuleIdList,omitempty"`
}

func (s DeleteAutoTagRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoTagRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoTagRulesRequest) GetRuleIdList() *string {
	return s.RuleIdList
}

func (s *DeleteAutoTagRulesRequest) SetRuleIdList(v string) *DeleteAutoTagRulesRequest {
	s.RuleIdList = &v
	return s
}

func (s *DeleteAutoTagRulesRequest) Validate() error {
	return dara.Validate(s)
}
