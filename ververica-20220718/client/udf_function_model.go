// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUdfFunction interface {
	dara.Model
	String() string
	GoString() string
	SetClassName(v string) *UdfFunction
	GetClassName() *string
	SetFunctionName(v string) *UdfFunction
	GetFunctionName() *string
	SetUdfArtifactName(v string) *UdfFunction
	GetUdfArtifactName() *string
}

type UdfFunction struct {
	ClassName       *string `json:"className,omitempty" xml:"className,omitempty"`
	FunctionName    *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	UdfArtifactName *string `json:"udfArtifactName,omitempty" xml:"udfArtifactName,omitempty"`
}

func (s UdfFunction) String() string {
	return dara.Prettify(s)
}

func (s UdfFunction) GoString() string {
	return s.String()
}

func (s *UdfFunction) GetClassName() *string {
	return s.ClassName
}

func (s *UdfFunction) GetFunctionName() *string {
	return s.FunctionName
}

func (s *UdfFunction) GetUdfArtifactName() *string {
	return s.UdfArtifactName
}

func (s *UdfFunction) SetClassName(v string) *UdfFunction {
	s.ClassName = &v
	return s
}

func (s *UdfFunction) SetFunctionName(v string) *UdfFunction {
	s.FunctionName = &v
	return s
}

func (s *UdfFunction) SetUdfArtifactName(v string) *UdfFunction {
	s.UdfArtifactName = &v
	return s
}

func (s *UdfFunction) Validate() error {
	return dara.Validate(s)
}
