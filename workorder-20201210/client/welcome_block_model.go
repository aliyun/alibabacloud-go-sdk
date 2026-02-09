// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWelcomeBlock interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *WelcomeBlock
	GetDescription() *string
	SetDisplayType(v string) *WelcomeBlock
	GetDisplayType() *string
	SetIcon(v string) *WelcomeBlock
	GetIcon() *string
	SetLabel(v string) *WelcomeBlock
	GetLabel() *string
	SetType(v string) *WelcomeBlock
	GetType() *string
	SetValue(v string) *WelcomeBlock
	GetValue() *string
}

type WelcomeBlock struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayType *string `json:"DisplayType,omitempty" xml:"DisplayType,omitempty"`
	Icon        *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s WelcomeBlock) String() string {
	return dara.Prettify(s)
}

func (s WelcomeBlock) GoString() string {
	return s.String()
}

func (s *WelcomeBlock) GetDescription() *string {
	return s.Description
}

func (s *WelcomeBlock) GetDisplayType() *string {
	return s.DisplayType
}

func (s *WelcomeBlock) GetIcon() *string {
	return s.Icon
}

func (s *WelcomeBlock) GetLabel() *string {
	return s.Label
}

func (s *WelcomeBlock) GetType() *string {
	return s.Type
}

func (s *WelcomeBlock) GetValue() *string {
	return s.Value
}

func (s *WelcomeBlock) SetDescription(v string) *WelcomeBlock {
	s.Description = &v
	return s
}

func (s *WelcomeBlock) SetDisplayType(v string) *WelcomeBlock {
	s.DisplayType = &v
	return s
}

func (s *WelcomeBlock) SetIcon(v string) *WelcomeBlock {
	s.Icon = &v
	return s
}

func (s *WelcomeBlock) SetLabel(v string) *WelcomeBlock {
	s.Label = &v
	return s
}

func (s *WelcomeBlock) SetType(v string) *WelcomeBlock {
	s.Type = &v
	return s
}

func (s *WelcomeBlock) SetValue(v string) *WelcomeBlock {
	s.Value = &v
	return s
}

func (s *WelcomeBlock) Validate() error {
	return dara.Validate(s)
}
