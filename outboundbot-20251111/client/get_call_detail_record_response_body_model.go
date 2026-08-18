// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallDetailRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCallDetailRecordResponseBody
	GetCode() *string
	SetData(v *GetCallDetailRecordResponseBodyData) *GetCallDetailRecordResponseBody
	GetData() *GetCallDetailRecordResponseBodyData
	SetHttpStatusCode(v int32) *GetCallDetailRecordResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCallDetailRecordResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetCallDetailRecordResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetCallDetailRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCallDetailRecordResponseBody
	GetSuccess() *bool
}

type GetCallDetailRecordResponseBody struct {
	// The error code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The call detail data.
	Data *GetCallDetailRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCallDetailRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCallDetailRecordResponseBody) GetData() *GetCallDetailRecordResponseBodyData {
	return s.Data
}

func (s *GetCallDetailRecordResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCallDetailRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCallDetailRecordResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetCallDetailRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCallDetailRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCallDetailRecordResponseBody) SetCode(v string) *GetCallDetailRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetData(v *GetCallDetailRecordResponseBodyData) *GetCallDetailRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetHttpStatusCode(v int32) *GetCallDetailRecordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetMessage(v string) *GetCallDetailRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetParams(v []*string) *GetCallDetailRecordResponseBody {
	s.Params = v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetRequestId(v string) *GetCallDetailRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetSuccess(v bool) *GetCallDetailRecordResponseBody {
	s.Success = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCallDetailRecordResponseBodyData struct {
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
	// The callee number.
	//
	// example:
	//
	// 13612617599
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// The caller number.
	//
	// example:
	//
	// 02867871030
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The case ID.
	//
	// example:
	//
	// 6fb3a6c4-c3e6-4722-8c71-e5fde4e2253e
	CaseId *string `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
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
	// Indicates whether this is a draft version.
	//
	// example:
	//
	// true
	DraftVersion *bool `json:"DraftVersion,omitempty" xml:"DraftVersion,omitempty"`
	// The total duration.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The end time of the call.
	//
	// example:
	//
	// 1786960840667
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of labels.
	Labels []*GetCallDetailRecordResponseBodyDataLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The party that initiated the hang-up.
	//
	// example:
	//
	// Customer
	ReleaseInitiator *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	// The call session ID.
	//
	// example:
	//
	// job-0b84bf6f-73dc-4462-bd8f-916e3a34c419
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The start time of the call.
	//
	// example:
	//
	// 1786960840667
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The talk time.
	//
	// example:
	//
	// 1
	TalkTime *int64 `json:"TalkTime,omitempty" xml:"TalkTime,omitempty"`
	// The number of conversation turns.
	//
	// example:
	//
	// 1
	TalkTurns *int64 `json:"TalkTurns,omitempty" xml:"TalkTurns,omitempty"`
	// Indicates whether the task was completed.
	//
	// example:
	//
	// false
	TaskCompleted *bool `json:"TaskCompleted,omitempty" xml:"TaskCompleted,omitempty"`
	// The conversation transcripts.
	Transcripts []*GetCallDetailRecordResponseBodyDataTranscripts `json:"Transcripts,omitempty" xml:"Transcripts,omitempty" type:"Repeated"`
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

func (s GetCallDetailRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyData) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *GetCallDetailRecordResponseBodyData) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *GetCallDetailRecordResponseBodyData) GetCallee() *string {
	return s.Callee
}

func (s *GetCallDetailRecordResponseBodyData) GetCaller() *string {
	return s.Caller
}

func (s *GetCallDetailRecordResponseBodyData) GetCaseId() *string {
	return s.CaseId
}

func (s *GetCallDetailRecordResponseBodyData) GetDispositionCode() *string {
	return s.DispositionCode
}

func (s *GetCallDetailRecordResponseBodyData) GetDispositionReason() *string {
	return s.DispositionReason
}

func (s *GetCallDetailRecordResponseBodyData) GetDraftVersion() *bool {
	return s.DraftVersion
}

func (s *GetCallDetailRecordResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *GetCallDetailRecordResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetCallDetailRecordResponseBodyData) GetLabels() []*GetCallDetailRecordResponseBodyDataLabels {
	return s.Labels
}

func (s *GetCallDetailRecordResponseBodyData) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *GetCallDetailRecordResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *GetCallDetailRecordResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetCallDetailRecordResponseBodyData) GetTalkTime() *int64 {
	return s.TalkTime
}

func (s *GetCallDetailRecordResponseBodyData) GetTalkTurns() *int64 {
	return s.TalkTurns
}

func (s *GetCallDetailRecordResponseBodyData) GetTaskCompleted() *bool {
	return s.TaskCompleted
}

func (s *GetCallDetailRecordResponseBodyData) GetTranscripts() []*GetCallDetailRecordResponseBodyDataTranscripts {
	return s.Transcripts
}

func (s *GetCallDetailRecordResponseBodyData) GetTransferTarget() *string {
	return s.TransferTarget
}

func (s *GetCallDetailRecordResponseBodyData) GetTransferType() *string {
	return s.TransferType
}

func (s *GetCallDetailRecordResponseBodyData) SetAccessChannelId(v string) *GetCallDetailRecordResponseBodyData {
	s.AccessChannelId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetAccessChannelType(v string) *GetCallDetailRecordResponseBodyData {
	s.AccessChannelType = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCallee(v string) *GetCallDetailRecordResponseBodyData {
	s.Callee = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCaller(v string) *GetCallDetailRecordResponseBodyData {
	s.Caller = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCaseId(v string) *GetCallDetailRecordResponseBodyData {
	s.CaseId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetDispositionCode(v string) *GetCallDetailRecordResponseBodyData {
	s.DispositionCode = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetDispositionReason(v string) *GetCallDetailRecordResponseBodyData {
	s.DispositionReason = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetDraftVersion(v bool) *GetCallDetailRecordResponseBodyData {
	s.DraftVersion = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetDuration(v int64) *GetCallDetailRecordResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetEndTime(v int64) *GetCallDetailRecordResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetLabels(v []*GetCallDetailRecordResponseBodyDataLabels) *GetCallDetailRecordResponseBodyData {
	s.Labels = v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetReleaseInitiator(v string) *GetCallDetailRecordResponseBodyData {
	s.ReleaseInitiator = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetSessionId(v string) *GetCallDetailRecordResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetStartTime(v int64) *GetCallDetailRecordResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetTalkTime(v int64) *GetCallDetailRecordResponseBodyData {
	s.TalkTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetTalkTurns(v int64) *GetCallDetailRecordResponseBodyData {
	s.TalkTurns = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetTaskCompleted(v bool) *GetCallDetailRecordResponseBodyData {
	s.TaskCompleted = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetTranscripts(v []*GetCallDetailRecordResponseBodyDataTranscripts) *GetCallDetailRecordResponseBodyData {
	s.Transcripts = v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetTransferTarget(v string) *GetCallDetailRecordResponseBodyData {
	s.TransferTarget = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetTransferType(v string) *GetCallDetailRecordResponseBodyData {
	s.TransferType = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Transcripts != nil {
		for _, item := range s.Transcripts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCallDetailRecordResponseBodyDataLabels struct {
	// The set of preset values for the label.
	CandidateValues []*string `json:"CandidateValues,omitempty" xml:"CandidateValues,omitempty" type:"Repeated"`
	// Indicates whether the label has been collected.
	//
	// example:
	//
	// false
	Collected *bool `json:"Collected,omitempty" xml:"Collected,omitempty"`
	// The label description.
	//
	// example:
	//
	// batch_install_docker
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The matched value.
	//
	// example:
	//
	// 123
	MatchedValue *string `json:"MatchedValue,omitempty" xml:"MatchedValue,omitempty"`
	// The label name.
	//
	// example:
	//
	// MemberCollection_20251215_161122_Copy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the label is a system label.
	//
	// example:
	//
	// false
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataLabels) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataLabels) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataLabels) GetCandidateValues() []*string {
	return s.CandidateValues
}

func (s *GetCallDetailRecordResponseBodyDataLabels) GetCollected() *bool {
	return s.Collected
}

func (s *GetCallDetailRecordResponseBodyDataLabels) GetDescription() *string {
	return s.Description
}

func (s *GetCallDetailRecordResponseBodyDataLabels) GetMatchedValue() *string {
	return s.MatchedValue
}

func (s *GetCallDetailRecordResponseBodyDataLabels) GetName() *string {
	return s.Name
}

func (s *GetCallDetailRecordResponseBodyDataLabels) GetSystem() *bool {
	return s.System
}

func (s *GetCallDetailRecordResponseBodyDataLabels) SetCandidateValues(v []*string) *GetCallDetailRecordResponseBodyDataLabels {
	s.CandidateValues = v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataLabels) SetCollected(v bool) *GetCallDetailRecordResponseBodyDataLabels {
	s.Collected = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataLabels) SetDescription(v string) *GetCallDetailRecordResponseBodyDataLabels {
	s.Description = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataLabels) SetMatchedValue(v string) *GetCallDetailRecordResponseBodyDataLabels {
	s.MatchedValue = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataLabels) SetName(v string) *GetCallDetailRecordResponseBodyDataLabels {
	s.Name = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataLabels) SetSystem(v bool) *GetCallDetailRecordResponseBodyDataLabels {
	s.System = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataLabels) Validate() error {
	return dara.Validate(s)
}

type GetCallDetailRecordResponseBodyDataTranscripts struct {
	// The assistant answer.
	//
	// example:
	//
	// Sorry, I cannot find your phone bill
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// Indicates whether the transcript is a backchannel response.
	//
	// example:
	//
	// false
	Backchannels *bool `json:"Backchannels,omitempty" xml:"Backchannels,omitempty"`
	// The begin time.
	//
	// example:
	//
	// 1760667651655
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The list of control parameters.
	//
	// example:
	//
	// {"transferCode":"Transfer02","type":"Transfer"}]
	ControlParamsList *string `json:"ControlParamsList,omitempty" xml:"ControlParamsList,omitempty"`
	// The end time of the call.
	//
	// example:
	//
	// 1760667651655
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event time.
	//
	// example:
	//
	// 1786960840667
	EventTime *int64 `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// The extended information.
	//
	// example:
	//
	// {}
	Extras *string `json:"Extras,omitempty" xml:"Extras,omitempty"`
	// The number of input tokens.
	//
	// example:
	//
	// 0
	InputTokens *int32 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// Indicates whether the response was interrupted.
	//
	// example:
	//
	// false
	Interrupted *bool `json:"Interrupted,omitempty" xml:"Interrupted,omitempty"`
	// Indicates whether the transcript is from the legacy version.
	//
	// example:
	//
	// false
	Legacy *bool `json:"Legacy,omitempty" xml:"Legacy,omitempty"`
	// The model.
	//
	// example:
	//
	// model1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The number of output tokens.
	//
	// example:
	//
	// 0
	OutputTokens *int32 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// The played text.
	//
	// example:
	//
	// Sorry, check
	PlayedWords *string `json:"PlayedWords,omitempty" xml:"PlayedWords,omitempty"`
	// The role.
	//
	// example:
	//
	// USER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The call session ID.
	//
	// example:
	//
	// job-0b84bf6f-73dc-4462-bd8f-916e3a34c419
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The stream ID.
	//
	// example:
	//
	// 1
	StreamId *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
	// The total number of tokens.
	//
	// example:
	//
	// 0
	TotalTokens *int32 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
	// The user utterance.
	//
	// example:
	//
	// Can you check my phone bill?
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	// The vendor parameters.
	//
	// example:
	//
	// {}
	VendorParams *string `json:"VendorParams,omitempty" xml:"VendorParams,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataTranscripts) String() string {
	return dara.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataTranscripts) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetAnswer() *string {
	return s.Answer
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetBackchannels() *bool {
	return s.Backchannels
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetControlParamsList() *string {
	return s.ControlParamsList
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetEventTime() *int64 {
	return s.EventTime
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetExtras() *string {
	return s.Extras
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetInterrupted() *bool {
	return s.Interrupted
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetLegacy() *bool {
	return s.Legacy
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetModel() *string {
	return s.Model
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetPlayedWords() *string {
	return s.PlayedWords
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetRole() *string {
	return s.Role
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetSessionId() *string {
	return s.SessionId
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetStreamId() *string {
	return s.StreamId
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetUtterance() *string {
	return s.Utterance
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetVendorParams() *string {
	return s.VendorParams
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetAnswer(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.Answer = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetBackchannels(v bool) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.Backchannels = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetBeginTime(v int64) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.BeginTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetControlParamsList(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.ControlParamsList = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetEndTime(v int64) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.EndTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetEventTime(v int64) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.EventTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetExtras(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.Extras = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetInputTokens(v int32) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.InputTokens = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetInterrupted(v bool) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.Interrupted = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetLegacy(v bool) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.Legacy = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetModel(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.Model = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetOutputTokens(v int32) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.OutputTokens = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetPlayedWords(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.PlayedWords = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetRole(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.Role = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetSessionId(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.SessionId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetStreamId(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.StreamId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetTotalTokens(v int32) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.TotalTokens = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetUtterance(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.Utterance = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetVendorParams(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.VendorParams = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) Validate() error {
	return dara.Validate(s)
}
