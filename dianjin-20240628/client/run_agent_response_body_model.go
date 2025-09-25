// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *RunAgentResponseBody
	GetCost() *int64
	SetData(v *RunAgentResponseBodyData) *RunAgentResponseBody
	GetData() *RunAgentResponseBodyData
	SetDataType(v string) *RunAgentResponseBody
	GetDataType() *string
	SetErrCode(v string) *RunAgentResponseBody
	GetErrCode() *string
	SetMessage(v string) *RunAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunAgentResponseBody
	GetSuccess() *bool
	SetTime(v string) *RunAgentResponseBody
	GetTime() *string
}

type RunAgentResponseBody struct {
	// example:
	//
	// null
	Cost *int64                    `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *RunAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RunAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunAgentResponseBody) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *RunAgentResponseBody) GetData() *RunAgentResponseBodyData {
	return s.Data
}

func (s *RunAgentResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *RunAgentResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RunAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunAgentResponseBody) GetTime() *string {
	return s.Time
}

func (s *RunAgentResponseBody) SetCost(v int64) *RunAgentResponseBody {
	s.Cost = &v
	return s
}

func (s *RunAgentResponseBody) SetData(v *RunAgentResponseBodyData) *RunAgentResponseBody {
	s.Data = v
	return s
}

func (s *RunAgentResponseBody) SetDataType(v string) *RunAgentResponseBody {
	s.DataType = &v
	return s
}

func (s *RunAgentResponseBody) SetErrCode(v string) *RunAgentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RunAgentResponseBody) SetMessage(v string) *RunAgentResponseBody {
	s.Message = &v
	return s
}

func (s *RunAgentResponseBody) SetRequestId(v string) *RunAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunAgentResponseBody) SetSuccess(v bool) *RunAgentResponseBody {
	s.Success = &v
	return s
}

func (s *RunAgentResponseBody) SetTime(v string) *RunAgentResponseBody {
	s.Time = &v
	return s
}

func (s *RunAgentResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunAgentResponseBodyData struct {
	FunctionCallResponses []*RunAgentResponseBodyDataFunctionCallResponses `json:"functionCallResponses,omitempty" xml:"functionCallResponses,omitempty" type:"Repeated"`
	// example:
	//
	// 766
	InputTokens *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 988
	OutputTokens *int32                            `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	Response     *RunAgentResponseBodyDataResponse `json:"response,omitempty" xml:"response,omitempty" type:"Struct"`
	// example:
	//
	// 4vlag5ken3
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// example:
	//
	// 5bdb9809856c58acb92001f8ae65773c
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
	// example:
	//
	// w4paqoezm2
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s RunAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RunAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBodyData) GetFunctionCallResponses() []*RunAgentResponseBodyDataFunctionCallResponses {
	return s.FunctionCallResponses
}

func (s *RunAgentResponseBodyData) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *RunAgentResponseBodyData) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *RunAgentResponseBodyData) GetResponse() *RunAgentResponseBodyDataResponse {
	return s.Response
}

func (s *RunAgentResponseBodyData) GetThreadId() *string {
	return s.ThreadId
}

func (s *RunAgentResponseBodyData) GetTraceId() *string {
	return s.TraceId
}

func (s *RunAgentResponseBodyData) GetVersionId() *string {
	return s.VersionId
}

func (s *RunAgentResponseBodyData) SetFunctionCallResponses(v []*RunAgentResponseBodyDataFunctionCallResponses) *RunAgentResponseBodyData {
	s.FunctionCallResponses = v
	return s
}

func (s *RunAgentResponseBodyData) SetInputTokens(v int32) *RunAgentResponseBodyData {
	s.InputTokens = &v
	return s
}

func (s *RunAgentResponseBodyData) SetOutputTokens(v int32) *RunAgentResponseBodyData {
	s.OutputTokens = &v
	return s
}

func (s *RunAgentResponseBodyData) SetResponse(v *RunAgentResponseBodyDataResponse) *RunAgentResponseBodyData {
	s.Response = v
	return s
}

func (s *RunAgentResponseBodyData) SetThreadId(v string) *RunAgentResponseBodyData {
	s.ThreadId = &v
	return s
}

func (s *RunAgentResponseBodyData) SetTraceId(v string) *RunAgentResponseBodyData {
	s.TraceId = &v
	return s
}

func (s *RunAgentResponseBodyData) SetVersionId(v string) *RunAgentResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *RunAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RunAgentResponseBodyDataFunctionCallResponses struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 2025-01-21 16:37:14
	EndTime      *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FunctionArgs *string `json:"functionArgs,omitempty" xml:"functionArgs,omitempty"`
	// example:
	//
	// web_search
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	Result       *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 2025-01-21 16:37:14
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s RunAgentResponseBodyDataFunctionCallResponses) String() string {
	return dara.Prettify(s)
}

func (s RunAgentResponseBodyDataFunctionCallResponses) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) GetDisplayName() *string {
	return s.DisplayName
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) GetEndTime() *string {
	return s.EndTime
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) GetFunctionArgs() *string {
	return s.FunctionArgs
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) GetFunctionName() *string {
	return s.FunctionName
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) GetResult() *string {
	return s.Result
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) GetStartTime() *string {
	return s.StartTime
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetDisplayName(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.DisplayName = &v
	return s
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetEndTime(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.EndTime = &v
	return s
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetFunctionArgs(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.FunctionArgs = &v
	return s
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetFunctionName(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.FunctionName = &v
	return s
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetResult(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.Result = &v
	return s
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetStartTime(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.StartTime = &v
	return s
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) Validate() error {
	return dara.Validate(s)
}

type RunAgentResponseBodyDataResponse struct {
	Choices []*RunAgentResponseBodyDataResponseChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1737448637
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// d91d9afa-7cfc-4235-b012-a6f8e6ffa443
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 2025-01-21T16:37:17.497206762
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RunAgentResponseBodyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s RunAgentResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBodyDataResponse) GetChoices() []*RunAgentResponseBodyDataResponseChoices {
	return s.Choices
}

func (s *RunAgentResponseBodyDataResponse) GetCreated() *int64 {
	return s.Created
}

func (s *RunAgentResponseBodyDataResponse) GetId() *string {
	return s.Id
}

func (s *RunAgentResponseBodyDataResponse) GetModelId() *string {
	return s.ModelId
}

func (s *RunAgentResponseBodyDataResponse) GetTime() *string {
	return s.Time
}

func (s *RunAgentResponseBodyDataResponse) SetChoices(v []*RunAgentResponseBodyDataResponseChoices) *RunAgentResponseBodyDataResponse {
	s.Choices = v
	return s
}

func (s *RunAgentResponseBodyDataResponse) SetCreated(v int64) *RunAgentResponseBodyDataResponse {
	s.Created = &v
	return s
}

func (s *RunAgentResponseBodyDataResponse) SetId(v string) *RunAgentResponseBodyDataResponse {
	s.Id = &v
	return s
}

func (s *RunAgentResponseBodyDataResponse) SetModelId(v string) *RunAgentResponseBodyDataResponse {
	s.ModelId = &v
	return s
}

func (s *RunAgentResponseBodyDataResponse) SetTime(v string) *RunAgentResponseBodyDataResponse {
	s.Time = &v
	return s
}

func (s *RunAgentResponseBodyDataResponse) Validate() error {
	return dara.Validate(s)
}

type RunAgentResponseBodyDataResponseChoices struct {
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int32                                          `json:"index,omitempty" xml:"index,omitempty"`
	Message *RunAgentResponseBodyDataResponseChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s RunAgentResponseBodyDataResponseChoices) String() string {
	return dara.Prettify(s)
}

func (s RunAgentResponseBodyDataResponseChoices) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBodyDataResponseChoices) GetFinishReason() *string {
	return s.FinishReason
}

func (s *RunAgentResponseBodyDataResponseChoices) GetIndex() *int32 {
	return s.Index
}

func (s *RunAgentResponseBodyDataResponseChoices) GetMessage() *RunAgentResponseBodyDataResponseChoicesMessage {
	return s.Message
}

func (s *RunAgentResponseBodyDataResponseChoices) SetFinishReason(v string) *RunAgentResponseBodyDataResponseChoices {
	s.FinishReason = &v
	return s
}

func (s *RunAgentResponseBodyDataResponseChoices) SetIndex(v int32) *RunAgentResponseBodyDataResponseChoices {
	s.Index = &v
	return s
}

func (s *RunAgentResponseBodyDataResponseChoices) SetMessage(v *RunAgentResponseBodyDataResponseChoicesMessage) *RunAgentResponseBodyDataResponseChoices {
	s.Message = v
	return s
}

func (s *RunAgentResponseBodyDataResponseChoices) Validate() error {
	return dara.Validate(s)
}

type RunAgentResponseBodyDataResponseChoicesMessage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// assistant
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// assistant
	RoleDisplayName *string `json:"roleDisplayName,omitempty" xml:"roleDisplayName,omitempty"`
}

func (s RunAgentResponseBodyDataResponseChoicesMessage) String() string {
	return dara.Prettify(s)
}

func (s RunAgentResponseBodyDataResponseChoicesMessage) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBodyDataResponseChoicesMessage) GetContent() *string {
	return s.Content
}

func (s *RunAgentResponseBodyDataResponseChoicesMessage) GetRole() *string {
	return s.Role
}

func (s *RunAgentResponseBodyDataResponseChoicesMessage) GetRoleDisplayName() *string {
	return s.RoleDisplayName
}

func (s *RunAgentResponseBodyDataResponseChoicesMessage) SetContent(v string) *RunAgentResponseBodyDataResponseChoicesMessage {
	s.Content = &v
	return s
}

func (s *RunAgentResponseBodyDataResponseChoicesMessage) SetRole(v string) *RunAgentResponseBodyDataResponseChoicesMessage {
	s.Role = &v
	return s
}

func (s *RunAgentResponseBodyDataResponseChoicesMessage) SetRoleDisplayName(v string) *RunAgentResponseBodyDataResponseChoicesMessage {
	s.RoleDisplayName = &v
	return s
}

func (s *RunAgentResponseBodyDataResponseChoicesMessage) Validate() error {
	return dara.Validate(s)
}
