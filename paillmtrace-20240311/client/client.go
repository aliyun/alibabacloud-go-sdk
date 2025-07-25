// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EvaluationConfig struct {
	Answer  *EvaluationConfigAnswer  `json:"Answer,omitempty" xml:"Answer,omitempty" type:"Struct"`
	Context *EvaluationConfigContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	Query   *EvaluationConfigQuery   `json:"Query,omitempty" xml:"Query,omitempty" type:"Struct"`
}

func (s EvaluationConfig) String() string {
	return tea.Prettify(s)
}

func (s EvaluationConfig) GoString() string {
	return s.String()
}

func (s *EvaluationConfig) SetAnswer(v *EvaluationConfigAnswer) *EvaluationConfig {
	s.Answer = v
	return s
}

func (s *EvaluationConfig) SetContext(v *EvaluationConfigContext) *EvaluationConfig {
	s.Context = v
	return s
}

func (s *EvaluationConfig) SetQuery(v *EvaluationConfigQuery) *EvaluationConfig {
	s.Query = v
	return s
}

type EvaluationConfigAnswer struct {
	JsonPathInSpan      *string `json:"JsonPathInSpan,omitempty" xml:"JsonPathInSpan,omitempty"`
	JsonPathInSpanValue *string `json:"JsonPathInSpanValue,omitempty" xml:"JsonPathInSpanValue,omitempty"`
	SpanName            *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
}

func (s EvaluationConfigAnswer) String() string {
	return tea.Prettify(s)
}

func (s EvaluationConfigAnswer) GoString() string {
	return s.String()
}

func (s *EvaluationConfigAnswer) SetJsonPathInSpan(v string) *EvaluationConfigAnswer {
	s.JsonPathInSpan = &v
	return s
}

func (s *EvaluationConfigAnswer) SetJsonPathInSpanValue(v string) *EvaluationConfigAnswer {
	s.JsonPathInSpanValue = &v
	return s
}

func (s *EvaluationConfigAnswer) SetSpanName(v string) *EvaluationConfigAnswer {
	s.SpanName = &v
	return s
}

type EvaluationConfigContext struct {
	JsonPathInSpan      *string `json:"JsonPathInSpan,omitempty" xml:"JsonPathInSpan,omitempty"`
	JsonPathInSpanValue *string `json:"JsonPathInSpanValue,omitempty" xml:"JsonPathInSpanValue,omitempty"`
	SpanName            *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
}

func (s EvaluationConfigContext) String() string {
	return tea.Prettify(s)
}

func (s EvaluationConfigContext) GoString() string {
	return s.String()
}

func (s *EvaluationConfigContext) SetJsonPathInSpan(v string) *EvaluationConfigContext {
	s.JsonPathInSpan = &v
	return s
}

func (s *EvaluationConfigContext) SetJsonPathInSpanValue(v string) *EvaluationConfigContext {
	s.JsonPathInSpanValue = &v
	return s
}

func (s *EvaluationConfigContext) SetSpanName(v string) *EvaluationConfigContext {
	s.SpanName = &v
	return s
}

type EvaluationConfigQuery struct {
	JsonPathInSpan      *string `json:"JsonPathInSpan,omitempty" xml:"JsonPathInSpan,omitempty"`
	JsonPathInSpanValue *string `json:"JsonPathInSpanValue,omitempty" xml:"JsonPathInSpanValue,omitempty"`
	SpanName            *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
}

func (s EvaluationConfigQuery) String() string {
	return tea.Prettify(s)
}

func (s EvaluationConfigQuery) GoString() string {
	return s.String()
}

func (s *EvaluationConfigQuery) SetJsonPathInSpan(v string) *EvaluationConfigQuery {
	s.JsonPathInSpan = &v
	return s
}

func (s *EvaluationConfigQuery) SetJsonPathInSpanValue(v string) *EvaluationConfigQuery {
	s.JsonPathInSpanValue = &v
	return s
}

func (s *EvaluationConfigQuery) SetSpanName(v string) *EvaluationConfigQuery {
	s.SpanName = &v
	return s
}

type ModelConfig struct {
	ApiKey          *string  `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	Endpoint        *string  `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	IsSelfHost      *bool    `json:"IsSelfHost,omitempty" xml:"IsSelfHost,omitempty"`
	Name            *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Temperature     *float32 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	TopP            *float32 `json:"TopP,omitempty" xml:"TopP,omitempty"`
	UseFunctionCall *bool    `json:"UseFunctionCall,omitempty" xml:"UseFunctionCall,omitempty"`
}

func (s ModelConfig) String() string {
	return tea.Prettify(s)
}

func (s ModelConfig) GoString() string {
	return s.String()
}

func (s *ModelConfig) SetApiKey(v string) *ModelConfig {
	s.ApiKey = &v
	return s
}

func (s *ModelConfig) SetEndpoint(v string) *ModelConfig {
	s.Endpoint = &v
	return s
}

func (s *ModelConfig) SetIsSelfHost(v bool) *ModelConfig {
	s.IsSelfHost = &v
	return s
}

func (s *ModelConfig) SetName(v string) *ModelConfig {
	s.Name = &v
	return s
}

func (s *ModelConfig) SetTemperature(v float32) *ModelConfig {
	s.Temperature = &v
	return s
}

func (s *ModelConfig) SetTopP(v float32) *ModelConfig {
	s.TopP = &v
	return s
}

func (s *ModelConfig) SetUseFunctionCall(v bool) *ModelConfig {
	s.UseFunctionCall = &v
	return s
}

type CreateOnlineEvalTaskRequest struct {
	// The request data.
	Body *CreateOnlineEvalTaskRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s CreateOnlineEvalTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOnlineEvalTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateOnlineEvalTaskRequest) SetBody(v *CreateOnlineEvalTaskRequestBody) *CreateOnlineEvalTaskRequest {
	s.Body = v
	return s
}

type CreateOnlineEvalTaskRequestBody struct {
	// The name of the user application in the trace data.
	//
	// example:
	//
	// my-best-llm-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the task.
	//
	// example:
	//
	// April to June data assessment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end time of the trace data, in UTC format.
	//
	// example:
	//
	// 2025-06-05 14:00:01
	//
	// 2025-06-05
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This configuration structure defines the JSON paths needed to extract specific values from trace data in JSON format. EvaluationConfig defines these JSON paths.
	EvaluationConfig *EvaluationConfig `json:"EvaluationConfig,omitempty" xml:"EvaluationConfig,omitempty"`
	// The evaluation task must search for a certain amount of trace data generated by the user application as input data for the evaluation. This list defines the search filter conditions.
	Filters []*CreateOnlineEvalTaskRequestBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The access configuration structure of the model used in the evaluation.
	ModelConfig *ModelConfig `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// The evaluation task must search for a certain amount of trace data generated by the user application as input data for the evaluation. This is the width of the time window for each search of input data.
	//
	// example:
	//
	// 9
	SamplingFrequencyMinutes *int32 `json:"SamplingFrequencyMinutes,omitempty" xml:"SamplingFrequencyMinutes,omitempty"`
	// The percentage of data found in a time window that truly serves as evaluation input data. For example, 100 indicates that all data searched is used as evaluation input. 20 indicates that 20% of the found data is randomly selected as evaluation input.
	//
	// example:
	//
	// 50
	SamplingRatio *int32 `json:"SamplingRatio,omitempty" xml:"SamplingRatio,omitempty"`
	// The start time of the trace data, in UTC format.
	//
	// example:
	//
	// 2025-04-05 14:00:01
	//
	// 2025-04-05
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task name.
	//
	// example:
	//
	// my-llm-app-eval-task-1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateOnlineEvalTaskRequestBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOnlineEvalTaskRequestBody) GoString() string {
	return s.String()
}

func (s *CreateOnlineEvalTaskRequestBody) SetAppName(v string) *CreateOnlineEvalTaskRequestBody {
	s.AppName = &v
	return s
}

func (s *CreateOnlineEvalTaskRequestBody) SetDescription(v string) *CreateOnlineEvalTaskRequestBody {
	s.Description = &v
	return s
}

func (s *CreateOnlineEvalTaskRequestBody) SetEndTime(v string) *CreateOnlineEvalTaskRequestBody {
	s.EndTime = &v
	return s
}

func (s *CreateOnlineEvalTaskRequestBody) SetEvaluationConfig(v *EvaluationConfig) *CreateOnlineEvalTaskRequestBody {
	s.EvaluationConfig = v
	return s
}

func (s *CreateOnlineEvalTaskRequestBody) SetFilters(v []*CreateOnlineEvalTaskRequestBodyFilters) *CreateOnlineEvalTaskRequestBody {
	s.Filters = v
	return s
}

func (s *CreateOnlineEvalTaskRequestBody) SetModelConfig(v *ModelConfig) *CreateOnlineEvalTaskRequestBody {
	s.ModelConfig = v
	return s
}

func (s *CreateOnlineEvalTaskRequestBody) SetSamplingFrequencyMinutes(v int32) *CreateOnlineEvalTaskRequestBody {
	s.SamplingFrequencyMinutes = &v
	return s
}

func (s *CreateOnlineEvalTaskRequestBody) SetSamplingRatio(v int32) *CreateOnlineEvalTaskRequestBody {
	s.SamplingRatio = &v
	return s
}

func (s *CreateOnlineEvalTaskRequestBody) SetStartTime(v string) *CreateOnlineEvalTaskRequestBody {
	s.StartTime = &v
	return s
}

func (s *CreateOnlineEvalTaskRequestBody) SetTaskName(v string) *CreateOnlineEvalTaskRequestBody {
	s.TaskName = &v
	return s
}

type CreateOnlineEvalTaskRequestBodyFilters struct {
	// The key of the filter condition.
	//
	// Valid values:
	//
	// 	- Status
	//
	// 	- SpanName
	//
	// 	- Input
	//
	// 	- TraceType
	//
	// 	- SpanType
	//
	// 	- ServiceName
	//
	// 	- Output
	//
	// 	- TraceName
	//
	// 	- ServiceId
	//
	// example:
	//
	// ServiceId
	//
	// ServiceName
	//
	// Input
	//
	// Output
	//
	// Status
	//
	// TraceType
	//
	// SpanType
	//
	// TraceName
	//
	// SpanName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching operator of the filter condition.
	//
	// Valid values:
	//
	// 	- Contains
	//
	// 	- \\=
	//
	// 	- StartsWith
	//
	// example:
	//
	// =
	//
	// StartsWith
	//
	// Contains
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the filter condition.
	//
	// example:
	//
	// foo
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOnlineEvalTaskRequestBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s CreateOnlineEvalTaskRequestBodyFilters) GoString() string {
	return s.String()
}

func (s *CreateOnlineEvalTaskRequestBodyFilters) SetKey(v string) *CreateOnlineEvalTaskRequestBodyFilters {
	s.Key = &v
	return s
}

func (s *CreateOnlineEvalTaskRequestBodyFilters) SetOperator(v string) *CreateOnlineEvalTaskRequestBodyFilters {
	s.Operator = &v
	return s
}

func (s *CreateOnlineEvalTaskRequestBodyFilters) SetValue(v string) *CreateOnlineEvalTaskRequestBodyFilters {
	s.Value = &v
	return s
}

type CreateOnlineEvalTaskShrinkRequest struct {
	// The request data.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOnlineEvalTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOnlineEvalTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOnlineEvalTaskShrinkRequest) SetBodyShrink(v string) *CreateOnlineEvalTaskShrinkRequest {
	s.BodyShrink = &v
	return s
}

type CreateOnlineEvalTaskResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// EvaluationConfig.Answer.SpanName is required.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the created trace evaluation task.
	//
	// example:
	//
	// 711ef9112343286810abbfce04e161ee
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateOnlineEvalTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOnlineEvalTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOnlineEvalTaskResponseBody) SetCode(v string) *CreateOnlineEvalTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOnlineEvalTaskResponseBody) SetMessage(v string) *CreateOnlineEvalTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOnlineEvalTaskResponseBody) SetRequestId(v string) *CreateOnlineEvalTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOnlineEvalTaskResponseBody) SetTaskId(v string) *CreateOnlineEvalTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateOnlineEvalTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOnlineEvalTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOnlineEvalTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOnlineEvalTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOnlineEvalTaskResponse) SetHeaders(v map[string]*string) *CreateOnlineEvalTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOnlineEvalTaskResponse) SetStatusCode(v int32) *CreateOnlineEvalTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOnlineEvalTaskResponse) SetBody(v *CreateOnlineEvalTaskResponseBody) *CreateOnlineEvalTaskResponse {
	s.Body = v
	return s
}

type CreateServiceIdentityRoleResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// User don\\"t have permission to create SLR.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The role details.
	//
	// example:
	//
	// AliyunServiceRoleForPaiLLMTrace
	RoleDetails *string `json:"RoleDetails,omitempty" xml:"RoleDetails,omitempty"`
	// The name of the service-linked role. Default value: AliyunServiceRoleForPaiLLMTrace.
	//
	// example:
	//
	// AliyunServiceRoleForPaiLLMTrace
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateServiceIdentityRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceIdentityRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceIdentityRoleResponseBody) SetCode(v string) *CreateServiceIdentityRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetMessage(v string) *CreateServiceIdentityRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRequestId(v string) *CreateServiceIdentityRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRoleDetails(v string) *CreateServiceIdentityRoleResponseBody {
	s.RoleDetails = &v
	return s
}

func (s *CreateServiceIdentityRoleResponseBody) SetRoleName(v string) *CreateServiceIdentityRoleResponseBody {
	s.RoleName = &v
	return s
}

type CreateServiceIdentityRoleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceIdentityRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceIdentityRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceIdentityRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceIdentityRoleResponse) SetHeaders(v map[string]*string) *CreateServiceIdentityRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceIdentityRoleResponse) SetStatusCode(v int32) *CreateServiceIdentityRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceIdentityRoleResponse) SetBody(v *CreateServiceIdentityRoleResponseBody) *CreateServiceIdentityRoleResponse {
	s.Body = v
	return s
}

type DeleteOnlineEvalTaskResponseBody struct {
	// Internal error code. Set only when the response is in error.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response error message. Set only when the response is in error.
	//
	// example:
	//
	// task id is empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOnlineEvalTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOnlineEvalTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOnlineEvalTaskResponseBody) SetCode(v string) *DeleteOnlineEvalTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOnlineEvalTaskResponseBody) SetMessage(v string) *DeleteOnlineEvalTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOnlineEvalTaskResponseBody) SetRequestId(v string) *DeleteOnlineEvalTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteOnlineEvalTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOnlineEvalTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOnlineEvalTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOnlineEvalTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteOnlineEvalTaskResponse) SetHeaders(v map[string]*string) *DeleteOnlineEvalTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteOnlineEvalTaskResponse) SetStatusCode(v int32) *DeleteOnlineEvalTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOnlineEvalTaskResponse) SetBody(v *DeleteOnlineEvalTaskResponseBody) *DeleteOnlineEvalTaskResponse {
	s.Body = v
	return s
}

type EvaluateTraceRequest struct {
	// The name of the application to which the trace belongs.
	//
	// example:
	//
	// my-llm-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// If the value retrieved at the JSON path is itself a JSON string, further JSON path definitions within this JSON are necessary to get the actual value.
	//
	// This parameter is required.
	EvaluationConfig *EvaluationConfig `json:"EvaluationConfig,omitempty" xml:"EvaluationConfig,omitempty"`
	// The ID of the evaluation task. If not specified, the system randomly generates and returns an ID. You can use this ID to quickly search for evaluation results.
	//
	// example:
	//
	// 44aea0ee00000000be5be24b2abb8f98
	EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
	// The end time of the search time range, in UTC format.
	//
	// example:
	//
	// 2025-04-05 13:24:25
	//
	// 2025-04-05
	MaxTime *string `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	// The start time of the search time range, in UTC format.
	//
	// example:
	//
	// 2025-04-05 13:24:25
	//
	// 2025-04-05
	MinTime *string `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	// The configuration structure to access the model used internally by the evaluation trace.
	ModelConfig *ModelConfig `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
}

func (s EvaluateTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s EvaluateTraceRequest) GoString() string {
	return s.String()
}

func (s *EvaluateTraceRequest) SetAppName(v string) *EvaluateTraceRequest {
	s.AppName = &v
	return s
}

func (s *EvaluateTraceRequest) SetEvaluationConfig(v *EvaluationConfig) *EvaluateTraceRequest {
	s.EvaluationConfig = v
	return s
}

func (s *EvaluateTraceRequest) SetEvaluationId(v string) *EvaluateTraceRequest {
	s.EvaluationId = &v
	return s
}

func (s *EvaluateTraceRequest) SetMaxTime(v string) *EvaluateTraceRequest {
	s.MaxTime = &v
	return s
}

func (s *EvaluateTraceRequest) SetMinTime(v string) *EvaluateTraceRequest {
	s.MinTime = &v
	return s
}

func (s *EvaluateTraceRequest) SetModelConfig(v *ModelConfig) *EvaluateTraceRequest {
	s.ModelConfig = v
	return s
}

type EvaluateTraceResponseBody struct {
	// The internal error code. This parameter is returned if an exception occurred.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// the task ID of the evaluation task to which the trace belongs.
	//
	// example:
	//
	// 6000043e103011f0922edec44617e03c
	EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
	// The error message. This parameter is returned if an exception occurred.
	//
	// example:
	//
	// eval_request missing dataset id or times
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F1AB295E-0D1F-5ECE-9FFA-98ABB4CB5DF5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EvaluateTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EvaluateTraceResponseBody) GoString() string {
	return s.String()
}

func (s *EvaluateTraceResponseBody) SetCode(v string) *EvaluateTraceResponseBody {
	s.Code = &v
	return s
}

func (s *EvaluateTraceResponseBody) SetEvaluationId(v string) *EvaluateTraceResponseBody {
	s.EvaluationId = &v
	return s
}

func (s *EvaluateTraceResponseBody) SetMessage(v string) *EvaluateTraceResponseBody {
	s.Message = &v
	return s
}

func (s *EvaluateTraceResponseBody) SetRequestId(v string) *EvaluateTraceResponseBody {
	s.RequestId = &v
	return s
}

type EvaluateTraceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EvaluateTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvaluateTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s EvaluateTraceResponse) GoString() string {
	return s.String()
}

func (s *EvaluateTraceResponse) SetHeaders(v map[string]*string) *EvaluateTraceResponse {
	s.Headers = v
	return s
}

func (s *EvaluateTraceResponse) SetStatusCode(v int32) *EvaluateTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *EvaluateTraceResponse) SetBody(v *EvaluateTraceResponseBody) *EvaluateTraceResponse {
	s.Body = v
	return s
}

type GetEvaluationTemplatesResponseBody struct {
	// Internal error code. Set only when the response has an error.
	//
	// example:
	//
	// ExecutionFailure
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A series of templates used internally by the evaluation system to construct LLM interaction information.
	EvaluationTemplates []interface{} `json:"EvaluationTemplates,omitempty" xml:"EvaluationTemplates,omitempty" type:"Repeated"`
	// Response error message. Set only when the response has an error.
	//
	// example:
	//
	// cannot get data back.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEvaluationTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEvaluationTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *GetEvaluationTemplatesResponseBody) SetCode(v string) *GetEvaluationTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *GetEvaluationTemplatesResponseBody) SetEvaluationTemplates(v []interface{}) *GetEvaluationTemplatesResponseBody {
	s.EvaluationTemplates = v
	return s
}

func (s *GetEvaluationTemplatesResponseBody) SetMessage(v string) *GetEvaluationTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *GetEvaluationTemplatesResponseBody) SetRequestId(v string) *GetEvaluationTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type GetEvaluationTemplatesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEvaluationTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEvaluationTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEvaluationTemplatesResponse) GoString() string {
	return s.String()
}

func (s *GetEvaluationTemplatesResponse) SetHeaders(v map[string]*string) *GetEvaluationTemplatesResponse {
	s.Headers = v
	return s
}

func (s *GetEvaluationTemplatesResponse) SetStatusCode(v int32) *GetEvaluationTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEvaluationTemplatesResponse) SetBody(v *GetEvaluationTemplatesResponseBody) *GetEvaluationTemplatesResponse {
	s.Body = v
	return s
}

type GetOnlineEvalTaskResponseBody struct {
	// Internal error code. Set only when the response is in error.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response error message. Set only when the response is in error.
	//
	// example:
	//
	// task id is empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// POP request ID
	//
	// example:
	//
	// 6F352A02-9C0D-54A7-B57C-663CF71D5714
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task information
	Task *GetOnlineEvalTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s GetOnlineEvalTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineEvalTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetOnlineEvalTaskResponseBody) SetCode(v string) *GetOnlineEvalTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBody) SetMessage(v string) *GetOnlineEvalTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBody) SetRequestId(v string) *GetOnlineEvalTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBody) SetTask(v *GetOnlineEvalTaskResponseBodyTask) *GetOnlineEvalTaskResponseBody {
	s.Task = v
	return s
}

type GetOnlineEvalTaskResponseBodyTask struct {
	// The Alibaba Cloud account (primary account) of the task creator.
	//
	// example:
	//
	// 1195531608511111
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The name of the user application targeted by this task.
	//
	// example:
	//
	// my-llm-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Task description information
	//
	// example:
	//
	// 这个任务做了这些事。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Deprecated. Will be removed.
	//
	// example:
	//
	// Deprecated. Will be removed.
	EvalResults *string `json:"EvalResults,omitempty" xml:"EvalResults,omitempty"`
	// Extract specific path values from JSON-formatted trace data as input for the evaluation operation. These JSON paths are defined within this EvaluationConfig structure.
	EvaluationConfig *EvaluationConfig `json:"EvaluationConfig,omitempty" xml:"EvaluationConfig,omitempty"`
	// The evaluation task needs to search for a certain amount of trace data generated by the user application as input data for the evaluation operation. This is a list that defines the search filter conditions.
	Filters []*GetOnlineEvalTaskResponseBodyTaskFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// UTC creation time of the task.
	//
	// example:
	//
	// 2024-07-31 08:30:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// UTC end time of the trace data.
	//
	// example:
	//
	// 2024-08-10 13:20:00
	GmtEndTime *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	// UTC upper bound of the last sampling window
	//
	// example:
	//
	// 2024-08-10 13:14:00
	GmtLastSamplingWindowEndTime *string `json:"GmtLastSamplingWindowEndTime,omitempty" xml:"GmtLastSamplingWindowEndTime,omitempty"`
	// UTC lower bound of the last sampling window.
	//
	// example:
	//
	// 2024-08-10 13:11:00
	GmtLastSamplingWindowStartTime *string `json:"GmtLastSamplingWindowStartTime,omitempty" xml:"GmtLastSamplingWindowStartTime,omitempty"`
	// UTC start time of the trace data.
	//
	// example:
	//
	// 2024-08-02
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// Task ID
	//
	// example:
	//
	// 0839a02d-aa24-4174-90bb-7a773885934d
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Access configuration structure for the large model used internally by the evaluation task.
	ModelConfig *ModelConfig `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// Task name.
	//
	// example:
	//
	// my-eval-task-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Number of evaluation records
	//
	// example:
	//
	// 999
	RecordCount *int32 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The evaluation task needs to search for a certain amount of trace data generated by the user application as input data for the evaluation operation. This defines the width of the time window for each search of input data.
	//
	// example:
	//
	// 3
	SamplingFrequencyMinutes *int32 `json:"SamplingFrequencyMinutes,omitempty" xml:"SamplingFrequencyMinutes,omitempty"`
	// The percentage of the data found within a time window that is actually used as input for the evaluation task. For example, 100 means all the found data is used as input, 20 means 20% of the found data is randomly selected as input.
	//
	// example:
	//
	// 70
	SamplingRatio *int32 `json:"SamplingRatio,omitempty" xml:"SamplingRatio,omitempty"`
	// Task status
	//
	// example:
	//
	// CREATED
	//
	// RUNNING
	//
	// FINISHED
	//
	// USER_CANCELED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Alibaba Cloud sub-account of the task creator.
	//
	// example:
	//
	// 222222222222222222
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetOnlineEvalTaskResponseBodyTask) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineEvalTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetAliyunUid(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.AliyunUid = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetAppName(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.AppName = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetDescription(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.Description = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetEvalResults(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.EvalResults = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetEvaluationConfig(v *EvaluationConfig) *GetOnlineEvalTaskResponseBodyTask {
	s.EvaluationConfig = v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetFilters(v []*GetOnlineEvalTaskResponseBodyTaskFilters) *GetOnlineEvalTaskResponseBodyTask {
	s.Filters = v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetGmtCreateTime(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.GmtCreateTime = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetGmtEndTime(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.GmtEndTime = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetGmtLastSamplingWindowEndTime(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.GmtLastSamplingWindowEndTime = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetGmtLastSamplingWindowStartTime(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.GmtLastSamplingWindowStartTime = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetGmtStartTime(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.GmtStartTime = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetId(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.Id = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetModelConfig(v *ModelConfig) *GetOnlineEvalTaskResponseBodyTask {
	s.ModelConfig = v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetName(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.Name = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetRecordCount(v int32) *GetOnlineEvalTaskResponseBodyTask {
	s.RecordCount = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetSamplingFrequencyMinutes(v int32) *GetOnlineEvalTaskResponseBodyTask {
	s.SamplingFrequencyMinutes = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetSamplingRatio(v int32) *GetOnlineEvalTaskResponseBodyTask {
	s.SamplingRatio = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetStatus(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTask) SetUserId(v string) *GetOnlineEvalTaskResponseBodyTask {
	s.UserId = &v
	return s
}

type GetOnlineEvalTaskResponseBodyTaskFilters struct {
	// Key of the filter condition.
	//
	// example:
	//
	// ServiceId
	//
	// ServiceName
	//
	// Input
	//
	// Output
	//
	// Status
	//
	// TraceType
	//
	// SpanType
	//
	// TraceName
	//
	// SpanName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Filter condition match operator.
	//
	// example:
	//
	// =
	//
	// StartsWith
	//
	// Contains
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Value of the filter condition.
	//
	// example:
	//
	// foo
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetOnlineEvalTaskResponseBodyTaskFilters) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineEvalTaskResponseBodyTaskFilters) GoString() string {
	return s.String()
}

func (s *GetOnlineEvalTaskResponseBodyTaskFilters) SetKey(v string) *GetOnlineEvalTaskResponseBodyTaskFilters {
	s.Key = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTaskFilters) SetOperator(v string) *GetOnlineEvalTaskResponseBodyTaskFilters {
	s.Operator = &v
	return s
}

func (s *GetOnlineEvalTaskResponseBodyTaskFilters) SetValue(v string) *GetOnlineEvalTaskResponseBodyTaskFilters {
	s.Value = &v
	return s
}

type GetOnlineEvalTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOnlineEvalTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOnlineEvalTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOnlineEvalTaskResponse) GoString() string {
	return s.String()
}

func (s *GetOnlineEvalTaskResponse) SetHeaders(v map[string]*string) *GetOnlineEvalTaskResponse {
	s.Headers = v
	return s
}

func (s *GetOnlineEvalTaskResponse) SetStatusCode(v int32) *GetOnlineEvalTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOnlineEvalTaskResponse) SetBody(v *GetOnlineEvalTaskResponseBody) *GetOnlineEvalTaskResponse {
	s.Body = v
	return s
}

type GetServiceIdentityRoleResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// EntityNotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// Serivce role does not exit.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The role details.
	//
	// example:
	//
	// AliyunServiceRoleForPaiLLMTrace
	RoleDetail *string `json:"RoleDetail,omitempty" xml:"RoleDetail,omitempty"`
	// The name of the service-linked role. Default value: AliyunServiceRoleForPaiLLMTrace.
	//
	// example:
	//
	// AliyunServiceRoleForPaiLLMTrace
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetServiceIdentityRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceIdentityRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceIdentityRoleResponseBody) SetCode(v string) *GetServiceIdentityRoleResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetMessage(v string) *GetServiceIdentityRoleResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRequestId(v string) *GetServiceIdentityRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRoleDetail(v string) *GetServiceIdentityRoleResponseBody {
	s.RoleDetail = &v
	return s
}

func (s *GetServiceIdentityRoleResponseBody) SetRoleName(v string) *GetServiceIdentityRoleResponseBody {
	s.RoleName = &v
	return s
}

type GetServiceIdentityRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceIdentityRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceIdentityRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceIdentityRoleResponse) GoString() string {
	return s.String()
}

func (s *GetServiceIdentityRoleResponse) SetHeaders(v map[string]*string) *GetServiceIdentityRoleResponse {
	s.Headers = v
	return s
}

func (s *GetServiceIdentityRoleResponse) SetStatusCode(v int32) *GetServiceIdentityRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceIdentityRoleResponse) SetBody(v *GetServiceIdentityRoleResponseBody) *GetServiceIdentityRoleResponse {
	s.Body = v
	return s
}

type GetXtraceTokenResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The gRPC endpoint used for uploading ARM traces.
	//
	// example:
	//
	// http://tracing-analysis-dc-hz.aliyuncs.com:8090
	GrpcEndpoint *string `json:"GrpcEndpoint,omitempty" xml:"GrpcEndpoint,omitempty"`
	// The internal gRPC endpoint used for uploading ARMS traces used by Alibaba Cloud.
	//
	// example:
	//
	// http://tracing-analysis-dc-hz-internal.aliyuncs.com:8090
	GrpcInternalEndpoint *string `json:"GrpcInternalEndpoint,omitempty" xml:"GrpcInternalEndpoint,omitempty"`
	// The endpoint used for uploading ARMS traces.
	//
	// example:
	//
	// http://tracing-analysis-dc-hz.aliyuncs.com/aaa@bbb@ccc/api/otlp/traces
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// The internal endpoint used for uploading ARMS traces used by Alibaba Cloud.
	//
	// example:
	//
	// http://tracing-analysis-dc-hz-internal.aliyuncs.com/aaa@bbb@ccc/api/otlp/traces
	HttpInternalEndpoint *string `json:"HttpInternalEndpoint,omitempty" xml:"HttpInternalEndpoint,omitempty"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// get_xtrace_token: failed, ERROR: NoPermission
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The token used for uploading ARMS traces.
	//
	// example:
	//
	// h1abcw7@5abcb_h1abcw7@5abc01
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetXtraceTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetXtraceTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetXtraceTokenResponseBody) SetCode(v string) *GetXtraceTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetGrpcEndpoint(v string) *GetXtraceTokenResponseBody {
	s.GrpcEndpoint = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetGrpcInternalEndpoint(v string) *GetXtraceTokenResponseBody {
	s.GrpcInternalEndpoint = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetHttpEndpoint(v string) *GetXtraceTokenResponseBody {
	s.HttpEndpoint = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetHttpInternalEndpoint(v string) *GetXtraceTokenResponseBody {
	s.HttpInternalEndpoint = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetMessage(v string) *GetXtraceTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetRequestId(v string) *GetXtraceTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetXtraceTokenResponseBody) SetToken(v string) *GetXtraceTokenResponseBody {
	s.Token = &v
	return s
}

type GetXtraceTokenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetXtraceTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetXtraceTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetXtraceTokenResponse) GoString() string {
	return s.String()
}

func (s *GetXtraceTokenResponse) SetHeaders(v map[string]*string) *GetXtraceTokenResponse {
	s.Headers = v
	return s
}

func (s *GetXtraceTokenResponse) SetStatusCode(v int32) *GetXtraceTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetXtraceTokenResponse) SetBody(v *GetXtraceTokenResponseBody) *GetXtraceTokenResponse {
	s.Body = v
	return s
}

type ListEvalResultsRequest struct {
	// The task ID of the evaluation task to which the trace belongs.
	//
	// example:
	//
	// 0bb05ae2a2dc11ef9757faaa2a1ec0c6
	EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
	// The keyword to query from the evaluation inputs.
	//
	// example:
	//
	// foo
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Page starts from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The trace data IDs.
	RecordIds []*string `json:"RecordIds,omitempty" xml:"RecordIds,omitempty" type:"Repeated"`
}

func (s ListEvalResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEvalResultsRequest) GoString() string {
	return s.String()
}

func (s *ListEvalResultsRequest) SetEvaluationId(v string) *ListEvalResultsRequest {
	s.EvaluationId = &v
	return s
}

func (s *ListEvalResultsRequest) SetKeyword(v string) *ListEvalResultsRequest {
	s.Keyword = &v
	return s
}

func (s *ListEvalResultsRequest) SetPageNumber(v int32) *ListEvalResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEvalResultsRequest) SetPageSize(v int32) *ListEvalResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEvalResultsRequest) SetRecordIds(v []*string) *ListEvalResultsRequest {
	s.RecordIds = v
	return s
}

type ListEvalResultsShrinkRequest struct {
	// The task ID of the evaluation task to which the trace belongs.
	//
	// example:
	//
	// 0bb05ae2a2dc11ef9757faaa2a1ec0c6
	EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
	// The keyword to query from the evaluation inputs.
	//
	// example:
	//
	// foo
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Page starts from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The trace data IDs.
	RecordIdsShrink *string `json:"RecordIds,omitempty" xml:"RecordIds,omitempty"`
}

func (s ListEvalResultsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEvalResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListEvalResultsShrinkRequest) SetEvaluationId(v string) *ListEvalResultsShrinkRequest {
	s.EvaluationId = &v
	return s
}

func (s *ListEvalResultsShrinkRequest) SetKeyword(v string) *ListEvalResultsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListEvalResultsShrinkRequest) SetPageNumber(v int32) *ListEvalResultsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEvalResultsShrinkRequest) SetPageSize(v int32) *ListEvalResultsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListEvalResultsShrinkRequest) SetRecordIdsShrink(v string) *ListEvalResultsShrinkRequest {
	s.RecordIdsShrink = &v
	return s
}

type ListEvalResultsResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// ExecutionFailure
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The evaluation results.
	EvaluationResults []*string `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Repeated"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// cannot get data back.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the POP request.
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of results that meet the condition.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEvalResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEvalResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvalResultsResponseBody) SetCode(v string) *ListEvalResultsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEvalResultsResponseBody) SetEvaluationResults(v []*string) *ListEvalResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListEvalResultsResponseBody) SetMessage(v string) *ListEvalResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEvalResultsResponseBody) SetRequestId(v string) *ListEvalResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvalResultsResponseBody) SetTotalCount(v int32) *ListEvalResultsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEvalResultsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvalResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvalResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEvalResultsResponse) GoString() string {
	return s.String()
}

func (s *ListEvalResultsResponse) SetHeaders(v map[string]*string) *ListEvalResultsResponse {
	s.Headers = v
	return s
}

func (s *ListEvalResultsResponse) SetStatusCode(v int32) *ListEvalResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvalResultsResponse) SetBody(v *ListEvalResultsResponseBody) *ListEvalResultsResponse {
	s.Body = v
	return s
}

type ListOnlineEvalTaskResultsRequest struct {
	// The ID of the evaluation task. At least one of the trace ID or task ID must be set.
	//
	// example:
	//
	// 0bb05ae8888c11ef9757faaa2a1ec0c6
	EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
	// The same trace data may have been evaluated by different tasks. If no task ID is specified and there are multiple evaluation results for the same trace ID, this parameter specifies whether to return only the most recent evaluation result.
	//
	// example:
	//
	// True
	MostRecentResultsOnly *bool `json:"MostRecentResultsOnly,omitempty" xml:"MostRecentResultsOnly,omitempty"`
	// The current page number. Value range: integers greater than 0. Default value: 1.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, default is 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specify a set of trace IDs, and only return the evaluation results for these traces. At least one of the trace ID or task ID must be set.
	TraceIds []*string `json:"TraceIds,omitempty" xml:"TraceIds,omitempty" type:"Repeated"`
}

func (s ListOnlineEvalTaskResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineEvalTaskResultsRequest) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTaskResultsRequest) SetEvaluationId(v string) *ListOnlineEvalTaskResultsRequest {
	s.EvaluationId = &v
	return s
}

func (s *ListOnlineEvalTaskResultsRequest) SetMostRecentResultsOnly(v bool) *ListOnlineEvalTaskResultsRequest {
	s.MostRecentResultsOnly = &v
	return s
}

func (s *ListOnlineEvalTaskResultsRequest) SetPageNumber(v int32) *ListOnlineEvalTaskResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOnlineEvalTaskResultsRequest) SetPageSize(v int32) *ListOnlineEvalTaskResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOnlineEvalTaskResultsRequest) SetTraceIds(v []*string) *ListOnlineEvalTaskResultsRequest {
	s.TraceIds = v
	return s
}

type ListOnlineEvalTaskResultsShrinkRequest struct {
	// The ID of the evaluation task. At least one of the trace ID or task ID must be set.
	//
	// example:
	//
	// 0bb05ae8888c11ef9757faaa2a1ec0c6
	EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
	// The same trace data may have been evaluated by different tasks. If no task ID is specified and there are multiple evaluation results for the same trace ID, this parameter specifies whether to return only the most recent evaluation result.
	//
	// example:
	//
	// True
	MostRecentResultsOnly *bool `json:"MostRecentResultsOnly,omitempty" xml:"MostRecentResultsOnly,omitempty"`
	// The current page number. Value range: integers greater than 0. Default value: 1.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, default is 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specify a set of trace IDs, and only return the evaluation results for these traces. At least one of the trace ID or task ID must be set.
	TraceIdsShrink *string `json:"TraceIds,omitempty" xml:"TraceIds,omitempty"`
}

func (s ListOnlineEvalTaskResultsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineEvalTaskResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) SetEvaluationId(v string) *ListOnlineEvalTaskResultsShrinkRequest {
	s.EvaluationId = &v
	return s
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) SetMostRecentResultsOnly(v bool) *ListOnlineEvalTaskResultsShrinkRequest {
	s.MostRecentResultsOnly = &v
	return s
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) SetPageNumber(v int32) *ListOnlineEvalTaskResultsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) SetPageSize(v int32) *ListOnlineEvalTaskResultsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListOnlineEvalTaskResultsShrinkRequest) SetTraceIdsShrink(v string) *ListOnlineEvalTaskResultsShrinkRequest {
	s.TraceIdsShrink = &v
	return s
}

type ListOnlineEvalTaskResultsResponseBody struct {
	// Internal error code. Set only when the response has an error.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// List of evaluation results.
	EvaluationResults []*string `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Repeated"`
	// Response error message. Set only when the response has an error.
	//
	// example:
	//
	// must provide trace_id(s) or eval_id
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 22BA9A5A-E3D8-5B4C-90FC-F33F6E5853F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of evaluation results that meet the criteria.
	//
	// example:
	//
	// 123
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOnlineEvalTaskResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineEvalTaskResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTaskResultsResponseBody) SetCode(v string) *ListOnlineEvalTaskResultsResponseBody {
	s.Code = &v
	return s
}

func (s *ListOnlineEvalTaskResultsResponseBody) SetEvaluationResults(v []*string) *ListOnlineEvalTaskResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListOnlineEvalTaskResultsResponseBody) SetMessage(v string) *ListOnlineEvalTaskResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListOnlineEvalTaskResultsResponseBody) SetRequestId(v string) *ListOnlineEvalTaskResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOnlineEvalTaskResultsResponseBody) SetTotalCount(v int32) *ListOnlineEvalTaskResultsResponseBody {
	s.TotalCount = &v
	return s
}

type ListOnlineEvalTaskResultsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOnlineEvalTaskResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOnlineEvalTaskResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineEvalTaskResultsResponse) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTaskResultsResponse) SetHeaders(v map[string]*string) *ListOnlineEvalTaskResultsResponse {
	s.Headers = v
	return s
}

func (s *ListOnlineEvalTaskResultsResponse) SetStatusCode(v int32) *ListOnlineEvalTaskResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOnlineEvalTaskResultsResponse) SetBody(v *ListOnlineEvalTaskResultsResponseBody) *ListOnlineEvalTaskResultsResponse {
	s.Body = v
	return s
}

type ListOnlineEvalTasksRequest struct {
	// Search keyword. It will match on fields such as task name, application name (appName), task description, and evaluation metric name.
	//
	// example:
	//
	// foo
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The UTC end time of the search time range
	//
	// example:
	//
	// 2025-04-07 13:24:25
	//
	// 2025-04-07
	MaxTime *string `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	// The UTC start time of the search time range
	//
	// example:
	//
	// 2025-04-05 13:24:25
	//
	// 2025-04-05
	MinTime *string `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	// The current page number. Value range: integers greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, default is 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOnlineEvalTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineEvalTasksRequest) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTasksRequest) SetKeyword(v string) *ListOnlineEvalTasksRequest {
	s.Keyword = &v
	return s
}

func (s *ListOnlineEvalTasksRequest) SetMaxTime(v string) *ListOnlineEvalTasksRequest {
	s.MaxTime = &v
	return s
}

func (s *ListOnlineEvalTasksRequest) SetMinTime(v string) *ListOnlineEvalTasksRequest {
	s.MinTime = &v
	return s
}

func (s *ListOnlineEvalTasksRequest) SetPageNumber(v int32) *ListOnlineEvalTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOnlineEvalTasksRequest) SetPageSize(v int32) *ListOnlineEvalTasksRequest {
	s.PageSize = &v
	return s
}

type ListOnlineEvalTasksResponseBody struct {
	// Internal error code. Set only when the response has an error.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response error message. Set only when the response has an error.
	//
	// example:
	//
	// page number should be greater than 0
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 6F352A02-9C0D-54A7-B57C-663CF71D5714
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of tasks.
	Tasks []*ListOnlineEvalTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// Total number of tasks that meet the criteria.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOnlineEvalTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineEvalTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTasksResponseBody) SetCode(v string) *ListOnlineEvalTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBody) SetMessage(v string) *ListOnlineEvalTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBody) SetRequestId(v string) *ListOnlineEvalTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBody) SetTasks(v []*ListOnlineEvalTasksResponseBodyTasks) *ListOnlineEvalTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListOnlineEvalTasksResponseBody) SetTotalCount(v int32) *ListOnlineEvalTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListOnlineEvalTasksResponseBodyTasks struct {
	// The Alibaba Cloud account (primary account) of the task creator.
	//
	// example:
	//
	// 1512522691911111
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The name of the user application targeted by this task.
	//
	// example:
	//
	// my-llm-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Task description information
	//
	// example:
	//
	// this task is for application evaluation
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Extract specific path values from JSON-formatted trace data as input for the evaluation operation. These JSON paths are defined in this EvaluationConfig structure.
	EvaluationConfig *EvaluationConfig `json:"EvaluationConfig,omitempty" xml:"EvaluationConfig,omitempty"`
	// The list define the search filter conditions for the evaluation task to search a certain amount of trace data generated by the user application, which serves as input data for the evaluation operation.
	Filters []*ListOnlineEvalTasksResponseBodyTasksFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The UTC creation time of the task.
	//
	// example:
	//
	// 2025-04-07 13:24:35
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Task UTC end time.
	//
	// example:
	//
	// 2025-04-09 13:24:35
	GmtEndTime *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	// The UTC start time of the task.
	//
	// example:
	//
	// 2025-04-08 13:24:35
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 9f50cd72efcf36535152ee811a911115
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Access configuration structure for the large model used internally by the evaluation task.
	ModelConfig *ModelConfig `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// Task name.
	//
	// example:
	//
	// my-foo-evaluation-task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The evaluation task needs to search for a certain amount of trace data generated by the user\\"s application as input data for the evaluation operation. This defines the time window for each data search.
	//
	// example:
	//
	// 12
	SamplingFrequencyMinutes *int32 `json:"SamplingFrequencyMinutes,omitempty" xml:"SamplingFrequencyMinutes,omitempty"`
	// The percentage of the data searched within a time window that is used as input data for the evaluation. For example, 100 means all the searched data is used as input, 20 means 20% of the searched data is randomly selected as input.
	//
	// example:
	//
	// 50
	SamplingRatio *int32 `json:"SamplingRatio,omitempty" xml:"SamplingRatio,omitempty"`
	// Task status
	//
	// example:
	//
	// CREATED
	//
	// RUNNING
	//
	// FINISHED
	//
	// USER_CANCELED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Alibaba Cloud sub-account of the task creator.
	//
	// example:
	//
	// 2222222222
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListOnlineEvalTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineEvalTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetAliyunUid(v string) *ListOnlineEvalTasksResponseBodyTasks {
	s.AliyunUid = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetAppName(v string) *ListOnlineEvalTasksResponseBodyTasks {
	s.AppName = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetDescription(v string) *ListOnlineEvalTasksResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetEvaluationConfig(v *EvaluationConfig) *ListOnlineEvalTasksResponseBodyTasks {
	s.EvaluationConfig = v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetFilters(v []*ListOnlineEvalTasksResponseBodyTasksFilters) *ListOnlineEvalTasksResponseBodyTasks {
	s.Filters = v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetGmtCreateTime(v string) *ListOnlineEvalTasksResponseBodyTasks {
	s.GmtCreateTime = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetGmtEndTime(v string) *ListOnlineEvalTasksResponseBodyTasks {
	s.GmtEndTime = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetGmtStartTime(v string) *ListOnlineEvalTasksResponseBodyTasks {
	s.GmtStartTime = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetId(v string) *ListOnlineEvalTasksResponseBodyTasks {
	s.Id = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetModelConfig(v *ModelConfig) *ListOnlineEvalTasksResponseBodyTasks {
	s.ModelConfig = v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetName(v string) *ListOnlineEvalTasksResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetSamplingFrequencyMinutes(v int32) *ListOnlineEvalTasksResponseBodyTasks {
	s.SamplingFrequencyMinutes = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetSamplingRatio(v int32) *ListOnlineEvalTasksResponseBodyTasks {
	s.SamplingRatio = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetStatus(v string) *ListOnlineEvalTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasks) SetUserId(v string) *ListOnlineEvalTasksResponseBodyTasks {
	s.UserId = &v
	return s
}

type ListOnlineEvalTasksResponseBodyTasksFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// ServiceId
	//
	// ServiceName
	//
	// Input
	//
	// Output
	//
	// Status
	//
	// TraceType
	//
	// SpanType
	//
	// TraceName
	//
	// SpanName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching operator of the filter condition.
	//
	// example:
	//
	// =
	//
	// StartsWith
	//
	// Contains
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the filter condition.
	//
	// example:
	//
	// foo
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListOnlineEvalTasksResponseBodyTasksFilters) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineEvalTasksResponseBodyTasksFilters) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTasksResponseBodyTasksFilters) SetKey(v string) *ListOnlineEvalTasksResponseBodyTasksFilters {
	s.Key = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasksFilters) SetOperator(v string) *ListOnlineEvalTasksResponseBodyTasksFilters {
	s.Operator = &v
	return s
}

func (s *ListOnlineEvalTasksResponseBodyTasksFilters) SetValue(v string) *ListOnlineEvalTasksResponseBodyTasksFilters {
	s.Value = &v
	return s
}

type ListOnlineEvalTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOnlineEvalTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOnlineEvalTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOnlineEvalTasksResponse) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTasksResponse) SetHeaders(v map[string]*string) *ListOnlineEvalTasksResponse {
	s.Headers = v
	return s
}

func (s *ListOnlineEvalTasksResponse) SetStatusCode(v int32) *ListOnlineEvalTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOnlineEvalTasksResponse) SetBody(v *ListOnlineEvalTasksResponseBody) *ListOnlineEvalTasksResponse {
	s.Body = v
	return s
}

type ListTracesDatasRequest struct {
	// The value of the attributes.service.app.user_id field in the trace record. It can contain upper and lower case letters, digits, dot (.), hyphen (-), and underscore (_). It is empty by default.
	//
	// example:
	//
	// end-user.12345
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Other filter parameters
	Filters []*ListTracesDatasRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// Whether to return only trace records containing spans with a non-empty events. Example: Suppose a trace has 3 spans. If this parameter is True, this trace meets the condition when any one of the 3 spans has a non-empty events. The default value is False. The events is not used for filtering.
	//
	// example:
	//
	// False
	HasEvents *bool `json:"HasEvents,omitempty" xml:"HasEvents,omitempty"`
	// Whether to return only trace records containing spans with a non-empty statusMessage. Example: Suppose a trace has 3 spans. If this parameter is True, this trace meets the condition when any one of the 3 spans has a non-empty statusMessage. The default value is False. The statusMessage is not used for filtering.
	//
	// example:
	//
	// False
	HasStatusMessage *bool `json:"HasStatusMessage,omitempty" xml:"HasStatusMessage,omitempty"`
	// The value of the resources.service.app.name field in the trace record. It can contain upper and lower case letters, digits, dot (.), hyphen (-), and underscore (_). Must be an exact match. It is empty by default.
	//
	// example:
	//
	// My.super_LLM-app2
	LlmAppName  *string  `json:"LlmAppName,omitempty" xml:"LlmAppName,omitempty"`
	MaxDuration *float32 `json:"MaxDuration,omitempty" xml:"MaxDuration,omitempty"`
	// The upper limit of the search time range, in UTC format (YYYY-mm-dd or YYYY-MM-DD HH:mm:ss). By default, the value is (current time +10 minutes)
	//
	// example:
	//
	// 2024-01-31
	//
	// 2024-12-31 23:59:59
	MaxTime     *string  `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	MinDuration *float32 `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	// The lower limit of the search time range, in UTC format (YYYY-mm-dd or YYYY-MM-DD HH:mm:ss). By default, the value is (current time - 2 days).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-31
	//
	// 2024-12-31 23:59:59
	MinTime *string `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	// Whether the returned JSON data can be directly converted to OpenTelemetry TracesData protobuf object. Default value: False. JSON data that is compatible with OpenTelemetry is more complex. Such data is generally not required unless you want to generate a protobuf object of OpenTelemetry.
	//
	// example:
	//
	// False
	OpentelemetryCompatible *bool   `json:"OpentelemetryCompatible,omitempty" xml:"OpentelemetryCompatible,omitempty"`
	OwnerId                 *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The value of the resources.service.owner.sub_id field in the trace record. It can contain upper and lower case letters, digits, dot (.), hyphen (-), and underscore (_). It is empty by default.
	//
	// example:
	//
	// 123456789
	OwnerSubId *string `json:"OwnerSubId,omitempty" xml:"OwnerSubId,omitempty"`
	// The page number. Page starts from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field used to sort the returned results. Valid values: StartTime and Duration.
	//
	// example:
	//
	// StartTime
	//
	// Duration
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- **ASC**
	//
	// 	- **DESC*	- (default)
	//
	// example:
	//
	// DESC
	//
	// ASC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The list of span IDs. Each trace record contains one or more spans.
	SpanIds  []*string `json:"SpanIds,omitempty" xml:"SpanIds,omitempty" type:"Repeated"`
	SpanName *string   `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	// The list of trace IDs.
	TraceIds []*string `json:"TraceIds,omitempty" xml:"TraceIds,omitempty" type:"Repeated"`
	// The content simplification method for returned trace data to reduce the data volume.
	//
	// REMOVE_EMBEDDING: Removes all embedding array contents.
	//
	// ROOT_ONLY: Returns only the root span for each trace, with the root span content also having the REMOVE_EMBEDDING applied.
	//
	// Blank: Maintains the original data without simplification.
	//
	// example:
	//
	// REMOVE_EMBEDDING
	//
	// ROOT_ONLY
	TraceReduceMethod *string `json:"TraceReduceMethod,omitempty" xml:"TraceReduceMethod,omitempty"`
}

func (s ListTracesDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTracesDatasRequest) GoString() string {
	return s.String()
}

func (s *ListTracesDatasRequest) SetEndUserId(v string) *ListTracesDatasRequest {
	s.EndUserId = &v
	return s
}

func (s *ListTracesDatasRequest) SetFilters(v []*ListTracesDatasRequestFilters) *ListTracesDatasRequest {
	s.Filters = v
	return s
}

func (s *ListTracesDatasRequest) SetHasEvents(v bool) *ListTracesDatasRequest {
	s.HasEvents = &v
	return s
}

func (s *ListTracesDatasRequest) SetHasStatusMessage(v bool) *ListTracesDatasRequest {
	s.HasStatusMessage = &v
	return s
}

func (s *ListTracesDatasRequest) SetLlmAppName(v string) *ListTracesDatasRequest {
	s.LlmAppName = &v
	return s
}

func (s *ListTracesDatasRequest) SetMaxDuration(v float32) *ListTracesDatasRequest {
	s.MaxDuration = &v
	return s
}

func (s *ListTracesDatasRequest) SetMaxTime(v string) *ListTracesDatasRequest {
	s.MaxTime = &v
	return s
}

func (s *ListTracesDatasRequest) SetMinDuration(v float32) *ListTracesDatasRequest {
	s.MinDuration = &v
	return s
}

func (s *ListTracesDatasRequest) SetMinTime(v string) *ListTracesDatasRequest {
	s.MinTime = &v
	return s
}

func (s *ListTracesDatasRequest) SetOpentelemetryCompatible(v bool) *ListTracesDatasRequest {
	s.OpentelemetryCompatible = &v
	return s
}

func (s *ListTracesDatasRequest) SetOwnerId(v string) *ListTracesDatasRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTracesDatasRequest) SetOwnerSubId(v string) *ListTracesDatasRequest {
	s.OwnerSubId = &v
	return s
}

func (s *ListTracesDatasRequest) SetPageNumber(v int32) *ListTracesDatasRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTracesDatasRequest) SetPageSize(v int32) *ListTracesDatasRequest {
	s.PageSize = &v
	return s
}

func (s *ListTracesDatasRequest) SetSortBy(v string) *ListTracesDatasRequest {
	s.SortBy = &v
	return s
}

func (s *ListTracesDatasRequest) SetSortOrder(v string) *ListTracesDatasRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTracesDatasRequest) SetSpanIds(v []*string) *ListTracesDatasRequest {
	s.SpanIds = v
	return s
}

func (s *ListTracesDatasRequest) SetSpanName(v string) *ListTracesDatasRequest {
	s.SpanName = &v
	return s
}

func (s *ListTracesDatasRequest) SetTraceIds(v []*string) *ListTracesDatasRequest {
	s.TraceIds = v
	return s
}

func (s *ListTracesDatasRequest) SetTraceReduceMethod(v string) *ListTracesDatasRequest {
	s.TraceReduceMethod = &v
	return s
}

type ListTracesDatasRequestFilters struct {
	// The name of the filter parameter, case-insensitive. Supported parameters: \\"serviceid\\", \\"servicename\\", \\"input\\", \\"output\\", \\"status\\", \\"tracetype\\", and \\"tracename\\".
	//
	// The otel span attributes corresponding to the parameters:
	//
	// serviceid: resources.service.id
	//
	// servicename: resources.service.name
	//
	// input: attributes.input.value
	//
	// output: attributes.output.value
	//
	// status: statusCode
	//
	// tracetype: the attributes.gen_ai.span.kind of span whose parentSpanId is 0
	//
	// tracename: the spanName of span whose parentSpanId is 0
	//
	// Valid values:
	//
	// 	- Status
	//
	// 	- SpanName
	//
	// 	- Input
	//
	// 	- TraceType
	//
	// 	- SpanType
	//
	// 	- ServiceName
	//
	// 	- Output
	//
	// 	- TraceName
	//
	// 	- ServiceId
	//
	// example:
	//
	// output
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The parameter operator. Case-insensitive. Supported operators: \\"=\\", \\"contains\\", and \\"startswith\\".
	//
	// Valid values:
	//
	// 	- contains
	//
	// 	- \\=
	//
	// 	- startsWith
	//
	// example:
	//
	// contains
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the filter parameter. For the contains operation, it is case-sensitive. For other operations, it is case-insensitive.
	//
	// example:
	//
	// cretain filter string
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTracesDatasRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListTracesDatasRequestFilters) GoString() string {
	return s.String()
}

func (s *ListTracesDatasRequestFilters) SetKey(v string) *ListTracesDatasRequestFilters {
	s.Key = &v
	return s
}

func (s *ListTracesDatasRequestFilters) SetOperator(v string) *ListTracesDatasRequestFilters {
	s.Operator = &v
	return s
}

func (s *ListTracesDatasRequestFilters) SetValue(v string) *ListTracesDatasRequestFilters {
	s.Value = &v
	return s
}

type ListTracesDatasShrinkRequest struct {
	// The value of the attributes.service.app.user_id field in the trace record. It can contain upper and lower case letters, digits, dot (.), hyphen (-), and underscore (_). It is empty by default.
	//
	// example:
	//
	// end-user.12345
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Other filter parameters
	FiltersShrink *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// Whether to return only trace records containing spans with a non-empty events. Example: Suppose a trace has 3 spans. If this parameter is True, this trace meets the condition when any one of the 3 spans has a non-empty events. The default value is False. The events is not used for filtering.
	//
	// example:
	//
	// False
	HasEvents *bool `json:"HasEvents,omitempty" xml:"HasEvents,omitempty"`
	// Whether to return only trace records containing spans with a non-empty statusMessage. Example: Suppose a trace has 3 spans. If this parameter is True, this trace meets the condition when any one of the 3 spans has a non-empty statusMessage. The default value is False. The statusMessage is not used for filtering.
	//
	// example:
	//
	// False
	HasStatusMessage *bool `json:"HasStatusMessage,omitempty" xml:"HasStatusMessage,omitempty"`
	// The value of the resources.service.app.name field in the trace record. It can contain upper and lower case letters, digits, dot (.), hyphen (-), and underscore (_). Must be an exact match. It is empty by default.
	//
	// example:
	//
	// My.super_LLM-app2
	LlmAppName  *string  `json:"LlmAppName,omitempty" xml:"LlmAppName,omitempty"`
	MaxDuration *float32 `json:"MaxDuration,omitempty" xml:"MaxDuration,omitempty"`
	// The upper limit of the search time range, in UTC format (YYYY-mm-dd or YYYY-MM-DD HH:mm:ss). By default, the value is (current time +10 minutes)
	//
	// example:
	//
	// 2024-01-31
	//
	// 2024-12-31 23:59:59
	MaxTime     *string  `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	MinDuration *float32 `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	// The lower limit of the search time range, in UTC format (YYYY-mm-dd or YYYY-MM-DD HH:mm:ss). By default, the value is (current time - 2 days).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-31
	//
	// 2024-12-31 23:59:59
	MinTime *string `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	// Whether the returned JSON data can be directly converted to OpenTelemetry TracesData protobuf object. Default value: False. JSON data that is compatible with OpenTelemetry is more complex. Such data is generally not required unless you want to generate a protobuf object of OpenTelemetry.
	//
	// example:
	//
	// False
	OpentelemetryCompatible *bool   `json:"OpentelemetryCompatible,omitempty" xml:"OpentelemetryCompatible,omitempty"`
	OwnerId                 *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The value of the resources.service.owner.sub_id field in the trace record. It can contain upper and lower case letters, digits, dot (.), hyphen (-), and underscore (_). It is empty by default.
	//
	// example:
	//
	// 123456789
	OwnerSubId *string `json:"OwnerSubId,omitempty" xml:"OwnerSubId,omitempty"`
	// The page number. Page starts from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field used to sort the returned results. Valid values: StartTime and Duration.
	//
	// example:
	//
	// StartTime
	//
	// Duration
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- **ASC**
	//
	// 	- **DESC*	- (default)
	//
	// example:
	//
	// DESC
	//
	// ASC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The list of span IDs. Each trace record contains one or more spans.
	SpanIdsShrink *string `json:"SpanIds,omitempty" xml:"SpanIds,omitempty"`
	SpanName      *string `json:"SpanName,omitempty" xml:"SpanName,omitempty"`
	// The list of trace IDs.
	TraceIdsShrink *string `json:"TraceIds,omitempty" xml:"TraceIds,omitempty"`
	// The content simplification method for returned trace data to reduce the data volume.
	//
	// REMOVE_EMBEDDING: Removes all embedding array contents.
	//
	// ROOT_ONLY: Returns only the root span for each trace, with the root span content also having the REMOVE_EMBEDDING applied.
	//
	// Blank: Maintains the original data without simplification.
	//
	// example:
	//
	// REMOVE_EMBEDDING
	//
	// ROOT_ONLY
	TraceReduceMethod *string `json:"TraceReduceMethod,omitempty" xml:"TraceReduceMethod,omitempty"`
}

func (s ListTracesDatasShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTracesDatasShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTracesDatasShrinkRequest) SetEndUserId(v string) *ListTracesDatasShrinkRequest {
	s.EndUserId = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetFiltersShrink(v string) *ListTracesDatasShrinkRequest {
	s.FiltersShrink = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetHasEvents(v bool) *ListTracesDatasShrinkRequest {
	s.HasEvents = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetHasStatusMessage(v bool) *ListTracesDatasShrinkRequest {
	s.HasStatusMessage = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetLlmAppName(v string) *ListTracesDatasShrinkRequest {
	s.LlmAppName = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetMaxDuration(v float32) *ListTracesDatasShrinkRequest {
	s.MaxDuration = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetMaxTime(v string) *ListTracesDatasShrinkRequest {
	s.MaxTime = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetMinDuration(v float32) *ListTracesDatasShrinkRequest {
	s.MinDuration = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetMinTime(v string) *ListTracesDatasShrinkRequest {
	s.MinTime = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetOpentelemetryCompatible(v bool) *ListTracesDatasShrinkRequest {
	s.OpentelemetryCompatible = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetOwnerId(v string) *ListTracesDatasShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetOwnerSubId(v string) *ListTracesDatasShrinkRequest {
	s.OwnerSubId = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetPageNumber(v int32) *ListTracesDatasShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetPageSize(v int32) *ListTracesDatasShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetSortBy(v string) *ListTracesDatasShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetSortOrder(v string) *ListTracesDatasShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetSpanIdsShrink(v string) *ListTracesDatasShrinkRequest {
	s.SpanIdsShrink = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetSpanName(v string) *ListTracesDatasShrinkRequest {
	s.SpanName = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetTraceIdsShrink(v string) *ListTracesDatasShrinkRequest {
	s.TraceIdsShrink = &v
	return s
}

func (s *ListTracesDatasShrinkRequest) SetTraceReduceMethod(v string) *ListTracesDatasShrinkRequest {
	s.TraceReduceMethod = &v
	return s
}

type ListTracesDatasResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// ExecutionFailure
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// failed to get trace data
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// POP request id
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of traces that meet the condition.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The JSON array with each element being a trace\\"s JSON string. Length of the array is equal to or less than the page size parameter value.
	Traces []interface{} `json:"Traces,omitempty" xml:"Traces,omitempty" type:"Repeated"`
}

func (s ListTracesDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTracesDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ListTracesDatasResponseBody) SetCode(v string) *ListTracesDatasResponseBody {
	s.Code = &v
	return s
}

func (s *ListTracesDatasResponseBody) SetMessage(v string) *ListTracesDatasResponseBody {
	s.Message = &v
	return s
}

func (s *ListTracesDatasResponseBody) SetRequestId(v string) *ListTracesDatasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTracesDatasResponseBody) SetTotalCount(v int32) *ListTracesDatasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTracesDatasResponseBody) SetTraces(v []interface{}) *ListTracesDatasResponseBody {
	s.Traces = v
	return s
}

type ListTracesDatasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTracesDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTracesDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTracesDatasResponse) GoString() string {
	return s.String()
}

func (s *ListTracesDatasResponse) SetHeaders(v map[string]*string) *ListTracesDatasResponse {
	s.Headers = v
	return s
}

func (s *ListTracesDatasResponse) SetStatusCode(v int32) *ListTracesDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTracesDatasResponse) SetBody(v *ListTracesDatasResponseBody) *ListTracesDatasResponse {
	s.Body = v
	return s
}

type StopOnlineEvalTaskResponseBody struct {
	// Internal error code. Set only when the response is in error.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response error message. Set only when the response is in error.
	//
	// example:
	//
	// task id is empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the POP request
	//
	// example:
	//
	// 31E5FBC2-792D-5B5C-A5EB-3019984ABFC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopOnlineEvalTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopOnlineEvalTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopOnlineEvalTaskResponseBody) SetCode(v string) *StopOnlineEvalTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopOnlineEvalTaskResponseBody) SetMessage(v string) *StopOnlineEvalTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopOnlineEvalTaskResponseBody) SetRequestId(v string) *StopOnlineEvalTaskResponseBody {
	s.RequestId = &v
	return s
}

type StopOnlineEvalTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopOnlineEvalTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopOnlineEvalTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopOnlineEvalTaskResponse) GoString() string {
	return s.String()
}

func (s *StopOnlineEvalTaskResponse) SetHeaders(v map[string]*string) *StopOnlineEvalTaskResponse {
	s.Headers = v
	return s
}

func (s *StopOnlineEvalTaskResponse) SetStatusCode(v int32) *StopOnlineEvalTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopOnlineEvalTaskResponse) SetBody(v *StopOnlineEvalTaskResponseBody) *StopOnlineEvalTaskResponse {
	s.Body = v
	return s
}

type UpdateOnlineEvalTaskRequest struct {
	// The user-defined name of the LLM application.
	//
	// example:
	//
	// my-llm-one
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the task.
	//
	// example:
	//
	// The purpose of the task is xxx.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end time of the trace data, in UTC format.
	//
	// example:
	//
	// 2024-09-02 22:24:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This configuration structure defines the JSON paths needed to extract specific values from trace data in JSON format. EvaluationConfig defines these JSON paths.
	EvaluationConfig *EvaluationConfig `json:"EvaluationConfig,omitempty" xml:"EvaluationConfig,omitempty"`
	// The evaluation task must search for a certain amount of trace data generated by the user application as input data for the evaluation. This list defines the search filter conditions.
	Filters []*UpdateOnlineEvalTaskRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The access configuration structure of the model used in the evaluation.
	ModelConfig *ModelConfig `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// The evaluation task must search for a certain amount of trace data generated by the user application as input data for the evaluation. This is the width of the time window for each search of input data.
	//
	// example:
	//
	// 10
	SamplingFrequencyMinutes *int32 `json:"SamplingFrequencyMinutes,omitempty" xml:"SamplingFrequencyMinutes,omitempty"`
	// The percentage of data found in a time window that truly serves as evaluation input data. For example, 100 indicates that all data searched is used as evaluation input. 20 indicates that 20% of the found data is randomly selected as evaluation input.
	//
	// example:
	//
	// 50
	SamplingRatio *int32 `json:"SamplingRatio,omitempty" xml:"SamplingRatio,omitempty"`
	// The start time of the trace data, in UTC format.
	//
	// example:
	//
	// 2024-07-31 08:30:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task name.
	//
	// example:
	//
	// foo model application performance evaluation
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateOnlineEvalTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOnlineEvalTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateOnlineEvalTaskRequest) SetAppName(v string) *UpdateOnlineEvalTaskRequest {
	s.AppName = &v
	return s
}

func (s *UpdateOnlineEvalTaskRequest) SetDescription(v string) *UpdateOnlineEvalTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateOnlineEvalTaskRequest) SetEndTime(v string) *UpdateOnlineEvalTaskRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateOnlineEvalTaskRequest) SetEvaluationConfig(v *EvaluationConfig) *UpdateOnlineEvalTaskRequest {
	s.EvaluationConfig = v
	return s
}

func (s *UpdateOnlineEvalTaskRequest) SetFilters(v []*UpdateOnlineEvalTaskRequestFilters) *UpdateOnlineEvalTaskRequest {
	s.Filters = v
	return s
}

func (s *UpdateOnlineEvalTaskRequest) SetModelConfig(v *ModelConfig) *UpdateOnlineEvalTaskRequest {
	s.ModelConfig = v
	return s
}

func (s *UpdateOnlineEvalTaskRequest) SetSamplingFrequencyMinutes(v int32) *UpdateOnlineEvalTaskRequest {
	s.SamplingFrequencyMinutes = &v
	return s
}

func (s *UpdateOnlineEvalTaskRequest) SetSamplingRatio(v int32) *UpdateOnlineEvalTaskRequest {
	s.SamplingRatio = &v
	return s
}

func (s *UpdateOnlineEvalTaskRequest) SetStartTime(v string) *UpdateOnlineEvalTaskRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateOnlineEvalTaskRequest) SetTaskName(v string) *UpdateOnlineEvalTaskRequest {
	s.TaskName = &v
	return s
}

type UpdateOnlineEvalTaskRequestFilters struct {
	// The key of the filter condition.
	//
	// Valid values:
	//
	// 	- Status
	//
	// 	- SpanName
	//
	// 	- Input
	//
	// 	- TraceType
	//
	// 	- SpanType
	//
	// 	- ServiceName
	//
	// 	- Output
	//
	// 	- TraceName
	//
	// 	- ServiceId
	//
	// example:
	//
	// ServiceId
	//
	// ServiceName
	//
	// Input
	//
	// Output
	//
	// Status
	//
	// TraceType
	//
	// SpanType
	//
	// TraceName
	//
	// SpanName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching operator of the filter condition.
	//
	// Valid values:
	//
	// 	- Contains
	//
	// 	- \\=
	//
	// 	- StartsWith
	//
	// example:
	//
	// =
	//
	// StartsWith
	//
	// Contains
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the filter condition.
	//
	// example:
	//
	// foo
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateOnlineEvalTaskRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s UpdateOnlineEvalTaskRequestFilters) GoString() string {
	return s.String()
}

func (s *UpdateOnlineEvalTaskRequestFilters) SetKey(v string) *UpdateOnlineEvalTaskRequestFilters {
	s.Key = &v
	return s
}

func (s *UpdateOnlineEvalTaskRequestFilters) SetOperator(v string) *UpdateOnlineEvalTaskRequestFilters {
	s.Operator = &v
	return s
}

func (s *UpdateOnlineEvalTaskRequestFilters) SetValue(v string) *UpdateOnlineEvalTaskRequestFilters {
	s.Value = &v
	return s
}

type UpdateOnlineEvalTaskResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// InvalidInputParams
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// cannot modify a stopped task
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the POP request
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOnlineEvalTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOnlineEvalTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOnlineEvalTaskResponseBody) SetCode(v string) *UpdateOnlineEvalTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateOnlineEvalTaskResponseBody) SetMessage(v string) *UpdateOnlineEvalTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateOnlineEvalTaskResponseBody) SetRequestId(v string) *UpdateOnlineEvalTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOnlineEvalTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOnlineEvalTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOnlineEvalTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOnlineEvalTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateOnlineEvalTaskResponse) SetHeaders(v map[string]*string) *UpdateOnlineEvalTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateOnlineEvalTaskResponse) SetStatusCode(v int32) *UpdateOnlineEvalTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOnlineEvalTaskResponse) SetBody(v *UpdateOnlineEvalTaskResponseBody) *UpdateOnlineEvalTaskResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("paillmtrace"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// Creates a trace evaluation task. The system will sample some data from the user\\"s trace data based on the task\\"s configuration. Then, an LLM is used to evaluate the performance of these traces, and the evaluation results are recorded.
//
// @param tmpReq - CreateOnlineEvalTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOnlineEvalTaskResponse
func (client *Client) CreateOnlineEvalTaskWithOptions(tmpReq *CreateOnlineEvalTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateOnlineEvalTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateOnlineEvalTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOnlineEvalTask"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/onlineevaltasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOnlineEvalTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a trace evaluation task. The system will sample some data from the user\\"s trace data based on the task\\"s configuration. Then, an LLM is used to evaluate the performance of these traces, and the evaluation results are recorded.
//
// @param request - CreateOnlineEvalTaskRequest
//
// @return CreateOnlineEvalTaskResponse
func (client *Client) CreateOnlineEvalTask(request *CreateOnlineEvalTaskRequest) (_result *CreateOnlineEvalTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOnlineEvalTaskResponse{}
	_body, _err := client.CreateOnlineEvalTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service-linked role required for the PaiLLMTrace service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceIdentityRoleResponse
func (client *Client) CreateServiceIdentityRoleWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceIdentityRoleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceIdentityRole"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/ServiceIdentityRole"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceIdentityRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role required for the PaiLLMTrace service.
//
// @return CreateServiceIdentityRoleResponse
func (client *Client) CreateServiceIdentityRole() (_result *CreateServiceIdentityRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceIdentityRoleResponse{}
	_body, _err := client.CreateServiceIdentityRoleWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete an online evaluation task
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOnlineEvalTaskResponse
func (client *Client) DeleteOnlineEvalTaskWithOptions(TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteOnlineEvalTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOnlineEvalTask"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/onlineevaltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteOnlineEvalTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete an online evaluation task
//
// @return DeleteOnlineEvalTaskResponse
func (client *Client) DeleteOnlineEvalTask(TaskId *string) (_result *DeleteOnlineEvalTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteOnlineEvalTaskResponse{}
	_body, _err := client.DeleteOnlineEvalTaskWithOptions(TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Evaluates a specified piece of trace data.
//
// @param request - EvaluateTraceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluateTraceResponse
func (client *Client) EvaluateTraceWithOptions(TraceId *string, request *EvaluateTraceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EvaluateTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.EvaluationConfig)) {
		body["EvaluationConfig"] = request.EvaluationConfig
	}

	if !tea.BoolValue(util.IsUnset(request.EvaluationId)) {
		body["EvaluationId"] = request.EvaluationId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTime)) {
		body["MaxTime"] = request.MaxTime
	}

	if !tea.BoolValue(util.IsUnset(request.MinTime)) {
		body["MinTime"] = request.MinTime
	}

	if !tea.BoolValue(util.IsUnset(request.ModelConfig)) {
		body["ModelConfig"] = request.ModelConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EvaluateTrace"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/eval/trace/" + tea.StringValue(openapiutil.GetEncodeParam(TraceId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EvaluateTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Evaluates a specified piece of trace data.
//
// @param request - EvaluateTraceRequest
//
// @return EvaluateTraceResponse
func (client *Client) EvaluateTrace(TraceId *string, request *EvaluateTraceRequest) (_result *EvaluateTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EvaluateTraceResponse{}
	_body, _err := client.EvaluateTraceWithOptions(TraceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get the content of prompt templates used for evaluation
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEvaluationTemplatesResponse
func (client *Client) GetEvaluationTemplatesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEvaluationTemplatesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetEvaluationTemplates"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/eval/templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEvaluationTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the content of prompt templates used for evaluation
//
// @return GetEvaluationTemplatesResponse
func (client *Client) GetEvaluationTemplates() (_result *GetEvaluationTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEvaluationTemplatesResponse{}
	_body, _err := client.GetEvaluationTemplatesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get the details of an online evaluation task
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOnlineEvalTaskResponse
func (client *Client) GetOnlineEvalTaskWithOptions(TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOnlineEvalTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOnlineEvalTask"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/onlineevaltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOnlineEvalTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the details of an online evaluation task
//
// @return GetOnlineEvalTaskResponse
func (client *Client) GetOnlineEvalTask(TaskId *string) (_result *GetOnlineEvalTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOnlineEvalTaskResponse{}
	_body, _err := client.GetOnlineEvalTaskWithOptions(TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information related to the service-linked role.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceIdentityRoleResponse
func (client *Client) GetServiceIdentityRoleWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceIdentityRoleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceIdentityRole"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/ServiceIdentityRole"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceIdentityRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information related to the service-linked role.
//
// @return GetServiceIdentityRoleResponse
func (client *Client) GetServiceIdentityRole() (_result *GetServiceIdentityRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceIdentityRoleResponse{}
	_body, _err := client.GetServiceIdentityRoleWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the token used in the Xtrace service and the endpoint required for uploading trace data.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetXtraceTokenResponse
func (client *Client) GetXtraceTokenWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetXtraceTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetXtraceToken"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/XtraceToken"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetXtraceTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the token used in the Xtrace service and the endpoint required for uploading trace data.
//
// @return GetXtraceTokenResponse
func (client *Client) GetXtraceToken() (_result *GetXtraceTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetXtraceTokenResponse{}
	_body, _err := client.GetXtraceTokenWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the list of results for trace evaluation. This API is used together with EvaluateTrace. EvaluateTrace starts the evaluation. ListEvalResults obtains the results.
//
// @param tmpReq - ListEvalResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvalResultsResponse
func (client *Client) ListEvalResultsWithOptions(tmpReq *ListEvalResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEvalResultsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListEvalResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RecordIds)) {
		request.RecordIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordIds, tea.String("RecordIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EvaluationId)) {
		query["EvaluationId"] = request.EvaluationId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RecordIdsShrink)) {
		query["RecordIds"] = request.RecordIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEvalResults"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/eval/results"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEvalResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list of results for trace evaluation. This API is used together with EvaluateTrace. EvaluateTrace starts the evaluation. ListEvalResults obtains the results.
//
// @param request - ListEvalResultsRequest
//
// @return ListEvalResultsResponse
func (client *Client) ListEvalResults(request *ListEvalResultsRequest) (_result *ListEvalResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEvalResultsResponse{}
	_body, _err := client.ListEvalResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List the results of online evaluation tasks that meet the criteria
//
// @param tmpReq - ListOnlineEvalTaskResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOnlineEvalTaskResultsResponse
func (client *Client) ListOnlineEvalTaskResultsWithOptions(tmpReq *ListOnlineEvalTaskResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOnlineEvalTaskResultsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListOnlineEvalTaskResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TraceIds)) {
		request.TraceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TraceIds, tea.String("TraceIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EvaluationId)) {
		query["EvaluationId"] = request.EvaluationId
	}

	if !tea.BoolValue(util.IsUnset(request.MostRecentResultsOnly)) {
		query["MostRecentResultsOnly"] = request.MostRecentResultsOnly
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TraceIdsShrink)) {
		query["TraceIds"] = request.TraceIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOnlineEvalTaskResults"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/onlineevaltaskresults"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOnlineEvalTaskResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List the results of online evaluation tasks that meet the criteria
//
// @param request - ListOnlineEvalTaskResultsRequest
//
// @return ListOnlineEvalTaskResultsResponse
func (client *Client) ListOnlineEvalTaskResults(request *ListOnlineEvalTaskResultsRequest) (_result *ListOnlineEvalTaskResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOnlineEvalTaskResultsResponse{}
	_body, _err := client.ListOnlineEvalTaskResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # View online evaluation tasks that meet the criteria
//
// @param request - ListOnlineEvalTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOnlineEvalTasksResponse
func (client *Client) ListOnlineEvalTasksWithOptions(request *ListOnlineEvalTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOnlineEvalTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTime)) {
		query["MaxTime"] = request.MaxTime
	}

	if !tea.BoolValue(util.IsUnset(request.MinTime)) {
		query["MinTime"] = request.MinTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOnlineEvalTasks"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/onlineevaltasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOnlineEvalTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View online evaluation tasks that meet the criteria
//
// @param request - ListOnlineEvalTasksRequest
//
// @return ListOnlineEvalTasksResponse
func (client *Client) ListOnlineEvalTasks(request *ListOnlineEvalTasksRequest) (_result *ListOnlineEvalTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOnlineEvalTasksResponse{}
	_body, _err := client.ListOnlineEvalTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of trace data based on specified criteria.
//
// @param tmpReq - ListTracesDatasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTracesDatasResponse
func (client *Client) ListTracesDatasWithOptions(tmpReq *ListTracesDatasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTracesDatasResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTracesDatasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Filters)) {
		request.FiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filters, tea.String("Filters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SpanIds)) {
		request.SpanIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SpanIds, tea.String("SpanIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TraceIds)) {
		request.TraceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TraceIds, tea.String("TraceIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.FiltersShrink)) {
		query["Filters"] = request.FiltersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HasEvents)) {
		query["HasEvents"] = request.HasEvents
	}

	if !tea.BoolValue(util.IsUnset(request.HasStatusMessage)) {
		query["HasStatusMessage"] = request.HasStatusMessage
	}

	if !tea.BoolValue(util.IsUnset(request.LlmAppName)) {
		query["LlmAppName"] = request.LlmAppName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxDuration)) {
		query["MaxDuration"] = request.MaxDuration
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTime)) {
		query["MaxTime"] = request.MaxTime
	}

	if !tea.BoolValue(util.IsUnset(request.MinDuration)) {
		query["MinDuration"] = request.MinDuration
	}

	if !tea.BoolValue(util.IsUnset(request.MinTime)) {
		query["MinTime"] = request.MinTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpentelemetryCompatible)) {
		query["OpentelemetryCompatible"] = request.OpentelemetryCompatible
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerSubId)) {
		query["OwnerSubId"] = request.OwnerSubId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.SpanIdsShrink)) {
		query["SpanIds"] = request.SpanIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SpanName)) {
		query["SpanName"] = request.SpanName
	}

	if !tea.BoolValue(util.IsUnset(request.TraceIdsShrink)) {
		query["TraceIds"] = request.TraceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TraceReduceMethod)) {
		query["TraceReduceMethod"] = request.TraceReduceMethod
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTracesDatas"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/TracesDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTracesDatasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of trace data based on specified criteria.
//
// @param request - ListTracesDatasRequest
//
// @return ListTracesDatasResponse
func (client *Client) ListTracesDatas(request *ListTracesDatasRequest) (_result *ListTracesDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTracesDatasResponse{}
	_body, _err := client.ListTracesDatasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Stop the execution of an online evaluation task
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopOnlineEvalTaskResponse
func (client *Client) StopOnlineEvalTaskWithOptions(TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopOnlineEvalTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopOnlineEvalTask"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/onlineevaltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopOnlineEvalTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Stop the execution of an online evaluation task
//
// @return StopOnlineEvalTaskResponse
func (client *Client) StopOnlineEvalTask(TaskId *string) (_result *StopOnlineEvalTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopOnlineEvalTaskResponse{}
	_body, _err := client.StopOnlineEvalTaskWithOptions(TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the configuration of a trace evaluation task.
//
// @param request - UpdateOnlineEvalTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOnlineEvalTaskResponse
func (client *Client) UpdateOnlineEvalTaskWithOptions(TaskId *string, request *UpdateOnlineEvalTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateOnlineEvalTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EvaluationConfig)) {
		body["EvaluationConfig"] = request.EvaluationConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		body["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.ModelConfig)) {
		body["ModelConfig"] = request.ModelConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SamplingFrequencyMinutes)) {
		body["SamplingFrequencyMinutes"] = request.SamplingFrequencyMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.SamplingRatio)) {
		body["SamplingRatio"] = request.SamplingRatio
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOnlineEvalTask"),
		Version:     tea.String("2024-03-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/PAILLMTrace/onlineevaltasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOnlineEvalTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the configuration of a trace evaluation task.
//
// @param request - UpdateOnlineEvalTaskRequest
//
// @return UpdateOnlineEvalTaskResponse
func (client *Client) UpdateOnlineEvalTask(TaskId *string, request *UpdateOnlineEvalTaskRequest) (_result *UpdateOnlineEvalTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateOnlineEvalTaskResponse{}
	_body, _err := client.UpdateOnlineEvalTaskWithOptions(TaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
