// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddBuiltInApplicationCustomActionsShrink(v string) *CreateApplicationTemplateShrinkRequest
	GetAddBuiltInApplicationCustomActionsShrink() *string
	SetApplicationTemplateName(v string) *CreateApplicationTemplateShrinkRequest
	GetApplicationTemplateName() *string
	SetBuiltInApplicationStatus(v string) *CreateApplicationTemplateShrinkRequest
	GetBuiltInApplicationStatus() *string
	SetCustomRuleStatus(v string) *CreateApplicationTemplateShrinkRequest
	GetCustomRuleStatus() *string
	SetDescription(v string) *CreateApplicationTemplateShrinkRequest
	GetDescription() *string
	SetTemplateStatus(v string) *CreateApplicationTemplateShrinkRequest
	GetTemplateStatus() *string
	SetUnknownAction(v string) *CreateApplicationTemplateShrinkRequest
	GetUnknownAction() *string
}

type CreateApplicationTemplateShrinkRequest struct {
	// The list of custom actions to configure for built-in applications. The application IDs in the list must be unique.
	AddBuiltInApplicationCustomActionsShrink *string `json:"AddBuiltInApplicationCustomActions,omitempty" xml:"AddBuiltInApplicationCustomActions,omitempty"`
	// The name of the application template. The name must be 1 to 64 characters in length and can contain Chinese characters, letters, digits, hyphens (-), underscores (_), and periods (.). It cannot contain spaces or be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// ApplicationTemplate01
	ApplicationTemplateName *string `json:"ApplicationTemplateName,omitempty" xml:"ApplicationTemplateName,omitempty"`
	// The switch status of built-in application control. Only **Enable*	- is supported, which indicates that built-in application control is enabled. If you pass Disable, an error is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	BuiltInApplicationStatus *string `json:"BuiltInApplicationStatus,omitempty" xml:"BuiltInApplicationStatus,omitempty"`
	// The switch status of custom application rules. Valid values:
	//
	// - **Enable**: Enables the feature.
	//
	// - **Disable**: Disables the feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	CustomRuleStatus *string `json:"CustomRuleStatus,omitempty" xml:"CustomRuleStatus,omitempty"`
	// The description of the application template. The description can be up to 128 characters in length and can be empty.
	//
	// example:
	//
	// Office network application control
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The overall switch status of the application template. Valid values:
	//
	// - **Enable**: Enables the template.
	//
	// - **Disable**: Disables the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	TemplateStatus *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The action to take on unidentified applications. Valid values:
	//
	// - **Allow**: Allows the traffic.
	//
	// - **Alert**: Monitors the traffic.
	//
	// - **Deny**: Blocks the traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alert
	UnknownAction *string `json:"UnknownAction,omitempty" xml:"UnknownAction,omitempty"`
}

func (s CreateApplicationTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationTemplateShrinkRequest) GetAddBuiltInApplicationCustomActionsShrink() *string {
	return s.AddBuiltInApplicationCustomActionsShrink
}

func (s *CreateApplicationTemplateShrinkRequest) GetApplicationTemplateName() *string {
	return s.ApplicationTemplateName
}

func (s *CreateApplicationTemplateShrinkRequest) GetBuiltInApplicationStatus() *string {
	return s.BuiltInApplicationStatus
}

func (s *CreateApplicationTemplateShrinkRequest) GetCustomRuleStatus() *string {
	return s.CustomRuleStatus
}

func (s *CreateApplicationTemplateShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationTemplateShrinkRequest) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *CreateApplicationTemplateShrinkRequest) GetUnknownAction() *string {
	return s.UnknownAction
}

func (s *CreateApplicationTemplateShrinkRequest) SetAddBuiltInApplicationCustomActionsShrink(v string) *CreateApplicationTemplateShrinkRequest {
	s.AddBuiltInApplicationCustomActionsShrink = &v
	return s
}

func (s *CreateApplicationTemplateShrinkRequest) SetApplicationTemplateName(v string) *CreateApplicationTemplateShrinkRequest {
	s.ApplicationTemplateName = &v
	return s
}

func (s *CreateApplicationTemplateShrinkRequest) SetBuiltInApplicationStatus(v string) *CreateApplicationTemplateShrinkRequest {
	s.BuiltInApplicationStatus = &v
	return s
}

func (s *CreateApplicationTemplateShrinkRequest) SetCustomRuleStatus(v string) *CreateApplicationTemplateShrinkRequest {
	s.CustomRuleStatus = &v
	return s
}

func (s *CreateApplicationTemplateShrinkRequest) SetDescription(v string) *CreateApplicationTemplateShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationTemplateShrinkRequest) SetTemplateStatus(v string) *CreateApplicationTemplateShrinkRequest {
	s.TemplateStatus = &v
	return s
}

func (s *CreateApplicationTemplateShrinkRequest) SetUnknownAction(v string) *CreateApplicationTemplateShrinkRequest {
	s.UnknownAction = &v
	return s
}

func (s *CreateApplicationTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
