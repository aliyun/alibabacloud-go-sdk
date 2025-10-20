// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFunctionDefinition interface {
	dara.Model
	String() string
	GoString() string
	SetClassName(v string) *FunctionDefinition
	GetClassName() *string
	SetDefinition(v string) *FunctionDefinition
	GetDefinition() *string
	SetFileResources(v []*FunctionFileResource) *FunctionDefinition
	GetFileResources() []*FunctionFileResource
	SetFunctionName(v string) *FunctionDefinition
	GetFunctionName() *string
	SetLanguage(v string) *FunctionDefinition
	GetLanguage() *string
	SetType(v string) *FunctionDefinition
	GetType() *string
}

type FunctionDefinition struct {
	// required in FileFunctionDefinition
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// required in SQLFunctionDefinition/LambdaFunctionDefinition
	Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
	// required in FileFunctionDefinition
	FileResources []*FunctionFileResource `json:"fileResources,omitempty" xml:"fileResources,omitempty" type:"Repeated"`
	// required in FileFunctionDefinition
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// required in LambdaFunctionDefinition/FileFunctionDefinition
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FunctionDefinition) String() string {
	return dara.Prettify(s)
}

func (s FunctionDefinition) GoString() string {
	return s.String()
}

func (s *FunctionDefinition) GetClassName() *string {
	return s.ClassName
}

func (s *FunctionDefinition) GetDefinition() *string {
	return s.Definition
}

func (s *FunctionDefinition) GetFileResources() []*FunctionFileResource {
	return s.FileResources
}

func (s *FunctionDefinition) GetFunctionName() *string {
	return s.FunctionName
}

func (s *FunctionDefinition) GetLanguage() *string {
	return s.Language
}

func (s *FunctionDefinition) GetType() *string {
	return s.Type
}

func (s *FunctionDefinition) SetClassName(v string) *FunctionDefinition {
	s.ClassName = &v
	return s
}

func (s *FunctionDefinition) SetDefinition(v string) *FunctionDefinition {
	s.Definition = &v
	return s
}

func (s *FunctionDefinition) SetFileResources(v []*FunctionFileResource) *FunctionDefinition {
	s.FileResources = v
	return s
}

func (s *FunctionDefinition) SetFunctionName(v string) *FunctionDefinition {
	s.FunctionName = &v
	return s
}

func (s *FunctionDefinition) SetLanguage(v string) *FunctionDefinition {
	s.Language = &v
	return s
}

func (s *FunctionDefinition) SetType(v string) *FunctionDefinition {
	s.Type = &v
	return s
}

func (s *FunctionDefinition) Validate() error {
	if s.FileResources != nil {
		for _, item := range s.FileResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
