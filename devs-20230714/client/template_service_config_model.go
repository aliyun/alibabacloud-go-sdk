// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplateServiceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetArtifact(v *ArtifactMeta) *TemplateServiceConfig
	GetArtifact() *ArtifactMeta
	SetBuild(v *BuildConfig) *TemplateServiceConfig
	GetBuild() *BuildConfig
	SetComponent(v string) *TemplateServiceConfig
	GetComponent() *string
	SetProps(v map[string]interface{}) *TemplateServiceConfig
	GetProps() map[string]interface{}
	SetSource(v *SourceConfig) *TemplateServiceConfig
	GetSource() *SourceConfig
	SetType(v string) *TemplateServiceConfig
	GetType() *string
	SetVariables(v map[string]*TemplateParameterSchema) *TemplateServiceConfig
	GetVariables() map[string]*TemplateParameterSchema
}

type TemplateServiceConfig struct {
	Artifact *ArtifactMeta `json:"artifact,omitempty" xml:"artifact,omitempty"`
	Build    *BuildConfig  `json:"build,omitempty" xml:"build,omitempty"`
	// example:
	//
	// fc3@1.0.0
	Component *string                `json:"component,omitempty" xml:"component,omitempty"`
	Props     map[string]interface{} `json:"props,omitempty" xml:"props,omitempty"`
	Source    *SourceConfig          `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// Function
	Type      *string                             `json:"type,omitempty" xml:"type,omitempty"`
	Variables map[string]*TemplateParameterSchema `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s TemplateServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s TemplateServiceConfig) GoString() string {
	return s.String()
}

func (s *TemplateServiceConfig) GetArtifact() *ArtifactMeta {
	return s.Artifact
}

func (s *TemplateServiceConfig) GetBuild() *BuildConfig {
	return s.Build
}

func (s *TemplateServiceConfig) GetComponent() *string {
	return s.Component
}

func (s *TemplateServiceConfig) GetProps() map[string]interface{} {
	return s.Props
}

func (s *TemplateServiceConfig) GetSource() *SourceConfig {
	return s.Source
}

func (s *TemplateServiceConfig) GetType() *string {
	return s.Type
}

func (s *TemplateServiceConfig) GetVariables() map[string]*TemplateParameterSchema {
	return s.Variables
}

func (s *TemplateServiceConfig) SetArtifact(v *ArtifactMeta) *TemplateServiceConfig {
	s.Artifact = v
	return s
}

func (s *TemplateServiceConfig) SetBuild(v *BuildConfig) *TemplateServiceConfig {
	s.Build = v
	return s
}

func (s *TemplateServiceConfig) SetComponent(v string) *TemplateServiceConfig {
	s.Component = &v
	return s
}

func (s *TemplateServiceConfig) SetProps(v map[string]interface{}) *TemplateServiceConfig {
	s.Props = v
	return s
}

func (s *TemplateServiceConfig) SetSource(v *SourceConfig) *TemplateServiceConfig {
	s.Source = v
	return s
}

func (s *TemplateServiceConfig) SetType(v string) *TemplateServiceConfig {
	s.Type = &v
	return s
}

func (s *TemplateServiceConfig) SetVariables(v map[string]*TemplateParameterSchema) *TemplateServiceConfig {
	s.Variables = v
	return s
}

func (s *TemplateServiceConfig) Validate() error {
	return dara.Validate(s)
}
