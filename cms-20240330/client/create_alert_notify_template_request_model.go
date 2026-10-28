// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertNotifyTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertNotifyTemplateId(v string) *CreateAlertNotifyTemplateRequest
	GetAlertNotifyTemplateId() *string
	SetAlertNotifyTemplateName(v string) *CreateAlertNotifyTemplateRequest
	GetAlertNotifyTemplateName() *string
	SetProgramLang(v string) *CreateAlertNotifyTemplateRequest
	GetProgramLang() *string
	SetTemplates(v map[string]map[string]interface{}) *CreateAlertNotifyTemplateRequest
	GetTemplates() map[string]map[string]interface{}
	SetType(v string) *CreateAlertNotifyTemplateRequest
	GetType() *string
}

type CreateAlertNotifyTemplateRequest struct {
	// The ID of the notification template.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AlertNotifyTemplateId *string `json:"alertNotifyTemplateId,omitempty" xml:"alertNotifyTemplateId,omitempty"`
	// The name of the notification template.
	//
	// example:
	//
	// test
	AlertNotifyTemplateName *string `json:"alertNotifyTemplateName,omitempty" xml:"alertNotifyTemplateName,omitempty"`
	// The template syntax type. Currently, only Go is supported. Templates in the title and content fields use the Go template syntax.
	//
	// example:
	//
	// GO
	ProgramLang *string `json:"programLang,omitempty" xml:"programLang,omitempty"`
	// The multilingual notification templates. The key is the language code: zh for Chinese and en for English. The value is the template configuration object for the specified language. When a notification is sent, the corresponding template is selected based on the recipient\\"s language. The content format varies by type: Markdown for DING, WEIXIN, and FEISHU; HTML for EMAIL; plain text for SMS, CALL, and SLACK; and JSON for CUSTOM.
	//
	// example:
	//
	// {"zh":{"title":"CloudMonitor Alert: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}\\n**Events:*	- {{len .events}}","incidentLifeCycleNotifyTemplates":"{\\"IncidentCreationNotify\\":{\\"title\\":\\"CloudMonitor Alert: {{.incident.title}}\\",\\"content\\":\\"**Alert Title:*	- {{.incident.title}}\\\\n**Severity:*	- {{.incident.severity}}\\",\\"enable\\":true},\\"IncidentResolutionNotify\\":{\\"title\\":\\"CloudMonitor Alert Resolved: {{.incident.title}}\\",\\"content\\":\\"**Alert Title:*	- {{.incident.title}}\\\\n**Severity:*	- {{.incident.severity}}\\",\\"enable\\":true}}"},"en":{"title":"CloudMonitor Alert: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}\\n**Events:*	- {{len .events}}","incidentLifeCycleNotifyTemplates":"{\\"IncidentCreationNotify\\":{\\"title\\":\\"CloudMonitor Alert: {{.incident.title}}\\",\\"content\\":\\"**Alert Title:*	- {{.incident.title}}\\\\n**Severity:*	- {{.incident.severity}}\\",\\"enable\\":true},\\"IncidentResolutionNotify\\":{\\"title\\":\\"CloudMonitor Alert Resolved: {{.incident.title}}\\",\\"content\\":\\"**Alert Title:*	- {{.incident.title}}\\\\n**Severity:*	- {{.incident.severity}}\\",\\"enable\\":true}}"}}
	Templates map[string]map[string]interface{} `json:"templates,omitempty" xml:"templates,omitempty"`
	// The notification channel type. Valid values:
	//
	// - DING
	//
	// - WEIXIN
	//
	// - FEISHU
	//
	// - SLACK
	//
	// - CUSTOM
	//
	// - TEAMS
	//
	// - SMS
	//
	// - EMAIL
	//
	// - CALL
	//
	// - PAGER_DUTY
	//
	// example:
	//
	// DING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAlertNotifyTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertNotifyTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertNotifyTemplateRequest) GetAlertNotifyTemplateId() *string {
	return s.AlertNotifyTemplateId
}

func (s *CreateAlertNotifyTemplateRequest) GetAlertNotifyTemplateName() *string {
	return s.AlertNotifyTemplateName
}

func (s *CreateAlertNotifyTemplateRequest) GetProgramLang() *string {
	return s.ProgramLang
}

func (s *CreateAlertNotifyTemplateRequest) GetTemplates() map[string]map[string]interface{} {
	return s.Templates
}

func (s *CreateAlertNotifyTemplateRequest) GetType() *string {
	return s.Type
}

func (s *CreateAlertNotifyTemplateRequest) SetAlertNotifyTemplateId(v string) *CreateAlertNotifyTemplateRequest {
	s.AlertNotifyTemplateId = &v
	return s
}

func (s *CreateAlertNotifyTemplateRequest) SetAlertNotifyTemplateName(v string) *CreateAlertNotifyTemplateRequest {
	s.AlertNotifyTemplateName = &v
	return s
}

func (s *CreateAlertNotifyTemplateRequest) SetProgramLang(v string) *CreateAlertNotifyTemplateRequest {
	s.ProgramLang = &v
	return s
}

func (s *CreateAlertNotifyTemplateRequest) SetTemplates(v map[string]map[string]interface{}) *CreateAlertNotifyTemplateRequest {
	s.Templates = v
	return s
}

func (s *CreateAlertNotifyTemplateRequest) SetType(v string) *CreateAlertNotifyTemplateRequest {
	s.Type = &v
	return s
}

func (s *CreateAlertNotifyTemplateRequest) Validate() error {
	return dara.Validate(s)
}
