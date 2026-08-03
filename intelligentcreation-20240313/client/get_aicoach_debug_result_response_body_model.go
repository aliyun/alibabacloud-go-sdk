// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachDebugResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunSubId(v string) *GetAICoachDebugResultResponseBody
	GetAliyunSubId() *string
	SetDataId(v string) *GetAICoachDebugResultResponseBody
	GetDataId() *string
	SetDataType(v int64) *GetAICoachDebugResultResponseBody
	GetDataType() *int64
	SetDialogueList(v []*GetAICoachDebugResultResponseBodyDialogueList) *GetAICoachDebugResultResponseBody
	GetDialogueList() []*GetAICoachDebugResultResponseBodyDialogueList
	SetErrorCode(v string) *GetAICoachDebugResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAICoachDebugResultResponseBody
	GetErrorMessage() *string
	SetFinishTime(v string) *GetAICoachDebugResultResponseBody
	GetFinishTime() *string
	SetGmtCreate(v string) *GetAICoachDebugResultResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetAICoachDebugResultResponseBody
	GetGmtModified() *string
	SetRequestId(v string) *GetAICoachDebugResultResponseBody
	GetRequestId() *string
	SetScriptDebugId(v string) *GetAICoachDebugResultResponseBody
	GetScriptDebugId() *string
	SetStatus(v int64) *GetAICoachDebugResultResponseBody
	GetStatus() *int64
	SetSuccess(v bool) *GetAICoachDebugResultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetAICoachDebugResultResponseBody
	GetTaskId() *string
	SetTaskReport(v *GetAICoachDebugResultResponseBodyTaskReport) *GetAICoachDebugResultResponseBody
	GetTaskReport() *GetAICoachDebugResultResponseBodyTaskReport
}

type GetAICoachDebugResultResponseBody struct {
	AliyunSubId   *string                                          `json:"aliyunSubId,omitempty" xml:"aliyunSubId,omitempty"`
	DataId        *string                                          `json:"dataId,omitempty" xml:"dataId,omitempty"`
	DataType      *int64                                           `json:"dataType,omitempty" xml:"dataType,omitempty"`
	DialogueList  []*GetAICoachDebugResultResponseBodyDialogueList `json:"dialogueList,omitempty" xml:"dialogueList,omitempty" type:"Repeated"`
	ErrorCode     *string                                          `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage  *string                                          `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	FinishTime    *string                                          `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	GmtCreate     *string                                          `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *string                                          `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	RequestId     *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScriptDebugId *string                                          `json:"scriptDebugId,omitempty" xml:"scriptDebugId,omitempty"`
	Status        *int64                                           `json:"status,omitempty" xml:"status,omitempty"`
	Success       *bool                                            `json:"success,omitempty" xml:"success,omitempty"`
	TaskId        *string                                          `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskReport    *GetAICoachDebugResultResponseBodyTaskReport     `json:"taskReport,omitempty" xml:"taskReport,omitempty" type:"Struct"`
}

func (s GetAICoachDebugResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachDebugResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICoachDebugResultResponseBody) GetAliyunSubId() *string {
	return s.AliyunSubId
}

func (s *GetAICoachDebugResultResponseBody) GetDataId() *string {
	return s.DataId
}

func (s *GetAICoachDebugResultResponseBody) GetDataType() *int64 {
	return s.DataType
}

func (s *GetAICoachDebugResultResponseBody) GetDialogueList() []*GetAICoachDebugResultResponseBodyDialogueList {
	return s.DialogueList
}

func (s *GetAICoachDebugResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAICoachDebugResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAICoachDebugResultResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetAICoachDebugResultResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetAICoachDebugResultResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetAICoachDebugResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAICoachDebugResultResponseBody) GetScriptDebugId() *string {
	return s.ScriptDebugId
}

func (s *GetAICoachDebugResultResponseBody) GetStatus() *int64 {
	return s.Status
}

func (s *GetAICoachDebugResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAICoachDebugResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAICoachDebugResultResponseBody) GetTaskReport() *GetAICoachDebugResultResponseBodyTaskReport {
	return s.TaskReport
}

func (s *GetAICoachDebugResultResponseBody) SetAliyunSubId(v string) *GetAICoachDebugResultResponseBody {
	s.AliyunSubId = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetDataId(v string) *GetAICoachDebugResultResponseBody {
	s.DataId = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetDataType(v int64) *GetAICoachDebugResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetDialogueList(v []*GetAICoachDebugResultResponseBodyDialogueList) *GetAICoachDebugResultResponseBody {
	s.DialogueList = v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetErrorCode(v string) *GetAICoachDebugResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetErrorMessage(v string) *GetAICoachDebugResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetFinishTime(v string) *GetAICoachDebugResultResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetGmtCreate(v string) *GetAICoachDebugResultResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetGmtModified(v string) *GetAICoachDebugResultResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetRequestId(v string) *GetAICoachDebugResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetScriptDebugId(v string) *GetAICoachDebugResultResponseBody {
	s.ScriptDebugId = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetStatus(v int64) *GetAICoachDebugResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetSuccess(v bool) *GetAICoachDebugResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetTaskId(v string) *GetAICoachDebugResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetAICoachDebugResultResponseBody) SetTaskReport(v *GetAICoachDebugResultResponseBodyTaskReport) *GetAICoachDebugResultResponseBody {
	s.TaskReport = v
	return s
}

func (s *GetAICoachDebugResultResponseBody) Validate() error {
	if s.DialogueList != nil {
		for _, item := range s.DialogueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskReport != nil {
		if err := s.TaskReport.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAICoachDebugResultResponseBodyDialogueList struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetAICoachDebugResultResponseBodyDialogueList) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachDebugResultResponseBodyDialogueList) GoString() string {
	return s.String()
}

func (s *GetAICoachDebugResultResponseBodyDialogueList) GetMessage() *string {
	return s.Message
}

func (s *GetAICoachDebugResultResponseBodyDialogueList) GetRole() *string {
	return s.Role
}

func (s *GetAICoachDebugResultResponseBodyDialogueList) SetMessage(v string) *GetAICoachDebugResultResponseBodyDialogueList {
	s.Message = &v
	return s
}

func (s *GetAICoachDebugResultResponseBodyDialogueList) SetRole(v string) *GetAICoachDebugResultResponseBodyDialogueList {
	s.Role = &v
	return s
}

func (s *GetAICoachDebugResultResponseBodyDialogueList) Validate() error {
	return dara.Validate(s)
}

type GetAICoachDebugResultResponseBodyTaskReport struct {
	DeductionRule  *GetAICoachDebugResultResponseBodyTaskReportDeductionRule  `json:"deductionRule,omitempty" xml:"deductionRule,omitempty" type:"Struct"`
	Expressiveness *GetAICoachDebugResultResponseBodyTaskReportExpressiveness `json:"expressiveness,omitempty" xml:"expressiveness,omitempty" type:"Struct"`
	Point          *GetAICoachDebugResultResponseBodyTaskReportPoint          `json:"point,omitempty" xml:"point,omitempty" type:"Struct"`
}

func (s GetAICoachDebugResultResponseBodyTaskReport) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachDebugResultResponseBodyTaskReport) GoString() string {
	return s.String()
}

func (s *GetAICoachDebugResultResponseBodyTaskReport) GetDeductionRule() *GetAICoachDebugResultResponseBodyTaskReportDeductionRule {
	return s.DeductionRule
}

func (s *GetAICoachDebugResultResponseBodyTaskReport) GetExpressiveness() *GetAICoachDebugResultResponseBodyTaskReportExpressiveness {
	return s.Expressiveness
}

func (s *GetAICoachDebugResultResponseBodyTaskReport) GetPoint() *GetAICoachDebugResultResponseBodyTaskReportPoint {
	return s.Point
}

func (s *GetAICoachDebugResultResponseBodyTaskReport) SetDeductionRule(v *GetAICoachDebugResultResponseBodyTaskReportDeductionRule) *GetAICoachDebugResultResponseBodyTaskReport {
	s.DeductionRule = v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReport) SetExpressiveness(v *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) *GetAICoachDebugResultResponseBodyTaskReport {
	s.Expressiveness = v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReport) SetPoint(v *GetAICoachDebugResultResponseBodyTaskReportPoint) *GetAICoachDebugResultResponseBodyTaskReport {
	s.Point = v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReport) Validate() error {
	if s.DeductionRule != nil {
		if err := s.DeductionRule.Validate(); err != nil {
			return err
		}
	}
	if s.Expressiveness != nil {
		if err := s.Expressiveness.Validate(); err != nil {
			return err
		}
	}
	if s.Point != nil {
		if err := s.Point.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAICoachDebugResultResponseBodyTaskReportDeductionRule struct {
	Hit    *bool     `json:"hit,omitempty" xml:"hit,omitempty"`
	Name   *string   `json:"name,omitempty" xml:"name,omitempty"`
	Reason []*string `json:"reason,omitempty" xml:"reason,omitempty" type:"Repeated"`
}

func (s GetAICoachDebugResultResponseBodyTaskReportDeductionRule) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachDebugResultResponseBodyTaskReportDeductionRule) GoString() string {
	return s.String()
}

func (s *GetAICoachDebugResultResponseBodyTaskReportDeductionRule) GetHit() *bool {
	return s.Hit
}

func (s *GetAICoachDebugResultResponseBodyTaskReportDeductionRule) GetName() *string {
	return s.Name
}

func (s *GetAICoachDebugResultResponseBodyTaskReportDeductionRule) GetReason() []*string {
	return s.Reason
}

func (s *GetAICoachDebugResultResponseBodyTaskReportDeductionRule) SetHit(v bool) *GetAICoachDebugResultResponseBodyTaskReportDeductionRule {
	s.Hit = &v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportDeductionRule) SetName(v string) *GetAICoachDebugResultResponseBodyTaskReportDeductionRule {
	s.Name = &v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportDeductionRule) SetReason(v []*string) *GetAICoachDebugResultResponseBodyTaskReportDeductionRule {
	s.Reason = v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportDeductionRule) Validate() error {
	return dara.Validate(s)
}

type GetAICoachDebugResultResponseBodyTaskReportExpressiveness struct {
	Name        *string   `json:"name,omitempty" xml:"name,omitempty"`
	Reason      []*string `json:"reason,omitempty" xml:"reason,omitempty" type:"Repeated"`
	ScoreRounds *int32    `json:"scoreRounds,omitempty" xml:"scoreRounds,omitempty"`
	Status      *string   `json:"status,omitempty" xml:"status,omitempty"`
	TotalRounds *int32    `json:"totalRounds,omitempty" xml:"totalRounds,omitempty"`
}

func (s GetAICoachDebugResultResponseBodyTaskReportExpressiveness) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachDebugResultResponseBodyTaskReportExpressiveness) GoString() string {
	return s.String()
}

func (s *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) GetName() *string {
	return s.Name
}

func (s *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) GetReason() []*string {
	return s.Reason
}

func (s *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) GetScoreRounds() *int32 {
	return s.ScoreRounds
}

func (s *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) GetStatus() *string {
	return s.Status
}

func (s *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) GetTotalRounds() *int32 {
	return s.TotalRounds
}

func (s *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) SetName(v string) *GetAICoachDebugResultResponseBodyTaskReportExpressiveness {
	s.Name = &v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) SetReason(v []*string) *GetAICoachDebugResultResponseBodyTaskReportExpressiveness {
	s.Reason = v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) SetScoreRounds(v int32) *GetAICoachDebugResultResponseBodyTaskReportExpressiveness {
	s.ScoreRounds = &v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) SetStatus(v string) *GetAICoachDebugResultResponseBodyTaskReportExpressiveness {
	s.Status = &v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) SetTotalRounds(v int32) *GetAICoachDebugResultResponseBodyTaskReportExpressiveness {
	s.TotalRounds = &v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportExpressiveness) Validate() error {
	return dara.Validate(s)
}

type GetAICoachDebugResultResponseBodyTaskReportPoint struct {
	AnswerList []*GetAICoachDebugResultResponseBodyTaskReportPointAnswerList `json:"answerList,omitempty" xml:"answerList,omitempty" type:"Repeated"`
	Name       *string                                                       `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAICoachDebugResultResponseBodyTaskReportPoint) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachDebugResultResponseBodyTaskReportPoint) GoString() string {
	return s.String()
}

func (s *GetAICoachDebugResultResponseBodyTaskReportPoint) GetAnswerList() []*GetAICoachDebugResultResponseBodyTaskReportPointAnswerList {
	return s.AnswerList
}

func (s *GetAICoachDebugResultResponseBodyTaskReportPoint) GetName() *string {
	return s.Name
}

func (s *GetAICoachDebugResultResponseBodyTaskReportPoint) SetAnswerList(v []*GetAICoachDebugResultResponseBodyTaskReportPointAnswerList) *GetAICoachDebugResultResponseBodyTaskReportPoint {
	s.AnswerList = v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportPoint) SetName(v string) *GetAICoachDebugResultResponseBodyTaskReportPoint {
	s.Name = &v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportPoint) Validate() error {
	if s.AnswerList != nil {
		for _, item := range s.AnswerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAICoachDebugResultResponseBodyTaskReportPointAnswerList struct {
	Reason []*string `json:"reason,omitempty" xml:"reason,omitempty" type:"Repeated"`
	Status *int64    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAICoachDebugResultResponseBodyTaskReportPointAnswerList) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachDebugResultResponseBodyTaskReportPointAnswerList) GoString() string {
	return s.String()
}

func (s *GetAICoachDebugResultResponseBodyTaskReportPointAnswerList) GetReason() []*string {
	return s.Reason
}

func (s *GetAICoachDebugResultResponseBodyTaskReportPointAnswerList) GetStatus() *int64 {
	return s.Status
}

func (s *GetAICoachDebugResultResponseBodyTaskReportPointAnswerList) SetReason(v []*string) *GetAICoachDebugResultResponseBodyTaskReportPointAnswerList {
	s.Reason = v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportPointAnswerList) SetStatus(v int64) *GetAICoachDebugResultResponseBodyTaskReportPointAnswerList {
	s.Status = &v
	return s
}

func (s *GetAICoachDebugResultResponseBodyTaskReportPointAnswerList) Validate() error {
	return dara.Validate(s)
}
