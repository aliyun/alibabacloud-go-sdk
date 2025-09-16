// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilesConfigVariablesValue interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *FilesConfigVariablesValue
	GetDescription() *string
	SetDisableModify(v bool) *FilesConfigVariablesValue
	GetDisableModify() *bool
	SetIsModify(v bool) *FilesConfigVariablesValue
	GetIsModify() *bool
	SetType(v string) *FilesConfigVariablesValue
	GetType() *string
	SetValue(v string) *FilesConfigVariablesValue
	GetValue() *string
}

type FilesConfigVariablesValue struct {
	// The description of the variable.
	//
	// example:
	//
	// Custom variable
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
	// true
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

func (s FilesConfigVariablesValue) String() string {
	return dara.Prettify(s)
}

func (s FilesConfigVariablesValue) GoString() string {
	return s.String()
}

func (s *FilesConfigVariablesValue) GetDescription() *string {
	return s.Description
}

func (s *FilesConfigVariablesValue) GetDisableModify() *bool {
	return s.DisableModify
}

func (s *FilesConfigVariablesValue) GetIsModify() *bool {
	return s.IsModify
}

func (s *FilesConfigVariablesValue) GetType() *string {
	return s.Type
}

func (s *FilesConfigVariablesValue) GetValue() *string {
	return s.Value
}

func (s *FilesConfigVariablesValue) SetDescription(v string) *FilesConfigVariablesValue {
	s.Description = &v
	return s
}

func (s *FilesConfigVariablesValue) SetDisableModify(v bool) *FilesConfigVariablesValue {
	s.DisableModify = &v
	return s
}

func (s *FilesConfigVariablesValue) SetIsModify(v bool) *FilesConfigVariablesValue {
	s.IsModify = &v
	return s
}

func (s *FilesConfigVariablesValue) SetType(v string) *FilesConfigVariablesValue {
	s.Type = &v
	return s
}

func (s *FilesConfigVariablesValue) SetValue(v string) *FilesConfigVariablesValue {
	s.Value = &v
	return s
}

func (s *FilesConfigVariablesValue) Validate() error {
	return dara.Validate(s)
}
