// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProhibitedPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *CreateProhibitedPolicyResponseBodyPolicy) *CreateProhibitedPolicyResponseBody
	GetPolicy() *CreateProhibitedPolicyResponseBodyPolicy
	SetRequestId(v string) *CreateProhibitedPolicyResponseBody
	GetRequestId() *string
}

type CreateProhibitedPolicyResponseBody struct {
	// The details of the software prohibition policy.
	Policy *CreateProhibitedPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 215060E3-03D2-548D-A014-17941EA3B6C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProhibitedPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProhibitedPolicyResponseBody) GetPolicy() *CreateProhibitedPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *CreateProhibitedPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProhibitedPolicyResponseBody) SetPolicy(v *CreateProhibitedPolicyResponseBodyPolicy) *CreateProhibitedPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *CreateProhibitedPolicyResponseBody) SetRequestId(v string) *CreateProhibitedPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProhibitedPolicyResponseBodyPolicy struct {
	// Specifies whether end users are allowed to submit a report request for this policy. Valid values:
	//
	// - **true**: Reporting is allowed. The pop-up window on the endpoint provides a reporting entry.
	//
	// - **false**: Reporting is not allowed.
	//
	// example:
	//
	// false
	AllowReport *bool `json:"AllowReport,omitempty" xml:"AllowReport,omitempty"`
	// The time when the software prohibition policy was created, in the yyyy-MM-dd HH:mm:ss format. The time is displayed in UTC+8.
	//
	// example:
	//
	// 2021-07-29 11:26:02
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the software prohibition policy.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the policy is enabled. Valid values:
	//
	// - **true**: Enabled. The policy is delivered to the endpoint and takes effect.
	//
	// - **false**: Disabled. The policy retains its configuration but is not delivered to the endpoint.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Specifies whether to forcibly terminate running software processes. Valid values:
	//
	// - **true**: The terminal immediately terminates the running process of the software when the policy is hit.
	//
	// - **false**: Running processes are not terminated. Only subsequent launches are blocked.
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
	// The scope of the policy. Valid values:
	//
	// - **UserGroupAll**: The policy takes effect for all users under the current Alibaba Cloud account. You do not need to specify user groups.
	//
	// - **UserGroupNormal**: The policy takes effect only for users in the user groups specified by UserGroupIds.
	//
	// example:
	//
	// UserGroupNormal
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
	// The Policy Name of the software disable policy.
	//
	// example:
	//
	// autotest_846acf98
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The object type of the controlled target. Valid values:
	//
	// example:
	//
	// App
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The ID of the software prohibition policy.
	//
	// example:
	//
	// pid-dcbfd33cb004****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The action to take. Valid values:
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
	// The Chinese prompt content displayed in the pop-up window on the endpoint.
	//
	// example:
	//
	// test
	PromptCh *string `json:"PromptCh,omitempty" xml:"PromptCh,omitempty"`
	// The English prompt content displayed in the pop-up window on the endpoint.
	//
	// example:
	//
	// L0 auto test prompt
	PromptEn *string `json:"PromptEn,omitempty" xml:"PromptEn,omitempty"`
	// The collection of banned software directly controlled by this policy.
	SoftwareIds []*CreateProhibitedPolicyResponseBodyPolicySoftwareIds `json:"SoftwareIds,omitempty" xml:"SoftwareIds,omitempty" type:"Repeated"`
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
	// The IDs of the user groups to which the policy applies.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of exempted usernames.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s CreateProhibitedPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetAllowReport() *bool {
	return s.AllowReport
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetDescription() *string {
	return s.Description
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetForceKill() *bool {
	return s.ForceKill
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetMainButtonTextCh() *string {
	return s.MainButtonTextCh
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetMainButtonTextEn() *string {
	return s.MainButtonTextEn
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetMinorButtonTextCh() *string {
	return s.MinorButtonTextCh
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetMinorButtonTextEn() *string {
	return s.MinorButtonTextEn
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetName() *string {
	return s.Name
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetObjectType() *string {
	return s.ObjectType
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetPromptCh() *string {
	return s.PromptCh
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetPromptEn() *string {
	return s.PromptEn
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetSoftwareIds() []*CreateProhibitedPolicyResponseBodyPolicySoftwareIds {
	return s.SoftwareIds
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetTagIds() []*string {
	return s.TagIds
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetTitleCh() *string {
	return s.TitleCh
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetTitleEn() *string {
	return s.TitleEn
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetAllowReport(v bool) *CreateProhibitedPolicyResponseBodyPolicy {
	s.AllowReport = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetCreateTime(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.CreateTime = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetDescription(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetEnabled(v bool) *CreateProhibitedPolicyResponseBodyPolicy {
	s.Enabled = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetForceKill(v bool) *CreateProhibitedPolicyResponseBodyPolicy {
	s.ForceKill = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetMainButtonTextCh(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.MainButtonTextCh = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetMainButtonTextEn(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.MainButtonTextEn = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetMatchMode(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.MatchMode = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetMinorButtonTextCh(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.MinorButtonTextCh = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetMinorButtonTextEn(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.MinorButtonTextEn = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetName(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.Name = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetObjectType(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.ObjectType = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetPolicyId(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetPolicyType(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetPriority(v int32) *CreateProhibitedPolicyResponseBodyPolicy {
	s.Priority = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetPromptCh(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.PromptCh = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetPromptEn(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.PromptEn = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetSoftwareIds(v []*CreateProhibitedPolicyResponseBodyPolicySoftwareIds) *CreateProhibitedPolicyResponseBodyPolicy {
	s.SoftwareIds = v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetTagIds(v []*string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.TagIds = v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetTitleCh(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.TitleCh = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetTitleEn(v string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.TitleEn = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetUserGroupIds(v []*string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.UserGroupIds = v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) SetWhitelist(v []*string) *CreateProhibitedPolicyResponseBodyPolicy {
	s.Whitelist = v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicy) Validate() error {
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

type CreateProhibitedPolicyResponseBodyPolicySoftwareIds struct {
	// Indicates whether the banned software is a system built-in banned software. Valid values:
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
	// swb-9a0bfde19662****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
}

func (s CreateProhibitedPolicyResponseBodyPolicySoftwareIds) String() string {
	return dara.Prettify(s)
}

func (s CreateProhibitedPolicyResponseBodyPolicySoftwareIds) GoString() string {
	return s.String()
}

func (s *CreateProhibitedPolicyResponseBodyPolicySoftwareIds) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *CreateProhibitedPolicyResponseBodyPolicySoftwareIds) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *CreateProhibitedPolicyResponseBodyPolicySoftwareIds) SetIsDefault(v bool) *CreateProhibitedPolicyResponseBodyPolicySoftwareIds {
	s.IsDefault = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicySoftwareIds) SetSoftwareId(v string) *CreateProhibitedPolicyResponseBodyPolicySoftwareIds {
	s.SoftwareId = &v
	return s
}

func (s *CreateProhibitedPolicyResponseBodyPolicySoftwareIds) Validate() error {
	return dara.Validate(s)
}
