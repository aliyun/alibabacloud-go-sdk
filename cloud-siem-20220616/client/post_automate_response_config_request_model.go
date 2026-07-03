// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostAutomateResponseConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionConfig(v string) *PostAutomateResponseConfigRequest
	GetActionConfig() *string
	SetActionType(v string) *PostAutomateResponseConfigRequest
	GetActionType() *string
	SetAutoResponseType(v string) *PostAutomateResponseConfigRequest
	GetAutoResponseType() *string
	SetExecutionCondition(v string) *PostAutomateResponseConfigRequest
	GetExecutionCondition() *string
	SetId(v int64) *PostAutomateResponseConfigRequest
	GetId() *int64
	SetRegionId(v string) *PostAutomateResponseConfigRequest
	GetRegionId() *string
	SetRoleFor(v int64) *PostAutomateResponseConfigRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *PostAutomateResponseConfigRequest
	GetRoleType() *int32
	SetRuleName(v string) *PostAutomateResponseConfigRequest
	GetRuleName() *string
	SetSubUserId(v int64) *PostAutomateResponseConfigRequest
	GetSubUserId() *int64
}

type PostAutomateResponseConfigRequest struct {
	// The configuration of the action that is specified in the automated response rule. The value is a JSON array.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "actionType": "doPlaybook",
	//
	//             "playbookName": "WafBlockIP",
	//
	//             "playbookUuid": "bdad6220-6584-41b2-9704-fc6584568758"
	//
	//       }
	//
	// ]
	ActionConfig *string `json:"ActionConfig,omitempty" xml:"ActionConfig,omitempty"`
	// The type of the action. Separate multiple values with commas. Valid values:
	//
	// - **doPlaybook**: runs a playbook
	//
	// - **changeEventStatus**: changes the status of the event
	//
	// - **changeThreatLevel**: changes the threat level of the event
	//
	// example:
	//
	// doPlaybook,changeEventStatus
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The type of the automated response. Valid values:
	//
	// - **event**: event
	//
	// - **alert**: alert
	//
	// example:
	//
	// event
	AutoResponseType *string `json:"AutoResponseType,omitempty" xml:"AutoResponseType,omitempty"`
	// The trigger condition of the automated response rule. The value is in the JSON format.
	//
	// example:
	//
	// [{"left":{"value":"alert_name"},"operator":"containsString","right":{"value":"webshell_online"}}]
	ExecutionCondition *string `json:"ExecutionCondition,omitempty" xml:"ExecutionCondition,omitempty"`
	// The ID of the automated response rule.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The region where the Data Management center of Threat Analysis is located. Select a region based on the location of your assets. Valid values:
	//
	// - **cn-hangzhou**: your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - **ap-southeast-1**: your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user that the administrator uses to switch the view. This parameter is used when an administrator switches to the perspective of a member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts that are managed by the administrator.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The name of the automated response rule.
	//
	// example:
	//
	// cfw kill quara book
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The ID of the user who created the rule.
	//
	// example:
	//
	// 17108579417****
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s PostAutomateResponseConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PostAutomateResponseConfigRequest) GoString() string {
	return s.String()
}

func (s *PostAutomateResponseConfigRequest) GetActionConfig() *string {
	return s.ActionConfig
}

func (s *PostAutomateResponseConfigRequest) GetActionType() *string {
	return s.ActionType
}

func (s *PostAutomateResponseConfigRequest) GetAutoResponseType() *string {
	return s.AutoResponseType
}

func (s *PostAutomateResponseConfigRequest) GetExecutionCondition() *string {
	return s.ExecutionCondition
}

func (s *PostAutomateResponseConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *PostAutomateResponseConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PostAutomateResponseConfigRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *PostAutomateResponseConfigRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *PostAutomateResponseConfigRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PostAutomateResponseConfigRequest) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *PostAutomateResponseConfigRequest) SetActionConfig(v string) *PostAutomateResponseConfigRequest {
	s.ActionConfig = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetActionType(v string) *PostAutomateResponseConfigRequest {
	s.ActionType = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetAutoResponseType(v string) *PostAutomateResponseConfigRequest {
	s.AutoResponseType = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetExecutionCondition(v string) *PostAutomateResponseConfigRequest {
	s.ExecutionCondition = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetId(v int64) *PostAutomateResponseConfigRequest {
	s.Id = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetRegionId(v string) *PostAutomateResponseConfigRequest {
	s.RegionId = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetRoleFor(v int64) *PostAutomateResponseConfigRequest {
	s.RoleFor = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetRoleType(v int32) *PostAutomateResponseConfigRequest {
	s.RoleType = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetRuleName(v string) *PostAutomateResponseConfigRequest {
	s.RuleName = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) SetSubUserId(v int64) *PostAutomateResponseConfigRequest {
	s.SubUserId = &v
	return s
}

func (s *PostAutomateResponseConfigRequest) Validate() error {
	return dara.Validate(s)
}
