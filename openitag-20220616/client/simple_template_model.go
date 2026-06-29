// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetAbandonReasons(v string) *SimpleTemplate
	GetAbandonReasons() *string
	SetDescription(v string) *SimpleTemplate
	GetDescription() *string
	SetGmtCreateTime(v string) *SimpleTemplate
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *SimpleTemplate
	GetGmtModifiedTime() *string
	SetSharedMode(v string) *SimpleTemplate
	GetSharedMode() *string
	SetStatus(v string) *SimpleTemplate
	GetStatus() *string
	SetTags(v []*string) *SimpleTemplate
	GetTags() []*string
	SetTemplateId(v string) *SimpleTemplate
	GetTemplateId() *string
	SetTemplateName(v string) *SimpleTemplate
	GetTemplateName() *string
	SetTenantId(v string) *SimpleTemplate
	GetTenantId() *string
	SetType(v string) *SimpleTemplate
	GetType() *string
}

type SimpleTemplate struct {
	// Reasons for template deprecation
	//
	// example:
	//
	// ["无效"]
	AbandonReasons *string `json:"AbandonReasons,omitempty" xml:"AbandonReasons,omitempty"`
	// Template description
	//
	// example:
	//
	// 图片分割的模板
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2022-07-12 14:21:08
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Update Time
	//
	// example:
	//
	// 2022-07-12 14:21:08
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Is shared
	//
	// example:
	//
	// ALL
	SharedMode *string `json:"SharedMode,omitempty" xml:"SharedMode,omitempty"`
	// Status
	//
	// example:
	//
	// DRAFT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// List of tags
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Template ID
	//
	// example:
	//
	// 154***1431673270272
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template Name
	//
	// example:
	//
	// 图片分割组合77aa
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Tenant ID of the template
	//
	// example:
	//
	// GA***W134
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Type
	//
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SimpleTemplate) String() string {
	return dara.Prettify(s)
}

func (s SimpleTemplate) GoString() string {
	return s.String()
}

func (s *SimpleTemplate) GetAbandonReasons() *string {
	return s.AbandonReasons
}

func (s *SimpleTemplate) GetDescription() *string {
	return s.Description
}

func (s *SimpleTemplate) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *SimpleTemplate) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *SimpleTemplate) GetSharedMode() *string {
	return s.SharedMode
}

func (s *SimpleTemplate) GetStatus() *string {
	return s.Status
}

func (s *SimpleTemplate) GetTags() []*string {
	return s.Tags
}

func (s *SimpleTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SimpleTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SimpleTemplate) GetTenantId() *string {
	return s.TenantId
}

func (s *SimpleTemplate) GetType() *string {
	return s.Type
}

func (s *SimpleTemplate) SetAbandonReasons(v string) *SimpleTemplate {
	s.AbandonReasons = &v
	return s
}

func (s *SimpleTemplate) SetDescription(v string) *SimpleTemplate {
	s.Description = &v
	return s
}

func (s *SimpleTemplate) SetGmtCreateTime(v string) *SimpleTemplate {
	s.GmtCreateTime = &v
	return s
}

func (s *SimpleTemplate) SetGmtModifiedTime(v string) *SimpleTemplate {
	s.GmtModifiedTime = &v
	return s
}

func (s *SimpleTemplate) SetSharedMode(v string) *SimpleTemplate {
	s.SharedMode = &v
	return s
}

func (s *SimpleTemplate) SetStatus(v string) *SimpleTemplate {
	s.Status = &v
	return s
}

func (s *SimpleTemplate) SetTags(v []*string) *SimpleTemplate {
	s.Tags = v
	return s
}

func (s *SimpleTemplate) SetTemplateId(v string) *SimpleTemplate {
	s.TemplateId = &v
	return s
}

func (s *SimpleTemplate) SetTemplateName(v string) *SimpleTemplate {
	s.TemplateName = &v
	return s
}

func (s *SimpleTemplate) SetTenantId(v string) *SimpleTemplate {
	s.TenantId = &v
	return s
}

func (s *SimpleTemplate) SetType(v string) *SimpleTemplate {
	s.Type = &v
	return s
}

func (s *SimpleTemplate) Validate() error {
	return dara.Validate(s)
}
