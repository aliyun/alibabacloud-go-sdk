// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddBuiltInApplicationCustomActionsShrink(v string) *UpdateApplicationTemplateShrinkRequest
	GetAddBuiltInApplicationCustomActionsShrink() *string
	SetApplicationTemplateId(v string) *UpdateApplicationTemplateShrinkRequest
	GetApplicationTemplateId() *string
	SetApplicationTemplateName(v string) *UpdateApplicationTemplateShrinkRequest
	GetApplicationTemplateName() *string
	SetBuiltInApplicationStatus(v string) *UpdateApplicationTemplateShrinkRequest
	GetBuiltInApplicationStatus() *string
	SetCustomRuleStatus(v string) *UpdateApplicationTemplateShrinkRequest
	GetCustomRuleStatus() *string
	SetDescription(v string) *UpdateApplicationTemplateShrinkRequest
	GetDescription() *string
	SetRemoveBuiltInApplicationCustomActionsShrink(v string) *UpdateApplicationTemplateShrinkRequest
	GetRemoveBuiltInApplicationCustomActionsShrink() *string
	SetTemplateStatus(v string) *UpdateApplicationTemplateShrinkRequest
	GetTemplateStatus() *string
	SetUnknownAction(v string) *UpdateApplicationTemplateShrinkRequest
	GetUnknownAction() *string
}

type UpdateApplicationTemplateShrinkRequest struct {
	// The list of custom actions to be added or updated for built-in applications. This list is merged with RemoveBuiltInApplicationCustomActions by removing items first and then adding items. The same application ID cannot appear in both lists at the same time.
	AddBuiltInApplicationCustomActionsShrink *string `json:"AddBuiltInApplicationCustomActions,omitempty" xml:"AddBuiltInApplicationCustomActions,omitempty"`
	// The ID of the application template to be modified. The template must exist within the current Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ati-4f0691c56b8d4f06
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The name of the application template. The name must be 1 to 64 characters in length and can contain Chinese characters, letters, digits, hyphens (-), underscores (_), and periods (.). It cannot contain spaces. If this parameter is not specified or is left empty, the original value is retained.
	//
	// example:
	//
	// Application template 02
	ApplicationTemplateName *string `json:"ApplicationTemplateName,omitempty" xml:"ApplicationTemplateName,omitempty"`
	// The switch status of built-in application control. Valid values: only **Enable*	- is supported. If you specify Disable, an error is returned. If this parameter is not specified or is left empty, the original value is retained.
	//
	// example:
	//
	// Enable
	BuiltInApplicationStatus *string `json:"BuiltInApplicationStatus,omitempty" xml:"BuiltInApplicationStatus,omitempty"`
	// The switch status of custom rules for applications. Valid values:
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// If this parameter is not specified or is left empty, the original value is retained.
	//
	// example:
	//
	// Enable
	CustomRuleStatus *string `json:"CustomRuleStatus,omitempty" xml:"CustomRuleStatus,omitempty"`
	// The description of the application template. The description can be up to 128 characters in length. If this parameter is not specified, the original value is retained. If an empty string is specified, the description is cleared.
	//
	// example:
	//
	// Office network application control
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of built-in applications from which to remove custom actions. After removal, these applications revert to the default actions of the template. The same application ID cannot appear in AddBuiltInApplicationCustomActions at the same time.
	RemoveBuiltInApplicationCustomActionsShrink *string `json:"RemoveBuiltInApplicationCustomActions,omitempty" xml:"RemoveBuiltInApplicationCustomActions,omitempty"`
	// The master switch status of the application template. Valid values:
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// If this parameter is not specified or is left empty, the original value is retained.
	//
	// example:
	//
	// Enable
	TemplateStatus *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The action for unidentified applications. Valid values:
	//
	// - **Allow**: Allow.
	//
	// - **Alert**: Monitor.
	//
	// - **Deny**: Block.
	//
	// If this parameter is not specified or is left empty, the original value is retained.
	//
	// example:
	//
	// Alert
	UnknownAction *string `json:"UnknownAction,omitempty" xml:"UnknownAction,omitempty"`
}

func (s UpdateApplicationTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationTemplateShrinkRequest) GetAddBuiltInApplicationCustomActionsShrink() *string {
	return s.AddBuiltInApplicationCustomActionsShrink
}

func (s *UpdateApplicationTemplateShrinkRequest) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *UpdateApplicationTemplateShrinkRequest) GetApplicationTemplateName() *string {
	return s.ApplicationTemplateName
}

func (s *UpdateApplicationTemplateShrinkRequest) GetBuiltInApplicationStatus() *string {
	return s.BuiltInApplicationStatus
}

func (s *UpdateApplicationTemplateShrinkRequest) GetCustomRuleStatus() *string {
	return s.CustomRuleStatus
}

func (s *UpdateApplicationTemplateShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationTemplateShrinkRequest) GetRemoveBuiltInApplicationCustomActionsShrink() *string {
	return s.RemoveBuiltInApplicationCustomActionsShrink
}

func (s *UpdateApplicationTemplateShrinkRequest) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *UpdateApplicationTemplateShrinkRequest) GetUnknownAction() *string {
	return s.UnknownAction
}

func (s *UpdateApplicationTemplateShrinkRequest) SetAddBuiltInApplicationCustomActionsShrink(v string) *UpdateApplicationTemplateShrinkRequest {
	s.AddBuiltInApplicationCustomActionsShrink = &v
	return s
}

func (s *UpdateApplicationTemplateShrinkRequest) SetApplicationTemplateId(v string) *UpdateApplicationTemplateShrinkRequest {
	s.ApplicationTemplateId = &v
	return s
}

func (s *UpdateApplicationTemplateShrinkRequest) SetApplicationTemplateName(v string) *UpdateApplicationTemplateShrinkRequest {
	s.ApplicationTemplateName = &v
	return s
}

func (s *UpdateApplicationTemplateShrinkRequest) SetBuiltInApplicationStatus(v string) *UpdateApplicationTemplateShrinkRequest {
	s.BuiltInApplicationStatus = &v
	return s
}

func (s *UpdateApplicationTemplateShrinkRequest) SetCustomRuleStatus(v string) *UpdateApplicationTemplateShrinkRequest {
	s.CustomRuleStatus = &v
	return s
}

func (s *UpdateApplicationTemplateShrinkRequest) SetDescription(v string) *UpdateApplicationTemplateShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationTemplateShrinkRequest) SetRemoveBuiltInApplicationCustomActionsShrink(v string) *UpdateApplicationTemplateShrinkRequest {
	s.RemoveBuiltInApplicationCustomActionsShrink = &v
	return s
}

func (s *UpdateApplicationTemplateShrinkRequest) SetTemplateStatus(v string) *UpdateApplicationTemplateShrinkRequest {
	s.TemplateStatus = &v
	return s
}

func (s *UpdateApplicationTemplateShrinkRequest) SetUnknownAction(v string) *UpdateApplicationTemplateShrinkRequest {
	s.UnknownAction = &v
	return s
}

func (s *UpdateApplicationTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
