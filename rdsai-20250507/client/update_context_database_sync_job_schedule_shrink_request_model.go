// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncJobScheduleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest
	GetClientToken() *string
	SetDailyTimesShrink(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest
	GetDailyTimesShrink() *string
	SetIntervalMinutes(v int32) *UpdateContextDatabaseSyncJobScheduleShrinkRequest
	GetIntervalMinutes() *int32
	SetJobId(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest
	GetJobId() *string
	SetMissedPolicy(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest
	GetMissedPolicy() *string
	SetType(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest
	GetType() *string
	SetWorkspaceId(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest
	GetWorkspaceId() *string
}

type UpdateContextDatabaseSyncJobScheduleShrinkRequest struct {
	// example:
	//
	// sync-connection-20260923-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// ["09:00", "18:00"]
	DailyTimesShrink *string `json:"DailyTimes,omitempty" xml:"DailyTimes,omitempty"`
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

func (s UpdateContextDatabaseSyncJobScheduleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncJobScheduleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) GetDailyTimesShrink() *string {
	return s.DailyTimesShrink
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) GetIntervalMinutes() *int32 {
	return s.IntervalMinutes
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) GetMissedPolicy() *string {
	return s.MissedPolicy
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) GetType() *string {
	return s.Type
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) SetClientToken(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) SetDailyTimesShrink(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest {
	s.DailyTimesShrink = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) SetIntervalMinutes(v int32) *UpdateContextDatabaseSyncJobScheduleShrinkRequest {
	s.IntervalMinutes = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) SetJobId(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest {
	s.JobId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) SetMissedPolicy(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest {
	s.MissedPolicy = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) SetType(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest {
	s.Type = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) SetWorkspaceId(v string) *UpdateContextDatabaseSyncJobScheduleShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
