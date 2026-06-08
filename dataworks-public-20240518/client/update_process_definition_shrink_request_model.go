// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProcessDefinitionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalNodesShrink(v string) *UpdateProcessDefinitionShrinkRequest
	GetApprovalNodesShrink() *string
	SetClientToken(v string) *UpdateProcessDefinitionShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateProcessDefinitionShrinkRequest
	GetDescription() *string
	SetId(v string) *UpdateProcessDefinitionShrinkRequest
	GetId() *string
	SetName(v string) *UpdateProcessDefinitionShrinkRequest
	GetName() *string
	SetNotificationServicesShrink(v string) *UpdateProcessDefinitionShrinkRequest
	GetNotificationServicesShrink() *string
	SetRuleConditionsShrink(v string) *UpdateProcessDefinitionShrinkRequest
	GetRuleConditionsShrink() *string
}

type UpdateProcessDefinitionShrinkRequest struct {
	ApprovalNodesShrink *string `json:"ApprovalNodes,omitempty" xml:"ApprovalNodes,omitempty"`
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 177554881536128
	Id                         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NotificationServicesShrink *string `json:"NotificationServices,omitempty" xml:"NotificationServices,omitempty"`
	RuleConditionsShrink       *string `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty"`
}

func (s UpdateProcessDefinitionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionShrinkRequest) GetApprovalNodesShrink() *string {
	return s.ApprovalNodesShrink
}

func (s *UpdateProcessDefinitionShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateProcessDefinitionShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProcessDefinitionShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateProcessDefinitionShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProcessDefinitionShrinkRequest) GetNotificationServicesShrink() *string {
	return s.NotificationServicesShrink
}

func (s *UpdateProcessDefinitionShrinkRequest) GetRuleConditionsShrink() *string {
	return s.RuleConditionsShrink
}

func (s *UpdateProcessDefinitionShrinkRequest) SetApprovalNodesShrink(v string) *UpdateProcessDefinitionShrinkRequest {
	s.ApprovalNodesShrink = &v
	return s
}

func (s *UpdateProcessDefinitionShrinkRequest) SetClientToken(v string) *UpdateProcessDefinitionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProcessDefinitionShrinkRequest) SetDescription(v string) *UpdateProcessDefinitionShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateProcessDefinitionShrinkRequest) SetId(v string) *UpdateProcessDefinitionShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateProcessDefinitionShrinkRequest) SetName(v string) *UpdateProcessDefinitionShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionShrinkRequest) SetNotificationServicesShrink(v string) *UpdateProcessDefinitionShrinkRequest {
	s.NotificationServicesShrink = &v
	return s
}

func (s *UpdateProcessDefinitionShrinkRequest) SetRuleConditionsShrink(v string) *UpdateProcessDefinitionShrinkRequest {
	s.RuleConditionsShrink = &v
	return s
}

func (s *UpdateProcessDefinitionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
