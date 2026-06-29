// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateDTO interface {
	dara.Model
	String() string
	GoString() string
	SetClassify(v string) *TemplateDTO
	GetClassify() *string
	SetDescription(v string) *TemplateDTO
	GetDescription() *string
	SetExif(v map[string]interface{}) *TemplateDTO
	GetExif() map[string]interface{}
	SetQuestionConfigs(v []*QuestionPlugin) *TemplateDTO
	GetQuestionConfigs() []*QuestionPlugin
	SetRobotConfigs(v []map[string]interface{}) *TemplateDTO
	GetRobotConfigs() []map[string]interface{}
	SetSharedMode(v string) *TemplateDTO
	GetSharedMode() *string
	SetTags(v []*string) *TemplateDTO
	GetTags() []*string
	SetTemplateId(v string) *TemplateDTO
	GetTemplateId() *string
	SetTemplateName(v string) *TemplateDTO
	GetTemplateName() *string
	SetViewConfigs(v *TemplateDTOViewConfigs) *TemplateDTO
	GetViewConfigs() *TemplateDTOViewConfigs
}

type TemplateDTO struct {
	// Template categorization
	//
	// example:
	//
	// picture
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// Template description
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Template additional information
	//
	// example:
	//
	// false
	Exif map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// List of question widget configurations
	//
	// This parameter is required.
	QuestionConfigs []*QuestionPlugin `json:"QuestionConfigs,omitempty" xml:"QuestionConfigs,omitempty" type:"Repeated"`
	// List of assisted annotation configurations
	RobotConfigs []map[string]interface{} `json:"RobotConfigs,omitempty" xml:"RobotConfigs,omitempty" type:"Repeated"`
	// Template shared mode
	//
	// example:
	//
	// true
	SharedMode *string `json:"SharedMode,omitempty" xml:"SharedMode,omitempty"`
	// List of tag information
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Template ID
	//
	// example:
	//
	// 1529***48342353920
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template Name
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// View layer configuration
	//
	// This parameter is required.
	ViewConfigs *TemplateDTOViewConfigs `json:"ViewConfigs,omitempty" xml:"ViewConfigs,omitempty" type:"Struct"`
}

func (s TemplateDTO) String() string {
	return dara.Prettify(s)
}

func (s TemplateDTO) GoString() string {
	return s.String()
}

func (s *TemplateDTO) GetClassify() *string {
	return s.Classify
}

func (s *TemplateDTO) GetDescription() *string {
	return s.Description
}

func (s *TemplateDTO) GetExif() map[string]interface{} {
	return s.Exif
}

func (s *TemplateDTO) GetQuestionConfigs() []*QuestionPlugin {
	return s.QuestionConfigs
}

func (s *TemplateDTO) GetRobotConfigs() []map[string]interface{} {
	return s.RobotConfigs
}

func (s *TemplateDTO) GetSharedMode() *string {
	return s.SharedMode
}

func (s *TemplateDTO) GetTags() []*string {
	return s.Tags
}

func (s *TemplateDTO) GetTemplateId() *string {
	return s.TemplateId
}

func (s *TemplateDTO) GetTemplateName() *string {
	return s.TemplateName
}

func (s *TemplateDTO) GetViewConfigs() *TemplateDTOViewConfigs {
	return s.ViewConfigs
}

func (s *TemplateDTO) SetClassify(v string) *TemplateDTO {
	s.Classify = &v
	return s
}

func (s *TemplateDTO) SetDescription(v string) *TemplateDTO {
	s.Description = &v
	return s
}

func (s *TemplateDTO) SetExif(v map[string]interface{}) *TemplateDTO {
	s.Exif = v
	return s
}

func (s *TemplateDTO) SetQuestionConfigs(v []*QuestionPlugin) *TemplateDTO {
	s.QuestionConfigs = v
	return s
}

func (s *TemplateDTO) SetRobotConfigs(v []map[string]interface{}) *TemplateDTO {
	s.RobotConfigs = v
	return s
}

func (s *TemplateDTO) SetSharedMode(v string) *TemplateDTO {
	s.SharedMode = &v
	return s
}

func (s *TemplateDTO) SetTags(v []*string) *TemplateDTO {
	s.Tags = v
	return s
}

func (s *TemplateDTO) SetTemplateId(v string) *TemplateDTO {
	s.TemplateId = &v
	return s
}

func (s *TemplateDTO) SetTemplateName(v string) *TemplateDTO {
	s.TemplateName = &v
	return s
}

func (s *TemplateDTO) SetViewConfigs(v *TemplateDTOViewConfigs) *TemplateDTO {
	s.ViewConfigs = v
	return s
}

func (s *TemplateDTO) Validate() error {
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

type TemplateDTOViewConfigs struct {
	// List of view widgets
	ViewPlugins []*ViewPlugin `json:"ViewPlugins,omitempty" xml:"ViewPlugins,omitempty" type:"Repeated"`
}

func (s TemplateDTOViewConfigs) String() string {
	return dara.Prettify(s)
}

func (s TemplateDTOViewConfigs) GoString() string {
	return s.String()
}

func (s *TemplateDTOViewConfigs) GetViewPlugins() []*ViewPlugin {
	return s.ViewPlugins
}

func (s *TemplateDTOViewConfigs) SetViewPlugins(v []*ViewPlugin) *TemplateDTOViewConfigs {
	s.ViewPlugins = v
	return s
}

func (s *TemplateDTOViewConfigs) Validate() error {
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
