// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type OpenApiMultiResponse struct {
	// example:
	//
	// []
	Data       []*OpenApiMultiResponseData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrCode    *string                     `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string                     `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenApiMultiResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenApiMultiResponse) GoString() string {
	return s.String()
}

func (s *OpenApiMultiResponse) SetData(v []*OpenApiMultiResponseData) *OpenApiMultiResponse {
	s.Data = v
	return s
}

func (s *OpenApiMultiResponse) SetErrCode(v string) *OpenApiMultiResponse {
	s.ErrCode = &v
	return s
}

func (s *OpenApiMultiResponse) SetErrMessage(v string) *OpenApiMultiResponse {
	s.ErrMessage = &v
	return s
}

func (s *OpenApiMultiResponse) SetHttpStatusCode(v int32) *OpenApiMultiResponse {
	s.HttpStatusCode = &v
	return s
}

func (s *OpenApiMultiResponse) SetRequestId(v string) *OpenApiMultiResponse {
	s.RequestId = &v
	return s
}

func (s *OpenApiMultiResponse) SetSuccess(v bool) *OpenApiMultiResponse {
	s.Success = &v
	return s
}

type OpenApiMultiResponseData struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 456
	Id                  *string                                     `json:"id,omitempty" xml:"id,omitempty"`
	ImageUrl            []*string                                   `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	InferenceImageCount *int32                                      `json:"inferenceImageCount,omitempty" xml:"inferenceImageCount,omitempty"`
	InferenceJobList    []*OpenApiMultiResponseDataInferenceJobList `json:"inferenceJobList,omitempty" xml:"inferenceJobList,omitempty" type:"Repeated"`
	// example:
	//
	// TRAINING
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 0.5
	JobTrainProgress *float64 `json:"jobTrainProgress,omitempty" xml:"jobTrainProgress,omitempty"`
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 可爱熊猫模型训练任务
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// panda
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s OpenApiMultiResponseData) String() string {
	return tea.Prettify(s)
}

func (s OpenApiMultiResponseData) GoString() string {
	return s.String()
}

func (s *OpenApiMultiResponseData) SetCreateTime(v string) *OpenApiMultiResponseData {
	s.CreateTime = &v
	return s
}

func (s *OpenApiMultiResponseData) SetId(v string) *OpenApiMultiResponseData {
	s.Id = &v
	return s
}

func (s *OpenApiMultiResponseData) SetImageUrl(v []*string) *OpenApiMultiResponseData {
	s.ImageUrl = v
	return s
}

func (s *OpenApiMultiResponseData) SetInferenceImageCount(v int32) *OpenApiMultiResponseData {
	s.InferenceImageCount = &v
	return s
}

func (s *OpenApiMultiResponseData) SetInferenceJobList(v []*OpenApiMultiResponseDataInferenceJobList) *OpenApiMultiResponseData {
	s.InferenceJobList = v
	return s
}

func (s *OpenApiMultiResponseData) SetJobStatus(v string) *OpenApiMultiResponseData {
	s.JobStatus = &v
	return s
}

func (s *OpenApiMultiResponseData) SetJobTrainProgress(v float64) *OpenApiMultiResponseData {
	s.JobTrainProgress = &v
	return s
}

func (s *OpenApiMultiResponseData) SetModelId(v string) *OpenApiMultiResponseData {
	s.ModelId = &v
	return s
}

func (s *OpenApiMultiResponseData) SetName(v string) *OpenApiMultiResponseData {
	s.Name = &v
	return s
}

func (s *OpenApiMultiResponseData) SetObjectType(v string) *OpenApiMultiResponseData {
	s.ObjectType = &v
	return s
}

type OpenApiMultiResponseDataInferenceJobList struct {
	// example:
	//
	// 2023-12-25T12:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3220
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// FINISHED
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 0.5
	JobTrainProgress *float64 `json:"jobTrainProgress,omitempty" xml:"jobTrainProgress,omitempty"`
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// promptId-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
	// example:
	//
	// 0000.png
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s OpenApiMultiResponseDataInferenceJobList) String() string {
	return tea.Prettify(s)
}

func (s OpenApiMultiResponseDataInferenceJobList) GoString() string {
	return s.String()
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetCreateTime(v string) *OpenApiMultiResponseDataInferenceJobList {
	s.CreateTime = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetId(v string) *OpenApiMultiResponseDataInferenceJobList {
	s.Id = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetJobStatus(v string) *OpenApiMultiResponseDataInferenceJobList {
	s.JobStatus = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetJobTrainProgress(v float64) *OpenApiMultiResponseDataInferenceJobList {
	s.JobTrainProgress = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetModelId(v string) *OpenApiMultiResponseDataInferenceJobList {
	s.ModelId = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetPromptId(v string) *OpenApiMultiResponseDataInferenceJobList {
	s.PromptId = &v
	return s
}

func (s *OpenApiMultiResponseDataInferenceJobList) SetResultImageUrl(v []*string) *OpenApiMultiResponseDataInferenceJobList {
	s.ResultImageUrl = v
	return s
}

type OpenApiSingleResponse struct {
	// example:
	//
	// []
	Data       *OpenApiSingleResponseData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode    *string                    `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string                    `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenApiSingleResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenApiSingleResponse) GoString() string {
	return s.String()
}

func (s *OpenApiSingleResponse) SetData(v *OpenApiSingleResponseData) *OpenApiSingleResponse {
	s.Data = v
	return s
}

func (s *OpenApiSingleResponse) SetErrCode(v string) *OpenApiSingleResponse {
	s.ErrCode = &v
	return s
}

func (s *OpenApiSingleResponse) SetErrMessage(v string) *OpenApiSingleResponse {
	s.ErrMessage = &v
	return s
}

func (s *OpenApiSingleResponse) SetHttpStatusCode(v int32) *OpenApiSingleResponse {
	s.HttpStatusCode = &v
	return s
}

func (s *OpenApiSingleResponse) SetRequestId(v string) *OpenApiSingleResponse {
	s.RequestId = &v
	return s
}

func (s *OpenApiSingleResponse) SetSuccess(v bool) *OpenApiSingleResponse {
	s.Success = &v
	return s
}

type OpenApiSingleResponseData struct {
	// example:
	//
	// FINISHED
	ModelTrainStatus *string `json:"modelTrainStatus,omitempty" xml:"modelTrainStatus,omitempty"`
}

func (s OpenApiSingleResponseData) String() string {
	return tea.Prettify(s)
}

func (s OpenApiSingleResponseData) GoString() string {
	return s.String()
}

func (s *OpenApiSingleResponseData) SetModelTrainStatus(v string) *OpenApiSingleResponseData {
	s.ModelTrainStatus = &v
	return s
}

type Personalizedtxt2imgAddInferenceJobCmd struct {
	// example:
	//
	// 1
	ImageNumber *int32 `json:"imageNumber,omitempty" xml:"imageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a <special-token> in the snow
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// 1
	Seed *int32 `json:"seed,omitempty" xml:"seed,omitempty"`
}

func (s Personalizedtxt2imgAddInferenceJobCmd) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobCmd) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetImageNumber(v int32) *Personalizedtxt2imgAddInferenceJobCmd {
	s.ImageNumber = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetModelId(v string) *Personalizedtxt2imgAddInferenceJobCmd {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetPrompt(v string) *Personalizedtxt2imgAddInferenceJobCmd {
	s.Prompt = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobCmd) SetSeed(v int32) *Personalizedtxt2imgAddInferenceJobCmd {
	s.Seed = &v
	return s
}

type Personalizedtxt2imgAddModelTrainJobCmd struct {
	// This parameter is required.
	ImageUrl []*string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 熊猫图片生成
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dog
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s Personalizedtxt2imgAddModelTrainJobCmd) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobCmd) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) SetImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobCmd {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) SetName(v string) *Personalizedtxt2imgAddModelTrainJobCmd {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) SetObjectType(v string) *Personalizedtxt2imgAddModelTrainJobCmd {
	s.ObjectType = &v
	return s
}

type Personalizedtxt2imgInferenceJobInfoDTO struct {
	// example:
	//
	// 123456
	CreateUserId *string `json:"createUserId,omitempty" xml:"createUserId,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 123456
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 123456
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 123456
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s Personalizedtxt2imgInferenceJobInfoDTO) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgInferenceJobInfoDTO) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetCreateUserId(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.CreateUserId = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetId(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetJobStatus(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetModelId(v string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgInferenceJobInfoDTO) SetResultImageUrl(v []*string) *Personalizedtxt2imgInferenceJobInfoDTO {
	s.ResultImageUrl = v
	return s
}

type Personalizedtxt2imgModelTrainJobInfoDTO struct {
	CreateTime       *string                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUserId     *string                                   `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	ImageUrl         []*string                                 `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" type:"Repeated"`
	InferenceJobList []*Personalizedtxt2imgInferenceJobInfoDTO `json:"InferenceJobList,omitempty" xml:"InferenceJobList,omitempty" type:"Repeated"`
	JobStatus        *string                                   `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	Name             *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectType       *string                                   `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s Personalizedtxt2imgModelTrainJobInfoDTO) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgModelTrainJobInfoDTO) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetCreateTime(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetCreateUserId(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.CreateUserId = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetImageUrl(v []*string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetInferenceJobList(v []*Personalizedtxt2imgInferenceJobInfoDTO) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.InferenceJobList = v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetJobStatus(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetName(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetObjectType(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.ObjectType = &v
	return s
}

func (s *Personalizedtxt2imgModelTrainJobInfoDTO) SetId(v string) *Personalizedtxt2imgModelTrainJobInfoDTO {
	s.Id = &v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody struct {
	// example:
	//
	// []
	Data []*AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetData(v []*AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.Data = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetErrCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetErrMessage(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetHttpStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetRequestId(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) SetSuccess(v bool) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody {
	s.Success = &v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData struct {
	// example:
	//
	// 10
	FreeConcurrencyCount *int32 `json:"FreeConcurrencyCount,omitempty" xml:"FreeConcurrencyCount,omitempty"`
	// example:
	//
	// 100
	FreeCount *int32 `json:"FreeCount,omitempty" xml:"FreeCount,omitempty"`
	// example:
	//
	// online_ai_algorithm_personalized_text_to_image_call_count
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// AI算法模型-个性化文生图-在线按量调用
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetFreeConcurrencyCount(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.FreeConcurrencyCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetFreeCount(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.FreeCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetServiceCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData) SetServiceName(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBodyData {
	s.ServiceName = &v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) SetHeaders(v map[string]*string) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse {
	s.Headers = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) SetStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse {
	s.StatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) SetBody(v *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponseBody) *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse {
	s.Body = v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody struct {
	// example:
	//
	// []
	Data []*AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetData(v []*AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.Data = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetErrCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetErrMessage(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetHttpStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetRequestId(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetSuccess(v bool) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
	s.Success = &v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData struct {
	// example:
	//
	// 10
	FreeConcurrencyCount *int32 `json:"FreeConcurrencyCount,omitempty" xml:"FreeConcurrencyCount,omitempty"`
	// example:
	//
	// 100
	FreeCount *int32 `json:"FreeCount,omitempty" xml:"FreeCount,omitempty"`
	// example:
	//
	// online_ai_algorithm_personalized_text_to_image_call_count
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// AI算法模型-个性化文生图-在线按量调用
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) SetFreeConcurrencyCount(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData {
	s.FreeConcurrencyCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) SetFreeCount(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData {
	s.FreeCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) SetServiceCode(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData) SetServiceName(v string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBodyData {
	s.ServiceName = &v
	return s
}

type AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) SetHeaders(v map[string]*string) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse {
	s.Headers = v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) SetStatusCode(v int32) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse {
	s.StatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse) SetBody(v *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse {
	s.Body = v
	return s
}

type PersonalizedTextToImageAddInferenceJobRequest struct {
	// example:
	//
	// 1
	ImageNumber *int32 `json:"imageNumber,omitempty" xml:"imageNumber,omitempty"`
	// This parameter is required.
	ImageUrl []*string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a <special-token> in the snow
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// 1
	Seed *int64 `json:"seed,omitempty" xml:"seed,omitempty"`
	// example:
	//
	// 1
	Strength *float64 `json:"strength,omitempty" xml:"strength,omitempty"`
	// example:
	//
	// 800
	TrainSteps *int32 `json:"trainSteps,omitempty" xml:"trainSteps,omitempty"`
}

func (s PersonalizedTextToImageAddInferenceJobRequest) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobRequest) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetImageNumber(v int32) *PersonalizedTextToImageAddInferenceJobRequest {
	s.ImageNumber = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetImageUrl(v []*string) *PersonalizedTextToImageAddInferenceJobRequest {
	s.ImageUrl = v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetPrompt(v string) *PersonalizedTextToImageAddInferenceJobRequest {
	s.Prompt = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetSeed(v int64) *PersonalizedTextToImageAddInferenceJobRequest {
	s.Seed = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetStrength(v float64) *PersonalizedTextToImageAddInferenceJobRequest {
	s.Strength = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobRequest) SetTrainSteps(v int32) *PersonalizedTextToImageAddInferenceJobRequest {
	s.TrainSteps = &v
	return s
}

type PersonalizedTextToImageAddInferenceJobResponseBody struct {
	// example:
	//
	// []
	Data *PersonalizedTextToImageAddInferenceJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PersonalizedTextToImageAddInferenceJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobResponseBody) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetData(v *PersonalizedTextToImageAddInferenceJobResponseBodyData) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.Data = v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetErrCode(v string) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetErrMessage(v string) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetHttpStatusCode(v int32) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetRequestId(v string) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBody) SetSuccess(v bool) *PersonalizedTextToImageAddInferenceJobResponseBody {
	s.Success = &v
	return s
}

type PersonalizedTextToImageAddInferenceJobResponseBodyData struct {
	// example:
	//
	// 2023-12-25T12:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3220
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// FINISHED
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 0.5
	JobTrainProgress *float64 `json:"jobTrainProgress,omitempty" xml:"jobTrainProgress,omitempty"`
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// promptId
	//
	// example:
	//
	// promptId-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
	// example:
	//
	// 0000.png
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s PersonalizedTextToImageAddInferenceJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetCreateTime(v string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetId(v string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetJobStatus(v string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetJobTrainProgress(v float64) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetModelId(v string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetPromptId(v string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.PromptId = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponseBodyData) SetResultImageUrl(v []*string) *PersonalizedTextToImageAddInferenceJobResponseBodyData {
	s.ResultImageUrl = v
	return s
}

type PersonalizedTextToImageAddInferenceJobResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PersonalizedTextToImageAddInferenceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PersonalizedTextToImageAddInferenceJobResponse) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobResponse) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) SetHeaders(v map[string]*string) *PersonalizedTextToImageAddInferenceJobResponse {
	s.Headers = v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) SetStatusCode(v int32) *PersonalizedTextToImageAddInferenceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) SetBody(v *PersonalizedTextToImageAddInferenceJobResponseBody) *PersonalizedTextToImageAddInferenceJobResponse {
	s.Body = v
	return s
}

type PersonalizedTextToImageQueryImageAssetRequest struct {
	// example:
	//
	// base64
	EncodeFormat *string `json:"encodeFormat,omitempty" xml:"encodeFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0000.png
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
}

func (s PersonalizedTextToImageQueryImageAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryImageAssetRequest) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryImageAssetRequest) SetEncodeFormat(v string) *PersonalizedTextToImageQueryImageAssetRequest {
	s.EncodeFormat = &v
	return s
}

func (s *PersonalizedTextToImageQueryImageAssetRequest) SetImageId(v string) *PersonalizedTextToImageQueryImageAssetRequest {
	s.ImageId = &v
	return s
}

type PersonalizedTextToImageQueryImageAssetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       interface{}        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PersonalizedTextToImageQueryImageAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryImageAssetResponse) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) SetHeaders(v map[string]*string) *PersonalizedTextToImageQueryImageAssetResponse {
	s.Headers = v
	return s
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) SetStatusCode(v int32) *PersonalizedTextToImageQueryImageAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *PersonalizedTextToImageQueryImageAssetResponse) SetBody(v interface{}) *PersonalizedTextToImageQueryImageAssetResponse {
	s.Body = v
	return s
}

type PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// girl-xxxx-xxxx-xxxx-xxxx
	InferenceJobId *string `json:"inferenceJobId,omitempty" xml:"inferenceJobId,omitempty"`
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) SetInferenceJobId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest {
	s.InferenceJobId = &v
	return s
}

type PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody struct {
	// example:
	//
	// []
	Data *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetData(v *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.Data = v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetErrCode(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetErrMessage(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetHttpStatusCode(v int32) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetRequestId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) SetSuccess(v bool) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	s.Success = &v
	return s
}

type PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData struct {
	// example:
	//
	// 2023-12-25T12:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3220
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// FINISHED
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 0.5
	JobTrainProgress *float64 `json:"jobTrainProgress,omitempty" xml:"jobTrainProgress,omitempty"`
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// promptId
	//
	// example:
	//
	// promptId-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
	// example:
	//
	// 0000.png
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetCreateTime(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetJobStatus(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetJobTrainProgress(v float64) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetModelId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetPromptId(v string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.PromptId = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData) SetResultImageUrl(v []*string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBodyData {
	s.ResultImageUrl = v
	return s
}

type PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse struct {
	Headers    map[string]*string                                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) SetHeaders(v map[string]*string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse {
	s.Headers = v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) SetStatusCode(v int32) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) SetBody(v *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse {
	s.Body = v
	return s
}

type Personalizedtxt2imgAddInferenceJobRequest struct {
	// example:
	//
	// 1
	ImageNumber *int32 `json:"imageNumber,omitempty" xml:"imageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a <special-token> in the snow
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// 1
	Seed *int64 `json:"seed,omitempty" xml:"seed,omitempty"`
}

func (s Personalizedtxt2imgAddInferenceJobRequest) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetImageNumber(v int32) *Personalizedtxt2imgAddInferenceJobRequest {
	s.ImageNumber = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetModelId(v string) *Personalizedtxt2imgAddInferenceJobRequest {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetPrompt(v string) *Personalizedtxt2imgAddInferenceJobRequest {
	s.Prompt = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobRequest) SetSeed(v int64) *Personalizedtxt2imgAddInferenceJobRequest {
	s.Seed = &v
	return s
}

type Personalizedtxt2imgAddInferenceJobResponseBody struct {
	// example:
	//
	// []
	Data *Personalizedtxt2imgAddInferenceJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s Personalizedtxt2imgAddInferenceJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobResponseBody) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetData(v *Personalizedtxt2imgAddInferenceJobResponseBodyData) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.Data = v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetErrCode(v string) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetErrMessage(v string) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetHttpStatusCode(v int32) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetRequestId(v string) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBody) SetSuccess(v bool) *Personalizedtxt2imgAddInferenceJobResponseBody {
	s.Success = &v
	return s
}

type Personalizedtxt2imgAddInferenceJobResponseBodyData struct {
	// example:
	//
	// 2023-12-25T12:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3220
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// FINISHED
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 0.5
	JobTrainProgress *float64 `json:"jobTrainProgress,omitempty" xml:"jobTrainProgress,omitempty"`
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// promptId-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
	// example:
	//
	// 0000.png
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s Personalizedtxt2imgAddInferenceJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetCreateTime(v string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetId(v string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetJobStatus(v string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetJobTrainProgress(v float64) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetModelId(v string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetPromptId(v string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.PromptId = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponseBodyData) SetResultImageUrl(v []*string) *Personalizedtxt2imgAddInferenceJobResponseBodyData {
	s.ResultImageUrl = v
	return s
}

type Personalizedtxt2imgAddInferenceJobResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgAddInferenceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgAddInferenceJobResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgAddInferenceJobResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) SetStatusCode(v int32) *Personalizedtxt2imgAddInferenceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) SetBody(v *Personalizedtxt2imgAddInferenceJobResponseBody) *Personalizedtxt2imgAddInferenceJobResponse {
	s.Body = v
	return s
}

type Personalizedtxt2imgAddModelTrainJobRequest struct {
	// This parameter is required.
	ImageUrl []*string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 熊猫图片生成
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dog
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// example:
	//
	// 800
	TrainSteps *int32 `json:"trainSteps,omitempty" xml:"trainSteps,omitempty"`
}

func (s Personalizedtxt2imgAddModelTrainJobRequest) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetName(v string) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetObjectType(v string) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.ObjectType = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetTrainSteps(v int32) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.TrainSteps = &v
	return s
}

type Personalizedtxt2imgAddModelTrainJobResponseBody struct {
	// example:
	//
	// []
	Data *Personalizedtxt2imgAddModelTrainJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBody) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetData(v *Personalizedtxt2imgAddModelTrainJobResponseBodyData) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.Data = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetErrCode(v string) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetErrMessage(v string) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetHttpStatusCode(v int32) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetRequestId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBody) SetSuccess(v bool) *Personalizedtxt2imgAddModelTrainJobResponseBody {
	s.Success = &v
	return s
}

type Personalizedtxt2imgAddModelTrainJobResponseBodyData struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 456
	Id                  *string                                                                `json:"id,omitempty" xml:"id,omitempty"`
	ImageUrl            []*string                                                              `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	InferenceImageCount *int32                                                                 `json:"inferenceImageCount,omitempty" xml:"inferenceImageCount,omitempty"`
	InferenceJobList    []*Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList `json:"inferenceJobList,omitempty" xml:"inferenceJobList,omitempty" type:"Repeated"`
	// example:
	//
	// TRAINING
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 0.5
	JobTrainProgress *float64 `json:"jobTrainProgress,omitempty" xml:"jobTrainProgress,omitempty"`
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 可爱熊猫模型训练任务
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// panda
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetCreateTime(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetInferenceImageCount(v int32) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.InferenceImageCount = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetInferenceJobList(v []*Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.InferenceJobList = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetJobStatus(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetJobTrainProgress(v float64) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetModelId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetName(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyData) SetObjectType(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyData {
	s.ObjectType = &v
	return s
}

type Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList struct {
	// example:
	//
	// 2023-12-25T12:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3220
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// FINISHED
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 0.5
	JobTrainProgress *float64 `json:"jobTrainProgress,omitempty" xml:"jobTrainProgress,omitempty"`
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// promptId-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
	// example:
	//
	// 0000.png
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetCreateTime(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetJobStatus(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetJobTrainProgress(v float64) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetModelId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetPromptId(v string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.PromptId = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList) SetResultImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobResponseBodyDataInferenceJobList {
	s.ResultImageUrl = v
	return s
}

type Personalizedtxt2imgAddModelTrainJobResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgAddModelTrainJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgAddModelTrainJobResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgAddModelTrainJobResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) SetStatusCode(v int32) *Personalizedtxt2imgAddModelTrainJobResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) SetBody(v *Personalizedtxt2imgAddModelTrainJobResponseBody) *Personalizedtxt2imgAddModelTrainJobResponse {
	s.Body = v
	return s
}

type Personalizedtxt2imgQueryImageAssetRequest struct {
	// example:
	//
	// base64
	EncodeFormat *string `json:"encodeFormat,omitempty" xml:"encodeFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0000.png
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// girl-xxxx-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
}

func (s Personalizedtxt2imgQueryImageAssetRequest) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryImageAssetRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetEncodeFormat(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.EncodeFormat = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetImageId(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.ImageId = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetModelId(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetPromptId(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.PromptId = &v
	return s
}

type Personalizedtxt2imgQueryImageAssetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       interface{}        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryImageAssetResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryImageAssetResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryImageAssetResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryImageAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetResponse) SetBody(v interface{}) *Personalizedtxt2imgQueryImageAssetResponse {
	s.Body = v
	return s
}

type Personalizedtxt2imgQueryInferenceJobInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 180
	InferenceJobId *string `json:"inferenceJobId,omitempty" xml:"inferenceJobId,omitempty"`
}

func (s Personalizedtxt2imgQueryInferenceJobInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoRequest) SetInferenceJobId(v string) *Personalizedtxt2imgQueryInferenceJobInfoRequest {
	s.InferenceJobId = &v
	return s
}

type Personalizedtxt2imgQueryInferenceJobInfoResponseBody struct {
	// example:
	//
	// []
	Data *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetData(v *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.Data = v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetErrCode(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.ErrCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetErrMessage(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetHttpStatusCode(v int32) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetRequestId(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) SetSuccess(v bool) *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	s.Success = &v
	return s
}

type Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData struct {
	// example:
	//
	// 2023-12-25T12:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3220
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// FINISHED
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 0.5
	JobTrainProgress *float64 `json:"jobTrainProgress,omitempty" xml:"jobTrainProgress,omitempty"`
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// promptId-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
	// example:
	//
	// 0000.png
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetCreateTime(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetId(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetJobStatus(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetJobTrainProgress(v float64) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetModelId(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetPromptId(v string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.PromptId = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData) SetResultImageUrl(v []*string) *Personalizedtxt2imgQueryInferenceJobInfoResponseBodyData {
	s.ResultImageUrl = v
	return s
}

type Personalizedtxt2imgQueryInferenceJobInfoResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgQueryInferenceJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryInferenceJobInfoResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryInferenceJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) SetBody(v *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) *Personalizedtxt2imgQueryInferenceJobInfoResponse {
	s.Body = v
	return s
}

type Personalizedtxt2imgQueryModelTrainJobListResponseBody struct {
	// example:
	//
	// []
	Data []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBody) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetData(v []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.Data = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetErrCode(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetErrMessage(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetHttpStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetRequestId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBody) SetSuccess(v bool) *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	s.Success = &v
	return s
}

type Personalizedtxt2imgQueryModelTrainJobListResponseBodyData struct {
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 456
	Id                  *string                                                                      `json:"id,omitempty" xml:"id,omitempty"`
	ImageUrl            []*string                                                                    `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	InferenceImageCount *int32                                                                       `json:"inferenceImageCount,omitempty" xml:"inferenceImageCount,omitempty"`
	InferenceJobList    []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList `json:"inferenceJobList,omitempty" xml:"inferenceJobList,omitempty" type:"Repeated"`
	// example:
	//
	// TRAINING
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 0.5
	JobTrainProgress *float64 `json:"jobTrainProgress,omitempty" xml:"jobTrainProgress,omitempty"`
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 可爱熊猫模型训练任务
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// panda
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetCreateTime(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetImageUrl(v []*string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetInferenceImageCount(v int32) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.InferenceImageCount = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetInferenceJobList(v []*Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.InferenceJobList = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetJobStatus(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetJobTrainProgress(v float64) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetModelId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetName(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData) SetObjectType(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyData {
	s.ObjectType = &v
	return s
}

type Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList struct {
	// example:
	//
	// 2023-12-25T12:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3220
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// FINISHED
	JobStatus *string `json:"jobStatus,omitempty" xml:"jobStatus,omitempty"`
	// example:
	//
	// 0.5
	JobTrainProgress *float64 `json:"jobTrainProgress,omitempty" xml:"jobTrainProgress,omitempty"`
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// promptId-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
	// example:
	//
	// 0000.png
	ResultImageUrl []*string `json:"resultImageUrl,omitempty" xml:"resultImageUrl,omitempty" type:"Repeated"`
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetCreateTime(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.CreateTime = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.Id = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetJobStatus(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.JobStatus = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetJobTrainProgress(v float64) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.JobTrainProgress = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetModelId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetPromptId(v string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.PromptId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList) SetResultImageUrl(v []*string) *Personalizedtxt2imgQueryModelTrainJobListResponseBodyDataInferenceJobList {
	s.ResultImageUrl = v
	return s
}

type Personalizedtxt2imgQueryModelTrainJobListResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgQueryModelTrainJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryModelTrainJobListResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) SetBody(v *Personalizedtxt2imgQueryModelTrainJobListResponseBody) *Personalizedtxt2imgQueryModelTrainJobListResponse {
	s.Body = v
	return s
}

type Personalizedtxt2imgQueryModelTrainStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusRequest) SetModelId(v string) *Personalizedtxt2imgQueryModelTrainStatusRequest {
	s.ModelId = &v
	return s
}

type Personalizedtxt2imgQueryModelTrainStatusResponseBody struct {
	// example:
	//
	// []
	Data *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBody) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetData(v *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.Data = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetErrCode(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetErrMessage(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetHttpStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetRequestId(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBody) SetSuccess(v bool) *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	s.Success = &v
	return s
}

type Personalizedtxt2imgQueryModelTrainStatusResponseBodyData struct {
	// example:
	//
	// FINISHED
	ModelTrainStatus *string `json:"modelTrainStatus,omitempty" xml:"modelTrainStatus,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData) SetModelTrainStatus(v string) *Personalizedtxt2imgQueryModelTrainStatusResponseBodyData {
	s.ModelTrainStatus = &v
	return s
}

type Personalizedtxt2imgQueryModelTrainStatusResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgQueryModelTrainStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryModelTrainStatusResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) SetBody(v *Personalizedtxt2imgQueryModelTrainStatusResponseBody) *Personalizedtxt2imgQueryModelTrainStatusResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aicontent"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 阿里云控制台/列出阿里云控制台上可使用的服务列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServcieListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AliyunConsoleOpenApiQueryAliyunConsoleServcieList"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunconsole/queryAliyunConsoleServcieList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 阿里云控制台/列出阿里云控制台上可使用的服务列表
//
// @return AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServcieList() (_result *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse{}
	_body, _err := client.AliyunConsoleOpenApiQueryAliyunConsoleServcieListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 阿里云控制台/列出阿里云控制台上可使用的服务列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServiceListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AliyunConsoleOpenApiQueryAliyunConsoleServiceList"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunconsole/queryAliyunConsoleServiceList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 阿里云控制台/列出阿里云控制台上可使用的服务列表
//
// @return AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServiceList() (_result *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponse{}
	_body, _err := client.AliyunConsoleOpenApiQueryAliyunConsoleServiceListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/基于一个预训练模型创建图片推理任务
//
// @param request - PersonalizedTextToImageAddInferenceJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PersonalizedTextToImageAddInferenceJobResponse
func (client *Client) PersonalizedTextToImageAddInferenceJobWithOptions(request *PersonalizedTextToImageAddInferenceJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PersonalizedTextToImageAddInferenceJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageNumber)) {
		body["imageNumber"] = request.ImageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["imageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.Seed)) {
		body["seed"] = request.Seed
	}

	if !tea.BoolValue(util.IsUnset(request.Strength)) {
		body["strength"] = request.Strength
	}

	if !tea.BoolValue(util.IsUnset(request.TrainSteps)) {
		body["trainSteps"] = request.TrainSteps
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PersonalizedTextToImageAddInferenceJob"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/addPreModelInferenceJob"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PersonalizedTextToImageAddInferenceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/基于一个预训练模型创建图片推理任务
//
// @param request - PersonalizedTextToImageAddInferenceJobRequest
//
// @return PersonalizedTextToImageAddInferenceJobResponse
func (client *Client) PersonalizedTextToImageAddInferenceJob(request *PersonalizedTextToImageAddInferenceJobRequest) (_result *PersonalizedTextToImageAddInferenceJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PersonalizedTextToImageAddInferenceJobResponse{}
	_body, _err := client.PersonalizedTextToImageAddInferenceJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/通过唯一的图片编号获取图片内容
//
// @param request - PersonalizedTextToImageQueryImageAssetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PersonalizedTextToImageQueryImageAssetResponse
func (client *Client) PersonalizedTextToImageQueryImageAssetWithOptions(request *PersonalizedTextToImageQueryImageAssetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PersonalizedTextToImageQueryImageAssetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodeFormat)) {
		query["encodeFormat"] = request.EncodeFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["imageId"] = request.ImageId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PersonalizedTextToImageQueryImageAsset"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryImageAssetFromImageId"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("any"),
	}
	_result = &PersonalizedTextToImageQueryImageAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/通过唯一的图片编号获取图片内容
//
// @param request - PersonalizedTextToImageQueryImageAssetRequest
//
// @return PersonalizedTextToImageQueryImageAssetResponse
func (client *Client) PersonalizedTextToImageQueryImageAsset(request *PersonalizedTextToImageQueryImageAssetRequest) (_result *PersonalizedTextToImageQueryImageAssetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PersonalizedTextToImageQueryImageAssetResponse{}
	_body, _err := client.PersonalizedTextToImageQueryImageAssetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询预制模型推理任务的状态
//
// @param request - PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse
func (client *Client) PersonalizedTextToImageQueryPreModelInferenceJobInfoWithOptions(request *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InferenceJobId)) {
		query["inferenceJobId"] = request.InferenceJobId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PersonalizedTextToImageQueryPreModelInferenceJobInfo"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryPreModelInferenceJobInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询预制模型推理任务的状态
//
// @param request - PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest
//
// @return PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse
func (client *Client) PersonalizedTextToImageQueryPreModelInferenceJobInfo(request *PersonalizedTextToImageQueryPreModelInferenceJobInfoRequest) (_result *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse{}
	_body, _err := client.PersonalizedTextToImageQueryPreModelInferenceJobInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/基于一个模型创建图片推理任务
//
// @param request - Personalizedtxt2imgAddInferenceJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgAddInferenceJobResponse
func (client *Client) Personalizedtxt2imgAddInferenceJobWithOptions(request *Personalizedtxt2imgAddInferenceJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgAddInferenceJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageNumber)) {
		body["imageNumber"] = request.ImageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.Seed)) {
		body["seed"] = request.Seed
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgAddInferenceJob"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/addInferenceJob"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &Personalizedtxt2imgAddInferenceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/基于一个模型创建图片推理任务
//
// @param request - Personalizedtxt2imgAddInferenceJobRequest
//
// @return Personalizedtxt2imgAddInferenceJobResponse
func (client *Client) Personalizedtxt2imgAddInferenceJob(request *Personalizedtxt2imgAddInferenceJobRequest) (_result *Personalizedtxt2imgAddInferenceJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgAddInferenceJobResponse{}
	_body, _err := client.Personalizedtxt2imgAddInferenceJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/创建一个模型训练任务
//
// @param request - Personalizedtxt2imgAddModelTrainJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgAddModelTrainJobResponse
func (client *Client) Personalizedtxt2imgAddModelTrainJobWithOptions(request *Personalizedtxt2imgAddModelTrainJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgAddModelTrainJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["imageUrl"] = request.ImageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		body["objectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.TrainSteps)) {
		body["trainSteps"] = request.TrainSteps
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgAddModelTrainJob"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/addModelTrainJob"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &Personalizedtxt2imgAddModelTrainJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/创建一个模型训练任务
//
// @param request - Personalizedtxt2imgAddModelTrainJobRequest
//
// @return Personalizedtxt2imgAddModelTrainJobResponse
func (client *Client) Personalizedtxt2imgAddModelTrainJob(request *Personalizedtxt2imgAddModelTrainJobRequest) (_result *Personalizedtxt2imgAddModelTrainJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgAddModelTrainJobResponse{}
	_body, _err := client.Personalizedtxt2imgAddModelTrainJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/图片二进制内容获取
//
// @param request - Personalizedtxt2imgQueryImageAssetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryImageAssetResponse
func (client *Client) Personalizedtxt2imgQueryImageAssetWithOptions(request *Personalizedtxt2imgQueryImageAssetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgQueryImageAssetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodeFormat)) {
		query["encodeFormat"] = request.EncodeFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["imageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.PromptId)) {
		query["promptId"] = request.PromptId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgQueryImageAsset"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryImageAsset"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("any"),
	}
	_result = &Personalizedtxt2imgQueryImageAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/图片二进制内容获取
//
// @param request - Personalizedtxt2imgQueryImageAssetRequest
//
// @return Personalizedtxt2imgQueryImageAssetResponse
func (client *Client) Personalizedtxt2imgQueryImageAsset(request *Personalizedtxt2imgQueryImageAssetRequest) (_result *Personalizedtxt2imgQueryImageAssetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgQueryImageAssetResponse{}
	_body, _err := client.Personalizedtxt2imgQueryImageAssetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询模型推理任务的状态和结果信息
//
// @param request - Personalizedtxt2imgQueryInferenceJobInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryInferenceJobInfoResponse
func (client *Client) Personalizedtxt2imgQueryInferenceJobInfoWithOptions(request *Personalizedtxt2imgQueryInferenceJobInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgQueryInferenceJobInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InferenceJobId)) {
		query["inferenceJobId"] = request.InferenceJobId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgQueryInferenceJobInfo"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryInferenceJobInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &Personalizedtxt2imgQueryInferenceJobInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询模型推理任务的状态和结果信息
//
// @param request - Personalizedtxt2imgQueryInferenceJobInfoRequest
//
// @return Personalizedtxt2imgQueryInferenceJobInfoResponse
func (client *Client) Personalizedtxt2imgQueryInferenceJobInfo(request *Personalizedtxt2imgQueryInferenceJobInfoRequest) (_result *Personalizedtxt2imgQueryInferenceJobInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgQueryInferenceJobInfoResponse{}
	_body, _err := client.Personalizedtxt2imgQueryInferenceJobInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询模型训练任务列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryModelTrainJobListResponse
func (client *Client) Personalizedtxt2imgQueryModelTrainJobListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgQueryModelTrainJobListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgQueryModelTrainJobList"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryModelTrainJobList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &Personalizedtxt2imgQueryModelTrainJobListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/查询模型训练任务列表
//
// @return Personalizedtxt2imgQueryModelTrainJobListResponse
func (client *Client) Personalizedtxt2imgQueryModelTrainJobList() (_result *Personalizedtxt2imgQueryModelTrainJobListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgQueryModelTrainJobListResponse{}
	_body, _err := client.Personalizedtxt2imgQueryModelTrainJobListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 个性化文生图/模型训练状态查询
//
// @param request - Personalizedtxt2imgQueryModelTrainStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Personalizedtxt2imgQueryModelTrainStatusResponse
func (client *Client) Personalizedtxt2imgQueryModelTrainStatusWithOptions(request *Personalizedtxt2imgQueryModelTrainStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *Personalizedtxt2imgQueryModelTrainStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["modelId"] = request.ModelId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Personalizedtxt2imgQueryModelTrainStatus"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/personalizedtxt2img/queryModelTrainStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &Personalizedtxt2imgQueryModelTrainStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 个性化文生图/模型训练状态查询
//
// @param request - Personalizedtxt2imgQueryModelTrainStatusRequest
//
// @return Personalizedtxt2imgQueryModelTrainStatusResponse
func (client *Client) Personalizedtxt2imgQueryModelTrainStatus(request *Personalizedtxt2imgQueryModelTrainStatusRequest) (_result *Personalizedtxt2imgQueryModelTrainStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &Personalizedtxt2imgQueryModelTrainStatusResponse{}
	_body, _err := client.Personalizedtxt2imgQueryModelTrainStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
