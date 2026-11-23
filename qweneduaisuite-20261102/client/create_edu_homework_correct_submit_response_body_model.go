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
	// The API status.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The import command.
	Data *CreateEduHomeworkCorrectSubmitResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID on the caller side.
	//
	// example:
	//
	// req-20261116-0001
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
	// The clockwise rotation angle of the image, in degrees.
	//
	// example:
	//
	// 0
	Angle *int64 `json:"angle,omitempty" xml:"angle,omitempty"`
	// The URL of the corrected image. If `dewarp` is set to `false`, the URL of the original image is returned.
	//
	// example:
	//
	// https://example.com/dewarp.jpg
	DewarpUrl *string `json:"dewarpUrl,omitempty" xml:"dewarpUrl,omitempty"`
	// The array of question recognition and grading results.
	QuestionResults []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResults `json:"questionResults,omitempty" xml:"questionResults,omitempty" type:"Repeated"`
	// The task status.
	//
	// example:
	//
	// Examples: submitted, processing, finished, or failed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The unique identifier of the task. You must include this ID in subsequent queries.
	//
	// example:
	//
	// task-20261116-abc123
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
	// The array of answering and grading results for the question.
	AnswerResults []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResults `json:"answerResults,omitempty" xml:"answerResults,omitempty" type:"Repeated"`
	// Specifies whether the task is completed:
	//
	// - **true**: The task is completed.
	//
	// - **false**: The task is not completed.
	//
	// example:
	//
	// true
	Finish *bool `json:"finish,omitempty" xml:"finish,omitempty"`
	// The text content of the question.
	//
	// example:
	//
	// 1 + 1 = ?
	QuestionContent *string `json:"questionContent,omitempty" xml:"questionContent,omitempty"`
	// The question ID.
	//
	// example:
	//
	// q-001
	QuestionId *string `json:"questionId,omitempty" xml:"questionId,omitempty"`
	// The coordinate group of the question area. It generally contains four vertex coordinates arranged in the order of top-left, top-right, bottom-right, and bottom-left.
	//
	// example:
	//
	// Formatted as `[[{"x": 0, "y": 0}, ...]]`
	QuestionPoints [][]*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsQuestionPoints `json:"questionPoints,omitempty" xml:"questionPoints,omitempty" type:"Repeated"`
	// The question type.
	//
	// example:
	//
	// choice
	QuestionType *string `json:"questionType,omitempty" xml:"questionType,omitempty"`
	// The question analysis and solution.
	//
	// example:
	//
	// The question analysis and solution
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
	// The coordinate group of the user\\"s answering area. It generally contains four vertex coordinates arranged in the order of top-left, top-right, bottom-right, and bottom-left, formatted as `[{"x": ..., "y": ...}, ...]`.
	AnswerPoints []*CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsAnswerPoints `json:"answerPoints,omitempty" xml:"answerPoints,omitempty" type:"Repeated"`
	// The coordinates of the grading mark point, formatted as `{ "x": ..., "y": ... }`.
	CheckPoint *CreateEduHomeworkCorrectSubmitResponseBodyDataQuestionResultsAnswerResultsCheckPoint `json:"checkPoint,omitempty" xml:"checkPoint,omitempty" type:"Struct"`
	// The grading conclusion, such as `correct`, `wrong`, or `partial`.
	//
	// example:
	//
	// correct
	CheckResult *string `json:"checkResult,omitempty" xml:"checkResult,omitempty"`
	// The ground truth.
	//
	// example:
	//
	// 2
	StandardAnswer *string `json:"standardAnswer,omitempty" xml:"standardAnswer,omitempty"`
	// The student\\"s answer.
	//
	// example:
	//
	// 2
	UserAnswer *string `json:"userAnswer,omitempty" xml:"userAnswer,omitempty"`
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
	// The X-axis coordinate value.
	//
	// example:
	//
	// 120
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// The Y-axis coordinate value.
	//
	// example:
	//
	// 980
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
	// The X-axis coordinate.
	//
	// example:
	//
	// 120
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// The Y-axis coordinate.
	//
	// example:
	//
	// 980
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
	// The X-axis coordinate value.
	//
	// example:
	//
	// 120
	X *int64 `json:"x,omitempty" xml:"x,omitempty"`
	// The Y-axis coordinate value.
	//
	// example:
	//
	// 980
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
