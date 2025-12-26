// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *FlowTemplate
	GetAccessibility() *string
	SetAliyunDocumentId(v string) *FlowTemplate
	GetAliyunDocumentId() *string
	SetDescription(v string) *FlowTemplate
	GetDescription() *string
	SetDescriptionI18N(v map[string]*string) *FlowTemplate
	GetDescriptionI18N() map[string]*string
	SetDisplayName(v string) *FlowTemplate
	GetDisplayName() *string
	SetDisplayNameI18N(v map[string]*string) *FlowTemplate
	GetDisplayNameI18N() map[string]*string
	SetFlowFiles(v string) *FlowTemplate
	GetFlowFiles() *string
	SetFlowStoragePath(v string) *FlowTemplate
	GetFlowStoragePath() *string
	SetFlowTemplateId(v string) *FlowTemplate
	GetFlowTemplateId() *string
	SetFlowType(v string) *FlowTemplate
	GetFlowType() *string
	SetLocale(v string) *FlowTemplate
	GetLocale() *string
	SetReferenceCount(v int32) *FlowTemplate
	GetReferenceCount() *int32
	SetTemplateGroup(v string) *FlowTemplate
	GetTemplateGroup() *string
	SetTemplateName(v string) *FlowTemplate
	GetTemplateName() *string
	SetUrl(v string) *FlowTemplate
	GetUrl() *string
	SetVersion(v string) *FlowTemplate
	GetVersion() *string
	SetWorkspaceId(v string) *FlowTemplate
	GetWorkspaceId() *string
}

type FlowTemplate struct {
	Accessibility    *string            `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	AliyunDocumentId *string            `json:"AliyunDocumentId,omitempty" xml:"AliyunDocumentId,omitempty"`
	Description      *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	DescriptionI18N  map[string]*string `json:"DescriptionI18N,omitempty" xml:"DescriptionI18N,omitempty"`
	DisplayName      *string            `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	DisplayNameI18N  map[string]*string `json:"DisplayNameI18N,omitempty" xml:"DisplayNameI18N,omitempty"`
	FlowFiles        *string            `json:"FlowFiles,omitempty" xml:"FlowFiles,omitempty"`
	FlowStoragePath  *string            `json:"FlowStoragePath,omitempty" xml:"FlowStoragePath,omitempty"`
	FlowTemplateId   *string            `json:"FlowTemplateId,omitempty" xml:"FlowTemplateId,omitempty"`
	FlowType         *string            `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	Locale           *string            `json:"Locale,omitempty" xml:"Locale,omitempty"`
	ReferenceCount   *int32             `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	TemplateGroup    *string            `json:"TemplateGroup,omitempty" xml:"TemplateGroup,omitempty"`
	TemplateName     *string            `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	Url              *string            `json:"Url,omitempty" xml:"Url,omitempty"`
	Version          *string            `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkspaceId      *string            `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s FlowTemplate) String() string {
	return dara.Prettify(s)
}

func (s FlowTemplate) GoString() string {
	return s.String()
}

func (s *FlowTemplate) GetAccessibility() *string {
	return s.Accessibility
}

func (s *FlowTemplate) GetAliyunDocumentId() *string {
	return s.AliyunDocumentId
}

func (s *FlowTemplate) GetDescription() *string {
	return s.Description
}

func (s *FlowTemplate) GetDescriptionI18N() map[string]*string {
	return s.DescriptionI18N
}

func (s *FlowTemplate) GetDisplayName() *string {
	return s.DisplayName
}

func (s *FlowTemplate) GetDisplayNameI18N() map[string]*string {
	return s.DisplayNameI18N
}

func (s *FlowTemplate) GetFlowFiles() *string {
	return s.FlowFiles
}

func (s *FlowTemplate) GetFlowStoragePath() *string {
	return s.FlowStoragePath
}

func (s *FlowTemplate) GetFlowTemplateId() *string {
	return s.FlowTemplateId
}

func (s *FlowTemplate) GetFlowType() *string {
	return s.FlowType
}

func (s *FlowTemplate) GetLocale() *string {
	return s.Locale
}

func (s *FlowTemplate) GetReferenceCount() *int32 {
	return s.ReferenceCount
}

func (s *FlowTemplate) GetTemplateGroup() *string {
	return s.TemplateGroup
}

func (s *FlowTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *FlowTemplate) GetUrl() *string {
	return s.Url
}

func (s *FlowTemplate) GetVersion() *string {
	return s.Version
}

func (s *FlowTemplate) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *FlowTemplate) SetAccessibility(v string) *FlowTemplate {
	s.Accessibility = &v
	return s
}

func (s *FlowTemplate) SetAliyunDocumentId(v string) *FlowTemplate {
	s.AliyunDocumentId = &v
	return s
}

func (s *FlowTemplate) SetDescription(v string) *FlowTemplate {
	s.Description = &v
	return s
}

func (s *FlowTemplate) SetDescriptionI18N(v map[string]*string) *FlowTemplate {
	s.DescriptionI18N = v
	return s
}

func (s *FlowTemplate) SetDisplayName(v string) *FlowTemplate {
	s.DisplayName = &v
	return s
}

func (s *FlowTemplate) SetDisplayNameI18N(v map[string]*string) *FlowTemplate {
	s.DisplayNameI18N = v
	return s
}

func (s *FlowTemplate) SetFlowFiles(v string) *FlowTemplate {
	s.FlowFiles = &v
	return s
}

func (s *FlowTemplate) SetFlowStoragePath(v string) *FlowTemplate {
	s.FlowStoragePath = &v
	return s
}

func (s *FlowTemplate) SetFlowTemplateId(v string) *FlowTemplate {
	s.FlowTemplateId = &v
	return s
}

func (s *FlowTemplate) SetFlowType(v string) *FlowTemplate {
	s.FlowType = &v
	return s
}

func (s *FlowTemplate) SetLocale(v string) *FlowTemplate {
	s.Locale = &v
	return s
}

func (s *FlowTemplate) SetReferenceCount(v int32) *FlowTemplate {
	s.ReferenceCount = &v
	return s
}

func (s *FlowTemplate) SetTemplateGroup(v string) *FlowTemplate {
	s.TemplateGroup = &v
	return s
}

func (s *FlowTemplate) SetTemplateName(v string) *FlowTemplate {
	s.TemplateName = &v
	return s
}

func (s *FlowTemplate) SetUrl(v string) *FlowTemplate {
	s.Url = &v
	return s
}

func (s *FlowTemplate) SetVersion(v string) *FlowTemplate {
	s.Version = &v
	return s
}

func (s *FlowTemplate) SetWorkspaceId(v string) *FlowTemplate {
	s.WorkspaceId = &v
	return s
}

func (s *FlowTemplate) Validate() error {
	return dara.Validate(s)
}
