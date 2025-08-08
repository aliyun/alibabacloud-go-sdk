// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateSpec interface {
	dara.Model
	String() string
	GoString() string
	SetAuthor(v string) *TemplateSpec
	GetAuthor() *string
	SetCategory(v string) *TemplateSpec
	GetCategory() *string
	SetLicense(v string) *TemplateSpec
	GetLicense() *string
	SetPackageName(v string) *TemplateSpec
	GetPackageName() *string
	SetReadme(v string) *TemplateSpec
	GetReadme() *string
	SetRegistryToken(v string) *TemplateSpec
	GetRegistryToken() *string
	SetServices(v map[string]*TemplateServiceConfig) *TemplateSpec
	GetServices() map[string]*TemplateServiceConfig
	SetSource(v *TemplateSpecSource) *TemplateSpec
	GetSource() *TemplateSpecSource
	SetVariables(v map[string]*TemplateParameterSchema) *TemplateSpec
	GetVariables() map[string]*TemplateParameterSchema
	SetVersion(v string) *TemplateSpec
	GetVersion() *string
}

type TemplateSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// CAP
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AI
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// Apache-1.0
	License *string `json:"license,omitempty" xml:"license,omitempty"`
	// example:
	//
	// demo-package
	PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty"`
	// This parameter is required.
	Readme        *string                             `json:"readme,omitempty" xml:"readme,omitempty"`
	RegistryToken *string                             `json:"registryToken,omitempty" xml:"registryToken,omitempty"`
	Services      map[string]*TemplateServiceConfig   `json:"services,omitempty" xml:"services,omitempty"`
	Source        *TemplateSpecSource                 `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
	Variables     map[string]*TemplateParameterSchema `json:"variables,omitempty" xml:"variables,omitempty"`
	Version       *string                             `json:"version,omitempty" xml:"version,omitempty"`
}

func (s TemplateSpec) String() string {
	return dara.Prettify(s)
}

func (s TemplateSpec) GoString() string {
	return s.String()
}

func (s *TemplateSpec) GetAuthor() *string {
	return s.Author
}

func (s *TemplateSpec) GetCategory() *string {
	return s.Category
}

func (s *TemplateSpec) GetLicense() *string {
	return s.License
}

func (s *TemplateSpec) GetPackageName() *string {
	return s.PackageName
}

func (s *TemplateSpec) GetReadme() *string {
	return s.Readme
}

func (s *TemplateSpec) GetRegistryToken() *string {
	return s.RegistryToken
}

func (s *TemplateSpec) GetServices() map[string]*TemplateServiceConfig {
	return s.Services
}

func (s *TemplateSpec) GetSource() *TemplateSpecSource {
	return s.Source
}

func (s *TemplateSpec) GetVariables() map[string]*TemplateParameterSchema {
	return s.Variables
}

func (s *TemplateSpec) GetVersion() *string {
	return s.Version
}

func (s *TemplateSpec) SetAuthor(v string) *TemplateSpec {
	s.Author = &v
	return s
}

func (s *TemplateSpec) SetCategory(v string) *TemplateSpec {
	s.Category = &v
	return s
}

func (s *TemplateSpec) SetLicense(v string) *TemplateSpec {
	s.License = &v
	return s
}

func (s *TemplateSpec) SetPackageName(v string) *TemplateSpec {
	s.PackageName = &v
	return s
}

func (s *TemplateSpec) SetReadme(v string) *TemplateSpec {
	s.Readme = &v
	return s
}

func (s *TemplateSpec) SetRegistryToken(v string) *TemplateSpec {
	s.RegistryToken = &v
	return s
}

func (s *TemplateSpec) SetServices(v map[string]*TemplateServiceConfig) *TemplateSpec {
	s.Services = v
	return s
}

func (s *TemplateSpec) SetSource(v *TemplateSpecSource) *TemplateSpec {
	s.Source = v
	return s
}

func (s *TemplateSpec) SetVariables(v map[string]*TemplateParameterSchema) *TemplateSpec {
	s.Variables = v
	return s
}

func (s *TemplateSpec) SetVersion(v string) *TemplateSpec {
	s.Version = &v
	return s
}

func (s *TemplateSpec) Validate() error {
	return dara.Validate(s)
}

type TemplateSpecSource struct {
	Repository *RepositorySourceConfig `json:"repository,omitempty" xml:"repository,omitempty"`
}

func (s TemplateSpecSource) String() string {
	return dara.Prettify(s)
}

func (s TemplateSpecSource) GoString() string {
	return s.String()
}

func (s *TemplateSpecSource) GetRepository() *RepositorySourceConfig {
	return s.Repository
}

func (s *TemplateSpecSource) SetRepository(v *RepositorySourceConfig) *TemplateSpecSource {
	s.Repository = v
	return s
}

func (s *TemplateSpecSource) Validate() error {
	return dara.Validate(s)
}
