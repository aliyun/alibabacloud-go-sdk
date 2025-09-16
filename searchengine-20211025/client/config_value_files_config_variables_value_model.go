// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigValueFilesConfigVariablesValue interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ConfigValueFilesConfigVariablesValue
	GetDescription() *string
	SetDisableModify(v bool) *ConfigValueFilesConfigVariablesValue
	GetDisableModify() *bool
	SetIsModify(v bool) *ConfigValueFilesConfigVariablesValue
	GetIsModify() *bool
	SetType(v string) *ConfigValueFilesConfigVariablesValue
	GetType() *string
	SetValue(v string) *ConfigValueFilesConfigVariablesValue
	GetValue() *string
}

type ConfigValueFilesConfigVariablesValue struct {
	// The description of the variable.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether the variable is not allowed to be modified.
	//
	// example:
	//
	// true
	DisableModify *bool `json:"disableModify,omitempty" xml:"disableModify,omitempty"`
	// Specifies whether the variable is modified.
	//
	// example:
	//
	// false
	IsModify *bool `json:"isModify,omitempty" xml:"isModify,omitempty"`
	// The variable type. Valid values: NORMAL: common variable. FUNCTION: function variable.
	//
	// example:
	//
	// NORMAL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The variable value.
	//
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ConfigValueFilesConfigVariablesValue) String() string {
	return dara.Prettify(s)
}

func (s ConfigValueFilesConfigVariablesValue) GoString() string {
	return s.String()
}

func (s *ConfigValueFilesConfigVariablesValue) GetDescription() *string {
	return s.Description
}

func (s *ConfigValueFilesConfigVariablesValue) GetDisableModify() *bool {
	return s.DisableModify
}

func (s *ConfigValueFilesConfigVariablesValue) GetIsModify() *bool {
	return s.IsModify
}

func (s *ConfigValueFilesConfigVariablesValue) GetType() *string {
	return s.Type
}

func (s *ConfigValueFilesConfigVariablesValue) GetValue() *string {
	return s.Value
}

func (s *ConfigValueFilesConfigVariablesValue) SetDescription(v string) *ConfigValueFilesConfigVariablesValue {
	s.Description = &v
	return s
}

func (s *ConfigValueFilesConfigVariablesValue) SetDisableModify(v bool) *ConfigValueFilesConfigVariablesValue {
	s.DisableModify = &v
	return s
}

func (s *ConfigValueFilesConfigVariablesValue) SetIsModify(v bool) *ConfigValueFilesConfigVariablesValue {
	s.IsModify = &v
	return s
}

func (s *ConfigValueFilesConfigVariablesValue) SetType(v string) *ConfigValueFilesConfigVariablesValue {
	s.Type = &v
	return s
}

func (s *ConfigValueFilesConfigVariablesValue) SetValue(v string) *ConfigValueFilesConfigVariablesValue {
	s.Value = &v
	return s
}

func (s *ConfigValueFilesConfigVariablesValue) Validate() error {
	return dara.Validate(s)
}
