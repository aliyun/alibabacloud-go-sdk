// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskTemplateConfig interface {
	dara.Model
	String() string
	GoString() string
	SetExif(v map[string]*string) *TaskTemplateConfig
	GetExif() map[string]*string
	SetResourceKey(v string) *TaskTemplateConfig
	GetResourceKey() *string
	SetSelectQuestions(v []*string) *TaskTemplateConfig
	GetSelectQuestions() []*string
	SetTemplateOptionMap(v map[string]*TaskTemplateOptionConfig) *TaskTemplateConfig
	GetTemplateOptionMap() map[string]*TaskTemplateOptionConfig
	SetTemplateRelationId(v string) *TaskTemplateConfig
	GetTemplateRelationId() *string
}

type TaskTemplateConfig struct {
	// Additional information for template configuration.
	Exif map[string]*string `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// Display field corresponding to the View.
	//
	// example:
	//
	// url
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	// List of questions in the template.
	SelectQuestions []*string `json:"SelectQuestions,omitempty" xml:"SelectQuestions,omitempty" type:"Repeated"`
	// Template options configuration.
	TemplateOptionMap map[string]*TaskTemplateOptionConfig `json:"TemplateOptionMap,omitempty" xml:"TemplateOptionMap,omitempty"`
	// Template ID on which this depends.
	//
	// example:
	//
	// 154***2391839854592
	TemplateRelationId *string `json:"TemplateRelationId,omitempty" xml:"TemplateRelationId,omitempty"`
}

func (s TaskTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s TaskTemplateConfig) GoString() string {
	return s.String()
}

func (s *TaskTemplateConfig) GetExif() map[string]*string {
	return s.Exif
}

func (s *TaskTemplateConfig) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *TaskTemplateConfig) GetSelectQuestions() []*string {
	return s.SelectQuestions
}

func (s *TaskTemplateConfig) GetTemplateOptionMap() map[string]*TaskTemplateOptionConfig {
	return s.TemplateOptionMap
}

func (s *TaskTemplateConfig) GetTemplateRelationId() *string {
	return s.TemplateRelationId
}

func (s *TaskTemplateConfig) SetExif(v map[string]*string) *TaskTemplateConfig {
	s.Exif = v
	return s
}

func (s *TaskTemplateConfig) SetResourceKey(v string) *TaskTemplateConfig {
	s.ResourceKey = &v
	return s
}

func (s *TaskTemplateConfig) SetSelectQuestions(v []*string) *TaskTemplateConfig {
	s.SelectQuestions = v
	return s
}

func (s *TaskTemplateConfig) SetTemplateOptionMap(v map[string]*TaskTemplateOptionConfig) *TaskTemplateConfig {
	s.TemplateOptionMap = v
	return s
}

func (s *TaskTemplateConfig) SetTemplateRelationId(v string) *TaskTemplateConfig {
	s.TemplateRelationId = &v
	return s
}

func (s *TaskTemplateConfig) Validate() error {
	return dara.Validate(s)
}
