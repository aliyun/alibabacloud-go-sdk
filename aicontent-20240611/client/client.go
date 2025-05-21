// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AliyunConsoleServiceInfoDTO struct {
	// example:
	//
	// https://www.aliyun.com/product/ai-algorithm
	BuyUrl *string `json:"buyUrl,omitempty" xml:"buyUrl,omitempty"`
	// example:
	//
	// https://www.aliyun.com/product/ai-algorithm
	DocumentUrl *string `json:"documentUrl,omitempty" xml:"documentUrl,omitempty"`
	// example:
	//
	// 10
	FreeConcurrencyCount *int32 `json:"freeConcurrencyCount,omitempty" xml:"freeConcurrencyCount,omitempty"`
	// example:
	//
	// 100
	FreeCount *int32 `json:"freeCount,omitempty" xml:"freeCount,omitempty"`
	// example:
	//
	// online_ai_algorithm_personalized_text_to_image_call_count
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	// example:
	//
	// AI算法模型-个性化文生图-在线按量调用
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s AliyunConsoleServiceInfoDTO) String() string {
	return tea.Prettify(s)
}

func (s AliyunConsoleServiceInfoDTO) GoString() string {
	return s.String()
}

func (s *AliyunConsoleServiceInfoDTO) SetBuyUrl(v string) *AliyunConsoleServiceInfoDTO {
	s.BuyUrl = &v
	return s
}

func (s *AliyunConsoleServiceInfoDTO) SetDocumentUrl(v string) *AliyunConsoleServiceInfoDTO {
	s.DocumentUrl = &v
	return s
}

func (s *AliyunConsoleServiceInfoDTO) SetFreeConcurrencyCount(v int32) *AliyunConsoleServiceInfoDTO {
	s.FreeConcurrencyCount = &v
	return s
}

func (s *AliyunConsoleServiceInfoDTO) SetFreeCount(v int32) *AliyunConsoleServiceInfoDTO {
	s.FreeCount = &v
	return s
}

func (s *AliyunConsoleServiceInfoDTO) SetServiceCode(v string) *AliyunConsoleServiceInfoDTO {
	s.ServiceCode = &v
	return s
}

func (s *AliyunConsoleServiceInfoDTO) SetServiceName(v string) *AliyunConsoleServiceInfoDTO {
	s.ServiceName = &v
	return s
}

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

type OralEvaluationStatisticsCallsCountRequest struct {
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	EndTime             *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Granularity         *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	ProjectId           *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	StartTime           *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s OralEvaluationStatisticsCallsCountRequest) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsCallsCountRequest) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsCallsCountRequest) SetApplicationAccessId(v string) *OralEvaluationStatisticsCallsCountRequest {
	s.ApplicationAccessId = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountRequest) SetEndTime(v string) *OralEvaluationStatisticsCallsCountRequest {
	s.EndTime = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountRequest) SetGranularity(v string) *OralEvaluationStatisticsCallsCountRequest {
	s.Granularity = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountRequest) SetProjectId(v string) *OralEvaluationStatisticsCallsCountRequest {
	s.ProjectId = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountRequest) SetStartTime(v string) *OralEvaluationStatisticsCallsCountRequest {
	s.StartTime = &v
	return s
}

type OralEvaluationStatisticsCallsCountResponse struct {
	// This parameter is required.
	ProjectData *OralEvaluationStatisticsCallsCountResponseProjectData `json:"projectData,omitempty" xml:"projectData,omitempty" type:"Struct"`
	// This parameter is required.
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s OralEvaluationStatisticsCallsCountResponse) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsCallsCountResponse) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsCallsCountResponse) SetProjectData(v *OralEvaluationStatisticsCallsCountResponseProjectData) *OralEvaluationStatisticsCallsCountResponse {
	s.ProjectData = v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponse) SetProjectId(v string) *OralEvaluationStatisticsCallsCountResponse {
	s.ProjectId = &v
	return s
}

type OralEvaluationStatisticsCallsCountResponseProjectData struct {
	ApplicationData []*OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData `json:"ApplicationData,omitempty" xml:"ApplicationData,omitempty" type:"Repeated"`
	// This parameter is required.
	ApplicationInternalId *string `json:"applicationInternalId,omitempty" xml:"applicationInternalId,omitempty"`
}

func (s OralEvaluationStatisticsCallsCountResponseProjectData) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsCallsCountResponseProjectData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectData) SetApplicationData(v []*OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) *OralEvaluationStatisticsCallsCountResponseProjectData {
	s.ApplicationData = v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectData) SetApplicationInternalId(v string) *OralEvaluationStatisticsCallsCountResponseProjectData {
	s.ApplicationInternalId = &v
	return s
}

type OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData struct {
	Data []*OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// This parameter is required.
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
}

func (s OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) SetData(v []*OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData {
	s.Data = v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData) SetApplicationAccessId(v string) *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationData {
	s.ApplicationAccessId = &v
	return s
}

type OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData struct {
	// This parameter is required.
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) SetCount(v int32) *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData {
	s.Count = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData) SetName(v string) *OralEvaluationStatisticsCallsCountResponseProjectDataApplicationDataData {
	s.Name = &v
	return s
}

type OralEvaluationStatisticsConcurrentCountRequest struct {
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	EndTime             *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Granularity         *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	ProjectId           *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	StartTime           *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s OralEvaluationStatisticsConcurrentCountRequest) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsConcurrentCountRequest) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) SetApplicationAccessId(v string) *OralEvaluationStatisticsConcurrentCountRequest {
	s.ApplicationAccessId = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) SetEndTime(v string) *OralEvaluationStatisticsConcurrentCountRequest {
	s.EndTime = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) SetGranularity(v string) *OralEvaluationStatisticsConcurrentCountRequest {
	s.Granularity = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) SetProjectId(v string) *OralEvaluationStatisticsConcurrentCountRequest {
	s.ProjectId = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) SetStartTime(v string) *OralEvaluationStatisticsConcurrentCountRequest {
	s.StartTime = &v
	return s
}

type OralEvaluationStatisticsConcurrentCountResponse struct {
	// This parameter is required.
	ProjectData *OralEvaluationStatisticsConcurrentCountResponseProjectData `json:"projectData,omitempty" xml:"projectData,omitempty" type:"Struct"`
	// This parameter is required.
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s OralEvaluationStatisticsConcurrentCountResponse) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsConcurrentCountResponse) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsConcurrentCountResponse) SetProjectData(v *OralEvaluationStatisticsConcurrentCountResponseProjectData) *OralEvaluationStatisticsConcurrentCountResponse {
	s.ProjectData = v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponse) SetProjectId(v string) *OralEvaluationStatisticsConcurrentCountResponse {
	s.ProjectId = &v
	return s
}

type OralEvaluationStatisticsConcurrentCountResponseProjectData struct {
	ApplicationData []*OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData `json:"ApplicationData,omitempty" xml:"ApplicationData,omitempty" type:"Repeated"`
	// This parameter is required.
	ApplicationInternalId *string `json:"applicationInternalId,omitempty" xml:"applicationInternalId,omitempty"`
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectData) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectData) SetApplicationData(v []*OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) *OralEvaluationStatisticsConcurrentCountResponseProjectData {
	s.ApplicationData = v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectData) SetApplicationInternalId(v string) *OralEvaluationStatisticsConcurrentCountResponseProjectData {
	s.ApplicationInternalId = &v
	return s
}

type OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData struct {
	Data []*OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// This parameter is required.
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) SetData(v []*OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData {
	s.Data = v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData) SetApplicationAccessId(v string) *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationData {
	s.ApplicationAccessId = &v
	return s
}

type OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData struct {
	// This parameter is required.
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) SetCount(v int32) *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData {
	s.Count = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData) SetName(v string) *OralEvaluationStatisticsConcurrentCountResponseProjectDataApplicationDataData {
	s.Name = &v
	return s
}

type OralEvaluationStatisticsErrorCountRequest struct {
	ApplicationAccessId *string   `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	EndTime             *string   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ErrorCode           []*string `json:"errorCode,omitempty" xml:"errorCode,omitempty" type:"Repeated"`
	Granularity         *string   `json:"granularity,omitempty" xml:"granularity,omitempty"`
	ProjectId           *string   `json:"projectId,omitempty" xml:"projectId,omitempty"`
	StartTime           *string   `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountRequest) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountRequest) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetApplicationAccessId(v string) *OralEvaluationStatisticsErrorCountRequest {
	s.ApplicationAccessId = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetEndTime(v string) *OralEvaluationStatisticsErrorCountRequest {
	s.EndTime = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetErrorCode(v []*string) *OralEvaluationStatisticsErrorCountRequest {
	s.ErrorCode = v
	return s
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetGranularity(v string) *OralEvaluationStatisticsErrorCountRequest {
	s.Granularity = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetProjectId(v string) *OralEvaluationStatisticsErrorCountRequest {
	s.ProjectId = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetStartTime(v string) *OralEvaluationStatisticsErrorCountRequest {
	s.StartTime = &v
	return s
}

type OralEvaluationStatisticsErrorCountResponse struct {
	ProjectData *OralEvaluationStatisticsErrorCountResponseProjectData `json:"ProjectData,omitempty" xml:"ProjectData,omitempty" type:"Struct"`
	// This parameter is required.
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountResponse) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountResponse) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountResponse) SetProjectData(v *OralEvaluationStatisticsErrorCountResponseProjectData) *OralEvaluationStatisticsErrorCountResponse {
	s.ProjectData = v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponse) SetProjectId(v string) *OralEvaluationStatisticsErrorCountResponse {
	s.ProjectId = &v
	return s
}

type OralEvaluationStatisticsErrorCountResponseProjectData struct {
	ApplicationData []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData `json:"ApplicationData,omitempty" xml:"ApplicationData,omitempty" type:"Repeated"`
	// This parameter is required.
	ApplicationInternalId *string `json:"applicationInternalId,omitempty" xml:"applicationInternalId,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountResponseProjectData) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountResponseProjectData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectData) SetApplicationData(v []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) *OralEvaluationStatisticsErrorCountResponseProjectData {
	s.ApplicationData = v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectData) SetApplicationInternalId(v string) *OralEvaluationStatisticsErrorCountResponseProjectData {
	s.ApplicationInternalId = &v
	return s
}

type OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData struct {
	Data []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// This parameter is required.
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) SetData(v []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData {
	s.Data = v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData) SetApplicationAccessId(v string) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationData {
	s.ApplicationAccessId = &v
	return s
}

type OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData struct {
	Data         []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *string                                                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                                                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) SetData(v []*OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData {
	s.Data = v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) SetErrorCode(v string) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData {
	s.ErrorCode = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData) SetErrorMessage(v string) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataData {
	s.ErrorMessage = &v
	return s
}

type OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData struct {
	// This parameter is required.
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) String() string {
	return tea.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) SetCount(v int32) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData {
	s.Count = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData) SetName(v string) *OralEvaluationStatisticsErrorCountResponseProjectDataApplicationDataDataData {
	s.Name = &v
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

type AITeacherExpansionPracticeTaskGenerateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 13
	Grade        *string   `json:"grade,omitempty" xml:"grade,omitempty"`
	KeySentences []*string `json:"keySentences,omitempty" xml:"keySentences,omitempty" type:"Repeated"`
	KeyWords     []*string `json:"keyWords,omitempty" xml:"keyWords,omitempty" type:"Repeated"`
	// example:
	//
	// Understanding unique professions such as dog walkers, hotel test sleepers, and food tasters, including their job responsibilities and the benefits or challenges associated with each role.
	LearningObject *string `json:"learningObject,omitempty" xml:"learningObject,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dog walker Dog walking, as a profession, originated in the US. Some may think that it\\"s a perfect job, because dog walkers won\\"t be imprisoned in an office. But it\\"s actually manual labour. At their busiest, dog walkers may have more than ten dogs to take care of in a day. Hotel test sleeper A hotel test sleeper, as the name suggests, has to write expert reviews about the facilities, locations, prices, dining and other services of hotels, in order to provide evaluations and guides for travelers. Hotel test sleepers don\\"t need to punch in for work and they get about ten thousand yuan as income every month. What a comfortable job! Food taster In ancient times, a food taster was a person who tasted foods (or drinks) to be served to someone else, to confirm that it was safe to eat. But now, those working as food tasters just get to taste various new foods and drinks aimed at specific regions across the world. They then give their opinions on these products to the companies and suggest improvements.
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	Textbook    *string `json:"textbook,omitempty" xml:"textbook,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// talk about your dream job.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6440xxxxxxxxxx5fafc98c421
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateRequest) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetGrade(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.Grade = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetKeySentences(v []*string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.KeySentences = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetKeyWords(v []*string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.KeyWords = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetLearningObject(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.LearningObject = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetTextContent(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.TextContent = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetTextbook(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.Textbook = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetTopic(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.Topic = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateRequest) SetUserId(v string) *AITeacherExpansionPracticeTaskGenerateRequest {
	s.UserId = &v
	return s
}

type AITeacherExpansionPracticeTaskGenerateResponseBody struct {
	Data *AITeacherExpansionPracticeTaskGenerateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
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

func (s AITeacherExpansionPracticeTaskGenerateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetData(v *AITeacherExpansionPracticeTaskGenerateResponseBodyData) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.Data = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetErrCode(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetErrMessage(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetHttpStatusCode(v int32) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetRequestId(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetSuccess(v bool) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.Success = &v
	return s
}

type AITeacherExpansionPracticeTaskGenerateResponseBodyData struct {
	// example:
	//
	// In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
	BackgroundDescription *string                                                        `json:"backgroundDescription,omitempty" xml:"backgroundDescription,omitempty"`
	RoleSet               *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet `json:"roleSet,omitempty" xml:"roleSet,omitempty" type:"Struct"`
	// example:
	//
	// Hey Jamie, do you know what a travel blogger does?
	StartSentence *string                                                              `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	TaskContent   []*AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent `json:"taskContent,omitempty" xml:"taskContent,omitempty" type:"Repeated"`
	// example:
	//
	// textbook_dialogue
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetBackgroundDescription(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.BackgroundDescription = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetRoleSet(v *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.RoleSet = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetStartSentence(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.StartSentence = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetTaskContent(v []*AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.TaskContent = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetTaskType(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.TaskType = &v
	return s
}

type AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet struct {
	// example:
	//
	// Alex
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// Jamie
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) SetAssistant(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet {
	s.Assistant = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) SetUser(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet {
	s.User = &v
	return s
}

type AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent struct {
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) SetAssistant(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent {
	s.Assistant = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) SetUser(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent {
	s.User = &v
	return s
}

type AITeacherExpansionPracticeTaskGenerateResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AITeacherExpansionPracticeTaskGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponse) String() string {
	return tea.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponse) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) SetHeaders(v map[string]*string) *AITeacherExpansionPracticeTaskGenerateResponse {
	s.Headers = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) SetStatusCode(v int32) *AITeacherExpansionPracticeTaskGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponse) SetBody(v *AITeacherExpansionPracticeTaskGenerateResponseBody) *AITeacherExpansionPracticeTaskGenerateResponse {
	s.Body = v
	return s
}

type AITeacherSyncPracticeTaskGenerateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 13
	Grade        *string   `json:"grade,omitempty" xml:"grade,omitempty"`
	KeySentences []*string `json:"keySentences,omitempty" xml:"keySentences,omitempty" type:"Repeated"`
	KeyWords     []*string `json:"keyWords,omitempty" xml:"keyWords,omitempty" type:"Repeated"`
	// example:
	//
	// Understanding unique professions such as dog walkers, hotel test sleepers, and food tasters, including their job responsibilities and the benefits or challenges associated with each role.
	LearningObject *string `json:"learningObject,omitempty" xml:"learningObject,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dog walker Dog walking, as a profession, originated in the US. Some may think that it\\"s a perfect job, because dog walkers won\\"t be imprisoned in an office. But it\\"s actually manual labour. At their busiest, dog walkers may have more than ten dogs to take care of in a day. Hotel test sleeper A hotel test sleeper, as the name suggests, has to write expert reviews about the facilities, locations, prices, dining and other services of hotels, in order to provide evaluations and guides for travelers. Hotel test sleepers don\\"t need to punch in for work and they get about ten thousand yuan as income every month. What a comfortable job! Food taster In ancient times, a food taster was a person who tasted foods (or drinks) to be served to someone else, to confirm that it was safe to eat. But now, those working as food tasters just get to taste various new foods and drinks aimed at specific regions across the world. They then give their opinions on these products to the companies and suggest improvements.
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	Textbook    *string `json:"textbook,omitempty" xml:"textbook,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// talk about your dream job.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6440xxxxxxxxxx5fafc98c421
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateRequest) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetGrade(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.Grade = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetKeySentences(v []*string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.KeySentences = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetKeyWords(v []*string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.KeyWords = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetLearningObject(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.LearningObject = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetTextContent(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.TextContent = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetTextbook(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.Textbook = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetTopic(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.Topic = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateRequest) SetUserId(v string) *AITeacherSyncPracticeTaskGenerateRequest {
	s.UserId = &v
	return s
}

type AITeacherSyncPracticeTaskGenerateResponseBody struct {
	Data *AITeacherSyncPracticeTaskGenerateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
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

func (s AITeacherSyncPracticeTaskGenerateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetData(v *AITeacherSyncPracticeTaskGenerateResponseBodyData) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.Data = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetErrCode(v string) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetErrMessage(v string) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetHttpStatusCode(v int32) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetRequestId(v string) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetSuccess(v bool) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.Success = &v
	return s
}

type AITeacherSyncPracticeTaskGenerateResponseBodyData struct {
	TaskContent []*AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent `json:"taskContent,omitempty" xml:"taskContent,omitempty" type:"Repeated"`
	// example:
	//
	// textbook_question_answering
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyData) SetTaskContent(v []*AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) *AITeacherSyncPracticeTaskGenerateResponseBodyData {
	s.TaskContent = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyData) SetTaskType(v string) *AITeacherSyncPracticeTaskGenerateResponseBodyData {
	s.TaskType = &v
	return s
}

type AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent struct {
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) String() string {
	return tea.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) SetAssistant(v string) *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent {
	s.Assistant = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) SetUser(v string) *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent {
	s.User = &v
	return s
}

type AITeacherSyncPracticeTaskGenerateResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AITeacherSyncPracticeTaskGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateResponse) String() string {
	return tea.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponse) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) SetHeaders(v map[string]*string) *AITeacherSyncPracticeTaskGenerateResponse {
	s.Headers = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) SetStatusCode(v int32) *AITeacherSyncPracticeTaskGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponse) SetBody(v *AITeacherSyncPracticeTaskGenerateResponseBody) *AITeacherSyncPracticeTaskGenerateResponse {
	s.Body = v
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
	Data []*AliyunConsoleServiceInfoDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody) SetData(v []*AliyunConsoleServiceInfoDTO) *AliyunConsoleOpenApiQueryAliyunConsoleServiceListResponseBody {
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

type CountOralEvaluationStatisticsCallsRequest struct {
	Body *OralEvaluationStatisticsCallsCountRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsCallsRequest) String() string {
	return tea.Prettify(s)
}

func (s CountOralEvaluationStatisticsCallsRequest) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsCallsRequest) SetBody(v *OralEvaluationStatisticsCallsCountRequest) *CountOralEvaluationStatisticsCallsRequest {
	s.Body = v
	return s
}

type CountOralEvaluationStatisticsCallsResponseBody struct {
	// example:
	//
	// []
	Data []*OralEvaluationStatisticsCallsCountResponse `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s CountOralEvaluationStatisticsCallsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountOralEvaluationStatisticsCallsResponseBody) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetData(v []*OralEvaluationStatisticsCallsCountResponse) *CountOralEvaluationStatisticsCallsResponseBody {
	s.Data = v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetErrCode(v string) *CountOralEvaluationStatisticsCallsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetErrMessage(v string) *CountOralEvaluationStatisticsCallsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetHttpStatusCode(v int32) *CountOralEvaluationStatisticsCallsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetRequestId(v string) *CountOralEvaluationStatisticsCallsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponseBody) SetSuccess(v bool) *CountOralEvaluationStatisticsCallsResponseBody {
	s.Success = &v
	return s
}

type CountOralEvaluationStatisticsCallsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountOralEvaluationStatisticsCallsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsCallsResponse) String() string {
	return tea.Prettify(s)
}

func (s CountOralEvaluationStatisticsCallsResponse) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsCallsResponse) SetHeaders(v map[string]*string) *CountOralEvaluationStatisticsCallsResponse {
	s.Headers = v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponse) SetStatusCode(v int32) *CountOralEvaluationStatisticsCallsResponse {
	s.StatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsCallsResponse) SetBody(v *CountOralEvaluationStatisticsCallsResponseBody) *CountOralEvaluationStatisticsCallsResponse {
	s.Body = v
	return s
}

type CountOralEvaluationStatisticsConcurrentRequest struct {
	Body *OralEvaluationStatisticsConcurrentCountRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsConcurrentRequest) String() string {
	return tea.Prettify(s)
}

func (s CountOralEvaluationStatisticsConcurrentRequest) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsConcurrentRequest) SetBody(v *OralEvaluationStatisticsConcurrentCountRequest) *CountOralEvaluationStatisticsConcurrentRequest {
	s.Body = v
	return s
}

type CountOralEvaluationStatisticsConcurrentResponseBody struct {
	// example:
	//
	// []
	Data []*OralEvaluationStatisticsConcurrentCountResponse `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s CountOralEvaluationStatisticsConcurrentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountOralEvaluationStatisticsConcurrentResponseBody) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetData(v []*OralEvaluationStatisticsConcurrentCountResponse) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.Data = v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetErrCode(v string) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetErrMessage(v string) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetHttpStatusCode(v int32) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetRequestId(v string) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponseBody) SetSuccess(v bool) *CountOralEvaluationStatisticsConcurrentResponseBody {
	s.Success = &v
	return s
}

type CountOralEvaluationStatisticsConcurrentResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountOralEvaluationStatisticsConcurrentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsConcurrentResponse) String() string {
	return tea.Prettify(s)
}

func (s CountOralEvaluationStatisticsConcurrentResponse) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsConcurrentResponse) SetHeaders(v map[string]*string) *CountOralEvaluationStatisticsConcurrentResponse {
	s.Headers = v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponse) SetStatusCode(v int32) *CountOralEvaluationStatisticsConcurrentResponse {
	s.StatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsConcurrentResponse) SetBody(v *CountOralEvaluationStatisticsConcurrentResponseBody) *CountOralEvaluationStatisticsConcurrentResponse {
	s.Body = v
	return s
}

type CountOralEvaluationStatisticsErrorRequest struct {
	Body *OralEvaluationStatisticsErrorCountRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsErrorRequest) String() string {
	return tea.Prettify(s)
}

func (s CountOralEvaluationStatisticsErrorRequest) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsErrorRequest) SetBody(v *OralEvaluationStatisticsErrorCountRequest) *CountOralEvaluationStatisticsErrorRequest {
	s.Body = v
	return s
}

type CountOralEvaluationStatisticsErrorResponseBody struct {
	// example:
	//
	// []
	Data []*OralEvaluationStatisticsErrorCountResponse `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s CountOralEvaluationStatisticsErrorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountOralEvaluationStatisticsErrorResponseBody) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetData(v []*OralEvaluationStatisticsErrorCountResponse) *CountOralEvaluationStatisticsErrorResponseBody {
	s.Data = v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetErrCode(v string) *CountOralEvaluationStatisticsErrorResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetErrMessage(v string) *CountOralEvaluationStatisticsErrorResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetHttpStatusCode(v int32) *CountOralEvaluationStatisticsErrorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetRequestId(v string) *CountOralEvaluationStatisticsErrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponseBody) SetSuccess(v bool) *CountOralEvaluationStatisticsErrorResponseBody {
	s.Success = &v
	return s
}

type CountOralEvaluationStatisticsErrorResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountOralEvaluationStatisticsErrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOralEvaluationStatisticsErrorResponse) String() string {
	return tea.Prettify(s)
}

func (s CountOralEvaluationStatisticsErrorResponse) GoString() string {
	return s.String()
}

func (s *CountOralEvaluationStatisticsErrorResponse) SetHeaders(v map[string]*string) *CountOralEvaluationStatisticsErrorResponse {
	s.Headers = v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponse) SetStatusCode(v int32) *CountOralEvaluationStatisticsErrorResponse {
	s.StatusCode = &v
	return s
}

func (s *CountOralEvaluationStatisticsErrorResponse) SetBody(v *CountOralEvaluationStatisticsErrorResponseBody) *CountOralEvaluationStatisticsErrorResponse {
	s.Body = v
	return s
}

type CreateAccessWarrantRequest struct {
	// example:
	//
	// a123
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// e32fac43df0b0b0be32fac43df0b0b0b
	RequestSign *string `json:"requestSign,omitempty" xml:"requestSign,omitempty"`
	// example:
	//
	// 1701000000
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// 110.25.23.12
	UserClientIp *string `json:"userClientIp,omitempty" xml:"userClientIp,omitempty"`
	// example:
	//
	// sn123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 7200
	WarrantAvailable *int32 `json:"warrantAvailable,omitempty" xml:"warrantAvailable,omitempty"`
}

func (s CreateAccessWarrantRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessWarrantRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessWarrantRequest) SetAppId(v string) *CreateAccessWarrantRequest {
	s.AppId = &v
	return s
}

func (s *CreateAccessWarrantRequest) SetRequestSign(v string) *CreateAccessWarrantRequest {
	s.RequestSign = &v
	return s
}

func (s *CreateAccessWarrantRequest) SetTimestamp(v string) *CreateAccessWarrantRequest {
	s.Timestamp = &v
	return s
}

func (s *CreateAccessWarrantRequest) SetUserClientIp(v string) *CreateAccessWarrantRequest {
	s.UserClientIp = &v
	return s
}

func (s *CreateAccessWarrantRequest) SetUserId(v string) *CreateAccessWarrantRequest {
	s.UserId = &v
	return s
}

func (s *CreateAccessWarrantRequest) SetWarrantAvailable(v int32) *CreateAccessWarrantRequest {
	s.WarrantAvailable = &v
	return s
}

type CreateAccessWarrantResponseBody struct {
	// example:
	//
	// []
	Data *CreateAccessWarrantResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateAccessWarrantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessWarrantResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessWarrantResponseBody) SetData(v *CreateAccessWarrantResponseBodyData) *CreateAccessWarrantResponseBody {
	s.Data = v
	return s
}

func (s *CreateAccessWarrantResponseBody) SetErrCode(v string) *CreateAccessWarrantResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateAccessWarrantResponseBody) SetErrMessage(v string) *CreateAccessWarrantResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateAccessWarrantResponseBody) SetHttpStatusCode(v int32) *CreateAccessWarrantResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAccessWarrantResponseBody) SetRequestId(v string) *CreateAccessWarrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessWarrantResponseBody) SetSuccess(v bool) *CreateAccessWarrantResponseBody {
	s.Success = &v
	return s
}

type CreateAccessWarrantResponseBodyData struct {
	// example:
	//
	// 1234567890
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// ex2xxxxxxxx
	AccessWarrantId *string `json:"AccessWarrantId,omitempty" xml:"AccessWarrantId,omitempty"`
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"ApplicationAccessId,omitempty" xml:"ApplicationAccessId,omitempty"`
	// example:
	//
	// 1672531200
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1672531200
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 1234567890
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateAccessWarrantResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessWarrantResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAccessWarrantResponseBodyData) SetAccessToken(v string) *CreateAccessWarrantResponseBodyData {
	s.AccessToken = &v
	return s
}

func (s *CreateAccessWarrantResponseBodyData) SetAccessWarrantId(v string) *CreateAccessWarrantResponseBodyData {
	s.AccessWarrantId = &v
	return s
}

func (s *CreateAccessWarrantResponseBodyData) SetApplicationAccessId(v string) *CreateAccessWarrantResponseBodyData {
	s.ApplicationAccessId = &v
	return s
}

func (s *CreateAccessWarrantResponseBodyData) SetCreateTime(v string) *CreateAccessWarrantResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateAccessWarrantResponseBodyData) SetExpireTime(v string) *CreateAccessWarrantResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *CreateAccessWarrantResponseBodyData) SetUserId(v string) *CreateAccessWarrantResponseBodyData {
	s.UserId = &v
	return s
}

type CreateAccessWarrantResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessWarrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessWarrantResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessWarrantResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessWarrantResponse) SetHeaders(v map[string]*string) *CreateAccessWarrantResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessWarrantResponse) SetStatusCode(v int32) *CreateAccessWarrantResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessWarrantResponse) SetBody(v *CreateAccessWarrantResponseBody) *CreateAccessWarrantResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	// example:
	//
	// MyProject
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// example:
	//
	// online_oral_evaluation_post_paid_call_count
	ProjectType *string `json:"projectType,omitempty" xml:"projectType,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetProjectType(v string) *CreateProjectRequest {
	s.ProjectType = &v
	return s
}

type CreateProjectResponseBody struct {
	// example:
	//
	// []
	Data *CreateProjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetData(v *CreateProjectResponseBodyData) *CreateProjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateProjectResponseBody) SetErrCode(v string) *CreateProjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateProjectResponseBody) SetErrMessage(v string) *CreateProjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateProjectResponseBody) SetHttpStatusCode(v int32) *CreateProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v bool) *CreateProjectResponseBody {
	s.Success = &v
	return s
}

type CreateProjectResponseBodyData struct {
	// example:
	//
	// 2023-02-15T09:17:39Z
	CreateTime  *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectApps []*CreateProjectResponseBodyDataProjectApps `json:"ProjectApps,omitempty" xml:"ProjectApps,omitempty" type:"Repeated"`
	// example:
	//
	// 124187
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// MyProject
	ProjectName *string                                    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectSDK  []*CreateProjectResponseBodyDataProjectSDK `json:"ProjectSDK,omitempty" xml:"ProjectSDK,omitempty" type:"Repeated"`
	// example:
	//
	// WebApplication
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
}

func (s CreateProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyData) SetCreateTime(v string) *CreateProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetProjectApps(v []*CreateProjectResponseBodyDataProjectApps) *CreateProjectResponseBodyData {
	s.ProjectApps = v
	return s
}

func (s *CreateProjectResponseBodyData) SetProjectId(v string) *CreateProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetProjectName(v string) *CreateProjectResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetProjectSDK(v []*CreateProjectResponseBodyDataProjectSDK) *CreateProjectResponseBodyData {
	s.ProjectSDK = v
	return s
}

func (s *CreateProjectResponseBodyData) SetProjectType(v string) *CreateProjectResponseBodyData {
	s.ProjectType = &v
	return s
}

type CreateProjectResponseBodyDataProjectApps struct {
	ApplicationAccessIds []*CreateProjectResponseBodyDataProjectAppsApplicationAccessIds `json:"ApplicationAccessIds,omitempty" xml:"ApplicationAccessIds,omitempty" type:"Repeated"`
	// example:
	//
	// 4867
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 4910
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateProjectResponseBodyDataProjectApps) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyDataProjectApps) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyDataProjectApps) SetApplicationAccessIds(v []*CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) *CreateProjectResponseBodyDataProjectApps {
	s.ApplicationAccessIds = v
	return s
}

func (s *CreateProjectResponseBodyDataProjectApps) SetId(v string) *CreateProjectResponseBodyDataProjectApps {
	s.Id = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectApps) SetProjectId(v string) *CreateProjectResponseBodyDataProjectApps {
	s.ProjectId = &v
	return s
}

type CreateProjectResponseBodyDataProjectAppsApplicationAccessIds struct {
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// MyAppSecret
	ApplicationAccessSecret *string `json:"applicationAccessSecret,omitempty" xml:"applicationAccessSecret,omitempty"`
}

func (s CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessId(v string) *CreateProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessId = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessSecret(v string) *CreateProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessSecret = &v
	return s
}

type CreateProjectResponseBodyDataProjectSDK struct {
	// example:
	//
	// 2023-02-15T09:17:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// http://demo.com/demo
	DemoUrl    *string `json:"DemoUrl,omitempty" xml:"DemoUrl,omitempty"`
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// example:
	//
	// C++
	DevelopLanguage *string `json:"DevelopLanguage,omitempty" xml:"DevelopLanguage,omitempty"`
	// example:
	//
	// http://demo.com/doc
	DocUrl *string `json:"DocUrl,omitempty" xml:"DocUrl,omitempty"`
	// example:
	//
	// C SDK
	SdkName *string `json:"SdkName,omitempty" xml:"SdkName,omitempty"`
	// example:
	//
	// http://demo.com/sdk.zip
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	// example:
	//
	// 4.12.8
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
}

func (s CreateProjectResponseBodyDataProjectSDK) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyDataProjectSDK) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetCreateTime(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.CreateTime = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetDemoUrl(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.DemoUrl = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetDeployMode(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.DeployMode = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetDevelopLanguage(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.DevelopLanguage = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetDocUrl(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.DocUrl = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetSdkName(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.SdkName = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetSdkUrl(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.SdkUrl = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetSdkVersion(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.SdkVersion = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest struct {
	EssayOutline *string `json:"essayOutline,omitempty" xml:"essayOutline,omitempty"`
	// This parameter is required.
	EssayRequirements *string `json:"essayRequirements,omitempty" xml:"essayRequirements,omitempty"`
	// This parameter is required.
	EssayTopic *string `json:"essayTopic,omitempty" xml:"essayTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// essay
	EssayType *string `json:"essayType,omitempty" xml:"essayType,omitempty"`
	// example:
	//
	// 100
	EssayWordCount *int64 `json:"essayWordCount,omitempty" xml:"essayWordCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	Grade *int64 `json:"grade,omitempty" xml:"grade,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// streaming
	ResponseMode *string `json:"responseMode,omitempty" xml:"responseMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxxxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetEssayOutline(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
	s.EssayOutline = &v
	return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetEssayRequirements(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
	s.EssayRequirements = &v
	return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetEssayTopic(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
	s.EssayTopic = &v
	return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetEssayType(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
	s.EssayType = &v
	return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetEssayWordCount(v int64) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
	s.EssayWordCount = &v
	return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetGrade(v int64) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
	s.Grade = &v
	return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetResponseMode(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
	s.ResponseMode = &v
	return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) SetUserId(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// message
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) SetContent(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) SetEvent(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody {
	s.Event = &v
	return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) SetRequestId(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse struct {
	Headers    map[string]*string                                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) SetStatusCode(v int32) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse) SetBody(v *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest struct {
	// example:
	//
	// Title: The Importance of Reading
	//
	// I. Introduction
	//
	// II. Body
	//
	// III. Conclusion
	EssayOutline *string `json:"essayOutline,omitempty" xml:"essayOutline,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// No less than 100 words
	EssayRequirements *string `json:"essayRequirements,omitempty" xml:"essayRequirements,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Discuss what to eat
	EssayTopic *string `json:"essayTopic,omitempty" xml:"essayTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// essay
	EssayType *string `json:"essayType,omitempty" xml:"essayType,omitempty"`
	// example:
	//
	// 100
	EssayWordCount *int64 `json:"essayWordCount,omitempty" xml:"essayWordCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	Grade *int64 `json:"grade,omitempty" xml:"grade,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// streaming
	ResponseMode *string `json:"responseMode,omitempty" xml:"responseMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetEssayOutline(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
	s.EssayOutline = &v
	return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetEssayRequirements(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
	s.EssayRequirements = &v
	return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetEssayTopic(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
	s.EssayTopic = &v
	return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetEssayType(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
	s.EssayType = &v
	return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetEssayWordCount(v int64) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
	s.EssayWordCount = &v
	return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetGrade(v int64) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
	s.Grade = &v
	return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetResponseMode(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
	s.ResponseMode = &v
	return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) SetUserId(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody struct {
	// example:
	//
	// hi
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// message
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) SetContent(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) SetEvent(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody {
	s.Event = &v
	return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) SetRequestId(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse struct {
	Headers    map[string]*string                                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) SetStatusCode(v int32) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse) SetBody(v *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherEnglishParaphraseChatMessageRequest struct {
	// example:
	//
	// 6788e0b475a4631ffc626722
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// How much is this?
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 3
	Grade *int64 `json:"grade,omitempty" xml:"grade,omitempty"`
	// example:
	//
	// xxxxxxxxx
	QuestionId *string `json:"questionId,omitempty" xml:"questionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// How to inquire about the price
	QuestionInfo *string `json:"questionInfo,omitempty" xml:"questionInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sreaming
	ResponseMode *string `json:"responseMode,omitempty" xml:"responseMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// How much is this?
	UserAnswer *string `json:"userAnswer,omitempty" xml:"userAnswer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetChatId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
	s.ChatId = &v
	return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetContent(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetGrade(v int64) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
	s.Grade = &v
	return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetQuestionId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
	s.QuestionId = &v
	return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetQuestionInfo(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
	s.QuestionInfo = &v
	return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetResponseMode(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
	s.ResponseMode = &v
	return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetUserAnswer(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
	s.UserAnswer = &v
	return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetUserId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherEnglishParaphraseChatMessageResponseBody struct {
	// example:
	//
	// how
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// message
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) SetContent(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) SetEvent(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody {
	s.Event = &v
	return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) SetRequestId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteAITeacherEnglishParaphraseChatMessageResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherEnglishParaphraseChatMessageResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) SetStatusCode(v int32) *ExecuteAITeacherEnglishParaphraseChatMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponse) SetBody(v *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) *ExecuteAITeacherEnglishParaphraseChatMessageResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherExpansionDialogueRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
	Background *string `json:"background,omitempty" xml:"background,omitempty"`
	// This parameter is required.
	DialogueTasks []*ExecuteAITeacherExpansionDialogueRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string                                            `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	Records      []*ExecuteAITeacherExpansionDialogueRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	RoleInfo *ExecuteAITeacherExpansionDialogueRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
	// example:
	//
	// Hello Lily, could you please come to the kitchen for a moment?
	StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Let\\"s talk about traffic rules.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetBackground(v string) *ExecuteAITeacherExpansionDialogueRequest {
	s.Background = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueRequest {
	s.DialogueTasks = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetLanguageCode(v string) *ExecuteAITeacherExpansionDialogueRequest {
	s.LanguageCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetRecords(v []*ExecuteAITeacherExpansionDialogueRequestRecords) *ExecuteAITeacherExpansionDialogueRequest {
	s.Records = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetRoleInfo(v *ExecuteAITeacherExpansionDialogueRequestRoleInfo) *ExecuteAITeacherExpansionDialogueRequest {
	s.RoleInfo = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueRequest {
	s.StartSentence = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetTopic(v string) *ExecuteAITeacherExpansionDialogueRequest {
	s.Topic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetUserId(v string) *ExecuteAITeacherExpansionDialogueRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequestRecords) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetContent(v string) *ExecuteAITeacherExpansionDialogueRequestRecords {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRequestRecords {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetRole(v string) *ExecuteAITeacherExpansionDialogueRequestRecords {
	s.Role = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRequestRoleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// Alex
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jamie
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRequestRoleInfo) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequestRoleInfo) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequestRoleInfo) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRequestRoleInfo {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRoleInfo) SetUser(v string) *ExecuteAITeacherExpansionDialogueRequestRoleInfo {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherExpansionDialogueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherExpansionDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetData(v *ExecuteAITeacherExpansionDialogueResponseBodyData) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetErrCode(v string) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetRequestId(v string) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBody) SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherExpansionDialogueResponseBodyData struct {
	// example:
	//
	// 1
	ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
	// example:
	//
	// 1
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
	// example:
	//
	// true
	IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
	// example:
	//
	// true
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// true
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// example:
	//
	// 2
	QuestionIndex *int32 `json:"questionIndex,omitempty" xml:"questionIndex,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetChineseResult(v string) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.ChineseResult = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetEnglishResult(v string) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.EnglishResult = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetIsFinish(v bool) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponseBodyData) SetQuestionIndex(v int32) *ExecuteAITeacherExpansionDialogueResponseBodyData {
	s.QuestionIndex = &v
	return s
}

type ExecuteAITeacherExpansionDialogueResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherExpansionDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponse) SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueResponse) SetBody(v *ExecuteAITeacherExpansionDialogueResponseBody) *ExecuteAITeacherExpansionDialogueResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
	Background *string `json:"background,omitempty" xml:"background,omitempty"`
	// This parameter is required.
	DialogueTasks []*ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	// This parameter is required.
	Records []*ExecuteAITeacherExpansionDialogueRefineRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	RoleInfo *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
	// example:
	//
	// Hello Lily, could you please come to the kitchen for a moment?
	StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// talk about your dream job.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetBackground(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.Background = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.DialogueTasks = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetLanguageCode(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.LanguageCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetRecords(v []*ExecuteAITeacherExpansionDialogueRefineRequestRecords) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.Records = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetRoleInfo(v *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.RoleInfo = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.StartSentence = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetTopic(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.Topic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetUserId(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestRecords) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetContent(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetRole(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
	s.Role = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// Jane, a caring mother
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Lily, a friendly student
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) SetUser(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherExpansionDialogueRefineResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherExpansionDialogueRefineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetData(v *ExecuteAITeacherExpansionDialogueRefineResponseBodyData) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetErrCode(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetRequestId(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBody) SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueRefineResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineResponseBodyData struct {
	// example:
	//
	// Yes, I\\"ll be right there.
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponseBodyData) SetResult(v string) *ExecuteAITeacherExpansionDialogueRefineResponseBodyData {
	s.Result = &v
	return s
}

type ExecuteAITeacherExpansionDialogueRefineResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherExpansionDialogueRefineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueRefineResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueRefineResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineResponse) SetBody(v *ExecuteAITeacherExpansionDialogueRefineResponseBody) *ExecuteAITeacherExpansionDialogueRefineResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// In this dialogue, you will be playing the role of Lily, a young girl. I will be Jane, Lily\\"s mother. We are in the kitchen, where I am preparing dinner. I am asking you about your food preferences, specifically if you like meat, fish, and milk. You like meat and milk, but you don\\"t like fish because of its smell. I explain to you the nutritional benefits of these foods and suggest alternatives for the ones you don\\"t like. Finally, I invite you to start eating.
	Background *string `json:"background,omitempty" xml:"background,omitempty"`
	// This parameter is required.
	DialogueTasks []*ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	Records       []*ExecuteAITeacherExpansionDialogueTranslateRequestRecords       `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	RoleInfo *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
	// example:
	//
	// Hello Lily, could you please come to the kitchen for a moment?
	StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// talk about food.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetBackground(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.Background = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.DialogueTasks = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetRecords(v []*ExecuteAITeacherExpansionDialogueTranslateRequestRecords) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.Records = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetRoleInfo(v *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.RoleInfo = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.StartSentence = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetTopic(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.Topic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetUserId(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRecords) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetContent(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetRole(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
	s.Role = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// Jane, a caring mother
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Lily, a friendly student
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) SetUser(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo {
	s.User = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetData(v *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetErrCode(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetErrMessage(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetRequestId(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBody) SetSuccess(v bool) *ExecuteAITeacherExpansionDialogueTranslateResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateResponseBodyData struct {
	// example:
	//
	// 太好了，谢谢你过来，莉莉。你喜欢吃肉吗？
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData) SetResult(v string) *ExecuteAITeacherExpansionDialogueTranslateResponseBodyData {
	s.Result = &v
	return s
}

type ExecuteAITeacherExpansionDialogueTranslateResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherExpansionDialogueTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherExpansionDialogueTranslateResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) SetStatusCode(v int32) *ExecuteAITeacherExpansionDialogueTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateResponse) SetBody(v *ExecuteAITeacherExpansionDialogueTranslateResponseBody) *ExecuteAITeacherExpansionDialogueTranslateResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherGrammarCheckRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i is good
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherGrammarCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherGrammarCheckRequest) SetContent(v string) *ExecuteAITeacherGrammarCheckRequest {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckRequest) SetUserId(v string) *ExecuteAITeacherGrammarCheckRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherGrammarCheckResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherGrammarCheckResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherGrammarCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetData(v *ExecuteAITeacherGrammarCheckResponseBodyData) *ExecuteAITeacherGrammarCheckResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetErrCode(v string) *ExecuteAITeacherGrammarCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetErrMessage(v string) *ExecuteAITeacherGrammarCheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherGrammarCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetRequestId(v string) *ExecuteAITeacherGrammarCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBody) SetSuccess(v bool) *ExecuteAITeacherGrammarCheckResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherGrammarCheckResponseBodyData struct {
	// example:
	//
	// 主语 "I" 对应的动词应该是 "am" 而不是 "is"。
	Analysis *string `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// example:
	//
	// I am good.
	Correction *string `json:"correction,omitempty" xml:"correction,omitempty"`
	// example:
	//
	// Has_Error
	CorrectionStatus *string `json:"correctionStatus,omitempty" xml:"correctionStatus,omitempty"`
	ErrorReason      *string `json:"errorReason,omitempty" xml:"errorReason,omitempty"`
}

func (s ExecuteAITeacherGrammarCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetAnalysis(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
	s.Analysis = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetCorrection(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
	s.Correction = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetCorrectionStatus(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
	s.CorrectionStatus = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponseBodyData) SetErrorReason(v string) *ExecuteAITeacherGrammarCheckResponseBodyData {
	s.ErrorReason = &v
	return s
}

type ExecuteAITeacherGrammarCheckResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherGrammarCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherGrammarCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherGrammarCheckResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherGrammarCheckResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponse) SetStatusCode(v int32) *ExecuteAITeacherGrammarCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherGrammarCheckResponse) SetBody(v *ExecuteAITeacherGrammarCheckResponseBody) *ExecuteAITeacherGrammarCheckResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherSyncDialogueRequest struct {
	// This parameter is required.
	DialogueTasks []*ExecuteAITeacherSyncDialogueRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string                                       `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	Records      []*ExecuteAITeacherSyncDialogueRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetDialogueTasks(v []*ExecuteAITeacherSyncDialogueRequestDialogueTasks) *ExecuteAITeacherSyncDialogueRequest {
	s.DialogueTasks = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetLanguageCode(v string) *ExecuteAITeacherSyncDialogueRequest {
	s.LanguageCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetRecords(v []*ExecuteAITeacherSyncDialogueRequestRecords) *ExecuteAITeacherSyncDialogueRequest {
	s.Records = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetUserId(v string) *ExecuteAITeacherSyncDialogueRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherSyncDialogueRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
	s.User = &v
	return s
}

type ExecuteAITeacherSyncDialogueRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueRequestRecords) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetContent(v string) *ExecuteAITeacherSyncDialogueRequestRecords {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherSyncDialogueRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherSyncDialogueRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetOrder(v int32) *ExecuteAITeacherSyncDialogueRequestRecords {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetRole(v string) *ExecuteAITeacherSyncDialogueRequestRecords {
	s.Role = &v
	return s
}

type ExecuteAITeacherSyncDialogueResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherSyncDialogueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherSyncDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetData(v *ExecuteAITeacherSyncDialogueResponseBodyData) *ExecuteAITeacherSyncDialogueResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetErrCode(v string) *ExecuteAITeacherSyncDialogueResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetErrMessage(v string) *ExecuteAITeacherSyncDialogueResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherSyncDialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetRequestId(v string) *ExecuteAITeacherSyncDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBody) SetSuccess(v bool) *ExecuteAITeacherSyncDialogueResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherSyncDialogueResponseBodyData struct {
	// example:
	//
	// Thanks, Lily. Do you like meat, Lily?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
	// example:
	//
	// true
	IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
	// example:
	//
	// true
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// example:
	//
	// 2
	QuestionIndex *int32 `json:"questionIndex,omitempty" xml:"questionIndex,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetEnglishResult(v string) *ExecuteAITeacherSyncDialogueResponseBodyData {
	s.EnglishResult = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetIsFinish(v bool) *ExecuteAITeacherSyncDialogueResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetIsOnTopic(v bool) *ExecuteAITeacherSyncDialogueResponseBodyData {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponseBodyData) SetQuestionIndex(v int32) *ExecuteAITeacherSyncDialogueResponseBodyData {
	s.QuestionIndex = &v
	return s
}

type ExecuteAITeacherSyncDialogueResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherSyncDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherSyncDialogueResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponse) SetStatusCode(v int32) *ExecuteAITeacherSyncDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueResponse) SetBody(v *ExecuteAITeacherSyncDialogueResponseBody) *ExecuteAITeacherSyncDialogueResponse {
	s.Body = v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateRequest struct {
	// This parameter is required.
	DialogueTasks []*ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	Records       []*ExecuteAITeacherSyncDialogueTranslateRequestRecords       `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateRequest) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) SetDialogueTasks(v []*ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) *ExecuteAITeacherSyncDialogueTranslateRequest {
	s.DialogueTasks = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) SetRecords(v []*ExecuteAITeacherSyncDialogueTranslateRequestRecords) *ExecuteAITeacherSyncDialogueTranslateRequest {
	s.Records = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) SetUserId(v string) *ExecuteAITeacherSyncDialogueTranslateRequest {
	s.UserId = &v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
	s.User = &v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateRequestRecords) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetContent(v string) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
	s.Content = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetOrder(v int32) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
	s.Order = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetRole(v string) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
	s.Role = &v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateResponseBody struct {
	// example:
	//
	// []
	Data *ExecuteAITeacherSyncDialogueTranslateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ExecuteAITeacherSyncDialogueTranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetData(v *ExecuteAITeacherSyncDialogueTranslateResponseBodyData) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetErrCode(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetErrMessage(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetHttpStatusCode(v int32) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetRequestId(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBody) SetSuccess(v bool) *ExecuteAITeacherSyncDialogueTranslateResponseBody {
	s.Success = &v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateResponseBodyData struct {
	// example:
	//
	// 太好了，谢谢你过来，莉莉。你喜欢吃肉吗？
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponseBodyData) SetResult(v string) *ExecuteAITeacherSyncDialogueTranslateResponseBodyData {
	s.Result = &v
	return s
}

type ExecuteAITeacherSyncDialogueTranslateResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteAITeacherSyncDialogueTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateResponse) GoString() string {
	return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) SetHeaders(v map[string]*string) *ExecuteAITeacherSyncDialogueTranslateResponse {
	s.Headers = v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) SetStatusCode(v int32) *ExecuteAITeacherSyncDialogueTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateResponse) SetBody(v *ExecuteAITeacherSyncDialogueTranslateResponseBody) *ExecuteAITeacherSyncDialogueTranslateResponse {
	s.Body = v
	return s
}

type ExecuteHundredThousandWhysDialogueRequest struct {
	// example:
	//
	// CHILDREN
	AgeGroup *string `json:"ageGroup,omitempty" xml:"ageGroup,omitempty"`
	// example:
	//
	// 2f28670c-eba6-4afb-9610-0942c434a021
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 700d4d9411efbe6e0
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 00:1A:2B:3C:4D:5E
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	// This parameter is required.
	Messages []*ExecuteHundredThousandWhysDialogueRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s ExecuteHundredThousandWhysDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteHundredThousandWhysDialogueRequest) GoString() string {
	return s.String()
}

func (s *ExecuteHundredThousandWhysDialogueRequest) SetAgeGroup(v string) *ExecuteHundredThousandWhysDialogueRequest {
	s.AgeGroup = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueRequest) SetChatId(v string) *ExecuteHundredThousandWhysDialogueRequest {
	s.ChatId = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueRequest) SetDeviceId(v string) *ExecuteHundredThousandWhysDialogueRequest {
	s.DeviceId = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueRequest) SetMacAddress(v string) *ExecuteHundredThousandWhysDialogueRequest {
	s.MacAddress = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueRequest) SetMessages(v []*ExecuteHundredThousandWhysDialogueRequestMessages) *ExecuteHundredThousandWhysDialogueRequest {
	s.Messages = v
	return s
}

type ExecuteHundredThousandWhysDialogueRequestMessages struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteHundredThousandWhysDialogueRequestMessages) String() string {
	return tea.Prettify(s)
}

func (s ExecuteHundredThousandWhysDialogueRequestMessages) GoString() string {
	return s.String()
}

func (s *ExecuteHundredThousandWhysDialogueRequestMessages) SetContent(v string) *ExecuteHundredThousandWhysDialogueRequestMessages {
	s.Content = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueRequestMessages) SetRole(v string) *ExecuteHundredThousandWhysDialogueRequestMessages {
	s.Role = &v
	return s
}

type ExecuteHundredThousandWhysDialogueResponseBody struct {
	Choices []*ExecuteHundredThousandWhysDialogueResponseBodyChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// null
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// null
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// example:
	//
	// null
	Object *string `json:"object,omitempty" xml:"object,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DBFA232A-1176-50E6-95AE-50F7A62A28AD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExecuteHundredThousandWhysDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteHundredThousandWhysDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteHundredThousandWhysDialogueResponseBody) SetChoices(v []*ExecuteHundredThousandWhysDialogueResponseBodyChoices) *ExecuteHundredThousandWhysDialogueResponseBody {
	s.Choices = v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueResponseBody) SetCreated(v int64) *ExecuteHundredThousandWhysDialogueResponseBody {
	s.Created = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueResponseBody) SetId(v string) *ExecuteHundredThousandWhysDialogueResponseBody {
	s.Id = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueResponseBody) SetModel(v string) *ExecuteHundredThousandWhysDialogueResponseBody {
	s.Model = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueResponseBody) SetObject(v string) *ExecuteHundredThousandWhysDialogueResponseBody {
	s.Object = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueResponseBody) SetRequestId(v string) *ExecuteHundredThousandWhysDialogueResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteHundredThousandWhysDialogueResponseBodyChoices struct {
	Delta *ExecuteHundredThousandWhysDialogueResponseBodyChoicesDelta `json:"delta,omitempty" xml:"delta,omitempty" type:"Struct"`
	// example:
	//
	// stop
	FinishReason *string `json:"finish_reason,omitempty" xml:"finish_reason,omitempty"`
	// example:
	//
	// 0
	Index *int64 `json:"index,omitempty" xml:"index,omitempty"`
}

func (s ExecuteHundredThousandWhysDialogueResponseBodyChoices) String() string {
	return tea.Prettify(s)
}

func (s ExecuteHundredThousandWhysDialogueResponseBodyChoices) GoString() string {
	return s.String()
}

func (s *ExecuteHundredThousandWhysDialogueResponseBodyChoices) SetDelta(v *ExecuteHundredThousandWhysDialogueResponseBodyChoicesDelta) *ExecuteHundredThousandWhysDialogueResponseBodyChoices {
	s.Delta = v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueResponseBodyChoices) SetFinishReason(v string) *ExecuteHundredThousandWhysDialogueResponseBodyChoices {
	s.FinishReason = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueResponseBodyChoices) SetIndex(v int64) *ExecuteHundredThousandWhysDialogueResponseBodyChoices {
	s.Index = &v
	return s
}

type ExecuteHundredThousandWhysDialogueResponseBodyChoicesDelta struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// assistant
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteHundredThousandWhysDialogueResponseBodyChoicesDelta) String() string {
	return tea.Prettify(s)
}

func (s ExecuteHundredThousandWhysDialogueResponseBodyChoicesDelta) GoString() string {
	return s.String()
}

func (s *ExecuteHundredThousandWhysDialogueResponseBodyChoicesDelta) SetContent(v string) *ExecuteHundredThousandWhysDialogueResponseBodyChoicesDelta {
	s.Content = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueResponseBodyChoicesDelta) SetRole(v string) *ExecuteHundredThousandWhysDialogueResponseBodyChoicesDelta {
	s.Role = &v
	return s
}

type ExecuteHundredThousandWhysDialogueResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteHundredThousandWhysDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteHundredThousandWhysDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteHundredThousandWhysDialogueResponse) GoString() string {
	return s.String()
}

func (s *ExecuteHundredThousandWhysDialogueResponse) SetHeaders(v map[string]*string) *ExecuteHundredThousandWhysDialogueResponse {
	s.Headers = v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueResponse) SetStatusCode(v int32) *ExecuteHundredThousandWhysDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteHundredThousandWhysDialogueResponse) SetBody(v *ExecuteHundredThousandWhysDialogueResponseBody) *ExecuteHundredThousandWhysDialogueResponse {
	s.Body = v
	return s
}

type ExecuteTextbookAssistantDialogueRequest struct {
	// This parameter is required.
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
	// This parameter is required.
	UserMessage *string `json:"userMessage,omitempty" xml:"userMessage,omitempty"`
}

func (s ExecuteTextbookAssistantDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantDialogueRequest) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantDialogueRequest) SetAuthToken(v string) *ExecuteTextbookAssistantDialogueRequest {
	s.AuthToken = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueRequest) SetChatId(v string) *ExecuteTextbookAssistantDialogueRequest {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueRequest) SetScenario(v string) *ExecuteTextbookAssistantDialogueRequest {
	s.Scenario = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueRequest) SetUserMessage(v string) *ExecuteTextbookAssistantDialogueRequest {
	s.UserMessage = &v
	return s
}

type ExecuteTextbookAssistantDialogueResponseBody struct {
	Data           *ExecuteTextbookAssistantDialogueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode        *string                                           `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage     *string                                           `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	HttpStatusCode *string                                           `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetData(v *ExecuteTextbookAssistantDialogueResponseBodyData) *ExecuteTextbookAssistantDialogueResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantDialogueResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantDialogueResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetHttpStatusCode(v string) *ExecuteTextbookAssistantDialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBody) SetSuccess(v string) *ExecuteTextbookAssistantDialogueResponseBody {
	s.Success = &v
	return s
}

type ExecuteTextbookAssistantDialogueResponseBodyData struct {
	Assistant *string                                                 `json:"assistant,omitempty" xml:"assistant,omitempty"`
	ChatId    *string                                                 `json:"chatId,omitempty" xml:"chatId,omitempty"`
	Result    *ExecuteTextbookAssistantDialogueResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	User      *string                                                 `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteTextbookAssistantDialogueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantDialogueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) SetAssistant(v string) *ExecuteTextbookAssistantDialogueResponseBodyData {
	s.Assistant = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) SetChatId(v string) *ExecuteTextbookAssistantDialogueResponseBodyData {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) SetResult(v *ExecuteTextbookAssistantDialogueResponseBodyDataResult) *ExecuteTextbookAssistantDialogueResponseBodyData {
	s.Result = v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyData) SetUser(v string) *ExecuteTextbookAssistantDialogueResponseBodyData {
	s.User = &v
	return s
}

type ExecuteTextbookAssistantDialogueResponseBodyDataResult struct {
	ChineseResult   *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
	EnglishResult   *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
	IsFinish        *bool   `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
	IsTaskCompleted *bool   `json:"isTaskCompleted,omitempty" xml:"isTaskCompleted,omitempty"`
}

func (s ExecuteTextbookAssistantDialogueResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantDialogueResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) SetChineseResult(v string) *ExecuteTextbookAssistantDialogueResponseBodyDataResult {
	s.ChineseResult = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) SetEnglishResult(v string) *ExecuteTextbookAssistantDialogueResponseBodyDataResult {
	s.EnglishResult = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) SetIsFinish(v bool) *ExecuteTextbookAssistantDialogueResponseBodyDataResult {
	s.IsFinish = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponseBodyDataResult) SetIsTaskCompleted(v bool) *ExecuteTextbookAssistantDialogueResponseBodyDataResult {
	s.IsTaskCompleted = &v
	return s
}

type ExecuteTextbookAssistantDialogueResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteTextbookAssistantDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantDialogueResponse) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantDialogueResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantDialogueResponse {
	s.Headers = v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantDialogueResponse) SetBody(v *ExecuteTextbookAssistantDialogueResponseBody) *ExecuteTextbookAssistantDialogueResponse {
	s.Body = v
	return s
}

type ExecuteTextbookAssistantDifficultyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// UP
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6788f4a6b54c5268c1b78a25
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6788e0b475a4631ffc626722
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ExecuteTextbookAssistantDifficultyRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantDifficultyRequest) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantDifficultyRequest) SetAction(v string) *ExecuteTextbookAssistantDifficultyRequest {
	s.Action = &v
	return s
}

func (s *ExecuteTextbookAssistantDifficultyRequest) SetAssistant(v string) *ExecuteTextbookAssistantDifficultyRequest {
	s.Assistant = &v
	return s
}

func (s *ExecuteTextbookAssistantDifficultyRequest) SetAuthToken(v string) *ExecuteTextbookAssistantDifficultyRequest {
	s.AuthToken = &v
	return s
}

func (s *ExecuteTextbookAssistantDifficultyRequest) SetChatId(v string) *ExecuteTextbookAssistantDifficultyRequest {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantDifficultyRequest) SetScenario(v string) *ExecuteTextbookAssistantDifficultyRequest {
	s.Scenario = &v
	return s
}

type ExecuteTextbookAssistantDifficultyResponseBody struct {
	Data *ExecuteTextbookAssistantDifficultyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0D7D382F-9475-572E-BE83-DDFBF5C5EB24
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantDifficultyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantDifficultyResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetData(v *ExecuteTextbookAssistantDifficultyResponseBodyData) *ExecuteTextbookAssistantDifficultyResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantDifficultyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantDifficultyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetHttpStatusCode(v int32) *ExecuteTextbookAssistantDifficultyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantDifficultyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteTextbookAssistantDifficultyResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantDifficultyResponseBody {
	s.Success = &v
	return s
}

type ExecuteTextbookAssistantDifficultyResponseBodyData struct {
	Result *ExecuteTextbookAssistantDifficultyResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteTextbookAssistantDifficultyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantDifficultyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantDifficultyResponseBodyData) SetResult(v *ExecuteTextbookAssistantDifficultyResponseBodyDataResult) *ExecuteTextbookAssistantDifficultyResponseBodyData {
	s.Result = v
	return s
}

type ExecuteTextbookAssistantDifficultyResponseBodyDataResult struct {
	// example:
	//
	// Let\\"s look at the text again. Mike says, \\"I\\"m Mike Black.\\" Can you try saying it like Mike?
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteTextbookAssistantDifficultyResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantDifficultyResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantDifficultyResponseBodyDataResult) SetResult(v string) *ExecuteTextbookAssistantDifficultyResponseBodyDataResult {
	s.Result = &v
	return s
}

type ExecuteTextbookAssistantDifficultyResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteTextbookAssistantDifficultyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantDifficultyResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantDifficultyResponse) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantDifficultyResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantDifficultyResponse {
	s.Headers = v
	return s
}

func (s *ExecuteTextbookAssistantDifficultyResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantDifficultyResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantDifficultyResponse) SetBody(v *ExecuteTextbookAssistantDifficultyResponseBody) *ExecuteTextbookAssistantDifficultyResponse {
	s.Body = v
	return s
}

type ExecuteTextbookAssistantGrammarCheckRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6788e0b475a4631ffc626722
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6788e0b45bdfc807f077a5a1
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteTextbookAssistantGrammarCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantGrammarCheckRequest) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) SetAuthToken(v string) *ExecuteTextbookAssistantGrammarCheckRequest {
	s.AuthToken = &v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) SetChatId(v string) *ExecuteTextbookAssistantGrammarCheckRequest {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) SetScenario(v string) *ExecuteTextbookAssistantGrammarCheckRequest {
	s.Scenario = &v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) SetUser(v string) *ExecuteTextbookAssistantGrammarCheckRequest {
	s.User = &v
	return s
}

type ExecuteTextbookAssistantGrammarCheckResponseBody struct {
	Data *ExecuteTextbookAssistantGrammarCheckResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0bc1e96d17091734639835114e12c8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetData(v *ExecuteTextbookAssistantGrammarCheckResponseBodyData) *ExecuteTextbookAssistantGrammarCheckResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantGrammarCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantGrammarCheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetHttpStatusCode(v int32) *ExecuteTextbookAssistantGrammarCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantGrammarCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantGrammarCheckResponseBody {
	s.Success = &v
	return s
}

type ExecuteTextbookAssistantGrammarCheckResponseBodyData struct {
	Result *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyData) SetResult(v *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) *ExecuteTextbookAssistantGrammarCheckResponseBodyData {
	s.Result = v
	return s
}

type ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult struct {
	Analysis *string `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// example:
	//
	// I am you.
	Correction *string `json:"correction,omitempty" xml:"correction,omitempty"`
	// example:
	//
	// Has_Error
	CorrectionStatus *string `json:"correctionStatus,omitempty" xml:"correctionStatus,omitempty"`
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) SetAnalysis(v string) *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult {
	s.Analysis = &v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) SetCorrection(v string) *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult {
	s.Correction = &v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult) SetCorrectionStatus(v string) *ExecuteTextbookAssistantGrammarCheckResponseBodyDataResult {
	s.CorrectionStatus = &v
	return s
}

type ExecuteTextbookAssistantGrammarCheckResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteTextbookAssistantGrammarCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantGrammarCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantGrammarCheckResponse) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantGrammarCheckResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantGrammarCheckResponse {
	s.Headers = v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantGrammarCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantGrammarCheckResponse) SetBody(v *ExecuteTextbookAssistantGrammarCheckResponseBody) *ExecuteTextbookAssistantGrammarCheckResponse {
	s.Body = v
	return s
}

type ExecuteTextbookAssistantRefineByContextRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6788e0b475a4631ffc626722
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6788e0b45bdfc807f077a5a1
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteTextbookAssistantRefineByContextRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantRefineByContextRequest) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) SetAuthToken(v string) *ExecuteTextbookAssistantRefineByContextRequest {
	s.AuthToken = &v
	return s
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) SetChatId(v string) *ExecuteTextbookAssistantRefineByContextRequest {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) SetScenario(v string) *ExecuteTextbookAssistantRefineByContextRequest {
	s.Scenario = &v
	return s
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) SetUser(v string) *ExecuteTextbookAssistantRefineByContextRequest {
	s.User = &v
	return s
}

type ExecuteTextbookAssistantRefineByContextResponseBody struct {
	Data *ExecuteTextbookAssistantRefineByContextResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *int32 `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6F73C114-A76E-51AD-99E3-BC7B941B69E0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantRefineByContextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantRefineByContextResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetData(v *ExecuteTextbookAssistantRefineByContextResponseBodyData) *ExecuteTextbookAssistantRefineByContextResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantRefineByContextResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetErrMessage(v int32) *ExecuteTextbookAssistantRefineByContextResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetHttpStatusCode(v string) *ExecuteTextbookAssistantRefineByContextResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantRefineByContextResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantRefineByContextResponseBody {
	s.Success = &v
	return s
}

type ExecuteTextbookAssistantRefineByContextResponseBodyData struct {
	Result *ExecuteTextbookAssistantRefineByContextResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteTextbookAssistantRefineByContextResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantRefineByContextResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBodyData) SetResult(v *ExecuteTextbookAssistantRefineByContextResponseBodyDataResult) *ExecuteTextbookAssistantRefineByContextResponseBodyData {
	s.Result = v
	return s
}

type ExecuteTextbookAssistantRefineByContextResponseBodyDataResult struct {
	// example:
	//
	// Good evening! From the book, how does Mike Black introduce himself?
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteTextbookAssistantRefineByContextResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantRefineByContextResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantRefineByContextResponseBodyDataResult) SetResult(v string) *ExecuteTextbookAssistantRefineByContextResponseBodyDataResult {
	s.Result = &v
	return s
}

type ExecuteTextbookAssistantRefineByContextResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteTextbookAssistantRefineByContextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantRefineByContextResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantRefineByContextResponse) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantRefineByContextResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantRefineByContextResponse {
	s.Headers = v
	return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantRefineByContextResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantRefineByContextResponse) SetBody(v *ExecuteTextbookAssistantRefineByContextResponseBody) *ExecuteTextbookAssistantRefineByContextResponse {
	s.Body = v
	return s
}

type ExecuteTextbookAssistantRetryConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6788e0b4b54c5268c1b78638
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6788e0b475a4631ffc626722
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ExecuteTextbookAssistantRetryConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantRetryConversationRequest) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) SetAssistant(v string) *ExecuteTextbookAssistantRetryConversationRequest {
	s.Assistant = &v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) SetAuthToken(v string) *ExecuteTextbookAssistantRetryConversationRequest {
	s.AuthToken = &v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) SetChatId(v string) *ExecuteTextbookAssistantRetryConversationRequest {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) SetScenario(v string) *ExecuteTextbookAssistantRetryConversationRequest {
	s.Scenario = &v
	return s
}

type ExecuteTextbookAssistantRetryConversationResponseBody struct {
	Data *ExecuteTextbookAssistantRetryConversationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2F2ABF4B-A4F6-5EC7-B287-7EF5B156F1ED
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantRetryConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantRetryConversationResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetData(v *ExecuteTextbookAssistantRetryConversationResponseBodyData) *ExecuteTextbookAssistantRetryConversationResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantRetryConversationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantRetryConversationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetHttpStatusCode(v int32) *ExecuteTextbookAssistantRetryConversationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantRetryConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantRetryConversationResponseBody {
	s.Success = &v
	return s
}

type ExecuteTextbookAssistantRetryConversationResponseBodyData struct {
	// example:
	//
	// 6788e0b4b54c5268c1b78638
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 6788e0b475a4631ffc626722
	ChatId *string                                                          `json:"chatId,omitempty" xml:"chatId,omitempty"`
	Result *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// 6788e0b45bdfc807f077a5a1
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteTextbookAssistantRetryConversationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantRetryConversationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) SetAssistant(v string) *ExecuteTextbookAssistantRetryConversationResponseBodyData {
	s.Assistant = &v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) SetChatId(v string) *ExecuteTextbookAssistantRetryConversationResponseBodyData {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) SetResult(v *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) *ExecuteTextbookAssistantRetryConversationResponseBodyData {
	s.Result = v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyData) SetUser(v string) *ExecuteTextbookAssistantRetryConversationResponseBodyData {
	s.User = &v
	return s
}

type ExecuteTextbookAssistantRetryConversationResponseBodyDataResult struct {
	ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
	// example:
	//
	// Good evening! From the book, how does Mike Black introduce himself?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
}

func (s ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) SetChineseResult(v string) *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult {
	s.ChineseResult = &v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult) SetEnglishResult(v string) *ExecuteTextbookAssistantRetryConversationResponseBodyDataResult {
	s.EnglishResult = &v
	return s
}

type ExecuteTextbookAssistantRetryConversationResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteTextbookAssistantRetryConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantRetryConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantRetryConversationResponse) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantRetryConversationResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantRetryConversationResponse {
	s.Headers = v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantRetryConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantRetryConversationResponse) SetBody(v *ExecuteTextbookAssistantRetryConversationResponseBody) *ExecuteTextbookAssistantRetryConversationResponse {
	s.Body = v
	return s
}

type ExecuteTextbookAssistantSseDialogueRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 67e374acb54c526c95c4fbd4
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EXPAND
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hello
	UserMessage *string `json:"userMessage,omitempty" xml:"userMessage,omitempty"`
}

func (s ExecuteTextbookAssistantSseDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantSseDialogueRequest) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) SetAuthToken(v string) *ExecuteTextbookAssistantSseDialogueRequest {
	s.AuthToken = &v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) SetChatId(v string) *ExecuteTextbookAssistantSseDialogueRequest {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) SetScenario(v string) *ExecuteTextbookAssistantSseDialogueRequest {
	s.Scenario = &v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) SetUserMessage(v string) *ExecuteTextbookAssistantSseDialogueRequest {
	s.UserMessage = &v
	return s
}

type ExecuteTextbookAssistantSseDialogueResponseBody struct {
	// example:
	//
	// 67e4c9d95bdfc83cd742ae7c
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 67e374acb54c526c95c4fbd4
	ChatId *string                                              `json:"chatId,omitempty" xml:"chatId,omitempty"`
	Data   *ExecuteTextbookAssistantSseDialogueResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// BIZ_ERROR
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 67e4c9d6b54c526c95c53925
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteTextbookAssistantSseDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantSseDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetAssistant(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
	s.Assistant = &v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetChatId(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetData(v *ExecuteTextbookAssistantSseDialogueResponseBodyData) *ExecuteTextbookAssistantSseDialogueResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantSseDialogueResponseBody {
	s.Success = &v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBody) SetUser(v string) *ExecuteTextbookAssistantSseDialogueResponseBody {
	s.User = &v
	return s
}

type ExecuteTextbookAssistantSseDialogueResponseBodyData struct {
	// example:
	//
	// Thanks, Lily. Do you like meat, Lily?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
	// example:
	//
	// true
	IsFinish *bool `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
}

func (s ExecuteTextbookAssistantSseDialogueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantSseDialogueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBodyData) SetEnglishResult(v string) *ExecuteTextbookAssistantSseDialogueResponseBodyData {
	s.EnglishResult = &v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponseBodyData) SetIsFinish(v bool) *ExecuteTextbookAssistantSseDialogueResponseBodyData {
	s.IsFinish = &v
	return s
}

type ExecuteTextbookAssistantSseDialogueResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteTextbookAssistantSseDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantSseDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantSseDialogueResponse) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantSseDialogueResponse {
	s.Headers = v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantSseDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantSseDialogueResponse) SetBody(v *ExecuteTextbookAssistantSseDialogueResponseBody) *ExecuteTextbookAssistantSseDialogueResponse {
	s.Body = v
	return s
}

type ExecuteTextbookAssistantStartConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0c05700d4d9411efbe6e0c42a106bb02
	ArticleId *string `json:"articleId,omitempty" xml:"articleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ExecuteTextbookAssistantStartConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantStartConversationRequest) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantStartConversationRequest) SetArticleId(v string) *ExecuteTextbookAssistantStartConversationRequest {
	s.ArticleId = &v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationRequest) SetAuthToken(v string) *ExecuteTextbookAssistantStartConversationRequest {
	s.AuthToken = &v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationRequest) SetScenario(v string) *ExecuteTextbookAssistantStartConversationRequest {
	s.Scenario = &v
	return s
}

type ExecuteTextbookAssistantStartConversationResponseBody struct {
	Data *ExecuteTextbookAssistantStartConversationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// B_USER_NOT_FOUND_EXCEPTION
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6F73C114-A76E-51AD-99E3-BC7B941B69E0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantStartConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantStartConversationResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetData(v *ExecuteTextbookAssistantStartConversationResponseBodyData) *ExecuteTextbookAssistantStartConversationResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantStartConversationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantStartConversationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetHttpStatusCode(v int32) *ExecuteTextbookAssistantStartConversationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantStartConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantStartConversationResponseBody {
	s.Success = &v
	return s
}

type ExecuteTextbookAssistantStartConversationResponseBodyData struct {
	// example:
	//
	// 6788e0b4b54c5268c1b78638
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 6788e0b475a4631ffc626722
	ChatId *string                                                          `json:"chatId,omitempty" xml:"chatId,omitempty"`
	Result *ExecuteTextbookAssistantStartConversationResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// 6788e0b45bdfc807f077a5a1
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteTextbookAssistantStartConversationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantStartConversationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) SetAssistant(v string) *ExecuteTextbookAssistantStartConversationResponseBodyData {
	s.Assistant = &v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) SetChatId(v string) *ExecuteTextbookAssistantStartConversationResponseBodyData {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) SetResult(v *ExecuteTextbookAssistantStartConversationResponseBodyDataResult) *ExecuteTextbookAssistantStartConversationResponseBodyData {
	s.Result = v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyData) SetUser(v string) *ExecuteTextbookAssistantStartConversationResponseBodyData {
	s.User = &v
	return s
}

type ExecuteTextbookAssistantStartConversationResponseBodyDataResult struct {
	ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
	// example:
	//
	// Good evening! From the book, how does Mike Black introduce himself?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
}

func (s ExecuteTextbookAssistantStartConversationResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantStartConversationResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyDataResult) SetChineseResult(v string) *ExecuteTextbookAssistantStartConversationResponseBodyDataResult {
	s.ChineseResult = &v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationResponseBodyDataResult) SetEnglishResult(v string) *ExecuteTextbookAssistantStartConversationResponseBodyDataResult {
	s.EnglishResult = &v
	return s
}

type ExecuteTextbookAssistantStartConversationResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteTextbookAssistantStartConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantStartConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantStartConversationResponse) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantStartConversationResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantStartConversationResponse {
	s.Headers = v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantStartConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantStartConversationResponse) SetBody(v *ExecuteTextbookAssistantStartConversationResponseBody) *ExecuteTextbookAssistantStartConversationResponse {
	s.Body = v
	return s
}

type ExecuteTextbookAssistantSuggestionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6788e0b4b54c5268c1b78638
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6788e0b4b54c5268c1b78638
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ExecuteTextbookAssistantSuggestionRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantSuggestionRequest) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantSuggestionRequest) SetAssistant(v string) *ExecuteTextbookAssistantSuggestionRequest {
	s.Assistant = &v
	return s
}

func (s *ExecuteTextbookAssistantSuggestionRequest) SetAuthToken(v string) *ExecuteTextbookAssistantSuggestionRequest {
	s.AuthToken = &v
	return s
}

func (s *ExecuteTextbookAssistantSuggestionRequest) SetChatId(v string) *ExecuteTextbookAssistantSuggestionRequest {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantSuggestionRequest) SetScenario(v string) *ExecuteTextbookAssistantSuggestionRequest {
	s.Scenario = &v
	return s
}

type ExecuteTextbookAssistantSuggestionResponseBody struct {
	Data *ExecuteTextbookAssistantSuggestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpstatusCode *int32 `json:"httpstatusCode,omitempty" xml:"httpstatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0D7D382F-9475-572E-BE83-DDFBF5C5EB24
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantSuggestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantSuggestionResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetData(v *ExecuteTextbookAssistantSuggestionResponseBodyData) *ExecuteTextbookAssistantSuggestionResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantSuggestionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantSuggestionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetHttpstatusCode(v int32) *ExecuteTextbookAssistantSuggestionResponseBody {
	s.HttpstatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantSuggestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantSuggestionResponseBody {
	s.Success = &v
	return s
}

type ExecuteTextbookAssistantSuggestionResponseBodyData struct {
	Result *ExecuteTextbookAssistantSuggestionResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteTextbookAssistantSuggestionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantSuggestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantSuggestionResponseBodyData) SetResult(v *ExecuteTextbookAssistantSuggestionResponseBodyDataResult) *ExecuteTextbookAssistantSuggestionResponseBodyData {
	s.Result = v
	return s
}

type ExecuteTextbookAssistantSuggestionResponseBodyDataResult struct {
	ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
	// example:
	//
	// Good evening! From the book, how does Mike Black introduce himself?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
}

func (s ExecuteTextbookAssistantSuggestionResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantSuggestionResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantSuggestionResponseBodyDataResult) SetChineseResult(v string) *ExecuteTextbookAssistantSuggestionResponseBodyDataResult {
	s.ChineseResult = &v
	return s
}

func (s *ExecuteTextbookAssistantSuggestionResponseBodyDataResult) SetEnglishResult(v string) *ExecuteTextbookAssistantSuggestionResponseBodyDataResult {
	s.EnglishResult = &v
	return s
}

type ExecuteTextbookAssistantSuggestionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteTextbookAssistantSuggestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantSuggestionResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantSuggestionResponse) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantSuggestionResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantSuggestionResponse {
	s.Headers = v
	return s
}

func (s *ExecuteTextbookAssistantSuggestionResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantSuggestionResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantSuggestionResponse) SetBody(v *ExecuteTextbookAssistantSuggestionResponseBody) *ExecuteTextbookAssistantSuggestionResponse {
	s.Body = v
	return s
}

type ExecuteTextbookAssistantTranslateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6788e0b4b54c5268c1b78638
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6788e0b475a4631ffc626722
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ExecuteTextbookAssistantTranslateRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantTranslateRequest) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantTranslateRequest) SetAssistant(v string) *ExecuteTextbookAssistantTranslateRequest {
	s.Assistant = &v
	return s
}

func (s *ExecuteTextbookAssistantTranslateRequest) SetAuthToken(v string) *ExecuteTextbookAssistantTranslateRequest {
	s.AuthToken = &v
	return s
}

func (s *ExecuteTextbookAssistantTranslateRequest) SetChatId(v string) *ExecuteTextbookAssistantTranslateRequest {
	s.ChatId = &v
	return s
}

func (s *ExecuteTextbookAssistantTranslateRequest) SetScenario(v string) *ExecuteTextbookAssistantTranslateRequest {
	s.Scenario = &v
	return s
}

type ExecuteTextbookAssistantTranslateResponseBody struct {
	Data *ExecuteTextbookAssistantTranslateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ""
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9EB79C1E-36C2-5777-BED6-C23A98DF0637
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteTextbookAssistantTranslateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetData(v *ExecuteTextbookAssistantTranslateResponseBodyData) *ExecuteTextbookAssistantTranslateResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetErrCode(v string) *ExecuteTextbookAssistantTranslateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetErrMessage(v string) *ExecuteTextbookAssistantTranslateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetHttpStatusCode(v int32) *ExecuteTextbookAssistantTranslateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetRequestId(v string) *ExecuteTextbookAssistantTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteTextbookAssistantTranslateResponseBody) SetSuccess(v bool) *ExecuteTextbookAssistantTranslateResponseBody {
	s.Success = &v
	return s
}

type ExecuteTextbookAssistantTranslateResponseBodyData struct {
	Result *ExecuteTextbookAssistantTranslateResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ExecuteTextbookAssistantTranslateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantTranslateResponseBodyData) SetResult(v *ExecuteTextbookAssistantTranslateResponseBodyDataResult) *ExecuteTextbookAssistantTranslateResponseBodyData {
	s.Result = v
	return s
}

type ExecuteTextbookAssistantTranslateResponseBodyDataResult struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteTextbookAssistantTranslateResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantTranslateResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantTranslateResponseBodyDataResult) SetResult(v string) *ExecuteTextbookAssistantTranslateResponseBodyDataResult {
	s.Result = &v
	return s
}

type ExecuteTextbookAssistantTranslateResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteTextbookAssistantTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteTextbookAssistantTranslateResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTextbookAssistantTranslateResponse) GoString() string {
	return s.String()
}

func (s *ExecuteTextbookAssistantTranslateResponse) SetHeaders(v map[string]*string) *ExecuteTextbookAssistantTranslateResponse {
	s.Headers = v
	return s
}

func (s *ExecuteTextbookAssistantTranslateResponse) SetStatusCode(v int32) *ExecuteTextbookAssistantTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteTextbookAssistantTranslateResponse) SetBody(v *ExecuteTextbookAssistantTranslateResponseBody) *ExecuteTextbookAssistantTranslateResponse {
	s.Body = v
	return s
}

type GetAITeacherExpansionDialogueSuggestionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
	Background *string `json:"background,omitempty" xml:"background,omitempty"`
	// This parameter is required.
	DialogueTasks []*GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	// This parameter is required.
	Records []*GetAITeacherExpansionDialogueSuggestionRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	RoleInfo *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
	// example:
	//
	// Hello Lily, could you please come to the kitchen for a moment?
	StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Let\\"s talk about traffic rules.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequest) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetBackground(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.Background = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetDialogueTasks(v []*GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.DialogueTasks = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetLanguageCode(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.LanguageCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetRecords(v []*GetAITeacherExpansionDialogueSuggestionRequestRecords) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.Records = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetRoleInfo(v *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.RoleInfo = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetStartSentence(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.StartSentence = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetTopic(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.Topic = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetUserId(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.UserId = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetAssistant(v string) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetAssistantTranslate(v string) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetOrder(v int32) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetUser(v string) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.User = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequestRecords) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetContent(v string) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.Content = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetIsOffTopicControl(v bool) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetIsOnTopic(v bool) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetOrder(v int32) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.Order = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetRole(v string) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.Role = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionRequestRoleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// Alex
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jamie
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) SetAssistant(v string) *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo {
	s.Assistant = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) SetUser(v string) *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo {
	s.User = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionResponseBody struct {
	// example:
	//
	// []
	Data *GetAITeacherExpansionDialogueSuggestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetAITeacherExpansionDialogueSuggestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetData(v *GetAITeacherExpansionDialogueSuggestionResponseBodyData) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.Data = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetErrCode(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetErrMessage(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetHttpStatusCode(v int32) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetRequestId(v string) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBody) SetSuccess(v bool) *GetAITeacherExpansionDialogueSuggestionResponseBody {
	s.Success = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionResponseBodyData struct {
	// example:
	//
	// 谢谢莉莉.你喜欢吃肉吗，莉莉？
	ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
	// example:
	//
	// Thanks, Lily. Do you like meat, Lily?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBodyData) SetChineseResult(v string) *GetAITeacherExpansionDialogueSuggestionResponseBodyData {
	s.ChineseResult = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponseBodyData) SetEnglishResult(v string) *GetAITeacherExpansionDialogueSuggestionResponseBodyData {
	s.EnglishResult = &v
	return s
}

type GetAITeacherExpansionDialogueSuggestionResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAITeacherExpansionDialogueSuggestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionResponse) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) SetHeaders(v map[string]*string) *GetAITeacherExpansionDialogueSuggestionResponse {
	s.Headers = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) SetStatusCode(v int32) *GetAITeacherExpansionDialogueSuggestionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionResponse) SetBody(v *GetAITeacherExpansionDialogueSuggestionResponseBody) *GetAITeacherExpansionDialogueSuggestionResponse {
	s.Body = v
	return s
}

type GetAITeacherSyncDialogueSuggestionRequest struct {
	// This parameter is required.
	DialogueTasks []*GetAITeacherSyncDialogueSuggestionRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	// This parameter is required.
	Records []*GetAITeacherSyncDialogueSuggestionRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionRequest) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetDialogueTasks(v []*GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) *GetAITeacherSyncDialogueSuggestionRequest {
	s.DialogueTasks = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetLanguageCode(v string) *GetAITeacherSyncDialogueSuggestionRequest {
	s.LanguageCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetRecords(v []*GetAITeacherSyncDialogueSuggestionRequestRecords) *GetAITeacherSyncDialogueSuggestionRequest {
	s.Records = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetUserId(v string) *GetAITeacherSyncDialogueSuggestionRequest {
	s.UserId = &v
	return s
}

type GetAITeacherSyncDialogueSuggestionRequestDialogueTasks struct {
	// This parameter is required.
	//
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// 为什么有些人认为遛狗是份好差事?
	AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetAssistant(v string) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetAssistantTranslate(v string) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetOrder(v int32) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetUser(v string) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.User = &v
	return s
}

type GetAITeacherSyncDialogueSuggestionRequestRecords struct {
	// This parameter is required.
	//
	// example:
	//
	// Ask Mark if he has thought about what his dream job might be.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 跑题：true, 不跑题：false
	IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
	// example:
	//
	// 扣题：true, 不扣题：false
	IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 老师：assistant；学生：user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionRequestRecords) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetContent(v string) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.Content = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetIsOffTopicControl(v bool) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetIsOnTopic(v bool) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetOrder(v int32) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.Order = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetRole(v string) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.Role = &v
	return s
}

type GetAITeacherSyncDialogueSuggestionResponseBody struct {
	// example:
	//
	// []
	Data *GetAITeacherSyncDialogueSuggestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetAITeacherSyncDialogueSuggestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetData(v *GetAITeacherSyncDialogueSuggestionResponseBodyData) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.Data = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetErrCode(v string) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetErrMessage(v string) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetHttpStatusCode(v int32) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetRequestId(v string) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBody) SetSuccess(v bool) *GetAITeacherSyncDialogueSuggestionResponseBody {
	s.Success = &v
	return s
}

type GetAITeacherSyncDialogueSuggestionResponseBodyData struct {
	// example:
	//
	// 谢谢莉莉.你喜欢吃肉吗，莉莉？
	ChineseResult *string `json:"chineseResult,omitempty" xml:"chineseResult,omitempty"`
	// example:
	//
	// Thanks, Lily. Do you like meat, Lily?
	EnglishResult *string `json:"englishResult,omitempty" xml:"englishResult,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBodyData) SetChineseResult(v string) *GetAITeacherSyncDialogueSuggestionResponseBodyData {
	s.ChineseResult = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponseBodyData) SetEnglishResult(v string) *GetAITeacherSyncDialogueSuggestionResponseBodyData {
	s.EnglishResult = &v
	return s
}

type GetAITeacherSyncDialogueSuggestionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAITeacherSyncDialogueSuggestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionResponse) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) SetHeaders(v map[string]*string) *GetAITeacherSyncDialogueSuggestionResponse {
	s.Headers = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) SetStatusCode(v int32) *GetAITeacherSyncDialogueSuggestionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionResponse) SetBody(v *GetAITeacherSyncDialogueSuggestionResponseBody) *GetAITeacherSyncDialogueSuggestionResponse {
	s.Body = v
	return s
}

type GetTextbookAssistantTokenRequest struct {
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 700d4d9411efbe6e0
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 25032PS56C
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
}

func (s GetTextbookAssistantTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTextbookAssistantTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTextbookAssistantTokenRequest) SetDeviceId(v string) *GetTextbookAssistantTokenRequest {
	s.DeviceId = &v
	return s
}

func (s *GetTextbookAssistantTokenRequest) SetModel(v string) *GetTextbookAssistantTokenRequest {
	s.Model = &v
	return s
}

type GetTextbookAssistantTokenResponseBody struct {
	Data *GetTextbookAssistantTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0A5E9849-A2F0-551D-A7D8-1A8118557BAB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTextbookAssistantTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTextbookAssistantTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextbookAssistantTokenResponseBody) SetData(v *GetTextbookAssistantTokenResponseBodyData) *GetTextbookAssistantTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTextbookAssistantTokenResponseBody) SetErrCode(v string) *GetTextbookAssistantTokenResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBody) SetErrMessage(v string) *GetTextbookAssistantTokenResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBody) SetHttpStatusCode(v int32) *GetTextbookAssistantTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBody) SetRequestId(v string) *GetTextbookAssistantTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBody) SetSuccess(v bool) *GetTextbookAssistantTokenResponseBody {
	s.Success = &v
	return s
}

type GetTextbookAssistantTokenResponseBodyData struct {
	// example:
	//
	// tc_197bf5bb81889cc79eb51ae9b8c0cea3
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// example:
	//
	// 5400
	Expire *int32 `json:"expire,omitempty" xml:"expire,omitempty"`
}

func (s GetTextbookAssistantTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTextbookAssistantTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTextbookAssistantTokenResponseBodyData) SetAuthToken(v string) *GetTextbookAssistantTokenResponseBodyData {
	s.AuthToken = &v
	return s
}

func (s *GetTextbookAssistantTokenResponseBodyData) SetExpire(v int32) *GetTextbookAssistantTokenResponseBodyData {
	s.Expire = &v
	return s
}

type GetTextbookAssistantTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextbookAssistantTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextbookAssistantTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTextbookAssistantTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTextbookAssistantTokenResponse) SetHeaders(v map[string]*string) *GetTextbookAssistantTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTextbookAssistantTokenResponse) SetStatusCode(v int32) *GetTextbookAssistantTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextbookAssistantTokenResponse) SetBody(v *GetTextbookAssistantTokenResponseBody) *GetTextbookAssistantTokenResponse {
	s.Body = v
	return s
}

type ListTextbookAssistantArticleDetailsRequest struct {
	ArticleIdList []*string `json:"articleIdList,omitempty" xml:"articleIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsRequest) SetArticleIdList(v []*string) *ListTextbookAssistantArticleDetailsRequest {
	s.ArticleIdList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsRequest) SetAuthToken(v string) *ListTextbookAssistantArticleDetailsRequest {
	s.AuthToken = &v
	return s
}

type ListTextbookAssistantArticleDetailsResponseBody struct {
	Data []*ListTextbookAssistantArticleDetailsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetData(v []*ListTextbookAssistantArticleDetailsResponseBodyData) *ListTextbookAssistantArticleDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetErrCode(v string) *ListTextbookAssistantArticleDetailsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetErrMessage(v string) *ListTextbookAssistantArticleDetailsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantArticleDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetRequestId(v string) *ListTextbookAssistantArticleDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBody) SetSuccess(v bool) *ListTextbookAssistantArticleDetailsResponseBody {
	s.Success = &v
	return s
}

type ListTextbookAssistantArticleDetailsResponseBodyData struct {
	// example:
	//
	// 0c05700d4d9411efbe6e0c42a106bb02
	ArticleId    *string                                                            `json:"articleId,omitempty" xml:"articleId,omitempty"`
	QuestionList []*ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList `json:"questionList,omitempty" xml:"questionList,omitempty" type:"Repeated"`
	SceneList    []*ListTextbookAssistantArticleDetailsResponseBodyDataSceneList    `json:"sceneList,omitempty" xml:"sceneList,omitempty" type:"Repeated"`
	SentenceList []*ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	Target       *string                                                            `json:"target,omitempty" xml:"target,omitempty"`
	Theme        *ListTextbookAssistantArticleDetailsResponseBodyDataTheme          `json:"theme,omitempty" xml:"theme,omitempty" type:"Struct"`
	Topic        *ListTextbookAssistantArticleDetailsResponseBodyDataTopic          `json:"topic,omitempty" xml:"topic,omitempty" type:"Struct"`
	WordList     []*ListTextbookAssistantArticleDetailsResponseBodyDataWordList     `json:"wordList,omitempty" xml:"wordList,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetArticleId(v string) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.ArticleId = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetQuestionList(v []*ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.QuestionList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetSceneList(v []*ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.SceneList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetSentenceList(v []*ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.SentenceList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetTarget(v string) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.Target = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetTheme(v *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.Theme = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetTopic(v *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.Topic = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyData) SetWordList(v []*ListTextbookAssistantArticleDetailsResponseBodyDataWordList) *ListTextbookAssistantArticleDetailsResponseBodyData {
	s.WordList = v
	return s
}

type ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList struct {
	// example:
	//
	// I\\"m Mike Black
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty"`
	// example:
	//
	// From the book, how does Mike Black introduce himself?
	Question          *string `json:"question,omitempty" xml:"question,omitempty"`
	QuestionTranslate *string `json:"questionTranslate,omitempty" xml:"questionTranslate,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) SetAnswer(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList {
	s.Answer = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) SetQuestion(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList {
	s.Question = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList) SetQuestionTranslate(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataQuestionList {
	s.QuestionTranslate = &v
	return s
}

type ListTextbookAssistantArticleDetailsResponseBodyDataSceneList struct {
	// example:
	//
	// In the park, you introduce yourself to John and ask his name.
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 38cddd70509911efbe6e0c42a106bb02
	SceneId        *string   `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
	SceneImageList []*string `json:"sceneImageList,omitempty" xml:"sceneImageList,omitempty" type:"Repeated"`
	SceneTranslate *string   `json:"sceneTranslate,omitempty" xml:"sceneTranslate,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) SetScene(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList {
	s.Scene = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) SetSceneId(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList {
	s.SceneId = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) SetSceneImageList(v []*string) *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList {
	s.SceneImageList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList) SetSceneTranslate(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSceneList {
	s.SceneTranslate = &v
	return s
}

type ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList struct {
	SentenceAnalysis *string `json:"sentenceAnalysis,omitempty" xml:"sentenceAnalysis,omitempty"`
	// example:
	//
	// 4de677d2509811efbe6e0c42a106bb02
	SentenceId *string `json:"sentenceId,omitempty" xml:"sentenceId,omitempty"`
	// example:
	//
	// I\\"m Mike Black
	SentenceText *string `json:"sentenceText,omitempty" xml:"sentenceText,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) SetSentenceAnalysis(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList {
	s.SentenceAnalysis = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) SetSentenceId(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList {
	s.SentenceId = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList) SetSentenceText(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataSentenceList {
	s.SentenceText = &v
	return s
}

type ListTextbookAssistantArticleDetailsResponseBodyDataTheme struct {
	ThemeImageList []*string `json:"themeImageList,omitempty" xml:"themeImageList,omitempty" type:"Repeated"`
	ThemeName      *string   `json:"themeName,omitempty" xml:"themeName,omitempty"`
	// example:
	//
	// Self-awareness, self-management, self-improvement
	ThemeTranslate *string `json:"themeTranslate,omitempty" xml:"themeTranslate,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataTheme) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataTheme) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) SetThemeImageList(v []*string) *ListTextbookAssistantArticleDetailsResponseBodyDataTheme {
	s.ThemeImageList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) SetThemeName(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataTheme {
	s.ThemeName = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTheme) SetThemeTranslate(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataTheme {
	s.ThemeTranslate = &v
	return s
}

type ListTextbookAssistantArticleDetailsResponseBodyDataTopic struct {
	TopicImageList []*string `json:"topicImageList,omitempty" xml:"topicImageList,omitempty" type:"Repeated"`
	TopicName      *string   `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// example:
	//
	// Greetings and self-introduction
	TopicTranslate *string `json:"topicTranslate,omitempty" xml:"topicTranslate,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataTopic) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataTopic) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) SetTopicImageList(v []*string) *ListTextbookAssistantArticleDetailsResponseBodyDataTopic {
	s.TopicImageList = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) SetTopicName(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataTopic {
	s.TopicName = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataTopic) SetTopicTranslate(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataTopic {
	s.TopicTranslate = &v
	return s
}

type ListTextbookAssistantArticleDetailsResponseBodyDataWordList struct {
	WordAnalysis *string `json:"wordAnalysis,omitempty" xml:"wordAnalysis,omitempty"`
	// example:
	//
	// a94df134ed8c11eebe6e0c42a106bb02
	WordId *string `json:"wordId,omitempty" xml:"wordId,omitempty"`
	// example:
	//
	// nice
	WordText *string `json:"wordText,omitempty" xml:"wordText,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataWordList) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponseBodyDataWordList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataWordList) SetWordAnalysis(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataWordList {
	s.WordAnalysis = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataWordList) SetWordId(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataWordList {
	s.WordId = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponseBodyDataWordList) SetWordText(v string) *ListTextbookAssistantArticleDetailsResponseBodyDataWordList {
	s.WordText = &v
	return s
}

type ListTextbookAssistantArticleDetailsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantArticleDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantArticleDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticleDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticleDetailsResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantArticleDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponse) SetStatusCode(v int32) *ListTextbookAssistantArticleDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantArticleDetailsResponse) SetBody(v *ListTextbookAssistantArticleDetailsResponseBody) *ListTextbookAssistantArticleDetailsResponse {
	s.Body = v
	return s
}

type ListTextbookAssistantArticlesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tc_a893b8492c4be046cbc906c566aeb8c9
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 90aa861b4d9311efbe6e0c42a106bb02
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
}

func (s ListTextbookAssistantArticlesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticlesRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticlesRequest) SetAuthToken(v string) *ListTextbookAssistantArticlesRequest {
	s.AuthToken = &v
	return s
}

func (s *ListTextbookAssistantArticlesRequest) SetDirectoryId(v string) *ListTextbookAssistantArticlesRequest {
	s.DirectoryId = &v
	return s
}

type ListTextbookAssistantArticlesResponseBody struct {
	Data []*ListTextbookAssistantArticlesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 70412360-4272-571A-827D-84C2C07C450F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTextbookAssistantArticlesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticlesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticlesResponseBody) SetData(v []*ListTextbookAssistantArticlesResponseBodyData) *ListTextbookAssistantArticlesResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBody) SetErrCode(v string) *ListTextbookAssistantArticlesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBody) SetErrMessage(v string) *ListTextbookAssistantArticlesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantArticlesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBody) SetRequestId(v string) *ListTextbookAssistantArticlesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponseBody) SetSuccess(v bool) *ListTextbookAssistantArticlesResponseBody {
	s.Success = &v
	return s
}

type ListTextbookAssistantArticlesResponseBodyData struct {
	// example:
	//
	// 0c05700d4d9411efbe6e0c42a106bb02
	ArticleId *string `json:"articleId,omitempty" xml:"articleId,omitempty"`
}

func (s ListTextbookAssistantArticlesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticlesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticlesResponseBodyData) SetArticleId(v string) *ListTextbookAssistantArticlesResponseBodyData {
	s.ArticleId = &v
	return s
}

type ListTextbookAssistantArticlesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantArticlesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantArticlesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantArticlesResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantArticlesResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantArticlesResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantArticlesResponse) SetStatusCode(v int32) *ListTextbookAssistantArticlesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantArticlesResponse) SetBody(v *ListTextbookAssistantArticlesResponseBody) *ListTextbookAssistantArticlesResponse {
	s.Body = v
	return s
}

type ListTextbookAssistantBookDirectoriesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 231698
	BookId *string `json:"bookId,omitempty" xml:"bookId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesRequest) SetAuthToken(v string) *ListTextbookAssistantBookDirectoriesRequest {
	s.AuthToken = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesRequest) SetBookId(v string) *ListTextbookAssistantBookDirectoriesRequest {
	s.BookId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesRequest) SetScenario(v string) *ListTextbookAssistantBookDirectoriesRequest {
	s.Scenario = &v
	return s
}

type ListTextbookAssistantBookDirectoriesResponseBody struct {
	Data *ListTextbookAssistantBookDirectoriesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// B_USER_NOT_FOUND_EXCEPTION
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0A5E9849-A2F0-551D-A7D8-1A8118557BAB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetData(v *ListTextbookAssistantBookDirectoriesResponseBodyData) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetErrCode(v string) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetErrMessage(v string) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetRequestId(v string) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBody) SetSuccess(v bool) *ListTextbookAssistantBookDirectoriesResponseBody {
	s.Success = &v
	return s
}

type ListTextbookAssistantBookDirectoriesResponseBodyData struct {
	DirectoryTree []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree `json:"directoryTree,omitempty" xml:"directoryTree,omitempty" type:"Repeated"`
	EditionInfo   *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo     `json:"editionInfo,omitempty" xml:"editionInfo,omitempty" type:"Struct"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyData) SetDirectoryTree(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) *ListTextbookAssistantBookDirectoriesResponseBodyData {
	s.DirectoryTree = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyData) SetEditionInfo(v *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) *ListTextbookAssistantBookDirectoriesResponseBodyData {
	s.EditionInfo = v
	return s
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree struct {
	// example:
	//
	// 05758807ed8e11eebe6e0c42a106bb02
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// example:
	//
	// 2 Jobs
	DirectoryName *string                                                                   `json:"directoryName,omitempty" xml:"directoryName,omitempty"`
	Topic         []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic `json:"topic,omitempty" xml:"topic,omitempty" type:"Repeated"`
	Unit          []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit  `json:"unit,omitempty" xml:"unit,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) SetDirectoryId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree {
	s.DirectoryId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) SetDirectoryName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree {
	s.DirectoryName = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) SetTopic(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree {
	s.Topic = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree) SetUnit(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTree {
	s.Unit = v
	return s
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic struct {
	// example:
	//
	// 1323
	LabelId   *string `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) SetLabelId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic {
	s.LabelId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic) SetLabelName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeTopic {
	s.LabelName = &v
	return s
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit struct {
	DirectoryId   *string                                                                         `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	DirectoryName *string                                                                         `json:"directoryName,omitempty" xml:"directoryName,omitempty"`
	Section       []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection `json:"section,omitempty" xml:"section,omitempty" type:"Repeated"`
	Topic         []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic   `json:"topic,omitempty" xml:"topic,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) SetDirectoryId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit {
	s.DirectoryId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) SetDirectoryName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit {
	s.DirectoryName = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) SetSection(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit {
	s.Section = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit) SetTopic(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnit {
	s.Topic = v
	return s
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection struct {
	Children      interface{}                                                                          `json:"children,omitempty" xml:"children,omitempty"`
	DirectoryId   *string                                                                              `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	DirectoryName *string                                                                              `json:"directoryName,omitempty" xml:"directoryName,omitempty"`
	Topic         []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic `json:"topic,omitempty" xml:"topic,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) SetChildren(v interface{}) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection {
	s.Children = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) SetDirectoryId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection {
	s.DirectoryId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) SetDirectoryName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection {
	s.DirectoryName = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection) SetTopic(v []*ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSection {
	s.Topic = v
	return s
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic struct {
	LabelId   *string `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) SetLabelId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic {
	s.LabelId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic) SetLabelName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitSectionTopic {
	s.LabelName = &v
	return s
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic struct {
	LabelId   *string `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) SetLabelId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic {
	s.LabelId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic) SetLabelName(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataDirectoryTreeUnitTopic {
	s.LabelName = &v
	return s
}

type ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo struct {
	// example:
	//
	// 55857
	BookId *string `json:"bookId,omitempty" xml:"bookId,omitempty"`
	// example:
	//
	// 1
	BookVolume *string `json:"bookVolume,omitempty" xml:"bookVolume,omitempty"`
	// example:
	//
	// 2010-1(2)
	Edition *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// example:
	//
	// 3
	Grade *string `json:"grade,omitempty" xml:"grade,omitempty"`
	// example:
	//
	// 2019-1(10)
	Impression *string `json:"impression,omitempty" xml:"impression,omitempty"`
	// example:
	//
	// 9787544413695
	Isbn      *string `json:"isbn,omitempty" xml:"isbn,omitempty"`
	Publisher *string `json:"publisher,omitempty" xml:"publisher,omitempty"`
	// example:
	//
	// ENGLISH
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetBookId(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.BookId = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetBookVolume(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.BookVolume = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetEdition(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Edition = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetGrade(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Grade = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetImpression(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Impression = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetIsbn(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Isbn = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetPublisher(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Publisher = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetSubject(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Subject = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo) SetVersion(v string) *ListTextbookAssistantBookDirectoriesResponseBodyDataEditionInfo {
	s.Version = &v
	return s
}

type ListTextbookAssistantBookDirectoriesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantBookDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantBookDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBookDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBookDirectoriesResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantBookDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponse) SetStatusCode(v int32) *ListTextbookAssistantBookDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantBookDirectoriesResponse) SetBody(v *ListTextbookAssistantBookDirectoriesResponseBody) *ListTextbookAssistantBookDirectoriesResponse {
	s.Body = v
	return s
}

type ListTextbookAssistantBooksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tc_197bf5bb81889cc79eb51ae9b8c0cea3
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// example:
	//
	// 231698
	BookId *string `json:"bookId,omitempty" xml:"bookId,omitempty"`
	// example:
	//
	// 1
	Grade *string `json:"grade,omitempty" xml:"grade,omitempty"`
	// example:
	//
	// 20
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	Page    *string `json:"page,omitempty" xml:"page,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 1
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s ListTextbookAssistantBooksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBooksRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksRequest) SetAuthToken(v string) *ListTextbookAssistantBooksRequest {
	s.AuthToken = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetBookId(v string) *ListTextbookAssistantBooksRequest {
	s.BookId = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetGrade(v string) *ListTextbookAssistantBooksRequest {
	s.Grade = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetMaxResults(v string) *ListTextbookAssistantBooksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetPage(v string) *ListTextbookAssistantBooksRequest {
	s.Page = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetVersion(v string) *ListTextbookAssistantBooksRequest {
	s.Version = &v
	return s
}

func (s *ListTextbookAssistantBooksRequest) SetVolume(v string) *ListTextbookAssistantBooksRequest {
	s.Volume = &v
	return s
}

type ListTextbookAssistantBooksResponseBody struct {
	Data *ListTextbookAssistantBooksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// B_USER_NOT_FOUND_EXCEPTION
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B695B377-7029-5805-9DE2-1AAE06C1BF6B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTextbookAssistantBooksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBooksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksResponseBody) SetData(v *ListTextbookAssistantBooksResponseBodyData) *ListTextbookAssistantBooksResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantBooksResponseBody) SetErrCode(v string) *ListTextbookAssistantBooksResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBody) SetErrMessage(v string) *ListTextbookAssistantBooksResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantBooksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBody) SetRequestId(v string) *ListTextbookAssistantBooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBody) SetSuccess(v bool) *ListTextbookAssistantBooksResponseBody {
	s.Success = &v
	return s
}

type ListTextbookAssistantBooksResponseBodyData struct {
	BookList       []*ListTextbookAssistantBooksResponseBodyDataBookList     `json:"bookList,omitempty" xml:"bookList,omitempty" type:"Repeated"`
	PaginationData *ListTextbookAssistantBooksResponseBodyDataPaginationData `json:"paginationData,omitempty" xml:"paginationData,omitempty" type:"Struct"`
}

func (s ListTextbookAssistantBooksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBooksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksResponseBodyData) SetBookList(v []*ListTextbookAssistantBooksResponseBodyDataBookList) *ListTextbookAssistantBooksResponseBodyData {
	s.BookList = v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyData) SetPaginationData(v *ListTextbookAssistantBooksResponseBodyDataPaginationData) *ListTextbookAssistantBooksResponseBodyData {
	s.PaginationData = v
	return s
}

type ListTextbookAssistantBooksResponseBodyDataBookList struct {
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// example:
	//
	// 231698
	BookId   *string `json:"bookId,omitempty" xml:"bookId,omitempty"`
	BookName *string `json:"bookName,omitempty" xml:"bookName,omitempty"`
	// example:
	//
	// null
	CoverImage *string `json:"coverImage,omitempty" xml:"coverImage,omitempty"`
	// example:
	//
	// 2024-7（1）
	Edition *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// example:
	//
	// 3
	Grade *string `json:"grade,omitempty" xml:"grade,omitempty"`
	// example:
	//
	// 2024-7（1）
	Impression *string `json:"impression,omitempty" xml:"impression,omitempty"`
	// example:
	//
	// 9787107382505
	Isbn      *string `json:"isbn,omitempty" xml:"isbn,omitempty"`
	Publisher *string `json:"publisher,omitempty" xml:"publisher,omitempty"`
	// example:
	//
	// ENGLISH
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 0
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s ListTextbookAssistantBooksResponseBodyDataBookList) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBooksResponseBodyDataBookList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetAuthor(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Author = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetBookId(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.BookId = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetBookName(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.BookName = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetCoverImage(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.CoverImage = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetEdition(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Edition = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetGrade(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Grade = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetImpression(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Impression = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetIsbn(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Isbn = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetPublisher(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Publisher = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetSubject(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Subject = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetVersion(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Version = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataBookList) SetVolume(v string) *ListTextbookAssistantBooksResponseBodyDataBookList {
	s.Volume = &v
	return s
}

type ListTextbookAssistantBooksResponseBodyDataPaginationData struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTextbookAssistantBooksResponseBodyDataPaginationData) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBooksResponseBodyDataPaginationData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksResponseBodyDataPaginationData) SetCurrentPage(v int32) *ListTextbookAssistantBooksResponseBodyDataPaginationData {
	s.CurrentPage = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataPaginationData) SetMaxResults(v int32) *ListTextbookAssistantBooksResponseBodyDataPaginationData {
	s.MaxResults = &v
	return s
}

func (s *ListTextbookAssistantBooksResponseBodyDataPaginationData) SetTotalCount(v int32) *ListTextbookAssistantBooksResponseBodyDataPaginationData {
	s.TotalCount = &v
	return s
}

type ListTextbookAssistantBooksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantBooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantBooksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantBooksResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantBooksResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantBooksResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantBooksResponse) SetStatusCode(v int32) *ListTextbookAssistantBooksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantBooksResponse) SetBody(v *ListTextbookAssistantBooksResponseBody) *ListTextbookAssistantBooksResponse {
	s.Body = v
	return s
}

type ListTextbookAssistantGradeVolumesRequest struct {
	// example:
	//
	// tc_197bf5bb81889cc79eb51ae9b8c0cea3
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ListTextbookAssistantGradeVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantGradeVolumesRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantGradeVolumesRequest) SetAuthToken(v string) *ListTextbookAssistantGradeVolumesRequest {
	s.AuthToken = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesRequest) SetScenario(v string) *ListTextbookAssistantGradeVolumesRequest {
	s.Scenario = &v
	return s
}

type ListTextbookAssistantGradeVolumesResponseBody struct {
	Data []*ListTextbookAssistantGradeVolumesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F3B1AAF2-3041-5AA7-A352-BD5F998FA465
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTextbookAssistantGradeVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantGradeVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetData(v []*ListTextbookAssistantGradeVolumesResponseBodyData) *ListTextbookAssistantGradeVolumesResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetErrCode(v string) *ListTextbookAssistantGradeVolumesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetErrMessage(v string) *ListTextbookAssistantGradeVolumesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantGradeVolumesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetRequestId(v string) *ListTextbookAssistantGradeVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetSuccess(v bool) *ListTextbookAssistantGradeVolumesResponseBody {
	s.Success = &v
	return s
}

type ListTextbookAssistantGradeVolumesResponseBodyData struct {
	GradeVolumes []*ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes `json:"gradeVolumes,omitempty" xml:"gradeVolumes,omitempty" type:"Repeated"`
	// example:
	//
	// 人教版
	TextbookVersion *string `json:"textbookVersion,omitempty" xml:"textbookVersion,omitempty"`
}

func (s ListTextbookAssistantGradeVolumesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantGradeVolumesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyData) SetGradeVolumes(v []*ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) *ListTextbookAssistantGradeVolumesResponseBodyData {
	s.GradeVolumes = v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyData) SetTextbookVersion(v string) *ListTextbookAssistantGradeVolumesResponseBodyData {
	s.TextbookVersion = &v
	return s
}

type ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes struct {
	// example:
	//
	// 3
	Grade *string `json:"grade,omitempty" xml:"grade,omitempty"`
	// example:
	//
	// 1
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) SetGrade(v string) *ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes {
	s.Grade = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) SetVolume(v string) *ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes {
	s.Volume = &v
	return s
}

type ListTextbookAssistantGradeVolumesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantGradeVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantGradeVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantGradeVolumesResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantGradeVolumesResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantGradeVolumesResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponse) SetStatusCode(v int32) *ListTextbookAssistantGradeVolumesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponse) SetBody(v *ListTextbookAssistantGradeVolumesResponseBody) *ListTextbookAssistantGradeVolumesResponse {
	s.Body = v
	return s
}

type ListTextbookAssistantSceneDetailsRequest struct {
	// example:
	//
	// tc_e6dc70c890866f4028ca685b6fa29874
	AuthToken   *string   `json:"authToken,omitempty" xml:"authToken,omitempty"`
	SceneIdList []*string `json:"sceneIdList,omitempty" xml:"sceneIdList,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantSceneDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsRequest) SetAuthToken(v string) *ListTextbookAssistantSceneDetailsRequest {
	s.AuthToken = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsRequest) SetSceneIdList(v []*string) *ListTextbookAssistantSceneDetailsRequest {
	s.SceneIdList = v
	return s
}

type ListTextbookAssistantSceneDetailsResponseBody struct {
	Data []*ListTextbookAssistantSceneDetailsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetData(v []*ListTextbookAssistantSceneDetailsResponseBodyData) *ListTextbookAssistantSceneDetailsResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetErrCode(v string) *ListTextbookAssistantSceneDetailsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetErrMessage(v string) *ListTextbookAssistantSceneDetailsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantSceneDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetRequestId(v string) *ListTextbookAssistantSceneDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBody) SetSuccess(v bool) *ListTextbookAssistantSceneDetailsResponseBody {
	s.Success = &v
	return s
}

type ListTextbookAssistantSceneDetailsResponseBodyData struct {
	RoleList []*ListTextbookAssistantSceneDetailsResponseBodyDataRoleList `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
	// example:
	//
	// At school, Carl sees a photo and asks you about your family.
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 38c41b7b509911efbe6e0c42a106bb02
	SceneId        *string                                                           `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
	SceneImageList []*string                                                         `json:"sceneImageList,omitempty" xml:"sceneImageList,omitempty" type:"Repeated"`
	SceneTaskList  []*ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList `json:"sceneTaskList,omitempty" xml:"sceneTaskList,omitempty" type:"Repeated"`
	SceneTranslate *string                                                           `json:"sceneTranslate,omitempty" xml:"sceneTranslate,omitempty"`
	SentenceList   []*ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList  `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	Target   *string                                                      `json:"target,omitempty" xml:"target,omitempty"`
	Theme    *ListTextbookAssistantSceneDetailsResponseBodyDataTheme      `json:"theme,omitempty" xml:"theme,omitempty" type:"Struct"`
	Topic    *ListTextbookAssistantSceneDetailsResponseBodyDataTopic      `json:"topic,omitempty" xml:"topic,omitempty" type:"Struct"`
	WordList []*ListTextbookAssistantSceneDetailsResponseBodyDataWordList `json:"wordList,omitempty" xml:"wordList,omitempty" type:"Repeated"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetRoleList(v []*ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.RoleList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetScene(v string) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.Scene = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetSceneId(v string) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.SceneId = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetSceneImageList(v []*string) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.SceneImageList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetSceneTaskList(v []*ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.SceneTaskList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetSceneTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.SceneTranslate = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetSentenceList(v []*ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.SentenceList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetTarget(v string) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.Target = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetTheme(v *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.Theme = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetTopic(v *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.Topic = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyData) SetWordList(v []*ListTextbookAssistantSceneDetailsResponseBodyDataWordList) *ListTextbookAssistantSceneDetailsResponseBodyData {
	s.WordList = v
	return s
}

type ListTextbookAssistantSceneDetailsResponseBodyDataRoleList struct {
	// example:
	//
	// Carl, a curious boy
	Introduction          *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	IntroductionTranslate *string `json:"introductionTranslate,omitempty" xml:"introductionTranslate,omitempty"`
	// example:
	//
	// Hi Noah, who is that in the photo?
	Promoting          *string `json:"promoting,omitempty" xml:"promoting,omitempty"`
	PromotingTranslate *string `json:"promotingTranslate,omitempty" xml:"promotingTranslate,omitempty"`
	// example:
	//
	// Carl
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// example:
	//
	// Carl
	RoleNameTranslate *string `json:"roleNameTranslate,omitempty" xml:"roleNameTranslate,omitempty"`
	// example:
	//
	// 0
	RoleType *string `json:"roleType,omitempty" xml:"roleType,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetIntroduction(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.Introduction = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetIntroductionTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.IntroductionTranslate = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetPromoting(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.Promoting = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetPromotingTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.PromotingTranslate = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetRoleName(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.RoleName = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetRoleNameTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.RoleNameTranslate = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList) SetRoleType(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataRoleList {
	s.RoleType = &v
	return s
}

type ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList struct {
	// example:
	//
	// Say that this is your dad\\"s brother.
	SceneTask          *string `json:"sceneTask,omitempty" xml:"sceneTask,omitempty"`
	SceneTaskTranslate *string `json:"sceneTaskTranslate,omitempty" xml:"sceneTaskTranslate,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) SetSceneTask(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList {
	s.SceneTask = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList) SetSceneTaskTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataSceneTaskList {
	s.SceneTaskTranslate = &v
	return s
}

type ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList struct {
	SentenceAnalysis *string `json:"sentenceAnalysis,omitempty" xml:"sentenceAnalysis,omitempty"`
	// example:
	//
	// a774c6d09c4511eebe6e0c42a106bb02
	SentenceId *string `json:"sentenceId,omitempty" xml:"sentenceId,omitempty"`
	// example:
	//
	// Is this your sister?
	SentenceText *string `json:"sentenceText,omitempty" xml:"sentenceText,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) SetSentenceAnalysis(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList {
	s.SentenceAnalysis = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) SetSentenceId(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList {
	s.SentenceId = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList) SetSentenceText(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataSentenceList {
	s.SentenceText = &v
	return s
}

type ListTextbookAssistantSceneDetailsResponseBodyDataTheme struct {
	ThemeImageList []*string `json:"themeImageList,omitempty" xml:"themeImageList,omitempty" type:"Repeated"`
	ThemeName      *string   `json:"themeName,omitempty" xml:"themeName,omitempty"`
	// example:
	//
	// Family and family life
	ThemeTranslate *string `json:"themeTranslate,omitempty" xml:"themeTranslate,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataTheme) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataTheme) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) SetThemeImageList(v []*string) *ListTextbookAssistantSceneDetailsResponseBodyDataTheme {
	s.ThemeImageList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) SetThemeName(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataTheme {
	s.ThemeName = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTheme) SetThemeTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataTheme {
	s.ThemeTranslate = &v
	return s
}

type ListTextbookAssistantSceneDetailsResponseBodyDataTopic struct {
	TopicImageList []*string `json:"topicImageList,omitempty" xml:"topicImageList,omitempty" type:"Repeated"`
	TopicName      *string   `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// example:
	//
	// Introducing family members
	TopicTranslate *string `json:"topicTranslate,omitempty" xml:"topicTranslate,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataTopic) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataTopic) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) SetTopicImageList(v []*string) *ListTextbookAssistantSceneDetailsResponseBodyDataTopic {
	s.TopicImageList = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) SetTopicName(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataTopic {
	s.TopicName = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataTopic) SetTopicTranslate(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataTopic {
	s.TopicTranslate = &v
	return s
}

type ListTextbookAssistantSceneDetailsResponseBodyDataWordList struct {
	WordAnalysis *string `json:"wordAnalysis,omitempty" xml:"wordAnalysis,omitempty"`
	// example:
	//
	// a94c3337ed8c11eebe6e0c42a106bb02
	WordId *string `json:"wordId,omitempty" xml:"wordId,omitempty"`
	// example:
	//
	// family
	WordText *string `json:"wordText,omitempty" xml:"wordText,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataWordList) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponseBodyDataWordList) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataWordList) SetWordAnalysis(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataWordList {
	s.WordAnalysis = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataWordList) SetWordId(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataWordList {
	s.WordId = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponseBodyDataWordList) SetWordText(v string) *ListTextbookAssistantSceneDetailsResponseBodyDataWordList {
	s.WordText = &v
	return s
}

type ListTextbookAssistantSceneDetailsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTextbookAssistantSceneDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextbookAssistantSceneDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTextbookAssistantSceneDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantSceneDetailsResponse) SetHeaders(v map[string]*string) *ListTextbookAssistantSceneDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponse) SetStatusCode(v int32) *ListTextbookAssistantSceneDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextbookAssistantSceneDetailsResponse) SetBody(v *ListTextbookAssistantSceneDetailsResponseBody) *ListTextbookAssistantSceneDetailsResponse {
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

type QueryApplicationAccessIdRequest struct {
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
}

func (s QueryApplicationAccessIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationAccessIdRequest) GoString() string {
	return s.String()
}

func (s *QueryApplicationAccessIdRequest) SetApplicationAccessId(v string) *QueryApplicationAccessIdRequest {
	s.ApplicationAccessId = &v
	return s
}

type QueryApplicationAccessIdResponseBody struct {
	// example:
	//
	// []
	Data *QueryApplicationAccessIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s QueryApplicationAccessIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationAccessIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryApplicationAccessIdResponseBody) SetData(v *QueryApplicationAccessIdResponseBodyData) *QueryApplicationAccessIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryApplicationAccessIdResponseBody) SetErrCode(v string) *QueryApplicationAccessIdResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBody) SetErrMessage(v string) *QueryApplicationAccessIdResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBody) SetHttpStatusCode(v int32) *QueryApplicationAccessIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBody) SetRequestId(v string) *QueryApplicationAccessIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBody) SetSuccess(v bool) *QueryApplicationAccessIdResponseBody {
	s.Success = &v
	return s
}

type QueryApplicationAccessIdResponseBodyData struct {
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// MyAppSecret
	ApplicationAccessSecret *string `json:"applicationAccessSecret,omitempty" xml:"applicationAccessSecret,omitempty"`
}

func (s QueryApplicationAccessIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationAccessIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryApplicationAccessIdResponseBodyData) SetApplicationAccessId(v string) *QueryApplicationAccessIdResponseBodyData {
	s.ApplicationAccessId = &v
	return s
}

func (s *QueryApplicationAccessIdResponseBodyData) SetApplicationAccessSecret(v string) *QueryApplicationAccessIdResponseBodyData {
	s.ApplicationAccessSecret = &v
	return s
}

type QueryApplicationAccessIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryApplicationAccessIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryApplicationAccessIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryApplicationAccessIdResponse) GoString() string {
	return s.String()
}

func (s *QueryApplicationAccessIdResponse) SetHeaders(v map[string]*string) *QueryApplicationAccessIdResponse {
	s.Headers = v
	return s
}

func (s *QueryApplicationAccessIdResponse) SetStatusCode(v int32) *QueryApplicationAccessIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryApplicationAccessIdResponse) SetBody(v *QueryApplicationAccessIdResponseBody) *QueryApplicationAccessIdResponse {
	s.Body = v
	return s
}

type QueryProjectRequest struct {
	// example:
	//
	// 123
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s QueryProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectRequest) GoString() string {
	return s.String()
}

func (s *QueryProjectRequest) SetProjectId(v string) *QueryProjectRequest {
	s.ProjectId = &v
	return s
}

type QueryProjectResponseBody struct {
	// example:
	//
	// []
	Data *QueryProjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s QueryProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBody) SetData(v *QueryProjectResponseBodyData) *QueryProjectResponseBody {
	s.Data = v
	return s
}

func (s *QueryProjectResponseBody) SetErrCode(v string) *QueryProjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryProjectResponseBody) SetErrMessage(v string) *QueryProjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryProjectResponseBody) SetHttpStatusCode(v int32) *QueryProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryProjectResponseBody) SetRequestId(v string) *QueryProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProjectResponseBody) SetSuccess(v bool) *QueryProjectResponseBody {
	s.Success = &v
	return s
}

type QueryProjectResponseBodyData struct {
	// example:
	//
	// 2024-11-01T13:40:53Z
	CreateTime  *string                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectApps []*QueryProjectResponseBodyDataProjectApps `json:"ProjectApps,omitempty" xml:"ProjectApps,omitempty" type:"Repeated"`
	// example:
	//
	// 67055
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// MyProject
	ProjectName *string                                   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectSDK  []*QueryProjectResponseBodyDataProjectSDK `json:"ProjectSDK,omitempty" xml:"ProjectSDK,omitempty" type:"Repeated"`
	// example:
	//
	// WebApplication
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
}

func (s QueryProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyData) SetCreateTime(v string) *QueryProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectResponseBodyData) SetProjectApps(v []*QueryProjectResponseBodyDataProjectApps) *QueryProjectResponseBodyData {
	s.ProjectApps = v
	return s
}

func (s *QueryProjectResponseBodyData) SetProjectId(v string) *QueryProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *QueryProjectResponseBodyData) SetProjectName(v string) *QueryProjectResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *QueryProjectResponseBodyData) SetProjectSDK(v []*QueryProjectResponseBodyDataProjectSDK) *QueryProjectResponseBodyData {
	s.ProjectSDK = v
	return s
}

func (s *QueryProjectResponseBodyData) SetProjectType(v string) *QueryProjectResponseBodyData {
	s.ProjectType = &v
	return s
}

type QueryProjectResponseBodyDataProjectApps struct {
	ApplicationAccessIds []*QueryProjectResponseBodyDataProjectAppsApplicationAccessIds `json:"ApplicationAccessIds,omitempty" xml:"ApplicationAccessIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2144
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 159
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s QueryProjectResponseBodyDataProjectApps) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseBodyDataProjectApps) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyDataProjectApps) SetApplicationAccessIds(v []*QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) *QueryProjectResponseBodyDataProjectApps {
	s.ApplicationAccessIds = v
	return s
}

func (s *QueryProjectResponseBodyDataProjectApps) SetId(v string) *QueryProjectResponseBodyDataProjectApps {
	s.Id = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectApps) SetProjectId(v string) *QueryProjectResponseBodyDataProjectApps {
	s.ProjectId = &v
	return s
}

type QueryProjectResponseBodyDataProjectAppsApplicationAccessIds struct {
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// MyAppSecret
	ApplicationAccessSecret *string `json:"applicationAccessSecret,omitempty" xml:"applicationAccessSecret,omitempty"`
}

func (s QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessId(v string) *QueryProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessId = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessSecret(v string) *QueryProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessSecret = &v
	return s
}

type QueryProjectResponseBodyDataProjectSDK struct {
	// example:
	//
	// 2024-11-01T13:40:53Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// http://demo.com/demo
	DemoUrl    *string `json:"DemoUrl,omitempty" xml:"DemoUrl,omitempty"`
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// example:
	//
	// JAVA
	DevelopLanguage *string `json:"DevelopLanguage,omitempty" xml:"DevelopLanguage,omitempty"`
	// example:
	//
	// http://demo.com/doc
	DocUrl *string `json:"DocUrl,omitempty" xml:"DocUrl,omitempty"`
	// example:
	//
	// JSSDK
	SdkName *string `json:"SdkName,omitempty" xml:"SdkName,omitempty"`
	// example:
	//
	// http://demo.com/sdk.zip
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	// example:
	//
	// 5.1.0
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
}

func (s QueryProjectResponseBodyDataProjectSDK) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseBodyDataProjectSDK) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetCreateTime(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetDemoUrl(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.DemoUrl = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetDeployMode(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.DeployMode = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetDevelopLanguage(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.DevelopLanguage = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetDocUrl(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.DocUrl = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetSdkName(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.SdkName = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetSdkUrl(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.SdkUrl = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetSdkVersion(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.SdkVersion = &v
	return s
}

type QueryProjectResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponse) GoString() string {
	return s.String()
}

func (s *QueryProjectResponse) SetHeaders(v map[string]*string) *QueryProjectResponse {
	s.Headers = v
	return s
}

func (s *QueryProjectResponse) SetStatusCode(v int32) *QueryProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProjectResponse) SetBody(v *QueryProjectResponseBody) *QueryProjectResponse {
	s.Body = v
	return s
}

type QueryProjectListResponseBody struct {
	// example:
	//
	// []
	Data []*QueryProjectListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s QueryProjectListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponseBody) SetData(v []*QueryProjectListResponseBodyData) *QueryProjectListResponseBody {
	s.Data = v
	return s
}

func (s *QueryProjectListResponseBody) SetErrCode(v string) *QueryProjectListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryProjectListResponseBody) SetErrMessage(v string) *QueryProjectListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryProjectListResponseBody) SetHttpStatusCode(v int32) *QueryProjectListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryProjectListResponseBody) SetRequestId(v string) *QueryProjectListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProjectListResponseBody) SetSuccess(v bool) *QueryProjectListResponseBody {
	s.Success = &v
	return s
}

type QueryProjectListResponseBodyData struct {
	// example:
	//
	// 2025-02-18 12:10:22
	CreateTime  *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectApps []*QueryProjectListResponseBodyDataProjectApps `json:"ProjectApps,omitempty" xml:"ProjectApps,omitempty" type:"Repeated"`
	// example:
	//
	// 4910
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// MyProject
	ProjectName *string                                       `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectSDK  []*QueryProjectListResponseBodyDataProjectSDK `json:"ProjectSDK,omitempty" xml:"ProjectSDK,omitempty" type:"Repeated"`
	// example:
	//
	// WebApplication
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
}

func (s QueryProjectListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponseBodyData) SetCreateTime(v string) *QueryProjectListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectListResponseBodyData) SetProjectApps(v []*QueryProjectListResponseBodyDataProjectApps) *QueryProjectListResponseBodyData {
	s.ProjectApps = v
	return s
}

func (s *QueryProjectListResponseBodyData) SetProjectId(v string) *QueryProjectListResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *QueryProjectListResponseBodyData) SetProjectName(v string) *QueryProjectListResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *QueryProjectListResponseBodyData) SetProjectSDK(v []*QueryProjectListResponseBodyDataProjectSDK) *QueryProjectListResponseBodyData {
	s.ProjectSDK = v
	return s
}

func (s *QueryProjectListResponseBodyData) SetProjectType(v string) *QueryProjectListResponseBodyData {
	s.ProjectType = &v
	return s
}

type QueryProjectListResponseBodyDataProjectApps struct {
	ApplicationAccessIds []*QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds `json:"ApplicationAccessIds,omitempty" xml:"ApplicationAccessIds,omitempty" type:"Repeated"`
	// example:
	//
	// 4700
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 4747
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s QueryProjectListResponseBodyDataProjectApps) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectListResponseBodyDataProjectApps) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponseBodyDataProjectApps) SetApplicationAccessIds(v []*QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) *QueryProjectListResponseBodyDataProjectApps {
	s.ApplicationAccessIds = v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectApps) SetId(v string) *QueryProjectListResponseBodyDataProjectApps {
	s.Id = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectApps) SetProjectId(v string) *QueryProjectListResponseBodyDataProjectApps {
	s.ProjectId = &v
	return s
}

type QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds struct {
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// MyAppSecret
	ApplicationAccessSecret *string `json:"applicationAccessSecret,omitempty" xml:"applicationAccessSecret,omitempty"`
}

func (s QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessId(v string) *QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessId = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessSecret(v string) *QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessSecret = &v
	return s
}

type QueryProjectListResponseBodyDataProjectSDK struct {
	// example:
	//
	// 2024-07-16T08:23:19Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// http://demo.com/demo
	DemoUrl    *string `json:"DemoUrl,omitempty" xml:"DemoUrl,omitempty"`
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// example:
	//
	// JAVA
	DevelopLanguage *string `json:"DevelopLanguage,omitempty" xml:"DevelopLanguage,omitempty"`
	// example:
	//
	// http://demo.com/doc
	DocUrl *string `json:"DocUrl,omitempty" xml:"DocUrl,omitempty"`
	// example:
	//
	// GO AUTH
	SdkName *string `json:"SdkName,omitempty" xml:"SdkName,omitempty"`
	// example:
	//
	// http://demo.com/sdk.zip
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	// example:
	//
	// .3.52
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
}

func (s QueryProjectListResponseBodyDataProjectSDK) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectListResponseBodyDataProjectSDK) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetCreateTime(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetDemoUrl(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.DemoUrl = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetDeployMode(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.DeployMode = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetDevelopLanguage(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.DevelopLanguage = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetDocUrl(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.DocUrl = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetSdkName(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.SdkName = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetSdkUrl(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.SdkUrl = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetSdkVersion(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.SdkVersion = &v
	return s
}

type QueryProjectListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProjectListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProjectListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectListResponse) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponse) SetHeaders(v map[string]*string) *QueryProjectListResponse {
	s.Headers = v
	return s
}

func (s *QueryProjectListResponse) SetStatusCode(v int32) *QueryProjectListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProjectListResponse) SetBody(v *QueryProjectListResponseBody) *QueryProjectListResponse {
	s.Body = v
	return s
}

type QueryPurchasedServiceResponseBody struct {
	// example:
	//
	// []
	Data []*AliyunConsoleServiceInfoDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s QueryPurchasedServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedServiceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPurchasedServiceResponseBody) SetData(v []*AliyunConsoleServiceInfoDTO) *QueryPurchasedServiceResponseBody {
	s.Data = v
	return s
}

func (s *QueryPurchasedServiceResponseBody) SetErrCode(v string) *QueryPurchasedServiceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryPurchasedServiceResponseBody) SetErrMessage(v string) *QueryPurchasedServiceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryPurchasedServiceResponseBody) SetHttpStatusCode(v int32) *QueryPurchasedServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryPurchasedServiceResponseBody) SetRequestId(v string) *QueryPurchasedServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPurchasedServiceResponseBody) SetSuccess(v bool) *QueryPurchasedServiceResponseBody {
	s.Success = &v
	return s
}

type QueryPurchasedServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPurchasedServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPurchasedServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedServiceResponse) GoString() string {
	return s.String()
}

func (s *QueryPurchasedServiceResponse) SetHeaders(v map[string]*string) *QueryPurchasedServiceResponse {
	s.Headers = v
	return s
}

func (s *QueryPurchasedServiceResponse) SetStatusCode(v int32) *QueryPurchasedServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPurchasedServiceResponse) SetBody(v *QueryPurchasedServiceResponseBody) *QueryPurchasedServiceResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	// example:
	//
	// 123
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// MyProject
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetProjectId(v string) *UpdateProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectName(v string) *UpdateProjectRequest {
	s.ProjectName = &v
	return s
}

type UpdateProjectResponseBody struct {
	// example:
	//
	// []
	Data *UpdateProjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetData(v *UpdateProjectResponseBodyData) *UpdateProjectResponseBody {
	s.Data = v
	return s
}

func (s *UpdateProjectResponseBody) SetErrCode(v string) *UpdateProjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateProjectResponseBody) SetErrMessage(v string) *UpdateProjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateProjectResponseBody) SetHttpStatusCode(v int32) *UpdateProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetSuccess(v bool) *UpdateProjectResponseBody {
	s.Success = &v
	return s
}

type UpdateProjectResponseBodyData struct {
	// example:
	//
	// 2024-12-10T02:07:16Z
	CreateTime  *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectApps []*UpdateProjectResponseBodyDataProjectApps `json:"ProjectApps,omitempty" xml:"ProjectApps,omitempty" type:"Repeated"`
	// example:
	//
	// 56160
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// MyProject
	ProjectName *string                                    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectSDK  []*UpdateProjectResponseBodyDataProjectSDK `json:"ProjectSDK,omitempty" xml:"ProjectSDK,omitempty" type:"Repeated"`
	// example:
	//
	// WebApplication
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
}

func (s UpdateProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBodyData) SetCreateTime(v string) *UpdateProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateProjectResponseBodyData) SetProjectApps(v []*UpdateProjectResponseBodyDataProjectApps) *UpdateProjectResponseBodyData {
	s.ProjectApps = v
	return s
}

func (s *UpdateProjectResponseBodyData) SetProjectId(v string) *UpdateProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectResponseBodyData) SetProjectName(v string) *UpdateProjectResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *UpdateProjectResponseBodyData) SetProjectSDK(v []*UpdateProjectResponseBodyDataProjectSDK) *UpdateProjectResponseBodyData {
	s.ProjectSDK = v
	return s
}

func (s *UpdateProjectResponseBodyData) SetProjectType(v string) *UpdateProjectResponseBodyData {
	s.ProjectType = &v
	return s
}

type UpdateProjectResponseBodyDataProjectApps struct {
	ApplicationAccessIds []*UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds `json:"ApplicationAccessIds,omitempty" xml:"ApplicationAccessIds,omitempty" type:"Repeated"`
	// example:
	//
	// 4498
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1889
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateProjectResponseBodyDataProjectApps) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBodyDataProjectApps) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBodyDataProjectApps) SetApplicationAccessIds(v []*UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) *UpdateProjectResponseBodyDataProjectApps {
	s.ApplicationAccessIds = v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectApps) SetId(v string) *UpdateProjectResponseBodyDataProjectApps {
	s.Id = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectApps) SetProjectId(v string) *UpdateProjectResponseBodyDataProjectApps {
	s.ProjectId = &v
	return s
}

type UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds struct {
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// MyAppSecret
	ApplicationAccessSecret *string `json:"applicationAccessSecret,omitempty" xml:"applicationAccessSecret,omitempty"`
}

func (s UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessId(v string) *UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessId = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessSecret(v string) *UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessSecret = &v
	return s
}

type UpdateProjectResponseBodyDataProjectSDK struct {
	// example:
	//
	// 2024-11-01T13:40:53Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// http://demo.com/demo
	DemoUrl    *string `json:"DemoUrl,omitempty" xml:"DemoUrl,omitempty"`
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// example:
	//
	// PHP
	DevelopLanguage *string `json:"DevelopLanguage,omitempty" xml:"DevelopLanguage,omitempty"`
	// example:
	//
	// http://demo.com/doc
	DocUrl  *string `json:"DocUrl,omitempty" xml:"DocUrl,omitempty"`
	SdkName *string `json:"SdkName,omitempty" xml:"SdkName,omitempty"`
	// example:
	//
	// http://demo.com/sdk.zip
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	// example:
	//
	// 4.13.0
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
}

func (s UpdateProjectResponseBodyDataProjectSDK) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBodyDataProjectSDK) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetCreateTime(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.CreateTime = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetDemoUrl(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.DemoUrl = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetDeployMode(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.DeployMode = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetDevelopLanguage(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.DevelopLanguage = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetDocUrl(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.DocUrl = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetSdkName(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.SdkName = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetSdkUrl(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.SdkUrl = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetSdkVersion(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.SdkVersion = &v
	return s
}

type UpdateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetStatusCode(v int32) *UpdateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
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
// 拓展练问答对生成
//
// @param request - AITeacherExpansionPracticeTaskGenerateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AITeacherExpansionPracticeTaskGenerateResponse
func (client *Client) AITeacherExpansionPracticeTaskGenerateWithOptions(request *AITeacherExpansionPracticeTaskGenerateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AITeacherExpansionPracticeTaskGenerateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Grade)) {
		body["grade"] = request.Grade
	}

	if !tea.BoolValue(util.IsUnset(request.KeySentences)) {
		body["keySentences"] = request.KeySentences
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWords)) {
		body["keyWords"] = request.KeyWords
	}

	if !tea.BoolValue(util.IsUnset(request.LearningObject)) {
		body["learningObject"] = request.LearningObject
	}

	if !tea.BoolValue(util.IsUnset(request.TextContent)) {
		body["textContent"] = request.TextContent
	}

	if !tea.BoolValue(util.IsUnset(request.Textbook)) {
		body["textbook"] = request.Textbook
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AITeacherExpansionPracticeTaskGenerate"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/expansionPractice/generateTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AITeacherExpansionPracticeTaskGenerateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拓展练问答对生成
//
// @param request - AITeacherExpansionPracticeTaskGenerateRequest
//
// @return AITeacherExpansionPracticeTaskGenerateResponse
func (client *Client) AITeacherExpansionPracticeTaskGenerate(request *AITeacherExpansionPracticeTaskGenerateRequest) (_result *AITeacherExpansionPracticeTaskGenerateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AITeacherExpansionPracticeTaskGenerateResponse{}
	_body, _err := client.AITeacherExpansionPracticeTaskGenerateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步基础练问答对生成
//
// @param request - AITeacherSyncPracticeTaskGenerateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AITeacherSyncPracticeTaskGenerateResponse
func (client *Client) AITeacherSyncPracticeTaskGenerateWithOptions(request *AITeacherSyncPracticeTaskGenerateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AITeacherSyncPracticeTaskGenerateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Grade)) {
		body["grade"] = request.Grade
	}

	if !tea.BoolValue(util.IsUnset(request.KeySentences)) {
		body["keySentences"] = request.KeySentences
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWords)) {
		body["keyWords"] = request.KeyWords
	}

	if !tea.BoolValue(util.IsUnset(request.LearningObject)) {
		body["learningObject"] = request.LearningObject
	}

	if !tea.BoolValue(util.IsUnset(request.TextContent)) {
		body["textContent"] = request.TextContent
	}

	if !tea.BoolValue(util.IsUnset(request.Textbook)) {
		body["textbook"] = request.Textbook
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AITeacherSyncPracticeTaskGenerate"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/syncPractice/generateTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AITeacherSyncPracticeTaskGenerateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步基础练问答对生成
//
// @param request - AITeacherSyncPracticeTaskGenerateRequest
//
// @return AITeacherSyncPracticeTaskGenerateResponse
func (client *Client) AITeacherSyncPracticeTaskGenerate(request *AITeacherSyncPracticeTaskGenerateRequest) (_result *AITeacherSyncPracticeTaskGenerateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AITeacherSyncPracticeTaskGenerateResponse{}
	_body, _err := client.AITeacherSyncPracticeTaskGenerateWithOptions(request, headers, runtime)
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
		Pathname:    tea.String("/api/v1/aliyunConsole/queryAliyunConsoleServiceList"),
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
// 智能批改/口语评测/统计/调用量
//
// @param request - CountOralEvaluationStatisticsCallsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountOralEvaluationStatisticsCallsResponse
func (client *Client) CountOralEvaluationStatisticsCallsWithOptions(request *CountOralEvaluationStatisticsCallsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CountOralEvaluationStatisticsCallsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CountOralEvaluationStatisticsCalls"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunConsole/countOralEvaluationStatisticsCalls"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CountOralEvaluationStatisticsCallsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能批改/口语评测/统计/调用量
//
// @param request - CountOralEvaluationStatisticsCallsRequest
//
// @return CountOralEvaluationStatisticsCallsResponse
func (client *Client) CountOralEvaluationStatisticsCalls(request *CountOralEvaluationStatisticsCallsRequest) (_result *CountOralEvaluationStatisticsCallsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CountOralEvaluationStatisticsCallsResponse{}
	_body, _err := client.CountOralEvaluationStatisticsCallsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能批改/口语评测/统计/并发数
//
// @param request - CountOralEvaluationStatisticsConcurrentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountOralEvaluationStatisticsConcurrentResponse
func (client *Client) CountOralEvaluationStatisticsConcurrentWithOptions(request *CountOralEvaluationStatisticsConcurrentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CountOralEvaluationStatisticsConcurrentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CountOralEvaluationStatisticsConcurrent"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunConsole/countOralEvaluationStatisticsConcurrent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CountOralEvaluationStatisticsConcurrentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能批改/口语评测/统计/并发数
//
// @param request - CountOralEvaluationStatisticsConcurrentRequest
//
// @return CountOralEvaluationStatisticsConcurrentResponse
func (client *Client) CountOralEvaluationStatisticsConcurrent(request *CountOralEvaluationStatisticsConcurrentRequest) (_result *CountOralEvaluationStatisticsConcurrentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CountOralEvaluationStatisticsConcurrentResponse{}
	_body, _err := client.CountOralEvaluationStatisticsConcurrentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 智能批改/口语评测/统计/调用错误
//
// @param request - CountOralEvaluationStatisticsErrorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountOralEvaluationStatisticsErrorResponse
func (client *Client) CountOralEvaluationStatisticsErrorWithOptions(request *CountOralEvaluationStatisticsErrorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CountOralEvaluationStatisticsErrorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CountOralEvaluationStatisticsError"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunConsole/countOralEvaluationStatisticsError"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CountOralEvaluationStatisticsErrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 智能批改/口语评测/统计/调用错误
//
// @param request - CountOralEvaluationStatisticsErrorRequest
//
// @return CountOralEvaluationStatisticsErrorResponse
func (client *Client) CountOralEvaluationStatisticsError(request *CountOralEvaluationStatisticsErrorRequest) (_result *CountOralEvaluationStatisticsErrorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CountOralEvaluationStatisticsErrorResponse{}
	_body, _err := client.CountOralEvaluationStatisticsErrorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 阿里云控制台/授权凭证创建
//
// @param request - CreateAccessWarrantRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessWarrantResponse
func (client *Client) CreateAccessWarrantWithOptions(request *CreateAccessWarrantRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAccessWarrantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestSign)) {
		body["requestSign"] = request.RequestSign
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		body["userClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WarrantAvailable)) {
		body["warrantAvailable"] = request.WarrantAvailable
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessWarrant"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunConsole/createAccessWarrant"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessWarrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 阿里云控制台/授权凭证创建
//
// @param request - CreateAccessWarrantRequest
//
// @return CreateAccessWarrantResponse
func (client *Client) CreateAccessWarrant(request *CreateAccessWarrantRequest) (_result *CreateAccessWarrantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAccessWarrantResponse{}
	_body, _err := client.CreateAccessWarrantWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 阿里云控制台/创建项目
//
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["projectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectType)) {
		body["projectType"] = request.ProjectType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunConsole/createProject"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 阿里云控制台/创建项目
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 中文作文辅导
//
// @param request - ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse
func (client *Client) ExecuteAITeacherChineseCompositionTutoringWorkflowRunWithOptions(request *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EssayOutline)) {
		body["essayOutline"] = request.EssayOutline
	}

	if !tea.BoolValue(util.IsUnset(request.EssayRequirements)) {
		body["essayRequirements"] = request.EssayRequirements
	}

	if !tea.BoolValue(util.IsUnset(request.EssayTopic)) {
		body["essayTopic"] = request.EssayTopic
	}

	if !tea.BoolValue(util.IsUnset(request.EssayType)) {
		body["essayType"] = request.EssayType
	}

	if !tea.BoolValue(util.IsUnset(request.EssayWordCount)) {
		body["essayWordCount"] = request.EssayWordCount
	}

	if !tea.BoolValue(util.IsUnset(request.Grade)) {
		body["grade"] = request.Grade
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseMode)) {
		body["responseMode"] = request.ResponseMode
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherChineseCompositionTutoringWorkflowRun"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/api/v1/intelligentAgent/chineseCompositionTutoring/workflowRun"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 中文作文辅导
//
// @param request - ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest
//
// @return ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse
func (client *Client) ExecuteAITeacherChineseCompositionTutoringWorkflowRun(request *ExecuteAITeacherChineseCompositionTutoringWorkflowRunRequest) (_result *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponse{}
	_body, _err := client.ExecuteAITeacherChineseCompositionTutoringWorkflowRunWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 英语作文辅导
//
// @param request - ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse
func (client *Client) ExecuteAITeacherEnglishCompositionTutoringWorkflowRunWithOptions(request *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EssayOutline)) {
		body["essayOutline"] = request.EssayOutline
	}

	if !tea.BoolValue(util.IsUnset(request.EssayRequirements)) {
		body["essayRequirements"] = request.EssayRequirements
	}

	if !tea.BoolValue(util.IsUnset(request.EssayTopic)) {
		body["essayTopic"] = request.EssayTopic
	}

	if !tea.BoolValue(util.IsUnset(request.EssayType)) {
		body["essayType"] = request.EssayType
	}

	if !tea.BoolValue(util.IsUnset(request.EssayWordCount)) {
		body["essayWordCount"] = request.EssayWordCount
	}

	if !tea.BoolValue(util.IsUnset(request.Grade)) {
		body["grade"] = request.Grade
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseMode)) {
		body["responseMode"] = request.ResponseMode
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherEnglishCompositionTutoringWorkflowRun"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/api/v1/intelligentAgent/englishCompositionTutoring/workflowRun"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 英语作文辅导
//
// @param request - ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest
//
// @return ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse
func (client *Client) ExecuteAITeacherEnglishCompositionTutoringWorkflowRun(request *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunRequest) (_result *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponse{}
	_body, _err := client.ExecuteAITeacherEnglishCompositionTutoringWorkflowRunWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 英文释义
//
// @param request - ExecuteAITeacherEnglishParaphraseChatMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherEnglishParaphraseChatMessageResponse
func (client *Client) ExecuteAITeacherEnglishParaphraseChatMessageWithOptions(request *ExecuteAITeacherEnglishParaphraseChatMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherEnglishParaphraseChatMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatId)) {
		body["chatId"] = request.ChatId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Grade)) {
		body["grade"] = request.Grade
	}

	if !tea.BoolValue(util.IsUnset(request.QuestionId)) {
		body["questionId"] = request.QuestionId
	}

	if !tea.BoolValue(util.IsUnset(request.QuestionInfo)) {
		body["questionInfo"] = request.QuestionInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseMode)) {
		body["responseMode"] = request.ResponseMode
	}

	if !tea.BoolValue(util.IsUnset(request.UserAnswer)) {
		body["userAnswer"] = request.UserAnswer
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherEnglishParaphraseChatMessage"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/api/v1/intelligentAgent/englishParaphrase/chatMessage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherEnglishParaphraseChatMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 英文释义
//
// @param request - ExecuteAITeacherEnglishParaphraseChatMessageRequest
//
// @return ExecuteAITeacherEnglishParaphraseChatMessageResponse
func (client *Client) ExecuteAITeacherEnglishParaphraseChatMessage(request *ExecuteAITeacherEnglishParaphraseChatMessageRequest) (_result *ExecuteAITeacherEnglishParaphraseChatMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherEnglishParaphraseChatMessageResponse{}
	_body, _err := client.ExecuteAITeacherEnglishParaphraseChatMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 进行拓展练对话
//
// @param request - ExecuteAITeacherExpansionDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherExpansionDialogueResponse
func (client *Client) ExecuteAITeacherExpansionDialogueWithOptions(request *ExecuteAITeacherExpansionDialogueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherExpansionDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Background)) {
		body["background"] = request.Background
	}

	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		body["languageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.RoleInfo)) {
		body["roleInfo"] = request.RoleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.StartSentence)) {
		body["startSentence"] = request.StartSentence
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherExpansionDialogue"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/expansionPractice/executeExpansionTraining"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherExpansionDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 进行拓展练对话
//
// @param request - ExecuteAITeacherExpansionDialogueRequest
//
// @return ExecuteAITeacherExpansionDialogueResponse
func (client *Client) ExecuteAITeacherExpansionDialogue(request *ExecuteAITeacherExpansionDialogueRequest) (_result *ExecuteAITeacherExpansionDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherExpansionDialogueResponse{}
	_body, _err := client.ExecuteAITeacherExpansionDialogueWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拓展练根据上下文进行润色
//
// @param request - ExecuteAITeacherExpansionDialogueRefineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherExpansionDialogueRefineResponse
func (client *Client) ExecuteAITeacherExpansionDialogueRefineWithOptions(request *ExecuteAITeacherExpansionDialogueRefineRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherExpansionDialogueRefineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Background)) {
		body["background"] = request.Background
	}

	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		body["languageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.RoleInfo)) {
		body["roleInfo"] = request.RoleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.StartSentence)) {
		body["startSentence"] = request.StartSentence
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherExpansionDialogueRefine"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/expansionPractice/refineByContext"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherExpansionDialogueRefineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拓展练根据上下文进行润色
//
// @param request - ExecuteAITeacherExpansionDialogueRefineRequest
//
// @return ExecuteAITeacherExpansionDialogueRefineResponse
func (client *Client) ExecuteAITeacherExpansionDialogueRefine(request *ExecuteAITeacherExpansionDialogueRefineRequest) (_result *ExecuteAITeacherExpansionDialogueRefineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherExpansionDialogueRefineResponse{}
	_body, _err := client.ExecuteAITeacherExpansionDialogueRefineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拓展练语境翻译
//
// @param request - ExecuteAITeacherExpansionDialogueTranslateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherExpansionDialogueTranslateResponse
func (client *Client) ExecuteAITeacherExpansionDialogueTranslateWithOptions(request *ExecuteAITeacherExpansionDialogueTranslateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherExpansionDialogueTranslateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Background)) {
		body["background"] = request.Background
	}

	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.RoleInfo)) {
		body["roleInfo"] = request.RoleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.StartSentence)) {
		body["startSentence"] = request.StartSentence
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherExpansionDialogueTranslate"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/expansionPractice/translate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherExpansionDialogueTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拓展练语境翻译
//
// @param request - ExecuteAITeacherExpansionDialogueTranslateRequest
//
// @return ExecuteAITeacherExpansionDialogueTranslateResponse
func (client *Client) ExecuteAITeacherExpansionDialogueTranslate(request *ExecuteAITeacherExpansionDialogueTranslateRequest) (_result *ExecuteAITeacherExpansionDialogueTranslateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherExpansionDialogueTranslateResponse{}
	_body, _err := client.ExecuteAITeacherExpansionDialogueTranslateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 语法检测
//
// @param request - ExecuteAITeacherGrammarCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherGrammarCheckResponse
func (client *Client) ExecuteAITeacherGrammarCheckWithOptions(request *ExecuteAITeacherGrammarCheckRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherGrammarCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherGrammarCheck"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/common/grammarChecking"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherGrammarCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语法检测
//
// @param request - ExecuteAITeacherGrammarCheckRequest
//
// @return ExecuteAITeacherGrammarCheckResponse
func (client *Client) ExecuteAITeacherGrammarCheck(request *ExecuteAITeacherGrammarCheckRequest) (_result *ExecuteAITeacherGrammarCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherGrammarCheckResponse{}
	_body, _err := client.ExecuteAITeacherGrammarCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 进行同步练对话
//
// @param request - ExecuteAITeacherSyncDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherSyncDialogueResponse
func (client *Client) ExecuteAITeacherSyncDialogueWithOptions(request *ExecuteAITeacherSyncDialogueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherSyncDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		body["languageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherSyncDialogue"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/syncPractice/executeSyncTraining"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherSyncDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 进行同步练对话
//
// @param request - ExecuteAITeacherSyncDialogueRequest
//
// @return ExecuteAITeacherSyncDialogueResponse
func (client *Client) ExecuteAITeacherSyncDialogue(request *ExecuteAITeacherSyncDialogueRequest) (_result *ExecuteAITeacherSyncDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherSyncDialogueResponse{}
	_body, _err := client.ExecuteAITeacherSyncDialogueWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步练语境翻译
//
// @param request - ExecuteAITeacherSyncDialogueTranslateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAITeacherSyncDialogueTranslateResponse
func (client *Client) ExecuteAITeacherSyncDialogueTranslateWithOptions(request *ExecuteAITeacherSyncDialogueTranslateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteAITeacherSyncDialogueTranslateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteAITeacherSyncDialogueTranslate"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/syncPractice/translate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteAITeacherSyncDialogueTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步练语境翻译
//
// @param request - ExecuteAITeacherSyncDialogueTranslateRequest
//
// @return ExecuteAITeacherSyncDialogueTranslateResponse
func (client *Client) ExecuteAITeacherSyncDialogueTranslate(request *ExecuteAITeacherSyncDialogueTranslateRequest) (_result *ExecuteAITeacherSyncDialogueTranslateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteAITeacherSyncDialogueTranslateResponse{}
	_body, _err := client.ExecuteAITeacherSyncDialogueTranslateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 十万个为什么对话接入
//
// @param request - ExecuteHundredThousandWhysDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteHundredThousandWhysDialogueResponse
func (client *Client) ExecuteHundredThousandWhysDialogueWithOptions(request *ExecuteHundredThousandWhysDialogueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteHundredThousandWhysDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgeGroup)) {
		body["ageGroup"] = request.AgeGroup
	}

	if !tea.BoolValue(util.IsUnset(request.ChatId)) {
		body["chatId"] = request.ChatId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.MacAddress)) {
		body["macAddress"] = request.MacAddress
	}

	if !tea.BoolValue(util.IsUnset(request.Messages)) {
		body["messages"] = request.Messages
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteHundredThousandWhysDialogue"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/pop/api/v1/intelligentAgent/tenWWhys/executeDialogue"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteHundredThousandWhysDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 十万个为什么对话接入
//
// @param request - ExecuteHundredThousandWhysDialogueRequest
//
// @return ExecuteHundredThousandWhysDialogueResponse
func (client *Client) ExecuteHundredThousandWhysDialogue(request *ExecuteHundredThousandWhysDialogueRequest) (_result *ExecuteHundredThousandWhysDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteHundredThousandWhysDialogueResponse{}
	_body, _err := client.ExecuteHundredThousandWhysDialogueWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 进行AI对话
//
// @param request - ExecuteTextbookAssistantDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantDialogueResponse
func (client *Client) ExecuteTextbookAssistantDialogueWithOptions(request *ExecuteTextbookAssistantDialogueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteTextbookAssistantDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.ChatId)) {
		body["chatId"] = request.ChatId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	if !tea.BoolValue(util.IsUnset(request.UserMessage)) {
		body["userMessage"] = request.UserMessage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteTextbookAssistantDialogue"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/dialogue/ExecuteDialogue"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteTextbookAssistantDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 进行AI对话
//
// @param request - ExecuteTextbookAssistantDialogueRequest
//
// @return ExecuteTextbookAssistantDialogueResponse
func (client *Client) ExecuteTextbookAssistantDialogue(request *ExecuteTextbookAssistantDialogueRequest) (_result *ExecuteTextbookAssistantDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTextbookAssistantDialogueResponse{}
	_body, _err := client.ExecuteTextbookAssistantDialogueWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调整难度
//
// @param request - ExecuteTextbookAssistantDifficultyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantDifficultyResponse
func (client *Client) ExecuteTextbookAssistantDifficultyWithOptions(request *ExecuteTextbookAssistantDifficultyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteTextbookAssistantDifficultyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.Assistant)) {
		body["assistant"] = request.Assistant
	}

	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.ChatId)) {
		body["chatId"] = request.ChatId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteTextbookAssistantDifficulty"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/dialogue/ExecuteDifficulty"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteTextbookAssistantDifficultyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调整难度
//
// @param request - ExecuteTextbookAssistantDifficultyRequest
//
// @return ExecuteTextbookAssistantDifficultyResponse
func (client *Client) ExecuteTextbookAssistantDifficulty(request *ExecuteTextbookAssistantDifficultyRequest) (_result *ExecuteTextbookAssistantDifficultyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTextbookAssistantDifficultyResponse{}
	_body, _err := client.ExecuteTextbookAssistantDifficultyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 语法检测
//
// @param request - ExecuteTextbookAssistantGrammarCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantGrammarCheckResponse
func (client *Client) ExecuteTextbookAssistantGrammarCheckWithOptions(request *ExecuteTextbookAssistantGrammarCheckRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteTextbookAssistantGrammarCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.ChatId)) {
		body["chatId"] = request.ChatId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		body["user"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteTextbookAssistantGrammarCheck"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/dialogue/ExecuteGrammarCheck"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteTextbookAssistantGrammarCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语法检测
//
// @param request - ExecuteTextbookAssistantGrammarCheckRequest
//
// @return ExecuteTextbookAssistantGrammarCheckResponse
func (client *Client) ExecuteTextbookAssistantGrammarCheck(request *ExecuteTextbookAssistantGrammarCheckRequest) (_result *ExecuteTextbookAssistantGrammarCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTextbookAssistantGrammarCheckResponse{}
	_body, _err := client.ExecuteTextbookAssistantGrammarCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 句子润色
//
// @param request - ExecuteTextbookAssistantRefineByContextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantRefineByContextResponse
func (client *Client) ExecuteTextbookAssistantRefineByContextWithOptions(request *ExecuteTextbookAssistantRefineByContextRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteTextbookAssistantRefineByContextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.ChatId)) {
		body["chatId"] = request.ChatId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		body["user"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteTextbookAssistantRefineByContext"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/dialogue/RefineByContext"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteTextbookAssistantRefineByContextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 句子润色
//
// @param request - ExecuteTextbookAssistantRefineByContextRequest
//
// @return ExecuteTextbookAssistantRefineByContextResponse
func (client *Client) ExecuteTextbookAssistantRefineByContext(request *ExecuteTextbookAssistantRefineByContextRequest) (_result *ExecuteTextbookAssistantRefineByContextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTextbookAssistantRefineByContextResponse{}
	_body, _err := client.ExecuteTextbookAssistantRefineByContextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对话重试
//
// @param request - ExecuteTextbookAssistantRetryConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantRetryConversationResponse
func (client *Client) ExecuteTextbookAssistantRetryConversationWithOptions(request *ExecuteTextbookAssistantRetryConversationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteTextbookAssistantRetryConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Assistant)) {
		body["assistant"] = request.Assistant
	}

	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.ChatId)) {
		body["chatId"] = request.ChatId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteTextbookAssistantRetryConversation"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/dialogue/RetryConversation"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteTextbookAssistantRetryConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对话重试
//
// @param request - ExecuteTextbookAssistantRetryConversationRequest
//
// @return ExecuteTextbookAssistantRetryConversationResponse
func (client *Client) ExecuteTextbookAssistantRetryConversation(request *ExecuteTextbookAssistantRetryConversationRequest) (_result *ExecuteTextbookAssistantRetryConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTextbookAssistantRetryConversationResponse{}
	_body, _err := client.ExecuteTextbookAssistantRetryConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 进行对话-流式输出
//
// @param request - ExecuteTextbookAssistantSseDialogueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantSseDialogueResponse
func (client *Client) ExecuteTextbookAssistantSseDialogueWithOptions(request *ExecuteTextbookAssistantSseDialogueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteTextbookAssistantSseDialogueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.ChatId)) {
		body["chatId"] = request.ChatId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	if !tea.BoolValue(util.IsUnset(request.UserMessage)) {
		body["userMessage"] = request.UserMessage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteTextbookAssistantSseDialogue"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/dialogue/ExecuteSseDialogue"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteTextbookAssistantSseDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 进行对话-流式输出
//
// @param request - ExecuteTextbookAssistantSseDialogueRequest
//
// @return ExecuteTextbookAssistantSseDialogueResponse
func (client *Client) ExecuteTextbookAssistantSseDialogue(request *ExecuteTextbookAssistantSseDialogueRequest) (_result *ExecuteTextbookAssistantSseDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTextbookAssistantSseDialogueResponse{}
	_body, _err := client.ExecuteTextbookAssistantSseDialogueWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启自由对话
//
// @param request - ExecuteTextbookAssistantStartConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantStartConversationResponse
func (client *Client) ExecuteTextbookAssistantStartConversationWithOptions(request *ExecuteTextbookAssistantStartConversationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteTextbookAssistantStartConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArticleId)) {
		body["articleId"] = request.ArticleId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteTextbookAssistantStartConversation"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/dialogue/StartConversation"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteTextbookAssistantStartConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启自由对话
//
// @param request - ExecuteTextbookAssistantStartConversationRequest
//
// @return ExecuteTextbookAssistantStartConversationResponse
func (client *Client) ExecuteTextbookAssistantStartConversation(request *ExecuteTextbookAssistantStartConversationRequest) (_result *ExecuteTextbookAssistantStartConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTextbookAssistantStartConversationResponse{}
	_body, _err := client.ExecuteTextbookAssistantStartConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取鉴权参数
//
// @param request - ExecuteTextbookAssistantSuggestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantSuggestionResponse
func (client *Client) ExecuteTextbookAssistantSuggestionWithOptions(request *ExecuteTextbookAssistantSuggestionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteTextbookAssistantSuggestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Assistant)) {
		body["assistant"] = request.Assistant
	}

	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.ChatId)) {
		body["chatId"] = request.ChatId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteTextbookAssistantSuggestion"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/dialogue/Suggestion"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteTextbookAssistantSuggestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取鉴权参数
//
// @param request - ExecuteTextbookAssistantSuggestionRequest
//
// @return ExecuteTextbookAssistantSuggestionResponse
func (client *Client) ExecuteTextbookAssistantSuggestion(request *ExecuteTextbookAssistantSuggestionRequest) (_result *ExecuteTextbookAssistantSuggestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTextbookAssistantSuggestionResponse{}
	_body, _err := client.ExecuteTextbookAssistantSuggestionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 翻译消息内容
//
// @param request - ExecuteTextbookAssistantTranslateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTextbookAssistantTranslateResponse
func (client *Client) ExecuteTextbookAssistantTranslateWithOptions(request *ExecuteTextbookAssistantTranslateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteTextbookAssistantTranslateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Assistant)) {
		body["assistant"] = request.Assistant
	}

	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.ChatId)) {
		body["chatId"] = request.ChatId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteTextbookAssistantTranslate"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/dialogue/ExecuteTranslate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteTextbookAssistantTranslateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 翻译消息内容
//
// @param request - ExecuteTextbookAssistantTranslateRequest
//
// @return ExecuteTextbookAssistantTranslateResponse
func (client *Client) ExecuteTextbookAssistantTranslate(request *ExecuteTextbookAssistantTranslateRequest) (_result *ExecuteTextbookAssistantTranslateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteTextbookAssistantTranslateResponse{}
	_body, _err := client.ExecuteTextbookAssistantTranslateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拓展练小助手
//
// @param request - GetAITeacherExpansionDialogueSuggestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAITeacherExpansionDialogueSuggestionResponse
func (client *Client) GetAITeacherExpansionDialogueSuggestionWithOptions(request *GetAITeacherExpansionDialogueSuggestionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAITeacherExpansionDialogueSuggestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Background)) {
		body["background"] = request.Background
	}

	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		body["languageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.RoleInfo)) {
		body["roleInfo"] = request.RoleInfo
	}

	if !tea.BoolValue(util.IsUnset(request.StartSentence)) {
		body["startSentence"] = request.StartSentence
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAITeacherExpansionDialogueSuggestion"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/expansionPractice/suggestion"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAITeacherExpansionDialogueSuggestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拓展练小助手
//
// @param request - GetAITeacherExpansionDialogueSuggestionRequest
//
// @return GetAITeacherExpansionDialogueSuggestionResponse
func (client *Client) GetAITeacherExpansionDialogueSuggestion(request *GetAITeacherExpansionDialogueSuggestionRequest) (_result *GetAITeacherExpansionDialogueSuggestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAITeacherExpansionDialogueSuggestionResponse{}
	_body, _err := client.GetAITeacherExpansionDialogueSuggestionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步练小助手
//
// @param request - GetAITeacherSyncDialogueSuggestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAITeacherSyncDialogueSuggestionResponse
func (client *Client) GetAITeacherSyncDialogueSuggestionWithOptions(request *GetAITeacherSyncDialogueSuggestionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAITeacherSyncDialogueSuggestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DialogueTasks)) {
		body["dialogueTasks"] = request.DialogueTasks
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageCode)) {
		body["languageCode"] = request.LanguageCode
	}

	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAITeacherSyncDialogueSuggestion"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aiteacher/syncPractice/suggestion"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAITeacherSyncDialogueSuggestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步练小助手
//
// @param request - GetAITeacherSyncDialogueSuggestionRequest
//
// @return GetAITeacherSyncDialogueSuggestionResponse
func (client *Client) GetAITeacherSyncDialogueSuggestion(request *GetAITeacherSyncDialogueSuggestionRequest) (_result *GetAITeacherSyncDialogueSuggestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAITeacherSyncDialogueSuggestionResponse{}
	_body, _err := client.GetAITeacherSyncDialogueSuggestionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取请求鉴权参数
//
// @param request - GetTextbookAssistantTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextbookAssistantTokenResponse
func (client *Client) GetTextbookAssistantTokenWithOptions(request *GetTextbookAssistantTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTextbookAssistantTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["model"] = request.Model
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTextbookAssistantToken"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/teachingResource/GetToken"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTextbookAssistantTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取请求鉴权参数
//
// @param request - GetTextbookAssistantTokenRequest
//
// @return GetTextbookAssistantTokenResponse
func (client *Client) GetTextbookAssistantToken(request *GetTextbookAssistantTokenRequest) (_result *GetTextbookAssistantTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextbookAssistantTokenResponse{}
	_body, _err := client.GetTextbookAssistantTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取文章详情
//
// @param request - ListTextbookAssistantArticleDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantArticleDetailsResponse
func (client *Client) ListTextbookAssistantArticleDetailsWithOptions(request *ListTextbookAssistantArticleDetailsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTextbookAssistantArticleDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArticleIdList)) {
		body["articleIdList"] = request.ArticleIdList
	}

	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTextbookAssistantArticleDetails"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/teachingResource/ListArticleDetails"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTextbookAssistantArticleDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取文章详情
//
// @param request - ListTextbookAssistantArticleDetailsRequest
//
// @return ListTextbookAssistantArticleDetailsResponse
func (client *Client) ListTextbookAssistantArticleDetails(request *ListTextbookAssistantArticleDetailsRequest) (_result *ListTextbookAssistantArticleDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTextbookAssistantArticleDetailsResponse{}
	_body, _err := client.ListTextbookAssistantArticleDetailsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文章列表
//
// @param request - ListTextbookAssistantArticlesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantArticlesResponse
func (client *Client) ListTextbookAssistantArticlesWithOptions(request *ListTextbookAssistantArticlesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTextbookAssistantArticlesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		body["directoryId"] = request.DirectoryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTextbookAssistantArticles"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/teachingResource/ListArticles"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTextbookAssistantArticlesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文章列表
//
// @param request - ListTextbookAssistantArticlesRequest
//
// @return ListTextbookAssistantArticlesResponse
func (client *Client) ListTextbookAssistantArticles(request *ListTextbookAssistantArticlesRequest) (_result *ListTextbookAssistantArticlesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTextbookAssistantArticlesResponse{}
	_body, _err := client.ListTextbookAssistantArticlesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取书本下的目录信息
//
// @param request - ListTextbookAssistantBookDirectoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantBookDirectoriesResponse
func (client *Client) ListTextbookAssistantBookDirectoriesWithOptions(request *ListTextbookAssistantBookDirectoriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTextbookAssistantBookDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.BookId)) {
		body["bookId"] = request.BookId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTextbookAssistantBookDirectories"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/teachingResource/ListBookDirectories"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTextbookAssistantBookDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取书本下的目录信息
//
// @param request - ListTextbookAssistantBookDirectoriesRequest
//
// @return ListTextbookAssistantBookDirectoriesResponse
func (client *Client) ListTextbookAssistantBookDirectories(request *ListTextbookAssistantBookDirectoriesRequest) (_result *ListTextbookAssistantBookDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTextbookAssistantBookDirectoriesResponse{}
	_body, _err := client.ListTextbookAssistantBookDirectoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取包含年级下的书本列表
//
// @param request - ListTextbookAssistantBooksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantBooksResponse
func (client *Client) ListTextbookAssistantBooksWithOptions(request *ListTextbookAssistantBooksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTextbookAssistantBooksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.BookId)) {
		body["bookId"] = request.BookId
	}

	if !tea.BoolValue(util.IsUnset(request.Grade)) {
		body["grade"] = request.Grade
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		body["volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTextbookAssistantBooks"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/teachingResource/ListBooks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTextbookAssistantBooksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取包含年级下的书本列表
//
// @param request - ListTextbookAssistantBooksRequest
//
// @return ListTextbookAssistantBooksResponse
func (client *Client) ListTextbookAssistantBooks(request *ListTextbookAssistantBooksRequest) (_result *ListTextbookAssistantBooksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTextbookAssistantBooksResponse{}
	_body, _err := client.ListTextbookAssistantBooksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取有资源的年级信息
//
// @param request - ListTextbookAssistantGradeVolumesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantGradeVolumesResponse
func (client *Client) ListTextbookAssistantGradeVolumesWithOptions(request *ListTextbookAssistantGradeVolumesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTextbookAssistantGradeVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		body["scenario"] = request.Scenario
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTextbookAssistantGradeVolumes"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/teachingResource/ListGradeVolumes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTextbookAssistantGradeVolumesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取有资源的年级信息
//
// @param request - ListTextbookAssistantGradeVolumesRequest
//
// @return ListTextbookAssistantGradeVolumesResponse
func (client *Client) ListTextbookAssistantGradeVolumes(request *ListTextbookAssistantGradeVolumesRequest) (_result *ListTextbookAssistantGradeVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTextbookAssistantGradeVolumesResponse{}
	_body, _err := client.ListTextbookAssistantGradeVolumesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文章内容详情
//
// @param request - ListTextbookAssistantSceneDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextbookAssistantSceneDetailsResponse
func (client *Client) ListTextbookAssistantSceneDetailsWithOptions(request *ListTextbookAssistantSceneDetailsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTextbookAssistantSceneDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthToken)) {
		body["authToken"] = request.AuthToken
	}

	if !tea.BoolValue(util.IsUnset(request.SceneIdList)) {
		body["sceneIdList"] = request.SceneIdList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTextbookAssistantSceneDetails"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/textbookAssistant/teachingResource/ListSceneDetails"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTextbookAssistantSceneDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文章内容详情
//
// @param request - ListTextbookAssistantSceneDetailsRequest
//
// @return ListTextbookAssistantSceneDetailsResponse
func (client *Client) ListTextbookAssistantSceneDetails(request *ListTextbookAssistantSceneDetailsRequest) (_result *ListTextbookAssistantSceneDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTextbookAssistantSceneDetailsResponse{}
	_body, _err := client.ListTextbookAssistantSceneDetailsWithOptions(request, headers, runtime)
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

// Summary:
//
// 阿里云控制台/获取应用访问识别码(appkey)信息
//
// @param request - QueryApplicationAccessIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryApplicationAccessIdResponse
func (client *Client) QueryApplicationAccessIdWithOptions(request *QueryApplicationAccessIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryApplicationAccessIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationAccessId)) {
		query["applicationAccessId"] = request.ApplicationAccessId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryApplicationAccessId"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunConsole/queryApplicationAccessId"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryApplicationAccessIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 阿里云控制台/获取应用访问识别码(appkey)信息
//
// @param request - QueryApplicationAccessIdRequest
//
// @return QueryApplicationAccessIdResponse
func (client *Client) QueryApplicationAccessId(request *QueryApplicationAccessIdRequest) (_result *QueryApplicationAccessIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryApplicationAccessIdResponse{}
	_body, _err := client.QueryApplicationAccessIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 阿里云控制台/获取项目列表
//
// @param request - QueryProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProjectResponse
func (client *Client) QueryProjectWithOptions(request *QueryProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["projectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryProject"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunConsole/queryProject"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 阿里云控制台/获取项目列表
//
// @param request - QueryProjectRequest
//
// @return QueryProjectResponse
func (client *Client) QueryProject(request *QueryProjectRequest) (_result *QueryProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryProjectResponse{}
	_body, _err := client.QueryProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 阿里云控制台/获取项目列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProjectListResponse
func (client *Client) QueryProjectListWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryProjectListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("QueryProjectList"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunConsole/queryProjectList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProjectListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 阿里云控制台/获取项目列表
//
// @return QueryProjectListResponse
func (client *Client) QueryProjectList() (_result *QueryProjectListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryProjectListResponse{}
	_body, _err := client.QueryProjectListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 阿里云控制台/已经购买过的服务项目
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPurchasedServiceResponse
func (client *Client) QueryPurchasedServiceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryPurchasedServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPurchasedService"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunConsole/queryPurchasedService"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPurchasedServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 阿里云控制台/已经购买过的服务项目
//
// @return QueryPurchasedServiceResponse
func (client *Client) QueryPurchasedService() (_result *QueryPurchasedServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPurchasedServiceResponse{}
	_body, _err := client.QueryPurchasedServiceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 阿里云控制台/更新项目信息
//
// @param request - UpdateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithOptions(request *UpdateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["projectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("20240611"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/aliyunConsole/updateProject"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 阿里云控制台/更新项目信息
//
// @param request - UpdateProjectRequest
//
// @return UpdateProjectResponse
func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
