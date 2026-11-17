// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationTemplate(v *GetApplicationTemplateResponseBodyApplicationTemplate) *GetApplicationTemplateResponseBody
	GetApplicationTemplate() *GetApplicationTemplateResponseBodyApplicationTemplate
	SetRequestId(v string) *GetApplicationTemplateResponseBody
	GetRequestId() *string
}

type GetApplicationTemplateResponseBody struct {
	// The application template details.
	ApplicationTemplate *GetApplicationTemplateResponseBodyApplicationTemplate `json:"ApplicationTemplate,omitempty" xml:"ApplicationTemplate,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationTemplateResponseBody) GetApplicationTemplate() *GetApplicationTemplateResponseBodyApplicationTemplate {
	return s.ApplicationTemplate
}

func (s *GetApplicationTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationTemplateResponseBody) SetApplicationTemplate(v *GetApplicationTemplateResponseBodyApplicationTemplate) *GetApplicationTemplateResponseBody {
	s.ApplicationTemplate = v
	return s
}

func (s *GetApplicationTemplateResponseBody) SetRequestId(v string) *GetApplicationTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationTemplateResponseBody) Validate() error {
	if s.ApplicationTemplate != nil {
		if err := s.ApplicationTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationTemplateResponseBodyApplicationTemplate struct {
	// The ID of the application template.
	//
	// example:
	//
	// ati-4f0691c56b8d4f06
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The name of the application template.
	//
	// example:
	//
	// Application Template 01
	ApplicationTemplateName *string `json:"ApplicationTemplateName,omitempty" xml:"ApplicationTemplateName,omitempty"`
	// The list of built-in application actions.
	BuiltInApplicationActions []*GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions `json:"BuiltInApplicationActions,omitempty" xml:"BuiltInApplicationActions,omitempty" type:"Repeated"`
	// The switch status for built-in application control. Valid values:
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// example:
	//
	// Enable
	BuiltInApplicationStatus *string `json:"BuiltInApplicationStatus,omitempty" xml:"BuiltInApplicationStatus,omitempty"`
	// The creation time of the application template. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1731465600
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The switch status for custom application rules. Valid values:
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// example:
	//
	// Enable
	CustomRuleStatus *string `json:"CustomRuleStatus,omitempty" xml:"CustomRuleStatus,omitempty"`
	// The description of the application template.
	//
	// example:
	//
	// Office network application control
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The overall switch status of the application template. Valid values:
	//
	// - **Enable**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// example:
	//
	// Enable
	TemplateStatus *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The action for unidentified applications. Valid values:
	//
	// - **Allow**: Allows the traffic.
	//
	// - **Alert**: Monitors the traffic.
	//
	// - **Deny**: Blocks the traffic.
	//
	// example:
	//
	// Alert
	UnknownAction *string `json:"UnknownAction,omitempty" xml:"UnknownAction,omitempty"`
	// The last update time of the application template. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1731469200
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetApplicationTemplateResponseBodyApplicationTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationTemplateResponseBodyApplicationTemplate) GoString() string {
	return s.String()
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetApplicationTemplateName() *string {
	return s.ApplicationTemplateName
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetBuiltInApplicationActions() []*GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions {
	return s.BuiltInApplicationActions
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetBuiltInApplicationStatus() *string {
	return s.BuiltInApplicationStatus
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetCustomRuleStatus() *string {
	return s.CustomRuleStatus
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetUnknownAction() *string {
	return s.UnknownAction
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetApplicationTemplateId(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.ApplicationTemplateId = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetApplicationTemplateName(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.ApplicationTemplateName = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetBuiltInApplicationActions(v []*GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.BuiltInApplicationActions = v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetBuiltInApplicationStatus(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.BuiltInApplicationStatus = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetCreateTime(v int64) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.CreateTime = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetCustomRuleStatus(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.CustomRuleStatus = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetDescription(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.Description = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetTemplateStatus(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.TemplateStatus = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetUnknownAction(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.UnknownAction = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetUpdateTime(v int64) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.UpdateTime = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) Validate() error {
	if s.BuiltInApplicationActions != nil {
		for _, item := range s.BuiltInApplicationActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions struct {
	// The action performed on the application. Valid values:
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
	// The ID of the application in the built-in application signature library.
	//
	// example:
	//
	// 10012150000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions) GoString() string {
	return s.String()
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions) GetAction() *string {
	return s.Action
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions) SetAction(v string) *GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions {
	s.Action = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions) SetApplicationId(v int64) *GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplateBuiltInApplicationActions) Validate() error {
	return dara.Validate(s)
}
