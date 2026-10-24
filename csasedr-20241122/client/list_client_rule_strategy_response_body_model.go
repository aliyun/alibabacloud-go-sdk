// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientRuleStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListClientRuleStrategyResponseBodyList) *ListClientRuleStrategyResponseBody
	GetList() []*ListClientRuleStrategyResponseBodyList
	SetPageInfo(v *ListClientRuleStrategyResponseBodyPageInfo) *ListClientRuleStrategyResponseBody
	GetPageInfo() *ListClientRuleStrategyResponseBodyPageInfo
	SetRequestId(v string) *ListClientRuleStrategyResponseBody
	GetRequestId() *string
}

type ListClientRuleStrategyResponseBody struct {
	// The list of policies.
	List []*ListClientRuleStrategyResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListClientRuleStrategyResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 45B2D3AF-0E5F-5D67-944C-2AD5E354B5C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClientRuleStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClientRuleStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientRuleStrategyResponseBody) GetList() []*ListClientRuleStrategyResponseBodyList {
	return s.List
}

func (s *ListClientRuleStrategyResponseBody) GetPageInfo() *ListClientRuleStrategyResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListClientRuleStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClientRuleStrategyResponseBody) SetList(v []*ListClientRuleStrategyResponseBodyList) *ListClientRuleStrategyResponseBody {
	s.List = v
	return s
}

func (s *ListClientRuleStrategyResponseBody) SetPageInfo(v *ListClientRuleStrategyResponseBodyPageInfo) *ListClientRuleStrategyResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListClientRuleStrategyResponseBody) SetRequestId(v string) *ListClientRuleStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientRuleStrategyResponseBody) Validate() error {
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

type ListClientRuleStrategyResponseBodyList struct {
	// The policy action. Valid values:
	//
	// - **0**: Monitor mode.
	//
	// - **1**: Block startup.
	//
	// - **2**: Block mode.
	//
	// Valid values for custom rules of proactive defense (PolicyType=2):
	//
	// - **0**: Allow.
	//
	// - **1**: Block.
	//
	// example:
	//
	// 1
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// Specifies whether the policy applies to all user groups. Valid values:
	//
	// - **0**: No.
	//
	// - **1**: Yes.
	//
	// example:
	//
	// 0
	AllGroup *int32 `json:"AllGroup,omitempty" xml:"AllGroup,omitempty"`
	// The instance ID of the approval.
	//
	// example:
	//
	// approval-3f9bcedb9643b956
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// The timestamp when the policy was created. Unit: milliseconds.
	//
	// example:
	//
	// 1792627200000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of usernames to which the policy does not apply.
	ExcludeUserNameList []*string `json:"ExcludeUserNameList,omitempty" xml:"ExcludeUserNameList,omitempty" type:"Repeated"`
	// The ID of the policy.
	//
	// example:
	//
	// 1938
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The switch for inbound network traffic.
	//
	// example:
	//
	// true
	InSwitch *bool `json:"InSwitch,omitempty" xml:"InSwitch,omitempty"`
	// The list of targets in the effective scope.
	IncludeGroupIdList []*string `json:"IncludeGroupIdList,omitempty" xml:"IncludeGroupIdList,omitempty" type:"Repeated"`
	// The text of the primary button in Chinese.
	//
	// example:
	//
	// test
	MainButtonTextCh *string `json:"MainButtonTextCh,omitempty" xml:"MainButtonTextCh,omitempty"`
	// The text of the primary button in English.
	//
	// example:
	//
	// test
	MainButtonTextEn *string `json:"MainButtonTextEn,omitempty" xml:"MainButtonTextEn,omitempty"`
	// The effective scope mode. Valid values:
	//
	// - **UserGroupAll**: All user groups.
	//
	// - **UserGroupNormal**: Specified user groups.
	//
	// - **DeviceGroupNormal**: Specified device groups.
	//
	// - **DevTagNormal**: Specified device tags.
	//
	// The meaning of MatchTargetIds is determined by this field.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The list of targets in the effective scope. UserGroupNormal corresponds to user group IDs, DeviceGroupNormal corresponds to device group IDs, and DevTagNormal corresponds to device tags. If MatchMode is set to UserGroupAll, you do not need to specify targets.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
	// The text of the secondary button in Chinese.
	//
	// example:
	//
	// test
	MinorButtonTextCh *string `json:"MinorButtonTextCh,omitempty" xml:"MinorButtonTextCh,omitempty"`
	// The text of the secondary button in English.
	//
	// example:
	//
	// test
	MinorButtonTextEn *string `json:"MinorButtonTextEn,omitempty" xml:"MinorButtonTextEn,omitempty"`
	// The Policy Name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The switch for outbound network traffic.
	//
	// example:
	//
	// true
	OutSwitch *bool `json:"OutSwitch,omitempty" xml:"OutSwitch,omitempty"`
	// The policy type. Valid values:
	//
	// - **0**: Software hardening.
	//
	// - **1**: Firewall.
	//
	// - **2**: Proactive defense.
	//
	// example:
	//
	// 0
	PolicyType *int32 `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The list of enabled ports.
	PortOnList []*string `json:"PortOnList,omitempty" xml:"PortOnList,omitempty" type:"Repeated"`
	// The policy priority. A smaller value indicates a higher priority. Valid values: greater than 0.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The association key of the policy configuration, which is used to associate the policy with its process or switch configuration.
	//
	// example:
	//
	// a8d64565-ab8b-41a6-bd3d-14d0fc639285
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
	// The list of main process names of the software for which integrity check is enabled. The elements correspond to the Process field returned by GetSwitchInfoListByType.
	SoftWareOnList []*string `json:"SoftWareOnList,omitempty" xml:"SoftWareOnList,omitempty" type:"Repeated"`
	// The status of the policy. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The source of the rule for proactive defense (PolicyType=2). Valid values:
	//
	// - **SYSTEM**: System rule.
	//
	// - **CUSTOM**: Custom rule.
	//
	// example:
	//
	// SYSTEM
	StrategySourceType *string `json:"StrategySourceType,omitempty" xml:"StrategySourceType,omitempty"`
	// Specifies whether reporting and approval are supported. Valid values:
	//
	// - **0**: Not supported.
	//
	// - **1**: Supported.
	//
	// example:
	//
	// 1
	SupportsApproval *int32 `json:"SupportsApproval,omitempty" xml:"SupportsApproval,omitempty"`
	// The content of the pop-up prompt in English.
	//
	// example:
	//
	// test
	UsContent *string `json:"UsContent,omitempty" xml:"UsContent,omitempty"`
	// The title of the pop-up prompt in English.
	//
	// example:
	//
	// test
	UsTitle *string `json:"UsTitle,omitempty" xml:"UsTitle,omitempty"`
	// The content of the pop-up prompt in Chinese.
	//
	// example:
	//
	// test
	ZhContent *string `json:"ZhContent,omitempty" xml:"ZhContent,omitempty"`
	// The title of the pop-up prompt in Chinese.
	//
	// example:
	//
	// test
	ZhTitle *string `json:"ZhTitle,omitempty" xml:"ZhTitle,omitempty"`
}

func (s ListClientRuleStrategyResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListClientRuleStrategyResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListClientRuleStrategyResponseBodyList) GetActionType() *int32 {
	return s.ActionType
}

func (s *ListClientRuleStrategyResponseBodyList) GetAllGroup() *int32 {
	return s.AllGroup
}

func (s *ListClientRuleStrategyResponseBodyList) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *ListClientRuleStrategyResponseBodyList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListClientRuleStrategyResponseBodyList) GetDescription() *string {
	return s.Description
}

func (s *ListClientRuleStrategyResponseBodyList) GetExcludeUserNameList() []*string {
	return s.ExcludeUserNameList
}

func (s *ListClientRuleStrategyResponseBodyList) GetId() *int64 {
	return s.Id
}

func (s *ListClientRuleStrategyResponseBodyList) GetInSwitch() *bool {
	return s.InSwitch
}

func (s *ListClientRuleStrategyResponseBodyList) GetIncludeGroupIdList() []*string {
	return s.IncludeGroupIdList
}

func (s *ListClientRuleStrategyResponseBodyList) GetMainButtonTextCh() *string {
	return s.MainButtonTextCh
}

func (s *ListClientRuleStrategyResponseBodyList) GetMainButtonTextEn() *string {
	return s.MainButtonTextEn
}

func (s *ListClientRuleStrategyResponseBodyList) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListClientRuleStrategyResponseBodyList) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *ListClientRuleStrategyResponseBodyList) GetMinorButtonTextCh() *string {
	return s.MinorButtonTextCh
}

func (s *ListClientRuleStrategyResponseBodyList) GetMinorButtonTextEn() *string {
	return s.MinorButtonTextEn
}

func (s *ListClientRuleStrategyResponseBodyList) GetName() *string {
	return s.Name
}

func (s *ListClientRuleStrategyResponseBodyList) GetOutSwitch() *bool {
	return s.OutSwitch
}

func (s *ListClientRuleStrategyResponseBodyList) GetPolicyType() *int32 {
	return s.PolicyType
}

func (s *ListClientRuleStrategyResponseBodyList) GetPortOnList() []*string {
	return s.PortOnList
}

func (s *ListClientRuleStrategyResponseBodyList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListClientRuleStrategyResponseBodyList) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *ListClientRuleStrategyResponseBodyList) GetSoftWareOnList() []*string {
	return s.SoftWareOnList
}

func (s *ListClientRuleStrategyResponseBodyList) GetStatus() *int32 {
	return s.Status
}

func (s *ListClientRuleStrategyResponseBodyList) GetStrategySourceType() *string {
	return s.StrategySourceType
}

func (s *ListClientRuleStrategyResponseBodyList) GetSupportsApproval() *int32 {
	return s.SupportsApproval
}

func (s *ListClientRuleStrategyResponseBodyList) GetUsContent() *string {
	return s.UsContent
}

func (s *ListClientRuleStrategyResponseBodyList) GetUsTitle() *string {
	return s.UsTitle
}

func (s *ListClientRuleStrategyResponseBodyList) GetZhContent() *string {
	return s.ZhContent
}

func (s *ListClientRuleStrategyResponseBodyList) GetZhTitle() *string {
	return s.ZhTitle
}

func (s *ListClientRuleStrategyResponseBodyList) SetActionType(v int32) *ListClientRuleStrategyResponseBodyList {
	s.ActionType = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetAllGroup(v int32) *ListClientRuleStrategyResponseBodyList {
	s.AllGroup = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetApprovalId(v string) *ListClientRuleStrategyResponseBodyList {
	s.ApprovalId = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetCreateTime(v int64) *ListClientRuleStrategyResponseBodyList {
	s.CreateTime = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetDescription(v string) *ListClientRuleStrategyResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetExcludeUserNameList(v []*string) *ListClientRuleStrategyResponseBodyList {
	s.ExcludeUserNameList = v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetId(v int64) *ListClientRuleStrategyResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetInSwitch(v bool) *ListClientRuleStrategyResponseBodyList {
	s.InSwitch = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetIncludeGroupIdList(v []*string) *ListClientRuleStrategyResponseBodyList {
	s.IncludeGroupIdList = v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetMainButtonTextCh(v string) *ListClientRuleStrategyResponseBodyList {
	s.MainButtonTextCh = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetMainButtonTextEn(v string) *ListClientRuleStrategyResponseBodyList {
	s.MainButtonTextEn = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetMatchMode(v string) *ListClientRuleStrategyResponseBodyList {
	s.MatchMode = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetMatchTargetIds(v []*string) *ListClientRuleStrategyResponseBodyList {
	s.MatchTargetIds = v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetMinorButtonTextCh(v string) *ListClientRuleStrategyResponseBodyList {
	s.MinorButtonTextCh = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetMinorButtonTextEn(v string) *ListClientRuleStrategyResponseBodyList {
	s.MinorButtonTextEn = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetName(v string) *ListClientRuleStrategyResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetOutSwitch(v bool) *ListClientRuleStrategyResponseBodyList {
	s.OutSwitch = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetPolicyType(v int32) *ListClientRuleStrategyResponseBodyList {
	s.PolicyType = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetPortOnList(v []*string) *ListClientRuleStrategyResponseBodyList {
	s.PortOnList = v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetPriority(v int32) *ListClientRuleStrategyResponseBodyList {
	s.Priority = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetSelectionKey(v string) *ListClientRuleStrategyResponseBodyList {
	s.SelectionKey = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetSoftWareOnList(v []*string) *ListClientRuleStrategyResponseBodyList {
	s.SoftWareOnList = v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetStatus(v int32) *ListClientRuleStrategyResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetStrategySourceType(v string) *ListClientRuleStrategyResponseBodyList {
	s.StrategySourceType = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetSupportsApproval(v int32) *ListClientRuleStrategyResponseBodyList {
	s.SupportsApproval = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetUsContent(v string) *ListClientRuleStrategyResponseBodyList {
	s.UsContent = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetUsTitle(v string) *ListClientRuleStrategyResponseBodyList {
	s.UsTitle = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetZhContent(v string) *ListClientRuleStrategyResponseBodyList {
	s.ZhContent = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) SetZhTitle(v string) *ListClientRuleStrategyResponseBodyList {
	s.ZhTitle = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListClientRuleStrategyResponseBodyPageInfo struct {
	// The page number of the current page for the paged query with paging.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries returned per page for the paged query with paging.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClientRuleStrategyResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListClientRuleStrategyResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListClientRuleStrategyResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClientRuleStrategyResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClientRuleStrategyResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListClientRuleStrategyResponseBodyPageInfo) SetCurrentPage(v int32) *ListClientRuleStrategyResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyPageInfo) SetPageSize(v int32) *ListClientRuleStrategyResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyPageInfo) SetTotalCount(v int32) *ListClientRuleStrategyResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListClientRuleStrategyResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
