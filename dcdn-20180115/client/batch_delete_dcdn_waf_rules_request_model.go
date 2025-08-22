// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnWafRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleIds(v string) *BatchDeleteDcdnWafRulesRequest
	GetRuleIds() *string
}

type BatchDeleteDcdnWafRulesRequest struct {
	// The IDs of the protection rules that you want to delete. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 20000001,20000002
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
}

func (s BatchDeleteDcdnWafRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnWafRulesRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnWafRulesRequest) GetRuleIds() *string {
	return s.RuleIds
}

func (s *BatchDeleteDcdnWafRulesRequest) SetRuleIds(v string) *BatchDeleteDcdnWafRulesRequest {
	s.RuleIds = &v
	return s
}

func (s *BatchDeleteDcdnWafRulesRequest) Validate() error {
	return dara.Validate(s)
}
