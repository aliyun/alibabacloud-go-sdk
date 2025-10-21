// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJob interface {
	dara.Model
	String() string
	GoString() string
	SetArtifact(v *Artifact) *Job
	GetArtifact() *Artifact
	SetBatchResourceSetting(v *BatchResourceSetting) *Job
	GetBatchResourceSetting() *BatchResourceSetting
	SetCreatedAt(v string) *Job
	GetCreatedAt() *string
	SetCreator(v string) *Job
	GetCreator() *string
	SetCreatorName(v string) *Job
	GetCreatorName() *string
	SetDeploymentId(v string) *Job
	GetDeploymentId() *string
	SetDeploymentName(v string) *Job
	GetDeploymentName() *string
	SetEndTime(v int64) *Job
	GetEndTime() *int64
	SetEngineVersion(v string) *Job
	GetEngineVersion() *string
	SetExecutionMode(v string) *Job
	GetExecutionMode() *string
	SetFlinkConf(v map[string]interface{}) *Job
	GetFlinkConf() map[string]interface{}
	SetJobId(v string) *Job
	GetJobId() *string
	SetLocalVariables(v []*LocalVariable) *Job
	GetLocalVariables() []*LocalVariable
	SetLogging(v *Logging) *Job
	GetLogging() *Logging
	SetMetric(v *JobMetric) *Job
	GetMetric() *JobMetric
	SetModifiedAt(v string) *Job
	GetModifiedAt() *string
	SetModifier(v string) *Job
	GetModifier() *string
	SetModifierName(v string) *Job
	GetModifierName() *string
	SetNamespace(v string) *Job
	GetNamespace() *string
	SetRestoreStrategy(v *DeploymentRestoreStrategy) *Job
	GetRestoreStrategy() *DeploymentRestoreStrategy
	SetSessionClusterName(v string) *Job
	GetSessionClusterName() *string
	SetStartTime(v int64) *Job
	GetStartTime() *int64
	SetStatus(v *JobStatus) *Job
	GetStatus() *JobStatus
	SetStreamingResourceSetting(v *StreamingResourceSetting) *Job
	GetStreamingResourceSetting() *StreamingResourceSetting
	SetUserFlinkConf(v map[string]interface{}) *Job
	GetUserFlinkConf() map[string]interface{}
	SetWorkspace(v string) *Job
	GetWorkspace() *string
}

type Job struct {
	Artifact             *Artifact             `json:"artifact,omitempty" xml:"artifact,omitempty"`
	BatchResourceSetting *BatchResourceSetting `json:"batchResourceSetting,omitempty" xml:"batchResourceSetting,omitempty"`
	CreatedAt            *string               `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
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
	// 354dde66-a3ae-463e-967a-0b4107fd****
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// example:
	//
	// flinktest
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// example:
	//
	// 1660277235
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// vvr-4.0.14-flink-1.13
	EngineVersion *string `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	// example:
	//
	// BATCH
	ExecutionMode *string                `json:"executionMode,omitempty" xml:"executionMode,omitempty"`
	FlinkConf     map[string]interface{} `json:"flinkConf,omitempty" xml:"flinkConf,omitempty"`
	// example:
	//
	// 354dde66-a3ae-463e-967a-0b4107fd****
	JobId          *string          `json:"jobId,omitempty" xml:"jobId,omitempty"`
	LocalVariables []*LocalVariable `json:"localVariables,omitempty" xml:"localVariables,omitempty" type:"Repeated"`
	Logging        *Logging         `json:"logging,omitempty" xml:"logging,omitempty"`
	Metric         *JobMetric       `json:"metric,omitempty" xml:"metric,omitempty"`
	ModifiedAt     *string          `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
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
	// namespacetest
	Namespace       *string                    `json:"namespace,omitempty" xml:"namespace,omitempty"`
	RestoreStrategy *DeploymentRestoreStrategy `json:"restoreStrategy,omitempty" xml:"restoreStrategy,omitempty"`
	// example:
	//
	// preview
	SessionClusterName *string `json:"sessionClusterName,omitempty" xml:"sessionClusterName,omitempty"`
	// example:
	//
	// 1660190835
	StartTime                *int64                    `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status                   *JobStatus                `json:"status,omitempty" xml:"status,omitempty"`
	StreamingResourceSetting *StreamingResourceSetting `json:"streamingResourceSetting,omitempty" xml:"streamingResourceSetting,omitempty"`
	UserFlinkConf            map[string]interface{}    `json:"userFlinkConf,omitempty" xml:"userFlinkConf,omitempty"`
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s Job) String() string {
	return dara.Prettify(s)
}

func (s Job) GoString() string {
	return s.String()
}

func (s *Job) GetArtifact() *Artifact {
	return s.Artifact
}

func (s *Job) GetBatchResourceSetting() *BatchResourceSetting {
	return s.BatchResourceSetting
}

func (s *Job) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Job) GetCreator() *string {
	return s.Creator
}

func (s *Job) GetCreatorName() *string {
	return s.CreatorName
}

func (s *Job) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *Job) GetDeploymentName() *string {
	return s.DeploymentName
}

func (s *Job) GetEndTime() *int64 {
	return s.EndTime
}

func (s *Job) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *Job) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *Job) GetFlinkConf() map[string]interface{} {
	return s.FlinkConf
}

func (s *Job) GetJobId() *string {
	return s.JobId
}

func (s *Job) GetLocalVariables() []*LocalVariable {
	return s.LocalVariables
}

func (s *Job) GetLogging() *Logging {
	return s.Logging
}

func (s *Job) GetMetric() *JobMetric {
	return s.Metric
}

func (s *Job) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *Job) GetModifier() *string {
	return s.Modifier
}

func (s *Job) GetModifierName() *string {
	return s.ModifierName
}

func (s *Job) GetNamespace() *string {
	return s.Namespace
}

func (s *Job) GetRestoreStrategy() *DeploymentRestoreStrategy {
	return s.RestoreStrategy
}

func (s *Job) GetSessionClusterName() *string {
	return s.SessionClusterName
}

func (s *Job) GetStartTime() *int64 {
	return s.StartTime
}

func (s *Job) GetStatus() *JobStatus {
	return s.Status
}

func (s *Job) GetStreamingResourceSetting() *StreamingResourceSetting {
	return s.StreamingResourceSetting
}

func (s *Job) GetUserFlinkConf() map[string]interface{} {
	return s.UserFlinkConf
}

func (s *Job) GetWorkspace() *string {
	return s.Workspace
}

func (s *Job) SetArtifact(v *Artifact) *Job {
	s.Artifact = v
	return s
}

func (s *Job) SetBatchResourceSetting(v *BatchResourceSetting) *Job {
	s.BatchResourceSetting = v
	return s
}

func (s *Job) SetCreatedAt(v string) *Job {
	s.CreatedAt = &v
	return s
}

func (s *Job) SetCreator(v string) *Job {
	s.Creator = &v
	return s
}

func (s *Job) SetCreatorName(v string) *Job {
	s.CreatorName = &v
	return s
}

func (s *Job) SetDeploymentId(v string) *Job {
	s.DeploymentId = &v
	return s
}

func (s *Job) SetDeploymentName(v string) *Job {
	s.DeploymentName = &v
	return s
}

func (s *Job) SetEndTime(v int64) *Job {
	s.EndTime = &v
	return s
}

func (s *Job) SetEngineVersion(v string) *Job {
	s.EngineVersion = &v
	return s
}

func (s *Job) SetExecutionMode(v string) *Job {
	s.ExecutionMode = &v
	return s
}

func (s *Job) SetFlinkConf(v map[string]interface{}) *Job {
	s.FlinkConf = v
	return s
}

func (s *Job) SetJobId(v string) *Job {
	s.JobId = &v
	return s
}

func (s *Job) SetLocalVariables(v []*LocalVariable) *Job {
	s.LocalVariables = v
	return s
}

func (s *Job) SetLogging(v *Logging) *Job {
	s.Logging = v
	return s
}

func (s *Job) SetMetric(v *JobMetric) *Job {
	s.Metric = v
	return s
}

func (s *Job) SetModifiedAt(v string) *Job {
	s.ModifiedAt = &v
	return s
}

func (s *Job) SetModifier(v string) *Job {
	s.Modifier = &v
	return s
}

func (s *Job) SetModifierName(v string) *Job {
	s.ModifierName = &v
	return s
}

func (s *Job) SetNamespace(v string) *Job {
	s.Namespace = &v
	return s
}

func (s *Job) SetRestoreStrategy(v *DeploymentRestoreStrategy) *Job {
	s.RestoreStrategy = v
	return s
}

func (s *Job) SetSessionClusterName(v string) *Job {
	s.SessionClusterName = &v
	return s
}

func (s *Job) SetStartTime(v int64) *Job {
	s.StartTime = &v
	return s
}

func (s *Job) SetStatus(v *JobStatus) *Job {
	s.Status = v
	return s
}

func (s *Job) SetStreamingResourceSetting(v *StreamingResourceSetting) *Job {
	s.StreamingResourceSetting = v
	return s
}

func (s *Job) SetUserFlinkConf(v map[string]interface{}) *Job {
	s.UserFlinkConf = v
	return s
}

func (s *Job) SetWorkspace(v string) *Job {
	s.Workspace = &v
	return s
}

func (s *Job) Validate() error {
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
	if s.Metric != nil {
		if err := s.Metric.Validate(); err != nil {
			return err
		}
	}
	if s.RestoreStrategy != nil {
		if err := s.RestoreStrategy.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
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
