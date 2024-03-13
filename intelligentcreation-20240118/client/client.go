// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActualDeductResourceCmd struct {
	Cost         *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	ExtraInfo    *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	TaskId       *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ActualDeductResourceCmd) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourceCmd) GoString() string {
	return s.String()
}

func (s *ActualDeductResourceCmd) SetCost(v int64) *ActualDeductResourceCmd {
	s.Cost = &v
	return s
}

func (s *ActualDeductResourceCmd) SetExtraInfo(v string) *ActualDeductResourceCmd {
	s.ExtraInfo = &v
	return s
}

func (s *ActualDeductResourceCmd) SetIdempotentId(v string) *ActualDeductResourceCmd {
	s.IdempotentId = &v
	return s
}

func (s *ActualDeductResourceCmd) SetTaskId(v string) *ActualDeductResourceCmd {
	s.TaskId = &v
	return s
}

type ActualDeductResourceResult struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ActualDeductResourceResult) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourceResult) GoString() string {
	return s.String()
}

func (s *ActualDeductResourceResult) SetErrorCode(v string) *ActualDeductResourceResult {
	s.ErrorCode = &v
	return s
}

func (s *ActualDeductResourceResult) SetErrorMessage(v string) *ActualDeductResourceResult {
	s.ErrorMessage = &v
	return s
}

func (s *ActualDeductResourceResult) SetRequestId(v string) *ActualDeductResourceResult {
	s.RequestId = &v
	return s
}

func (s *ActualDeductResourceResult) SetSuccess(v bool) *ActualDeductResourceResult {
	s.Success = &v
	return s
}

type DigitalHumanLiveBroadcastQACmd struct {
	AccountId    *string                                  `json:"accountId,omitempty" xml:"accountId,omitempty"`
	History      []*DigitalHumanLiveBroadcastQACmdHistory `json:"history,omitempty" xml:"history,omitempty" type:"Repeated"`
	Question     *string                                  `json:"question,omitempty" xml:"question,omitempty"`
	SessionId    *string                                  `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Stream       *bool                                    `json:"stream,omitempty" xml:"stream,omitempty"`
	SubAccountId *string                                  `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
}

func (s DigitalHumanLiveBroadcastQACmd) String() string {
	return tea.Prettify(s)
}

func (s DigitalHumanLiveBroadcastQACmd) GoString() string {
	return s.String()
}

func (s *DigitalHumanLiveBroadcastQACmd) SetAccountId(v string) *DigitalHumanLiveBroadcastQACmd {
	s.AccountId = &v
	return s
}

func (s *DigitalHumanLiveBroadcastQACmd) SetHistory(v []*DigitalHumanLiveBroadcastQACmdHistory) *DigitalHumanLiveBroadcastQACmd {
	s.History = v
	return s
}

func (s *DigitalHumanLiveBroadcastQACmd) SetQuestion(v string) *DigitalHumanLiveBroadcastQACmd {
	s.Question = &v
	return s
}

func (s *DigitalHumanLiveBroadcastQACmd) SetSessionId(v string) *DigitalHumanLiveBroadcastQACmd {
	s.SessionId = &v
	return s
}

func (s *DigitalHumanLiveBroadcastQACmd) SetStream(v bool) *DigitalHumanLiveBroadcastQACmd {
	s.Stream = &v
	return s
}

func (s *DigitalHumanLiveBroadcastQACmd) SetSubAccountId(v string) *DigitalHumanLiveBroadcastQACmd {
	s.SubAccountId = &v
	return s
}

type DigitalHumanLiveBroadcastQACmdHistory struct {
	Bot  *string `json:"bot,omitempty" xml:"bot,omitempty"`
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s DigitalHumanLiveBroadcastQACmdHistory) String() string {
	return tea.Prettify(s)
}

func (s DigitalHumanLiveBroadcastQACmdHistory) GoString() string {
	return s.String()
}

func (s *DigitalHumanLiveBroadcastQACmdHistory) SetBot(v string) *DigitalHumanLiveBroadcastQACmdHistory {
	s.Bot = &v
	return s
}

func (s *DigitalHumanLiveBroadcastQACmdHistory) SetUser(v string) *DigitalHumanLiveBroadcastQACmdHistory {
	s.User = &v
	return s
}

type DigitalHumanLiveBroadcastQAResult struct {
	Content      *string `json:"content,omitempty" xml:"content,omitempty"`
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	SessionId    *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DigitalHumanLiveBroadcastQAResult) String() string {
	return tea.Prettify(s)
}

func (s DigitalHumanLiveBroadcastQAResult) GoString() string {
	return s.String()
}

func (s *DigitalHumanLiveBroadcastQAResult) SetContent(v string) *DigitalHumanLiveBroadcastQAResult {
	s.Content = &v
	return s
}

func (s *DigitalHumanLiveBroadcastQAResult) SetErrorCode(v string) *DigitalHumanLiveBroadcastQAResult {
	s.ErrorCode = &v
	return s
}

func (s *DigitalHumanLiveBroadcastQAResult) SetErrorMessage(v string) *DigitalHumanLiveBroadcastQAResult {
	s.ErrorMessage = &v
	return s
}

func (s *DigitalHumanLiveBroadcastQAResult) SetSessionId(v string) *DigitalHumanLiveBroadcastQAResult {
	s.SessionId = &v
	return s
}

func (s *DigitalHumanLiveBroadcastQAResult) SetSuccess(v bool) *DigitalHumanLiveBroadcastQAResult {
	s.Success = &v
	return s
}

type DirectDeductResourceCmd struct {
	AccountId    *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Cost         *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	DeductScene  *string `json:"deductScene,omitempty" xml:"deductScene,omitempty"`
	ExtraInfo    *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	ResourceType *int64  `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
	Token        *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s DirectDeductResourceCmd) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourceCmd) GoString() string {
	return s.String()
}

func (s *DirectDeductResourceCmd) SetAccountId(v string) *DirectDeductResourceCmd {
	s.AccountId = &v
	return s
}

func (s *DirectDeductResourceCmd) SetCost(v int64) *DirectDeductResourceCmd {
	s.Cost = &v
	return s
}

func (s *DirectDeductResourceCmd) SetDeductScene(v string) *DirectDeductResourceCmd {
	s.DeductScene = &v
	return s
}

func (s *DirectDeductResourceCmd) SetExtraInfo(v string) *DirectDeductResourceCmd {
	s.ExtraInfo = &v
	return s
}

func (s *DirectDeductResourceCmd) SetIdempotentId(v string) *DirectDeductResourceCmd {
	s.IdempotentId = &v
	return s
}

func (s *DirectDeductResourceCmd) SetResourceType(v int64) *DirectDeductResourceCmd {
	s.ResourceType = &v
	return s
}

func (s *DirectDeductResourceCmd) SetSubAccountId(v string) *DirectDeductResourceCmd {
	s.SubAccountId = &v
	return s
}

func (s *DirectDeductResourceCmd) SetToken(v string) *DirectDeductResourceCmd {
	s.Token = &v
	return s
}

type DirectDeductResourceResult struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DirectDeductResourceResult) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourceResult) GoString() string {
	return s.String()
}

func (s *DirectDeductResourceResult) SetErrorCode(v string) *DirectDeductResourceResult {
	s.ErrorCode = &v
	return s
}

func (s *DirectDeductResourceResult) SetErrorMessage(v string) *DirectDeductResourceResult {
	s.ErrorMessage = &v
	return s
}

func (s *DirectDeductResourceResult) SetRequestId(v string) *DirectDeductResourceResult {
	s.RequestId = &v
	return s
}

func (s *DirectDeductResourceResult) SetSuccess(v bool) *DirectDeductResourceResult {
	s.Success = &v
	return s
}

type ExpectDeductResourceCmd struct {
	AccountId    *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Cost         *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	DeductScene  *string `json:"deductScene,omitempty" xml:"deductScene,omitempty"`
	ExtraInfo    *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	ResourceType *int64  `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
	Token        *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ExpectDeductResourceCmd) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourceCmd) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourceCmd) SetAccountId(v string) *ExpectDeductResourceCmd {
	s.AccountId = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetCost(v int64) *ExpectDeductResourceCmd {
	s.Cost = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetDeductScene(v string) *ExpectDeductResourceCmd {
	s.DeductScene = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetExtraInfo(v string) *ExpectDeductResourceCmd {
	s.ExtraInfo = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetIdempotentId(v string) *ExpectDeductResourceCmd {
	s.IdempotentId = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetResourceType(v int64) *ExpectDeductResourceCmd {
	s.ResourceType = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetSubAccountId(v string) *ExpectDeductResourceCmd {
	s.SubAccountId = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetToken(v string) *ExpectDeductResourceCmd {
	s.Token = &v
	return s
}

type ExpectDeductResourceResult struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TaskId       *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ExpectDeductResourceResult) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourceResult) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourceResult) SetErrorCode(v string) *ExpectDeductResourceResult {
	s.ErrorCode = &v
	return s
}

func (s *ExpectDeductResourceResult) SetErrorMessage(v string) *ExpectDeductResourceResult {
	s.ErrorMessage = &v
	return s
}

func (s *ExpectDeductResourceResult) SetRequestId(v string) *ExpectDeductResourceResult {
	s.RequestId = &v
	return s
}

func (s *ExpectDeductResourceResult) SetSuccess(v bool) *ExpectDeductResourceResult {
	s.Success = &v
	return s
}

func (s *ExpectDeductResourceResult) SetTaskId(v string) *ExpectDeductResourceResult {
	s.TaskId = &v
	return s
}

type SubmitBulletQuestionsCmd struct {
	AccountId    *string                              `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Questions    []*SubmitBulletQuestionsCmdQuestions `json:"questions,omitempty" xml:"questions,omitempty" type:"Repeated"`
	RoomId       *string                              `json:"roomId,omitempty" xml:"roomId,omitempty"`
	SubAccountId *string                              `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
}

func (s SubmitBulletQuestionsCmd) String() string {
	return tea.Prettify(s)
}

func (s SubmitBulletQuestionsCmd) GoString() string {
	return s.String()
}

func (s *SubmitBulletQuestionsCmd) SetAccountId(v string) *SubmitBulletQuestionsCmd {
	s.AccountId = &v
	return s
}

func (s *SubmitBulletQuestionsCmd) SetQuestions(v []*SubmitBulletQuestionsCmdQuestions) *SubmitBulletQuestionsCmd {
	s.Questions = v
	return s
}

func (s *SubmitBulletQuestionsCmd) SetRoomId(v string) *SubmitBulletQuestionsCmd {
	s.RoomId = &v
	return s
}

func (s *SubmitBulletQuestionsCmd) SetSubAccountId(v string) *SubmitBulletQuestionsCmd {
	s.SubAccountId = &v
	return s
}

type SubmitBulletQuestionsCmdQuestions struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Time     *int64  `json:"time,omitempty" xml:"time,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s SubmitBulletQuestionsCmdQuestions) String() string {
	return tea.Prettify(s)
}

func (s SubmitBulletQuestionsCmdQuestions) GoString() string {
	return s.String()
}

func (s *SubmitBulletQuestionsCmdQuestions) SetContent(v string) *SubmitBulletQuestionsCmdQuestions {
	s.Content = &v
	return s
}

func (s *SubmitBulletQuestionsCmdQuestions) SetId(v string) *SubmitBulletQuestionsCmdQuestions {
	s.Id = &v
	return s
}

func (s *SubmitBulletQuestionsCmdQuestions) SetTime(v int64) *SubmitBulletQuestionsCmdQuestions {
	s.Time = &v
	return s
}

func (s *SubmitBulletQuestionsCmdQuestions) SetUsername(v string) *SubmitBulletQuestionsCmdQuestions {
	s.Username = &v
	return s
}

type SubmitBulletQuestionsQAResult struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitBulletQuestionsQAResult) String() string {
	return tea.Prettify(s)
}

func (s SubmitBulletQuestionsQAResult) GoString() string {
	return s.String()
}

func (s *SubmitBulletQuestionsQAResult) SetErrorCode(v string) *SubmitBulletQuestionsQAResult {
	s.ErrorCode = &v
	return s
}

func (s *SubmitBulletQuestionsQAResult) SetErrorMessage(v string) *SubmitBulletQuestionsQAResult {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitBulletQuestionsQAResult) SetSuccess(v bool) *SubmitBulletQuestionsQAResult {
	s.Success = &v
	return s
}

type ActualDeductResourceRequest struct {
	Body *ActualDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActualDeductResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourceRequest) GoString() string {
	return s.String()
}

func (s *ActualDeductResourceRequest) SetBody(v *ActualDeductResourceCmd) *ActualDeductResourceRequest {
	s.Body = v
	return s
}

type ActualDeductResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActualDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActualDeductResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourceResponse) GoString() string {
	return s.String()
}

func (s *ActualDeductResourceResponse) SetHeaders(v map[string]*string) *ActualDeductResourceResponse {
	s.Headers = v
	return s
}

func (s *ActualDeductResourceResponse) SetStatusCode(v int32) *ActualDeductResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ActualDeductResourceResponse) SetBody(v *ActualDeductResourceResult) *ActualDeductResourceResponse {
	s.Body = v
	return s
}

type ActualDeductResourcesRequest struct {
	Body *ActualDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActualDeductResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourcesRequest) GoString() string {
	return s.String()
}

func (s *ActualDeductResourcesRequest) SetBody(v *ActualDeductResourceCmd) *ActualDeductResourcesRequest {
	s.Body = v
	return s
}

type ActualDeductResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActualDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActualDeductResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourcesResponse) GoString() string {
	return s.String()
}

func (s *ActualDeductResourcesResponse) SetHeaders(v map[string]*string) *ActualDeductResourcesResponse {
	s.Headers = v
	return s
}

func (s *ActualDeductResourcesResponse) SetStatusCode(v int32) *ActualDeductResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ActualDeductResourcesResponse) SetBody(v *ActualDeductResourceResult) *ActualDeductResourcesResponse {
	s.Body = v
	return s
}

type CopywritingQARequest struct {
	AccountId *string                          `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Histories []*CopywritingQARequestHistories `json:"histories,omitempty" xml:"histories,omitempty" type:"Repeated"`
	// Deprecated
	History      *CopywritingQARequestHistory `json:"history,omitempty" xml:"history,omitempty" type:"Struct"`
	Question     *string                      `json:"question,omitempty" xml:"question,omitempty"`
	SessionId    *string                      `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Stream       *bool                        `json:"stream,omitempty" xml:"stream,omitempty"`
	SubAccountId *string                      `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
}

func (s CopywritingQARequest) String() string {
	return tea.Prettify(s)
}

func (s CopywritingQARequest) GoString() string {
	return s.String()
}

func (s *CopywritingQARequest) SetAccountId(v string) *CopywritingQARequest {
	s.AccountId = &v
	return s
}

func (s *CopywritingQARequest) SetHistories(v []*CopywritingQARequestHistories) *CopywritingQARequest {
	s.Histories = v
	return s
}

func (s *CopywritingQARequest) SetHistory(v *CopywritingQARequestHistory) *CopywritingQARequest {
	s.History = v
	return s
}

func (s *CopywritingQARequest) SetQuestion(v string) *CopywritingQARequest {
	s.Question = &v
	return s
}

func (s *CopywritingQARequest) SetSessionId(v string) *CopywritingQARequest {
	s.SessionId = &v
	return s
}

func (s *CopywritingQARequest) SetStream(v bool) *CopywritingQARequest {
	s.Stream = &v
	return s
}

func (s *CopywritingQARequest) SetSubAccountId(v string) *CopywritingQARequest {
	s.SubAccountId = &v
	return s
}

type CopywritingQARequestHistories struct {
	Bot  *string `json:"bot,omitempty" xml:"bot,omitempty"`
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s CopywritingQARequestHistories) String() string {
	return tea.Prettify(s)
}

func (s CopywritingQARequestHistories) GoString() string {
	return s.String()
}

func (s *CopywritingQARequestHistories) SetBot(v string) *CopywritingQARequestHistories {
	s.Bot = &v
	return s
}

func (s *CopywritingQARequestHistories) SetUser(v string) *CopywritingQARequestHistories {
	s.User = &v
	return s
}

type CopywritingQARequestHistory struct {
	Bot  *string `json:"bot,omitempty" xml:"bot,omitempty"`
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s CopywritingQARequestHistory) String() string {
	return tea.Prettify(s)
}

func (s CopywritingQARequestHistory) GoString() string {
	return s.String()
}

func (s *CopywritingQARequestHistory) SetBot(v string) *CopywritingQARequestHistory {
	s.Bot = &v
	return s
}

func (s *CopywritingQARequestHistory) SetUser(v string) *CopywritingQARequestHistory {
	s.User = &v
	return s
}

type CopywritingQAShrinkRequest struct {
	AccountId       *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	HistoriesShrink *string `json:"histories,omitempty" xml:"histories,omitempty"`
	// Deprecated
	HistoryShrink *string `json:"history,omitempty" xml:"history,omitempty"`
	Question      *string `json:"question,omitempty" xml:"question,omitempty"`
	SessionId     *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Stream        *bool   `json:"stream,omitempty" xml:"stream,omitempty"`
	SubAccountId  *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
}

func (s CopywritingQAShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CopywritingQAShrinkRequest) GoString() string {
	return s.String()
}

func (s *CopywritingQAShrinkRequest) SetAccountId(v string) *CopywritingQAShrinkRequest {
	s.AccountId = &v
	return s
}

func (s *CopywritingQAShrinkRequest) SetHistoriesShrink(v string) *CopywritingQAShrinkRequest {
	s.HistoriesShrink = &v
	return s
}

func (s *CopywritingQAShrinkRequest) SetHistoryShrink(v string) *CopywritingQAShrinkRequest {
	s.HistoryShrink = &v
	return s
}

func (s *CopywritingQAShrinkRequest) SetQuestion(v string) *CopywritingQAShrinkRequest {
	s.Question = &v
	return s
}

func (s *CopywritingQAShrinkRequest) SetSessionId(v string) *CopywritingQAShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *CopywritingQAShrinkRequest) SetStream(v bool) *CopywritingQAShrinkRequest {
	s.Stream = &v
	return s
}

func (s *CopywritingQAShrinkRequest) SetSubAccountId(v string) *CopywritingQAShrinkRequest {
	s.SubAccountId = &v
	return s
}

type CopywritingQAResponseBody struct {
	Content      *string `json:"content,omitempty" xml:"content,omitempty"`
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	SessionId    *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CopywritingQAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopywritingQAResponseBody) GoString() string {
	return s.String()
}

func (s *CopywritingQAResponseBody) SetContent(v string) *CopywritingQAResponseBody {
	s.Content = &v
	return s
}

func (s *CopywritingQAResponseBody) SetErrorCode(v string) *CopywritingQAResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CopywritingQAResponseBody) SetErrorMessage(v string) *CopywritingQAResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CopywritingQAResponseBody) SetSessionId(v string) *CopywritingQAResponseBody {
	s.SessionId = &v
	return s
}

func (s *CopywritingQAResponseBody) SetSuccess(v bool) *CopywritingQAResponseBody {
	s.Success = &v
	return s
}

type CopywritingQAResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopywritingQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopywritingQAResponse) String() string {
	return tea.Prettify(s)
}

func (s CopywritingQAResponse) GoString() string {
	return s.String()
}

func (s *CopywritingQAResponse) SetHeaders(v map[string]*string) *CopywritingQAResponse {
	s.Headers = v
	return s
}

func (s *CopywritingQAResponse) SetStatusCode(v int32) *CopywritingQAResponse {
	s.StatusCode = &v
	return s
}

func (s *CopywritingQAResponse) SetBody(v *CopywritingQAResponseBody) *CopywritingQAResponse {
	s.Body = v
	return s
}

type CopywritingQAV1Request struct {
	Body *DigitalHumanLiveBroadcastQACmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopywritingQAV1Request) String() string {
	return tea.Prettify(s)
}

func (s CopywritingQAV1Request) GoString() string {
	return s.String()
}

func (s *CopywritingQAV1Request) SetBody(v *DigitalHumanLiveBroadcastQACmd) *CopywritingQAV1Request {
	s.Body = v
	return s
}

type CopywritingQAV1Response struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DigitalHumanLiveBroadcastQAResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopywritingQAV1Response) String() string {
	return tea.Prettify(s)
}

func (s CopywritingQAV1Response) GoString() string {
	return s.String()
}

func (s *CopywritingQAV1Response) SetHeaders(v map[string]*string) *CopywritingQAV1Response {
	s.Headers = v
	return s
}

func (s *CopywritingQAV1Response) SetStatusCode(v int32) *CopywritingQAV1Response {
	s.StatusCode = &v
	return s
}

func (s *CopywritingQAV1Response) SetBody(v *DigitalHumanLiveBroadcastQAResult) *CopywritingQAV1Response {
	s.Body = v
	return s
}

type DeleteDigitalVideoRequest struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	VideoId   *string `json:"videoId,omitempty" xml:"videoId,omitempty"`
}

func (s DeleteDigitalVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDigitalVideoRequest) GoString() string {
	return s.String()
}

func (s *DeleteDigitalVideoRequest) SetAccountId(v string) *DeleteDigitalVideoRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteDigitalVideoRequest) SetVideoId(v string) *DeleteDigitalVideoRequest {
	s.VideoId = &v
	return s
}

type DeleteDigitalVideoResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDigitalVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDigitalVideoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDigitalVideoResponseBody) SetErrorCode(v string) *DeleteDigitalVideoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDigitalVideoResponseBody) SetErrorMessage(v string) *DeleteDigitalVideoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDigitalVideoResponseBody) SetRequestId(v string) *DeleteDigitalVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDigitalVideoResponseBody) SetSuccess(v bool) *DeleteDigitalVideoResponseBody {
	s.Success = &v
	return s
}

type DeleteDigitalVideoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDigitalVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDigitalVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDigitalVideoResponse) GoString() string {
	return s.String()
}

func (s *DeleteDigitalVideoResponse) SetHeaders(v map[string]*string) *DeleteDigitalVideoResponse {
	s.Headers = v
	return s
}

func (s *DeleteDigitalVideoResponse) SetStatusCode(v int32) *DeleteDigitalVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDigitalVideoResponse) SetBody(v *DeleteDigitalVideoResponseBody) *DeleteDigitalVideoResponse {
	s.Body = v
	return s
}

type DirectDeductResourceRequest struct {
	Body *DirectDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DirectDeductResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourceRequest) GoString() string {
	return s.String()
}

func (s *DirectDeductResourceRequest) SetBody(v *DirectDeductResourceCmd) *DirectDeductResourceRequest {
	s.Body = v
	return s
}

type DirectDeductResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DirectDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DirectDeductResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourceResponse) GoString() string {
	return s.String()
}

func (s *DirectDeductResourceResponse) SetHeaders(v map[string]*string) *DirectDeductResourceResponse {
	s.Headers = v
	return s
}

func (s *DirectDeductResourceResponse) SetStatusCode(v int32) *DirectDeductResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DirectDeductResourceResponse) SetBody(v *DirectDeductResourceResult) *DirectDeductResourceResponse {
	s.Body = v
	return s
}

type DirectDeductResourcesRequest struct {
	Body *DirectDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DirectDeductResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourcesRequest) GoString() string {
	return s.String()
}

func (s *DirectDeductResourcesRequest) SetBody(v *DirectDeductResourceCmd) *DirectDeductResourcesRequest {
	s.Body = v
	return s
}

type DirectDeductResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DirectDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DirectDeductResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourcesResponse) GoString() string {
	return s.String()
}

func (s *DirectDeductResourcesResponse) SetHeaders(v map[string]*string) *DirectDeductResourcesResponse {
	s.Headers = v
	return s
}

func (s *DirectDeductResourcesResponse) SetStatusCode(v int32) *DirectDeductResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DirectDeductResourcesResponse) SetBody(v *DirectDeductResourceResult) *DirectDeductResourcesResponse {
	s.Body = v
	return s
}

type ExpectDeductResourceRequest struct {
	Body *ExpectDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpectDeductResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourceRequest) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourceRequest) SetBody(v *ExpectDeductResourceCmd) *ExpectDeductResourceRequest {
	s.Body = v
	return s
}

type ExpectDeductResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExpectDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpectDeductResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourceResponse) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourceResponse) SetHeaders(v map[string]*string) *ExpectDeductResourceResponse {
	s.Headers = v
	return s
}

func (s *ExpectDeductResourceResponse) SetStatusCode(v int32) *ExpectDeductResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ExpectDeductResourceResponse) SetBody(v *ExpectDeductResourceResult) *ExpectDeductResourceResponse {
	s.Body = v
	return s
}

type ExpectDeductResourcesRequest struct {
	Body *ExpectDeductResourceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpectDeductResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourcesRequest) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourcesRequest) SetBody(v *ExpectDeductResourceCmd) *ExpectDeductResourcesRequest {
	s.Body = v
	return s
}

type ExpectDeductResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExpectDeductResourceResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpectDeductResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourcesResponse) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourcesResponse) SetHeaders(v map[string]*string) *ExpectDeductResourcesResponse {
	s.Headers = v
	return s
}

func (s *ExpectDeductResourcesResponse) SetStatusCode(v int32) *ExpectDeductResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ExpectDeductResourcesResponse) SetBody(v *ExpectDeductResourceResult) *ExpectDeductResourcesResponse {
	s.Body = v
	return s
}

type GetRemainResourceRequest struct {
	AccountId    *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
}

func (s GetRemainResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRemainResourceRequest) GoString() string {
	return s.String()
}

func (s *GetRemainResourceRequest) SetAccountId(v string) *GetRemainResourceRequest {
	s.AccountId = &v
	return s
}

func (s *GetRemainResourceRequest) SetResourceType(v string) *GetRemainResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *GetRemainResourceRequest) SetSubAccountId(v string) *GetRemainResourceRequest {
	s.SubAccountId = &v
	return s
}

type GetRemainResourceResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	RemainCount  *int32  `json:"remainCount,omitempty" xml:"remainCount,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetRemainResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRemainResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetRemainResourceResponseBody) SetErrorCode(v string) *GetRemainResourceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRemainResourceResponseBody) SetErrorMessage(v string) *GetRemainResourceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRemainResourceResponseBody) SetRemainCount(v int32) *GetRemainResourceResponseBody {
	s.RemainCount = &v
	return s
}

func (s *GetRemainResourceResponseBody) SetRequestId(v string) *GetRemainResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRemainResourceResponseBody) SetSuccess(v bool) *GetRemainResourceResponseBody {
	s.Success = &v
	return s
}

type GetRemainResourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRemainResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRemainResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRemainResourceResponse) GoString() string {
	return s.String()
}

func (s *GetRemainResourceResponse) SetHeaders(v map[string]*string) *GetRemainResourceResponse {
	s.Headers = v
	return s
}

func (s *GetRemainResourceResponse) SetStatusCode(v int32) *GetRemainResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRemainResourceResponse) SetBody(v *GetRemainResourceResponseBody) *GetRemainResourceResponse {
	s.Body = v
	return s
}

type SubmitBulletQuestionsRequest struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// questions
	Questions    []*SubmitBulletQuestionsRequestQuestions `json:"questions,omitempty" xml:"questions,omitempty" type:"Repeated"`
	RoomId       *string                                  `json:"roomId,omitempty" xml:"roomId,omitempty"`
	SubAccountId *string                                  `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
}

func (s SubmitBulletQuestionsRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitBulletQuestionsRequest) GoString() string {
	return s.String()
}

func (s *SubmitBulletQuestionsRequest) SetAccountId(v string) *SubmitBulletQuestionsRequest {
	s.AccountId = &v
	return s
}

func (s *SubmitBulletQuestionsRequest) SetQuestions(v []*SubmitBulletQuestionsRequestQuestions) *SubmitBulletQuestionsRequest {
	s.Questions = v
	return s
}

func (s *SubmitBulletQuestionsRequest) SetRoomId(v string) *SubmitBulletQuestionsRequest {
	s.RoomId = &v
	return s
}

func (s *SubmitBulletQuestionsRequest) SetSubAccountId(v string) *SubmitBulletQuestionsRequest {
	s.SubAccountId = &v
	return s
}

type SubmitBulletQuestionsRequestQuestions struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Time     *int64  `json:"time,omitempty" xml:"time,omitempty"`
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s SubmitBulletQuestionsRequestQuestions) String() string {
	return tea.Prettify(s)
}

func (s SubmitBulletQuestionsRequestQuestions) GoString() string {
	return s.String()
}

func (s *SubmitBulletQuestionsRequestQuestions) SetContent(v string) *SubmitBulletQuestionsRequestQuestions {
	s.Content = &v
	return s
}

func (s *SubmitBulletQuestionsRequestQuestions) SetId(v string) *SubmitBulletQuestionsRequestQuestions {
	s.Id = &v
	return s
}

func (s *SubmitBulletQuestionsRequestQuestions) SetTime(v int64) *SubmitBulletQuestionsRequestQuestions {
	s.Time = &v
	return s
}

func (s *SubmitBulletQuestionsRequestQuestions) SetUsername(v string) *SubmitBulletQuestionsRequestQuestions {
	s.Username = &v
	return s
}

type SubmitBulletQuestionsShrinkRequest struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// questions
	QuestionsShrink *string `json:"questions,omitempty" xml:"questions,omitempty"`
	RoomId          *string `json:"roomId,omitempty" xml:"roomId,omitempty"`
	SubAccountId    *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
}

func (s SubmitBulletQuestionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitBulletQuestionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitBulletQuestionsShrinkRequest) SetAccountId(v string) *SubmitBulletQuestionsShrinkRequest {
	s.AccountId = &v
	return s
}

func (s *SubmitBulletQuestionsShrinkRequest) SetQuestionsShrink(v string) *SubmitBulletQuestionsShrinkRequest {
	s.QuestionsShrink = &v
	return s
}

func (s *SubmitBulletQuestionsShrinkRequest) SetRoomId(v string) *SubmitBulletQuestionsShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *SubmitBulletQuestionsShrinkRequest) SetSubAccountId(v string) *SubmitBulletQuestionsShrinkRequest {
	s.SubAccountId = &v
	return s
}

type SubmitBulletQuestionsResponseBody struct {
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitBulletQuestionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitBulletQuestionsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitBulletQuestionsResponseBody) SetErrorCode(v string) *SubmitBulletQuestionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitBulletQuestionsResponseBody) SetErrorMessage(v string) *SubmitBulletQuestionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitBulletQuestionsResponseBody) SetSuccess(v bool) *SubmitBulletQuestionsResponseBody {
	s.Success = &v
	return s
}

type SubmitBulletQuestionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitBulletQuestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitBulletQuestionsResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitBulletQuestionsResponse) GoString() string {
	return s.String()
}

func (s *SubmitBulletQuestionsResponse) SetHeaders(v map[string]*string) *SubmitBulletQuestionsResponse {
	s.Headers = v
	return s
}

func (s *SubmitBulletQuestionsResponse) SetStatusCode(v int32) *SubmitBulletQuestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitBulletQuestionsResponse) SetBody(v *SubmitBulletQuestionsResponseBody) *SubmitBulletQuestionsResponse {
	s.Body = v
	return s
}

type SubmitBulletQuestionsV1Request struct {
	Body *SubmitBulletQuestionsCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitBulletQuestionsV1Request) String() string {
	return tea.Prettify(s)
}

func (s SubmitBulletQuestionsV1Request) GoString() string {
	return s.String()
}

func (s *SubmitBulletQuestionsV1Request) SetBody(v *SubmitBulletQuestionsCmd) *SubmitBulletQuestionsV1Request {
	s.Body = v
	return s
}

type SubmitBulletQuestionsV1Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitBulletQuestionsQAResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitBulletQuestionsV1Response) String() string {
	return tea.Prettify(s)
}

func (s SubmitBulletQuestionsV1Response) GoString() string {
	return s.String()
}

func (s *SubmitBulletQuestionsV1Response) SetHeaders(v map[string]*string) *SubmitBulletQuestionsV1Response {
	s.Headers = v
	return s
}

func (s *SubmitBulletQuestionsV1Response) SetStatusCode(v int32) *SubmitBulletQuestionsV1Response {
	s.StatusCode = &v
	return s
}

func (s *SubmitBulletQuestionsV1Response) SetBody(v *SubmitBulletQuestionsQAResult) *SubmitBulletQuestionsV1Response {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("intelligentcreation"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActualDeductResourceWithOptions(request *ActualDeductResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ActualDeductResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ActualDeductResource"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/digital-human/commands/actual-deduct"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ActualDeductResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActualDeductResource(request *ActualDeductResourceRequest) (_result *ActualDeductResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ActualDeductResourceResponse{}
	_body, _err := client.ActualDeductResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActualDeductResourcesWithOptions(request *ActualDeductResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ActualDeductResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ActualDeductResources"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/actualDeductResources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ActualDeductResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActualDeductResources(request *ActualDeductResourcesRequest) (_result *ActualDeductResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ActualDeductResourcesResponse{}
	_body, _err := client.ActualDeductResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopywritingQAWithOptions(tmpReq *CopywritingQARequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CopywritingQAResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CopywritingQAShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Histories)) {
		request.HistoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Histories, tea.String("histories"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.History)) {
		request.HistoryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.History, tea.String("history"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.HistoriesShrink)) {
		query["histories"] = request.HistoriesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryShrink)) {
		query["history"] = request.HistoryShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Question)) {
		query["question"] = request.Question
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		query["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.SubAccountId)) {
		query["subAccountId"] = request.SubAccountId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CopywritingQA"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/copywritingQA"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CopywritingQAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopywritingQA(request *CopywritingQARequest) (_result *CopywritingQAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CopywritingQAResponse{}
	_body, _err := client.CopywritingQAWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopywritingQAV1WithOptions(request *CopywritingQAV1Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CopywritingQAV1Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CopywritingQAV1"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/copywritingQAV1"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CopywritingQAV1Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopywritingQAV1(request *CopywritingQAV1Request) (_result *CopywritingQAV1Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CopywritingQAV1Response{}
	_body, _err := client.CopywritingQAV1WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDigitalVideoWithOptions(request *DeleteDigitalVideoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDigitalVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoId)) {
		body["videoId"] = request.VideoId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDigitalVideo"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/videos"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDigitalVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDigitalVideo(request *DeleteDigitalVideoRequest) (_result *DeleteDigitalVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDigitalVideoResponse{}
	_body, _err := client.DeleteDigitalVideoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DirectDeductResourceWithOptions(request *DirectDeductResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DirectDeductResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("DirectDeductResource"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/digital-human/commands/direct-deduct"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DirectDeductResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DirectDeductResource(request *DirectDeductResourceRequest) (_result *DirectDeductResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DirectDeductResourceResponse{}
	_body, _err := client.DirectDeductResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DirectDeductResourcesWithOptions(request *DirectDeductResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DirectDeductResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("DirectDeductResources"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/directDeductResources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DirectDeductResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DirectDeductResources(request *DirectDeductResourcesRequest) (_result *DirectDeductResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DirectDeductResourcesResponse{}
	_body, _err := client.DirectDeductResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExpectDeductResourceWithOptions(request *ExpectDeductResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExpectDeductResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExpectDeductResource"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/v1/digital-human/commands/expect-deduct"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExpectDeductResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExpectDeductResource(request *ExpectDeductResourceRequest) (_result *ExpectDeductResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExpectDeductResourceResponse{}
	_body, _err := client.ExpectDeductResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExpectDeductResourcesWithOptions(request *ExpectDeductResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExpectDeductResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExpectDeductResources"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/expectDeductResources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExpectDeductResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExpectDeductResources(request *ExpectDeductResourcesRequest) (_result *ExpectDeductResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExpectDeductResourcesResponse{}
	_body, _err := client.ExpectDeductResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRemainResourceWithOptions(request *GetRemainResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRemainResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SubAccountId)) {
		query["subAccountId"] = request.SubAccountId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRemainResource"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/getRemainResource"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRemainResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRemainResource(request *GetRemainResourceRequest) (_result *GetRemainResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRemainResourceResponse{}
	_body, _err := client.GetRemainResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitBulletQuestionsWithOptions(tmpReq *SubmitBulletQuestionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitBulletQuestionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitBulletQuestionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Questions)) {
		request.QuestionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Questions, tea.String("questions"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.QuestionsShrink)) {
		query["questions"] = request.QuestionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RoomId)) {
		query["roomId"] = request.RoomId
	}

	if !tea.BoolValue(util.IsUnset(request.SubAccountId)) {
		query["subAccountId"] = request.SubAccountId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitBulletQuestions"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/submitBulletQuestions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitBulletQuestionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitBulletQuestions(request *SubmitBulletQuestionsRequest) (_result *SubmitBulletQuestionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitBulletQuestionsResponse{}
	_body, _err := client.SubmitBulletQuestionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitBulletQuestionsV1WithOptions(request *SubmitBulletQuestionsV1Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitBulletQuestionsV1Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitBulletQuestionsV1"),
		Version:     tea.String("2024-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/yic/yic-console/openService/v1/digitalHuman/commands/submitBulletQuestionsV1"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitBulletQuestionsV1Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitBulletQuestionsV1(request *SubmitBulletQuestionsV1Request) (_result *SubmitBulletQuestionsV1Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitBulletQuestionsV1Response{}
	_body, _err := client.SubmitBulletQuestionsV1WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
