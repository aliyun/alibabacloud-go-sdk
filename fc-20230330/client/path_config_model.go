// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPathConfig interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *PathConfig
	GetFunctionName() *string
	SetMethods(v []*string) *PathConfig
	GetMethods() []*string
	SetPath(v string) *PathConfig
	GetPath() *string
	SetQualifier(v string) *PathConfig
	GetQualifier() *string
	SetRewriteConfig(v *RewriteConfig) *PathConfig
	GetRewriteConfig() *RewriteConfig
}

type PathConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// myFunction
	FunctionName *string   `json:"functionName,omitempty" xml:"functionName,omitempty"`
	Methods      []*string `json:"methods" xml:"methods" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// /api/*
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// myAlias
	Qualifier     *string        `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	RewriteConfig *RewriteConfig `json:"rewriteConfig,omitempty" xml:"rewriteConfig,omitempty"`
}

func (s PathConfig) String() string {
	return dara.Prettify(s)
}

func (s PathConfig) GoString() string {
	return s.String()
}

func (s *PathConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *PathConfig) GetMethods() []*string {
	return s.Methods
}

func (s *PathConfig) GetPath() *string {
	return s.Path
}

func (s *PathConfig) GetQualifier() *string {
	return s.Qualifier
}

func (s *PathConfig) GetRewriteConfig() *RewriteConfig {
	return s.RewriteConfig
}

func (s *PathConfig) SetFunctionName(v string) *PathConfig {
	s.FunctionName = &v
	return s
}

func (s *PathConfig) SetMethods(v []*string) *PathConfig {
	s.Methods = v
	return s
}

func (s *PathConfig) SetPath(v string) *PathConfig {
	s.Path = &v
	return s
}

func (s *PathConfig) SetQualifier(v string) *PathConfig {
	s.Qualifier = &v
	return s
}

func (s *PathConfig) SetRewriteConfig(v *RewriteConfig) *PathConfig {
	s.RewriteConfig = v
	return s
}

func (s *PathConfig) Validate() error {
	return dara.Validate(s)
}
