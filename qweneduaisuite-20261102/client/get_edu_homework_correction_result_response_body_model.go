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
	// The API status.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The business data. This field is usually empty when the task fails.
	//
	// example:
	//
	// {
	//
	//   "request_id": "req-20261116-0001",
	//
	//   "code": "SUCCESS",
	//
	//   "message": "success",
	//
	//   "data": {
	//
	//     "task_id": "task-20261116-abc123",
	//
	//     "status": "finished",
	//
	//     "course": "math",
	//
	//     "dewarp_url": "https://example.com/dewarp.jpg",
	//
	//     "angle": 0,
	//
	//     "render_url": "https://example.com/render.jpg",
	//
	//     "question_results": [
	//
	//       {
	//
	//         "question_id": "q-001",
	//
	//         "question_type": "choice",
	//
	//         "question_content": "1 + 1 = ?",
	//
	//         "question_points": [
	//
	//           [
	//
	//             {"x": 100, "y": 100},
	//
	//             {"x": 200, "y": 100},
	//
	//             {"x": 200, "y": 150},
	//
	//             {"x": 100, "y": 150}
	//
	//           ]
	//
	//         ],
	//
	//         "finish": true,
	//
	//         "answer_results": [
	//
	//           {
	//
	//             "check_result": "correct",
	//
	//             "check_point": {"x": 150, "y": 125},
	//
	//             "answer_points": [
	//
	//               {"x": 120, "y": 120},
	//
	//               {"x": 180, "y": 120},
	//
	//               {"x": 180, "y": 130},
	//
	//               {"x": 120, "y": 130}
	//
	//             ],
	//
	//             "user_answer": "2",
	//
	//             "standard_answer": "2"
	//
	//           }
	//
	//         ],
	//
	//         "solution": "1 + 1 = 2"
	//
	//       }
	//
	//     ],
	//
	//     "usage": {
	//
	//       "prompt_tokens": 1000,
	//
	//       "completion_tokens": 200,
	//
	//       "total_tokens": 1200
	//
	//     }
	//
	//   }
	//
	// }
	Data *GetEduHomeworkCorrectionResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// req-20261116-0001
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
	// The rotation angle of the image in the clockwise direction, in degrees.
	//
	// example:
	//
	// 0
	Angle *int64 `json:"angle,omitempty" xml:"angle,omitempty"`
	// The URL of the dewarped image. If dewarp is set to false, the URL of the original image is returned.
	//
	// example:
	//
	// https://example.com/dewarp.jpg
	DewarpUrl *string `json:"dewarpUrl,omitempty" xml:"dewarpUrl,omitempty"`
	// The array of question recognition and grading results. For more information, see the following table.
	QuestionResults []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResults `json:"questionResults,omitempty" xml:"questionResults,omitempty" type:"Repeated"`
	// The task status.
	//
	// example:
	//
	// submitted, processing, finished, or failed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The unique identifier of the task.
	//
	// example:
	//
	// task-20261116-abc123
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
	// The array of answering and grading results for the question. For more information, see the following table.
	AnswerResults []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResults `json:"answerResults,omitempty" xml:"answerResults,omitempty" type:"Repeated"`
	// Indicates whether the task is finished:
	//
	// - **true**: The task is finished.
	//
	// - **false**: The task is not finished.
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
	// The coordinate group of the question area. It generally contains four vertex coordinate values, arranged in the order of top-left, top-right, bottom-right, and bottom-left.
	//
	// example:
	//
	// `[[{"x": 120, "y": 980}, ...]]`
	QuestionPoints [][]*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsQuestionPoints `json:"questionPoints,omitempty" xml:"questionPoints,omitempty" type:"Repeated"`
	// The question type.
	//
	// example:
	//
	// Multiple-choice question, applied subjective question, problem-solving question, calculation question, etc
	QuestionType *string `json:"questionType,omitempty" xml:"questionType,omitempty"`
	// The question analysis or solution.
	//
	// example:
	//
	// Question analysis or solution
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
	// The coordinate group of the user\\"s answering area. It generally contains four vertex coordinate values, arranged in the order of top-left, top-right, bottom-right, and bottom-left, such as `[{"x": ..., "y": ...}, ...]`.
	AnswerPoints []*GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsAnswerPoints `json:"answerPoints,omitempty" xml:"answerPoints,omitempty" type:"Repeated"`
	// The coordinates of the grading mark point, such as `{ "x": ..., "y": ... }`.
	CheckPoint *GetEduHomeworkCorrectionResultResponseBodyDataQuestionResultsAnswerResultsCheckPoint `json:"checkPoint,omitempty" xml:"checkPoint,omitempty" type:"Struct"`
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
	// 120
	StandardAnswer *string `json:"standardAnswer,omitempty" xml:"standardAnswer,omitempty"`
	// The student\\"s answer.
	//
	// example:
	//
	// 980
	UserAnswer *string `json:"userAnswer,omitempty" xml:"userAnswer,omitempty"`
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
