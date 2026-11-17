// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEduHomeworkCorrectionResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEduHomeworkCorrectionResultResponseBody
	GetCode() *string
	SetData(v *GetEduHomeworkCorrectionResultResponseBodyData) *GetEduHomeworkCorrectionResultResponseBody
	GetData() *GetEduHomeworkCorrectionResultResponseBodyData
	SetMessage(v string) *GetEduHomeworkCorrectionResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEduHomeworkCorrectionResultResponseBody
	GetRequestId() *string
}

type GetEduHomeworkCorrectionResultResponseBody struct {
	// example:
	//
	// OK
	Code *string                                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetEduHomeworkCorrectionResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEduHomeworkCorrectionResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEduHomeworkCorrectionResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetEduHomeworkCorrectionResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEduHomeworkCorrectionResultResponseBody) GetData() *GetEduHomeworkCorrectionResultResponseBodyData {
	return s.Data
}

func (s *GetEduHomeworkCorrectionResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEduHomeworkCorrectionResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEduHomeworkCorrectionResultResponseBody) SetCode(v string) *GetEduHomeworkCorrectionResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBody) SetData(v *GetEduHomeworkCorrectionResultResponseBodyData) *GetEduHomeworkCorrectionResultResponseBody {
	s.Data = v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBody) SetMessage(v string) *GetEduHomeworkCorrectionResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBody) SetRequestId(v string) *GetEduHomeworkCorrectionResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEduHomeworkCorrectionResultResponseBodyData struct {
	Angle           *int64                                                           `json:"angle,omitempty" xml:"angle,omitempty"`
	DewarpUrl       *string                                                          `json:"dewarpUrl,omitempty" xml:"dewarpUrl,omitempty"`
	QuestionResults []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults `json:"questionResults,omitempty" xml:"questionResults,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetEduHomeworkCorrectionResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEduHomeworkCorrectionResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEduHomeworkCorrectionResultResponseBodyData) GetAngle() *int64 {
	return s.Angle
}

func (s *GetEduHomeworkCorrectionResultResponseBodyData) GetDewarpUrl() *string {
	return s.DewarpUrl
}

func (s *GetEduHomeworkCorrectionResultResponseBodyData) GetQuestionResults() []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults {
	return s.QuestionResults
}

func (s *GetEduHomeworkCorrectionResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetEduHomeworkCorrectionResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetEduHomeworkCorrectionResultResponseBodyData) SetAngle(v int64) *GetEduHomeworkCorrectionResultResponseBodyData {
	s.Angle = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyData) SetDewarpUrl(v string) *GetEduHomeworkCorrectionResultResponseBodyData {
	s.DewarpUrl = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyData) SetQuestionResults(v []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) *GetEduHomeworkCorrectionResultResponseBodyData {
	s.QuestionResults = v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyData) SetStatus(v string) *GetEduHomeworkCorrectionResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyData) SetTaskId(v string) *GetEduHomeworkCorrectionResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyData) Validate() error {
	if s.QuestionResults != nil {
		for _, item := range s.QuestionResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults struct {
	AnswerResults []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults `json:"answerResults,omitempty" xml:"answerResults,omitempty" type:"Repeated"`
	// example:
	//
	// True
	Finish          *bool                                                                            `json:"finish,omitempty" xml:"finish,omitempty"`
	QuestionContent *string                                                                          `json:"questionContent,omitempty" xml:"questionContent,omitempty"`
	QuestionId      *string                                                                          `json:"questionId,omitempty" xml:"questionId,omitempty"`
	QuestionPoints  [][]*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints `json:"questionPoints,omitempty" xml:"questionPoints,omitempty" type:"Repeated"`
	QuestionType    *string                                                                          `json:"questionType,omitempty" xml:"questionType,omitempty"`
	// example:
	//
	// 已处理：接口耗时已恢复低于阈值（09:53后分钟平均约1.8s～3.3s）
	Solution *string `json:"solution,omitempty" xml:"solution,omitempty"`
}

func (s GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) String() string {
	return dara.Prettify(s)
}

func (s GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) GoString() string {
	return s.String()
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) GetAnswerResults() []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults {
	return s.AnswerResults
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) GetFinish() *bool {
	return s.Finish
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) GetQuestionContent() *string {
	return s.QuestionContent
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) GetQuestionId() *string {
	return s.QuestionId
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) GetQuestionPoints() [][]*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints {
	return s.QuestionPoints
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) GetQuestionType() *string {
	return s.QuestionType
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) GetSolution() *string {
	return s.Solution
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) SetAnswerResults(v []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults {
	s.AnswerResults = v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) SetFinish(v bool) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults {
	s.Finish = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) SetQuestionContent(v string) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults {
	s.QuestionContent = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) SetQuestionId(v string) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults {
	s.QuestionId = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) SetQuestionPoints(v [][]*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults {
	s.QuestionPoints = v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) SetQuestionType(v string) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults {
	s.QuestionType = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) SetSolution(v string) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults {
	s.Solution = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults) Validate() error {
	if s.AnswerResults != nil {
		for _, item := range s.AnswerResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults struct {
	AnswerPoints   []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints `json:"answerPoints,omitempty" xml:"answerPoints,omitempty" type:"Repeated"`
	CheckPoint     *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint     `json:"checkPoint,omitempty" xml:"checkPoint,omitempty" type:"Struct"`
	CheckResult    *string                                                                                   `json:"checkResult,omitempty" xml:"checkResult,omitempty"`
	StandardAnswer *string                                                                                   `json:"standardAnswer,omitempty" xml:"standardAnswer,omitempty"`
	UserAnswer     *string                                                                                   `json:"userAnswer,omitempty" xml:"userAnswer,omitempty"`
}

func (s GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) String() string {
	return dara.Prettify(s)
}

func (s GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) GoString() string {
	return s.String()
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) GetAnswerPoints() []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints {
	return s.AnswerPoints
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) GetCheckPoint() *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint {
	return s.CheckPoint
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) GetCheckResult() *string {
	return s.CheckResult
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) GetStandardAnswer() *string {
	return s.StandardAnswer
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) GetUserAnswer() *string {
	return s.UserAnswer
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) SetAnswerPoints(v []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults {
	s.AnswerPoints = v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) SetCheckPoint(v *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults {
	s.CheckPoint = v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) SetCheckResult(v string) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults {
	s.CheckResult = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) SetStandardAnswer(v string) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults {
	s.StandardAnswer = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) SetUserAnswer(v string) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults {
	s.UserAnswer = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults) Validate() error {
	if s.AnswerPoints != nil {
		for _, item := range s.AnswerPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CheckPoint != nil {
		if err := s.CheckPoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints struct {
	// example:
	//
	// 0.0
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 0.03
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) String() string {
	return dara.Prettify(s)
}

func (s GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) GoString() string {
	return s.String()
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) GetX() *int64 {
	return s.X
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) GetY() *int64 {
	return s.Y
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) SetX(v int64) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints {
	s.X = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) SetY(v int64) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints {
	s.Y = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) Validate() error {
	return dara.Validate(s)
}

type GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint struct {
	// example:
	//
	// 0.0
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 0.90
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint) String() string {
	return dara.Prettify(s)
}

func (s GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint) GoString() string {
	return s.String()
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint) GetX() *int64 {
	return s.X
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint) GetY() *int64 {
	return s.Y
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint) SetX(v int64) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint {
	s.X = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint) SetY(v int64) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint {
	s.Y = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint) Validate() error {
	return dara.Validate(s)
}

type GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints struct {
	// example:
	//
	// 0.0
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 0.0
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints) String() string {
	return dara.Prettify(s)
}

func (s GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints) GoString() string {
	return s.String()
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints) GetX() *int64 {
	return s.X
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints) GetY() *int64 {
	return s.Y
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints) SetX(v int64) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints {
	s.X = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints) SetY(v int64) *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints {
	s.Y = &v
	return s
}

func (s *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints) Validate() error {
	return dara.Validate(s)
}
