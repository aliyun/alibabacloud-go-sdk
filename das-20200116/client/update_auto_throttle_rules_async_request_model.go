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
	// The duration threshold of the anomaly that triggers automatic SQL throttling. The value must be a positive integer greater than or equal to 2. Unit: minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	AbnormalDuration *float64 `json:"AbnormalDuration,omitempty" xml:"AbnormalDuration,omitempty"`
	// The active sessions threshold.
	//
	// - If the relationship with the CPU utilization threshold is **OR**, the value must be greater than or equal to 16.
	//
	// - If the relationship with the CPU utilization threshold is **AND**, the value must be greater than or equal to 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16
	ActiveSessions *int64 `json:"ActiveSessions,omitempty" xml:"ActiveSessions,omitempty"`
	// The end time of the throttling time window (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 23:59Z
	AllowThrottleEndTime *string `json:"AllowThrottleEndTime,omitempty" xml:"AllowThrottleEndTime,omitempty"`
	// The start time of the throttling time window (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 00:00Z
	AllowThrottleStartTime *string `json:"AllowThrottleStartTime,omitempty" xml:"AllowThrottleStartTime,omitempty"`
	// Specifies whether to simultaneously kill abnormal SQL statements that are being executed.
	//
	// > Abnormal SQL statements are those that match the SQL templates to be throttled.
	//
	// Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoKillSession *bool `json:"AutoKillSession,omitempty" xml:"AutoKillSession,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The logical relationship between the CPU utilization threshold and the active sessions threshold. Valid values:
	//
	// - **AND**: both conditions must be met.
	//
	// - **OR**: either condition must be met.
	//
	// This parameter is required.
	//
	// example:
	//
	// OR
	CpuSessionRelation *string `json:"CpuSessionRelation,omitempty" xml:"CpuSessionRelation,omitempty"`
	// The CPU utilization threshold. Valid values: 70% to 100%.
	//
	// This parameter is required.
	//
	// example:
	//
	// 70
	CpuUsage *float64 `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	// The database instance IDs.
	//
	// > The data format is JSONArray, such as `[\\"Instance ID 1\\",\\"Instance ID 2\\"]`. Separate instance IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"rm-2ze8g2am97624****\\",\\"rm-2ze9xrhze0709****\\"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The maximum throttling duration. The value must be a positive integer. Unit: minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxThrottleTime *float64 `json:"MaxThrottleTime,omitempty" xml:"MaxThrottleTime,omitempty"`
	// The ID of the asynchronous request.
	//
	// > An asynchronous call does not immediately return complete results. First, call this operation to obtain the **ResultId**. Then, use the returned **ResultId*	- to initiate the call again until **isFinish*	- is **true**, at which point the complete results are returned. This means that you must call this operation at least twice to obtain complete data.
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
