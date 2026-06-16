// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleVersionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRuleVersionListRequest
	GetLang() *string
	SetConsoleRuleId(v int64) *DescribeRuleVersionListRequest
	GetConsoleRuleId() *int64
	SetCurrentPage(v int32) *DescribeRuleVersionListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeRuleVersionListRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeRuleVersionListRequest
	GetRegId() *string
	SetRuleId(v string) *DescribeRuleVersionListRequest
	GetRuleId() *string
}

type DescribeRuleVersionListRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The primary key ID of the policy.
	//
	// example:
	//
	// 6851
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 101804
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s DescribeRuleVersionListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleVersionListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleVersionListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRuleVersionListRequest) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *DescribeRuleVersionListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRuleVersionListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRuleVersionListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRuleVersionListRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleVersionListRequest) SetLang(v string) *DescribeRuleVersionListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRuleVersionListRequest) SetConsoleRuleId(v int64) *DescribeRuleVersionListRequest {
	s.ConsoleRuleId = &v
	return s
}

func (s *DescribeRuleVersionListRequest) SetCurrentPage(v int32) *DescribeRuleVersionListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRuleVersionListRequest) SetPageSize(v int32) *DescribeRuleVersionListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRuleVersionListRequest) SetRegId(v string) *DescribeRuleVersionListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRuleVersionListRequest) SetRuleId(v string) *DescribeRuleVersionListRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleVersionListRequest) Validate() error {
	return dara.Validate(s)
}
