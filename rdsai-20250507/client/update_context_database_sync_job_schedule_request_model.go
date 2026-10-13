// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncJobScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateContextDatabaseSyncJobScheduleRequest
	GetClientToken() *string
	SetDailyTimes(v []*string) *UpdateContextDatabaseSyncJobScheduleRequest
	GetDailyTimes() []*string
	SetIntervalMinutes(v int32) *UpdateContextDatabaseSyncJobScheduleRequest
	GetIntervalMinutes() *int32
	SetJobId(v string) *UpdateContextDatabaseSyncJobScheduleRequest
	GetJobId() *string
	SetMissedPolicy(v string) *UpdateContextDatabaseSyncJobScheduleRequest
	GetMissedPolicy() *string
	SetType(v string) *UpdateContextDatabaseSyncJobScheduleRequest
	GetType() *string
	SetWorkspaceId(v string) *UpdateContextDatabaseSyncJobScheduleRequest
	GetWorkspaceId() *string
}

type UpdateContextDatabaseSyncJobScheduleRequest struct {
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// ["09:00", "18:00"]
	DailyTimes []*string `json:"DailyTimes,omitempty" xml:"DailyTimes,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	IntervalMinutes *int32 `json:"IntervalMinutes,omitempty" xml:"IntervalMinutes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// skip
	MissedPolicy *string `json:"MissedPolicy,omitempty" xml:"MissedPolicy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// daily-times
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateContextDatabaseSyncJobScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncJobScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) GetDailyTimes() []*string {
	return s.DailyTimes
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) GetIntervalMinutes() *int32 {
	return s.IntervalMinutes
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) GetMissedPolicy() *string {
	return s.MissedPolicy
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) GetType() *string {
	return s.Type
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) SetClientToken(v string) *UpdateContextDatabaseSyncJobScheduleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) SetDailyTimes(v []*string) *UpdateContextDatabaseSyncJobScheduleRequest {
	s.DailyTimes = v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) SetIntervalMinutes(v int32) *UpdateContextDatabaseSyncJobScheduleRequest {
	s.IntervalMinutes = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) SetJobId(v string) *UpdateContextDatabaseSyncJobScheduleRequest {
	s.JobId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) SetMissedPolicy(v string) *UpdateContextDatabaseSyncJobScheduleRequest {
	s.MissedPolicy = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) SetType(v string) *UpdateContextDatabaseSyncJobScheduleRequest {
	s.Type = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) SetWorkspaceId(v string) *UpdateContextDatabaseSyncJobScheduleRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleRequest) Validate() error {
	return dara.Validate(s)
}
