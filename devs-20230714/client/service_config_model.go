// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetArtifact(v *ArtifactMeta) *ServiceConfig
	GetArtifact() *ArtifactMeta
	SetBuild(v *BuildConfig) *ServiceConfig
	GetBuild() *BuildConfig
	SetComponent(v string) *ServiceConfig
	GetComponent() *string
	SetProps(v map[string]interface{}) *ServiceConfig
	GetProps() map[string]interface{}
	SetSource(v *SourceConfig) *ServiceConfig
	GetSource() *SourceConfig
	SetType(v string) *ServiceConfig
	GetType() *string
	SetVariables(v map[string]*Variable) *ServiceConfig
	GetVariables() map[string]*Variable
}

type ServiceConfig struct {
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
	Type      *string              `json:"type,omitempty" xml:"type,omitempty"`
	Variables map[string]*Variable `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s ServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s ServiceConfig) GoString() string {
	return s.String()
}

func (s *ServiceConfig) GetArtifact() *ArtifactMeta {
	return s.Artifact
}

func (s *ServiceConfig) GetBuild() *BuildConfig {
	return s.Build
}

func (s *ServiceConfig) GetComponent() *string {
	return s.Component
}

func (s *ServiceConfig) GetProps() map[string]interface{} {
	return s.Props
}

func (s *ServiceConfig) GetSource() *SourceConfig {
	return s.Source
}

func (s *ServiceConfig) GetType() *string {
	return s.Type
}

func (s *ServiceConfig) GetVariables() map[string]*Variable {
	return s.Variables
}

func (s *ServiceConfig) SetArtifact(v *ArtifactMeta) *ServiceConfig {
	s.Artifact = v
	return s
}

func (s *ServiceConfig) SetBuild(v *BuildConfig) *ServiceConfig {
	s.Build = v
	return s
}

func (s *ServiceConfig) SetComponent(v string) *ServiceConfig {
	s.Component = &v
	return s
}

func (s *ServiceConfig) SetProps(v map[string]interface{}) *ServiceConfig {
	s.Props = v
	return s
}

func (s *ServiceConfig) SetSource(v *SourceConfig) *ServiceConfig {
	s.Source = v
	return s
}

func (s *ServiceConfig) SetType(v string) *ServiceConfig {
	s.Type = &v
	return s
}

func (s *ServiceConfig) SetVariables(v map[string]*Variable) *ServiceConfig {
	s.Variables = v
	return s
}

func (s *ServiceConfig) Validate() error {
	return dara.Validate(s)
}
