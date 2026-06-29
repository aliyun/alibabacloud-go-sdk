// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuestionPlugin interface {
	dara.Model
	String() string
	GoString() string
	SetCanSelect(v bool) *QuestionPlugin
	GetCanSelect() *bool
	SetChildren(v []*QuestionPlugin) *QuestionPlugin
	GetChildren() []*QuestionPlugin
	SetDefaultResult(v string) *QuestionPlugin
	GetDefaultResult() *string
	SetDisplay(v bool) *QuestionPlugin
	GetDisplay() *bool
	SetExif(v map[string]interface{}) *QuestionPlugin
	GetExif() map[string]interface{}
	SetHotKeyMap(v string) *QuestionPlugin
	GetHotKeyMap() *string
	SetMarkTitle(v string) *QuestionPlugin
	GetMarkTitle() *string
	SetMarkTitleAlias(v string) *QuestionPlugin
	GetMarkTitleAlias() *string
	SetMustFill(v bool) *QuestionPlugin
	GetMustFill() *bool
	SetOptions(v []*QuestionOption) *QuestionPlugin
	GetOptions() []*QuestionOption
	SetPreOptions(v []*string) *QuestionPlugin
	GetPreOptions() []*string
	SetQuestionId(v string) *QuestionPlugin
	GetQuestionId() *string
	SetRule(v string) *QuestionPlugin
	GetRule() *string
	SetSelectGroup(v string) *QuestionPlugin
	GetSelectGroup() *string
	SetSelected(v bool) *QuestionPlugin
	GetSelected() *bool
	SetType(v string) *QuestionPlugin
	GetType() *string
}

type QuestionPlugin struct {
	// Whether it can be selected
	//
	// example:
	//
	// False
	CanSelect *bool `json:"CanSelect,omitempty" xml:"CanSelect,omitempty"`
	// List of child widgets
	Children []*QuestionPlugin `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// Default result
	//
	// example:
	//
	// 1
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// Whether it is displayed
	//
	// This parameter is required.
	//
	// example:
	//
	// True
	Display *bool `json:"Display,omitempty" xml:"Display,omitempty"`
	// Additional remarks
	//
	// example:
	//
	// false
	Exif map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// Keyboard shortcut map
	//
	// example:
	//
	// ""
	HotKeyMap *string `json:"HotKeyMap,omitempty" xml:"HotKeyMap,omitempty"`
	// Widget title
	//
	// This parameter is required.
	//
	// example:
	//
	// 内部单选
	MarkTitle *string `json:"MarkTitle,omitempty" xml:"MarkTitle,omitempty"`
	// Question alias
	//
	// example:
	//
	// 单选
	MarkTitleAlias *string `json:"MarkTitleAlias,omitempty" xml:"MarkTitleAlias,omitempty"`
	// Whether it is required
	//
	// This parameter is required.
	//
	// example:
	//
	// False
	MustFill *bool `json:"MustFill,omitempty" xml:"MustFill,omitempty"`
	// List of options configuration
	//
	// This parameter is required.
	Options []*QuestionOption `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	// List of predefined options for fill-in-the-blank questions
	PreOptions []*string `json:"PreOptions,omitempty" xml:"PreOptions,omitempty" type:"Repeated"`
	// Question widget ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	QuestionId *string `json:"QuestionId,omitempty" xml:"QuestionId,omitempty"`
	// Regular expression, validation rule
	//
	// example:
	//
	// ""
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Selection group
	//
	// example:
	//
	// g1
	SelectGroup *string `json:"SelectGroup,omitempty" xml:"SelectGroup,omitempty"`
	// Whether it is selected
	//
	// example:
	//
	// False
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
	// Widget type
	//
	// This parameter is required.
	//
	// example:
	//
	// RADIO
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QuestionPlugin) String() string {
	return dara.Prettify(s)
}

func (s QuestionPlugin) GoString() string {
	return s.String()
}

func (s *QuestionPlugin) GetCanSelect() *bool {
	return s.CanSelect
}

func (s *QuestionPlugin) GetChildren() []*QuestionPlugin {
	return s.Children
}

func (s *QuestionPlugin) GetDefaultResult() *string {
	return s.DefaultResult
}

func (s *QuestionPlugin) GetDisplay() *bool {
	return s.Display
}

func (s *QuestionPlugin) GetExif() map[string]interface{} {
	return s.Exif
}

func (s *QuestionPlugin) GetHotKeyMap() *string {
	return s.HotKeyMap
}

func (s *QuestionPlugin) GetMarkTitle() *string {
	return s.MarkTitle
}

func (s *QuestionPlugin) GetMarkTitleAlias() *string {
	return s.MarkTitleAlias
}

func (s *QuestionPlugin) GetMustFill() *bool {
	return s.MustFill
}

func (s *QuestionPlugin) GetOptions() []*QuestionOption {
	return s.Options
}

func (s *QuestionPlugin) GetPreOptions() []*string {
	return s.PreOptions
}

func (s *QuestionPlugin) GetQuestionId() *string {
	return s.QuestionId
}

func (s *QuestionPlugin) GetRule() *string {
	return s.Rule
}

func (s *QuestionPlugin) GetSelectGroup() *string {
	return s.SelectGroup
}

func (s *QuestionPlugin) GetSelected() *bool {
	return s.Selected
}

func (s *QuestionPlugin) GetType() *string {
	return s.Type
}

func (s *QuestionPlugin) SetCanSelect(v bool) *QuestionPlugin {
	s.CanSelect = &v
	return s
}

func (s *QuestionPlugin) SetChildren(v []*QuestionPlugin) *QuestionPlugin {
	s.Children = v
	return s
}

func (s *QuestionPlugin) SetDefaultResult(v string) *QuestionPlugin {
	s.DefaultResult = &v
	return s
}

func (s *QuestionPlugin) SetDisplay(v bool) *QuestionPlugin {
	s.Display = &v
	return s
}

func (s *QuestionPlugin) SetExif(v map[string]interface{}) *QuestionPlugin {
	s.Exif = v
	return s
}

func (s *QuestionPlugin) SetHotKeyMap(v string) *QuestionPlugin {
	s.HotKeyMap = &v
	return s
}

func (s *QuestionPlugin) SetMarkTitle(v string) *QuestionPlugin {
	s.MarkTitle = &v
	return s
}

func (s *QuestionPlugin) SetMarkTitleAlias(v string) *QuestionPlugin {
	s.MarkTitleAlias = &v
	return s
}

func (s *QuestionPlugin) SetMustFill(v bool) *QuestionPlugin {
	s.MustFill = &v
	return s
}

func (s *QuestionPlugin) SetOptions(v []*QuestionOption) *QuestionPlugin {
	s.Options = v
	return s
}

func (s *QuestionPlugin) SetPreOptions(v []*string) *QuestionPlugin {
	s.PreOptions = v
	return s
}

func (s *QuestionPlugin) SetQuestionId(v string) *QuestionPlugin {
	s.QuestionId = &v
	return s
}

func (s *QuestionPlugin) SetRule(v string) *QuestionPlugin {
	s.Rule = &v
	return s
}

func (s *QuestionPlugin) SetSelectGroup(v string) *QuestionPlugin {
	s.SelectGroup = &v
	return s
}

func (s *QuestionPlugin) SetSelected(v bool) *QuestionPlugin {
	s.Selected = &v
	return s
}

func (s *QuestionPlugin) SetType(v string) *QuestionPlugin {
	s.Type = &v
	return s
}

func (s *QuestionPlugin) Validate() error {
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
