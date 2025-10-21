// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploymentDraft interface {
	dara.Model
	String() string
	GoString() string
	SetArtifact(v *Artifact) *DeploymentDraft
	GetArtifact() *Artifact
	SetCreatedAt(v int64) *DeploymentDraft
	GetCreatedAt() *int64
	SetCreator(v string) *DeploymentDraft
	GetCreator() *string
	SetCreatorName(v string) *DeploymentDraft
	GetCreatorName() *string
	SetDeploymentDraftId(v string) *DeploymentDraft
	GetDeploymentDraftId() *string
	SetEngineVersion(v string) *DeploymentDraft
	GetEngineVersion() *string
	SetExecutionMode(v string) *DeploymentDraft
	GetExecutionMode() *string
	SetLabels(v map[string]interface{}) *DeploymentDraft
	GetLabels() map[string]interface{}
	SetLocalVariables(v []*LocalVariable) *DeploymentDraft
	GetLocalVariables() []*LocalVariable
	SetLock(v *Lock) *DeploymentDraft
	GetLock() *Lock
	SetModifiedAt(v int64) *DeploymentDraft
	GetModifiedAt() *int64
	SetModifier(v string) *DeploymentDraft
	GetModifier() *string
	SetModifierName(v string) *DeploymentDraft
	GetModifierName() *string
	SetName(v string) *DeploymentDraft
	GetName() *string
	SetNamespace(v string) *DeploymentDraft
	GetNamespace() *string
	SetParentId(v string) *DeploymentDraft
	GetParentId() *string
	SetReferencedDeploymentId(v string) *DeploymentDraft
	GetReferencedDeploymentId() *string
	SetWorkspace(v string) *DeploymentDraft
	GetWorkspace() *string
}

type DeploymentDraft struct {
	Artifact  *Artifact `json:"artifact,omitempty" xml:"artifact,omitempty"`
	CreatedAt *int64    `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-00000012****
	DeploymentDraftId *string `json:"deploymentDraftId,omitempty" xml:"deploymentDraftId,omitempty"`
	// example:
	//
	// vvr-6.0.7-flink-1.15
	EngineVersion *string `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	// example:
	//
	// STREAMING
	ExecutionMode  *string                `json:"executionMode,omitempty" xml:"executionMode,omitempty"`
	Labels         map[string]interface{} `json:"labels,omitempty" xml:"labels,omitempty"`
	LocalVariables []*LocalVariable       `json:"localVariables,omitempty" xml:"localVariables,omitempty" type:"Repeated"`
	Lock           *Lock                  `json:"lock,omitempty" xml:"lock,omitempty"`
	ModifiedAt     *int64                 `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// example:
	//
	// 27846363877456****
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// ****@streamcompute.onaliyun.com
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// example:
	//
	// test-draft
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-00000013****
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-0000012312****
	ReferencedDeploymentId *string `json:"referencedDeploymentId,omitempty" xml:"referencedDeploymentId,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeploymentDraft) String() string {
	return dara.Prettify(s)
}

func (s DeploymentDraft) GoString() string {
	return s.String()
}

func (s *DeploymentDraft) GetArtifact() *Artifact {
	return s.Artifact
}

func (s *DeploymentDraft) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *DeploymentDraft) GetCreator() *string {
	return s.Creator
}

func (s *DeploymentDraft) GetCreatorName() *string {
	return s.CreatorName
}

func (s *DeploymentDraft) GetDeploymentDraftId() *string {
	return s.DeploymentDraftId
}

func (s *DeploymentDraft) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DeploymentDraft) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *DeploymentDraft) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *DeploymentDraft) GetLocalVariables() []*LocalVariable {
	return s.LocalVariables
}

func (s *DeploymentDraft) GetLock() *Lock {
	return s.Lock
}

func (s *DeploymentDraft) GetModifiedAt() *int64 {
	return s.ModifiedAt
}

func (s *DeploymentDraft) GetModifier() *string {
	return s.Modifier
}

func (s *DeploymentDraft) GetModifierName() *string {
	return s.ModifierName
}

func (s *DeploymentDraft) GetName() *string {
	return s.Name
}

func (s *DeploymentDraft) GetNamespace() *string {
	return s.Namespace
}

func (s *DeploymentDraft) GetParentId() *string {
	return s.ParentId
}

func (s *DeploymentDraft) GetReferencedDeploymentId() *string {
	return s.ReferencedDeploymentId
}

func (s *DeploymentDraft) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeploymentDraft) SetArtifact(v *Artifact) *DeploymentDraft {
	s.Artifact = v
	return s
}

func (s *DeploymentDraft) SetCreatedAt(v int64) *DeploymentDraft {
	s.CreatedAt = &v
	return s
}

func (s *DeploymentDraft) SetCreator(v string) *DeploymentDraft {
	s.Creator = &v
	return s
}

func (s *DeploymentDraft) SetCreatorName(v string) *DeploymentDraft {
	s.CreatorName = &v
	return s
}

func (s *DeploymentDraft) SetDeploymentDraftId(v string) *DeploymentDraft {
	s.DeploymentDraftId = &v
	return s
}

func (s *DeploymentDraft) SetEngineVersion(v string) *DeploymentDraft {
	s.EngineVersion = &v
	return s
}

func (s *DeploymentDraft) SetExecutionMode(v string) *DeploymentDraft {
	s.ExecutionMode = &v
	return s
}

func (s *DeploymentDraft) SetLabels(v map[string]interface{}) *DeploymentDraft {
	s.Labels = v
	return s
}

func (s *DeploymentDraft) SetLocalVariables(v []*LocalVariable) *DeploymentDraft {
	s.LocalVariables = v
	return s
}

func (s *DeploymentDraft) SetLock(v *Lock) *DeploymentDraft {
	s.Lock = v
	return s
}

func (s *DeploymentDraft) SetModifiedAt(v int64) *DeploymentDraft {
	s.ModifiedAt = &v
	return s
}

func (s *DeploymentDraft) SetModifier(v string) *DeploymentDraft {
	s.Modifier = &v
	return s
}

func (s *DeploymentDraft) SetModifierName(v string) *DeploymentDraft {
	s.ModifierName = &v
	return s
}

func (s *DeploymentDraft) SetName(v string) *DeploymentDraft {
	s.Name = &v
	return s
}

func (s *DeploymentDraft) SetNamespace(v string) *DeploymentDraft {
	s.Namespace = &v
	return s
}

func (s *DeploymentDraft) SetParentId(v string) *DeploymentDraft {
	s.ParentId = &v
	return s
}

func (s *DeploymentDraft) SetReferencedDeploymentId(v string) *DeploymentDraft {
	s.ReferencedDeploymentId = &v
	return s
}

func (s *DeploymentDraft) SetWorkspace(v string) *DeploymentDraft {
	s.Workspace = &v
	return s
}

func (s *DeploymentDraft) Validate() error {
	if s.Artifact != nil {
		if err := s.Artifact.Validate(); err != nil {
			return err
		}
	}
	if s.LocalVariables != nil {
		for _, item := range s.LocalVariables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Lock != nil {
		if err := s.Lock.Validate(); err != nil {
			return err
		}
	}
	return nil
}
