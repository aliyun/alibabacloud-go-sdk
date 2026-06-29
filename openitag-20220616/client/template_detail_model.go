// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAbandonReasons(v []*string) *TemplateDetail
	GetAbandonReasons() []*string
	SetClassify(v string) *TemplateDetail
	GetClassify() *string
	SetCreator(v *SimpleUser) *TemplateDetail
	GetCreator() *SimpleUser
	SetDescription(v string) *TemplateDetail
	GetDescription() *string
	SetExif(v map[string]interface{}) *TemplateDetail
	GetExif() map[string]interface{}
	SetGmtCreateTime(v string) *TemplateDetail
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *TemplateDetail
	GetGmtModifiedTime() *string
	SetModifier(v *SimpleUser) *TemplateDetail
	GetModifier() *SimpleUser
	SetQuestionConfigs(v []*QuestionPlugin) *TemplateDetail
	GetQuestionConfigs() []*QuestionPlugin
	SetSharedMode(v string) *TemplateDetail
	GetSharedMode() *string
	SetStatus(v string) *TemplateDetail
	GetStatus() *string
	SetTags(v []*string) *TemplateDetail
	GetTags() []*string
	SetTemplateId(v string) *TemplateDetail
	GetTemplateId() *string
	SetTemplateName(v string) *TemplateDetail
	GetTemplateName() *string
	SetTenantId(v string) *TemplateDetail
	GetTenantId() *string
	SetType(v string) *TemplateDetail
	GetType() *string
	SetViewConfigs(v *TemplateDetailViewConfigs) *TemplateDetail
	GetViewConfigs() *TemplateDetailViewConfigs
}

type TemplateDetail struct {
	// Reasons for deprecation.
	AbandonReasons []*string `json:"AbandonReasons,omitempty" xml:"AbandonReasons,omitempty" type:"Repeated"`
	// Template categorization.
	//
	// example:
	//
	// picture
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// Creator.
	Creator *SimpleUser `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// Template description.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Additional template information.
	//
	// example:
	//
	// false
	Exif map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// Creation Time.
	//
	// example:
	//
	// 2021-07-07 16:09:20
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Updated At.
	//
	// example:
	//
	// 2021-07-07 16:09:20
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Modifier.
	Modifier *SimpleUser `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// Question widget configuration.
	QuestionConfigs []*QuestionPlugin `json:"QuestionConfigs,omitempty" xml:"QuestionConfigs,omitempty" type:"Repeated"`
	// Template shared mode.
	//
	// example:
	//
	// true
	SharedMode *string `json:"SharedMode,omitempty" xml:"SharedMode,omitempty"`
	// Template status.
	//
	// example:
	//
	// CREATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Label information.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Template ID.
	//
	// example:
	//
	// 1529***48342353920
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template Name.
	//
	// example:
	//
	// demo
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Tenant where the template resides.
	//
	// example:
	//
	// GA***W134
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Template type.
	//
	// example:
	//
	// picture
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// View layer configuration.
	ViewConfigs *TemplateDetailViewConfigs `json:"ViewConfigs,omitempty" xml:"ViewConfigs,omitempty" type:"Struct"`
}

func (s TemplateDetail) String() string {
	return dara.Prettify(s)
}

func (s TemplateDetail) GoString() string {
	return s.String()
}

func (s *TemplateDetail) GetAbandonReasons() []*string {
	return s.AbandonReasons
}

func (s *TemplateDetail) GetClassify() *string {
	return s.Classify
}

func (s *TemplateDetail) GetCreator() *SimpleUser {
	return s.Creator
}

func (s *TemplateDetail) GetDescription() *string {
	return s.Description
}

func (s *TemplateDetail) GetExif() map[string]interface{} {
	return s.Exif
}

func (s *TemplateDetail) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *TemplateDetail) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *TemplateDetail) GetModifier() *SimpleUser {
	return s.Modifier
}

func (s *TemplateDetail) GetQuestionConfigs() []*QuestionPlugin {
	return s.QuestionConfigs
}

func (s *TemplateDetail) GetSharedMode() *string {
	return s.SharedMode
}

func (s *TemplateDetail) GetStatus() *string {
	return s.Status
}

func (s *TemplateDetail) GetTags() []*string {
	return s.Tags
}

func (s *TemplateDetail) GetTemplateId() *string {
	return s.TemplateId
}

func (s *TemplateDetail) GetTemplateName() *string {
	return s.TemplateName
}

func (s *TemplateDetail) GetTenantId() *string {
	return s.TenantId
}

func (s *TemplateDetail) GetType() *string {
	return s.Type
}

func (s *TemplateDetail) GetViewConfigs() *TemplateDetailViewConfigs {
	return s.ViewConfigs
}

func (s *TemplateDetail) SetAbandonReasons(v []*string) *TemplateDetail {
	s.AbandonReasons = v
	return s
}

func (s *TemplateDetail) SetClassify(v string) *TemplateDetail {
	s.Classify = &v
	return s
}

func (s *TemplateDetail) SetCreator(v *SimpleUser) *TemplateDetail {
	s.Creator = v
	return s
}

func (s *TemplateDetail) SetDescription(v string) *TemplateDetail {
	s.Description = &v
	return s
}

func (s *TemplateDetail) SetExif(v map[string]interface{}) *TemplateDetail {
	s.Exif = v
	return s
}

func (s *TemplateDetail) SetGmtCreateTime(v string) *TemplateDetail {
	s.GmtCreateTime = &v
	return s
}

func (s *TemplateDetail) SetGmtModifiedTime(v string) *TemplateDetail {
	s.GmtModifiedTime = &v
	return s
}

func (s *TemplateDetail) SetModifier(v *SimpleUser) *TemplateDetail {
	s.Modifier = v
	return s
}

func (s *TemplateDetail) SetQuestionConfigs(v []*QuestionPlugin) *TemplateDetail {
	s.QuestionConfigs = v
	return s
}

func (s *TemplateDetail) SetSharedMode(v string) *TemplateDetail {
	s.SharedMode = &v
	return s
}

func (s *TemplateDetail) SetStatus(v string) *TemplateDetail {
	s.Status = &v
	return s
}

func (s *TemplateDetail) SetTags(v []*string) *TemplateDetail {
	s.Tags = v
	return s
}

func (s *TemplateDetail) SetTemplateId(v string) *TemplateDetail {
	s.TemplateId = &v
	return s
}

func (s *TemplateDetail) SetTemplateName(v string) *TemplateDetail {
	s.TemplateName = &v
	return s
}

func (s *TemplateDetail) SetTenantId(v string) *TemplateDetail {
	s.TenantId = &v
	return s
}

func (s *TemplateDetail) SetType(v string) *TemplateDetail {
	s.Type = &v
	return s
}

func (s *TemplateDetail) SetViewConfigs(v *TemplateDetailViewConfigs) *TemplateDetail {
	s.ViewConfigs = v
	return s
}

func (s *TemplateDetail) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.Modifier != nil {
		if err := s.Modifier.Validate(); err != nil {
			return err
		}
	}
	if s.QuestionConfigs != nil {
		for _, item := range s.QuestionConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ViewConfigs != nil {
		if err := s.ViewConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TemplateDetailViewConfigs struct {
	// View widgets.
	ViewPlugins []*ViewPlugin `json:"ViewPlugins,omitempty" xml:"ViewPlugins,omitempty" type:"Repeated"`
}

func (s TemplateDetailViewConfigs) String() string {
	return dara.Prettify(s)
}

func (s TemplateDetailViewConfigs) GoString() string {
	return s.String()
}

func (s *TemplateDetailViewConfigs) GetViewPlugins() []*ViewPlugin {
	return s.ViewPlugins
}

func (s *TemplateDetailViewConfigs) SetViewPlugins(v []*ViewPlugin) *TemplateDetailViewConfigs {
	s.ViewPlugins = v
	return s
}

func (s *TemplateDetailViewConfigs) Validate() error {
	if s.ViewPlugins != nil {
		for _, item := range s.ViewPlugins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
