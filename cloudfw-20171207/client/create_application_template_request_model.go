// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddBuiltInApplicationCustomActions(v []*CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions) *CreateApplicationTemplateRequest
	GetAddBuiltInApplicationCustomActions() []*CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions
	SetApplicationTemplateName(v string) *CreateApplicationTemplateRequest
	GetApplicationTemplateName() *string
	SetBuiltInApplicationStatus(v string) *CreateApplicationTemplateRequest
	GetBuiltInApplicationStatus() *string
	SetCustomRuleStatus(v string) *CreateApplicationTemplateRequest
	GetCustomRuleStatus() *string
	SetDescription(v string) *CreateApplicationTemplateRequest
	GetDescription() *string
	SetTemplateStatus(v string) *CreateApplicationTemplateRequest
	GetTemplateStatus() *string
	SetUnknownAction(v string) *CreateApplicationTemplateRequest
	GetUnknownAction() *string
}

type CreateApplicationTemplateRequest struct {
	// The list of custom actions to configure for built-in applications. The application IDs in the list must be unique.
	AddBuiltInApplicationCustomActions []*CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions `json:"AddBuiltInApplicationCustomActions,omitempty" xml:"AddBuiltInApplicationCustomActions,omitempty" type:"Repeated"`
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

func (s CreateApplicationTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationTemplateRequest) GetAddBuiltInApplicationCustomActions() []*CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions {
	return s.AddBuiltInApplicationCustomActions
}

func (s *CreateApplicationTemplateRequest) GetApplicationTemplateName() *string {
	return s.ApplicationTemplateName
}

func (s *CreateApplicationTemplateRequest) GetBuiltInApplicationStatus() *string {
	return s.BuiltInApplicationStatus
}

func (s *CreateApplicationTemplateRequest) GetCustomRuleStatus() *string {
	return s.CustomRuleStatus
}

func (s *CreateApplicationTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationTemplateRequest) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *CreateApplicationTemplateRequest) GetUnknownAction() *string {
	return s.UnknownAction
}

func (s *CreateApplicationTemplateRequest) SetAddBuiltInApplicationCustomActions(v []*CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions) *CreateApplicationTemplateRequest {
	s.AddBuiltInApplicationCustomActions = v
	return s
}

func (s *CreateApplicationTemplateRequest) SetApplicationTemplateName(v string) *CreateApplicationTemplateRequest {
	s.ApplicationTemplateName = &v
	return s
}

func (s *CreateApplicationTemplateRequest) SetBuiltInApplicationStatus(v string) *CreateApplicationTemplateRequest {
	s.BuiltInApplicationStatus = &v
	return s
}

func (s *CreateApplicationTemplateRequest) SetCustomRuleStatus(v string) *CreateApplicationTemplateRequest {
	s.CustomRuleStatus = &v
	return s
}

func (s *CreateApplicationTemplateRequest) SetDescription(v string) *CreateApplicationTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationTemplateRequest) SetTemplateStatus(v string) *CreateApplicationTemplateRequest {
	s.TemplateStatus = &v
	return s
}

func (s *CreateApplicationTemplateRequest) SetUnknownAction(v string) *CreateApplicationTemplateRequest {
	s.UnknownAction = &v
	return s
}

func (s *CreateApplicationTemplateRequest) Validate() error {
	if s.AddBuiltInApplicationCustomActions != nil {
		for _, item := range s.AddBuiltInApplicationCustomActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions struct {
	// The action to perform on the application. Valid values:
	//
	// - **Allow**: Allows the traffic.
	//
	// - **Alert**: Monitors the traffic.
	//
	// - **Deny**: Blocks the traffic.
	//
	// example:
	//
	// Deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The ID of the application for which you want to configure a custom action. The application must exist in the built-in application signature library and the ID cannot be 0. You can call the ListApplications operation to query the ID.
	//
	// example:
	//
	// 10012150000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions) GoString() string {
	return s.String()
}

func (s *CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions) GetAction() *string {
	return s.Action
}

func (s *CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions) SetAction(v string) *CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions {
	s.Action = &v
	return s
}

func (s *CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions) SetApplicationId(v int64) *CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationTemplateRequestAddBuiltInApplicationCustomActions) Validate() error {
	return dara.Validate(s)
}
