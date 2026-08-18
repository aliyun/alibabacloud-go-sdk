// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaseDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCaseDetailResponseBody
	GetCode() *string
	SetData(v *GetCaseDetailResponseBodyData) *GetCaseDetailResponseBody
	GetData() *GetCaseDetailResponseBodyData
	SetHttpStatusCode(v int32) *GetCaseDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCaseDetailResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetCaseDetailResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetCaseDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCaseDetailResponseBody
	GetSuccess() *bool
}

type GetCaseDetailResponseBody struct {
	// The error code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The case details data.
	Data *GetCaseDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pass-through parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 91102861-AEB9-56C5-8F3A-A023A0E8B5F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCaseDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCaseDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCaseDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCaseDetailResponseBody) GetData() *GetCaseDetailResponseBodyData {
	return s.Data
}

func (s *GetCaseDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCaseDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCaseDetailResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetCaseDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCaseDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCaseDetailResponseBody) SetCode(v string) *GetCaseDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetCaseDetailResponseBody) SetData(v *GetCaseDetailResponseBodyData) *GetCaseDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetCaseDetailResponseBody) SetHttpStatusCode(v int32) *GetCaseDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCaseDetailResponseBody) SetMessage(v string) *GetCaseDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetCaseDetailResponseBody) SetParams(v []*string) *GetCaseDetailResponseBody {
	s.Params = v
	return s
}

func (s *GetCaseDetailResponseBody) SetRequestId(v string) *GetCaseDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCaseDetailResponseBody) SetSuccess(v bool) *GetCaseDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetCaseDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCaseDetailResponseBodyData struct {
	// The list of associated call detail records.
	CallDetailRecords []*GetCaseDetailResponseBodyDataCallDetailRecords `json:"CallDetailRecords,omitempty" xml:"CallDetailRecords,omitempty" type:"Repeated"`
	// The case information.
	Case *GetCaseDetailResponseBodyDataCase `json:"Case,omitempty" xml:"Case,omitempty" type:"Struct"`
}

func (s GetCaseDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCaseDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCaseDetailResponseBodyData) GetCallDetailRecords() []*GetCaseDetailResponseBodyDataCallDetailRecords {
	return s.CallDetailRecords
}

func (s *GetCaseDetailResponseBodyData) GetCase() *GetCaseDetailResponseBodyDataCase {
	return s.Case
}

func (s *GetCaseDetailResponseBodyData) SetCallDetailRecords(v []*GetCaseDetailResponseBodyDataCallDetailRecords) *GetCaseDetailResponseBodyData {
	s.CallDetailRecords = v
	return s
}

func (s *GetCaseDetailResponseBodyData) SetCase(v *GetCaseDetailResponseBodyDataCase) *GetCaseDetailResponseBodyData {
	s.Case = v
	return s
}

func (s *GetCaseDetailResponseBodyData) Validate() error {
	if s.CallDetailRecords != nil {
		for _, item := range s.CallDetailRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Case != nil {
		if err := s.Case.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCaseDetailResponseBodyDataCallDetailRecords struct {
	// The access channel ID.
	//
	// example:
	//
	// 2957188c-6cb3-42b4-beca-906fc0e752e5
	AccessChannelId *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	// The access channel type.
	//
	// example:
	//
	// Test
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	// The called number.
	//
	// example:
	//
	// 13510595079
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// The caller number.
	//
	// example:
	//
	// 02162300961
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The disposition code.
	//
	// example:
	//
	// Answered
	DispositionCode *string `json:"DispositionCode,omitempty" xml:"DispositionCode,omitempty"`
	// The disposition reason.
	//
	// example:
	//
	// Normal
	DispositionReason *string `json:"DispositionReason,omitempty" xml:"DispositionReason,omitempty"`
	// Indicates whether the version is a draft version.
	//
	// example:
	//
	// true
	DraftVersion *bool `json:"DraftVersion,omitempty" xml:"DraftVersion,omitempty"`
	// The total duration.
	//
	// example:
	//
	// 20
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The time when the call ended.
	//
	// example:
	//
	// 1786960840667
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The party that initiated the hangup.
	//
	// example:
	//
	// Customer
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The call session ID.
	//
	// example:
	//
	// job-893f8715-3658-4488-8cf0-6a8546124f96
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The time when the call started.
	//
	// example:
	//
	// 1786960840667
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The talk time.
	//
	// example:
	//
	// 10
	TalkTime *int64 `json:"TalkTime,omitempty" xml:"TalkTime,omitempty"`
	// The number of conversation turns.
	//
	// example:
	//
	// 2
	TalkTurns *int64 `json:"TalkTurns,omitempty" xml:"TalkTurns,omitempty"`
	// The transfer target.
	//
	// example:
	//
	// SkillGroup1
	TransferTarget *string `json:"TransferTarget,omitempty" xml:"TransferTarget,omitempty"`
	// The transfer type.
	//
	// example:
	//
	// SkillGroup
	TransferType *string `json:"TransferType,omitempty" xml:"TransferType,omitempty"`
}

func (s GetCaseDetailResponseBodyDataCallDetailRecords) String() string {
	return dara.Prettify(s)
}

func (s GetCaseDetailResponseBodyDataCallDetailRecords) GoString() string {
	return s.String()
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetCallee() *string {
	return s.Callee
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetCaller() *string {
	return s.Caller
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetDispositionCode() *string {
	return s.DispositionCode
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetDispositionReason() *string {
	return s.DispositionReason
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetDraftVersion() *bool {
	return s.DraftVersion
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetDuration() *int64 {
	return s.Duration
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetSessionId() *string {
	return s.SessionId
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetTalkTime() *int64 {
	return s.TalkTime
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetTalkTurns() *int64 {
	return s.TalkTurns
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetTransferTarget() *string {
	return s.TransferTarget
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) GetTransferType() *string {
	return s.TransferType
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetAccessChannelId(v string) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.AccessChannelId = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetAccessChannelType(v string) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.AccessChannelType = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetCallee(v string) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.Callee = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetCaller(v string) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.Caller = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetDispositionCode(v string) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.DispositionCode = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetDispositionReason(v string) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.DispositionReason = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetDraftVersion(v bool) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.DraftVersion = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetDuration(v int64) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.Duration = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetEndTime(v int64) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.EndTime = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetReleaseInitiator(v string) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.ReleaseInitiator = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetSessionId(v string) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.SessionId = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetStartTime(v int64) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.StartTime = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetTalkTime(v int64) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.TalkTime = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetTalkTurns(v int64) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.TalkTurns = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetTransferTarget(v string) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.TransferTarget = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) SetTransferType(v string) *GetCaseDetailResponseBodyDataCallDetailRecords {
	s.TransferType = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCallDetailRecords) Validate() error {
	return dara.Validate(s)
}

type GetCaseDetailResponseBodyDataCase struct {
	// The number of call attempts.
	//
	// example:
	//
	// 1
	AttemptedCount *int32 `json:"AttemptedCount,omitempty" xml:"AttemptedCount,omitempty"`
	// The caller number.
	//
	// example:
	//
	// 0571000018766
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// The ID of the outbound campaign.
	//
	// example:
	//
	// 5b5c7b4a-978e-4937-a192-02f4621bf67e
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The name of the outbound campaign.
	//
	// example:
	//
	// Test campaign
	CampaignName *string `json:"CampaignName,omitempty" xml:"CampaignName,omitempty"`
	// The case ID.
	//
	// example:
	//
	// 893f8715-3658-4488-8cf0-6a8546124f00
	CaseId *string `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-04-13T06:05:54Z
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The custom variables in JSON string format.
	//
	// example:
	//
	// {}
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// The dialing time.
	//
	// example:
	//
	// 1786960840667
	DialingTime *int64 `json:"DialingTime,omitempty" xml:"DialingTime,omitempty"`
	// The disposition code.
	//
	// example:
	//
	// Answered
	DispositionCode *string `json:"DispositionCode,omitempty" xml:"DispositionCode,omitempty"`
	// The disposition reason.
	//
	// example:
	//
	// Normal
	DispositionReason *string `json:"DispositionReason,omitempty" xml:"DispositionReason,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 893f8715-3658-4488-8cf0-6a8546124f96
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of labels.
	Labels []*GetCaseDetailResponseBodyDataCaseLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The time of the last call attempt.
	//
	// example:
	//
	// 1786960840667
	LastAttemptedTime *int64 `json:"LastAttemptedTime,omitempty" xml:"LastAttemptedTime,omitempty"`
	// The called number.
	//
	// example:
	//
	// 18512345678
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The priority.
	//
	// example:
	//
	// Daily
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The reference ID.
	//
	// example:
	//
	// 1529431297649278976
	ReferenceId *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
	// The ringing duration.
	//
	// example:
	//
	// 30
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
	// The ringing time.
	//
	// example:
	//
	// 1786960840667
	RingingTime *int64 `json:"RingingTime,omitempty" xml:"RingingTime,omitempty"`
	// The script ID.
	//
	// example:
	//
	// 64241e64-190c-45d1-af66-06f51c07b090
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The script name.
	//
	// example:
	//
	// XiaoHuan
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// The session ID of the last call.
	//
	// example:
	//
	// job-893f8715-3658-4488-8cf0-6a8546124f96
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The case state.
	//
	// example:
	//
	// Pending
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The talk time.
	//
	// example:
	//
	// 10
	TalkTime *int64 `json:"TalkTime,omitempty" xml:"TalkTime,omitempty"`
}

func (s GetCaseDetailResponseBodyDataCase) String() string {
	return dara.Prettify(s)
}

func (s GetCaseDetailResponseBodyDataCase) GoString() string {
	return s.String()
}

func (s *GetCaseDetailResponseBodyDataCase) GetAttemptedCount() *int32 {
	return s.AttemptedCount
}

func (s *GetCaseDetailResponseBodyDataCase) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *GetCaseDetailResponseBodyDataCase) GetCampaignId() *string {
	return s.CampaignId
}

func (s *GetCaseDetailResponseBodyDataCase) GetCampaignName() *string {
	return s.CampaignName
}

func (s *GetCaseDetailResponseBodyDataCase) GetCaseId() *string {
	return s.CaseId
}

func (s *GetCaseDetailResponseBodyDataCase) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetCaseDetailResponseBodyDataCase) GetCustomVariables() *string {
	return s.CustomVariables
}

func (s *GetCaseDetailResponseBodyDataCase) GetDialingTime() *int64 {
	return s.DialingTime
}

func (s *GetCaseDetailResponseBodyDataCase) GetDispositionCode() *string {
	return s.DispositionCode
}

func (s *GetCaseDetailResponseBodyDataCase) GetDispositionReason() *string {
	return s.DispositionReason
}

func (s *GetCaseDetailResponseBodyDataCase) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCaseDetailResponseBodyDataCase) GetLabels() []*GetCaseDetailResponseBodyDataCaseLabels {
	return s.Labels
}

func (s *GetCaseDetailResponseBodyDataCase) GetLastAttemptedTime() *int64 {
	return s.LastAttemptedTime
}

func (s *GetCaseDetailResponseBodyDataCase) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetCaseDetailResponseBodyDataCase) GetPriority() *int32 {
	return s.Priority
}

func (s *GetCaseDetailResponseBodyDataCase) GetReferenceId() *string {
	return s.ReferenceId
}

func (s *GetCaseDetailResponseBodyDataCase) GetRingingDuration() *int64 {
	return s.RingingDuration
}

func (s *GetCaseDetailResponseBodyDataCase) GetRingingTime() *int64 {
	return s.RingingTime
}

func (s *GetCaseDetailResponseBodyDataCase) GetScriptId() *string {
	return s.ScriptId
}

func (s *GetCaseDetailResponseBodyDataCase) GetScriptName() *string {
	return s.ScriptName
}

func (s *GetCaseDetailResponseBodyDataCase) GetSessionId() *string {
	return s.SessionId
}

func (s *GetCaseDetailResponseBodyDataCase) GetState() *string {
	return s.State
}

func (s *GetCaseDetailResponseBodyDataCase) GetTalkTime() *int64 {
	return s.TalkTime
}

func (s *GetCaseDetailResponseBodyDataCase) SetAttemptedCount(v int32) *GetCaseDetailResponseBodyDataCase {
	s.AttemptedCount = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetCallerNumber(v string) *GetCaseDetailResponseBodyDataCase {
	s.CallerNumber = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetCampaignId(v string) *GetCaseDetailResponseBodyDataCase {
	s.CampaignId = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetCampaignName(v string) *GetCaseDetailResponseBodyDataCase {
	s.CampaignName = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetCaseId(v string) *GetCaseDetailResponseBodyDataCase {
	s.CaseId = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetCreatedTime(v int64) *GetCaseDetailResponseBodyDataCase {
	s.CreatedTime = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetCustomVariables(v string) *GetCaseDetailResponseBodyDataCase {
	s.CustomVariables = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetDialingTime(v int64) *GetCaseDetailResponseBodyDataCase {
	s.DialingTime = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetDispositionCode(v string) *GetCaseDetailResponseBodyDataCase {
	s.DispositionCode = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetDispositionReason(v string) *GetCaseDetailResponseBodyDataCase {
	s.DispositionReason = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetInstanceId(v string) *GetCaseDetailResponseBodyDataCase {
	s.InstanceId = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetLabels(v []*GetCaseDetailResponseBodyDataCaseLabels) *GetCaseDetailResponseBodyDataCase {
	s.Labels = v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetLastAttemptedTime(v int64) *GetCaseDetailResponseBodyDataCase {
	s.LastAttemptedTime = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetPhoneNumber(v string) *GetCaseDetailResponseBodyDataCase {
	s.PhoneNumber = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetPriority(v int32) *GetCaseDetailResponseBodyDataCase {
	s.Priority = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetReferenceId(v string) *GetCaseDetailResponseBodyDataCase {
	s.ReferenceId = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetRingingDuration(v int64) *GetCaseDetailResponseBodyDataCase {
	s.RingingDuration = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetRingingTime(v int64) *GetCaseDetailResponseBodyDataCase {
	s.RingingTime = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetScriptId(v string) *GetCaseDetailResponseBodyDataCase {
	s.ScriptId = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetScriptName(v string) *GetCaseDetailResponseBodyDataCase {
	s.ScriptName = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetSessionId(v string) *GetCaseDetailResponseBodyDataCase {
	s.SessionId = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetState(v string) *GetCaseDetailResponseBodyDataCase {
	s.State = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) SetTalkTime(v int64) *GetCaseDetailResponseBodyDataCase {
	s.TalkTime = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCase) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCaseDetailResponseBodyDataCaseLabels struct {
	// The set of candidate values for the label.
	CandidateValues []*string `json:"CandidateValues,omitempty" xml:"CandidateValues,omitempty" type:"Repeated"`
	// Indicates whether the label was collected.
	//
	// example:
	//
	// true
	Collected *bool `json:"Collected,omitempty" xml:"Collected,omitempty"`
	// The label description.
	//
	// example:
	//
	// 123
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The matched value of the label.
	//
	// example:
	//
	// 123
	MatchedValue *string `json:"MatchedValue,omitempty" xml:"MatchedValue,omitempty"`
	// The label name.
	//
	// example:
	//
	// Contact customer to register acquaintance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the label is a system label.
	//
	// example:
	//
	// false
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
}

func (s GetCaseDetailResponseBodyDataCaseLabels) String() string {
	return dara.Prettify(s)
}

func (s GetCaseDetailResponseBodyDataCaseLabels) GoString() string {
	return s.String()
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) GetCandidateValues() []*string {
	return s.CandidateValues
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) GetCollected() *bool {
	return s.Collected
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) GetDescription() *string {
	return s.Description
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) GetMatchedValue() *string {
	return s.MatchedValue
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) GetName() *string {
	return s.Name
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) GetSystem() *bool {
	return s.System
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) SetCandidateValues(v []*string) *GetCaseDetailResponseBodyDataCaseLabels {
	s.CandidateValues = v
	return s
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) SetCollected(v bool) *GetCaseDetailResponseBodyDataCaseLabels {
	s.Collected = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) SetDescription(v string) *GetCaseDetailResponseBodyDataCaseLabels {
	s.Description = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) SetMatchedValue(v string) *GetCaseDetailResponseBodyDataCaseLabels {
	s.MatchedValue = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) SetName(v string) *GetCaseDetailResponseBodyDataCaseLabels {
	s.Name = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) SetSystem(v bool) *GetCaseDetailResponseBodyDataCaseLabels {
	s.System = &v
	return s
}

func (s *GetCaseDetailResponseBodyDataCaseLabels) Validate() error {
	return dara.Validate(s)
}
