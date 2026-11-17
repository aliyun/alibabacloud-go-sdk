// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddBuiltInApplicationCustomActions(v []*UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions) *UpdateApplicationTemplateRequest
	GetAddBuiltInApplicationCustomActions() []*UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions
	SetApplicationTemplateId(v string) *UpdateApplicationTemplateRequest
	GetApplicationTemplateId() *string
	SetApplicationTemplateName(v string) *UpdateApplicationTemplateRequest
	GetApplicationTemplateName() *string
	SetBuiltInApplicationStatus(v string) *UpdateApplicationTemplateRequest
	GetBuiltInApplicationStatus() *string
	SetCustomRuleStatus(v string) *UpdateApplicationTemplateRequest
	GetCustomRuleStatus() *string
	SetDescription(v string) *UpdateApplicationTemplateRequest
	GetDescription() *string
	SetRemoveBuiltInApplicationCustomActions(v []*UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions) *UpdateApplicationTemplateRequest
	GetRemoveBuiltInApplicationCustomActions() []*UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions
	SetTemplateStatus(v string) *UpdateApplicationTemplateRequest
	GetTemplateStatus() *string
	SetUnknownAction(v string) *UpdateApplicationTemplateRequest
	GetUnknownAction() *string
}

type UpdateApplicationTemplateRequest struct {
	// The list of custom actions to be added or updated for built-in applications. This list is merged with RemoveBuiltInApplicationCustomActions by removing items first and then adding items. The same application ID cannot appear in both lists at the same time.
	AddBuiltInApplicationCustomActions []*UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions `json:"AddBuiltInApplicationCustomActions,omitempty" xml:"AddBuiltInApplicationCustomActions,omitempty" type:"Repeated"`
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
	RemoveBuiltInApplicationCustomActions []*UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions `json:"RemoveBuiltInApplicationCustomActions,omitempty" xml:"RemoveBuiltInApplicationCustomActions,omitempty" type:"Repeated"`
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

func (s UpdateApplicationTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationTemplateRequest) GetAddBuiltInApplicationCustomActions() []*UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions {
	return s.AddBuiltInApplicationCustomActions
}

func (s *UpdateApplicationTemplateRequest) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *UpdateApplicationTemplateRequest) GetApplicationTemplateName() *string {
	return s.ApplicationTemplateName
}

func (s *UpdateApplicationTemplateRequest) GetBuiltInApplicationStatus() *string {
	return s.BuiltInApplicationStatus
}

func (s *UpdateApplicationTemplateRequest) GetCustomRuleStatus() *string {
	return s.CustomRuleStatus
}

func (s *UpdateApplicationTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationTemplateRequest) GetRemoveBuiltInApplicationCustomActions() []*UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions {
	return s.RemoveBuiltInApplicationCustomActions
}

func (s *UpdateApplicationTemplateRequest) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *UpdateApplicationTemplateRequest) GetUnknownAction() *string {
	return s.UnknownAction
}

func (s *UpdateApplicationTemplateRequest) SetAddBuiltInApplicationCustomActions(v []*UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions) *UpdateApplicationTemplateRequest {
	s.AddBuiltInApplicationCustomActions = v
	return s
}

func (s *UpdateApplicationTemplateRequest) SetApplicationTemplateId(v string) *UpdateApplicationTemplateRequest {
	s.ApplicationTemplateId = &v
	return s
}

func (s *UpdateApplicationTemplateRequest) SetApplicationTemplateName(v string) *UpdateApplicationTemplateRequest {
	s.ApplicationTemplateName = &v
	return s
}

func (s *UpdateApplicationTemplateRequest) SetBuiltInApplicationStatus(v string) *UpdateApplicationTemplateRequest {
	s.BuiltInApplicationStatus = &v
	return s
}

func (s *UpdateApplicationTemplateRequest) SetCustomRuleStatus(v string) *UpdateApplicationTemplateRequest {
	s.CustomRuleStatus = &v
	return s
}

func (s *UpdateApplicationTemplateRequest) SetDescription(v string) *UpdateApplicationTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationTemplateRequest) SetRemoveBuiltInApplicationCustomActions(v []*UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions) *UpdateApplicationTemplateRequest {
	s.RemoveBuiltInApplicationCustomActions = v
	return s
}

func (s *UpdateApplicationTemplateRequest) SetTemplateStatus(v string) *UpdateApplicationTemplateRequest {
	s.TemplateStatus = &v
	return s
}

func (s *UpdateApplicationTemplateRequest) SetUnknownAction(v string) *UpdateApplicationTemplateRequest {
	s.UnknownAction = &v
	return s
}

func (s *UpdateApplicationTemplateRequest) Validate() error {
	if s.AddBuiltInApplicationCustomActions != nil {
		for _, item := range s.AddBuiltInApplicationCustomActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RemoveBuiltInApplicationCustomActions != nil {
		for _, item := range s.RemoveBuiltInApplicationCustomActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions struct {
	// The action to be performed on the application. Valid values:
	//
	// - **Allow**: Allow.
	//
	// - **Alert**: Monitor.
	//
	// - **Deny**: Block.
	//
	// example:
	//
	// Deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The ID of the application for which to set custom actions. The application must exist in the built-in application signature library and the ID cannot be 0. You can call the ListApplications operation to query the ID.
	//
	// example:
	//
	// 10012150000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions) GoString() string {
	return s.String()
}

func (s *UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions) GetAction() *string {
	return s.Action
}

func (s *UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions) SetAction(v string) *UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions {
	s.Action = &v
	return s
}

func (s *UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions) SetApplicationId(v int64) *UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationTemplateRequestAddBuiltInApplicationCustomActions) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions struct {
	// The ID of the application from which to remove custom actions. The application must exist in the built-in application signature library.
	//
	// example:
	//
	// 10012150000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions) GoString() string {
	return s.String()
}

func (s *UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions) SetApplicationId(v int64) *UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationTemplateRequestRemoveBuiltInApplicationCustomActions) Validate() error {
	return dara.Validate(s)
}
