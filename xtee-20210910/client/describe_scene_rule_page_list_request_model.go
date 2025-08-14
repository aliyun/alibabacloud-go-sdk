// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneRulePageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSceneRulePageListRequest
	GetLang() *string
	SetCreateType(v string) *DescribeSceneRulePageListRequest
	GetCreateType() *string
	SetCurrentPage(v string) *DescribeSceneRulePageListRequest
	GetCurrentPage() *string
	SetEventCode(v string) *DescribeSceneRulePageListRequest
	GetEventCode() *string
	SetPageSize(v string) *DescribeSceneRulePageListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeSceneRulePageListRequest
	GetRegId() *string
	SetRuleAuthType(v string) *DescribeSceneRulePageListRequest
	GetRuleAuthType() *string
	SetRuleName(v string) *DescribeSceneRulePageListRequest
	GetRuleName() *string
	SetRuleStatus(v string) *DescribeSceneRulePageListRequest
	GetRuleStatus() *string
}

type DescribeSceneRulePageListRequest struct {
	// Set the language type for requests and received messages. Default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Creation type
	//
	// example:
	//
	// NOMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Event code
	//
	// example:
	//
	// de_ahgctb7098
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Number of items per page in the returned results. Default value: 20, minimum value: 1, maximum value: 50.
	//
	// example:
	//
	// 20
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Strategy type
	//
	// example:
	//
	// CUSTMER
	RuleAuthType *string `json:"ruleAuthType,omitempty" xml:"ruleAuthType,omitempty"`
	// Strategy name
	//
	// example:
	//
	// 营销风险识别
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Strategy status
	//
	// example:
	//
	// RUNNING
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
}

func (s DescribeSceneRulePageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneRulePageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSceneRulePageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSceneRulePageListRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeSceneRulePageListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeSceneRulePageListRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeSceneRulePageListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSceneRulePageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSceneRulePageListRequest) GetRuleAuthType() *string {
	return s.RuleAuthType
}

func (s *DescribeSceneRulePageListRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeSceneRulePageListRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeSceneRulePageListRequest) SetLang(v string) *DescribeSceneRulePageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSceneRulePageListRequest) SetCreateType(v string) *DescribeSceneRulePageListRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeSceneRulePageListRequest) SetCurrentPage(v string) *DescribeSceneRulePageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSceneRulePageListRequest) SetEventCode(v string) *DescribeSceneRulePageListRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeSceneRulePageListRequest) SetPageSize(v string) *DescribeSceneRulePageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSceneRulePageListRequest) SetRegId(v string) *DescribeSceneRulePageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSceneRulePageListRequest) SetRuleAuthType(v string) *DescribeSceneRulePageListRequest {
	s.RuleAuthType = &v
	return s
}

func (s *DescribeSceneRulePageListRequest) SetRuleName(v string) *DescribeSceneRulePageListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeSceneRulePageListRequest) SetRuleStatus(v string) *DescribeSceneRulePageListRequest {
	s.RuleStatus = &v
	return s
}

func (s *DescribeSceneRulePageListRequest) Validate() error {
	return dara.Validate(s)
}
