// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateQuestion interface {
	dara.Model
	String() string
	GoString() string
	SetChildren(v []*TemplateQuestion) *TemplateQuestion
	GetChildren() []*TemplateQuestion
	SetExif(v map[string]interface{}) *TemplateQuestion
	GetExif() map[string]interface{}
	SetMarkTitle(v string) *TemplateQuestion
	GetMarkTitle() *string
	SetOptions(v []*QuestionOption) *TemplateQuestion
	GetOptions() []*QuestionOption
	SetPreOptions(v []*string) *TemplateQuestion
	GetPreOptions() []*string
	SetQuestionId(v int64) *TemplateQuestion
	GetQuestionId() *int64
	SetType(v string) *TemplateQuestion
	GetType() *string
}

type TemplateQuestion struct {
	// List of child nodes
	Children []*TemplateQuestion `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// Additional properties
	//
	// example:
	//
	// null
	Exif map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// Title
	//
	// example:
	//
	// 题目1
	MarkTitle *string `json:"MarkTitle,omitempty" xml:"MarkTitle,omitempty"`
	// List of options
	Options []*QuestionOption `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	// List of pre-filled values
	PreOptions []*string `json:"PreOptions,omitempty" xml:"PreOptions,omitempty" type:"Repeated"`
	// Question ID
	//
	// example:
	//
	// 1
	QuestionId *int64 `json:"QuestionId,omitempty" xml:"QuestionId,omitempty"`
	// Type, including the following:
	//
	// - TEXT_EDIT
	//
	// - CHECKBOX
	//
	// - INPUT
	//
	// - PICTURE
	//
	// - VIDEO
	//
	// - OPEN_ENDED
	//
	// - SLOT
	//
	// - RADIO
	//
	// example:
	//
	// RADIO
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TemplateQuestion) String() string {
	return dara.Prettify(s)
}

func (s TemplateQuestion) GoString() string {
	return s.String()
}

func (s *TemplateQuestion) GetChildren() []*TemplateQuestion {
	return s.Children
}

func (s *TemplateQuestion) GetExif() map[string]interface{} {
	return s.Exif
}

func (s *TemplateQuestion) GetMarkTitle() *string {
	return s.MarkTitle
}

func (s *TemplateQuestion) GetOptions() []*QuestionOption {
	return s.Options
}

func (s *TemplateQuestion) GetPreOptions() []*string {
	return s.PreOptions
}

func (s *TemplateQuestion) GetQuestionId() *int64 {
	return s.QuestionId
}

func (s *TemplateQuestion) GetType() *string {
	return s.Type
}

func (s *TemplateQuestion) SetChildren(v []*TemplateQuestion) *TemplateQuestion {
	s.Children = v
	return s
}

func (s *TemplateQuestion) SetExif(v map[string]interface{}) *TemplateQuestion {
	s.Exif = v
	return s
}

func (s *TemplateQuestion) SetMarkTitle(v string) *TemplateQuestion {
	s.MarkTitle = &v
	return s
}

func (s *TemplateQuestion) SetOptions(v []*QuestionOption) *TemplateQuestion {
	s.Options = v
	return s
}

func (s *TemplateQuestion) SetPreOptions(v []*string) *TemplateQuestion {
	s.PreOptions = v
	return s
}

func (s *TemplateQuestion) SetQuestionId(v int64) *TemplateQuestion {
	s.QuestionId = &v
	return s
}

func (s *TemplateQuestion) SetType(v string) *TemplateQuestion {
	s.Type = &v
	return s
}

func (s *TemplateQuestion) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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
