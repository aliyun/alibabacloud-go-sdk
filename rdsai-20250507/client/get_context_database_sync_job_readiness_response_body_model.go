// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncJobReadinessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContextDatabaseSyncJobReadinessResponseBody
	GetCode() *string
	SetData(v *GetContextDatabaseSyncJobReadinessResponseBodyData) *GetContextDatabaseSyncJobReadinessResponseBody
	GetData() *GetContextDatabaseSyncJobReadinessResponseBodyData
	SetHttpStatusCode(v int64) *GetContextDatabaseSyncJobReadinessResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetContextDatabaseSyncJobReadinessResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContextDatabaseSyncJobReadinessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetContextDatabaseSyncJobReadinessResponseBody
	GetSuccess() *bool
}

type GetContextDatabaseSyncJobReadinessResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetContextDatabaseSyncJobReadinessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// Sample value
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContextDatabaseSyncJobReadinessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobReadinessResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) GetData() *GetContextDatabaseSyncJobReadinessResponseBodyData {
	return s.Data
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) SetCode(v string) *GetContextDatabaseSyncJobReadinessResponseBody {
	s.Code = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) SetData(v *GetContextDatabaseSyncJobReadinessResponseBodyData) *GetContextDatabaseSyncJobReadinessResponseBody {
	s.Data = v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) SetHttpStatusCode(v int64) *GetContextDatabaseSyncJobReadinessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) SetMessage(v string) *GetContextDatabaseSyncJobReadinessResponseBody {
	s.Message = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) SetRequestId(v string) *GetContextDatabaseSyncJobReadinessResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) SetSuccess(v bool) *GetContextDatabaseSyncJobReadinessResponseBody {
	s.Success = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextDatabaseSyncJobReadinessResponseBodyData struct {
	// The reason for the automatic pause.
	//
	// example:
	//
	// Three consecutive execution failures
	AutoPauseReason *string `json:"AutoPauseReason,omitempty" xml:"AutoPauseReason,omitempty"`
	// The summary of check items.
	//
	// example:
	//
	// {"jobEnabled":true,"connectionReady":true,"selectionConfirmed":true,"scheduleConfigured":true,"scheduleNotAutoPaused":true}
	Checks *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks `json:"Checks,omitempty" xml:"Checks,omitempty" type:"Struct"`
	// Indicates whether the task is ready.
	//
	// example:
	//
	// true
	Ready *bool `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// Indicates whether the task is ready for a manual run.
	//
	// example:
	//
	// true
	ReadyForManualRun *bool `json:"ReadyForManualRun,omitempty" xml:"ReadyForManualRun,omitempty"`
	// Indicates whether the task is ready for a scheduled run.
	//
	// example:
	//
	// true
	ReadyForScheduledRun *bool `json:"ReadyForScheduledRun,omitempty" xml:"ReadyForScheduledRun,omitempty"`
}

func (s GetContextDatabaseSyncJobReadinessResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobReadinessResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyData) GetAutoPauseReason() *string {
	return s.AutoPauseReason
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyData) GetChecks() *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks {
	return s.Checks
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyData) GetReady() *bool {
	return s.Ready
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyData) GetReadyForManualRun() *bool {
	return s.ReadyForManualRun
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyData) GetReadyForScheduledRun() *bool {
	return s.ReadyForScheduledRun
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyData) SetAutoPauseReason(v string) *GetContextDatabaseSyncJobReadinessResponseBodyData {
	s.AutoPauseReason = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyData) SetChecks(v *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) *GetContextDatabaseSyncJobReadinessResponseBodyData {
	s.Checks = v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyData) SetReady(v bool) *GetContextDatabaseSyncJobReadinessResponseBodyData {
	s.Ready = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyData) SetReadyForManualRun(v bool) *GetContextDatabaseSyncJobReadinessResponseBodyData {
	s.ReadyForManualRun = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyData) SetReadyForScheduledRun(v bool) *GetContextDatabaseSyncJobReadinessResponseBodyData {
	s.ReadyForScheduledRun = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyData) Validate() error {
	if s.Checks != nil {
		if err := s.Checks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextDatabaseSyncJobReadinessResponseBodyDataChecks struct {
	// Indicates whether the data source is ready.
	//
	// example:
	//
	// true
	ConnectionReady *bool `json:"ConnectionReady,omitempty" xml:"ConnectionReady,omitempty"`
	// Indicates whether the import task is enabled.
	//
	// example:
	//
	// true
	JobEnabled *bool `json:"JobEnabled,omitempty" xml:"JobEnabled,omitempty"`
	// Indicates whether scheduling is configured.
	//
	// example:
	//
	// true
	ScheduleConfigured *bool `json:"ScheduleConfigured,omitempty" xml:"ScheduleConfigured,omitempty"`
	// Indicates whether scheduling is not automatically paused.
	//
	// example:
	//
	// true
	ScheduleNotAutoPaused *bool `json:"ScheduleNotAutoPaused,omitempty" xml:"ScheduleNotAutoPaused,omitempty"`
	// Indicates whether the selection is confirmed.
	//
	// example:
	//
	// true
	SelectionConfirmed *bool `json:"SelectionConfirmed,omitempty" xml:"SelectionConfirmed,omitempty"`
	// Indicates whether the target is ready.
	TargetReady *bool `json:"TargetReady,omitempty" xml:"TargetReady,omitempty"`
}

func (s GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) GetConnectionReady() *bool {
	return s.ConnectionReady
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) GetJobEnabled() *bool {
	return s.JobEnabled
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) GetScheduleConfigured() *bool {
	return s.ScheduleConfigured
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) GetScheduleNotAutoPaused() *bool {
	return s.ScheduleNotAutoPaused
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) GetSelectionConfirmed() *bool {
	return s.SelectionConfirmed
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) GetTargetReady() *bool {
	return s.TargetReady
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) SetConnectionReady(v bool) *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks {
	s.ConnectionReady = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) SetJobEnabled(v bool) *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks {
	s.JobEnabled = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) SetScheduleConfigured(v bool) *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks {
	s.ScheduleConfigured = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) SetScheduleNotAutoPaused(v bool) *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks {
	s.ScheduleNotAutoPaused = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) SetSelectionConfirmed(v bool) *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks {
	s.SelectionConfirmed = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) SetTargetReady(v bool) *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks {
	s.TargetReady = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponseBodyDataChecks) Validate() error {
	return dara.Validate(s)
}
