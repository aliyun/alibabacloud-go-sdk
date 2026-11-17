// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebFilterRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *ListWebFilterRulesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListWebFilterRulesRequest
	GetPageSize() *int64
	SetRuleAction(v string) *ListWebFilterRulesRequest
	GetRuleAction() *string
	SetWebFilterRuleId(v string) *ListWebFilterRulesRequest
	GetWebFilterRuleId() *string
	SetWebFilterTemplateId(v string) *ListWebFilterRulesRequest
	GetWebFilterTemplateId() *string
}

type ListWebFilterRulesRequest struct {
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The value must be greater than or equal to 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The rule action for exact matching. Valid values:
	//
	// - **Allow**: Allows the request.
	//
	// - **Alert**: Monitors the request.
	//
	// - **Deny**: Blocks the request.
	//
	// example:
	//
	// Deny
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The ID of the web filtering rule for exact matching.
	//
	// example:
	//
	// wfr-8c1d2e3f4a5b6c7d
	WebFilterRuleId *string `json:"WebFilterRuleId,omitempty" xml:"WebFilterRuleId,omitempty"`
	// The ID of the web filtering template for exact matching.
	//
	// example:
	//
	// wft-3f2a9c1d4b5e6f70
	WebFilterTemplateId *string `json:"WebFilterTemplateId,omitempty" xml:"WebFilterTemplateId,omitempty"`
}

func (s ListWebFilterRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWebFilterRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWebFilterRulesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListWebFilterRulesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWebFilterRulesRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *ListWebFilterRulesRequest) GetWebFilterRuleId() *string {
	return s.WebFilterRuleId
}

func (s *ListWebFilterRulesRequest) GetWebFilterTemplateId() *string {
	return s.WebFilterTemplateId
}

func (s *ListWebFilterRulesRequest) SetPageNo(v int64) *ListWebFilterRulesRequest {
	s.PageNo = &v
	return s
}

func (s *ListWebFilterRulesRequest) SetPageSize(v int64) *ListWebFilterRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWebFilterRulesRequest) SetRuleAction(v string) *ListWebFilterRulesRequest {
	s.RuleAction = &v
	return s
}

func (s *ListWebFilterRulesRequest) SetWebFilterRuleId(v string) *ListWebFilterRulesRequest {
	s.WebFilterRuleId = &v
	return s
}

func (s *ListWebFilterRulesRequest) SetWebFilterTemplateId(v string) *ListWebFilterRulesRequest {
	s.WebFilterTemplateId = &v
	return s
}

func (s *ListWebFilterRulesRequest) Validate() error {
	return dara.Validate(s)
}
