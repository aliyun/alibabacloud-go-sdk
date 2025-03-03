// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ChatMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hello world！
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 723b7f0f-c28a-4123-86e7-094d3d3863f8
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c28a-4123-86e7
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ChatMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatMessageRequest) GoString() string {
	return s.String()
}

func (s *ChatMessageRequest) SetContent(v string) *ChatMessageRequest {
	s.Content = &v
	return s
}

func (s *ChatMessageRequest) SetConversationId(v string) *ChatMessageRequest {
	s.ConversationId = &v
	return s
}

func (s *ChatMessageRequest) SetUserId(v string) *ChatMessageRequest {
	s.UserId = &v
	return s
}

type ChatMessageResponseBody struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// Param.Invalid
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The specified parameter is not valid.
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 696acaa9-eb29-4c1f-b48a-1f901579acc5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ChatMessageResponseBody) SetContent(v string) *ChatMessageResponseBody {
	s.Content = &v
	return s
}

func (s *ChatMessageResponseBody) SetErrCode(v string) *ChatMessageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatMessageResponseBody) SetErrMsg(v string) *ChatMessageResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ChatMessageResponseBody) SetEventType(v string) *ChatMessageResponseBody {
	s.EventType = &v
	return s
}

func (s *ChatMessageResponseBody) SetRequestId(v string) *ChatMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatMessageResponseBody) SetSuccess(v bool) *ChatMessageResponseBody {
	s.Success = &v
	return s
}

type ChatMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatMessageResponse) GoString() string {
	return s.String()
}

func (s *ChatMessageResponse) SetHeaders(v map[string]*string) *ChatMessageResponse {
	s.Headers = v
	return s
}

func (s *ChatMessageResponse) SetStatusCode(v int32) *ChatMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatMessageResponse) SetBody(v *ChatMessageResponseBody) *ChatMessageResponse {
	s.Body = v
	return s
}

type CreateConversationRequest struct {
	ExerciseAnalysis *string `json:"ExerciseAnalysis,omitempty" xml:"ExerciseAnalysis,omitempty"`
	ExerciseAnswer   *string `json:"ExerciseAnswer,omitempty" xml:"ExerciseAnswer,omitempty"`
	// This parameter is required.
	ExerciseContent *string `json:"ExerciseContent,omitempty" xml:"ExerciseContent,omitempty"`
	ExerciseType    *string `json:"ExerciseType,omitempty" xml:"ExerciseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2r560eHAbsknrfHXVZO4L
	OuterBizId *string `json:"OuterBizId,omitempty" xml:"OuterBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// wx-xx-yy
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateConversationRequest) SetExerciseAnalysis(v string) *CreateConversationRequest {
	s.ExerciseAnalysis = &v
	return s
}

func (s *CreateConversationRequest) SetExerciseAnswer(v string) *CreateConversationRequest {
	s.ExerciseAnswer = &v
	return s
}

func (s *CreateConversationRequest) SetExerciseContent(v string) *CreateConversationRequest {
	s.ExerciseContent = &v
	return s
}

func (s *CreateConversationRequest) SetExerciseType(v string) *CreateConversationRequest {
	s.ExerciseType = &v
	return s
}

func (s *CreateConversationRequest) SetOuterBizId(v string) *CreateConversationRequest {
	s.OuterBizId = &v
	return s
}

func (s *CreateConversationRequest) SetUserId(v string) *CreateConversationRequest {
	s.UserId = &v
	return s
}

type CreateConversationResponseBody struct {
	// example:
	//
	// a499fef7-fef7-453c-a6b2-6a34089613e8
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 89C27D03-4C85-5420-9752-989130878F4D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConversationResponseBody) SetConversationId(v string) *CreateConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateConversationResponseBody) SetErrCode(v string) *CreateConversationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateConversationResponseBody) SetErrMsg(v string) *CreateConversationResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateConversationResponseBody) SetRequestId(v string) *CreateConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConversationResponseBody) SetSuccess(v bool) *CreateConversationResponseBody {
	s.Success = &v
	return s
}

type CreateConversationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateConversationResponse) SetHeaders(v map[string]*string) *CreateConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateConversationResponse) SetStatusCode(v int32) *CreateConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConversationResponse) SetBody(v *CreateConversationResponseBody) *CreateConversationResponse {
	s.Body = v
	return s
}

type CreateRelatedConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Ex_pop_1731848070815_funI
	ExerciseCode *string `json:"ExerciseCode,omitempty" xml:"ExerciseCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 725e5550-8f81-42e0-a4db-d2de1be52afc
	OuterBizId *string `json:"OuterBizId,omitempty" xml:"OuterBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pop_1731848070815
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateRelatedConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRelatedConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateRelatedConversationRequest) SetExerciseCode(v string) *CreateRelatedConversationRequest {
	s.ExerciseCode = &v
	return s
}

func (s *CreateRelatedConversationRequest) SetOuterBizId(v string) *CreateRelatedConversationRequest {
	s.OuterBizId = &v
	return s
}

func (s *CreateRelatedConversationRequest) SetUserId(v string) *CreateRelatedConversationRequest {
	s.UserId = &v
	return s
}

type CreateRelatedConversationResponseBody struct {
	// example:
	//
	// 96d36ed0-ebde-11ee-806f-f35ee6682ec5
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1d31b11b-4b82-4db1-b3c0-001529fc78eb
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRelatedConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRelatedConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRelatedConversationResponseBody) SetConversationId(v string) *CreateRelatedConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateRelatedConversationResponseBody) SetErrCode(v string) *CreateRelatedConversationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateRelatedConversationResponseBody) SetErrMsg(v string) *CreateRelatedConversationResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateRelatedConversationResponseBody) SetRequestId(v string) *CreateRelatedConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRelatedConversationResponseBody) SetSuccess(v bool) *CreateRelatedConversationResponseBody {
	s.Success = &v
	return s
}

type CreateRelatedConversationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRelatedConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRelatedConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRelatedConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateRelatedConversationResponse) SetHeaders(v map[string]*string) *CreateRelatedConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateRelatedConversationResponse) SetStatusCode(v int32) *CreateRelatedConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRelatedConversationResponse) SetBody(v *CreateRelatedConversationResponseBody) *CreateRelatedConversationResponse {
	s.Body = v
	return s
}

type GenAnalysisRequest struct {
	// This parameter is required.
	ExerciseContent *string `json:"ExerciseContent,omitempty" xml:"ExerciseContent,omitempty"`
}

func (s GenAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s GenAnalysisRequest) GoString() string {
	return s.String()
}

func (s *GenAnalysisRequest) SetExerciseContent(v string) *GenAnalysisRequest {
	s.ExerciseContent = &v
	return s
}

type GenAnalysisResponseBody struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 40020503
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Parameter value validation failed
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1d31b11b-4b82-4db1-b3c0-001529fc78eb
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *GenAnalysisResponseBody) SetContent(v string) *GenAnalysisResponseBody {
	s.Content = &v
	return s
}

func (s *GenAnalysisResponseBody) SetErrCode(v string) *GenAnalysisResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GenAnalysisResponseBody) SetErrMsg(v string) *GenAnalysisResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GenAnalysisResponseBody) SetEventType(v string) *GenAnalysisResponseBody {
	s.EventType = &v
	return s
}

func (s *GenAnalysisResponseBody) SetRequestId(v string) *GenAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type GenAnalysisResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s GenAnalysisResponse) GoString() string {
	return s.String()
}

func (s *GenAnalysisResponse) SetHeaders(v map[string]*string) *GenAnalysisResponse {
	s.Headers = v
	return s
}

func (s *GenAnalysisResponse) SetStatusCode(v int32) *GenAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *GenAnalysisResponse) SetBody(v *GenAnalysisResponseBody) *GenAnalysisResponse {
	s.Body = v
	return s
}

type GenStepRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Ex_pop_1731848070815_funI
	ExerciseCode *string `json:"ExerciseCode,omitempty" xml:"ExerciseCode,omitempty"`
}

func (s GenStepRequest) String() string {
	return tea.Prettify(s)
}

func (s GenStepRequest) GoString() string {
	return s.String()
}

func (s *GenStepRequest) SetExerciseCode(v string) *GenStepRequest {
	s.ExerciseCode = &v
	return s
}

type GenStepResponseBody struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 40020503
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Parameter value validation failed
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 696acaa9-eb29-4c1f-b48a-1f901579acc5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenStepResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenStepResponseBody) GoString() string {
	return s.String()
}

func (s *GenStepResponseBody) SetContent(v string) *GenStepResponseBody {
	s.Content = &v
	return s
}

func (s *GenStepResponseBody) SetErrCode(v string) *GenStepResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GenStepResponseBody) SetErrMsg(v string) *GenStepResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GenStepResponseBody) SetEventType(v string) *GenStepResponseBody {
	s.EventType = &v
	return s
}

func (s *GenStepResponseBody) SetRequestId(v string) *GenStepResponseBody {
	s.RequestId = &v
	return s
}

type GenStepResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenStepResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenStepResponse) String() string {
	return tea.Prettify(s)
}

func (s GenStepResponse) GoString() string {
	return s.String()
}

func (s *GenStepResponse) SetHeaders(v map[string]*string) *GenStepResponse {
	s.Headers = v
	return s
}

func (s *GenStepResponse) SetStatusCode(v int32) *GenStepResponse {
	s.StatusCode = &v
	return s
}

func (s *GenStepResponse) SetBody(v *GenStepResponseBody) *GenStepResponse {
	s.Body = v
	return s
}

type GlobalConfirmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Ex_pop_1731848070815_funI
	ExerciseCode *string `json:"ExerciseCode,omitempty" xml:"ExerciseCode,omitempty"`
	// example:
	//
	// 2024-11-18
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GlobalConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s GlobalConfirmRequest) GoString() string {
	return s.String()
}

func (s *GlobalConfirmRequest) SetExerciseCode(v string) *GlobalConfirmRequest {
	s.ExerciseCode = &v
	return s
}

func (s *GlobalConfirmRequest) SetTag(v string) *GlobalConfirmRequest {
	s.Tag = &v
	return s
}

type GlobalConfirmResponseBody struct {
	// example:
	//
	// 40020503
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Parameter value validation failed
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 696acaa9-eb29-4c1f-b48a-1f901579acc5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GlobalConfirmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GlobalConfirmResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalConfirmResponseBody) SetErrCode(v string) *GlobalConfirmResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GlobalConfirmResponseBody) SetErrMsg(v string) *GlobalConfirmResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GlobalConfirmResponseBody) SetRequestId(v string) *GlobalConfirmResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalConfirmResponseBody) SetSuccess(v bool) *GlobalConfirmResponseBody {
	s.Success = &v
	return s
}

type GlobalConfirmResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalConfirmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s GlobalConfirmResponse) GoString() string {
	return s.String()
}

func (s *GlobalConfirmResponse) SetHeaders(v map[string]*string) *GlobalConfirmResponse {
	s.Headers = v
	return s
}

func (s *GlobalConfirmResponse) SetStatusCode(v int32) *GlobalConfirmResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalConfirmResponse) SetBody(v *GlobalConfirmResponseBody) *GlobalConfirmResponse {
	s.Body = v
	return s
}

type UpdateAnalysisRequest struct {
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1549d636-b102-4fee-8a99-fcc552aa9aa9
	ContentCode *string `json:"ContentCode,omitempty" xml:"ContentCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Ex_pop_1731848070815_funI
	ExerciseCode *string `json:"ExerciseCode,omitempty" xml:"ExerciseCode,omitempty"`
}

func (s UpdateAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnalysisRequest) GoString() string {
	return s.String()
}

func (s *UpdateAnalysisRequest) SetContent(v string) *UpdateAnalysisRequest {
	s.Content = &v
	return s
}

func (s *UpdateAnalysisRequest) SetContentCode(v string) *UpdateAnalysisRequest {
	s.ContentCode = &v
	return s
}

func (s *UpdateAnalysisRequest) SetExerciseCode(v string) *UpdateAnalysisRequest {
	s.ExerciseCode = &v
	return s
}

type UpdateAnalysisResponseBody struct {
	// example:
	//
	// 9901100002
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Parameter value validation failed
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 89C27D03-4C85-5420-9752-989130878F4D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAnalysisResponseBody) SetErrCode(v string) *UpdateAnalysisResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateAnalysisResponseBody) SetErrMsg(v string) *UpdateAnalysisResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateAnalysisResponseBody) SetRequestId(v string) *UpdateAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAnalysisResponseBody) SetSuccess(v bool) *UpdateAnalysisResponseBody {
	s.Success = &v
	return s
}

type UpdateAnalysisResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAnalysisResponse) GoString() string {
	return s.String()
}

func (s *UpdateAnalysisResponse) SetHeaders(v map[string]*string) *UpdateAnalysisResponse {
	s.Headers = v
	return s
}

func (s *UpdateAnalysisResponse) SetStatusCode(v int32) *UpdateAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAnalysisResponse) SetBody(v *UpdateAnalysisResponseBody) *UpdateAnalysisResponse {
	s.Body = v
	return s
}

type UpdateStepRequest struct {
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1549d636-b102-4fee-8a99-fcc552aa9aa9
	ContentCode *string `json:"ContentCode,omitempty" xml:"ContentCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Ex_pop_1731848070815_funI
	ExerciseCode *string `json:"ExerciseCode,omitempty" xml:"ExerciseCode,omitempty"`
}

func (s UpdateStepRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStepRequest) GoString() string {
	return s.String()
}

func (s *UpdateStepRequest) SetContent(v string) *UpdateStepRequest {
	s.Content = &v
	return s
}

func (s *UpdateStepRequest) SetContentCode(v string) *UpdateStepRequest {
	s.ContentCode = &v
	return s
}

func (s *UpdateStepRequest) SetExerciseCode(v string) *UpdateStepRequest {
	s.ExerciseCode = &v
	return s
}

type UpdateStepResponseBody struct {
	// example:
	//
	// 9901100002
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Parameter value validation failed
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 15486286-243F-51E6-8DD3-B2D2365E5136
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateStepResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStepResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStepResponseBody) SetErrCode(v string) *UpdateStepResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateStepResponseBody) SetErrMsg(v string) *UpdateStepResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateStepResponseBody) SetRequestId(v string) *UpdateStepResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStepResponseBody) SetSuccess(v bool) *UpdateStepResponseBody {
	s.Success = &v
	return s
}

type UpdateStepResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStepResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStepResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStepResponse) GoString() string {
	return s.String()
}

func (s *UpdateStepResponse) SetHeaders(v map[string]*string) *UpdateStepResponse {
	s.Headers = v
	return s
}

func (s *UpdateStepResponse) SetStatusCode(v int32) *UpdateStepResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStepResponse) SetBody(v *UpdateStepResponseBody) *UpdateStepResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aimath"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 聊天消息API
//
// @param request - ChatMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatMessageResponse
func (client *Client) ChatMessageWithOptions(request *ChatMessageRequest, runtime *util.RuntimeOptions) (_result *ChatMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationId)) {
		body["ConversationId"] = request.ConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChatMessage"),
		Version:     tea.String("2024-11-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ChatMessageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ChatMessageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 聊天消息API
//
// @param request - ChatMessageRequest
//
// @return ChatMessageResponse
func (client *Client) ChatMessage(request *ChatMessageRequest) (_result *ChatMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChatMessageResponse{}
	_body, _err := client.ChatMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建试题相应的对话。
//
// @param request - CreateConversationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConversationResponse
func (client *Client) CreateConversationWithOptions(request *CreateConversationRequest, runtime *util.RuntimeOptions) (_result *CreateConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExerciseAnalysis)) {
		body["ExerciseAnalysis"] = request.ExerciseAnalysis
	}

	if !tea.BoolValue(util.IsUnset(request.ExerciseAnswer)) {
		body["ExerciseAnswer"] = request.ExerciseAnswer
	}

	if !tea.BoolValue(util.IsUnset(request.ExerciseContent)) {
		body["ExerciseContent"] = request.ExerciseContent
	}

	if !tea.BoolValue(util.IsUnset(request.ExerciseType)) {
		body["ExerciseType"] = request.ExerciseType
	}

	if !tea.BoolValue(util.IsUnset(request.OuterBizId)) {
		body["OuterBizId"] = request.OuterBizId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConversation"),
		Version:     tea.String("2024-11-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateConversationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateConversationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建试题相应的对话。
//
// @param request - CreateConversationRequest
//
// @return CreateConversationResponse
func (client *Client) CreateConversation(request *CreateConversationRequest) (_result *CreateConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConversationResponse{}
	_body, _err := client.CreateConversationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建关联的对话，输入试题code即可开启对话
//
// @param request - CreateRelatedConversationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRelatedConversationResponse
func (client *Client) CreateRelatedConversationWithOptions(request *CreateRelatedConversationRequest, runtime *util.RuntimeOptions) (_result *CreateRelatedConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExerciseCode)) {
		body["ExerciseCode"] = request.ExerciseCode
	}

	if !tea.BoolValue(util.IsUnset(request.OuterBizId)) {
		body["OuterBizId"] = request.OuterBizId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRelatedConversation"),
		Version:     tea.String("2024-11-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateRelatedConversationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateRelatedConversationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建关联的对话，输入试题code即可开启对话
//
// @param request - CreateRelatedConversationRequest
//
// @return CreateRelatedConversationResponse
func (client *Client) CreateRelatedConversation(request *CreateRelatedConversationRequest) (_result *CreateRelatedConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRelatedConversationResponse{}
	_body, _err := client.CreateRelatedConversationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成解题分析
//
// @param request - GenAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenAnalysisResponse
func (client *Client) GenAnalysisWithOptions(request *GenAnalysisRequest, runtime *util.RuntimeOptions) (_result *GenAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExerciseContent)) {
		body["ExerciseContent"] = request.ExerciseContent
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenAnalysis"),
		Version:     tea.String("2024-11-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenAnalysisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenAnalysisResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 生成解题分析
//
// @param request - GenAnalysisRequest
//
// @return GenAnalysisResponse
func (client *Client) GenAnalysis(request *GenAnalysisRequest) (_result *GenAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenAnalysisResponse{}
	_body, _err := client.GenAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成数学解题步骤
//
// @param request - GenStepRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenStepResponse
func (client *Client) GenStepWithOptions(request *GenStepRequest, runtime *util.RuntimeOptions) (_result *GenStepResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExerciseCode)) {
		body["ExerciseCode"] = request.ExerciseCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenStep"),
		Version:     tea.String("2024-11-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenStepResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenStepResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 生成数学解题步骤
//
// @param request - GenStepRequest
//
// @return GenStepResponse
func (client *Client) GenStep(request *GenStepRequest) (_result *GenStepResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenStepResponse{}
	_body, _err := client.GenStepWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 全局确认修订完成
//
// @param request - GlobalConfirmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalConfirmResponse
func (client *Client) GlobalConfirmWithOptions(request *GlobalConfirmRequest, runtime *util.RuntimeOptions) (_result *GlobalConfirmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExerciseCode)) {
		body["ExerciseCode"] = request.ExerciseCode
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GlobalConfirm"),
		Version:     tea.String("2024-11-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GlobalConfirmResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GlobalConfirmResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 全局确认修订完成
//
// @param request - GlobalConfirmRequest
//
// @return GlobalConfirmResponse
func (client *Client) GlobalConfirm(request *GlobalConfirmRequest) (_result *GlobalConfirmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GlobalConfirmResponse{}
	_body, _err := client.GlobalConfirmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修订解题分析
//
// @param request - UpdateAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAnalysisResponse
func (client *Client) UpdateAnalysisWithOptions(request *UpdateAnalysisRequest, runtime *util.RuntimeOptions) (_result *UpdateAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentCode)) {
		body["ContentCode"] = request.ContentCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExerciseCode)) {
		body["ExerciseCode"] = request.ExerciseCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAnalysis"),
		Version:     tea.String("2024-11-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateAnalysisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateAnalysisResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 修订解题分析
//
// @param request - UpdateAnalysisRequest
//
// @return UpdateAnalysisResponse
func (client *Client) UpdateAnalysis(request *UpdateAnalysisRequest) (_result *UpdateAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAnalysisResponse{}
	_body, _err := client.UpdateAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修订解题步骤
//
// @param request - UpdateStepRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStepResponse
func (client *Client) UpdateStepWithOptions(request *UpdateStepRequest, runtime *util.RuntimeOptions) (_result *UpdateStepResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentCode)) {
		body["ContentCode"] = request.ContentCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExerciseCode)) {
		body["ExerciseCode"] = request.ExerciseCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStep"),
		Version:     tea.String("2024-11-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateStepResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateStepResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 修订解题步骤
//
// @param request - UpdateStepRequest
//
// @return UpdateStepResponse
func (client *Client) UpdateStep(request *UpdateStepRequest) (_result *UpdateStepResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStepResponse{}
	_body, _err := client.UpdateStepWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
