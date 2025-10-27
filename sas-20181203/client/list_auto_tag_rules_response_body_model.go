// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoTagRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoTagRuleList(v []*ListAutoTagRulesResponseBodyAutoTagRuleList) *ListAutoTagRulesResponseBody
	GetAutoTagRuleList() []*ListAutoTagRulesResponseBodyAutoTagRuleList
	SetPageInfo(v *ListAutoTagRulesResponseBodyPageInfo) *ListAutoTagRulesResponseBody
	GetPageInfo() *ListAutoTagRulesResponseBodyPageInfo
	SetRequestId(v string) *ListAutoTagRulesResponseBody
	GetRequestId() *string
}

type ListAutoTagRulesResponseBody struct {
	// The asset auto-tagging rules.
	AutoTagRuleList []*ListAutoTagRulesResponseBodyAutoTagRuleList `json:"AutoTagRuleList,omitempty" xml:"AutoTagRuleList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListAutoTagRulesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1383B0DB-D5D6-4B0C-9E6B-75939C8E67FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAutoTagRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutoTagRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoTagRulesResponseBody) GetAutoTagRuleList() []*ListAutoTagRulesResponseBodyAutoTagRuleList {
	return s.AutoTagRuleList
}

func (s *ListAutoTagRulesResponseBody) GetPageInfo() *ListAutoTagRulesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListAutoTagRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAutoTagRulesResponseBody) SetAutoTagRuleList(v []*ListAutoTagRulesResponseBodyAutoTagRuleList) *ListAutoTagRulesResponseBody {
	s.AutoTagRuleList = v
	return s
}

func (s *ListAutoTagRulesResponseBody) SetPageInfo(v *ListAutoTagRulesResponseBodyPageInfo) *ListAutoTagRulesResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListAutoTagRulesResponseBody) SetRequestId(v string) *ListAutoTagRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoTagRulesResponseBody) Validate() error {
	if s.AutoTagRuleList != nil {
		for _, item := range s.AutoTagRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAutoTagRulesResponseBodyAutoTagRuleList struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 558463566374****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The timestamp when the rule was created. Unit: milliseconds.
	//
	// example:
	//
	// 1614674401000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The expression of the rule.
	//
	// example:
	//
	// [{\\"groups\\":\\"0\\",\\"fieldValueType\\":\\"string\\",\\"field\\":\\"internetIp\\",\\"operator\\":\\"equals\\",\\"value\\":\\"12.0.0.1\\"}]
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 900029
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The timestamp when the rule was last updated. Unit: milliseconds.
	//
	// example:
	//
	// 1614674401000
	ModifiedTimestamp *int64 `json:"ModifiedTimestamp,omitempty" xml:"ModifiedTimestamp,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// describe
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// auto_test_rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The tag specified by the operation type of the rule.
	//
	// example:
	//
	// {\\"tagId\\":4577447}
	TagContext *string `json:"TagContext,omitempty" xml:"TagContext,omitempty"`
	// The operation type of the rule. Valid values:
	//
	// 	- **group**
	//
	// 	- **tag**
	//
	// example:
	//
	// group
	TagType *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
}

func (s ListAutoTagRulesResponseBodyAutoTagRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListAutoTagRulesResponseBodyAutoTagRuleList) GoString() string {
	return s.String()
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) GetExpression() *string {
	return s.Expression
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) GetId() *int32 {
	return s.Id
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) GetModifiedTimestamp() *int64 {
	return s.ModifiedTimestamp
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) GetTagContext() *string {
	return s.TagContext
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) GetTagType() *string {
	return s.TagType
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) SetAliUid(v int64) *ListAutoTagRulesResponseBodyAutoTagRuleList {
	s.AliUid = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) SetCreateTimestamp(v int64) *ListAutoTagRulesResponseBodyAutoTagRuleList {
	s.CreateTimestamp = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) SetExpression(v string) *ListAutoTagRulesResponseBodyAutoTagRuleList {
	s.Expression = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) SetId(v int32) *ListAutoTagRulesResponseBodyAutoTagRuleList {
	s.Id = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) SetModifiedTimestamp(v int64) *ListAutoTagRulesResponseBodyAutoTagRuleList {
	s.ModifiedTimestamp = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) SetRuleDesc(v string) *ListAutoTagRulesResponseBodyAutoTagRuleList {
	s.RuleDesc = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) SetRuleName(v string) *ListAutoTagRulesResponseBodyAutoTagRuleList {
	s.RuleName = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) SetTagContext(v string) *ListAutoTagRulesResponseBodyAutoTagRuleList {
	s.TagContext = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) SetTagType(v string) *ListAutoTagRulesResponseBodyAutoTagRuleList {
	s.TagType = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyAutoTagRuleList) Validate() error {
	return dara.Validate(s)
}

type ListAutoTagRulesResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 196
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAutoTagRulesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAutoTagRulesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListAutoTagRulesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListAutoTagRulesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAutoTagRulesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAutoTagRulesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAutoTagRulesResponseBodyPageInfo) SetCount(v int32) *ListAutoTagRulesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyPageInfo) SetCurrentPage(v int32) *ListAutoTagRulesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyPageInfo) SetPageSize(v int32) *ListAutoTagRulesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyPageInfo) SetTotalCount(v int32) *ListAutoTagRulesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListAutoTagRulesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
