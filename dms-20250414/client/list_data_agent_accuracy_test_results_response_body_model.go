// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentAccuracyTestResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataAgentAccuracyTestResultsResponseBodyData) *ListDataAgentAccuracyTestResultsResponseBody
	GetData() *ListDataAgentAccuracyTestResultsResponseBodyData
	SetErrorCode(v string) *ListDataAgentAccuracyTestResultsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataAgentAccuracyTestResultsResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataAgentAccuracyTestResultsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentAccuracyTestResultsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDataAgentAccuracyTestResultsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataAgentAccuracyTestResultsResponseBody
	GetSuccess() *bool
}

type ListDataAgentAccuracyTestResultsResponseBody struct {
	// The response struct.
	Data *ListDataAgentAccuracyTestResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// no use
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// no use
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataAgentAccuracyTestResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) GetData() *ListDataAgentAccuracyTestResultsResponseBodyData {
	return s.Data
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) SetData(v *ListDataAgentAccuracyTestResultsResponseBodyData) *ListDataAgentAccuracyTestResultsResponseBody {
	s.Data = v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) SetErrorCode(v string) *ListDataAgentAccuracyTestResultsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) SetErrorMessage(v string) *ListDataAgentAccuracyTestResultsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) SetMaxResults(v int32) *ListDataAgentAccuracyTestResultsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) SetNextToken(v string) *ListDataAgentAccuracyTestResultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) SetRequestId(v string) *ListDataAgentAccuracyTestResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) SetSuccess(v bool) *ListDataAgentAccuracyTestResultsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataAgentAccuracyTestResultsResponseBodyData struct {
	// The accuracy rate.
	//
	// example:
	//
	// 90
	AccuracyRate *float64 `json:"AccuracyRate,omitempty" xml:"AccuracyRate,omitempty"`
	// The ID of the accuracy test task.
	//
	// example:
	//
	// 692abb8f-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	AccuracyTestTaskId *string `json:"AccuracyTestTaskId,omitempty" xml:"AccuracyTestTaskId,omitempty"`
	// The data content.
	Content []*ListDataAgentAccuracyTestResultsResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The number of test cases that passed evaluation.
	//
	// example:
	//
	// 9
	CorrectCount *int64  `json:"CorrectCount,omitempty" xml:"CorrectCount,omitempty"`
	FailedCount  *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PendingCount *string `json:"PendingCount,omitempty" xml:"PendingCount,omitempty"`
	// The total number of results.
	//
	// example:
	//
	// 10
	TotalElements *int32 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListDataAgentAccuracyTestResultsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) GetAccuracyRate() *float64 {
	return s.AccuracyRate
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) GetAccuracyTestTaskId() *string {
	return s.AccuracyTestTaskId
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) GetContent() []*ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	return s.Content
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) GetCorrectCount() *int64 {
	return s.CorrectCount
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) GetFailedCount() *string {
	return s.FailedCount
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) GetPendingCount() *string {
	return s.PendingCount
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) GetTotalElements() *int32 {
	return s.TotalElements
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) SetAccuracyRate(v float64) *ListDataAgentAccuracyTestResultsResponseBodyData {
	s.AccuracyRate = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) SetAccuracyTestTaskId(v string) *ListDataAgentAccuracyTestResultsResponseBodyData {
	s.AccuracyTestTaskId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) SetContent(v []*ListDataAgentAccuracyTestResultsResponseBodyDataContent) *ListDataAgentAccuracyTestResultsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) SetCorrectCount(v int64) *ListDataAgentAccuracyTestResultsResponseBodyData {
	s.CorrectCount = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) SetFailedCount(v string) *ListDataAgentAccuracyTestResultsResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) SetPageNumber(v int64) *ListDataAgentAccuracyTestResultsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) SetPageSize(v int64) *ListDataAgentAccuracyTestResultsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) SetPendingCount(v string) *ListDataAgentAccuracyTestResultsResponseBodyData {
	s.PendingCount = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) SetTotalElements(v int32) *ListDataAgentAccuracyTestResultsResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) SetTotalPages(v int32) *ListDataAgentAccuracyTestResultsResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataAgentAccuracyTestResultsResponseBodyDataContent struct {
	// The ID of the accuracy test task.
	//
	// example:
	//
	// 692abb8f-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	AccuracyTestTaskId *string `json:"AccuracyTestTaskId,omitempty" xml:"AccuracyTestTaskId,omitempty"`
	// The actual answer from the agent.
	//
	// example:
	//
	// 在公司历史职位记录中，共有97,750名员工曾拥有Senior Engineer头衔。
	AgentResult *string `json:"AgentResult,omitempty" xml:"AgentResult,omitempty"`
	AgentSql    *string `json:"AgentSql,omitempty" xml:"AgentSql,omitempty"`
	// The expected answer.
	//
	// example:
	//
	// 97750
	AnswerResult *string `json:"AnswerResult,omitempty" xml:"AnswerResult,omitempty"`
	// The expected SQL.
	//
	// example:
	//
	// SELECT COUNT(*) FROM titles WHERE title = \\"Senior Engineer\\";
	AnswerSql *string `json:"AnswerSql,omitempty" xml:"AnswerSql,omitempty"`
	// The AI evaluation result.
	//
	// example:
	//
	// true
	IsTrue *bool `json:"IsTrue,omitempty" xml:"IsTrue,omitempty"`
	// The test question.
	//
	// example:
	//
	// 拥有Senior Engineer头衔的员工有多少人？
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// The error reason.
	//
	// example:
	//
	// SQL 中不应该使用COUNT(DISTINCT)
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The improvement suggestion.
	//
	// example:
	//
	// 在问题中描述清楚是否需去重
	Recommendation *string `json:"Recommendation,omitempty" xml:"Recommendation,omitempty"`
	// The result ID.
	//
	// example:
	//
	// at-emhnbwewfngrxxxxxxxxxx
	ResultId  *string `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The subtask ID.
	//
	// example:
	//
	// f1eb8728-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	SubtaskId *string `json:"SubtaskId,omitempty" xml:"SubtaskId,omitempty"`
}

func (s ListDataAgentAccuracyTestResultsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestResultsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetAccuracyTestTaskId() *string {
	return s.AccuracyTestTaskId
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetAgentResult() *string {
	return s.AgentResult
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetAgentSql() *string {
	return s.AgentSql
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetAnswerResult() *string {
	return s.AnswerResult
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetAnswerSql() *string {
	return s.AnswerSql
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetIsTrue() *bool {
	return s.IsTrue
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetQuestion() *string {
	return s.Question
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetReason() *string {
	return s.Reason
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetRecommendation() *string {
	return s.Recommendation
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetResultId() *string {
	return s.ResultId
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetSessionId() *string {
	return s.SessionId
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) GetSubtaskId() *string {
	return s.SubtaskId
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetAccuracyTestTaskId(v string) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.AccuracyTestTaskId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetAgentResult(v string) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.AgentResult = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetAgentSql(v string) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.AgentSql = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetAnswerResult(v string) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.AnswerResult = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetAnswerSql(v string) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.AnswerSql = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetIsTrue(v bool) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.IsTrue = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetQuestion(v string) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.Question = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetReason(v string) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.Reason = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetRecommendation(v string) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.Recommendation = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetResultId(v string) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.ResultId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetSessionId(v string) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.SessionId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) SetSubtaskId(v string) *ListDataAgentAccuracyTestResultsResponseBodyDataContent {
	s.SubtaskId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
