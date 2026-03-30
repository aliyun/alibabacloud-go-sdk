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
}

type GetCallDetailRecordResponseBody struct {
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCallDetailRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// a2c26e67-5984-4935-984e-bcee52971993
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *GetCallDetailRecordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCallDetailRecordResponseBodyData struct {
	// example:
	//
	// 15100000001
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 15929990193
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// Answered
	DispositionCode *string `json:"DispositionCode,omitempty" xml:"DispositionCode,omitempty"`
	// example:
	//
	// ConcurrentLimitExceeded
	DispositionReason *string `json:"DispositionReason,omitempty" xml:"DispositionReason,omitempty"`
	// example:
	//
	// 160
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 1752283603978
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// false
	IssueResolved *bool `json:"IssueResolved,omitempty" xml:"IssueResolved,omitempty"`
	// example:
	//
	// Customer
	ReleaseInitiator   *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ResolutionEvidence *string `json:"ResolutionEvidence,omitempty" xml:"ResolutionEvidence,omitempty"`
	// example:
	//
	// 07d72ea0-6e38-48d4-8371-14deaaba996f
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1774858849987
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 100
	TalkTime *int32 `json:"TalkTime,omitempty" xml:"TalkTime,omitempty"`
	// example:
	//
	// 5
	TalkTurns   *int32                                            `json:"TalkTurns,omitempty" xml:"TalkTurns,omitempty"`
	Transcripts []*GetCallDetailRecordResponseBodyDataTranscripts `json:"Transcripts,omitempty" xml:"Transcripts,omitempty" type:"Repeated"`
	// example:
	//
	// default@testInstance1
	TransferTarget *string `json:"TransferTarget,omitempty" xml:"TransferTarget,omitempty"`
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

func (s *GetCallDetailRecordResponseBodyData) GetCallee() *string {
	return s.Callee
}

func (s *GetCallDetailRecordResponseBodyData) GetCaller() *string {
	return s.Caller
}

func (s *GetCallDetailRecordResponseBodyData) GetDispositionCode() *string {
	return s.DispositionCode
}

func (s *GetCallDetailRecordResponseBodyData) GetDispositionReason() *string {
	return s.DispositionReason
}

func (s *GetCallDetailRecordResponseBodyData) GetDuration() *int32 {
	return s.Duration
}

func (s *GetCallDetailRecordResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetCallDetailRecordResponseBodyData) GetIssueResolved() *bool {
	return s.IssueResolved
}

func (s *GetCallDetailRecordResponseBodyData) GetReleaseInitiator() *string {
	return s.ReleaseInitiator
}

func (s *GetCallDetailRecordResponseBodyData) GetResolutionEvidence() *string {
	return s.ResolutionEvidence
}

func (s *GetCallDetailRecordResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *GetCallDetailRecordResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetCallDetailRecordResponseBodyData) GetTalkTime() *int32 {
	return s.TalkTime
}

func (s *GetCallDetailRecordResponseBodyData) GetTalkTurns() *int32 {
	return s.TalkTurns
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

func (s *GetCallDetailRecordResponseBodyData) SetCallee(v string) *GetCallDetailRecordResponseBodyData {
	s.Callee = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCaller(v string) *GetCallDetailRecordResponseBodyData {
	s.Caller = &v
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

func (s *GetCallDetailRecordResponseBodyData) SetDuration(v int32) *GetCallDetailRecordResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetEndTime(v int64) *GetCallDetailRecordResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetIssueResolved(v bool) *GetCallDetailRecordResponseBodyData {
	s.IssueResolved = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetReleaseInitiator(v string) *GetCallDetailRecordResponseBodyData {
	s.ReleaseInitiator = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetResolutionEvidence(v string) *GetCallDetailRecordResponseBodyData {
	s.ResolutionEvidence = &v
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

func (s *GetCallDetailRecordResponseBodyData) SetTalkTime(v int32) *GetCallDetailRecordResponseBodyData {
	s.TalkTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetTalkTurns(v int32) *GetCallDetailRecordResponseBodyData {
	s.TalkTurns = &v
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

type GetCallDetailRecordResponseBodyDataTranscripts struct {
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// false
	Backchannels *bool `json:"Backchannels,omitempty" xml:"Backchannels,omitempty"`
	// example:
	//
	// 1748945414222
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// [{"endSilenceTimeout":500,"interruptible":false,"silenceDetectionTimeout":5,"type":"Voice"}]
	ControlParamsList *string `json:"ControlParamsList,omitempty" xml:"ControlParamsList,omitempty"`
	// example:
	//
	// 1767315903531
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1774859592165
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// example:
	//
	// [{"dataType":"Silence"}]
	Extras *string `json:"Extras,omitempty" xml:"Extras,omitempty"`
	// example:
	//
	// false
	Interrupted *bool `json:"Interrupted,omitempty" xml:"Interrupted,omitempty"`
	// example:
	//
	// false
	Legacy *bool `json:"Legacy,omitempty" xml:"Legacy,omitempty"`
	// example:
	//
	// 10
	PlayedWords *int32 `json:"PlayedWords,omitempty" xml:"PlayedWords,omitempty"`
	// example:
	//
	// USER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// true
	StreamEnd *bool `json:"StreamEnd,omitempty" xml:"StreamEnd,omitempty"`
	// example:
	//
	// 1
	StreamId  *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
	// example:
	//
	// {"test1":"1234"}
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

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetEventTime() *string {
	return s.EventTime
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetExtras() *string {
	return s.Extras
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetInterrupted() *bool {
	return s.Interrupted
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetLegacy() *bool {
	return s.Legacy
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetPlayedWords() *int32 {
	return s.PlayedWords
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetRole() *string {
	return s.Role
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetStreamEnd() *bool {
	return s.StreamEnd
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) GetStreamId() *string {
	return s.StreamId
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

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetEventTime(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.EventTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetExtras(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.Extras = &v
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

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetPlayedWords(v int32) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.PlayedWords = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetRole(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.Role = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetStreamEnd(v bool) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.StreamEnd = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataTranscripts) SetStreamId(v string) *GetCallDetailRecordResponseBodyDataTranscripts {
	s.StreamId = &v
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
