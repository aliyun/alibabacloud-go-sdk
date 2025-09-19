// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientUserDefineRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListClientUserDefineRulesResponseBodyPageInfo) *ListClientUserDefineRulesResponseBody
	GetPageInfo() *ListClientUserDefineRulesResponseBodyPageInfo
	SetRequestId(v string) *ListClientUserDefineRulesResponseBody
	GetRequestId() *string
	SetUserDefineRuleList(v []*ListClientUserDefineRulesResponseBodyUserDefineRuleList) *ListClientUserDefineRulesResponseBody
	GetUserDefineRuleList() []*ListClientUserDefineRulesResponseBodyUserDefineRuleList
}

type ListClientUserDefineRulesResponseBody struct {
	// The pagination information.
	PageInfo *ListClientUserDefineRulesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB393***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the rules.
	UserDefineRuleList []*ListClientUserDefineRulesResponseBodyUserDefineRuleList `json:"UserDefineRuleList,omitempty" xml:"UserDefineRuleList,omitempty" type:"Repeated"`
}

func (s ListClientUserDefineRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClientUserDefineRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientUserDefineRulesResponseBody) GetPageInfo() *ListClientUserDefineRulesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListClientUserDefineRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClientUserDefineRulesResponseBody) GetUserDefineRuleList() []*ListClientUserDefineRulesResponseBodyUserDefineRuleList {
	return s.UserDefineRuleList
}

func (s *ListClientUserDefineRulesResponseBody) SetPageInfo(v *ListClientUserDefineRulesResponseBodyPageInfo) *ListClientUserDefineRulesResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListClientUserDefineRulesResponseBody) SetRequestId(v string) *ListClientUserDefineRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientUserDefineRulesResponseBody) SetUserDefineRuleList(v []*ListClientUserDefineRulesResponseBodyUserDefineRuleList) *ListClientUserDefineRulesResponseBody {
	s.UserDefineRuleList = v
	return s
}

func (s *ListClientUserDefineRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClientUserDefineRulesResponseBodyPageInfo struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClientUserDefineRulesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListClientUserDefineRulesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListClientUserDefineRulesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClientUserDefineRulesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClientUserDefineRulesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListClientUserDefineRulesResponseBodyPageInfo) SetCurrentPage(v int32) *ListClientUserDefineRulesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListClientUserDefineRulesResponseBodyPageInfo) SetPageSize(v int32) *ListClientUserDefineRulesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListClientUserDefineRulesResponseBodyPageInfo) SetTotalCount(v int32) *ListClientUserDefineRulesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListClientUserDefineRulesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListClientUserDefineRulesResponseBodyUserDefineRuleList struct {
	// The action of the rule. Valid values:
	//
	// 	- **0**: allow
	//
	// 	- **1**: block
	//
	// example:
	//
	// 0
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 200****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// Rule\\*\\*\\*\\*
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- **windows**: Windows
	//
	// 	- **linux**: Linux
	//
	// 	- **all**: all types
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The switch ID of the rule.
	//
	// example:
	//
	// USER-DEFINE-RULE-SWITCH-TYPE_200****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **1**: Process hash
	//
	// 	- **2**: Command line
	//
	// 	- **3**: Process Network
	//
	// 	- **4**: File Read and Write
	//
	// 	- **5**: Operation on Registry
	//
	// 	- **6**: Dynamic-link Library Loading
	//
	// 	- **7**: File Renaming
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListClientUserDefineRulesResponseBodyUserDefineRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListClientUserDefineRulesResponseBodyUserDefineRuleList) GoString() string {
	return s.String()
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) GetActionType() *int32 {
	return s.ActionType
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) GetId() *int64 {
	return s.Id
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) GetName() *string {
	return s.Name
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) GetPlatform() *string {
	return s.Platform
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) GetSwitchId() *string {
	return s.SwitchId
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) GetType() *int32 {
	return s.Type
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) SetActionType(v int32) *ListClientUserDefineRulesResponseBodyUserDefineRuleList {
	s.ActionType = &v
	return s
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) SetId(v int64) *ListClientUserDefineRulesResponseBodyUserDefineRuleList {
	s.Id = &v
	return s
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) SetName(v string) *ListClientUserDefineRulesResponseBodyUserDefineRuleList {
	s.Name = &v
	return s
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) SetPlatform(v string) *ListClientUserDefineRulesResponseBodyUserDefineRuleList {
	s.Platform = &v
	return s
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) SetSwitchId(v string) *ListClientUserDefineRulesResponseBodyUserDefineRuleList {
	s.SwitchId = &v
	return s
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) SetType(v int32) *ListClientUserDefineRulesResponseBodyUserDefineRuleList {
	s.Type = &v
	return s
}

func (s *ListClientUserDefineRulesResponseBodyUserDefineRuleList) Validate() error {
	return dara.Validate(s)
}
