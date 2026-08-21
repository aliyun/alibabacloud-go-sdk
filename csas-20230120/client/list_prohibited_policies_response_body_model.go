// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicies(v []*ListProhibitedPoliciesResponseBodyPolicies) *ListProhibitedPoliciesResponseBody
	GetPolicies() []*ListProhibitedPoliciesResponseBodyPolicies
	SetRequestId(v string) *ListProhibitedPoliciesResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *ListProhibitedPoliciesResponseBody
	GetTotalNum() *int64
}

type ListProhibitedPoliciesResponseBody struct {
	// The list of software prohibition policies, sorted by priority from highest to lowest.
	Policies []*ListProhibitedPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 03362EE0-C6F7-51ED-91FF-0BFFA5A2AB67
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of software prohibition policies.
	//
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListProhibitedPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProhibitedPoliciesResponseBody) GetPolicies() []*ListProhibitedPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *ListProhibitedPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProhibitedPoliciesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListProhibitedPoliciesResponseBody) SetPolicies(v []*ListProhibitedPoliciesResponseBodyPolicies) *ListProhibitedPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListProhibitedPoliciesResponseBody) SetRequestId(v string) *ListProhibitedPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBody) SetTotalNum(v int64) *ListProhibitedPoliciesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBody) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProhibitedPoliciesResponseBodyPolicies struct {
	// Indicates whether endpoint users are allowed to submit a filing request for this policy. Valid values:
	//
	// - **true**: Filing is allowed. The endpoint pop-up notification provides a filing entry.
	//
	// - **false**: Filing is not allowed.
	//
	// example:
	//
	// false
	AllowReport *bool `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	// The creation time of the software prohibition policy, in the format of yyyy-MM-dd HH:mm:ss, using the UTC+8 time zone.
	//
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the software prohibition policy.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the policy is enabled. Valid values:
	//
	// - **true**: Enabled. The policy is delivered to endpoints and takes effect.
	//
	// - **false**: Disabled. The policy configuration is retained but not delivered to endpoints.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Specifies whether to forcibly terminate running software processes. Valid values:
	//
	// - **true**: The endpoint immediately terminates the running processes of the software when the policy is matched.
	//
	// - **false**: Running processes are not terminated. Only subsequent launches are blocked.
	//
	// example:
	//
	// false
	ForceKill *bool `json:"ForceKill,omitempty" xml:"ForceKill,omitempty"`
	// The Chinese text of the primary button in the endpoint pop-up notification.
	//
	// example:
	//
	// 去报备
	MainButtonTextCh *string `json:"MainButtonTextCh,omitempty" xml:"MainButtonTextCh,omitempty"`
	// The English text of the primary button in the endpoint pop-up notification.
	//
	// example:
	//
	// Report
	MainButtonTextEn *string `json:"MainButtonTextEn,omitempty" xml:"MainButtonTextEn,omitempty"`
	// The policy matching target type. Valid values:
	//
	// - **UserGroupAll**: Associated with all users.
	//
	// - **UserGroupNormal**: Associated with specific user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The Chinese text of the secondary button in the endpoint pop-up notification.
	//
	// example:
	//
	// 我知道了
	MinorButtonTextCh *string `json:"MinorButtonTextCh,omitempty" xml:"MinorButtonTextCh,omitempty"`
	// The English text of the secondary button in the endpoint pop-up notification.
	//
	// example:
	//
	// Got it
	MinorButtonTextEn *string `json:"MinorButtonTextEn,omitempty" xml:"MinorButtonTextEn,omitempty"`
	// The name of the software prohibition policy.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The object type of the controlled target. Valid values:
	//
	// - **App**: Controls by prohibited software. The controlled objects are specified by SoftwareIds.
	//
	// - **Tag**: Controls by prohibited software tag. The controlled objects are specified by TagIds. All prohibited software under the tag is controlled.
	//
	// example:
	//
	// App
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The ID of the software prohibition policy.
	//
	// example:
	//
	// pid-42f19f1b6a3e****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The action to take. Valid values:
	//
	// - **Ban**: Blocks the software from running and displays a pop-up notification on the endpoint to alert the user.
	//
	// - **BanSilent**: Blocks the software from running without notifying the user. The blocking is silent.
	//
	// - **Warn**: Only displays a pop-up notification on the endpoint to alert the user without blocking the software from running.
	//
	// example:
	//
	// Warn
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The policy priority. Valid values: 0 to 99. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The Chinese prompt content displayed in the endpoint pop-up notification.
	//
	// example:
	//
	// This software has been blocked by the enterprise security policy. To use it, submit an approval request
	PromptCh *string `json:"PromptCh,omitempty" xml:"PromptCh,omitempty"`
	// The English prompt content displayed in the endpoint pop-up notification.
	//
	// example:
	//
	// This software is blocked by your enterprise security policy.
	PromptEn *string `json:"PromptEn,omitempty" xml:"PromptEn,omitempty"`
	// The collection of prohibited software directly controlled by this policy.
	SoftwareIds []*ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds `json:"SoftwareIds,omitempty" xml:"SoftwareIds,omitempty" type:"Repeated"`
	// The collection of prohibited software tag IDs controlled by this policy.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The Chinese title of the endpoint pop-up notification.
	//
	// example:
	//
	// 软件禁用提醒
	TitleCh *string `json:"TitleCh,omitempty" xml:"TitleCh,omitempty"`
	// The English title of the endpoint pop-up notification.
	//
	// example:
	//
	// Software Blocked
	TitleEn *string `json:"TitleEn,omitempty" xml:"TitleEn,omitempty"`
	// The collection of user group IDs to which this policy applies.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of exempted usernames.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s ListProhibitedPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetAllowReport() *bool {
	return s.AllowReport
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetForceKill() *bool {
	return s.ForceKill
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetMainButtonTextCh() *string {
	return s.MainButtonTextCh
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetMainButtonTextEn() *string {
	return s.MainButtonTextEn
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetMinorButtonTextCh() *string {
	return s.MinorButtonTextCh
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetMinorButtonTextEn() *string {
	return s.MinorButtonTextEn
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetName() *string {
	return s.Name
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetObjectType() *string {
	return s.ObjectType
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetPromptCh() *string {
	return s.PromptCh
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetPromptEn() *string {
	return s.PromptEn
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetSoftwareIds() []*ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds {
	return s.SoftwareIds
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetTagIds() []*string {
	return s.TagIds
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetTitleCh() *string {
	return s.TitleCh
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetTitleEn() *string {
	return s.TitleEn
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetAllowReport(v bool) *ListProhibitedPoliciesResponseBodyPolicies {
	s.AllowReport = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetCreateTime(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetDescription(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.Description = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetEnabled(v bool) *ListProhibitedPoliciesResponseBodyPolicies {
	s.Enabled = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetForceKill(v bool) *ListProhibitedPoliciesResponseBodyPolicies {
	s.ForceKill = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetMainButtonTextCh(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.MainButtonTextCh = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetMainButtonTextEn(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.MainButtonTextEn = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetMatchMode(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.MatchMode = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetMinorButtonTextCh(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.MinorButtonTextCh = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetMinorButtonTextEn(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.MinorButtonTextEn = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetName(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.Name = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetObjectType(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.ObjectType = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetPolicyId(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetPolicyType(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.PolicyType = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetPriority(v int32) *ListProhibitedPoliciesResponseBodyPolicies {
	s.Priority = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetPromptCh(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.PromptCh = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetPromptEn(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.PromptEn = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetSoftwareIds(v []*ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds) *ListProhibitedPoliciesResponseBodyPolicies {
	s.SoftwareIds = v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetTagIds(v []*string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.TagIds = v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetTitleCh(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.TitleCh = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetTitleEn(v string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.TitleEn = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetUserGroupIds(v []*string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.UserGroupIds = v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) SetWhitelist(v []*string) *ListProhibitedPoliciesResponseBodyPolicies {
	s.Whitelist = v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPolicies) Validate() error {
	if s.SoftwareIds != nil {
		for _, item := range s.SoftwareIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds struct {
	// Indicates whether the prohibited software is a system built-in entry. Valid values:
	//
	// - **true**: A system built-in prohibited software entry shared by all Alibaba Cloud accounts. Modification and deletion are not supported.
	//
	// - **false**: Custom prohibited software under the current Alibaba Cloud account.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The prohibited software ID.
	//
	// example:
	//
	// swb-9a0bfde19662****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds) GoString() string {
	return s.String()
}

func (s *ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds) SetIsDefault(v bool) *ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds {
	s.IsDefault = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds) SetSoftwareId(v string) *ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds {
	s.SoftwareId = &v
	return s
}

func (s *ListProhibitedPoliciesResponseBodyPoliciesSoftwareIds) Validate() error {
	return dara.Validate(s)
}
