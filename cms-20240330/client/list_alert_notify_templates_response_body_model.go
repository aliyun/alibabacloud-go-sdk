// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertNotifyTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertNotifyTemplates(v []*ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) *ListAlertNotifyTemplatesResponseBody
	GetAlertNotifyTemplates() []*ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates
	SetPageNumber(v int64) *ListAlertNotifyTemplatesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAlertNotifyTemplatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListAlertNotifyTemplatesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAlertNotifyTemplatesResponseBody
	GetTotal() *int64
}

type ListAlertNotifyTemplatesResponseBody struct {
	// The list of notification templates.
	AlertNotifyTemplates []*ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates `json:"alertNotifyTemplates,omitempty" xml:"alertNotifyTemplates,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CD8BA7D6-995D-578D-9941-**********
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 33
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAlertNotifyTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotifyTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertNotifyTemplatesResponseBody) GetAlertNotifyTemplates() []*ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates {
	return s.AlertNotifyTemplates
}

func (s *ListAlertNotifyTemplatesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAlertNotifyTemplatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAlertNotifyTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertNotifyTemplatesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAlertNotifyTemplatesResponseBody) SetAlertNotifyTemplates(v []*ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) *ListAlertNotifyTemplatesResponseBody {
	s.AlertNotifyTemplates = v
	return s
}

func (s *ListAlertNotifyTemplatesResponseBody) SetPageNumber(v int64) *ListAlertNotifyTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlertNotifyTemplatesResponseBody) SetPageSize(v int64) *ListAlertNotifyTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlertNotifyTemplatesResponseBody) SetRequestId(v string) *ListAlertNotifyTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertNotifyTemplatesResponseBody) SetTotal(v int64) *ListAlertNotifyTemplatesResponseBody {
	s.Total = &v
	return s
}

func (s *ListAlertNotifyTemplatesResponseBody) Validate() error {
	if s.AlertNotifyTemplates != nil {
		for _, item := range s.AlertNotifyTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates struct {
	// The notification template ID.
	//
	// example:
	//
	// template-abc123
	AlertNotifyTemplateId *string `json:"alertNotifyTemplateId,omitempty" xml:"alertNotifyTemplateId,omitempty"`
	// The notification template name.
	//
	// example:
	//
	// High-availability alert template
	AlertNotifyTemplateName *string `json:"alertNotifyTemplateName,omitempty" xml:"alertNotifyTemplateName,omitempty"`
	// The template syntax type. The current template syntax is fixed to GO and uses the Go template syntax.
	//
	// example:
	//
	// GO
	ProgramLang *string `json:"programLang,omitempty" xml:"programLang,omitempty"`
	// The multilingual notification templates. The key is the language code: zh (Chinese) or en (English). The value is the template configuration object for the language. When a notification is sent, the corresponding template is selected based on the recipient\\"s language. The content format varies by type: Markdown for DING, WEIXIN, and FEISHU; HTML for EMAIL; plain text for SMS, CALL, and SLACK; and JSON for CUSTOM.
	//
	// example:
	//
	// {"zh":{"title":"CloudMonitor Alert: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}\\n**Events:*	- {{len .events}}","incidentLifeCycleNotifyTemplates":{"IncidentCreationNotify":{"title":"CloudMonitor Alert: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}","enable":true},"IncidentResolutionNotify":{"title":"CloudMonitor Alert Resolved: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}","enable":true}}},"en":{"title":"CloudMonitor Alert: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}\\n**Events:*	- {{len .events}}","incidentLifeCycleNotifyTemplates":{"IncidentCreationNotify":{"title":"CloudMonitor Alert: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}","enable":true},"IncidentResolutionNotify":{"title":"CloudMonitor Alert Resolved: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}","enable":true}}}}
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

func (s ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) GoString() string {
	return s.String()
}

func (s *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) GetAlertNotifyTemplateId() *string {
	return s.AlertNotifyTemplateId
}

func (s *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) GetAlertNotifyTemplateName() *string {
	return s.AlertNotifyTemplateName
}

func (s *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) GetProgramLang() *string {
	return s.ProgramLang
}

func (s *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) GetTemplates() map[string]map[string]interface{} {
	return s.Templates
}

func (s *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) GetType() *string {
	return s.Type
}

func (s *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) SetAlertNotifyTemplateId(v string) *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates {
	s.AlertNotifyTemplateId = &v
	return s
}

func (s *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) SetAlertNotifyTemplateName(v string) *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates {
	s.AlertNotifyTemplateName = &v
	return s
}

func (s *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) SetProgramLang(v string) *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates {
	s.ProgramLang = &v
	return s
}

func (s *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) SetTemplates(v map[string]map[string]interface{}) *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates {
	s.Templates = v
	return s
}

func (s *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) SetType(v string) *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates {
	s.Type = &v
	return s
}

func (s *ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates) Validate() error {
	return dara.Validate(s)
}
