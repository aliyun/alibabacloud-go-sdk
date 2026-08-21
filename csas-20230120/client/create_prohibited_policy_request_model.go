// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProhibitedPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowReport(v string) *CreateProhibitedPolicyRequest
	GetAllowReport() *string
	SetDescription(v string) *CreateProhibitedPolicyRequest
	GetDescription() *string
	SetEnabled(v bool) *CreateProhibitedPolicyRequest
	GetEnabled() *bool
	SetForceKill(v bool) *CreateProhibitedPolicyRequest
	GetForceKill() *bool
	SetMainButtonTextCh(v string) *CreateProhibitedPolicyRequest
	GetMainButtonTextCh() *string
	SetMainButtonTextEn(v string) *CreateProhibitedPolicyRequest
	GetMainButtonTextEn() *string
	SetMatchMode(v string) *CreateProhibitedPolicyRequest
	GetMatchMode() *string
	SetMinorButtonTextCh(v string) *CreateProhibitedPolicyRequest
	GetMinorButtonTextCh() *string
	SetMinorButtonTextEn(v string) *CreateProhibitedPolicyRequest
	GetMinorButtonTextEn() *string
	SetName(v string) *CreateProhibitedPolicyRequest
	GetName() *string
	SetObjectType(v string) *CreateProhibitedPolicyRequest
	GetObjectType() *string
	SetPolicyType(v string) *CreateProhibitedPolicyRequest
	GetPolicyType() *string
	SetPriority(v int32) *CreateProhibitedPolicyRequest
	GetPriority() *int32
	SetPromptCh(v string) *CreateProhibitedPolicyRequest
	GetPromptCh() *string
	SetPromptEn(v string) *CreateProhibitedPolicyRequest
	GetPromptEn() *string
	SetSoftwareIds(v []*CreateProhibitedPolicyRequestSoftwareIds) *CreateProhibitedPolicyRequest
	GetSoftwareIds() []*CreateProhibitedPolicyRequestSoftwareIds
	SetTagIds(v []*string) *CreateProhibitedPolicyRequest
	GetTagIds() []*string
	SetTitleCh(v string) *CreateProhibitedPolicyRequest
	GetTitleCh() *string
	SetTitleEn(v string) *CreateProhibitedPolicyRequest
	GetTitleEn() *string
	SetUserGroupIds(v []*string) *CreateProhibitedPolicyRequest
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *CreateProhibitedPolicyRequest
	GetWhitelist() []*string
}

type CreateProhibitedPolicyRequest struct {
	// Specifies whether end users are allowed to submit a report request for this policy. Valid values:
	//
	// - **true**: Reporting is allowed. The pop-up window on the endpoint provides a reporting entry.
	//
	// - **false**: Reporting is not allowed.
	//
	// example:
	//
	// false
	AllowReport *string `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	// The description of the software ban policy. The description can be up to 128 characters in length and can be left empty.
	//
	// example:
	//
	// project name pass the check
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the policy is enabled. Valid values:
	//
	// - **true**: Enabled. The policy is delivered to the endpoint and takes effect.
	//
	// - **false**: Disabled. The policy configuration is retained but not delivered to the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Specifies whether to forcefully terminate running software processes. Valid values:
	//
	// example:
	//
	// false
	ForceKill *bool `json:"ForceKill,omitempty" xml:"ForceKill,omitempty"`
	// The Chinese text of the primary button in the pop-up window on the endpoint.
	//
	// example:
	//
	// Submit Filing
	MainButtonTextCh *string `json:"MainButtonTextCh,omitempty" xml:"MainButtonTextCh,omitempty"`
	// The English text of the primary button in the pop-up window on the endpoint.
	//
	// example:
	//
	// Report
	MainButtonTextEn *string `json:"MainButtonTextEn,omitempty" xml:"MainButtonTextEn,omitempty"`
	// The scope in which the policy takes effect. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The Chinese text of the secondary button in the pop-up window on the endpoint.
	//
	// example:
	//
	// Got It
	MinorButtonTextCh *string `json:"MinorButtonTextCh,omitempty" xml:"MinorButtonTextCh,omitempty"`
	// The English text of the secondary button in the pop-up window on the endpoint.
	//
	// example:
	//
	// I know
	MinorButtonTextEn *string `json:"MinorButtonTextEn,omitempty" xml:"MinorButtonTextEn,omitempty"`
	// Policy Name of the software ban policy. Policy Name must be 1 to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest_a0344d22
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The object type. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// App
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The action to take. Valid values:
	//
	// - **Ban**: Blocks the software from running and displays a pop-up notification to the end user.
	//
	// - **BanSilent**: Blocks the software from running without notifying the end user (silent blocking).
	//
	// - **Warn**: Displays a pop-up notification to the end user without blocking the software from running.
	//
	// This parameter is required.
	//
	// example:
	//
	// Warn
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The policy priority. Valid values: 0 to 99. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 99
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The Chinese prompt content displayed in the pop-up window on the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	PromptCh *string `json:"PromptCh,omitempty" xml:"PromptCh,omitempty"`
	// The English prompt content displayed in the pop-up window on the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// Auto test prohibition prompt
	PromptEn *string `json:"PromptEn,omitempty" xml:"PromptEn,omitempty"`
	// The collection of banned software directly controlled by this policy.
	SoftwareIds []*CreateProhibitedPolicyRequestSoftwareIds `json:"SoftwareIds,omitempty" xml:"SoftwareIds,omitempty" type:"Repeated"`
	// The collection of banned software tag IDs controlled by this policy.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The Chinese title of the pop-up window on the endpoint.
	//
	// example:
	//
	// Software Ban Reminder
	TitleCh *string `json:"TitleCh,omitempty" xml:"TitleCh,omitempty"`
	// The English title of the pop-up window on the endpoint.
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

func (s CreateProhibitedPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateProhibitedPolicyRequest) GetAllowReport() *string {
	return s.AllowReport
}

func (s *CreateProhibitedPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProhibitedPolicyRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateProhibitedPolicyRequest) GetForceKill() *bool {
	return s.ForceKill
}

func (s *CreateProhibitedPolicyRequest) GetMainButtonTextCh() *string {
	return s.MainButtonTextCh
}

func (s *CreateProhibitedPolicyRequest) GetMainButtonTextEn() *string {
	return s.MainButtonTextEn
}

func (s *CreateProhibitedPolicyRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateProhibitedPolicyRequest) GetMinorButtonTextCh() *string {
	return s.MinorButtonTextCh
}

func (s *CreateProhibitedPolicyRequest) GetMinorButtonTextEn() *string {
	return s.MinorButtonTextEn
}

func (s *CreateProhibitedPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateProhibitedPolicyRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *CreateProhibitedPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateProhibitedPolicyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateProhibitedPolicyRequest) GetPromptCh() *string {
	return s.PromptCh
}

func (s *CreateProhibitedPolicyRequest) GetPromptEn() *string {
	return s.PromptEn
}

func (s *CreateProhibitedPolicyRequest) GetSoftwareIds() []*CreateProhibitedPolicyRequestSoftwareIds {
	return s.SoftwareIds
}

func (s *CreateProhibitedPolicyRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *CreateProhibitedPolicyRequest) GetTitleCh() *string {
	return s.TitleCh
}

func (s *CreateProhibitedPolicyRequest) GetTitleEn() *string {
	return s.TitleEn
}

func (s *CreateProhibitedPolicyRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateProhibitedPolicyRequest) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *CreateProhibitedPolicyRequest) SetAllowReport(v string) *CreateProhibitedPolicyRequest {
	s.AllowReport = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetDescription(v string) *CreateProhibitedPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetEnabled(v bool) *CreateProhibitedPolicyRequest {
	s.Enabled = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetForceKill(v bool) *CreateProhibitedPolicyRequest {
	s.ForceKill = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetMainButtonTextCh(v string) *CreateProhibitedPolicyRequest {
	s.MainButtonTextCh = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetMainButtonTextEn(v string) *CreateProhibitedPolicyRequest {
	s.MainButtonTextEn = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetMatchMode(v string) *CreateProhibitedPolicyRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetMinorButtonTextCh(v string) *CreateProhibitedPolicyRequest {
	s.MinorButtonTextCh = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetMinorButtonTextEn(v string) *CreateProhibitedPolicyRequest {
	s.MinorButtonTextEn = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetName(v string) *CreateProhibitedPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetObjectType(v string) *CreateProhibitedPolicyRequest {
	s.ObjectType = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetPolicyType(v string) *CreateProhibitedPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetPriority(v int32) *CreateProhibitedPolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetPromptCh(v string) *CreateProhibitedPolicyRequest {
	s.PromptCh = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetPromptEn(v string) *CreateProhibitedPolicyRequest {
	s.PromptEn = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetSoftwareIds(v []*CreateProhibitedPolicyRequestSoftwareIds) *CreateProhibitedPolicyRequest {
	s.SoftwareIds = v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetTagIds(v []*string) *CreateProhibitedPolicyRequest {
	s.TagIds = v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetTitleCh(v string) *CreateProhibitedPolicyRequest {
	s.TitleCh = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetTitleEn(v string) *CreateProhibitedPolicyRequest {
	s.TitleEn = &v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetUserGroupIds(v []*string) *CreateProhibitedPolicyRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateProhibitedPolicyRequest) SetWhitelist(v []*string) *CreateProhibitedPolicyRequest {
	s.Whitelist = v
	return s
}

func (s *CreateProhibitedPolicyRequest) Validate() error {
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

type CreateProhibitedPolicyRequestSoftwareIds struct {
	// Specifies whether the blocked software is a built-in blocked software entry. Valid values:
	//
	// - **true**: A built-in blocked software entry that is shared across all Alibaba Cloud accounts. Built-in entries cannot be modified or deleted.
	//
	// - **false**: A custom blocked software entry under the current Alibaba Cloud account.
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
	// swb-df1fa76d889b****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s CreateProhibitedPolicyRequestSoftwareIds) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedPolicyRequestSoftwareIds) GoString() string {
	return s.String()
}

func (s *CreateProhibitedPolicyRequestSoftwareIds) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *CreateProhibitedPolicyRequestSoftwareIds) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *CreateProhibitedPolicyRequestSoftwareIds) SetIsDefault(v bool) *CreateProhibitedPolicyRequestSoftwareIds {
	s.IsDefault = &v
	return s
}

func (s *CreateProhibitedPolicyRequestSoftwareIds) SetSoftwareId(v string) *CreateProhibitedPolicyRequestSoftwareIds {
	s.SoftwareId = &v
	return s
}

func (s *CreateProhibitedPolicyRequestSoftwareIds) Validate() error {
	return dara.Validate(s)
}
