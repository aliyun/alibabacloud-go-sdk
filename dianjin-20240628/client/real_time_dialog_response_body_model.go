// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRealTimeDialogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChoices(v []*RealTimeDialogResponseBodyChoices) *RealTimeDialogResponseBody
	GetChoices() []*RealTimeDialogResponseBodyChoices
	SetCreated(v string) *RealTimeDialogResponseBody
	GetCreated() *string
	SetId(v string) *RealTimeDialogResponseBody
	GetId() *string
	SetRequestId(v string) *RealTimeDialogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RealTimeDialogResponseBody
	GetSuccess() *bool
}

type RealTimeDialogResponseBody struct {
	Choices []*RealTimeDialogResponseBodyChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1735139569523
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RealTimeDialogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RealTimeDialogResponseBody) GoString() string {
	return s.String()
}

func (s *RealTimeDialogResponseBody) GetChoices() []*RealTimeDialogResponseBodyChoices {
	return s.Choices
}

func (s *RealTimeDialogResponseBody) GetCreated() *string {
	return s.Created
}

func (s *RealTimeDialogResponseBody) GetId() *string {
	return s.Id
}

func (s *RealTimeDialogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RealTimeDialogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RealTimeDialogResponseBody) SetChoices(v []*RealTimeDialogResponseBodyChoices) *RealTimeDialogResponseBody {
	s.Choices = v
	return s
}

func (s *RealTimeDialogResponseBody) SetCreated(v string) *RealTimeDialogResponseBody {
	s.Created = &v
	return s
}

func (s *RealTimeDialogResponseBody) SetId(v string) *RealTimeDialogResponseBody {
	s.Id = &v
	return s
}

func (s *RealTimeDialogResponseBody) SetRequestId(v string) *RealTimeDialogResponseBody {
	s.RequestId = &v
	return s
}

func (s *RealTimeDialogResponseBody) SetSuccess(v bool) *RealTimeDialogResponseBody {
	s.Success = &v
	return s
}

func (s *RealTimeDialogResponseBody) Validate() error {
	if s.Choices != nil {
		for _, item := range s.Choices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RealTimeDialogResponseBodyChoices struct {
	Delta *RealTimeDialogResponseBodyChoicesDelta `json:"delta,omitempty" xml:"delta,omitempty" type:"Struct"`
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int32                                    `json:"index,omitempty" xml:"index,omitempty"`
	Message *RealTimeDialogResponseBodyChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s RealTimeDialogResponseBodyChoices) String() string {
	return dara.Prettify(s)
}

func (s RealTimeDialogResponseBodyChoices) GoString() string {
	return s.String()
}

func (s *RealTimeDialogResponseBodyChoices) GetDelta() *RealTimeDialogResponseBodyChoicesDelta {
	return s.Delta
}

func (s *RealTimeDialogResponseBodyChoices) GetFinishReason() *string {
	return s.FinishReason
}

func (s *RealTimeDialogResponseBodyChoices) GetIndex() *int32 {
	return s.Index
}

func (s *RealTimeDialogResponseBodyChoices) GetMessage() *RealTimeDialogResponseBodyChoicesMessage {
	return s.Message
}

func (s *RealTimeDialogResponseBodyChoices) SetDelta(v *RealTimeDialogResponseBodyChoicesDelta) *RealTimeDialogResponseBodyChoices {
	s.Delta = v
	return s
}

func (s *RealTimeDialogResponseBodyChoices) SetFinishReason(v string) *RealTimeDialogResponseBodyChoices {
	s.FinishReason = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoices) SetIndex(v int32) *RealTimeDialogResponseBodyChoices {
	s.Index = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoices) SetMessage(v *RealTimeDialogResponseBodyChoicesMessage) *RealTimeDialogResponseBodyChoices {
	s.Message = v
	return s
}

func (s *RealTimeDialogResponseBodyChoices) Validate() error {
	if s.Delta != nil {
		if err := s.Delta.Validate(); err != nil {
			return err
		}
	}
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RealTimeDialogResponseBodyChoicesDelta struct {
	// example:
	//
	// null
	AnalysisProcess *string `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
	// time
	//
	// example:
	//
	// null
	CallTime *string `json:"callTime,omitempty" xml:"callTime,omitempty"`
	// example:
	//
	// false
	HangUpDialog *bool `json:"hangUpDialog,omitempty" xml:"hangUpDialog,omitempty"`
	// example:
	//
	// 1853360771162058752
	IntentionCode   *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	IntentionName   *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
	Interrupt       *bool   `json:"interrupt,omitempty" xml:"interrupt,omitempty"`
	// example:
	//
	// null
	RecommendIntention *string `json:"recommendIntention,omitempty" xml:"recommendIntention,omitempty"`
	// example:
	//
	// null
	RecommendScript               *string `json:"recommendScript,omitempty" xml:"recommendScript,omitempty"`
	SelfDirectedScript            *string `json:"selfDirectedScript,omitempty" xml:"selfDirectedScript,omitempty"`
	SelfDirectedScriptFullContent *string `json:"selfDirectedScriptFullContent,omitempty" xml:"selfDirectedScriptFullContent,omitempty"`
	SkipCurrentRecognize          *bool   `json:"skipCurrentRecognize,omitempty" xml:"skipCurrentRecognize,omitempty"`
}

func (s RealTimeDialogResponseBodyChoicesDelta) String() string {
	return dara.Prettify(s)
}

func (s RealTimeDialogResponseBodyChoicesDelta) GoString() string {
	return s.String()
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetAnalysisProcess() *string {
	return s.AnalysisProcess
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetCallTime() *string {
	return s.CallTime
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetHangUpDialog() *bool {
	return s.HangUpDialog
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetIntentionCode() *string {
	return s.IntentionCode
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetIntentionName() *string {
	return s.IntentionName
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetIntentionScript() *string {
	return s.IntentionScript
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetInterrupt() *bool {
	return s.Interrupt
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetRecommendIntention() *string {
	return s.RecommendIntention
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetRecommendScript() *string {
	return s.RecommendScript
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetSelfDirectedScript() *string {
	return s.SelfDirectedScript
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetSelfDirectedScriptFullContent() *string {
	return s.SelfDirectedScriptFullContent
}

func (s *RealTimeDialogResponseBodyChoicesDelta) GetSkipCurrentRecognize() *bool {
	return s.SkipCurrentRecognize
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetAnalysisProcess(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.AnalysisProcess = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetCallTime(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.CallTime = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetHangUpDialog(v bool) *RealTimeDialogResponseBodyChoicesDelta {
	s.HangUpDialog = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetIntentionCode(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.IntentionCode = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetIntentionName(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.IntentionName = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetIntentionScript(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.IntentionScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetInterrupt(v bool) *RealTimeDialogResponseBodyChoicesDelta {
	s.Interrupt = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetRecommendIntention(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.RecommendIntention = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetRecommendScript(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.RecommendScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetSelfDirectedScript(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.SelfDirectedScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetSelfDirectedScriptFullContent(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.SelfDirectedScriptFullContent = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetSkipCurrentRecognize(v bool) *RealTimeDialogResponseBodyChoicesDelta {
	s.SkipCurrentRecognize = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) Validate() error {
	return dara.Validate(s)
}

type RealTimeDialogResponseBodyChoicesMessage struct {
	// example:
	//
	// null
	AnalysisProcess *string `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
	// time
	//
	// example:
	//
	// 1735139569523
	CallTime *string `json:"callTime,omitempty" xml:"callTime,omitempty"`
	// example:
	//
	// false
	HangUpDialog *bool `json:"hangUpDialog,omitempty" xml:"hangUpDialog,omitempty"`
	// example:
	//
	// 1853360771162058752
	IntentionCode   *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	IntentionName   *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
	Interrupt       *bool   `json:"interrupt,omitempty" xml:"interrupt,omitempty"`
	// example:
	//
	// null
	RecommendIntention *string `json:"recommendIntention,omitempty" xml:"recommendIntention,omitempty"`
	// example:
	//
	// null
	RecommendScript *string `json:"recommendScript,omitempty" xml:"recommendScript,omitempty"`
	// example:
	//
	// null
	SelfDirectedScript            *string `json:"selfDirectedScript,omitempty" xml:"selfDirectedScript,omitempty"`
	SelfDirectedScriptFullContent *string `json:"selfDirectedScriptFullContent,omitempty" xml:"selfDirectedScriptFullContent,omitempty"`
	SkipCurrentRecognize          *bool   `json:"skipCurrentRecognize,omitempty" xml:"skipCurrentRecognize,omitempty"`
}

func (s RealTimeDialogResponseBodyChoicesMessage) String() string {
	return dara.Prettify(s)
}

func (s RealTimeDialogResponseBodyChoicesMessage) GoString() string {
	return s.String()
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetAnalysisProcess() *string {
	return s.AnalysisProcess
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetCallTime() *string {
	return s.CallTime
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetHangUpDialog() *bool {
	return s.HangUpDialog
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetIntentionCode() *string {
	return s.IntentionCode
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetIntentionName() *string {
	return s.IntentionName
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetIntentionScript() *string {
	return s.IntentionScript
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetInterrupt() *bool {
	return s.Interrupt
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetRecommendIntention() *string {
	return s.RecommendIntention
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetRecommendScript() *string {
	return s.RecommendScript
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetSelfDirectedScript() *string {
	return s.SelfDirectedScript
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetSelfDirectedScriptFullContent() *string {
	return s.SelfDirectedScriptFullContent
}

func (s *RealTimeDialogResponseBodyChoicesMessage) GetSkipCurrentRecognize() *bool {
	return s.SkipCurrentRecognize
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetAnalysisProcess(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.AnalysisProcess = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetCallTime(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.CallTime = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetHangUpDialog(v bool) *RealTimeDialogResponseBodyChoicesMessage {
	s.HangUpDialog = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetIntentionCode(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.IntentionCode = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetIntentionName(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.IntentionName = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetIntentionScript(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.IntentionScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetInterrupt(v bool) *RealTimeDialogResponseBodyChoicesMessage {
	s.Interrupt = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetRecommendIntention(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.RecommendIntention = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetRecommendScript(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.RecommendScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetSelfDirectedScript(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.SelfDirectedScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetSelfDirectedScriptFullContent(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.SelfDirectedScriptFullContent = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetSkipCurrentRecognize(v bool) *RealTimeDialogResponseBodyChoicesMessage {
	s.SkipCurrentRecognize = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) Validate() error {
	return dara.Validate(s)
}
