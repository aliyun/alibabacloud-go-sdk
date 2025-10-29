// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrationJob interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentState(v string) *MigrationJob
	GetCurrentState() *string
	SetDisableSourceClusterAuth(v bool) *MigrationJob
	GetDisableSourceClusterAuth() *bool
	SetDisableTargetClusterAuth(v bool) *MigrationJob
	GetDisableTargetClusterAuth() *bool
	SetEndTime(v int64) *MigrationJob
	GetEndTime() *int64
	SetMigrationJobId(v string) *MigrationJob
	GetMigrationJobId() *string
	SetPhase(v string) *MigrationJob
	GetPhase() *string
	SetSourceCluster(v *MigrationJobSourceCluster) *MigrationJob
	GetSourceCluster() *MigrationJobSourceCluster
	SetStartTime(v int64) *MigrationJob
	GetStartTime() *int64
	SetStatusResult(v []*MigrationJobStatusResult) *MigrationJob
	GetStatusResult() []*MigrationJobStatusResult
	SetTargetCluster(v *MigrationJobTargetCluster) *MigrationJob
	GetTargetCluster() *MigrationJobTargetCluster
	SetUpdateTime(v int64) *MigrationJob
	GetUpdateTime() *int64
}

type MigrationJob struct {
	CurrentState             *string                     `json:"currentState,omitempty" xml:"currentState,omitempty"`
	DisableSourceClusterAuth *bool                       `json:"disableSourceClusterAuth,omitempty" xml:"disableSourceClusterAuth,omitempty"`
	DisableTargetClusterAuth *bool                       `json:"disableTargetClusterAuth,omitempty" xml:"disableTargetClusterAuth,omitempty"`
	EndTime                  *int64                      `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MigrationJobId           *string                     `json:"migrationJobId,omitempty" xml:"migrationJobId,omitempty"`
	Phase                    *string                     `json:"phase,omitempty" xml:"phase,omitempty"`
	SourceCluster            *MigrationJobSourceCluster  `json:"sourceCluster,omitempty" xml:"sourceCluster,omitempty" type:"Struct"`
	StartTime                *int64                      `json:"startTime,omitempty" xml:"startTime,omitempty"`
	StatusResult             []*MigrationJobStatusResult `json:"statusResult,omitempty" xml:"statusResult,omitempty" type:"Repeated"`
	TargetCluster            *MigrationJobTargetCluster  `json:"targetCluster,omitempty" xml:"targetCluster,omitempty" type:"Struct"`
	UpdateTime               *int64                      `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s MigrationJob) String() string {
	return dara.Prettify(s)
}

func (s MigrationJob) GoString() string {
	return s.String()
}

func (s *MigrationJob) GetCurrentState() *string {
	return s.CurrentState
}

func (s *MigrationJob) GetDisableSourceClusterAuth() *bool {
	return s.DisableSourceClusterAuth
}

func (s *MigrationJob) GetDisableTargetClusterAuth() *bool {
	return s.DisableTargetClusterAuth
}

func (s *MigrationJob) GetEndTime() *int64 {
	return s.EndTime
}

func (s *MigrationJob) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *MigrationJob) GetPhase() *string {
	return s.Phase
}

func (s *MigrationJob) GetSourceCluster() *MigrationJobSourceCluster {
	return s.SourceCluster
}

func (s *MigrationJob) GetStartTime() *int64 {
	return s.StartTime
}

func (s *MigrationJob) GetStatusResult() []*MigrationJobStatusResult {
	return s.StatusResult
}

func (s *MigrationJob) GetTargetCluster() *MigrationJobTargetCluster {
	return s.TargetCluster
}

func (s *MigrationJob) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *MigrationJob) SetCurrentState(v string) *MigrationJob {
	s.CurrentState = &v
	return s
}

func (s *MigrationJob) SetDisableSourceClusterAuth(v bool) *MigrationJob {
	s.DisableSourceClusterAuth = &v
	return s
}

func (s *MigrationJob) SetDisableTargetClusterAuth(v bool) *MigrationJob {
	s.DisableTargetClusterAuth = &v
	return s
}

func (s *MigrationJob) SetEndTime(v int64) *MigrationJob {
	s.EndTime = &v
	return s
}

func (s *MigrationJob) SetMigrationJobId(v string) *MigrationJob {
	s.MigrationJobId = &v
	return s
}

func (s *MigrationJob) SetPhase(v string) *MigrationJob {
	s.Phase = &v
	return s
}

func (s *MigrationJob) SetSourceCluster(v *MigrationJobSourceCluster) *MigrationJob {
	s.SourceCluster = v
	return s
}

func (s *MigrationJob) SetStartTime(v int64) *MigrationJob {
	s.StartTime = &v
	return s
}

func (s *MigrationJob) SetStatusResult(v []*MigrationJobStatusResult) *MigrationJob {
	s.StatusResult = v
	return s
}

func (s *MigrationJob) SetTargetCluster(v *MigrationJobTargetCluster) *MigrationJob {
	s.TargetCluster = v
	return s
}

func (s *MigrationJob) SetUpdateTime(v int64) *MigrationJob {
	s.UpdateTime = &v
	return s
}

func (s *MigrationJob) Validate() error {
	if s.SourceCluster != nil {
		if err := s.SourceCluster.Validate(); err != nil {
			return err
		}
	}
	if s.StatusResult != nil {
		for _, item := range s.StatusResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TargetCluster != nil {
		if err := s.TargetCluster.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MigrationJobSourceCluster struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MigrationJobSourceCluster) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobSourceCluster) GoString() string {
	return s.String()
}

func (s *MigrationJobSourceCluster) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MigrationJobSourceCluster) GetType() *string {
	return s.Type
}

func (s *MigrationJobSourceCluster) SetInstanceId(v string) *MigrationJobSourceCluster {
	s.InstanceId = &v
	return s
}

func (s *MigrationJobSourceCluster) SetType(v string) *MigrationJobSourceCluster {
	s.Type = &v
	return s
}

func (s *MigrationJobSourceCluster) Validate() error {
	return dara.Validate(s)
}

type MigrationJobStatusResult struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s MigrationJobStatusResult) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobStatusResult) GoString() string {
	return s.String()
}

func (s *MigrationJobStatusResult) GetCode() *string {
	return s.Code
}

func (s *MigrationJobStatusResult) GetSuccess() *bool {
	return s.Success
}

func (s *MigrationJobStatusResult) SetCode(v string) *MigrationJobStatusResult {
	s.Code = &v
	return s
}

func (s *MigrationJobStatusResult) SetSuccess(v bool) *MigrationJobStatusResult {
	s.Success = &v
	return s
}

func (s *MigrationJobStatusResult) Validate() error {
	return dara.Validate(s)
}

type MigrationJobTargetCluster struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s MigrationJobTargetCluster) String() string {
	return dara.Prettify(s)
}

func (s MigrationJobTargetCluster) GoString() string {
	return s.String()
}

func (s *MigrationJobTargetCluster) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MigrationJobTargetCluster) GetType() *string {
	return s.Type
}

func (s *MigrationJobTargetCluster) SetInstanceId(v string) *MigrationJobTargetCluster {
	s.InstanceId = &v
	return s
}

func (s *MigrationJobTargetCluster) SetType(v string) *MigrationJobTargetCluster {
	s.Type = &v
	return s
}

func (s *MigrationJobTargetCluster) Validate() error {
	return dara.Validate(s)
}
