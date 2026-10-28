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
	// This parameter is required.
	//
	// example:
	//
	// test
	AlertNotifyTemplateId *string `json:"alertNotifyTemplateId,omitempty" xml:"alertNotifyTemplateId,omitempty"`
	// example:
	//
	// test
	AlertNotifyTemplateName *string `json:"alertNotifyTemplateName,omitempty" xml:"alertNotifyTemplateName,omitempty"`
	// example:
	//
	// GO
	ProgramLang *string `json:"programLang,omitempty" xml:"programLang,omitempty"`
	// example:
	//
	// {"zh":{"title":"云监控报警: {{.incident.title}}","content":"**报警标题:*	- {{.incident.title}}\\n**级别:*	- {{.incident.severity}}\\n**事件数:*	- {{len .events}}","incidentLifeCycleNotifyTemplates":"{\\"IncidentCreationNotify\\":{\\"title\\":\\"云监控报警: {{.incident.title}}\\",\\"content\\":\\"**报警标题:*	- {{.incident.title}}\\\\n**级别:*	- {{.incident.severity}}\\",\\"enable\\":true},\\"IncidentResolutionNotify\\":{\\"title\\":\\"云监控报警已解决: {{.incident.title}}\\",\\"content\\":\\"**报警标题:*	- {{.incident.title}}\\\\n**级别:*	- {{.incident.severity}}\\",\\"enable\\":true}}"},"en":{"title":"CloudMonitor Alert: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}\\n**Events:*	- {{len .events}}","incidentLifeCycleNotifyTemplates":"{\\"IncidentCreationNotify\\":{\\"title\\":\\"CloudMonitor Alert: {{.incident.title}}\\",\\"content\\":\\"**Alert Title:*	- {{.incident.title}}\\\\n**Severity:*	- {{.incident.severity}}\\",\\"enable\\":true},\\"IncidentResolutionNotify\\":{\\"title\\":\\"CloudMonitor Alert Resolved: {{.incident.title}}\\",\\"content\\":\\"**Alert Title:*	- {{.incident.title}}\\\\n**Severity:*	- {{.incident.severity}}\\",\\"enable\\":true}}"}}
	Templates map[string]map[string]interface{} `json:"templates,omitempty" xml:"templates,omitempty"`
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
