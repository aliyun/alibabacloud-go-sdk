// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallDetailRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessChannelId(v string) *ListCallDetailRecordsShrinkRequest
	GetAccessChannelId() *string
	SetAccessChannelType(v string) *ListCallDetailRecordsShrinkRequest
	GetAccessChannelType() *string
	SetCallee(v string) *ListCallDetailRecordsShrinkRequest
	GetCallee() *string
	SetCaller(v string) *ListCallDetailRecordsShrinkRequest
	GetCaller() *string
	SetDispositionCodesShrink(v string) *ListCallDetailRecordsShrinkRequest
	GetDispositionCodesShrink() *string
	SetDispositionReasonsShrink(v string) *ListCallDetailRecordsShrinkRequest
	GetDispositionReasonsShrink() *string
	SetDraftVersion(v bool) *ListCallDetailRecordsShrinkRequest
	GetDraftVersion() *bool
	SetEndTime(v string) *ListCallDetailRecordsShrinkRequest
	GetEndTime() *string
	SetInstanceId(v string) *ListCallDetailRecordsShrinkRequest
	GetInstanceId() *string
	SetIssueResolved(v bool) *ListCallDetailRecordsShrinkRequest
	GetIssueResolved() *bool
	SetMaxTalkTurns(v int32) *ListCallDetailRecordsShrinkRequest
	GetMaxTalkTurns() *int32
	SetMinTalkTurns(v int32) *ListCallDetailRecordsShrinkRequest
	GetMinTalkTurns() *int32
	SetPageNumber(v int32) *ListCallDetailRecordsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCallDetailRecordsShrinkRequest
	GetPageSize() *int32
	SetScriptId(v string) *ListCallDetailRecordsShrinkRequest
	GetScriptId() *string
	SetSessionIdsShrink(v string) *ListCallDetailRecordsShrinkRequest
	GetSessionIdsShrink() *string
	SetStartTime(v string) *ListCallDetailRecordsShrinkRequest
	GetStartTime() *string
}

type ListCallDetailRecordsShrinkRequest struct {
	// example:
	//
	// 33606503-c22c-4547-a51c-dda5e8d87962
	AccessChannelId *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	// example:
	//
	// PSTN
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	// example:
	//
	// 18104560087
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 09876543210,
	Caller                   *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	DispositionCodesShrink   *string `json:"DispositionCodes,omitempty" xml:"DispositionCodes,omitempty"`
	DispositionReasonsShrink *string `json:"DispositionReasons,omitempty" xml:"DispositionReasons,omitempty"`
	// example:
	//
	// false
	DraftVersion *bool `json:"DraftVersion,omitempty" xml:"DraftVersion,omitempty"`
	// example:
	//
	// 1582103299434
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// false
	IssueResolved *bool `json:"IssueResolved,omitempty" xml:"IssueResolved,omitempty"`
	// example:
	//
	// 10
	MaxTalkTurns *int32 `json:"MaxTalkTurns,omitempty" xml:"MaxTalkTurns,omitempty"`
	// example:
	//
	// 1
	MinTalkTurns *int32 `json:"MinTalkTurns,omitempty" xml:"MinTalkTurns,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b090
	ScriptId         *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	SessionIdsShrink *string `json:"SessionIds,omitempty" xml:"SessionIds,omitempty"`
	// example:
	//
	// 1774858849987
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListCallDetailRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsShrinkRequest) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *ListCallDetailRecordsShrinkRequest) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *ListCallDetailRecordsShrinkRequest) GetCallee() *string {
	return s.Callee
}

func (s *ListCallDetailRecordsShrinkRequest) GetCaller() *string {
	return s.Caller
}

func (s *ListCallDetailRecordsShrinkRequest) GetDispositionCodesShrink() *string {
	return s.DispositionCodesShrink
}

func (s *ListCallDetailRecordsShrinkRequest) GetDispositionReasonsShrink() *string {
	return s.DispositionReasonsShrink
}

func (s *ListCallDetailRecordsShrinkRequest) GetDraftVersion() *bool {
	return s.DraftVersion
}

func (s *ListCallDetailRecordsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListCallDetailRecordsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallDetailRecordsShrinkRequest) GetIssueResolved() *bool {
	return s.IssueResolved
}

func (s *ListCallDetailRecordsShrinkRequest) GetMaxTalkTurns() *int32 {
	return s.MaxTalkTurns
}

func (s *ListCallDetailRecordsShrinkRequest) GetMinTalkTurns() *int32 {
	return s.MinTalkTurns
}

func (s *ListCallDetailRecordsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallDetailRecordsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallDetailRecordsShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListCallDetailRecordsShrinkRequest) GetSessionIdsShrink() *string {
	return s.SessionIdsShrink
}

func (s *ListCallDetailRecordsShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListCallDetailRecordsShrinkRequest) SetAccessChannelId(v string) *ListCallDetailRecordsShrinkRequest {
	s.AccessChannelId = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetAccessChannelType(v string) *ListCallDetailRecordsShrinkRequest {
	s.AccessChannelType = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetCallee(v string) *ListCallDetailRecordsShrinkRequest {
	s.Callee = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetCaller(v string) *ListCallDetailRecordsShrinkRequest {
	s.Caller = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetDispositionCodesShrink(v string) *ListCallDetailRecordsShrinkRequest {
	s.DispositionCodesShrink = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetDispositionReasonsShrink(v string) *ListCallDetailRecordsShrinkRequest {
	s.DispositionReasonsShrink = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetDraftVersion(v bool) *ListCallDetailRecordsShrinkRequest {
	s.DraftVersion = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetEndTime(v string) *ListCallDetailRecordsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetInstanceId(v string) *ListCallDetailRecordsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetIssueResolved(v bool) *ListCallDetailRecordsShrinkRequest {
	s.IssueResolved = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetMaxTalkTurns(v int32) *ListCallDetailRecordsShrinkRequest {
	s.MaxTalkTurns = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetMinTalkTurns(v int32) *ListCallDetailRecordsShrinkRequest {
	s.MinTalkTurns = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetPageNumber(v int32) *ListCallDetailRecordsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetPageSize(v int32) *ListCallDetailRecordsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetScriptId(v string) *ListCallDetailRecordsShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetSessionIdsShrink(v string) *ListCallDetailRecordsShrinkRequest {
	s.SessionIdsShrink = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) SetStartTime(v string) *ListCallDetailRecordsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListCallDetailRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
