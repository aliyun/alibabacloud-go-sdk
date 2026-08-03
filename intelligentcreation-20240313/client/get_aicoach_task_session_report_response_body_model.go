// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachTaskSessionReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int64) *GetAICoachTaskSessionReportResponseBody
	GetDuration() *int64
	SetEndTime(v string) *GetAICoachTaskSessionReportResponseBody
	GetEndTime() *string
	SetEvaluationRating(v string) *GetAICoachTaskSessionReportResponseBody
	GetEvaluationRating() *string
	SetEvaluationResult(v string) *GetAICoachTaskSessionReportResponseBody
	GetEvaluationResult() *string
	SetExtendCustomNameMap(v map[string]*string) *GetAICoachTaskSessionReportResponseBody
	GetExtendCustomNameMap() map[string]*string
	SetFeedback(v bool) *GetAICoachTaskSessionReportResponseBody
	GetFeedback() *bool
	SetRequestId(v string) *GetAICoachTaskSessionReportResponseBody
	GetRequestId() *string
	SetScriptName(v string) *GetAICoachTaskSessionReportResponseBody
	GetScriptName() *string
	SetStartTime(v string) *GetAICoachTaskSessionReportResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetAICoachTaskSessionReportResponseBody
	GetStatus() *string
	SetUid(v string) *GetAICoachTaskSessionReportResponseBody
	GetUid() *string
}

type GetAICoachTaskSessionReportResponseBody struct {
	Duration            *int64             `json:"duration,omitempty" xml:"duration,omitempty"`
	EndTime             *string            `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EvaluationRating    *string            `json:"evaluationRating,omitempty" xml:"evaluationRating,omitempty"`
	EvaluationResult    *string            `json:"evaluationResult,omitempty" xml:"evaluationResult,omitempty"`
	ExtendCustomNameMap map[string]*string `json:"extendCustomNameMap,omitempty" xml:"extendCustomNameMap,omitempty"`
	Feedback            *bool              `json:"feedback,omitempty" xml:"feedback,omitempty"`
	RequestId           *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScriptName          *string            `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	StartTime           *string            `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status              *string            `json:"status,omitempty" xml:"status,omitempty"`
	Uid                 *string            `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetAICoachTaskSessionReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachTaskSessionReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionReportResponseBody) GetDuration() *int64 {
	return s.Duration
}

func (s *GetAICoachTaskSessionReportResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAICoachTaskSessionReportResponseBody) GetEvaluationRating() *string {
	return s.EvaluationRating
}

func (s *GetAICoachTaskSessionReportResponseBody) GetEvaluationResult() *string {
	return s.EvaluationResult
}

func (s *GetAICoachTaskSessionReportResponseBody) GetExtendCustomNameMap() map[string]*string {
	return s.ExtendCustomNameMap
}

func (s *GetAICoachTaskSessionReportResponseBody) GetFeedback() *bool {
	return s.Feedback
}

func (s *GetAICoachTaskSessionReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAICoachTaskSessionReportResponseBody) GetScriptName() *string {
	return s.ScriptName
}

func (s *GetAICoachTaskSessionReportResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAICoachTaskSessionReportResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAICoachTaskSessionReportResponseBody) GetUid() *string {
	return s.Uid
}

func (s *GetAICoachTaskSessionReportResponseBody) SetDuration(v int64) *GetAICoachTaskSessionReportResponseBody {
	s.Duration = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetEndTime(v string) *GetAICoachTaskSessionReportResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetEvaluationRating(v string) *GetAICoachTaskSessionReportResponseBody {
	s.EvaluationRating = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetEvaluationResult(v string) *GetAICoachTaskSessionReportResponseBody {
	s.EvaluationResult = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetExtendCustomNameMap(v map[string]*string) *GetAICoachTaskSessionReportResponseBody {
	s.ExtendCustomNameMap = v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetFeedback(v bool) *GetAICoachTaskSessionReportResponseBody {
	s.Feedback = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetRequestId(v string) *GetAICoachTaskSessionReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetScriptName(v string) *GetAICoachTaskSessionReportResponseBody {
	s.ScriptName = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetStartTime(v string) *GetAICoachTaskSessionReportResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetStatus(v string) *GetAICoachTaskSessionReportResponseBody {
	s.Status = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) SetUid(v string) *GetAICoachTaskSessionReportResponseBody {
	s.Uid = &v
	return s
}

func (s *GetAICoachTaskSessionReportResponseBody) Validate() error {
	return dara.Validate(s)
}
