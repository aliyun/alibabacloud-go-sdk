// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTuningHistory interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *TuningHistory
	GetActionType() *string
	SetAnnotations(v map[string]*string) *TuningHistory
	GetAnnotations() map[string]*string
	SetDeploymentName(v string) *TuningHistory
	GetDeploymentName() *string
	SetIsHotUpdate(v bool) *TuningHistory
	GetIsHotUpdate() *bool
	SetJobId(v string) *TuningHistory
	GetJobId() *string
	SetNewResourceSetting(v *TuningHistoryNewResourceSetting) *TuningHistory
	GetNewResourceSetting() *TuningHistoryNewResourceSetting
	SetOldResourceSetting(v *TuningHistoryOldResourceSetting) *TuningHistory
	GetOldResourceSetting() *TuningHistoryOldResourceSetting
	SetTriggerTime(v int64) *TuningHistory
	GetTriggerTime() *int64
	SetTuningId(v string) *TuningHistory
	GetTuningId() *string
	SetTuningMessage(v string) *TuningHistory
	GetTuningMessage() *string
	SetTuningState(v string) *TuningHistory
	GetTuningState() *string
}

type TuningHistory struct {
	// The action type. Valid values:
	//
	// - SCALE_UP_PARALLELISM: scales up parallelism.
	//
	// - SCALE_DOWN_PARALLELISM: scales down parallelism.
	//
	// - SCALE_UP_MEMORY: scales up memory.
	//
	// - RESTART: restarts the job.
	//
	// example:
	//
	// SCALE_UP_PARALLELISM
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// The additional annotations.
	Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
	// The full path name of the deployment.
	//
	// example:
	//
	// namespaces/ns-xxx/deployments/6aa0d4d1-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// Indicates whether this is a hot update. A value of true indicates that the change takes effect without restarting the job. A value of false indicates that the job must be restarted.
	//
	// example:
	//
	// true
	IsHotUpdate *bool `json:"isHotUpdate,omitempty" xml:"isHotUpdate,omitempty"`
	// The ID of the associated job.
	//
	// example:
	//
	// b462c053-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The resource configuration after tuning. This value may be null if the tuning failed.
	NewResourceSetting *TuningHistoryNewResourceSetting `json:"newResourceSetting,omitempty" xml:"newResourceSetting,omitempty" type:"Struct"`
	// The resource configuration before tuning.
	OldResourceSetting *TuningHistoryOldResourceSetting `json:"oldResourceSetting,omitempty" xml:"oldResourceSetting,omitempty" type:"Struct"`
	// The trigger timestamp in milliseconds.
	//
	// example:
	//
	// 1718270936000
	TriggerTime *int64 `json:"triggerTime,omitempty" xml:"triggerTime,omitempty"`
	// The UUID of the tuning record.
	//
	// example:
	//
	// 06d81ae2-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	TuningId *string `json:"tuningId,omitempty" xml:"tuningId,omitempty"`
	// The tuning message. This is an internationalized, human-readable string that is not recommended for programmatic parsing.
	//
	// example:
	//
	// Scale up parallelism from 2 to 4
	TuningMessage *string `json:"tuningMessage,omitempty" xml:"tuningMessage,omitempty"`
	// The tuning state. Valid values:
	//
	// - SUCCESS: The tuning succeeded.
	//
	// - FAILED: The tuning failed.
	//
	// - EXECUTING: The tuning is in progress.
	//
	// - TERMINATED: The tuning was terminated.
	//
	// - FAILED_WITH_ROLLBACK_SUCCESS: The tuning failed but the rollback succeeded.
	//
	// - FAILED_WITH_ROLLBACK_FAILED: The tuning failed and the rollback also failed.
	//
	// - FAILED_WITH_RESOURCE_LACK: The tuning failed due to insufficient resources.
	//
	// - FAILED_WITH_SAME_RESOURCE_SETTING: The tuning failed because the resource configuration did not change.
	//
	// example:
	//
	// SUCCESS
	TuningState *string `json:"tuningState,omitempty" xml:"tuningState,omitempty"`
}

func (s TuningHistory) String() string {
	return dara.Prettify(s)
}

func (s TuningHistory) GoString() string {
	return s.String()
}

func (s *TuningHistory) GetActionType() *string {
	return s.ActionType
}

func (s *TuningHistory) GetAnnotations() map[string]*string {
	return s.Annotations
}

func (s *TuningHistory) GetDeploymentName() *string {
	return s.DeploymentName
}

func (s *TuningHistory) GetIsHotUpdate() *bool {
	return s.IsHotUpdate
}

func (s *TuningHistory) GetJobId() *string {
	return s.JobId
}

func (s *TuningHistory) GetNewResourceSetting() *TuningHistoryNewResourceSetting {
	return s.NewResourceSetting
}

func (s *TuningHistory) GetOldResourceSetting() *TuningHistoryOldResourceSetting {
	return s.OldResourceSetting
}

func (s *TuningHistory) GetTriggerTime() *int64 {
	return s.TriggerTime
}

func (s *TuningHistory) GetTuningId() *string {
	return s.TuningId
}

func (s *TuningHistory) GetTuningMessage() *string {
	return s.TuningMessage
}

func (s *TuningHistory) GetTuningState() *string {
	return s.TuningState
}

func (s *TuningHistory) SetActionType(v string) *TuningHistory {
	s.ActionType = &v
	return s
}

func (s *TuningHistory) SetAnnotations(v map[string]*string) *TuningHistory {
	s.Annotations = v
	return s
}

func (s *TuningHistory) SetDeploymentName(v string) *TuningHistory {
	s.DeploymentName = &v
	return s
}

func (s *TuningHistory) SetIsHotUpdate(v bool) *TuningHistory {
	s.IsHotUpdate = &v
	return s
}

func (s *TuningHistory) SetJobId(v string) *TuningHistory {
	s.JobId = &v
	return s
}

func (s *TuningHistory) SetNewResourceSetting(v *TuningHistoryNewResourceSetting) *TuningHistory {
	s.NewResourceSetting = v
	return s
}

func (s *TuningHistory) SetOldResourceSetting(v *TuningHistoryOldResourceSetting) *TuningHistory {
	s.OldResourceSetting = v
	return s
}

func (s *TuningHistory) SetTriggerTime(v int64) *TuningHistory {
	s.TriggerTime = &v
	return s
}

func (s *TuningHistory) SetTuningId(v string) *TuningHistory {
	s.TuningId = &v
	return s
}

func (s *TuningHistory) SetTuningMessage(v string) *TuningHistory {
	s.TuningMessage = &v
	return s
}

func (s *TuningHistory) SetTuningState(v string) *TuningHistory {
	s.TuningState = &v
	return s
}

func (s *TuningHistory) Validate() error {
	if s.NewResourceSetting != nil {
		if err := s.NewResourceSetting.Validate(); err != nil {
			return err
		}
	}
	if s.OldResourceSetting != nil {
		if err := s.OldResourceSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TuningHistoryNewResourceSetting struct {
	// The number of CPU cores per TaskManager.
	//
	// example:
	//
	// 1.0
	Cpu *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The memory per TaskManager, in a format such as 4 Gi.
	//
	// example:
	//
	// 2 Gi
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// The parallelism.
	//
	// example:
	//
	// 4
	Parallelism *int32 `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
}

func (s TuningHistoryNewResourceSetting) String() string {
	return dara.Prettify(s)
}

func (s TuningHistoryNewResourceSetting) GoString() string {
	return s.String()
}

func (s *TuningHistoryNewResourceSetting) GetCpu() *float64 {
	return s.Cpu
}

func (s *TuningHistoryNewResourceSetting) GetMemory() *string {
	return s.Memory
}

func (s *TuningHistoryNewResourceSetting) GetParallelism() *int32 {
	return s.Parallelism
}

func (s *TuningHistoryNewResourceSetting) SetCpu(v float64) *TuningHistoryNewResourceSetting {
	s.Cpu = &v
	return s
}

func (s *TuningHistoryNewResourceSetting) SetMemory(v string) *TuningHistoryNewResourceSetting {
	s.Memory = &v
	return s
}

func (s *TuningHistoryNewResourceSetting) SetParallelism(v int32) *TuningHistoryNewResourceSetting {
	s.Parallelism = &v
	return s
}

func (s *TuningHistoryNewResourceSetting) Validate() error {
	return dara.Validate(s)
}

type TuningHistoryOldResourceSetting struct {
	// The number of CPU cores per TaskManager.
	//
	// example:
	//
	// 1.0
	Cpu *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The memory per TaskManager, in a format such as 4 Gi.
	//
	// example:
	//
	// 2 Gi
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// The parallelism.
	//
	// example:
	//
	// 2
	Parallelism *int32 `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
}

func (s TuningHistoryOldResourceSetting) String() string {
	return dara.Prettify(s)
}

func (s TuningHistoryOldResourceSetting) GoString() string {
	return s.String()
}

func (s *TuningHistoryOldResourceSetting) GetCpu() *float64 {
	return s.Cpu
}

func (s *TuningHistoryOldResourceSetting) GetMemory() *string {
	return s.Memory
}

func (s *TuningHistoryOldResourceSetting) GetParallelism() *int32 {
	return s.Parallelism
}

func (s *TuningHistoryOldResourceSetting) SetCpu(v float64) *TuningHistoryOldResourceSetting {
	s.Cpu = &v
	return s
}

func (s *TuningHistoryOldResourceSetting) SetMemory(v string) *TuningHistoryOldResourceSetting {
	s.Memory = &v
	return s
}

func (s *TuningHistoryOldResourceSetting) SetParallelism(v int32) *TuningHistoryOldResourceSetting {
	s.Parallelism = &v
	return s
}

func (s *TuningHistoryOldResourceSetting) Validate() error {
	return dara.Validate(s)
}
