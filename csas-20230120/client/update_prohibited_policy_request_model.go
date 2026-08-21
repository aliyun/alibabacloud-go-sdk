// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProhibitedPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowReport(v string) *UpdateProhibitedPolicyRequest
	GetAllowReport() *string
	SetDescription(v string) *UpdateProhibitedPolicyRequest
	GetDescription() *string
	SetEnabled(v bool) *UpdateProhibitedPolicyRequest
	GetEnabled() *bool
	SetForceKill(v bool) *UpdateProhibitedPolicyRequest
	GetForceKill() *bool
	SetMainButtonTextCh(v string) *UpdateProhibitedPolicyRequest
	GetMainButtonTextCh() *string
	SetMainButtonTextEn(v string) *UpdateProhibitedPolicyRequest
	GetMainButtonTextEn() *string
	SetMatchMode(v string) *UpdateProhibitedPolicyRequest
	GetMatchMode() *string
	SetMinorButtonTextCh(v string) *UpdateProhibitedPolicyRequest
	GetMinorButtonTextCh() *string
	SetMinorButtonTextEn(v string) *UpdateProhibitedPolicyRequest
	GetMinorButtonTextEn() *string
	SetName(v string) *UpdateProhibitedPolicyRequest
	GetName() *string
	SetObjectType(v string) *UpdateProhibitedPolicyRequest
	GetObjectType() *string
	SetPolicyId(v string) *UpdateProhibitedPolicyRequest
	GetPolicyId() *string
	SetPolicyType(v string) *UpdateProhibitedPolicyRequest
	GetPolicyType() *string
	SetPriority(v int32) *UpdateProhibitedPolicyRequest
	GetPriority() *int32
	SetPromptCh(v string) *UpdateProhibitedPolicyRequest
	GetPromptCh() *string
	SetPromptEn(v string) *UpdateProhibitedPolicyRequest
	GetPromptEn() *string
	SetSoftwareIds(v []*UpdateProhibitedPolicyRequestSoftwareIds) *UpdateProhibitedPolicyRequest
	GetSoftwareIds() []*UpdateProhibitedPolicyRequestSoftwareIds
	SetTagIds(v []*string) *UpdateProhibitedPolicyRequest
	GetTagIds() []*string
	SetTitleCh(v string) *UpdateProhibitedPolicyRequest
	GetTitleCh() *string
	SetTitleEn(v string) *UpdateProhibitedPolicyRequest
	GetTitleEn() *string
	SetUserGroupIds(v []*string) *UpdateProhibitedPolicyRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *UpdateProhibitedPolicyRequest
	GetWhitelist() []*string
}

type UpdateProhibitedPolicyRequest struct {
	// Specifies whether endpoint users are allowed to submit a filing request for this policy. Valid values:
	//
	// - **true**: Filing is allowed. A filing entry is provided in the pop-up notification on the endpoint.
	//
	// - **false**: Filing is not allowed.
	//
	// example:
	//
	// false
	AllowReport *string `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	// The description of the software prohibition policy. The description can contain Chinese characters, uppercase and lowercase letters, digits, spaces, periods (.), underscores (_), and hyphens (-). The description can be up to 128 characters in length and can be left empty.
	//
	// example:
	//
	// No description
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
	// Policy Name of the software prohibition policy. Policy Name must be 1 to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// example:
	//
	// ProhibitionPolicy
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
	// This parameter is required.
	//
	// example:
	//
	// pid-6a9f6adbee0a****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The action to take. If this parameter is not specified, the original value is retained. Valid values:
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
	// 99
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
	// test
	PromptEn *string `json:"PromptEn,omitempty" xml:"PromptEn,omitempty"`
	// The collection of prohibited software directly controlled by this policy.
	SoftwareIds []*UpdateProhibitedPolicyRequestSoftwareIds `json:"SoftwareIds,omitempty" xml:"SoftwareIds,omitempty" type:"Repeated"`
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
	// The collection of user group IDs for which the policy takes effect. Duplicate values are not allowed.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of exempted usernames. Duplicate values are not allowed.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s UpdateProhibitedPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedPolicyRequest) GetAllowReport() *string {
	return s.AllowReport
}

func (s *UpdateProhibitedPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProhibitedPolicyRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateProhibitedPolicyRequest) GetForceKill() *bool {
	return s.ForceKill
}

func (s *UpdateProhibitedPolicyRequest) GetMainButtonTextCh() *string {
	return s.MainButtonTextCh
}

func (s *UpdateProhibitedPolicyRequest) GetMainButtonTextEn() *string {
	return s.MainButtonTextEn
}

func (s *UpdateProhibitedPolicyRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateProhibitedPolicyRequest) GetMinorButtonTextCh() *string {
	return s.MinorButtonTextCh
}

func (s *UpdateProhibitedPolicyRequest) GetMinorButtonTextEn() *string {
	return s.MinorButtonTextEn
}

func (s *UpdateProhibitedPolicyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProhibitedPolicyRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *UpdateProhibitedPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateProhibitedPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateProhibitedPolicyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateProhibitedPolicyRequest) GetPromptCh() *string {
	return s.PromptCh
}

func (s *UpdateProhibitedPolicyRequest) GetPromptEn() *string {
	return s.PromptEn
}

func (s *UpdateProhibitedPolicyRequest) GetSoftwareIds() []*UpdateProhibitedPolicyRequestSoftwareIds {
	return s.SoftwareIds
}

func (s *UpdateProhibitedPolicyRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *UpdateProhibitedPolicyRequest) GetTitleCh() *string {
	return s.TitleCh
}

func (s *UpdateProhibitedPolicyRequest) GetTitleEn() *string {
	return s.TitleEn
}

func (s *UpdateProhibitedPolicyRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateProhibitedPolicyRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *UpdateProhibitedPolicyRequest) SetAllowReport(v string) *UpdateProhibitedPolicyRequest {
	s.AllowReport = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetDescription(v string) *UpdateProhibitedPolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetEnabled(v bool) *UpdateProhibitedPolicyRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetForceKill(v bool) *UpdateProhibitedPolicyRequest {
	s.ForceKill = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetMainButtonTextCh(v string) *UpdateProhibitedPolicyRequest {
	s.MainButtonTextCh = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetMainButtonTextEn(v string) *UpdateProhibitedPolicyRequest {
	s.MainButtonTextEn = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetMatchMode(v string) *UpdateProhibitedPolicyRequest {
	s.MatchMode = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetMinorButtonTextCh(v string) *UpdateProhibitedPolicyRequest {
	s.MinorButtonTextCh = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetMinorButtonTextEn(v string) *UpdateProhibitedPolicyRequest {
	s.MinorButtonTextEn = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetName(v string) *UpdateProhibitedPolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetObjectType(v string) *UpdateProhibitedPolicyRequest {
	s.ObjectType = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetPolicyId(v string) *UpdateProhibitedPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetPolicyType(v string) *UpdateProhibitedPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetPriority(v int32) *UpdateProhibitedPolicyRequest {
	s.Priority = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetPromptCh(v string) *UpdateProhibitedPolicyRequest {
	s.PromptCh = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetPromptEn(v string) *UpdateProhibitedPolicyRequest {
	s.PromptEn = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetSoftwareIds(v []*UpdateProhibitedPolicyRequestSoftwareIds) *UpdateProhibitedPolicyRequest {
	s.SoftwareIds = v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetTagIds(v []*string) *UpdateProhibitedPolicyRequest {
	s.TagIds = v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetTitleCh(v string) *UpdateProhibitedPolicyRequest {
	s.TitleCh = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetTitleEn(v string) *UpdateProhibitedPolicyRequest {
	s.TitleEn = &v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetUserGroupIds(v []*string) *UpdateProhibitedPolicyRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdateProhibitedPolicyRequest) SetWhitelist(v []*string) *UpdateProhibitedPolicyRequest {
	s.Whitelist = v
	return s
}

func (s *UpdateProhibitedPolicyRequest) Validate() error {
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

type UpdateProhibitedPolicyRequestSoftwareIds struct {
	// Indicates whether the prohibited software is a system built-in entry. Valid values:
	//
	// - **true**: A system built-in prohibited software entry shared across all Alibaba Cloud accounts. It cannot be modified or deleted.
	//
	// - **false**: A custom prohibited software entry under the current Alibaba Cloud account.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the prohibited software. You can obtain the value from the following operations:
	//
	// - [ListProhibitedSoftware](~~ListProhibitedSoftware~~): Lists prohibited software.
	//
	// - [CreateProhibitedSoftware](~~CreateProhibitedSoftware~~): Creates custom prohibited software.
	//
	// example:
	//
	// swb-23d749361c41****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s UpdateProhibitedPolicyRequestSoftwareIds) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedPolicyRequestSoftwareIds) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedPolicyRequestSoftwareIds) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *UpdateProhibitedPolicyRequestSoftwareIds) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *UpdateProhibitedPolicyRequestSoftwareIds) SetIsDefault(v bool) *UpdateProhibitedPolicyRequestSoftwareIds {
	s.IsDefault = &v
	return s
}

func (s *UpdateProhibitedPolicyRequestSoftwareIds) SetSoftwareId(v string) *UpdateProhibitedPolicyRequestSoftwareIds {
	s.SoftwareId = &v
	return s
}

func (s *UpdateProhibitedPolicyRequestSoftwareIds) Validate() error {
	return dara.Validate(s)
}
