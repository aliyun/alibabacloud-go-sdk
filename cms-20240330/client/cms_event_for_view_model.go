// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCmsEventForView interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotations(v map[string]interface{}) *CmsEventForView
	GetAnnotations() map[string]interface{}
	SetData(v interface{}) *CmsEventForView
	GetData() interface{}
	SetDatacontenttype(v string) *CmsEventForView
	GetDatacontenttype() *string
	SetDataschema(v string) *CmsEventForView
	GetDataschema() *string
	SetDedupId(v string) *CmsEventForView
	GetDedupId() *string
	SetId(v string) *CmsEventForView
	GetId() *string
	SetIntegrationUuid(v string) *CmsEventForView
	GetIntegrationUuid() *string
	SetLabels(v map[string]interface{}) *CmsEventForView
	GetLabels() map[string]interface{}
	SetReceiveTime(v int64) *CmsEventForView
	GetReceiveTime() *int64
	SetResource(v *EventResourceForEventView) *CmsEventForView
	GetResource() *EventResourceForEventView
	SetSeverity(v string) *CmsEventForView
	GetSeverity() *string
	SetSource(v string) *CmsEventForView
	GetSource() *string
	SetSourcetype(v string) *CmsEventForView
	GetSourcetype() *string
	SetStatus(v string) *CmsEventForView
	GetStatus() *string
	SetSubject(v string) *CmsEventForView
	GetSubject() *string
	SetSubtype(v string) *CmsEventForView
	GetSubtype() *string
	SetSysId(v string) *CmsEventForView
	GetSysId() *string
	SetTime(v string) *CmsEventForView
	GetTime() *string
	SetTimestamp(v int64) *CmsEventForView
	GetTimestamp() *int64
	SetType(v string) *CmsEventForView
	GetType() *string
	SetWorkspace(v string) *CmsEventForView
	GetWorkspace() *string
	SetWorkspaceTags(v map[string]interface{}) *CmsEventForView
	GetWorkspaceTags() map[string]interface{}
}

type CmsEventForView struct {
	Annotations     map[string]interface{}     `json:"annotations,omitempty" xml:"annotations,omitempty"`
	Data            interface{}                `json:"data,omitempty" xml:"data,omitempty"`
	Datacontenttype *string                    `json:"datacontenttype,omitempty" xml:"datacontenttype,omitempty"`
	Dataschema      *string                    `json:"dataschema,omitempty" xml:"dataschema,omitempty"`
	DedupId         *string                    `json:"dedupId,omitempty" xml:"dedupId,omitempty"`
	Id              *string                    `json:"id,omitempty" xml:"id,omitempty"`
	IntegrationUuid *string                    `json:"integrationUuid,omitempty" xml:"integrationUuid,omitempty"`
	Labels          map[string]interface{}     `json:"labels,omitempty" xml:"labels,omitempty"`
	ReceiveTime     *int64                     `json:"receiveTime,omitempty" xml:"receiveTime,omitempty"`
	Resource        *EventResourceForEventView `json:"resource,omitempty" xml:"resource,omitempty"`
	Severity        *string                    `json:"severity,omitempty" xml:"severity,omitempty"`
	Source          *string                    `json:"source,omitempty" xml:"source,omitempty"`
	Sourcetype      *string                    `json:"sourcetype,omitempty" xml:"sourcetype,omitempty"`
	Status          *string                    `json:"status,omitempty" xml:"status,omitempty"`
	Subject         *string                    `json:"subject,omitempty" xml:"subject,omitempty"`
	Subtype         *string                    `json:"subtype,omitempty" xml:"subtype,omitempty"`
	SysId           *string                    `json:"sysId,omitempty" xml:"sysId,omitempty"`
	Time            *string                    `json:"time,omitempty" xml:"time,omitempty"`
	Timestamp       *int64                     `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Type            *string                    `json:"type,omitempty" xml:"type,omitempty"`
	Workspace       *string                    `json:"workspace,omitempty" xml:"workspace,omitempty"`
	WorkspaceTags   map[string]interface{}     `json:"workspaceTags,omitempty" xml:"workspaceTags,omitempty"`
}

func (s CmsEventForView) String() string {
	return dara.Prettify(s)
}

func (s CmsEventForView) GoString() string {
	return s.String()
}

func (s *CmsEventForView) GetAnnotations() map[string]interface{} {
	return s.Annotations
}

func (s *CmsEventForView) GetData() interface{} {
	return s.Data
}

func (s *CmsEventForView) GetDatacontenttype() *string {
	return s.Datacontenttype
}

func (s *CmsEventForView) GetDataschema() *string {
	return s.Dataschema
}

func (s *CmsEventForView) GetDedupId() *string {
	return s.DedupId
}

func (s *CmsEventForView) GetId() *string {
	return s.Id
}

func (s *CmsEventForView) GetIntegrationUuid() *string {
	return s.IntegrationUuid
}

func (s *CmsEventForView) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *CmsEventForView) GetReceiveTime() *int64 {
	return s.ReceiveTime
}

func (s *CmsEventForView) GetResource() *EventResourceForEventView {
	return s.Resource
}

func (s *CmsEventForView) GetSeverity() *string {
	return s.Severity
}

func (s *CmsEventForView) GetSource() *string {
	return s.Source
}

func (s *CmsEventForView) GetSourcetype() *string {
	return s.Sourcetype
}

func (s *CmsEventForView) GetStatus() *string {
	return s.Status
}

func (s *CmsEventForView) GetSubject() *string {
	return s.Subject
}

func (s *CmsEventForView) GetSubtype() *string {
	return s.Subtype
}

func (s *CmsEventForView) GetSysId() *string {
	return s.SysId
}

func (s *CmsEventForView) GetTime() *string {
	return s.Time
}

func (s *CmsEventForView) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CmsEventForView) GetType() *string {
	return s.Type
}

func (s *CmsEventForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *CmsEventForView) GetWorkspaceTags() map[string]interface{} {
	return s.WorkspaceTags
}

func (s *CmsEventForView) SetAnnotations(v map[string]interface{}) *CmsEventForView {
	s.Annotations = v
	return s
}

func (s *CmsEventForView) SetData(v interface{}) *CmsEventForView {
	s.Data = v
	return s
}

func (s *CmsEventForView) SetDatacontenttype(v string) *CmsEventForView {
	s.Datacontenttype = &v
	return s
}

func (s *CmsEventForView) SetDataschema(v string) *CmsEventForView {
	s.Dataschema = &v
	return s
}

func (s *CmsEventForView) SetDedupId(v string) *CmsEventForView {
	s.DedupId = &v
	return s
}

func (s *CmsEventForView) SetId(v string) *CmsEventForView {
	s.Id = &v
	return s
}

func (s *CmsEventForView) SetIntegrationUuid(v string) *CmsEventForView {
	s.IntegrationUuid = &v
	return s
}

func (s *CmsEventForView) SetLabels(v map[string]interface{}) *CmsEventForView {
	s.Labels = v
	return s
}

func (s *CmsEventForView) SetReceiveTime(v int64) *CmsEventForView {
	s.ReceiveTime = &v
	return s
}

func (s *CmsEventForView) SetResource(v *EventResourceForEventView) *CmsEventForView {
	s.Resource = v
	return s
}

func (s *CmsEventForView) SetSeverity(v string) *CmsEventForView {
	s.Severity = &v
	return s
}

func (s *CmsEventForView) SetSource(v string) *CmsEventForView {
	s.Source = &v
	return s
}

func (s *CmsEventForView) SetSourcetype(v string) *CmsEventForView {
	s.Sourcetype = &v
	return s
}

func (s *CmsEventForView) SetStatus(v string) *CmsEventForView {
	s.Status = &v
	return s
}

func (s *CmsEventForView) SetSubject(v string) *CmsEventForView {
	s.Subject = &v
	return s
}

func (s *CmsEventForView) SetSubtype(v string) *CmsEventForView {
	s.Subtype = &v
	return s
}

func (s *CmsEventForView) SetSysId(v string) *CmsEventForView {
	s.SysId = &v
	return s
}

func (s *CmsEventForView) SetTime(v string) *CmsEventForView {
	s.Time = &v
	return s
}

func (s *CmsEventForView) SetTimestamp(v int64) *CmsEventForView {
	s.Timestamp = &v
	return s
}

func (s *CmsEventForView) SetType(v string) *CmsEventForView {
	s.Type = &v
	return s
}

func (s *CmsEventForView) SetWorkspace(v string) *CmsEventForView {
	s.Workspace = &v
	return s
}

func (s *CmsEventForView) SetWorkspaceTags(v map[string]interface{}) *CmsEventForView {
	s.WorkspaceTags = v
	return s
}

func (s *CmsEventForView) Validate() error {
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}
