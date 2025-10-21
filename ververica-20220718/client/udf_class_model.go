// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUdfClass interface {
	dara.Model
	String() string
	GoString() string
	SetClassName(v string) *UdfClass
	GetClassName() *string
	SetClassType(v string) *UdfClass
	GetClassType() *string
	SetFunctionNames(v []*string) *UdfClass
	GetFunctionNames() []*string
	SetUdfArtifactName(v string) *UdfClass
	GetUdfArtifactName() *string
}

type UdfClass struct {
	ClassName       *string   `json:"className,omitempty" xml:"className,omitempty"`
	ClassType       *string   `json:"classType,omitempty" xml:"classType,omitempty"`
	FunctionNames   []*string `json:"functionNames,omitempty" xml:"functionNames,omitempty" type:"Repeated"`
	UdfArtifactName *string   `json:"udfArtifactName,omitempty" xml:"udfArtifactName,omitempty"`
}

func (s UdfClass) String() string {
	return dara.Prettify(s)
}

func (s UdfClass) GoString() string {
	return s.String()
}

func (s *UdfClass) GetClassName() *string {
	return s.ClassName
}

func (s *UdfClass) GetClassType() *string {
	return s.ClassType
}

func (s *UdfClass) GetFunctionNames() []*string {
	return s.FunctionNames
}

func (s *UdfClass) GetUdfArtifactName() *string {
	return s.UdfArtifactName
}

func (s *UdfClass) SetClassName(v string) *UdfClass {
	s.ClassName = &v
	return s
}

func (s *UdfClass) SetClassType(v string) *UdfClass {
	s.ClassType = &v
	return s
}

func (s *UdfClass) SetFunctionNames(v []*string) *UdfClass {
	s.FunctionNames = v
	return s
}

func (s *UdfClass) SetUdfArtifactName(v string) *UdfClass {
	s.UdfArtifactName = &v
	return s
}

func (s *UdfClass) Validate() error {
	return dara.Validate(s)
}
