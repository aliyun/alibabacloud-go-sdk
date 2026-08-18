// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCasesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessChannelId(v string) *ListCasesShrinkRequest
	GetAccessChannelId() *string
	SetAccessChannelType(v string) *ListCasesShrinkRequest
	GetAccessChannelType() *string
	SetCaller(v string) *ListCasesShrinkRequest
	GetCaller() *string
	SetCampaignId(v string) *ListCasesShrinkRequest
	GetCampaignId() *string
	SetCaseCompleted(v bool) *ListCasesShrinkRequest
	GetCaseCompleted() *bool
	SetCaseIdsShrink(v string) *ListCasesShrinkRequest
	GetCaseIdsShrink() *string
	SetDispositionCodesShrink(v string) *ListCasesShrinkRequest
	GetDispositionCodesShrink() *string
	SetDispositionReasonsShrink(v string) *ListCasesShrinkRequest
	GetDispositionReasonsShrink() *string
	SetDraftVersion(v bool) *ListCasesShrinkRequest
	GetDraftVersion() *bool
	SetEndTime(v int64) *ListCasesShrinkRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListCasesShrinkRequest
	GetInstanceId() *string
	SetLabelSearchShrink(v string) *ListCasesShrinkRequest
	GetLabelSearchShrink() *string
	SetMaxRingingDuration(v int64) *ListCasesShrinkRequest
	GetMaxRingingDuration() *int64
	SetMaxTalkTime(v int64) *ListCasesShrinkRequest
	GetMaxTalkTime() *int64
	SetMaxTalkTurns(v int64) *ListCasesShrinkRequest
	GetMaxTalkTurns() *int64
	SetMinRingingDuration(v int64) *ListCasesShrinkRequest
	GetMinRingingDuration() *int64
	SetMinTalkTime(v int64) *ListCasesShrinkRequest
	GetMinTalkTime() *int64
	SetMinTalkTurns(v int64) *ListCasesShrinkRequest
	GetMinTalkTurns() *int64
	SetPageNumber(v int32) *ListCasesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCasesShrinkRequest
	GetPageSize() *int32
	SetPhoneNumber(v string) *ListCasesShrinkRequest
	GetPhoneNumber() *string
	SetScriptId(v string) *ListCasesShrinkRequest
	GetScriptId() *string
	SetStartTime(v int64) *ListCasesShrinkRequest
	GetStartTime() *int64
	SetStatesShrink(v string) *ListCasesShrinkRequest
	GetStatesShrink() *string
}

type ListCasesShrinkRequest struct {
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
	CaseIdsShrink *string `json:"CaseIds,omitempty" xml:"CaseIds,omitempty"`
	// The list of disposition codes.
	DispositionCodesShrink *string `json:"DispositionCodes,omitempty" xml:"DispositionCodes,omitempty"`
	// The list of disposition reasons.
	DispositionReasonsShrink *string `json:"DispositionReasons,omitempty" xml:"DispositionReasons,omitempty"`
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
	LabelSearchShrink *string `json:"LabelSearch,omitempty" xml:"LabelSearch,omitempty"`
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
	StatesShrink *string `json:"States,omitempty" xml:"States,omitempty"`
}

func (s ListCasesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCasesShrinkRequest) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *ListCasesShrinkRequest) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *ListCasesShrinkRequest) GetCaller() *string {
	return s.Caller
}

func (s *ListCasesShrinkRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListCasesShrinkRequest) GetCaseCompleted() *bool {
	return s.CaseCompleted
}

func (s *ListCasesShrinkRequest) GetCaseIdsShrink() *string {
	return s.CaseIdsShrink
}

func (s *ListCasesShrinkRequest) GetDispositionCodesShrink() *string {
	return s.DispositionCodesShrink
}

func (s *ListCasesShrinkRequest) GetDispositionReasonsShrink() *string {
	return s.DispositionReasonsShrink
}

func (s *ListCasesShrinkRequest) GetDraftVersion() *bool {
	return s.DraftVersion
}

func (s *ListCasesShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCasesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCasesShrinkRequest) GetLabelSearchShrink() *string {
	return s.LabelSearchShrink
}

func (s *ListCasesShrinkRequest) GetMaxRingingDuration() *int64 {
	return s.MaxRingingDuration
}

func (s *ListCasesShrinkRequest) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ListCasesShrinkRequest) GetMaxTalkTurns() *int64 {
	return s.MaxTalkTurns
}

func (s *ListCasesShrinkRequest) GetMinRingingDuration() *int64 {
	return s.MinRingingDuration
}

func (s *ListCasesShrinkRequest) GetMinTalkTime() *int64 {
	return s.MinTalkTime
}

func (s *ListCasesShrinkRequest) GetMinTalkTurns() *int64 {
	return s.MinTalkTurns
}

func (s *ListCasesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCasesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCasesShrinkRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListCasesShrinkRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListCasesShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCasesShrinkRequest) GetStatesShrink() *string {
	return s.StatesShrink
}

func (s *ListCasesShrinkRequest) SetAccessChannelId(v string) *ListCasesShrinkRequest {
	s.AccessChannelId = &v
	return s
}

func (s *ListCasesShrinkRequest) SetAccessChannelType(v string) *ListCasesShrinkRequest {
	s.AccessChannelType = &v
	return s
}

func (s *ListCasesShrinkRequest) SetCaller(v string) *ListCasesShrinkRequest {
	s.Caller = &v
	return s
}

func (s *ListCasesShrinkRequest) SetCampaignId(v string) *ListCasesShrinkRequest {
	s.CampaignId = &v
	return s
}

func (s *ListCasesShrinkRequest) SetCaseCompleted(v bool) *ListCasesShrinkRequest {
	s.CaseCompleted = &v
	return s
}

func (s *ListCasesShrinkRequest) SetCaseIdsShrink(v string) *ListCasesShrinkRequest {
	s.CaseIdsShrink = &v
	return s
}

func (s *ListCasesShrinkRequest) SetDispositionCodesShrink(v string) *ListCasesShrinkRequest {
	s.DispositionCodesShrink = &v
	return s
}

func (s *ListCasesShrinkRequest) SetDispositionReasonsShrink(v string) *ListCasesShrinkRequest {
	s.DispositionReasonsShrink = &v
	return s
}

func (s *ListCasesShrinkRequest) SetDraftVersion(v bool) *ListCasesShrinkRequest {
	s.DraftVersion = &v
	return s
}

func (s *ListCasesShrinkRequest) SetEndTime(v int64) *ListCasesShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListCasesShrinkRequest) SetInstanceId(v string) *ListCasesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCasesShrinkRequest) SetLabelSearchShrink(v string) *ListCasesShrinkRequest {
	s.LabelSearchShrink = &v
	return s
}

func (s *ListCasesShrinkRequest) SetMaxRingingDuration(v int64) *ListCasesShrinkRequest {
	s.MaxRingingDuration = &v
	return s
}

func (s *ListCasesShrinkRequest) SetMaxTalkTime(v int64) *ListCasesShrinkRequest {
	s.MaxTalkTime = &v
	return s
}

func (s *ListCasesShrinkRequest) SetMaxTalkTurns(v int64) *ListCasesShrinkRequest {
	s.MaxTalkTurns = &v
	return s
}

func (s *ListCasesShrinkRequest) SetMinRingingDuration(v int64) *ListCasesShrinkRequest {
	s.MinRingingDuration = &v
	return s
}

func (s *ListCasesShrinkRequest) SetMinTalkTime(v int64) *ListCasesShrinkRequest {
	s.MinTalkTime = &v
	return s
}

func (s *ListCasesShrinkRequest) SetMinTalkTurns(v int64) *ListCasesShrinkRequest {
	s.MinTalkTurns = &v
	return s
}

func (s *ListCasesShrinkRequest) SetPageNumber(v int32) *ListCasesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCasesShrinkRequest) SetPageSize(v int32) *ListCasesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListCasesShrinkRequest) SetPhoneNumber(v string) *ListCasesShrinkRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ListCasesShrinkRequest) SetScriptId(v string) *ListCasesShrinkRequest {
	s.ScriptId = &v
	return s
}

func (s *ListCasesShrinkRequest) SetStartTime(v int64) *ListCasesShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListCasesShrinkRequest) SetStatesShrink(v string) *ListCasesShrinkRequest {
	s.StatesShrink = &v
	return s
}

func (s *ListCasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
