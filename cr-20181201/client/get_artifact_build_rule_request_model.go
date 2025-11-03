// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactBuildRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *GetArtifactBuildRuleRequest
	GetArtifactType() *string
	SetBuildRuleId(v string) *GetArtifactBuildRuleRequest
	GetBuildRuleId() *string
	SetInstanceId(v string) *GetArtifactBuildRuleRequest
	GetInstanceId() *string
	SetScopeId(v string) *GetArtifactBuildRuleRequest
	GetScopeId() *string
	SetScopeType(v string) *GetArtifactBuildRuleRequest
	GetScopeType() *string
}

type GetArtifactBuildRuleRequest struct {
	// The type of the artifact.
	//
	// 	- `ACCELERATED_IMAGE`: accelerated images.
	//
	// example:
	//
	// ACCELERATED_IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The ID of the artifact building rule.
	//
	// example:
	//
	// crabr-o2670wqz2n70****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the effective range of the artifact building rule.
	//
	// 	- Set the value to the ID of the image repository.
	//
	// example:
	//
	// crr-8dz3aedjqlmk****
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// The effective range of the artifact building rule. Valid values:
	//
	// 	- `REPOSITORY`: The artifact building rule is effective in the repository level.
	//
	// example:
	//
	// REPOSITORY
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s GetArtifactBuildRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildRuleRequest) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *GetArtifactBuildRuleRequest) GetBuildRuleId() *string {
	return s.BuildRuleId
}

func (s *GetArtifactBuildRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetArtifactBuildRuleRequest) GetScopeId() *string {
	return s.ScopeId
}

func (s *GetArtifactBuildRuleRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *GetArtifactBuildRuleRequest) SetArtifactType(v string) *GetArtifactBuildRuleRequest {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactBuildRuleRequest) SetBuildRuleId(v string) *GetArtifactBuildRuleRequest {
	s.BuildRuleId = &v
	return s
}

func (s *GetArtifactBuildRuleRequest) SetInstanceId(v string) *GetArtifactBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactBuildRuleRequest) SetScopeId(v string) *GetArtifactBuildRuleRequest {
	s.ScopeId = &v
	return s
}

func (s *GetArtifactBuildRuleRequest) SetScopeType(v string) *GetArtifactBuildRuleRequest {
	s.ScopeType = &v
	return s
}

func (s *GetArtifactBuildRuleRequest) Validate() error {
	return dara.Validate(s)
}
