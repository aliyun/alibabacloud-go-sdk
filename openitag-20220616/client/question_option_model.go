// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuestionOption interface {
	dara.Model
	String() string
	GoString() string
	SetChildren(v []*QuestionOption) *QuestionOption
	GetChildren() []*QuestionOption
	SetColor(v string) *QuestionOption
	GetColor() *string
	SetKey(v string) *QuestionOption
	GetKey() *string
	SetLabel(v string) *QuestionOption
	GetLabel() *string
	SetRemark(v string) *QuestionOption
	GetRemark() *string
	SetShortcut(v string) *QuestionOption
	GetShortcut() *string
}

type QuestionOption struct {
	// List of child options.
	Children []*QuestionOption `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// Color.
	//
	// example:
	//
	// #239125
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// Tag Name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Label display name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dog
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Remark.
	//
	// example:
	//
	// 第一道题目
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Keyboard shortcut.
	//
	// example:
	//
	// 1
	Shortcut *string `json:"Shortcut,omitempty" xml:"Shortcut,omitempty"`
}

func (s QuestionOption) String() string {
	return dara.Prettify(s)
}

func (s QuestionOption) GoString() string {
	return s.String()
}

func (s *QuestionOption) GetChildren() []*QuestionOption {
	return s.Children
}

func (s *QuestionOption) GetColor() *string {
	return s.Color
}

func (s *QuestionOption) GetKey() *string {
	return s.Key
}

func (s *QuestionOption) GetLabel() *string {
	return s.Label
}

func (s *QuestionOption) GetRemark() *string {
	return s.Remark
}

func (s *QuestionOption) GetShortcut() *string {
	return s.Shortcut
}

func (s *QuestionOption) SetChildren(v []*QuestionOption) *QuestionOption {
	s.Children = v
	return s
}

func (s *QuestionOption) SetColor(v string) *QuestionOption {
	s.Color = &v
	return s
}

func (s *QuestionOption) SetKey(v string) *QuestionOption {
	s.Key = &v
	return s
}

func (s *QuestionOption) SetLabel(v string) *QuestionOption {
	s.Label = &v
	return s
}

func (s *QuestionOption) SetRemark(v string) *QuestionOption {
	s.Remark = &v
	return s
}

func (s *QuestionOption) SetShortcut(v string) *QuestionOption {
	s.Shortcut = &v
	return s
}

func (s *QuestionOption) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
