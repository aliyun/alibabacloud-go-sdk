// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySyncResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySyncResultResponseBody
	GetCode() *string
	SetCompletedAt(v string) *QuerySyncResultResponseBody
	GetCompletedAt() *string
	SetCorpId(v string) *QuerySyncResultResponseBody
	GetCorpId() *string
	SetDeptStats(v *QuerySyncResultResponseBodyDeptStats) *QuerySyncResultResponseBody
	GetDeptStats() *QuerySyncResultResponseBodyDeptStats
	SetDurationSeconds(v int64) *QuerySyncResultResponseBody
	GetDurationSeconds() *int64
	SetErrorMessage(v string) *QuerySyncResultResponseBody
	GetErrorMessage() *string
	SetMemberStats(v *QuerySyncResultResponseBodyMemberStats) *QuerySyncResultResponseBody
	GetMemberStats() *QuerySyncResultResponseBodyMemberStats
	SetMessage(v string) *QuerySyncResultResponseBody
	GetMessage() *string
	SetPlatformType(v string) *QuerySyncResultResponseBody
	GetPlatformType() *string
	SetRequestId(v string) *QuerySyncResultResponseBody
	GetRequestId() *string
	SetStartedAt(v string) *QuerySyncResultResponseBody
	GetStartedAt() *string
	SetStatus(v string) *QuerySyncResultResponseBody
	GetStatus() *string
	SetSubmittedAt(v string) *QuerySyncResultResponseBody
	GetSubmittedAt() *string
	SetSummary(v string) *QuerySyncResultResponseBody
	GetSummary() *string
	SetTaskId(v int64) *QuerySyncResultResponseBody
	GetTaskId() *int64
}

type QuerySyncResultResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The time when the task was completed (ISO 8601 format).
	//
	// example:
	//
	// string_value
	CompletedAt *string `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// The enterprise identifier.
	//
	// example:
	//
	// exampleCorpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// The department synchronization statistics. This field has a value when the task is completed.
	DeptStats *QuerySyncResultResponseBodyDeptStats `json:"deptStats,omitempty" xml:"deptStats,omitempty" type:"Struct"`
	// The execution duration, in seconds.
	//
	// example:
	//
	// 1
	DurationSeconds *int64 `json:"durationSeconds,omitempty" xml:"durationSeconds,omitempty"`
	// The error message.
	//
	// example:
	//
	// string_value
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The member synchronization statistics. This field has a value when syncMembers is set to true and the task is completed.
	MemberStats *QuerySyncResultResponseBodyMemberStats `json:"memberStats,omitempty" xml:"memberStats,omitempty" type:"Struct"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The platform type.
	//
	// example:
	//
	// string_value
	PlatformType *string `json:"platformType,omitempty" xml:"platformType,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The time when the task started (ISO 8601 format).
	//
	// example:
	//
	// string_value
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// The task status. Valid values: PENDING, RUNNING, COMPLETED, FAILED, TIMEOUT, and CANCELED.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the task was submitted (ISO 8601 format).
	//
	// example:
	//
	// string_value
	SubmittedAt *string `json:"submittedAt,omitempty" xml:"submittedAt,omitempty"`
	// The intelligent meeting summary content.
	//
	// example:
	//
	// string_value
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QuerySyncResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySyncResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySyncResultResponseBody) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *QuerySyncResultResponseBody) GetCorpId() *string {
	return s.CorpId
}

func (s *QuerySyncResultResponseBody) GetDeptStats() *QuerySyncResultResponseBodyDeptStats {
	return s.DeptStats
}

func (s *QuerySyncResultResponseBody) GetDurationSeconds() *int64 {
	return s.DurationSeconds
}

func (s *QuerySyncResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QuerySyncResultResponseBody) GetMemberStats() *QuerySyncResultResponseBodyMemberStats {
	return s.MemberStats
}

func (s *QuerySyncResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySyncResultResponseBody) GetPlatformType() *string {
	return s.PlatformType
}

func (s *QuerySyncResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySyncResultResponseBody) GetStartedAt() *string {
	return s.StartedAt
}

func (s *QuerySyncResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *QuerySyncResultResponseBody) GetSubmittedAt() *string {
	return s.SubmittedAt
}

func (s *QuerySyncResultResponseBody) GetSummary() *string {
	return s.Summary
}

func (s *QuerySyncResultResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QuerySyncResultResponseBody) SetCode(v string) *QuerySyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetCompletedAt(v string) *QuerySyncResultResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetCorpId(v string) *QuerySyncResultResponseBody {
	s.CorpId = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetDeptStats(v *QuerySyncResultResponseBodyDeptStats) *QuerySyncResultResponseBody {
	s.DeptStats = v
	return s
}

func (s *QuerySyncResultResponseBody) SetDurationSeconds(v int64) *QuerySyncResultResponseBody {
	s.DurationSeconds = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetErrorMessage(v string) *QuerySyncResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetMemberStats(v *QuerySyncResultResponseBodyMemberStats) *QuerySyncResultResponseBody {
	s.MemberStats = v
	return s
}

func (s *QuerySyncResultResponseBody) SetMessage(v string) *QuerySyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetPlatformType(v string) *QuerySyncResultResponseBody {
	s.PlatformType = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetRequestId(v string) *QuerySyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetStartedAt(v string) *QuerySyncResultResponseBody {
	s.StartedAt = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetStatus(v string) *QuerySyncResultResponseBody {
	s.Status = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetSubmittedAt(v string) *QuerySyncResultResponseBody {
	s.SubmittedAt = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetSummary(v string) *QuerySyncResultResponseBody {
	s.Summary = &v
	return s
}

func (s *QuerySyncResultResponseBody) SetTaskId(v int64) *QuerySyncResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *QuerySyncResultResponseBody) Validate() error {
	if s.DeptStats != nil {
		if err := s.DeptStats.Validate(); err != nil {
			return err
		}
	}
	if s.MemberStats != nil {
		if err := s.MemberStats.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySyncResultResponseBodyDeptStats struct {
	// The total number of external departments.
	//
	// example:
	//
	// 1
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// The number of user groups marked for deletion.
	//
	// example:
	//
	// 1
	Deleted *int64 `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// The number of moved user groups.
	//
	// example:
	//
	// 1
	Moved *int64 `json:"moved,omitempty" xml:"moved,omitempty"`
	// The number of renamed user groups.
	//
	// example:
	//
	// 1
	Renamed *int64 `json:"renamed,omitempty" xml:"renamed,omitempty"`
	// The number of skipped user groups.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"skipped,omitempty" xml:"skipped,omitempty"`
	// The total number of external departments.
	//
	// example:
	//
	// 1
	TotalExternal *int64 `json:"totalExternal,omitempty" xml:"totalExternal,omitempty"`
}

func (s QuerySyncResultResponseBodyDeptStats) String() string {
	return dara.Prettify(s)
}

func (s QuerySyncResultResponseBodyDeptStats) GoString() string {
	return s.String()
}

func (s *QuerySyncResultResponseBodyDeptStats) GetCreated() *int64 {
	return s.Created
}

func (s *QuerySyncResultResponseBodyDeptStats) GetDeleted() *int64 {
	return s.Deleted
}

func (s *QuerySyncResultResponseBodyDeptStats) GetMoved() *int64 {
	return s.Moved
}

func (s *QuerySyncResultResponseBodyDeptStats) GetRenamed() *int64 {
	return s.Renamed
}

func (s *QuerySyncResultResponseBodyDeptStats) GetSkipped() *int64 {
	return s.Skipped
}

func (s *QuerySyncResultResponseBodyDeptStats) GetTotalExternal() *int64 {
	return s.TotalExternal
}

func (s *QuerySyncResultResponseBodyDeptStats) SetCreated(v int64) *QuerySyncResultResponseBodyDeptStats {
	s.Created = &v
	return s
}

func (s *QuerySyncResultResponseBodyDeptStats) SetDeleted(v int64) *QuerySyncResultResponseBodyDeptStats {
	s.Deleted = &v
	return s
}

func (s *QuerySyncResultResponseBodyDeptStats) SetMoved(v int64) *QuerySyncResultResponseBodyDeptStats {
	s.Moved = &v
	return s
}

func (s *QuerySyncResultResponseBodyDeptStats) SetRenamed(v int64) *QuerySyncResultResponseBodyDeptStats {
	s.Renamed = &v
	return s
}

func (s *QuerySyncResultResponseBodyDeptStats) SetSkipped(v int64) *QuerySyncResultResponseBodyDeptStats {
	s.Skipped = &v
	return s
}

func (s *QuerySyncResultResponseBodyDeptStats) SetTotalExternal(v int64) *QuerySyncResultResponseBodyDeptStats {
	s.TotalExternal = &v
	return s
}

func (s *QuerySyncResultResponseBodyDeptStats) Validate() error {
	return dara.Validate(s)
}

type QuerySyncResultResponseBodyMemberStats struct {
	// The number of failed members.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"failed,omitempty" xml:"failed,omitempty"`
	// The number of added member relationships.
	//
	// example:
	//
	// 1
	RelationshipAdded *int64 `json:"relationshipAdded,omitempty" xml:"relationshipAdded,omitempty"`
	// The number of removed member relationships.
	//
	// example:
	//
	// 1
	RelationshipRemoved *int64 `json:"relationshipRemoved,omitempty" xml:"relationshipRemoved,omitempty"`
	// The total number of external members.
	//
	// example:
	//
	// 1
	TotalExternal *int64 `json:"totalExternal,omitempty" xml:"totalExternal,omitempty"`
	// The number of unchanged member relationships.
	//
	// example:
	//
	// 1
	Unchanged *int64 `json:"unchanged,omitempty" xml:"unchanged,omitempty"`
}

func (s QuerySyncResultResponseBodyMemberStats) String() string {
	return dara.Prettify(s)
}

func (s QuerySyncResultResponseBodyMemberStats) GoString() string {
	return s.String()
}

func (s *QuerySyncResultResponseBodyMemberStats) GetFailed() *int64 {
	return s.Failed
}

func (s *QuerySyncResultResponseBodyMemberStats) GetRelationshipAdded() *int64 {
	return s.RelationshipAdded
}

func (s *QuerySyncResultResponseBodyMemberStats) GetRelationshipRemoved() *int64 {
	return s.RelationshipRemoved
}

func (s *QuerySyncResultResponseBodyMemberStats) GetTotalExternal() *int64 {
	return s.TotalExternal
}

func (s *QuerySyncResultResponseBodyMemberStats) GetUnchanged() *int64 {
	return s.Unchanged
}

func (s *QuerySyncResultResponseBodyMemberStats) SetFailed(v int64) *QuerySyncResultResponseBodyMemberStats {
	s.Failed = &v
	return s
}

func (s *QuerySyncResultResponseBodyMemberStats) SetRelationshipAdded(v int64) *QuerySyncResultResponseBodyMemberStats {
	s.RelationshipAdded = &v
	return s
}

func (s *QuerySyncResultResponseBodyMemberStats) SetRelationshipRemoved(v int64) *QuerySyncResultResponseBodyMemberStats {
	s.RelationshipRemoved = &v
	return s
}

func (s *QuerySyncResultResponseBodyMemberStats) SetTotalExternal(v int64) *QuerySyncResultResponseBodyMemberStats {
	s.TotalExternal = &v
	return s
}

func (s *QuerySyncResultResponseBodyMemberStats) SetUnchanged(v int64) *QuerySyncResultResponseBodyMemberStats {
	s.Unchanged = &v
	return s
}

func (s *QuerySyncResultResponseBodyMemberStats) Validate() error {
	return dara.Validate(s)
}
