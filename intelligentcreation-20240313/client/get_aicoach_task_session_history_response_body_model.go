// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachTaskSessionHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConversationList(v []*GetAICoachTaskSessionHistoryResponseBodyConversationList) *GetAICoachTaskSessionHistoryResponseBody
	GetConversationList() []*GetAICoachTaskSessionHistoryResponseBodyConversationList
	SetDuration(v int64) *GetAICoachTaskSessionHistoryResponseBody
	GetDuration() *int64
	SetEndTime(v string) *GetAICoachTaskSessionHistoryResponseBody
	GetEndTime() *string
	SetPauseDuration(v int64) *GetAICoachTaskSessionHistoryResponseBody
	GetPauseDuration() *int64
	SetRequestId(v string) *GetAICoachTaskSessionHistoryResponseBody
	GetRequestId() *string
	SetScriptName(v string) *GetAICoachTaskSessionHistoryResponseBody
	GetScriptName() *string
	SetStartTime(v string) *GetAICoachTaskSessionHistoryResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetAICoachTaskSessionHistoryResponseBody
	GetStatus() *string
	SetTotal(v int32) *GetAICoachTaskSessionHistoryResponseBody
	GetTotal() *int32
	SetUid(v string) *GetAICoachTaskSessionHistoryResponseBody
	GetUid() *string
}

type GetAICoachTaskSessionHistoryResponseBody struct {
	ConversationList []*GetAICoachTaskSessionHistoryResponseBodyConversationList `json:"conversationList,omitempty" xml:"conversationList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 2024-11-08 09:33:21
	EndTime       *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	PauseDuration *int64  `json:"pauseDuration,omitempty" xml:"pauseDuration,omitempty"`
	// example:
	//
	// D5798660-1531-5D12-9C20-16FEE9D22351
	RequestId  *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// example:
	//
	// 2024-08-21 05:00:01
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	Total     *int32  `json:"total,omitempty" xml:"total,omitempty"`
	// example:
	//
	// 1579404690269235
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetAICoachTaskSessionHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachTaskSessionHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionHistoryResponseBody) GetConversationList() []*GetAICoachTaskSessionHistoryResponseBodyConversationList {
	return s.ConversationList
}

func (s *GetAICoachTaskSessionHistoryResponseBody) GetDuration() *int64 {
	return s.Duration
}

func (s *GetAICoachTaskSessionHistoryResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAICoachTaskSessionHistoryResponseBody) GetPauseDuration() *int64 {
	return s.PauseDuration
}

func (s *GetAICoachTaskSessionHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAICoachTaskSessionHistoryResponseBody) GetScriptName() *string {
	return s.ScriptName
}

func (s *GetAICoachTaskSessionHistoryResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAICoachTaskSessionHistoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAICoachTaskSessionHistoryResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetAICoachTaskSessionHistoryResponseBody) GetUid() *string {
	return s.Uid
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetConversationList(v []*GetAICoachTaskSessionHistoryResponseBodyConversationList) *GetAICoachTaskSessionHistoryResponseBody {
	s.ConversationList = v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetDuration(v int64) *GetAICoachTaskSessionHistoryResponseBody {
	s.Duration = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetEndTime(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetPauseDuration(v int64) *GetAICoachTaskSessionHistoryResponseBody {
	s.PauseDuration = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetRequestId(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetScriptName(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.ScriptName = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetStartTime(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetStatus(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.Status = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetTotal(v int32) *GetAICoachTaskSessionHistoryResponseBody {
	s.Total = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) SetUid(v string) *GetAICoachTaskSessionHistoryResponseBody {
	s.Uid = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBody) Validate() error {
	if s.ConversationList != nil {
		for _, item := range s.ConversationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAICoachTaskSessionHistoryResponseBodyConversationList struct {
	AudioUrl           *string `json:"audioUrl,omitempty" xml:"audioUrl,omitempty"`
	DateLabel          *string `json:"dateLabel,omitempty" xml:"dateLabel,omitempty"`
	EvaluationFeedback *string `json:"evaluationFeedback,omitempty" xml:"evaluationFeedback,omitempty"`
	// example:
	//
	// {}
	EvaluationResult *string `json:"evaluationResult,omitempty" xml:"evaluationResult,omitempty"`
	Message          *string `json:"message,omitempty" xml:"message,omitempty"`
	RecordId         *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	Role             *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetAICoachTaskSessionHistoryResponseBodyConversationList) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachTaskSessionHistoryResponseBodyConversationList) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) GetAudioUrl() *string {
	return s.AudioUrl
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) GetDateLabel() *string {
	return s.DateLabel
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) GetEvaluationFeedback() *string {
	return s.EvaluationFeedback
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) GetEvaluationResult() *string {
	return s.EvaluationResult
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) GetMessage() *string {
	return s.Message
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) GetRecordId() *string {
	return s.RecordId
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) GetRole() *string {
	return s.Role
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetAudioUrl(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.AudioUrl = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetDateLabel(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.DateLabel = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetEvaluationFeedback(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.EvaluationFeedback = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetEvaluationResult(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.EvaluationResult = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetMessage(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.Message = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetRecordId(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.RecordId = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) SetRole(v string) *GetAICoachTaskSessionHistoryResponseBodyConversationList {
	s.Role = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponseBodyConversationList) Validate() error {
	return dara.Validate(s)
}
