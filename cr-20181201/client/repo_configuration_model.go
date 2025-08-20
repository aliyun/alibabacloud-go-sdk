// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepoConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactBuildRuleParameters(v *RepoConfigurationArtifactBuildRuleParameters) *RepoConfiguration
	GetArtifactBuildRuleParameters() *RepoConfigurationArtifactBuildRuleParameters
	SetRepoType(v string) *RepoConfiguration
	GetRepoType() *string
	SetTagImmutability(v bool) *RepoConfiguration
	GetTagImmutability() *bool
}

type RepoConfiguration struct {
	ArtifactBuildRuleParameters *RepoConfigurationArtifactBuildRuleParameters `json:"ArtifactBuildRuleParameters,omitempty" xml:"ArtifactBuildRuleParameters,omitempty" type:"Struct"`
	// This parameter is required.
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// This parameter is required.
	TagImmutability *bool `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
}

func (s RepoConfiguration) String() string {
	return dara.Prettify(s)
}

func (s RepoConfiguration) GoString() string {
	return s.String()
}

func (s *RepoConfiguration) GetArtifactBuildRuleParameters() *RepoConfigurationArtifactBuildRuleParameters {
	return s.ArtifactBuildRuleParameters
}

func (s *RepoConfiguration) GetRepoType() *string {
	return s.RepoType
}

func (s *RepoConfiguration) GetTagImmutability() *bool {
	return s.TagImmutability
}

func (s *RepoConfiguration) SetArtifactBuildRuleParameters(v *RepoConfigurationArtifactBuildRuleParameters) *RepoConfiguration {
	s.ArtifactBuildRuleParameters = v
	return s
}

func (s *RepoConfiguration) SetRepoType(v string) *RepoConfiguration {
	s.RepoType = &v
	return s
}

func (s *RepoConfiguration) SetTagImmutability(v bool) *RepoConfiguration {
	s.TagImmutability = &v
	return s
}

func (s *RepoConfiguration) Validate() error {
	return dara.Validate(s)
}

type RepoConfigurationArtifactBuildRuleParameters struct {
	// This parameter is required.
	ImageIndexOnly *bool `json:"ImageIndexOnly,omitempty" xml:"ImageIndexOnly,omitempty"`
}

func (s RepoConfigurationArtifactBuildRuleParameters) String() string {
	return dara.Prettify(s)
}

func (s RepoConfigurationArtifactBuildRuleParameters) GoString() string {
	return s.String()
}

func (s *RepoConfigurationArtifactBuildRuleParameters) GetImageIndexOnly() *bool {
	return s.ImageIndexOnly
}

func (s *RepoConfigurationArtifactBuildRuleParameters) SetImageIndexOnly(v bool) *RepoConfigurationArtifactBuildRuleParameters {
	s.ImageIndexOnly = &v
	return s
}

func (s *RepoConfigurationArtifactBuildRuleParameters) Validate() error {
	return dara.Validate(s)
}
