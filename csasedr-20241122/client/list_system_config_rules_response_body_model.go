// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemConfigRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListSystemConfigRulesResponseBodyList) *ListSystemConfigRulesResponseBody
	GetList() []*ListSystemConfigRulesResponseBodyList
	SetPageInfo(v *ListSystemConfigRulesResponseBodyPageInfo) *ListSystemConfigRulesResponseBody
	GetPageInfo() *ListSystemConfigRulesResponseBodyPageInfo
	SetRequestId(v string) *ListSystemConfigRulesResponseBody
	GetRequestId() *string
}

type ListSystemConfigRulesResponseBody struct {
	// The rule list.
	List []*ListSystemConfigRulesResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The paging information displayed during a paged query.
	PageInfo *ListSystemConfigRulesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The unique identifier of the request, which is used to troubleshoot and locate issues.
	//
	// example:
	//
	// 5B1AACF9-E87B-577A-9B12-655A2E0C96FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSystemConfigRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSystemConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemConfigRulesResponseBody) GetList() []*ListSystemConfigRulesResponseBodyList {
	return s.List
}

func (s *ListSystemConfigRulesResponseBody) GetPageInfo() *ListSystemConfigRulesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListSystemConfigRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSystemConfigRulesResponseBody) SetList(v []*ListSystemConfigRulesResponseBodyList) *ListSystemConfigRulesResponseBody {
	s.List = v
	return s
}

func (s *ListSystemConfigRulesResponseBody) SetPageInfo(v *ListSystemConfigRulesResponseBodyPageInfo) *ListSystemConfigRulesResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListSystemConfigRulesResponseBody) SetRequestId(v string) *ListSystemConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemConfigRulesResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
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

type ListSystemConfigRulesResponseBodyList struct {
	// The rule ID.
	//
	// example:
	//
	// 204
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The platform name of the system rule. Valid values:
	//
	// - **1**: Linux.
	//
	// - **2**: Windows.
	//
	// - **3**: Mac.
	//
	// example:
	//
	// windows
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The rule name.
	//
	// example:
	//
	// Host defense link test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule type.
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The switch status of the rule under the specified SelectionKey. Valid values:
	//
	// - **add**: Enabled.
	//
	// - **del**: Disabled.
	//
	// example:
	//
	// add
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The switch identifier of the system rule.
	//
	// example:
	//
	// USER-ENABLE-SWITCH-TYPE_1
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
}

func (s ListSystemConfigRulesResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListSystemConfigRulesResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSystemConfigRulesResponseBodyList) GetId() *int32 {
	return s.Id
}

func (s *ListSystemConfigRulesResponseBodyList) GetPlatform() *string {
	return s.Platform
}

func (s *ListSystemConfigRulesResponseBodyList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSystemConfigRulesResponseBodyList) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ListSystemConfigRulesResponseBodyList) GetStatus() *string {
	return s.Status
}

func (s *ListSystemConfigRulesResponseBodyList) GetSwitchId() *string {
	return s.SwitchId
}

func (s *ListSystemConfigRulesResponseBodyList) SetId(v int32) *ListSystemConfigRulesResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListSystemConfigRulesResponseBodyList) SetPlatform(v string) *ListSystemConfigRulesResponseBodyList {
	s.Platform = &v
	return s
}

func (s *ListSystemConfigRulesResponseBodyList) SetRuleName(v string) *ListSystemConfigRulesResponseBodyList {
	s.RuleName = &v
	return s
}

func (s *ListSystemConfigRulesResponseBodyList) SetRuleType(v int32) *ListSystemConfigRulesResponseBodyList {
	s.RuleType = &v
	return s
}

func (s *ListSystemConfigRulesResponseBodyList) SetStatus(v string) *ListSystemConfigRulesResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListSystemConfigRulesResponseBodyList) SetSwitchId(v string) *ListSystemConfigRulesResponseBodyList {
	s.SwitchId = &v
	return s
}

func (s *ListSystemConfigRulesResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListSystemConfigRulesResponseBodyPageInfo struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries to return on each page during paging in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 122
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSystemConfigRulesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSystemConfigRulesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListSystemConfigRulesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListSystemConfigRulesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSystemConfigRulesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSystemConfigRulesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSystemConfigRulesResponseBodyPageInfo) SetCount(v int32) *ListSystemConfigRulesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListSystemConfigRulesResponseBodyPageInfo) SetCurrentPage(v int32) *ListSystemConfigRulesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListSystemConfigRulesResponseBodyPageInfo) SetPageSize(v int32) *ListSystemConfigRulesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListSystemConfigRulesResponseBodyPageInfo) SetTotalCount(v int32) *ListSystemConfigRulesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListSystemConfigRulesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
