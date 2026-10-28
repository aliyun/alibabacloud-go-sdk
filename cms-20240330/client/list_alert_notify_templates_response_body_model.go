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
	AlertNotifyTemplates []*ListAlertNotifyTemplatesResponseBodyAlertNotifyTemplates `json:"alertNotifyTemplates,omitempty" xml:"alertNotifyTemplates,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// CD8BA7D6-995D-578D-9941-**********
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	// example:
	//
	// template-abc123
	AlertNotifyTemplateId *string `json:"alertNotifyTemplateId,omitempty" xml:"alertNotifyTemplateId,omitempty"`
	// example:
	//
	// 高可用性告警模板
	AlertNotifyTemplateName *string `json:"alertNotifyTemplateName,omitempty" xml:"alertNotifyTemplateName,omitempty"`
	// example:
	//
	// GO
	ProgramLang *string `json:"programLang,omitempty" xml:"programLang,omitempty"`
	// example:
	//
	// {"zh":{"title":"云监控报警: {{.incident.title}}","content":"**报警标题:*	- {{.incident.title}}\\n**级别:*	- {{.incident.severity}}\\n**事件数:*	- {{len .events}}","incidentLifeCycleNotifyTemplates":{"IncidentCreationNotify":{"title":"云监控报警: {{.incident.title}}","content":"**报警标题:*	- {{.incident.title}}\\n**级别:*	- {{.incident.severity}}","enable":true},"IncidentResolutionNotify":{"title":"云监控报警已解决: {{.incident.title}}","content":"**报警标题:*	- {{.incident.title}}\\n**级别:*	- {{.incident.severity}}","enable":true}}},"en":{"title":"CloudMonitor Alert: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}\\n**Events:*	- {{len .events}}","incidentLifeCycleNotifyTemplates":{"IncidentCreationNotify":{"title":"CloudMonitor Alert: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}","enable":true},"IncidentResolutionNotify":{"title":"CloudMonitor Alert Resolved: {{.incident.title}}","content":"**Alert Title:*	- {{.incident.title}}\\n**Severity:*	- {{.incident.severity}}","enable":true}}}}
	Templates map[string]map[string]interface{} `json:"templates,omitempty" xml:"templates,omitempty"`
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
