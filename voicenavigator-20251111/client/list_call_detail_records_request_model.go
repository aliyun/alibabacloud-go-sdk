// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallDetailRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessChannelId(v string) *ListCallDetailRecordsRequest
	GetAccessChannelId() *string
	SetAccessChannelType(v string) *ListCallDetailRecordsRequest
	GetAccessChannelType() *string
	SetCallee(v string) *ListCallDetailRecordsRequest
	GetCallee() *string
	SetCaller(v string) *ListCallDetailRecordsRequest
	GetCaller() *string
	SetDispositionCodes(v []*string) *ListCallDetailRecordsRequest
	GetDispositionCodes() []*string
	SetDispositionReasons(v []*string) *ListCallDetailRecordsRequest
	GetDispositionReasons() []*string
	SetDraftVersion(v bool) *ListCallDetailRecordsRequest
	GetDraftVersion() *bool
	SetEndTime(v string) *ListCallDetailRecordsRequest
	GetEndTime() *string
	SetInstanceId(v string) *ListCallDetailRecordsRequest
	GetInstanceId() *string
	SetIssueResolved(v bool) *ListCallDetailRecordsRequest
	GetIssueResolved() *bool
	SetMaxTalkTurns(v int32) *ListCallDetailRecordsRequest
	GetMaxTalkTurns() *int32
	SetMinTalkTurns(v int32) *ListCallDetailRecordsRequest
	GetMinTalkTurns() *int32
	SetPageNumber(v int32) *ListCallDetailRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCallDetailRecordsRequest
	GetPageSize() *int32
	SetScriptId(v string) *ListCallDetailRecordsRequest
	GetScriptId() *string
	SetSessionIds(v []*string) *ListCallDetailRecordsRequest
	GetSessionIds() []*string
	SetStartTime(v string) *ListCallDetailRecordsRequest
	GetStartTime() *string
}

type ListCallDetailRecordsRequest struct {
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
	Caller             *string   `json:"Caller,omitempty" xml:"Caller,omitempty"`
	DispositionCodes   []*string `json:"DispositionCodes,omitempty" xml:"DispositionCodes,omitempty" type:"Repeated"`
	DispositionReasons []*string `json:"DispositionReasons,omitempty" xml:"DispositionReasons,omitempty" type:"Repeated"`
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
	ScriptId   *string   `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	SessionIds []*string `json:"SessionIds,omitempty" xml:"SessionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1774858849987
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListCallDetailRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsRequest) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *ListCallDetailRecordsRequest) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *ListCallDetailRecordsRequest) GetCallee() *string {
	return s.Callee
}

func (s *ListCallDetailRecordsRequest) GetCaller() *string {
	return s.Caller
}

func (s *ListCallDetailRecordsRequest) GetDispositionCodes() []*string {
	return s.DispositionCodes
}

func (s *ListCallDetailRecordsRequest) GetDispositionReasons() []*string {
	return s.DispositionReasons
}

func (s *ListCallDetailRecordsRequest) GetDraftVersion() *bool {
	return s.DraftVersion
}

func (s *ListCallDetailRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListCallDetailRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallDetailRecordsRequest) GetIssueResolved() *bool {
	return s.IssueResolved
}

func (s *ListCallDetailRecordsRequest) GetMaxTalkTurns() *int32 {
	return s.MaxTalkTurns
}

func (s *ListCallDetailRecordsRequest) GetMinTalkTurns() *int32 {
	return s.MinTalkTurns
}

func (s *ListCallDetailRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallDetailRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallDetailRecordsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListCallDetailRecordsRequest) GetSessionIds() []*string {
	return s.SessionIds
}

func (s *ListCallDetailRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListCallDetailRecordsRequest) SetAccessChannelId(v string) *ListCallDetailRecordsRequest {
	s.AccessChannelId = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetAccessChannelType(v string) *ListCallDetailRecordsRequest {
	s.AccessChannelType = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetCallee(v string) *ListCallDetailRecordsRequest {
	s.Callee = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetCaller(v string) *ListCallDetailRecordsRequest {
	s.Caller = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetDispositionCodes(v []*string) *ListCallDetailRecordsRequest {
	s.DispositionCodes = v
	return s
}

func (s *ListCallDetailRecordsRequest) SetDispositionReasons(v []*string) *ListCallDetailRecordsRequest {
	s.DispositionReasons = v
	return s
}

func (s *ListCallDetailRecordsRequest) SetDraftVersion(v bool) *ListCallDetailRecordsRequest {
	s.DraftVersion = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetEndTime(v string) *ListCallDetailRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetInstanceId(v string) *ListCallDetailRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetIssueResolved(v bool) *ListCallDetailRecordsRequest {
	s.IssueResolved = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetMaxTalkTurns(v int32) *ListCallDetailRecordsRequest {
	s.MaxTalkTurns = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetMinTalkTurns(v int32) *ListCallDetailRecordsRequest {
	s.MinTalkTurns = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetPageNumber(v int32) *ListCallDetailRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetPageSize(v int32) *ListCallDetailRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetScriptId(v string) *ListCallDetailRecordsRequest {
	s.ScriptId = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetSessionIds(v []*string) *ListCallDetailRecordsRequest {
	s.SessionIds = v
	return s
}

func (s *ListCallDetailRecordsRequest) SetStartTime(v string) *ListCallDetailRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *ListCallDetailRecordsRequest) Validate() error {
	return dara.Validate(s)
}
