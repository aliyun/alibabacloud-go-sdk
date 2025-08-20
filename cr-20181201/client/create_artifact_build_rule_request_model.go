// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactBuildRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *CreateArtifactBuildRuleRequest
	GetArtifactType() *string
	SetInstanceId(v string) *CreateArtifactBuildRuleRequest
	GetInstanceId() *string
	SetParameters(v map[string]interface{}) *CreateArtifactBuildRuleRequest
	GetParameters() map[string]interface{}
	SetScopeId(v string) *CreateArtifactBuildRuleRequest
	GetScopeId() *string
	SetScopeType(v string) *CreateArtifactBuildRuleRequest
	GetScopeType() *string
}

type CreateArtifactBuildRuleRequest struct {
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
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s CreateArtifactBuildRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactBuildRuleRequest) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *CreateArtifactBuildRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateArtifactBuildRuleRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateArtifactBuildRuleRequest) GetScopeId() *string {
	return s.ScopeId
}

func (s *CreateArtifactBuildRuleRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *CreateArtifactBuildRuleRequest) SetArtifactType(v string) *CreateArtifactBuildRuleRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactBuildRuleRequest) SetInstanceId(v string) *CreateArtifactBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateArtifactBuildRuleRequest) SetParameters(v map[string]interface{}) *CreateArtifactBuildRuleRequest {
	s.Parameters = v
	return s
}

func (s *CreateArtifactBuildRuleRequest) SetScopeId(v string) *CreateArtifactBuildRuleRequest {
	s.ScopeId = &v
	return s
}

func (s *CreateArtifactBuildRuleRequest) SetScopeType(v string) *CreateArtifactBuildRuleRequest {
	s.ScopeType = &v
	return s
}

func (s *CreateArtifactBuildRuleRequest) Validate() error {
	return dara.Validate(s)
}
