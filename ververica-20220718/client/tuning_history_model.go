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
	ActionType         *string                          `json:"actionType,omitempty" xml:"actionType,omitempty"`
	Annotations        map[string]*string               `json:"annotations,omitempty" xml:"annotations,omitempty"`
	DeploymentName     *string                          `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	IsHotUpdate        *bool                            `json:"isHotUpdate,omitempty" xml:"isHotUpdate,omitempty"`
	JobId              *string                          `json:"jobId,omitempty" xml:"jobId,omitempty"`
	NewResourceSetting *TuningHistoryNewResourceSetting `json:"newResourceSetting,omitempty" xml:"newResourceSetting,omitempty" type:"Struct"`
	OldResourceSetting *TuningHistoryOldResourceSetting `json:"oldResourceSetting,omitempty" xml:"oldResourceSetting,omitempty" type:"Struct"`
	TriggerTime        *int64                           `json:"triggerTime,omitempty" xml:"triggerTime,omitempty"`
	TuningId           *string                          `json:"tuningId,omitempty" xml:"tuningId,omitempty"`
	TuningMessage      *string                          `json:"tuningMessage,omitempty" xml:"tuningMessage,omitempty"`
	TuningState        *string                          `json:"tuningState,omitempty" xml:"tuningState,omitempty"`
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
	Cpu         *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Memory      *string  `json:"memory,omitempty" xml:"memory,omitempty"`
	Parallelism *int32   `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
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
	Cpu         *float64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Memory      *string  `json:"memory,omitempty" xml:"memory,omitempty"`
	Parallelism *int32   `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
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
