// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNextResultToVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNextResultToVerifyResponseBody
	GetCode() *string
	SetData(v *GetNextResultToVerifyResponseBodyData) *GetNextResultToVerifyResponseBody
	GetData() *GetNextResultToVerifyResponseBodyData
	SetMessage(v string) *GetNextResultToVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNextResultToVerifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNextResultToVerifyResponseBody
	GetSuccess() *bool
}

type GetNextResultToVerifyResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetNextResultToVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNextResultToVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNextResultToVerifyResponseBody) GetData() *GetNextResultToVerifyResponseBodyData {
	return s.Data
}

func (s *GetNextResultToVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNextResultToVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNextResultToVerifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNextResultToVerifyResponseBody) SetCode(v string) *GetNextResultToVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetData(v *GetNextResultToVerifyResponseBodyData) *GetNextResultToVerifyResponseBody {
	s.Data = v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetMessage(v string) *GetNextResultToVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetRequestId(v string) *GetNextResultToVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNextResultToVerifyResponseBody) SetSuccess(v bool) *GetNextResultToVerifyResponseBody {
	s.Success = &v
	return s
}

func (s *GetNextResultToVerifyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNextResultToVerifyResponseBodyData struct {
	// example:
	//
	// http
	AudioScheme *string `json:"AudioScheme,omitempty" xml:"AudioScheme,omitempty"`
	// example:
	//
	// sca-bucket.oss-cn-hangzhou.aliyuncs.com/upload_1173636551461420/dateset_1584674455133_SzC/%E4%BA%BA%E5%B7%A5%E6%A0%A1%E9%AA%8C%E6%B5%8B%E8%AF%95-%E6%9F%A5%E5%8C%97%E4%BA%AC%E5%A4%A9%E6%B0%94.wav?Expires=1584847372&amp;OSSAccessKeyId=*****&amp;Signature=HccAKnLOJwoYvzE*********
	AudioURL  *string                                         `json:"AudioURL,omitempty" xml:"AudioURL,omitempty"`
	Dialogues *GetNextResultToVerifyResponseBodyDataDialogues `json:"Dialogues,omitempty" xml:"Dialogues,omitempty" type:"Struct"`
	// example:
	//
	// 23421
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// e790e6c919d84b82b64ee*****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// xxx.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 23
	IncorrectWords *int32 `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	// example:
	//
	// 2
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// 0.97079998
	Precision *float32 `json:"Precision,omitempty" xml:"Precision,omitempty"`
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2020-03-20T11:26Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// true
	Verified *bool `json:"Verified,omitempty" xml:"Verified,omitempty"`
	// example:
	//
	// 2
	VerifiedCount *int32 `json:"VerifiedCount,omitempty" xml:"VerifiedCount,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyData) GetAudioScheme() *string {
	return s.AudioScheme
}

func (s *GetNextResultToVerifyResponseBodyData) GetAudioURL() *string {
	return s.AudioURL
}

func (s *GetNextResultToVerifyResponseBodyData) GetDialogues() *GetNextResultToVerifyResponseBodyDataDialogues {
	return s.Dialogues
}

func (s *GetNextResultToVerifyResponseBodyData) GetDuration() *int32 {
	return s.Duration
}

func (s *GetNextResultToVerifyResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *GetNextResultToVerifyResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetNextResultToVerifyResponseBodyData) GetIncorrectWords() *int32 {
	return s.IncorrectWords
}

func (s *GetNextResultToVerifyResponseBodyData) GetIndex() *int32 {
	return s.Index
}

func (s *GetNextResultToVerifyResponseBodyData) GetPrecision() *float32 {
	return s.Precision
}

func (s *GetNextResultToVerifyResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetNextResultToVerifyResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetNextResultToVerifyResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetNextResultToVerifyResponseBodyData) GetVerified() *bool {
	return s.Verified
}

func (s *GetNextResultToVerifyResponseBodyData) GetVerifiedCount() *int32 {
	return s.VerifiedCount
}

func (s *GetNextResultToVerifyResponseBodyData) SetAudioScheme(v string) *GetNextResultToVerifyResponseBodyData {
	s.AudioScheme = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetAudioURL(v string) *GetNextResultToVerifyResponseBodyData {
	s.AudioURL = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetDialogues(v *GetNextResultToVerifyResponseBodyDataDialogues) *GetNextResultToVerifyResponseBodyData {
	s.Dialogues = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetDuration(v int32) *GetNextResultToVerifyResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetFileId(v string) *GetNextResultToVerifyResponseBodyData {
	s.FileId = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetFileName(v string) *GetNextResultToVerifyResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetIncorrectWords(v int32) *GetNextResultToVerifyResponseBodyData {
	s.IncorrectWords = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetIndex(v int32) *GetNextResultToVerifyResponseBodyData {
	s.Index = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetPrecision(v float32) *GetNextResultToVerifyResponseBodyData {
	s.Precision = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetStatus(v int32) *GetNextResultToVerifyResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetTotalCount(v int32) *GetNextResultToVerifyResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetUpdateTime(v string) *GetNextResultToVerifyResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetVerified(v bool) *GetNextResultToVerifyResponseBodyData {
	s.Verified = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) SetVerifiedCount(v int32) *GetNextResultToVerifyResponseBodyData {
	s.VerifiedCount = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyData) Validate() error {
	if s.Dialogues != nil {
		if err := s.Dialogues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNextResultToVerifyResponseBodyDataDialogues struct {
	Dialogue []*GetNextResultToVerifyResponseBodyDataDialoguesDialogue `json:"Dialogue,omitempty" xml:"Dialogue,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialogues) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialogues) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialogues) GetDialogue() []*GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	return s.Dialogue
}

func (s *GetNextResultToVerifyResponseBodyDataDialogues) SetDialogue(v []*GetNextResultToVerifyResponseBodyDataDialoguesDialogue) *GetNextResultToVerifyResponseBodyDataDialogues {
	s.Dialogue = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialogues) Validate() error {
	if s.Dialogue != nil {
		for _, item := range s.Dialogue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogue struct {
	// example:
	//
	// 980
	Begin *int64 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// XXX
	BeginTime *string                                                       `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	Deltas    *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas `json:"Deltas,omitempty" xml:"Deltas,omitempty" type:"Struct"`
	// example:
	//
	// 6
	EmotionValue *int32 `json:"EmotionValue,omitempty" xml:"EmotionValue,omitempty"`
	// example:
	//
	// 3422
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 00:00:07
	HourMinSec *string `json:"HourMinSec,omitempty" xml:"HourMinSec,omitempty"`
	Identity   *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	// example:
	//
	// 2
	IncorrectWords *int32  `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	Role           *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 1
	SilenceDuration *int32  `json:"SilenceDuration,omitempty" xml:"SilenceDuration,omitempty"`
	SourceRole      *string `json:"SourceRole,omitempty" xml:"SourceRole,omitempty"`
	SourceWords     *string `json:"SourceWords,omitempty" xml:"SourceWords,omitempty"`
	// example:
	//
	// 332
	SpeechRate *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Words      *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogue) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetBegin() *int64 {
	return s.Begin
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetDeltas() *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas {
	return s.Deltas
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetEnd() *int64 {
	return s.End
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetHourMinSec() *string {
	return s.HourMinSec
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetIdentity() *string {
	return s.Identity
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetIncorrectWords() *int32 {
	return s.IncorrectWords
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetRole() *string {
	return s.Role
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetSilenceDuration() *int32 {
	return s.SilenceDuration
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetSourceRole() *string {
	return s.SourceRole
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetSourceWords() *string {
	return s.SourceWords
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) GetWords() *string {
	return s.Words
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetBegin(v int64) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Begin = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetBeginTime(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.BeginTime = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetDeltas(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Deltas = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetEmotionValue(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.EmotionValue = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetEnd(v int64) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.End = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetHourMinSec(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.HourMinSec = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetIdentity(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Identity = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetIncorrectWords(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.IncorrectWords = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetRole(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Role = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSilenceDuration(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SilenceDuration = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSourceRole(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SourceRole = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSourceWords(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SourceWords = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetSpeechRate(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.SpeechRate = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) SetWords(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogue {
	s.Words = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogue) Validate() error {
	if s.Deltas != nil {
		if err := s.Deltas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas struct {
	Delta []*GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta `json:"Delta,omitempty" xml:"Delta,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) GetDelta() []*GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta {
	return s.Delta
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) SetDelta(v []*GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas {
	s.Delta = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltas) Validate() error {
	if s.Delta != nil {
		for _, item := range s.Delta {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta struct {
	Source *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Target *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// CHANGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) GetSource() *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource {
	return s.Source
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) GetTarget() *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget {
	return s.Target
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) GetType() *string {
	return s.Type
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) SetSource(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta {
	s.Source = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) SetTarget(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta {
	s.Target = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) SetType(v string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta {
	s.Type = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDelta) Validate() error {
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource struct {
	Line *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	// example:
	//
	// 5
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) GetLine() *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine {
	return s.Line
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) GetPosition() *int32 {
	return s.Position
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) SetLine(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource {
	s.Line = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) SetPosition(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource {
	s.Position = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSource) Validate() error {
	if s.Line != nil {
		if err := s.Line.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) GetLine() []*string {
	return s.Line
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) SetLine(v []*string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine {
	s.Line = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaSourceLine) Validate() error {
	return dara.Validate(s)
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget struct {
	Line *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Struct"`
	// example:
	//
	// 5
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) GetLine() *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine {
	return s.Line
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) GetPosition() *int32 {
	return s.Position
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) SetLine(v *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget {
	s.Line = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) SetPosition(v int32) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget {
	s.Position = &v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTarget) Validate() error {
	if s.Line != nil {
		if err := s.Line.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine struct {
	Line []*string `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) GetLine() []*string {
	return s.Line
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) SetLine(v []*string) *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine {
	s.Line = v
	return s
}

func (s *GetNextResultToVerifyResponseBodyDataDialoguesDialogueDeltasDeltaTargetLine) Validate() error {
	return dara.Validate(s)
}
