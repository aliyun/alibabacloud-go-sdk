// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessChannelId(v string) *ListCasesRequest
	GetAccessChannelId() *string
	SetAccessChannelType(v string) *ListCasesRequest
	GetAccessChannelType() *string
	SetCaller(v string) *ListCasesRequest
	GetCaller() *string
	SetCampaignId(v string) *ListCasesRequest
	GetCampaignId() *string
	SetCaseCompleted(v bool) *ListCasesRequest
	GetCaseCompleted() *bool
	SetCaseIds(v []*string) *ListCasesRequest
	GetCaseIds() []*string
	SetDispositionCodes(v []*string) *ListCasesRequest
	GetDispositionCodes() []*string
	SetDispositionReasons(v []*string) *ListCasesRequest
	GetDispositionReasons() []*string
	SetDraftVersion(v bool) *ListCasesRequest
	GetDraftVersion() *bool
	SetEndTime(v int64) *ListCasesRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListCasesRequest
	GetInstanceId() *string
	SetLabelSearch(v map[string]*string) *ListCasesRequest
	GetLabelSearch() map[string]*string
	SetMaxRingingDuration(v int64) *ListCasesRequest
	GetMaxRingingDuration() *int64
	SetMaxTalkTime(v int64) *ListCasesRequest
	GetMaxTalkTime() *int64
	SetMaxTalkTurns(v int64) *ListCasesRequest
	GetMaxTalkTurns() *int64
	SetMinRingingDuration(v int64) *ListCasesRequest
	GetMinRingingDuration() *int64
	SetMinTalkTime(v int64) *ListCasesRequest
	GetMinTalkTime() *int64
	SetMinTalkTurns(v int64) *ListCasesRequest
	GetMinTalkTurns() *int64
	SetPageNumber(v int32) *ListCasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCasesRequest
	GetPageSize() *int32
	SetPhoneNumber(v string) *ListCasesRequest
	GetPhoneNumber() *string
	SetScriptId(v string) *ListCasesRequest
	GetScriptId() *string
	SetStartTime(v int64) *ListCasesRequest
	GetStartTime() *int64
	SetStates(v []*string) *ListCasesRequest
	GetStates() []*string
}

type ListCasesRequest struct {
	// The access channel ID.
	//
	// example:
	//
	// 33606503-c22c-4547-a51c-dda5e8d87962
	AccessChannelId *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	// The access channel type.
	//
	// example:
	//
	// Test
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	// The caller number.
	//
	// example:
	//
	// 01080862792
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The outbound campaign ID.
	//
	// example:
	//
	// 7607dae1-91ad-47ea-ad76-3d81ac34f729
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// Specifies whether the case is completed.
	//
	// example:
	//
	// true
	CaseCompleted *bool `json:"CaseCompleted,omitempty" xml:"CaseCompleted,omitempty"`
	// The list of case IDs.
	CaseIds []*string `json:"CaseIds,omitempty" xml:"CaseIds,omitempty" type:"Repeated"`
	// The list of disposition codes.
	DispositionCodes []*string `json:"DispositionCodes,omitempty" xml:"DispositionCodes,omitempty" type:"Repeated"`
	// The list of disposition reasons.
	DispositionReasons []*string `json:"DispositionReasons,omitempty" xml:"DispositionReasons,omitempty" type:"Repeated"`
	// Specifies whether the version is a draft version.
	//
	// example:
	//
	// true
	DraftVersion *bool `json:"DraftVersion,omitempty" xml:"DraftVersion,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1578995079000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a5fc6490-ef1e-4666-870a-07a4e586c414
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The label search condition.
	LabelSearch map[string]*string `json:"LabelSearch,omitempty" xml:"LabelSearch,omitempty"`
	// The maximum ringing duration.
	//
	// example:
	//
	// 2
	MaxRingingDuration *int64 `json:"MaxRingingDuration,omitempty" xml:"MaxRingingDuration,omitempty"`
	// The maximum talk time.
	//
	// example:
	//
	// 2
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	// The maximum number of conversation turns.
	//
	// example:
	//
	// 2
	MaxTalkTurns *int64 `json:"MaxTalkTurns,omitempty" xml:"MaxTalkTurns,omitempty"`
	// The minimum ringing duration.
	//
	// example:
	//
	// 1
	MinRingingDuration *int64 `json:"MinRingingDuration,omitempty" xml:"MinRingingDuration,omitempty"`
	// The minimum talk time.
	//
	// example:
	//
	// 1
	MinTalkTime *int64 `json:"MinTalkTime,omitempty" xml:"MinTalkTime,omitempty"`
	// The minimum number of conversation turns.
	//
	// example:
	//
	// 1
	MinTalkTurns *int64 `json:"MinTalkTurns,omitempty" xml:"MinTalkTurns,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The called number.
	//
	// example:
	//
	// 18512345678
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The script ID.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1578965079000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The list of case states.
	States []*string `json:"States,omitempty" xml:"States,omitempty" type:"Repeated"`
}

func (s ListCasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCasesRequest) GoString() string {
	return s.String()
}

func (s *ListCasesRequest) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *ListCasesRequest) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *ListCasesRequest) GetCaller() *string {
	return s.Caller
}

func (s *ListCasesRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListCasesRequest) GetCaseCompleted() *bool {
	return s.CaseCompleted
}

func (s *ListCasesRequest) GetCaseIds() []*string {
	return s.CaseIds
}

func (s *ListCasesRequest) GetDispositionCodes() []*string {
	return s.DispositionCodes
}

func (s *ListCasesRequest) GetDispositionReasons() []*string {
	return s.DispositionReasons
}

func (s *ListCasesRequest) GetDraftVersion() *bool {
	return s.DraftVersion
}

func (s *ListCasesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCasesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCasesRequest) GetLabelSearch() map[string]*string {
	return s.LabelSearch
}

func (s *ListCasesRequest) GetMaxRingingDuration() *int64 {
	return s.MaxRingingDuration
}

func (s *ListCasesRequest) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListCasesRequest) GetMaxTalkTurns() *int64 {
	return s.MaxTalkTurns
}

func (s *ListCasesRequest) GetMinRingingDuration() *int64 {
	return s.MinRingingDuration
}

func (s *ListCasesRequest) GetMinTalkTime() *int64 {
	return s.MinTalkTime
}

func (s *ListCasesRequest) GetMinTalkTurns() *int64 {
	return s.MinTalkTurns
}

func (s *ListCasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCasesRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListCasesRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListCasesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCasesRequest) GetStates() []*string {
	return s.States
}

func (s *ListCasesRequest) SetAccessChannelId(v string) *ListCasesRequest {
	s.AccessChannelId = &v
	return s
}

func (s *ListCasesRequest) SetAccessChannelType(v string) *ListCasesRequest {
	s.AccessChannelType = &v
	return s
}

func (s *ListCasesRequest) SetCaller(v string) *ListCasesRequest {
	s.Caller = &v
	return s
}

func (s *ListCasesRequest) SetCampaignId(v string) *ListCasesRequest {
	s.CampaignId = &v
	return s
}

func (s *ListCasesRequest) SetCaseCompleted(v bool) *ListCasesRequest {
	s.CaseCompleted = &v
	return s
}

func (s *ListCasesRequest) SetCaseIds(v []*string) *ListCasesRequest {
	s.CaseIds = v
	return s
}

func (s *ListCasesRequest) SetDispositionCodes(v []*string) *ListCasesRequest {
	s.DispositionCodes = v
	return s
}

func (s *ListCasesRequest) SetDispositionReasons(v []*string) *ListCasesRequest {
	s.DispositionReasons = v
	return s
}

func (s *ListCasesRequest) SetDraftVersion(v bool) *ListCasesRequest {
	s.DraftVersion = &v
	return s
}

func (s *ListCasesRequest) SetEndTime(v int64) *ListCasesRequest {
	s.EndTime = &v
	return s
}

func (s *ListCasesRequest) SetInstanceId(v string) *ListCasesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCasesRequest) SetLabelSearch(v map[string]*string) *ListCasesRequest {
	s.LabelSearch = v
	return s
}

func (s *ListCasesRequest) SetMaxRingingDuration(v int64) *ListCasesRequest {
	s.MaxRingingDuration = &v
	return s
}

func (s *ListCasesRequest) SetMaxTalkTime(v int64) *ListCasesRequest {
	s.MaxTalkTime = &v
	return s
}

func (s *ListCasesRequest) SetMaxTalkTurns(v int64) *ListCasesRequest {
	s.MaxTalkTurns = &v
	return s
}

func (s *ListCasesRequest) SetMinRingingDuration(v int64) *ListCasesRequest {
	s.MinRingingDuration = &v
	return s
}

func (s *ListCasesRequest) SetMinTalkTime(v int64) *ListCasesRequest {
	s.MinTalkTime = &v
	return s
}

func (s *ListCasesRequest) SetMinTalkTurns(v int64) *ListCasesRequest {
	s.MinTalkTurns = &v
	return s
}

func (s *ListCasesRequest) SetPageNumber(v int32) *ListCasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCasesRequest) SetPageSize(v int32) *ListCasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCasesRequest) SetPhoneNumber(v string) *ListCasesRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ListCasesRequest) SetScriptId(v string) *ListCasesRequest {
	s.ScriptId = &v
	return s
}

func (s *ListCasesRequest) SetStartTime(v int64) *ListCasesRequest {
	s.StartTime = &v
	return s
}

func (s *ListCasesRequest) SetStates(v []*string) *ListCasesRequest {
	s.States = v
	return s
}

func (s *ListCasesRequest) Validate() error {
	return dara.Validate(s)
}
