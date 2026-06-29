// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskTemplateOptionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultResult(v string) *TaskTemplateOptionConfig
	GetDefaultResult() *string
	SetOptions(v []*QuestionOption) *TaskTemplateOptionConfig
	GetOptions() []*QuestionOption
	SetPreOptions(v []*string) *TaskTemplateOptionConfig
	GetPreOptions() []*string
	SetRule(v string) *TaskTemplateOptionConfig
	GetRule() *string
}

type TaskTemplateOptionConfig struct {
	// The default value must be adapted according to the question type. For a Radio or text box question, directly enter the tag value. For a Multiple Choice question, configure it as ["{tag 1}", "{tag 2}"].
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 猫咪
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// Select the list of question options.
	//
	// if can be null:
	// false
	Options []*QuestionOption `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	// List of preset options for text-type questions.
	//
	// if can be null:
	// false
	PreOptions []*string `json:"PreOptions,omitempty" xml:"PreOptions,omitempty" type:"Repeated"`
	// Validation rule item; valid only for fill-in-the-blank text-type questions.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// \\w+
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
}

func (s TaskTemplateOptionConfig) String() string {
	return dara.Prettify(s)
}

func (s TaskTemplateOptionConfig) GoString() string {
	return s.String()
}

func (s *TaskTemplateOptionConfig) GetDefaultResult() *string {
	return s.DefaultResult
}

func (s *TaskTemplateOptionConfig) GetOptions() []*QuestionOption {
	return s.Options
}

func (s *TaskTemplateOptionConfig) GetPreOptions() []*string {
	return s.PreOptions
}

func (s *TaskTemplateOptionConfig) GetRule() *string {
	return s.Rule
}

func (s *TaskTemplateOptionConfig) SetDefaultResult(v string) *TaskTemplateOptionConfig {
	s.DefaultResult = &v
	return s
}

func (s *TaskTemplateOptionConfig) SetOptions(v []*QuestionOption) *TaskTemplateOptionConfig {
	s.Options = v
	return s
}

func (s *TaskTemplateOptionConfig) SetPreOptions(v []*string) *TaskTemplateOptionConfig {
	s.PreOptions = v
	return s
}

func (s *TaskTemplateOptionConfig) SetRule(v string) *TaskTemplateOptionConfig {
	s.Rule = &v
	return s
}

func (s *TaskTemplateOptionConfig) Validate() error {
	if s.Options != nil {
		for _, item := range s.Options {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
