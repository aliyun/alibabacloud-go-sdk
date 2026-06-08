// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcessDefinitionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalNodesShrink(v string) *CreateProcessDefinitionShrinkRequest
	GetApprovalNodesShrink() *string
	SetClientToken(v string) *CreateProcessDefinitionShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateProcessDefinitionShrinkRequest
	GetDescription() *string
	SetEnabled(v bool) *CreateProcessDefinitionShrinkRequest
	GetEnabled() *bool
	SetName(v string) *CreateProcessDefinitionShrinkRequest
	GetName() *string
	SetNotificationServicesShrink(v string) *CreateProcessDefinitionShrinkRequest
	GetNotificationServicesShrink() *string
	SetRuleConditionsShrink(v string) *CreateProcessDefinitionShrinkRequest
	GetRuleConditionsShrink() *string
	SetSubType(v string) *CreateProcessDefinitionShrinkRequest
	GetSubType() *string
	SetType(v string) *CreateProcessDefinitionShrinkRequest
	GetType() *string
}

type CreateProcessDefinitionShrinkRequest struct {
	// This parameter is required.
	ApprovalNodesShrink *string `json:"ApprovalNodes,omitempty" xml:"ApprovalNodes,omitempty"`
	// example:
	//
	// 0000-ABCD-EFG****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled     *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	Name                       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NotificationServicesShrink *string `json:"NotificationServices,omitempty" xml:"NotificationServices,omitempty"`
	// This parameter is required.
	RuleConditionsShrink *string `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty"`
	// example:
	//
	// Table
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// example:
	//
	// Extension
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateProcessDefinitionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionShrinkRequest) GetApprovalNodesShrink() *string {
	return s.ApprovalNodesShrink
}

func (s *CreateProcessDefinitionShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProcessDefinitionShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProcessDefinitionShrinkRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateProcessDefinitionShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateProcessDefinitionShrinkRequest) GetNotificationServicesShrink() *string {
	return s.NotificationServicesShrink
}

func (s *CreateProcessDefinitionShrinkRequest) GetRuleConditionsShrink() *string {
	return s.RuleConditionsShrink
}

func (s *CreateProcessDefinitionShrinkRequest) GetSubType() *string {
	return s.SubType
}

func (s *CreateProcessDefinitionShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateProcessDefinitionShrinkRequest) SetApprovalNodesShrink(v string) *CreateProcessDefinitionShrinkRequest {
	s.ApprovalNodesShrink = &v
	return s
}

func (s *CreateProcessDefinitionShrinkRequest) SetClientToken(v string) *CreateProcessDefinitionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProcessDefinitionShrinkRequest) SetDescription(v string) *CreateProcessDefinitionShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProcessDefinitionShrinkRequest) SetEnabled(v bool) *CreateProcessDefinitionShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *CreateProcessDefinitionShrinkRequest) SetName(v string) *CreateProcessDefinitionShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateProcessDefinitionShrinkRequest) SetNotificationServicesShrink(v string) *CreateProcessDefinitionShrinkRequest {
	s.NotificationServicesShrink = &v
	return s
}

func (s *CreateProcessDefinitionShrinkRequest) SetRuleConditionsShrink(v string) *CreateProcessDefinitionShrinkRequest {
	s.RuleConditionsShrink = &v
	return s
}

func (s *CreateProcessDefinitionShrinkRequest) SetSubType(v string) *CreateProcessDefinitionShrinkRequest {
	s.SubType = &v
	return s
}

func (s *CreateProcessDefinitionShrinkRequest) SetType(v string) *CreateProcessDefinitionShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateProcessDefinitionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
