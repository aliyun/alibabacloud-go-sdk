// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEduHomeworkCorrectSubmitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateEduHomeworkCorrectSubmitResponseBody
	GetCode() *string
	SetData(v *CreateEduHomeworkCorrectSubmitResponseBodyData) *CreateEduHomeworkCorrectSubmitResponseBody
	GetData() *CreateEduHomeworkCorrectSubmitResponseBodyData
	SetMessage(v string) *CreateEduHomeworkCorrectSubmitResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEduHomeworkCorrectSubmitResponseBody
	GetRequestId() *string
}

type CreateEduHomeworkCorrectSubmitResponseBody struct {
	// example:
	//
	// OK
	Code *string                                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateEduHomeworkCorrectSubmitResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateEduHomeworkCorrectSubmitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEduHomeworkCorrectSubmitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEduHomeworkCorrectSubmitResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateEduHomeworkCorrectSubmitResponseBody) GetData() *CreateEduHomeworkCorrectSubmitResponseBodyData {
	return s.Data
}

func (s *CreateEduHomeworkCorrectSubmitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEduHomeworkCorrectSubmitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEduHomeworkCorrectSubmitResponseBody) SetCode(v string) *CreateEduHomeworkCorrectSubmitResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBody) SetData(v *CreateEduHomeworkCorrectSubmitResponseBodyData) *CreateEduHomeworkCorrectSubmitResponseBody {
	s.Data = v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBody) SetMessage(v string) *CreateEduHomeworkCorrectSubmitResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBody) SetRequestId(v string) *CreateEduHomeworkCorrectSubmitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEduHomeworkCorrectSubmitResponseBodyData struct {
	Angle           *int64                                                           `json:"angle,omitempty" xml:"angle,omitempty"`
	DewarpUrl       *string                                                          `json:"dewarpUrl,omitempty" xml:"dewarpUrl,omitempty"`
	QuestionResults []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults `json:"questionResults,omitempty" xml:"questionResults,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyData) GetAngle() *int64 {
	return s.Angle
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyData) GetDewarpUrl() *string {
	return s.DewarpUrl
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyData) GetQuestionResults() []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults {
	return s.QuestionResults
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyData) SetAngle(v int64) *CreateEduHomeworkCorrectSubmitResponseBodyData {
	s.Angle = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyData) SetDewarpUrl(v string) *CreateEduHomeworkCorrectSubmitResponseBodyData {
	s.DewarpUrl = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyData) SetQuestionResults(v []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) *CreateEduHomeworkCorrectSubmitResponseBodyData {
	s.QuestionResults = v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyData) SetStatus(v string) *CreateEduHomeworkCorrectSubmitResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyData) SetTaskId(v string) *CreateEduHomeworkCorrectSubmitResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyData) Validate() error {
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

type CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults struct {
	AnswerResults []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults `json:"answerResults,omitempty" xml:"answerResults,omitempty" type:"Repeated"`
	// example:
	//
	// True
	Finish          *bool                                                                            `json:"finish,omitempty" xml:"finish,omitempty"`
	QuestionContent *string                                                                          `json:"questionContent,omitempty" xml:"questionContent,omitempty"`
	QuestionId      *string                                                                          `json:"questionId,omitempty" xml:"questionId,omitempty"`
	QuestionPoints  [][]*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints `json:"questionPoints,omitempty" xml:"questionPoints,omitempty" type:"Repeated"`
	QuestionType    *string                                                                          `json:"questionType,omitempty" xml:"questionType,omitempty"`
	// example:
	//
	// 已处理：接口耗时已恢复低于阈值（09:53后分钟平均约1.8s～3.3s）
	Solution *string `json:"solution,omitempty" xml:"solution,omitempty"`
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) String() string {
	return dara.Prettify(s)
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) GoString() string {
	return s.String()
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) GetAnswerResults() []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults {
	return s.AnswerResults
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) GetFinish() *bool {
	return s.Finish
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) GetQuestionContent() *string {
	return s.QuestionContent
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) GetQuestionId() *string {
	return s.QuestionId
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) GetQuestionPoints() [][]*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints {
	return s.QuestionPoints
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) GetQuestionType() *string {
	return s.QuestionType
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) GetSolution() *string {
	return s.Solution
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) SetAnswerResults(v []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults {
	s.AnswerResults = v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) SetFinish(v bool) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults {
	s.Finish = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) SetQuestionContent(v string) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults {
	s.QuestionContent = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) SetQuestionId(v string) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults {
	s.QuestionId = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) SetQuestionPoints(v [][]*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults {
	s.QuestionPoints = v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) SetQuestionType(v string) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults {
	s.QuestionType = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) SetSolution(v string) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults {
	s.Solution = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults) Validate() error {
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

type CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults struct {
	AnswerPoints   []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints `json:"answerPoints,omitempty" xml:"answerPoints,omitempty" type:"Repeated"`
	CheckPoint     *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint     `json:"checkPoint,omitempty" xml:"checkPoint,omitempty" type:"Struct"`
	CheckResult    *string                                                                                   `json:"checkResult,omitempty" xml:"checkResult,omitempty"`
	StandardAnswer *string                                                                                   `json:"standardAnswer,omitempty" xml:"standardAnswer,omitempty"`
	UserAnswer     *string                                                                                   `json:"userAnswer,omitempty" xml:"userAnswer,omitempty"`
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) String() string {
	return dara.Prettify(s)
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) GoString() string {
	return s.String()
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) GetAnswerPoints() []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints {
	return s.AnswerPoints
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) GetCheckPoint() *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint {
	return s.CheckPoint
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) GetCheckResult() *string {
	return s.CheckResult
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) GetStandardAnswer() *string {
	return s.StandardAnswer
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) GetUserAnswer() *string {
	return s.UserAnswer
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) SetAnswerPoints(v []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults {
	s.AnswerPoints = v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) SetCheckPoint(v *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults {
	s.CheckPoint = v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) SetCheckResult(v string) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults {
	s.CheckResult = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) SetStandardAnswer(v string) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults {
	s.StandardAnswer = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) SetUserAnswer(v string) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults {
	s.UserAnswer = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults) Validate() error {
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

type CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints struct {
	// example:
	//
	// 0.0
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 0.03
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) String() string {
	return dara.Prettify(s)
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) GoString() string {
	return s.String()
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) GetX() *int64 {
	return s.X
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) GetY() *int64 {
	return s.Y
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) SetX(v int64) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints {
	s.X = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) SetY(v int64) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints {
	s.Y = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints) Validate() error {
	return dara.Validate(s)
}

type CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint struct {
	// example:
	//
	// 0.0
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 0.90
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint) String() string {
	return dara.Prettify(s)
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint) GoString() string {
	return s.String()
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint) GetX() *int64 {
	return s.X
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint) GetY() *int64 {
	return s.Y
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint) SetX(v int64) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint {
	s.X = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint) SetY(v int64) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint {
	s.Y = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint) Validate() error {
	return dara.Validate(s)
}

type CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints struct {
	// example:
	//
	// 0.0
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// example:
	//
	// 0.0
	Y *int64 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints) String() string {
	return dara.Prettify(s)
}

func (s CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints) GoString() string {
	return s.String()
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints) GetX() *int64 {
	return s.X
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints) GetY() *int64 {
	return s.Y
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints) SetX(v int64) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints {
	s.X = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints) SetY(v int64) *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints {
	s.Y = &v
	return s
}

func (s *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints) Validate() error {
	return dara.Validate(s)
}
