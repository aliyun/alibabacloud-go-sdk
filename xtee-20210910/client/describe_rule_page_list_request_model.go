// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRulePageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRulePageListRequest
	GetLang() *string
	SetCreateType(v string) *DescribeRulePageListRequest
	GetCreateType() *string
	SetCurrentPage(v string) *DescribeRulePageListRequest
	GetCurrentPage() *string
	SetEventCode(v string) *DescribeRulePageListRequest
	GetEventCode() *string
	SetPageSize(v string) *DescribeRulePageListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeRulePageListRequest
	GetRegId() *string
	SetRuleAuthType(v string) *DescribeRulePageListRequest
	GetRuleAuthType() *string
	SetRuleName(v string) *DescribeRulePageListRequest
	GetRuleName() *string
	SetRuleStatus(v string) *DescribeRulePageListRequest
	GetRuleStatus() *string
	SetSort(v string) *DescribeRulePageListRequest
	GetSort() *string
}

type DescribeRulePageListRequest struct {
	// Set the language type for requests and responses, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Creation type.
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_asssce8122
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy type.
	//
	// example:
	//
	// NOMAL
	RuleAuthType *string `json:"ruleAuthType,omitempty" xml:"ruleAuthType,omitempty"`
	// Policy name.
	//
	// example:
	//
	// 营销风险识别评分
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Policy status.
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// Sorting method, default value is desc.
	//
	// - desc: descending order
	//
	// - asc: ascending order
	//
	// example:
	//
	// asc
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s DescribeRulePageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulePageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRulePageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRulePageListRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeRulePageListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeRulePageListRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeRulePageListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeRulePageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRulePageListRequest) GetRuleAuthType() *string {
	return s.RuleAuthType
}

func (s *DescribeRulePageListRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeRulePageListRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeRulePageListRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeRulePageListRequest) SetLang(v string) *DescribeRulePageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRulePageListRequest) SetCreateType(v string) *DescribeRulePageListRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeRulePageListRequest) SetCurrentPage(v string) *DescribeRulePageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRulePageListRequest) SetEventCode(v string) *DescribeRulePageListRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeRulePageListRequest) SetPageSize(v string) *DescribeRulePageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRulePageListRequest) SetRegId(v string) *DescribeRulePageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRulePageListRequest) SetRuleAuthType(v string) *DescribeRulePageListRequest {
	s.RuleAuthType = &v
	return s
}

func (s *DescribeRulePageListRequest) SetRuleName(v string) *DescribeRulePageListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeRulePageListRequest) SetRuleStatus(v string) *DescribeRulePageListRequest {
	s.RuleStatus = &v
	return s
}

func (s *DescribeRulePageListRequest) SetSort(v string) *DescribeRulePageListRequest {
	s.Sort = &v
	return s
}

func (s *DescribeRulePageListRequest) Validate() error {
	return dara.Validate(s)
}
