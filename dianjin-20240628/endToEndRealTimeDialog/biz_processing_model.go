// This file is auto-generated, don't edit it. Thanks.
package endToEndRealTimeDialog

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iBizProcessing interface {
  dara.Model
  String() string
  GoString() string
  SetId(v string) *BizProcessing
  GetId() *string 
  SetChoices(v []*BizProcessingChoices) *BizProcessing
  GetChoices() []*BizProcessingChoices 
  SetCreated(v string) *BizProcessing
  GetCreated() *string 
  SetSuccess(v bool) *BizProcessing
  GetSuccess() *bool 
  SetRequestId(v string) *BizProcessing
  GetRequestId() *string 
}

type BizProcessing struct {
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Choices []*BizProcessingChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
  Created *string `json:"created,omitempty" xml:"created,omitempty"`
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BizProcessing) String() string {
  return dara.Prettify(s)
}

func (s BizProcessing) GoString() string {
  return s.String()
}

func (s *BizProcessing) GetId() *string  {
  return s.Id
}

func (s *BizProcessing) GetChoices() []*BizProcessingChoices  {
  return s.Choices
}

func (s *BizProcessing) GetCreated() *string  {
  return s.Created
}

func (s *BizProcessing) GetSuccess() *bool  {
  return s.Success
}

func (s *BizProcessing) GetRequestId() *string  {
  return s.RequestId
}

func (s *BizProcessing) SetId(v string) *BizProcessing {
  s.Id = &v
  return s
}

func (s *BizProcessing) SetChoices(v []*BizProcessingChoices) *BizProcessing {
  s.Choices = v
  return s
}

func (s *BizProcessing) SetCreated(v string) *BizProcessing {
  s.Created = &v
  return s
}

func (s *BizProcessing) SetSuccess(v bool) *BizProcessing {
  s.Success = &v
  return s
}

func (s *BizProcessing) SetRequestId(v string) *BizProcessing {
  s.RequestId = &v
  return s
}

func (s *BizProcessing) Validate() error {
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

type BizProcessingChoices struct {
  FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
  Delta *BizProcessingChoicesDelta `json:"delta,omitempty" xml:"delta,omitempty" type:"Struct"`
  Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
  Message *BizProcessingChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s BizProcessingChoices) String() string {
  return dara.Prettify(s)
}

func (s BizProcessingChoices) GoString() string {
  return s.String()
}

func (s *BizProcessingChoices) GetFinishReason() *string  {
  return s.FinishReason
}

func (s *BizProcessingChoices) GetDelta() *BizProcessingChoicesDelta  {
  return s.Delta
}

func (s *BizProcessingChoices) GetIndex() *int32  {
  return s.Index
}

func (s *BizProcessingChoices) GetMessage() *BizProcessingChoicesMessage  {
  return s.Message
}

func (s *BizProcessingChoices) SetFinishReason(v string) *BizProcessingChoices {
  s.FinishReason = &v
  return s
}

func (s *BizProcessingChoices) SetDelta(v *BizProcessingChoicesDelta) *BizProcessingChoices {
  s.Delta = v
  return s
}

func (s *BizProcessingChoices) SetIndex(v int32) *BizProcessingChoices {
  s.Index = &v
  return s
}

func (s *BizProcessingChoices) SetMessage(v *BizProcessingChoicesMessage) *BizProcessingChoices {
  s.Message = v
  return s
}

func (s *BizProcessingChoices) Validate() error {
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

type BizProcessingChoicesDelta struct {
  RecommendIntention *string `json:"recommendIntention,omitempty" xml:"recommendIntention,omitempty"`
  SelfDirectedScriptFullContent *string `json:"selfDirectedScriptFullContent,omitempty" xml:"selfDirectedScriptFullContent,omitempty"`
  HangUpDialog *bool `json:"hangUpDialog,omitempty" xml:"hangUpDialog,omitempty"`
  SelfDirectedScript *string `json:"selfDirectedScript,omitempty" xml:"selfDirectedScript,omitempty"`
  AnalysisProcess *string `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
  Interrupt *bool `json:"interrupt,omitempty" xml:"interrupt,omitempty"`
  IntentionCode *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
  CallTime *string `json:"callTime,omitempty" xml:"callTime,omitempty"`
  IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
  IntentionName *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
  RecommendScript *string `json:"recommendScript,omitempty" xml:"recommendScript,omitempty"`
}

func (s BizProcessingChoicesDelta) String() string {
  return dara.Prettify(s)
}

func (s BizProcessingChoicesDelta) GoString() string {
  return s.String()
}

func (s *BizProcessingChoicesDelta) GetRecommendIntention() *string  {
  return s.RecommendIntention
}

func (s *BizProcessingChoicesDelta) GetSelfDirectedScriptFullContent() *string  {
  return s.SelfDirectedScriptFullContent
}

func (s *BizProcessingChoicesDelta) GetHangUpDialog() *bool  {
  return s.HangUpDialog
}

func (s *BizProcessingChoicesDelta) GetSelfDirectedScript() *string  {
  return s.SelfDirectedScript
}

func (s *BizProcessingChoicesDelta) GetAnalysisProcess() *string  {
  return s.AnalysisProcess
}

func (s *BizProcessingChoicesDelta) GetInterrupt() *bool  {
  return s.Interrupt
}

func (s *BizProcessingChoicesDelta) GetIntentionCode() *string  {
  return s.IntentionCode
}

func (s *BizProcessingChoicesDelta) GetCallTime() *string  {
  return s.CallTime
}

func (s *BizProcessingChoicesDelta) GetIntentionScript() *string  {
  return s.IntentionScript
}

func (s *BizProcessingChoicesDelta) GetIntentionName() *string  {
  return s.IntentionName
}

func (s *BizProcessingChoicesDelta) GetRecommendScript() *string  {
  return s.RecommendScript
}

func (s *BizProcessingChoicesDelta) SetRecommendIntention(v string) *BizProcessingChoicesDelta {
  s.RecommendIntention = &v
  return s
}

func (s *BizProcessingChoicesDelta) SetSelfDirectedScriptFullContent(v string) *BizProcessingChoicesDelta {
  s.SelfDirectedScriptFullContent = &v
  return s
}

func (s *BizProcessingChoicesDelta) SetHangUpDialog(v bool) *BizProcessingChoicesDelta {
  s.HangUpDialog = &v
  return s
}

func (s *BizProcessingChoicesDelta) SetSelfDirectedScript(v string) *BizProcessingChoicesDelta {
  s.SelfDirectedScript = &v
  return s
}

func (s *BizProcessingChoicesDelta) SetAnalysisProcess(v string) *BizProcessingChoicesDelta {
  s.AnalysisProcess = &v
  return s
}

func (s *BizProcessingChoicesDelta) SetInterrupt(v bool) *BizProcessingChoicesDelta {
  s.Interrupt = &v
  return s
}

func (s *BizProcessingChoicesDelta) SetIntentionCode(v string) *BizProcessingChoicesDelta {
  s.IntentionCode = &v
  return s
}

func (s *BizProcessingChoicesDelta) SetCallTime(v string) *BizProcessingChoicesDelta {
  s.CallTime = &v
  return s
}

func (s *BizProcessingChoicesDelta) SetIntentionScript(v string) *BizProcessingChoicesDelta {
  s.IntentionScript = &v
  return s
}

func (s *BizProcessingChoicesDelta) SetIntentionName(v string) *BizProcessingChoicesDelta {
  s.IntentionName = &v
  return s
}

func (s *BizProcessingChoicesDelta) SetRecommendScript(v string) *BizProcessingChoicesDelta {
  s.RecommendScript = &v
  return s
}

func (s *BizProcessingChoicesDelta) Validate() error {
  return dara.Validate(s)
}

type BizProcessingChoicesMessage struct {
  RecommendIntention *string `json:"recommendIntention,omitempty" xml:"recommendIntention,omitempty"`
  SelfDirectedScriptFullContent *string `json:"selfDirectedScriptFullContent,omitempty" xml:"selfDirectedScriptFullContent,omitempty"`
  HangUpDialog *bool `json:"hangUpDialog,omitempty" xml:"hangUpDialog,omitempty"`
  SelfDirectedScript *string `json:"selfDirectedScript,omitempty" xml:"selfDirectedScript,omitempty"`
  AnalysisProcess *string `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
  Interrupt *bool `json:"interrupt,omitempty" xml:"interrupt,omitempty"`
  IntentionCode *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
  CallTime *string `json:"callTime,omitempty" xml:"callTime,omitempty"`
  IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
  IntentionName *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
  RecommendScript *string `json:"recommendScript,omitempty" xml:"recommendScript,omitempty"`
}

func (s BizProcessingChoicesMessage) String() string {
  return dara.Prettify(s)
}

func (s BizProcessingChoicesMessage) GoString() string {
  return s.String()
}

func (s *BizProcessingChoicesMessage) GetRecommendIntention() *string  {
  return s.RecommendIntention
}

func (s *BizProcessingChoicesMessage) GetSelfDirectedScriptFullContent() *string  {
  return s.SelfDirectedScriptFullContent
}

func (s *BizProcessingChoicesMessage) GetHangUpDialog() *bool  {
  return s.HangUpDialog
}

func (s *BizProcessingChoicesMessage) GetSelfDirectedScript() *string  {
  return s.SelfDirectedScript
}

func (s *BizProcessingChoicesMessage) GetAnalysisProcess() *string  {
  return s.AnalysisProcess
}

func (s *BizProcessingChoicesMessage) GetInterrupt() *bool  {
  return s.Interrupt
}

func (s *BizProcessingChoicesMessage) GetIntentionCode() *string  {
  return s.IntentionCode
}

func (s *BizProcessingChoicesMessage) GetCallTime() *string  {
  return s.CallTime
}

func (s *BizProcessingChoicesMessage) GetIntentionScript() *string  {
  return s.IntentionScript
}

func (s *BizProcessingChoicesMessage) GetIntentionName() *string  {
  return s.IntentionName
}

func (s *BizProcessingChoicesMessage) GetRecommendScript() *string  {
  return s.RecommendScript
}

func (s *BizProcessingChoicesMessage) SetRecommendIntention(v string) *BizProcessingChoicesMessage {
  s.RecommendIntention = &v
  return s
}

func (s *BizProcessingChoicesMessage) SetSelfDirectedScriptFullContent(v string) *BizProcessingChoicesMessage {
  s.SelfDirectedScriptFullContent = &v
  return s
}

func (s *BizProcessingChoicesMessage) SetHangUpDialog(v bool) *BizProcessingChoicesMessage {
  s.HangUpDialog = &v
  return s
}

func (s *BizProcessingChoicesMessage) SetSelfDirectedScript(v string) *BizProcessingChoicesMessage {
  s.SelfDirectedScript = &v
  return s
}

func (s *BizProcessingChoicesMessage) SetAnalysisProcess(v string) *BizProcessingChoicesMessage {
  s.AnalysisProcess = &v
  return s
}

func (s *BizProcessingChoicesMessage) SetInterrupt(v bool) *BizProcessingChoicesMessage {
  s.Interrupt = &v
  return s
}

func (s *BizProcessingChoicesMessage) SetIntentionCode(v string) *BizProcessingChoicesMessage {
  s.IntentionCode = &v
  return s
}

func (s *BizProcessingChoicesMessage) SetCallTime(v string) *BizProcessingChoicesMessage {
  s.CallTime = &v
  return s
}

func (s *BizProcessingChoicesMessage) SetIntentionScript(v string) *BizProcessingChoicesMessage {
  s.IntentionScript = &v
  return s
}

func (s *BizProcessingChoicesMessage) SetIntentionName(v string) *BizProcessingChoicesMessage {
  s.IntentionName = &v
  return s
}

func (s *BizProcessingChoicesMessage) SetRecommendScript(v string) *BizProcessingChoicesMessage {
  s.RecommendScript = &v
  return s
}

func (s *BizProcessingChoicesMessage) Validate() error {
  return dara.Validate(s)
}

