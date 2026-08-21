// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProhibitedPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *UpdateProhibitedPolicyResponseBodyPolicy) *UpdateProhibitedPolicyResponseBody
	GetPolicy() *UpdateProhibitedPolicyResponseBodyPolicy
	SetRequestId(v string) *UpdateProhibitedPolicyResponseBody
	GetRequestId() *string
}

type UpdateProhibitedPolicyResponseBody struct {
	// The details of the software prohibition policy.
	Policy *UpdateProhibitedPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C4F48DD4-B70D-5342-80B9-2BF5498262FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProhibitedPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedPolicyResponseBody) GetPolicy() *UpdateProhibitedPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *UpdateProhibitedPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProhibitedPolicyResponseBody) SetPolicy(v *UpdateProhibitedPolicyResponseBodyPolicy) *UpdateProhibitedPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *UpdateProhibitedPolicyResponseBody) SetRequestId(v string) *UpdateProhibitedPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProhibitedPolicyResponseBodyPolicy struct {
	// Specifies whether endpoint users are allowed to submit a filing request for this policy. Valid values:
	//
	// - **true**: Filing is allowed. A filing entry is provided in the pop-up notification on the endpoint.
	//
	// - **false**: Filing is not allowed.
	//
	// example:
	//
	// true
	AllowReport *bool `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	// The time when the software prohibition policy was created, in the yyyy-MM-dd HH:mm:ss format. The time is in the UTC+8 time zone.
	//
	// example:
	//
	// 2026-08-19 10:24:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the software prohibition policy.
	//
	// example:
	//
	// OK
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
	// - **true**: The endpoint immediately terminates the running processes of the software when the policy is triggered.
	//
	// - **false**: Running processes are not terminated. Only subsequent launches are blocked.
	//
	// example:
	//
	// false
	ForceKill *bool `json:"ForceKill,omitempty" xml:"ForceKill,omitempty"`
	// The Chinese text of the primary button in the pop-up notification on the endpoint.
	//
	// example:
	//
	// Submit Filing
	MainButtonTextCh *string `json:"MainButtonTextCh,omitempty" xml:"MainButtonTextCh,omitempty"`
	// The English text of the primary button in the pop-up notification on the endpoint.
	//
	// example:
	//
	// Report
	MainButtonTextEn *string `json:"MainButtonTextEn,omitempty" xml:"MainButtonTextEn,omitempty"`
	// The scope in which the policy takes effect. Valid values:
	//
	// - **UserGroupAll**: Takes effect for all users under the current Alibaba Cloud account. No user group needs to be specified.
	//
	// - **UserGroupNormal**: Takes effect only for users in the user groups specified by UserGroupIds.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The Chinese text of the secondary button in the pop-up notification on the endpoint.
	//
	// example:
	//
	// Got It
	MinorButtonTextCh *string `json:"MinorButtonTextCh,omitempty" xml:"MinorButtonTextCh,omitempty"`
	// The English text of the secondary button in the pop-up notification on the endpoint.
	//
	// example:
	//
	// I know
	MinorButtonTextEn *string `json:"MinorButtonTextEn,omitempty" xml:"MinorButtonTextEn,omitempty"`
	// The name of the software prohibition policy.
	//
	// example:
	//
	// autotest_c51af82d
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The object type. Valid values:
	//
	// - **App**: Controls by prohibited software. The controlled objects are specified by SoftwareIds.
	//
	// - **Tag**: Controls by prohibited software tag. The controlled objects are specified by TagIds. All prohibited software under the specified tags is controlled.
	//
	// example:
	//
	// App
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The ID of the software prohibition policy.
	//
	// example:
	//
	// pid-ef8eb37cff62****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The action to take. Valid values:
	//
	// - **Ban**: Blocks the software from running and displays a pop-up notification on the endpoint to alert the user.
	//
	// - **BanSilent**: Blocks the software from running without notifying the user (silent blocking).
	//
	// - **Warn**: Displays a pop-up notification on the endpoint to alert the user without blocking the software from running.
	//
	// example:
	//
	// Ban
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The policy priority. Valid values: 0 to 99. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The Chinese prompt content displayed in the pop-up notification on the endpoint.
	//
	// example:
	//
	// test
	PromptCh *string `json:"PromptCh,omitempty" xml:"PromptCh,omitempty"`
	// The English prompt content displayed in the pop-up notification on the endpoint.
	//
	// example:
	//
	// After analysis and monitoring, the software has been disabled as a security risk.
	PromptEn *string `json:"PromptEn,omitempty" xml:"PromptEn,omitempty"`
	// The collection of prohibited software directly controlled by this policy.
	SoftwareIds []*UpdateProhibitedPolicyResponseBodyPolicySoftwareIds `json:"SoftwareIds,omitempty" xml:"SoftwareIds,omitempty" type:"Repeated"`
	// The collection of prohibited software tag IDs controlled by this policy.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The Chinese title of the pop-up notification on the endpoint.
	//
	// example:
	//
	// Software Prohibition Reminder
	TitleCh *string `json:"TitleCh,omitempty" xml:"TitleCh,omitempty"`
	// The English title of the pop-up notification on the endpoint.
	//
	// example:
	//
	// Software Blocked
	TitleEn *string `json:"TitleEn,omitempty" xml:"TitleEn,omitempty"`
	// The collection of user group IDs for which the policy takes effect.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of exempted usernames.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s UpdateProhibitedPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetAllowReport() *bool {
	return s.AllowReport
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetDescription() *string {
	return s.Description
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetForceKill() *bool {
	return s.ForceKill
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetMainButtonTextCh() *string {
	return s.MainButtonTextCh
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetMainButtonTextEn() *string {
	return s.MainButtonTextEn
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetMinorButtonTextCh() *string {
	return s.MinorButtonTextCh
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetMinorButtonTextEn() *string {
	return s.MinorButtonTextEn
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetName() *string {
	return s.Name
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetObjectType() *string {
	return s.ObjectType
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetPromptCh() *string {
	return s.PromptCh
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetPromptEn() *string {
	return s.PromptEn
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetSoftwareIds() []*UpdateProhibitedPolicyResponseBodyPolicySoftwareIds {
	return s.SoftwareIds
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetTagIds() []*string {
	return s.TagIds
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetTitleCh() *string {
	return s.TitleCh
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetTitleEn() *string {
	return s.TitleEn
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetAllowReport(v bool) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.AllowReport = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetCreateTime(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.CreateTime = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetDescription(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetEnabled(v bool) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.Enabled = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetForceKill(v bool) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.ForceKill = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetMainButtonTextCh(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.MainButtonTextCh = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetMainButtonTextEn(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.MainButtonTextEn = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetMatchMode(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.MatchMode = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetMinorButtonTextCh(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.MinorButtonTextCh = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetMinorButtonTextEn(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.MinorButtonTextEn = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetName(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.Name = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetObjectType(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.ObjectType = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetPolicyId(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetPolicyType(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetPriority(v int32) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.Priority = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetPromptCh(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.PromptCh = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetPromptEn(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.PromptEn = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetSoftwareIds(v []*UpdateProhibitedPolicyResponseBodyPolicySoftwareIds) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.SoftwareIds = v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetTagIds(v []*string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.TagIds = v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetTitleCh(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.TitleCh = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetTitleEn(v string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.TitleEn = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetUserGroupIds(v []*string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.UserGroupIds = v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) SetWhitelist(v []*string) *UpdateProhibitedPolicyResponseBodyPolicy {
	s.Whitelist = v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicy) Validate() error {
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

type UpdateProhibitedPolicyResponseBodyPolicySoftwareIds struct {
	// Indicates whether the prohibited software is a system built-in entry. Valid values:
	//
	// - **true**: A system built-in prohibited software entry shared across all Alibaba Cloud accounts. It cannot be modified or deleted.
	//
	// - **false**: A custom prohibited software entry under the current Alibaba Cloud account.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the prohibited software.
	//
	// example:
	//
	// swb-f024ee962344****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s UpdateProhibitedPolicyResponseBodyPolicySoftwareIds) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedPolicyResponseBodyPolicySoftwareIds) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedPolicyResponseBodyPolicySoftwareIds) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *UpdateProhibitedPolicyResponseBodyPolicySoftwareIds) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *UpdateProhibitedPolicyResponseBodyPolicySoftwareIds) SetIsDefault(v bool) *UpdateProhibitedPolicyResponseBodyPolicySoftwareIds {
	s.IsDefault = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicySoftwareIds) SetSoftwareId(v string) *UpdateProhibitedPolicyResponseBodyPolicySoftwareIds {
	s.SoftwareId = &v
	return s
}

func (s *UpdateProhibitedPolicyResponseBodyPolicySoftwareIds) Validate() error {
	return dara.Validate(s)
}
