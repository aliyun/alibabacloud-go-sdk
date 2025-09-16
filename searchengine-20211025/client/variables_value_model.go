// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVariablesValue interface {
	dara.Model
	String() string
	GoString() string
	SetDisableModify(v bool) *VariablesValue
	GetDisableModify() *bool
	SetIsModify(v bool) *VariablesValue
	GetIsModify() *bool
	SetValue(v string) *VariablesValue
	GetValue() *string
	SetDescription(v string) *VariablesValue
	GetDescription() *string
	SetTemplateValue(v string) *VariablesValue
	GetTemplateValue() *string
	SetType(v string) *VariablesValue
	GetType() *string
	SetFuncValue(v *VariablesValueFuncValue) *VariablesValue
	GetFuncValue() *VariablesValueFuncValue
}

type VariablesValue struct {
	// Specifies whether the variable is not allowed to be modified.
	//
	// example:
	//
	// false
	DisableModify *bool `json:"disableModify,omitempty" xml:"disableModify,omitempty"`
	// Specifies whether the variable is modified.
	//
	// example:
	//
	// false
	IsModify *bool `json:"isModify,omitempty" xml:"isModify,omitempty"`
	// The variable value.
	//
	// example:
	//
	// ""
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// The description of the variable.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The template value of the variable.
	//
	// example:
	//
	// ""
	TemplateValue *string `json:"templateValue,omitempty" xml:"templateValue,omitempty"`
	// The variable type. Valid values:
	//
	// 	- NORMAL: common variable
	//
	// 	- FUNCTION: function variable
	//
	// example:
	//
	// NORMAL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The function variables.
	FuncValue *VariablesValueFuncValue `json:"funcValue,omitempty" xml:"funcValue,omitempty" type:"Struct"`
}

func (s VariablesValue) String() string {
	return dara.Prettify(s)
}

func (s VariablesValue) GoString() string {
	return s.String()
}

func (s *VariablesValue) GetDisableModify() *bool {
	return s.DisableModify
}

func (s *VariablesValue) GetIsModify() *bool {
	return s.IsModify
}

func (s *VariablesValue) GetValue() *string {
	return s.Value
}

func (s *VariablesValue) GetDescription() *string {
	return s.Description
}

func (s *VariablesValue) GetTemplateValue() *string {
	return s.TemplateValue
}

func (s *VariablesValue) GetType() *string {
	return s.Type
}

func (s *VariablesValue) GetFuncValue() *VariablesValueFuncValue {
	return s.FuncValue
}

func (s *VariablesValue) SetDisableModify(v bool) *VariablesValue {
	s.DisableModify = &v
	return s
}

func (s *VariablesValue) SetIsModify(v bool) *VariablesValue {
	s.IsModify = &v
	return s
}

func (s *VariablesValue) SetValue(v string) *VariablesValue {
	s.Value = &v
	return s
}

func (s *VariablesValue) SetDescription(v string) *VariablesValue {
	s.Description = &v
	return s
}

func (s *VariablesValue) SetTemplateValue(v string) *VariablesValue {
	s.TemplateValue = &v
	return s
}

func (s *VariablesValue) SetType(v string) *VariablesValue {
	s.Type = &v
	return s
}

func (s *VariablesValue) SetFuncValue(v *VariablesValueFuncValue) *VariablesValue {
	s.FuncValue = v
	return s
}

func (s *VariablesValue) Validate() error {
	return dara.Validate(s)
}

type VariablesValueFuncValue struct {
	// The class name of the function variable.
	//
	// example:
	//
	// ""
	FuncClassName *string `json:"funcClassName,omitempty" xml:"funcClassName,omitempty"`
	// The template of the function variable.
	//
	// example:
	//
	// ""
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
}

func (s VariablesValueFuncValue) String() string {
	return dara.Prettify(s)
}

func (s VariablesValueFuncValue) GoString() string {
	return s.String()
}

func (s *VariablesValueFuncValue) GetFuncClassName() *string {
	return s.FuncClassName
}

func (s *VariablesValueFuncValue) GetTemplate() *string {
	return s.Template
}

func (s *VariablesValueFuncValue) SetFuncClassName(v string) *VariablesValueFuncValue {
	s.FuncClassName = &v
	return s
}

func (s *VariablesValueFuncValue) SetTemplate(v string) *VariablesValueFuncValue {
	s.Template = &v
	return s
}

func (s *VariablesValueFuncValue) Validate() error {
	return dara.Validate(s)
}
