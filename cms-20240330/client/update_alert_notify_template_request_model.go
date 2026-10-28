// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertNotifyTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertNotifyTemplateName(v string) *UpdateAlertNotifyTemplateRequest
	GetAlertNotifyTemplateName() *string
	SetTemplates(v map[string]map[string]interface{}) *UpdateAlertNotifyTemplateRequest
	GetTemplates() map[string]map[string]interface{}
}

type UpdateAlertNotifyTemplateRequest struct {
	// The name of the notification template.
	//
	// example:
	//
	// test
	AlertNotifyTemplateName *string `json:"alertNotifyTemplateName,omitempty" xml:"alertNotifyTemplateName,omitempty"`
	// The multilingual notification templates. The key is the language code: zh for Chinese and en for English. The value is the template configuration object for the specified language. When a notification is sent, the corresponding template is selected based on the recipient\\"s language. The body format varies by type: Markdown for DING, WEIXIN, and FEISHU; HTML for EMAIL; plain text for SMS, CALL, and SLACK; and JSON for CUSTOM.
	//
	// example:
	//
	// {"zh":{"title":"云监控报警: {{.incident.title}}","content":"**报警标题:*	- {{.incident.title}}\\n**级别:*	- {{.incident.severity}}\\n**事件数:*	- {{len .events}}","incidentLifeCycleNotifyTemplates":"{\\"IncidentCreationNotify\\":{\\"title\\":\\"云监控报警: {{.incident.title}}\\",\\"content\\":\\"**报警标题:*	- {{.incident.title}}\\\\n**级别:*	- {{.incident.severity}}\\",\\"enable\\":true},\\"IncidentResolutionNotify\\":{\\"title\\":\\"云监控报警已解决: {{.incident.title}}\\",\\"content\\":\\"**报警标题:*	- {{.incident.title}}\\\\n**级别:*	- {{.incident.severity}}\\",\\"enable\\":true}}"},"en":{"title":"CloudMonitor Alert: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}\\n**Events:*	- {{len .events}}","incidentLifeCycleNotifyTemplates":"{\\"IncidentCreationNotify\\":{\\"title\\":\\"CloudMonitor Alert: {{.incident.title}}\\",\\"content\\":\\"**Alert Title:*	- {{.incident.title}}\\\\n**Severity:*	- {{.incident.severity}}\\",\\"enable\\":true},\\"IncidentResolutionNotify\\":{\\"title\\":\\"CloudMonitor Alert Resolved: {{.incident.title}}\\",\\"content\\":\\"**Alert Title:*	- {{.incident.title}}\\\\n**Severity:*	- {{.incident.severity}}\\",\\"enable\\":true}}"}}
	Templates map[string]map[string]interface{} `json:"templates,omitempty" xml:"templates,omitempty"`
}

func (s UpdateAlertNotifyTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertNotifyTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertNotifyTemplateRequest) GetAlertNotifyTemplateName() *string {
	return s.AlertNotifyTemplateName
}

func (s *UpdateAlertNotifyTemplateRequest) GetTemplates() map[string]map[string]interface{} {
	return s.Templates
}

func (s *UpdateAlertNotifyTemplateRequest) SetAlertNotifyTemplateName(v string) *UpdateAlertNotifyTemplateRequest {
	s.AlertNotifyTemplateName = &v
	return s
}

func (s *UpdateAlertNotifyTemplateRequest) SetTemplates(v map[string]map[string]interface{}) *UpdateAlertNotifyTemplateRequest {
	s.Templates = v
	return s
}

func (s *UpdateAlertNotifyTemplateRequest) Validate() error {
	return dara.Validate(s)
}
