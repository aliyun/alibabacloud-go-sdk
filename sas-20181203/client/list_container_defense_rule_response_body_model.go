// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContainerDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContainerDefenseRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListContainerDefenseRuleResponseBody
	GetHttpStatusCode() *int32
	SetList(v []*ListContainerDefenseRuleResponseBodyList) *ListContainerDefenseRuleResponseBody
	GetList() []*ListContainerDefenseRuleResponseBodyList
	SetMessage(v string) *ListContainerDefenseRuleResponseBody
	GetMessage() *string
	SetPageInfo(v *ListContainerDefenseRuleResponseBodyPageInfo) *ListContainerDefenseRuleResponseBody
	GetPageInfo() *ListContainerDefenseRuleResponseBodyPageInfo
	SetRequestId(v string) *ListContainerDefenseRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContainerDefenseRuleResponseBody
	GetSuccess() *bool
}

type ListContainerDefenseRuleResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The rules.
	List []*ListContainerDefenseRuleResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListContainerDefenseRuleResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5E3A63BA-***843
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListContainerDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContainerDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListContainerDefenseRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContainerDefenseRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListContainerDefenseRuleResponseBody) GetList() []*ListContainerDefenseRuleResponseBodyList {
	return s.List
}

func (s *ListContainerDefenseRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContainerDefenseRuleResponseBody) GetPageInfo() *ListContainerDefenseRuleResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListContainerDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContainerDefenseRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContainerDefenseRuleResponseBody) SetCode(v string) *ListContainerDefenseRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBody) SetHttpStatusCode(v int32) *ListContainerDefenseRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBody) SetList(v []*ListContainerDefenseRuleResponseBodyList) *ListContainerDefenseRuleResponseBody {
	s.List = v
	return s
}

func (s *ListContainerDefenseRuleResponseBody) SetMessage(v string) *ListContainerDefenseRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBody) SetPageInfo(v *ListContainerDefenseRuleResponseBodyPageInfo) *ListContainerDefenseRuleResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListContainerDefenseRuleResponseBody) SetRequestId(v string) *ListContainerDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBody) SetSuccess(v bool) *ListContainerDefenseRuleResponseBody {
	s.Success = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBody) Validate() error {
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

type ListContainerDefenseRuleResponseBodyList struct {
	// The total number of clusters.
	//
	// example:
	//
	// 1
	ClusterCount *int32 `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	// The clusters specified in the rule.
	//
	// example:
	//
	// cfb41a8**8a106
	ClusterIdList *string `json:"ClusterIdList,omitempty" xml:"ClusterIdList,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// defense rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The action specified in the rule. Valid values:
	//
	// 	- **1**: alert
	//
	// 	- **2**: block
	//
	// example:
	//
	// 1
	RuleAction *int32 `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 181
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test-rule-01
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 0
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **1**: system rule
	//
	// 	- **2**: custom rule
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListContainerDefenseRuleResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListContainerDefenseRuleResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListContainerDefenseRuleResponseBodyList) GetClusterCount() *int32 {
	return s.ClusterCount
}

func (s *ListContainerDefenseRuleResponseBodyList) GetClusterIdList() *string {
	return s.ClusterIdList
}

func (s *ListContainerDefenseRuleResponseBodyList) GetDescription() *string {
	return s.Description
}

func (s *ListContainerDefenseRuleResponseBodyList) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *ListContainerDefenseRuleResponseBodyList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListContainerDefenseRuleResponseBodyList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListContainerDefenseRuleResponseBodyList) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *ListContainerDefenseRuleResponseBodyList) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ListContainerDefenseRuleResponseBodyList) SetClusterCount(v int32) *ListContainerDefenseRuleResponseBodyList {
	s.ClusterCount = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyList) SetClusterIdList(v string) *ListContainerDefenseRuleResponseBodyList {
	s.ClusterIdList = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyList) SetDescription(v string) *ListContainerDefenseRuleResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyList) SetRuleAction(v int32) *ListContainerDefenseRuleResponseBodyList {
	s.RuleAction = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyList) SetRuleId(v int64) *ListContainerDefenseRuleResponseBodyList {
	s.RuleId = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyList) SetRuleName(v string) *ListContainerDefenseRuleResponseBodyList {
	s.RuleName = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyList) SetRuleSwitch(v int32) *ListContainerDefenseRuleResponseBodyList {
	s.RuleSwitch = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyList) SetRuleType(v int32) *ListContainerDefenseRuleResponseBodyList {
	s.RuleType = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListContainerDefenseRuleResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 9
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The key of the last data entry.
	//
	// example:
	//
	// CAESGgo***jE2NDc4NjE=
	LastRowKey *string `json:"LastRowKey,omitempty" xml:"LastRowKey,omitempty"`
	// The query credential.
	//
	// example:
	//
	// B60***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// 45
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListContainerDefenseRuleResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListContainerDefenseRuleResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) GetLastRowKey() *string {
	return s.LastRowKey
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) SetCount(v int32) *ListContainerDefenseRuleResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) SetCurrentPage(v int32) *ListContainerDefenseRuleResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) SetLastRowKey(v string) *ListContainerDefenseRuleResponseBodyPageInfo {
	s.LastRowKey = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) SetNextToken(v string) *ListContainerDefenseRuleResponseBodyPageInfo {
	s.NextToken = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) SetPageSize(v int32) *ListContainerDefenseRuleResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) SetTotalCount(v int32) *ListContainerDefenseRuleResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListContainerDefenseRuleResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
