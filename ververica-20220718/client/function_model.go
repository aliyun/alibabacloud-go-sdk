// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunction interface {
	dara.Model
	String() string
	GoString() string
	SetClassName(v string) *Function
	GetClassName() *string
	SetDescription(v string) *Function
	GetDescription() *string
	SetFunctionLanguage(v string) *Function
	GetFunctionLanguage() *string
	SetFunctionType(v string) *Function
	GetFunctionType() *string
	SetGmtModified(v int64) *Function
	GetGmtModified() *int64
	SetName(v string) *Function
	GetName() *string
}

type Function struct {
	// This parameter is required.
	ClassName   *string `json:"className,omitempty" xml:"className,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	FunctionLanguage *string `json:"functionLanguage,omitempty" xml:"functionLanguage,omitempty"`
	// This parameter is required.
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	GmtModified  *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s Function) String() string {
	return dara.Prettify(s)
}

func (s Function) GoString() string {
	return s.String()
}

func (s *Function) GetClassName() *string {
	return s.ClassName
}

func (s *Function) GetDescription() *string {
	return s.Description
}

func (s *Function) GetFunctionLanguage() *string {
	return s.FunctionLanguage
}

func (s *Function) GetFunctionType() *string {
	return s.FunctionType
}

func (s *Function) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *Function) GetName() *string {
	return s.Name
}

func (s *Function) SetClassName(v string) *Function {
	s.ClassName = &v
	return s
}

func (s *Function) SetDescription(v string) *Function {
	s.Description = &v
	return s
}

func (s *Function) SetFunctionLanguage(v string) *Function {
	s.FunctionLanguage = &v
	return s
}

func (s *Function) SetFunctionType(v string) *Function {
	s.FunctionType = &v
	return s
}

func (s *Function) SetGmtModified(v int64) *Function {
	s.GmtModified = &v
	return s
}

func (s *Function) SetName(v string) *Function {
	s.Name = &v
	return s
}

func (s *Function) Validate() error {
	return dara.Validate(s)
}
