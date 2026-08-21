// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProhibitedPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *GetProhibitedPolicyResponseBodyPolicy) *GetProhibitedPolicyResponseBody
	GetPolicy() *GetProhibitedPolicyResponseBodyPolicy
	SetRequestId(v string) *GetProhibitedPolicyResponseBody
	GetRequestId() *string
}

type GetProhibitedPolicyResponseBody struct {
	// The details of the software prohibition policy.
	Policy *GetProhibitedPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// E966413B-7538-5332-99B4-C3DA016B9453
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProhibitedPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetProhibitedPolicyResponseBody) GetPolicy() *GetProhibitedPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *GetProhibitedPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProhibitedPolicyResponseBody) SetPolicy(v *GetProhibitedPolicyResponseBodyPolicy) *GetProhibitedPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetProhibitedPolicyResponseBody) SetRequestId(v string) *GetProhibitedPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProhibitedPolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProhibitedPolicyResponseBodyPolicy struct {
	// Indicates whether end users are allowed to submit a filing request for this policy. Valid values:
	//
	// - **true**: Filing is allowed. The terminal pop-up window provides a filing entry.
	//
	// - **false**: Filing is not allowed.
	//
	// example:
	//
	// false
	AllowReport *bool `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	// The creation time of the software prohibition policy, in the yyyy-MM-dd HH:mm:ss format using the UTC+8 time zone.
	//
	// example:
	//
	// 2021-07-29 11:26:02
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the software prohibition policy.
	//
	// example:
	//
	// completed
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the policy is enabled. Valid values:
	//
	// - **true**: Enabled. The policy is delivered to terminals and takes effect.
	//
	// - **false**: Disabled. The policy configuration is retained but not delivered to terminals.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether to forcibly terminate running software processes. Valid values:
	//
	// - **true**: The terminal immediately terminates the running processes of the software when the policy is triggered.
	//
	// - **false**: Running processes are not terminated. Only subsequent launches are blocked.
	//
	// example:
	//
	// false
	ForceKill *bool `json:"ForceKill,omitempty" xml:"ForceKill,omitempty"`
	// The Chinese text of the primary button in the terminal pop-up window.
	//
	// example:
	//
	// 前往报备
	MainButtonTextCh *string `json:"MainButtonTextCh,omitempty" xml:"MainButtonTextCh,omitempty"`
	// The English text of the primary button in the terminal pop-up window.
	//
	// example:
	//
	// Report
	MainButtonTextEn *string `json:"MainButtonTextEn,omitempty" xml:"MainButtonTextEn,omitempty"`
	// The policy matching target type. Valid values:
	//
	// - **UserGroupAll**: Associates with all users.
	//
	// - **UserGroupNormal**: Associates with specific user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The Chinese text of the secondary button in the terminal pop-up window.
	//
	// example:
	//
	// 我知道了
	MinorButtonTextCh *string `json:"MinorButtonTextCh,omitempty" xml:"MinorButtonTextCh,omitempty"`
	// The English text of the secondary button in the terminal pop-up window.
	//
	// example:
	//
	// I know
	MinorButtonTextEn *string `json:"MinorButtonTextEn,omitempty" xml:"MinorButtonTextEn,omitempty"`
	// The name of the software prohibition policy.
	//
	// example:
	//
	// PolicyC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The object type of the controlled target. Valid values:
	//
	// - **App**: Controls by prohibited software. The controlled objects are specified by SoftwareIds.
	//
	// - **Tag**: Controls by prohibited software labels. The controlled objects are specified by TagIds. All prohibited software under the labels is controlled.
	//
	// example:
	//
	// App
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The software prohibition policy ID.
	//
	// example:
	//
	// pid-36ee4a5869f3****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The action type. Valid values:
	//
	// - **Ban**: Blocks the software from running and displays a pop-up notification to the end user.
	//
	// - **BanSilent**: Blocks the software from running without notifying the end user (silent blocking).
	//
	// - **Warn**: Displays a pop-up notification to the end user without blocking the software from running.
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
	// The Chinese prompt content displayed in the terminal pop-up window.
	//
	// example:
	//
	// This software has been blocked by the enterprise security policy. To use it, submit an approval request
	PromptCh *string `json:"PromptCh,omitempty" xml:"PromptCh,omitempty"`
	// The English prompt content displayed in the terminal pop-up window.
	//
	// example:
	//
	// This software is blocked by your enterprise security policy.
	PromptEn *string `json:"PromptEn,omitempty" xml:"PromptEn,omitempty"`
	// The approval process ID bound to this policy. An empty string is returned if no approval process is bound. In this case, filing requests submitted by end users are approved by the IT administrator as a fallback. Approval processes are bound by using [AttachPolicy2ApprovalProcess](~~AttachPolicy2ApprovalProcess~~) and unbound by using [DetachPolicy2ApprovalProcess](~~DetachPolicy2ApprovalProcess~~). You can obtain this value from the following operation:
	//
	// - [ListApprovalProcesses](~~ListApprovalProcesses~~): Lists approval processes.
	//
	// example:
	//
	// approval-process-6c2f8a1b7d3e****
	ReportProcessId *string `json:"ReportProcessId,omitempty" xml:"ReportProcessId,omitempty"`
	// The collection of prohibited software directly controlled by this policy.
	SoftwareIds []*GetProhibitedPolicyResponseBodyPolicySoftwareIds `json:"SoftwareIds,omitempty" xml:"SoftwareIds,omitempty" type:"Repeated"`
	// The collection of prohibited software label IDs controlled by this policy.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The Chinese title of the terminal pop-up window.
	//
	// example:
	//
	// 软件禁用提醒
	TitleCh *string `json:"TitleCh,omitempty" xml:"TitleCh,omitempty"`
	// The English title of the terminal pop-up window.
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

func (s GetProhibitedPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetAllowReport() *bool {
	return s.AllowReport
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetDescription() *string {
	return s.Description
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetForceKill() *bool {
	return s.ForceKill
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetMainButtonTextCh() *string {
	return s.MainButtonTextCh
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetMainButtonTextEn() *string {
	return s.MainButtonTextEn
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetMatchMode() *string {
	return s.MatchMode
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetMinorButtonTextCh() *string {
	return s.MinorButtonTextCh
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetMinorButtonTextEn() *string {
	return s.MinorButtonTextEn
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetName() *string {
	return s.Name
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetPriority() *int32 {
	return s.Priority
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetPromptCh() *string {
	return s.PromptCh
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetPromptEn() *string {
	return s.PromptEn
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetReportProcessId() *string {
	return s.ReportProcessId
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetSoftwareIds() []*GetProhibitedPolicyResponseBodyPolicySoftwareIds {
	return s.SoftwareIds
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetTagIds() []*string {
	return s.TagIds
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetTitleCh() *string {
	return s.TitleCh
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetTitleEn() *string {
	return s.TitleEn
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *GetProhibitedPolicyResponseBodyPolicy) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetAllowReport(v bool) *GetProhibitedPolicyResponseBodyPolicy {
	s.AllowReport = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetCreateTime(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.CreateTime = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetDescription(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetEnabled(v bool) *GetProhibitedPolicyResponseBodyPolicy {
	s.Enabled = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetForceKill(v bool) *GetProhibitedPolicyResponseBodyPolicy {
	s.ForceKill = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetMainButtonTextCh(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.MainButtonTextCh = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetMainButtonTextEn(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.MainButtonTextEn = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetMatchMode(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.MatchMode = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetMinorButtonTextCh(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.MinorButtonTextCh = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetMinorButtonTextEn(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.MinorButtonTextEn = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetName(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.Name = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetObjectType(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.ObjectType = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetPolicyId(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetPolicyType(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetPriority(v int32) *GetProhibitedPolicyResponseBodyPolicy {
	s.Priority = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetPromptCh(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.PromptCh = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetPromptEn(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.PromptEn = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetReportProcessId(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.ReportProcessId = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetSoftwareIds(v []*GetProhibitedPolicyResponseBodyPolicySoftwareIds) *GetProhibitedPolicyResponseBodyPolicy {
	s.SoftwareIds = v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetTagIds(v []*string) *GetProhibitedPolicyResponseBodyPolicy {
	s.TagIds = v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetTitleCh(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.TitleCh = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetTitleEn(v string) *GetProhibitedPolicyResponseBodyPolicy {
	s.TitleEn = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetUserGroupIds(v []*string) *GetProhibitedPolicyResponseBodyPolicy {
	s.UserGroupIds = v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) SetWhitelist(v []*string) *GetProhibitedPolicyResponseBodyPolicy {
	s.Whitelist = v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicy) Validate() error {
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

type GetProhibitedPolicyResponseBodyPolicySoftwareIds struct {
	// Indicates whether the prohibited software is a system built-in entry. Valid values:
	//
	// - **true**: A system built-in prohibited software entry shared by all Alibaba Cloud accounts. Modification and deletion are not supported.
	//
	// - **false**: A custom prohibited software entry under the current Alibaba Cloud account.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The prohibited software ID.
	//
	// example:
	//
	// swb-050216aafaae****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s GetProhibitedPolicyResponseBodyPolicySoftwareIds) String() string {
	return dara.Prettify(s)
}

func (s GetProhibitedPolicyResponseBodyPolicySoftwareIds) GoString() string {
	return s.String()
}

func (s *GetProhibitedPolicyResponseBodyPolicySoftwareIds) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetProhibitedPolicyResponseBodyPolicySoftwareIds) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *GetProhibitedPolicyResponseBodyPolicySoftwareIds) SetIsDefault(v bool) *GetProhibitedPolicyResponseBodyPolicySoftwareIds {
	s.IsDefault = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicySoftwareIds) SetSoftwareId(v string) *GetProhibitedPolicyResponseBodyPolicySoftwareIds {
	s.SoftwareId = &v
	return s
}

func (s *GetProhibitedPolicyResponseBodyPolicySoftwareIds) Validate() error {
	return dara.Validate(s)
}
