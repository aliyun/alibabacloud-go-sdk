// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactBuildRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *CreateArtifactBuildRuleShrinkRequest
	GetArtifactType() *string
	SetInstanceId(v string) *CreateArtifactBuildRuleShrinkRequest
	GetInstanceId() *string
	SetParametersShrink(v string) *CreateArtifactBuildRuleShrinkRequest
	GetParametersShrink() *string
	SetScopeId(v string) *CreateArtifactBuildRuleShrinkRequest
	GetScopeId() *string
	SetScopeType(v string) *CreateArtifactBuildRuleShrinkRequest
	GetScopeType() *string
}

type CreateArtifactBuildRuleShrinkRequest struct {
	// The type of the artifact.
	//
	// 	- `ACCELERATED_IMAGE`: accelerated images.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACCELERATED_IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-cxreylqvcyje****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Additional parameters.
	//
	// example:
	//
	// {}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the effective range of the rule.
	//
	// 	- Set the value to the ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-pmajihou6cg0****
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// The effective range of the rule. Valid values:
	//
	// 	- `REPOSITORY`
	//
	// This parameter is required.
	//
	// example:
	//
	// REPOSITORY
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s CreateArtifactBuildRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactBuildRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactBuildRuleShrinkRequest) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *CreateArtifactBuildRuleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateArtifactBuildRuleShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *CreateArtifactBuildRuleShrinkRequest) GetScopeId() *string {
	return s.ScopeId
}

func (s *CreateArtifactBuildRuleShrinkRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *CreateArtifactBuildRuleShrinkRequest) SetArtifactType(v string) *CreateArtifactBuildRuleShrinkRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactBuildRuleShrinkRequest) SetInstanceId(v string) *CreateArtifactBuildRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateArtifactBuildRuleShrinkRequest) SetParametersShrink(v string) *CreateArtifactBuildRuleShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateArtifactBuildRuleShrinkRequest) SetScopeId(v string) *CreateArtifactBuildRuleShrinkRequest {
	s.ScopeId = &v
	return s
}

func (s *CreateArtifactBuildRuleShrinkRequest) SetScopeType(v string) *CreateArtifactBuildRuleShrinkRequest {
	s.ScopeType = &v
	return s
}

func (s *CreateArtifactBuildRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
