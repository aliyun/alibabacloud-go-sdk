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
	// Specifies whether to allow terminal users to submit an exception request for the policy. Valid values:
	//
	// - **true**: Exception requests are allowed. The terminal pop-up window provides an entry for the request.
	//
	// - **false**: Exception requests are not allowed.
	//
	// example:
	//
	// false
	AllowReport *string `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	// The description of the prohibited software policy. The description can be up to 128 characters in length and can be empty.
	//
	// example:
	//
	// project name pass the check
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the policy is enabled. Valid values:
	//
	// - **true**: The policy is enabled, delivered to terminals, and takes effect.
	//
	// - **false**: The policy is disabled. The configuration is retained but not delivered to terminals.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Specifies whether to forcibly terminate running processes of the software. Valid values:
	//
	// - **true**: The terminal immediately terminates the running processes of the software when the policy is matched.
	//
	// - **false**: The terminal does not terminate running processes but only prevents subsequent startups.
	//
	// example:
	//
	// false
	ForceKill *bool `json:"ForceKill,omitempty" xml:"ForceKill,omitempty"`
	// The Chinese text of the primary button in the terminal pop-up window.
	//
	// example:
	//
	// Apply for exemption
	MainButtonTextCh *string `json:"MainButtonTextCh,omitempty" xml:"MainButtonTextCh,omitempty"`
	// The English text of the primary button in the terminal pop-up window.
	//
	// example:
	//
	// Report
	MainButtonTextEn *string `json:"MainButtonTextEn,omitempty" xml:"MainButtonTextEn,omitempty"`
	// The effective scope. Valid values:
	//
	// - **UserGroupAll**: The policy takes effect for all users under the current Alibaba Cloud account. You do not need to specify user groups.
	//
	// - **UserGroupNormal**: The policy takes effect only for users in the user groups specified by UserGroupIds.
	//
	// This parameter is required.
	//
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The Chinese text of the secondary button in the terminal pop-up window.
	//
	// example:
	//
	// Got it
	MinorButtonTextCh *string `json:"MinorButtonTextCh,omitempty" xml:"MinorButtonTextCh,omitempty"`
	// The English text of the secondary button in the terminal pop-up window.
	//
	// example:
	//
	// I know
	MinorButtonTextEn *string `json:"MinorButtonTextEn,omitempty" xml:"MinorButtonTextEn,omitempty"`
	// Policy Name of the prohibited software policy. Policy Name must be 1 to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest_a0344d22
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The object type. Valid values:
	//
	// - **App**: controls disable software. The managed objects are specified by SoftwareIds.
	//
	// - **Tag**: controls disable software labels. The managed objects are specified by TagIds. All disable software within the labels is managed.
	//
	// This parameter is required.
	//
	// example:
	//
	// App
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The action to take. Valid values:
	//
	// - **Ban**: blocks the software from running and displays a pop-up window on the terminal to notify the user.
	//
	// - **BanSilent**: blocks the software from running without notifying the user.
	//
	// - **Warn**: displays a pop-up window on the terminal to notify the user without blocking the software from running.
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
	// The Chinese prompt content in the terminal pop-up window.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	PromptCh *string `json:"PromptCh,omitempty" xml:"PromptCh,omitempty"`
	// The English prompt content in the terminal pop-up window.
	//
	// This parameter is required.
	//
	// example:
	//
	// Auto test prohibition prompt
	PromptEn *string `json:"PromptEn,omitempty" xml:"PromptEn,omitempty"`
	// The collection of prohibited software directly managed by the policy.
	SoftwareIds []*CreateProhibitedPolicyRequestSoftwareIds `json:"SoftwareIds,omitempty" xml:"SoftwareIds,omitempty" type:"Repeated"`
	// The collection of prohibited software tag IDs managed by the policy.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The Chinese title of the terminal pop-up window.
	//
	// example:
	//
	// Prohibited software reminder
	TitleCh *string `json:"TitleCh,omitempty" xml:"TitleCh,omitempty"`
	// The English title of the terminal pop-up window.
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
	// Specifies whether the prohibited software is built-in. Valid values:
	//
	// - **true**: The software is built-in, shared by all Alibaba Cloud accounts, and cannot be modified or deleted.
	//
	// - **false**: The software is custom prohibited software under the current Alibaba Cloud account.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the prohibited software. You can obtain the value from the following sources:
	//
	// - [ListProhibitedSoftware](~~ListProhibitedSoftware~~): queries prohibited software in batches.
	//
	// - [CreateProhibitedSoftware](~~CreateProhibitedSoftware~~): creates custom prohibited software.
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
