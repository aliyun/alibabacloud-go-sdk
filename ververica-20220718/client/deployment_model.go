// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployment interface {
	dara.Model
	String() string
	GoString() string
	SetArtifact(v *Artifact) *Deployment
	GetArtifact() *Artifact
	SetBatchResourceSetting(v *BatchResourceSetting) *Deployment
	GetBatchResourceSetting() *BatchResourceSetting
	SetCreatedAt(v string) *Deployment
	GetCreatedAt() *string
	SetCreator(v string) *Deployment
	GetCreator() *string
	SetCreatorName(v string) *Deployment
	GetCreatorName() *string
	SetDeploymentHasChanged(v bool) *Deployment
	GetDeploymentHasChanged() *bool
	SetDeploymentId(v string) *Deployment
	GetDeploymentId() *string
	SetDeploymentTarget(v *BriefDeploymentTarget) *Deployment
	GetDeploymentTarget() *BriefDeploymentTarget
	SetDescription(v string) *Deployment
	GetDescription() *string
	SetEngineVersion(v string) *Deployment
	GetEngineVersion() *string
	SetExecutionMode(v string) *Deployment
	GetExecutionMode() *string
	SetFlinkConf(v map[string]interface{}) *Deployment
	GetFlinkConf() map[string]interface{}
	SetJobSummary(v *JobSummary) *Deployment
	GetJobSummary() *JobSummary
	SetLabels(v map[string]interface{}) *Deployment
	GetLabels() map[string]interface{}
	SetLocalVariables(v []*LocalVariable) *Deployment
	GetLocalVariables() []*LocalVariable
	SetLogging(v *Logging) *Deployment
	GetLogging() *Logging
	SetModifiedAt(v string) *Deployment
	GetModifiedAt() *string
	SetModifier(v string) *Deployment
	GetModifier() *string
	SetModifierName(v string) *Deployment
	GetModifierName() *string
	SetName(v string) *Deployment
	GetName() *string
	SetNamespace(v string) *Deployment
	GetNamespace() *string
	SetReferencedDeploymentDraftId(v string) *Deployment
	GetReferencedDeploymentDraftId() *string
	SetStreamingResourceSetting(v *StreamingResourceSetting) *Deployment
	GetStreamingResourceSetting() *StreamingResourceSetting
	SetWorkspace(v string) *Deployment
	GetWorkspace() *string
}

type Deployment struct {
	// The parameters that are required for starting a deployment.
	Artifact *Artifact `json:"artifact,omitempty" xml:"artifact,omitempty"`
	// The resource configuration of the batch deployment.
	BatchResourceSetting *BatchResourceSetting `json:"batchResourceSetting,omitempty" xml:"batchResourceSetting,omitempty"`
	// The time at which the deployment was created.
	//
	// example:
	//
	// 1714058507
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The ID of the account that is used to create the deployment.
	//
	// example:
	//
	// 27846363877456****
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The name of the account that is used to create the deployment.
	//
	// example:
	//
	// ****@streamcompute.onaliyun.com
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// Specifies whether the deployment is modified after the deployment is started.
	//
	// example:
	//
	// true
	DeploymentHasChanged *bool `json:"deploymentHasChanged,omitempty" xml:"deploymentHasChanged,omitempty"`
	// The ID of the deployment.
	//
	// example:
	//
	// 00000000-0000-0000-0000-0000012312****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// The cluster on which the deployment is deployed.
	DeploymentTarget *BriefDeploymentTarget `json:"deploymentTarget,omitempty" xml:"deploymentTarget,omitempty"`
	// The description of the deployment.
	//
	// example:
	//
	// this is a deployment description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The engine version of the deployment.
	//
	// example:
	//
	// vvr-6.0.0-flink-1.15
	EngineVersion *string `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	// The execution mode of the deployment. Valid values:
	//
	// 	- STREAMING
	//
	// 	- BATCH
	//
	// example:
	//
	// STREAMING
	ExecutionMode *string `json:"executionMode,omitempty" xml:"executionMode,omitempty"`
	// The Realtime Compute for Apache Flink configuration.
	//
	// example:
	//
	// {"taskmanager.numberOfTaskSlots":"1"}
	FlinkConf map[string]interface{} `json:"flinkConf,omitempty" xml:"flinkConf,omitempty"`
	// The summary of jobs in the deployment.
	JobSummary *JobSummary            `json:"jobSummary,omitempty" xml:"jobSummary,omitempty"`
	Labels     map[string]interface{} `json:"labels,omitempty" xml:"labels,omitempty"`
	// The variables of the deployment.
	LocalVariables []*LocalVariable `json:"localVariables,omitempty" xml:"localVariables,omitempty" type:"Repeated"`
	// The logging configuration.
	Logging *Logging `json:"logging,omitempty" xml:"logging,omitempty"`
	// The time at which the deployment was modified.
	//
	// example:
	//
	// 1714058843
	ModifiedAt *string `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	// The ID of the account that is used to modify the deployment.
	//
	// example:
	//
	// 27846363877456****
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// The name of the account that is used to modify the deployment.
	//
	// example:
	//
	// ****@streamcompute.onaliyun.com
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// The name of the deployment.
	//
	// example:
	//
	// deploymentName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 00000000-0000-0000-0000-000000000003
	ReferencedDeploymentDraftId *string `json:"referencedDeploymentDraftId,omitempty" xml:"referencedDeploymentDraftId,omitempty"`
	// The resource configuration of the streaming deployment.
	StreamingResourceSetting *StreamingResourceSetting `json:"streamingResourceSetting,omitempty" xml:"streamingResourceSetting,omitempty"`
	// The workspace to which the deployment belongs.
	//
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s Deployment) String() string {
	return dara.Prettify(s)
}

func (s Deployment) GoString() string {
	return s.String()
}

func (s *Deployment) GetArtifact() *Artifact {
	return s.Artifact
}

func (s *Deployment) GetBatchResourceSetting() *BatchResourceSetting {
	return s.BatchResourceSetting
}

func (s *Deployment) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Deployment) GetCreator() *string {
	return s.Creator
}

func (s *Deployment) GetCreatorName() *string {
	return s.CreatorName
}

func (s *Deployment) GetDeploymentHasChanged() *bool {
	return s.DeploymentHasChanged
}

func (s *Deployment) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *Deployment) GetDeploymentTarget() *BriefDeploymentTarget {
	return s.DeploymentTarget
}

func (s *Deployment) GetDescription() *string {
	return s.Description
}

func (s *Deployment) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *Deployment) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *Deployment) GetFlinkConf() map[string]interface{} {
	return s.FlinkConf
}

func (s *Deployment) GetJobSummary() *JobSummary {
	return s.JobSummary
}

func (s *Deployment) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *Deployment) GetLocalVariables() []*LocalVariable {
	return s.LocalVariables
}

func (s *Deployment) GetLogging() *Logging {
	return s.Logging
}

func (s *Deployment) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *Deployment) GetModifier() *string {
	return s.Modifier
}

func (s *Deployment) GetModifierName() *string {
	return s.ModifierName
}

func (s *Deployment) GetName() *string {
	return s.Name
}

func (s *Deployment) GetNamespace() *string {
	return s.Namespace
}

func (s *Deployment) GetReferencedDeploymentDraftId() *string {
	return s.ReferencedDeploymentDraftId
}

func (s *Deployment) GetStreamingResourceSetting() *StreamingResourceSetting {
	return s.StreamingResourceSetting
}

func (s *Deployment) GetWorkspace() *string {
	return s.Workspace
}

func (s *Deployment) SetArtifact(v *Artifact) *Deployment {
	s.Artifact = v
	return s
}

func (s *Deployment) SetBatchResourceSetting(v *BatchResourceSetting) *Deployment {
	s.BatchResourceSetting = v
	return s
}

func (s *Deployment) SetCreatedAt(v string) *Deployment {
	s.CreatedAt = &v
	return s
}

func (s *Deployment) SetCreator(v string) *Deployment {
	s.Creator = &v
	return s
}

func (s *Deployment) SetCreatorName(v string) *Deployment {
	s.CreatorName = &v
	return s
}

func (s *Deployment) SetDeploymentHasChanged(v bool) *Deployment {
	s.DeploymentHasChanged = &v
	return s
}

func (s *Deployment) SetDeploymentId(v string) *Deployment {
	s.DeploymentId = &v
	return s
}

func (s *Deployment) SetDeploymentTarget(v *BriefDeploymentTarget) *Deployment {
	s.DeploymentTarget = v
	return s
}

func (s *Deployment) SetDescription(v string) *Deployment {
	s.Description = &v
	return s
}

func (s *Deployment) SetEngineVersion(v string) *Deployment {
	s.EngineVersion = &v
	return s
}

func (s *Deployment) SetExecutionMode(v string) *Deployment {
	s.ExecutionMode = &v
	return s
}

func (s *Deployment) SetFlinkConf(v map[string]interface{}) *Deployment {
	s.FlinkConf = v
	return s
}

func (s *Deployment) SetJobSummary(v *JobSummary) *Deployment {
	s.JobSummary = v
	return s
}

func (s *Deployment) SetLabels(v map[string]interface{}) *Deployment {
	s.Labels = v
	return s
}

func (s *Deployment) SetLocalVariables(v []*LocalVariable) *Deployment {
	s.LocalVariables = v
	return s
}

func (s *Deployment) SetLogging(v *Logging) *Deployment {
	s.Logging = v
	return s
}

func (s *Deployment) SetModifiedAt(v string) *Deployment {
	s.ModifiedAt = &v
	return s
}

func (s *Deployment) SetModifier(v string) *Deployment {
	s.Modifier = &v
	return s
}

func (s *Deployment) SetModifierName(v string) *Deployment {
	s.ModifierName = &v
	return s
}

func (s *Deployment) SetName(v string) *Deployment {
	s.Name = &v
	return s
}

func (s *Deployment) SetNamespace(v string) *Deployment {
	s.Namespace = &v
	return s
}

func (s *Deployment) SetReferencedDeploymentDraftId(v string) *Deployment {
	s.ReferencedDeploymentDraftId = &v
	return s
}

func (s *Deployment) SetStreamingResourceSetting(v *StreamingResourceSetting) *Deployment {
	s.StreamingResourceSetting = v
	return s
}

func (s *Deployment) SetWorkspace(v string) *Deployment {
	s.Workspace = &v
	return s
}

func (s *Deployment) Validate() error {
	if s.Artifact != nil {
		if err := s.Artifact.Validate(); err != nil {
			return err
		}
	}
	if s.BatchResourceSetting != nil {
		if err := s.BatchResourceSetting.Validate(); err != nil {
			return err
		}
	}
	if s.DeploymentTarget != nil {
		if err := s.DeploymentTarget.Validate(); err != nil {
			return err
		}
	}
	if s.JobSummary != nil {
		if err := s.JobSummary.Validate(); err != nil {
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
	if s.Logging != nil {
		if err := s.Logging.Validate(); err != nil {
			return err
		}
	}
	if s.StreamingResourceSetting != nil {
		if err := s.StreamingResourceSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}
