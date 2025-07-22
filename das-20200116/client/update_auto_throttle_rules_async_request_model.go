// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoThrottleRulesAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbnormalDuration(v float64) *UpdateAutoThrottleRulesAsyncRequest
	GetAbnormalDuration() *float64
	SetActiveSessions(v int64) *UpdateAutoThrottleRulesAsyncRequest
	GetActiveSessions() *int64
	SetAllowThrottleEndTime(v string) *UpdateAutoThrottleRulesAsyncRequest
	GetAllowThrottleEndTime() *string
	SetAllowThrottleStartTime(v string) *UpdateAutoThrottleRulesAsyncRequest
	GetAllowThrottleStartTime() *string
	SetAutoKillSession(v bool) *UpdateAutoThrottleRulesAsyncRequest
	GetAutoKillSession() *bool
	SetConsoleContext(v string) *UpdateAutoThrottleRulesAsyncRequest
	GetConsoleContext() *string
	SetCpuSessionRelation(v string) *UpdateAutoThrottleRulesAsyncRequest
	GetCpuSessionRelation() *string
	SetCpuUsage(v float64) *UpdateAutoThrottleRulesAsyncRequest
	GetCpuUsage() *float64
	SetInstanceIds(v string) *UpdateAutoThrottleRulesAsyncRequest
	GetInstanceIds() *string
	SetMaxThrottleTime(v float64) *UpdateAutoThrottleRulesAsyncRequest
	GetMaxThrottleTime() *float64
	SetResultId(v string) *UpdateAutoThrottleRulesAsyncRequest
	GetResultId() *string
}

type UpdateAutoThrottleRulesAsyncRequest struct {
	// The duration threshold for triggering automatic SQL throttling. Set this parameter to an integer that is greater than or equal to 2. Unit: minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	AbnormalDuration *float64 `json:"AbnormalDuration,omitempty" xml:"AbnormalDuration,omitempty"`
	// The threshold for the number of active sessions.
	//
	// 	- If this parameter and CpuUsage are in the **OR*	- relationship, set this parameter to an integer that is greater than or equal to 16.
	//
	// 	- If this parameter and CpuUsage are in the **AND*	- relationship, set this parameter to an integer that is greater than or equal to 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16
	ActiveSessions *int64 `json:"ActiveSessions,omitempty" xml:"ActiveSessions,omitempty"`
	// The end time of the throttling window. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23:59Z
	AllowThrottleEndTime *string `json:"AllowThrottleEndTime,omitempty" xml:"AllowThrottleEndTime,omitempty"`
	// The start time of the throttling window. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00:00Z
	AllowThrottleStartTime *string `json:"AllowThrottleStartTime,omitempty" xml:"AllowThrottleStartTime,omitempty"`
	// Specifies whether to terminate abnormal SQL statements in execution at the same time. Valid values:
	//
	// >  Abnormal SQL statements use the same template as the SQL statements to be throttled.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoKillSession *bool `json:"AutoKillSession,omitempty" xml:"AutoKillSession,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The logical relationship between the CPU utilization threshold and the maximum number of active sessions. Valid values:
	//
	// 	- **AND**
	//
	// 	- **OR**
	//
	// This parameter is required.
	//
	// example:
	//
	// OR
	CpuSessionRelation *string `json:"CpuSessionRelation,omitempty" xml:"CpuSessionRelation,omitempty"`
	// The threshold for CPU utilization. Valid values: 70% to 100%.
	//
	// This parameter is required.
	//
	// example:
	//
	// 70
	CpuUsage *float64 `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	// The database instance IDs.
	//
	// >  Set this parameter to a JSON array that consists of multiple instance IDs. Separate instance IDs with commas (,). Example: `[\\"Instance ID1\\", \\"Instance ID2\\"]`.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"rm-2ze8g2am97624****\\",\\"rm-2ze9xrhze0709****\\"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The maximum throttling duration. Set this parameter to a positive integer. Unit: minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxThrottleTime *float64 `json:"MaxThrottleTime,omitempty" xml:"MaxThrottleTime,omitempty"`
	// The ID of the asynchronous request.
	//
	// >  You can leave this parameter empty when you call the operation to initiate the request for the first time, and use the value of this parameter contained in the response to the first request for subsequent requests.
	//
	// example:
	//
	// async__507044db6c4eadfa2dab9b084e80****
	ResultId *string `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
}

func (s UpdateAutoThrottleRulesAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoThrottleRulesAsyncRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoThrottleRulesAsyncRequest) GetAbnormalDuration() *float64 {
	return s.AbnormalDuration
}

func (s *UpdateAutoThrottleRulesAsyncRequest) GetActiveSessions() *int64 {
	return s.ActiveSessions
}

func (s *UpdateAutoThrottleRulesAsyncRequest) GetAllowThrottleEndTime() *string {
	return s.AllowThrottleEndTime
}

func (s *UpdateAutoThrottleRulesAsyncRequest) GetAllowThrottleStartTime() *string {
	return s.AllowThrottleStartTime
}

func (s *UpdateAutoThrottleRulesAsyncRequest) GetAutoKillSession() *bool {
	return s.AutoKillSession
}

func (s *UpdateAutoThrottleRulesAsyncRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *UpdateAutoThrottleRulesAsyncRequest) GetCpuSessionRelation() *string {
	return s.CpuSessionRelation
}

func (s *UpdateAutoThrottleRulesAsyncRequest) GetCpuUsage() *float64 {
	return s.CpuUsage
}

func (s *UpdateAutoThrottleRulesAsyncRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *UpdateAutoThrottleRulesAsyncRequest) GetMaxThrottleTime() *float64 {
	return s.MaxThrottleTime
}

func (s *UpdateAutoThrottleRulesAsyncRequest) GetResultId() *string {
	return s.ResultId
}

func (s *UpdateAutoThrottleRulesAsyncRequest) SetAbnormalDuration(v float64) *UpdateAutoThrottleRulesAsyncRequest {
	s.AbnormalDuration = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncRequest) SetActiveSessions(v int64) *UpdateAutoThrottleRulesAsyncRequest {
	s.ActiveSessions = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncRequest) SetAllowThrottleEndTime(v string) *UpdateAutoThrottleRulesAsyncRequest {
	s.AllowThrottleEndTime = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncRequest) SetAllowThrottleStartTime(v string) *UpdateAutoThrottleRulesAsyncRequest {
	s.AllowThrottleStartTime = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncRequest) SetAutoKillSession(v bool) *UpdateAutoThrottleRulesAsyncRequest {
	s.AutoKillSession = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncRequest) SetConsoleContext(v string) *UpdateAutoThrottleRulesAsyncRequest {
	s.ConsoleContext = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncRequest) SetCpuSessionRelation(v string) *UpdateAutoThrottleRulesAsyncRequest {
	s.CpuSessionRelation = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncRequest) SetCpuUsage(v float64) *UpdateAutoThrottleRulesAsyncRequest {
	s.CpuUsage = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncRequest) SetInstanceIds(v string) *UpdateAutoThrottleRulesAsyncRequest {
	s.InstanceIds = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncRequest) SetMaxThrottleTime(v float64) *UpdateAutoThrottleRulesAsyncRequest {
	s.MaxThrottleTime = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncRequest) SetResultId(v string) *UpdateAutoThrottleRulesAsyncRequest {
	s.ResultId = &v
	return s
}

func (s *UpdateAutoThrottleRulesAsyncRequest) Validate() error {
	return dara.Validate(s)
}
