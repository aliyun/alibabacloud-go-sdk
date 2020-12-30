// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAsyncPredictRequest struct {
	ModelId      *int32  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	DetailTag    *string `json:"DetailTag,omitempty" xml:"DetailTag,omitempty"`
	TopK         *int32  `json:"TopK,omitempty" xml:"TopK,omitempty"`
	FileType     *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FileContent  *string `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	FetchContent *string `json:"FetchContent,omitempty" xml:"FetchContent,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s CreateAsyncPredictRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncPredictRequest) GoString() string {
	return s.String()
}

func (s *CreateAsyncPredictRequest) SetModelId(v int32) *CreateAsyncPredictRequest {
	s.ModelId = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetContent(v string) *CreateAsyncPredictRequest {
	s.Content = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetModelVersion(v string) *CreateAsyncPredictRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetDetailTag(v string) *CreateAsyncPredictRequest {
	s.DetailTag = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetTopK(v int32) *CreateAsyncPredictRequest {
	s.TopK = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetFileType(v string) *CreateAsyncPredictRequest {
	s.FileType = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetFileUrl(v string) *CreateAsyncPredictRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetFileContent(v string) *CreateAsyncPredictRequest {
	s.FileContent = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetFetchContent(v string) *CreateAsyncPredictRequest {
	s.FetchContent = &v
	return s
}

func (s *CreateAsyncPredictRequest) SetProduct(v string) *CreateAsyncPredictRequest {
	s.Product = &v
	return s
}

type CreateAsyncPredictResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AsyncPredictId *int64  `json:"AsyncPredictId,omitempty" xml:"AsyncPredictId,omitempty"`
}

func (s CreateAsyncPredictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncPredictResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAsyncPredictResponseBody) SetRequestId(v string) *CreateAsyncPredictResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAsyncPredictResponseBody) SetAsyncPredictId(v int64) *CreateAsyncPredictResponseBody {
	s.AsyncPredictId = &v
	return s
}

type CreateAsyncPredictResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAsyncPredictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAsyncPredictResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncPredictResponse) GoString() string {
	return s.String()
}

func (s *CreateAsyncPredictResponse) SetHeaders(v map[string]*string) *CreateAsyncPredictResponse {
	s.Headers = v
	return s
}

func (s *CreateAsyncPredictResponse) SetBody(v *CreateAsyncPredictResponseBody) *CreateAsyncPredictResponse {
	s.Body = v
	return s
}

type CreateDatasetRequest struct {
	ProjectId   *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Product     *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) SetProjectId(v int64) *CreateDatasetRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDatasetRequest) SetDatasetName(v string) *CreateDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateDatasetRequest) SetProduct(v string) *CreateDatasetRequest {
	s.Product = &v
	return s
}

type CreateDatasetResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
	DatasetId map[string]interface{} `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) SetMessage(v string) *CreateDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetResponseBody) SetCode(v int32) *CreateDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDatasetResponseBody) SetSuccess(v bool) *CreateDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDatasetResponseBody) SetDatasetId(v map[string]interface{}) *CreateDatasetResponseBody {
	s.DatasetId = v
	return s
}

type CreateDatasetResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponse) SetHeaders(v map[string]*string) *CreateDatasetResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetResponse) SetBody(v *CreateDatasetResponseBody) *CreateDatasetResponse {
	s.Body = v
	return s
}

type CreateDatasetRecordRequest struct {
	DatasetId     *int64  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetRecord *string `json:"DatasetRecord,omitempty" xml:"DatasetRecord,omitempty"`
	ProjectId     *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateDatasetRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRecordRequest) SetDatasetId(v int64) *CreateDatasetRecordRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateDatasetRecordRequest) SetDatasetRecord(v string) *CreateDatasetRecordRequest {
	s.DatasetRecord = &v
	return s
}

func (s *CreateDatasetRecordRequest) SetProjectId(v int64) *CreateDatasetRecordRequest {
	s.ProjectId = &v
	return s
}

type CreateDatasetRecordResponseBody struct {
	Message         *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DatasetRecordId map[string]interface{} `json:"DatasetRecordId,omitempty" xml:"DatasetRecordId,omitempty"`
	Code            *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success         *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDatasetRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetRecordResponseBody) SetMessage(v string) *CreateDatasetRecordResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDatasetRecordResponseBody) SetRequestId(v string) *CreateDatasetRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetRecordResponseBody) SetDatasetRecordId(v map[string]interface{}) *CreateDatasetRecordResponseBody {
	s.DatasetRecordId = v
	return s
}

func (s *CreateDatasetRecordResponseBody) SetCode(v int32) *CreateDatasetRecordResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDatasetRecordResponseBody) SetSuccess(v bool) *CreateDatasetRecordResponseBody {
	s.Success = &v
	return s
}

type CreateDatasetRecordResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDatasetRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDatasetRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetRecordResponse) SetHeaders(v map[string]*string) *CreateDatasetRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetRecordResponse) SetBody(v *CreateDatasetRecordResponseBody) *CreateDatasetRecordResponse {
	s.Body = v
	return s
}

type CreateModelRequest struct {
	ModelId            *int64                 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelType          *string                `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ProjectId          *int64                 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ModelName          *string                `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	DatasetIdList      map[string]interface{} `json:"DatasetIdList,omitempty" xml:"DatasetIdList,omitempty"`
	TestDatasetIdList  map[string]interface{} `json:"TestDatasetIdList,omitempty" xml:"TestDatasetIdList,omitempty"`
	IsIncrementalTrain *string                `json:"IsIncrementalTrain,omitempty" xml:"IsIncrementalTrain,omitempty"`
	Product            *string                `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s CreateModelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModelRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequest) SetModelId(v int64) *CreateModelRequest {
	s.ModelId = &v
	return s
}

func (s *CreateModelRequest) SetModelType(v string) *CreateModelRequest {
	s.ModelType = &v
	return s
}

func (s *CreateModelRequest) SetProjectId(v int64) *CreateModelRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateModelRequest) SetModelName(v string) *CreateModelRequest {
	s.ModelName = &v
	return s
}

func (s *CreateModelRequest) SetDatasetIdList(v map[string]interface{}) *CreateModelRequest {
	s.DatasetIdList = v
	return s
}

func (s *CreateModelRequest) SetTestDatasetIdList(v map[string]interface{}) *CreateModelRequest {
	s.TestDatasetIdList = v
	return s
}

func (s *CreateModelRequest) SetIsIncrementalTrain(v string) *CreateModelRequest {
	s.IsIncrementalTrain = &v
	return s
}

func (s *CreateModelRequest) SetProduct(v string) *CreateModelRequest {
	s.Product = &v
	return s
}

type CreateModelShrinkRequest struct {
	ModelId                 *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelType               *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	ProjectId               *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ModelName               *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	DatasetIdListShrink     *string `json:"DatasetIdList,omitempty" xml:"DatasetIdList,omitempty"`
	TestDatasetIdListShrink *string `json:"TestDatasetIdList,omitempty" xml:"TestDatasetIdList,omitempty"`
	IsIncrementalTrain      *string `json:"IsIncrementalTrain,omitempty" xml:"IsIncrementalTrain,omitempty"`
	Product                 *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s CreateModelShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateModelShrinkRequest) SetModelId(v int64) *CreateModelShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *CreateModelShrinkRequest) SetModelType(v string) *CreateModelShrinkRequest {
	s.ModelType = &v
	return s
}

func (s *CreateModelShrinkRequest) SetProjectId(v int64) *CreateModelShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateModelShrinkRequest) SetModelName(v string) *CreateModelShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *CreateModelShrinkRequest) SetDatasetIdListShrink(v string) *CreateModelShrinkRequest {
	s.DatasetIdListShrink = &v
	return s
}

func (s *CreateModelShrinkRequest) SetTestDatasetIdListShrink(v string) *CreateModelShrinkRequest {
	s.TestDatasetIdListShrink = &v
	return s
}

func (s *CreateModelShrinkRequest) SetIsIncrementalTrain(v string) *CreateModelShrinkRequest {
	s.IsIncrementalTrain = &v
	return s
}

func (s *CreateModelShrinkRequest) SetProduct(v string) *CreateModelShrinkRequest {
	s.Product = &v
	return s
}

type CreateModelResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBody) SetMessage(v string) *CreateModelResponseBody {
	s.Message = &v
	return s
}

func (s *CreateModelResponseBody) SetRequestId(v string) *CreateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelResponseBody) SetData(v map[string]interface{}) *CreateModelResponseBody {
	s.Data = v
	return s
}

func (s *CreateModelResponseBody) SetCode(v int32) *CreateModelResponseBody {
	s.Code = &v
	return s
}

func (s *CreateModelResponseBody) SetSuccess(v bool) *CreateModelResponseBody {
	s.Success = &v
	return s
}

type CreateModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateModelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponse) GoString() string {
	return s.String()
}

func (s *CreateModelResponse) SetHeaders(v map[string]*string) *CreateModelResponse {
	s.Headers = v
	return s
}

func (s *CreateModelResponse) SetBody(v *CreateModelResponseBody) *CreateModelResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	ProjectType        *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	ProjectName        *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	Product            *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetProjectType(v string) *CreateProjectRequest {
	s.ProjectType = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetProjectDescription(v string) *CreateProjectRequest {
	s.ProjectDescription = &v
	return s
}

func (s *CreateProjectRequest) SetProduct(v string) *CreateProjectRequest {
	s.Product = &v
	return s
}

type CreateProjectResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProjectId map[string]interface{} `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetMessage(v string) *CreateProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetProjectId(v map[string]interface{}) *CreateProjectResponseBody {
	s.ProjectId = v
	return s
}

func (s *CreateProjectResponseBody) SetCode(v int32) *CreateProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v bool) *CreateProjectResponseBody {
	s.Success = &v
	return s
}

type CreateProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type DeleteModelRequest struct {
	ModelId   *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ProjectId *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Product   *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s DeleteModelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelRequest) SetModelId(v int64) *DeleteModelRequest {
	s.ModelId = &v
	return s
}

func (s *DeleteModelRequest) SetProjectId(v int64) *DeleteModelRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteModelRequest) SetProduct(v string) *DeleteModelRequest {
	s.Product = &v
	return s
}

type DeleteModelResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBody) SetMessage(v string) *DeleteModelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteModelResponseBody) SetRequestId(v string) *DeleteModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelResponseBody) SetData(v map[string]interface{}) *DeleteModelResponseBody {
	s.Data = v
	return s
}

func (s *DeleteModelResponseBody) SetCode(v int32) *DeleteModelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteModelResponseBody) SetSuccess(v bool) *DeleteModelResponseBody {
	s.Success = &v
	return s
}

type DeleteModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelResponse) SetHeaders(v map[string]*string) *DeleteModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelResponse) SetBody(v *DeleteModelResponseBody) *DeleteModelResponse {
	s.Body = v
	return s
}

type DeployModelRequest struct {
	ModelId      *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	OptType      *string `json:"OptType,omitempty" xml:"OptType,omitempty"`
	ProjectId    *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s DeployModelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployModelRequest) GoString() string {
	return s.String()
}

func (s *DeployModelRequest) SetModelId(v int64) *DeployModelRequest {
	s.ModelId = &v
	return s
}

func (s *DeployModelRequest) SetOptType(v string) *DeployModelRequest {
	s.OptType = &v
	return s
}

func (s *DeployModelRequest) SetProjectId(v int64) *DeployModelRequest {
	s.ProjectId = &v
	return s
}

func (s *DeployModelRequest) SetModelVersion(v string) *DeployModelRequest {
	s.ModelVersion = &v
	return s
}

func (s *DeployModelRequest) SetProduct(v string) *DeployModelRequest {
	s.Product = &v
	return s
}

type DeployModelResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeployModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeployModelResponseBody) SetMessage(v string) *DeployModelResponseBody {
	s.Message = &v
	return s
}

func (s *DeployModelResponseBody) SetRequestId(v string) *DeployModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployModelResponseBody) SetData(v map[string]interface{}) *DeployModelResponseBody {
	s.Data = v
	return s
}

func (s *DeployModelResponseBody) SetCode(v int32) *DeployModelResponseBody {
	s.Code = &v
	return s
}

func (s *DeployModelResponseBody) SetSuccess(v bool) *DeployModelResponseBody {
	s.Success = &v
	return s
}

type DeployModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployModelResponse) GoString() string {
	return s.String()
}

func (s *DeployModelResponse) SetHeaders(v map[string]*string) *DeployModelResponse {
	s.Headers = v
	return s
}

func (s *DeployModelResponse) SetBody(v *DeployModelResponseBody) *DeployModelResponse {
	s.Body = v
	return s
}

type GetAsyncPredictRequest struct {
	AsyncPredictId *int32  `json:"AsyncPredictId,omitempty" xml:"AsyncPredictId,omitempty"`
	Product        *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s GetAsyncPredictRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncPredictRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncPredictRequest) SetAsyncPredictId(v int32) *GetAsyncPredictRequest {
	s.AsyncPredictId = &v
	return s
}

func (s *GetAsyncPredictRequest) SetProduct(v string) *GetAsyncPredictRequest {
	s.Product = &v
	return s
}

type GetAsyncPredictResponseBody struct {
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	AsyncPredictId *int32  `json:"AsyncPredictId,omitempty" xml:"AsyncPredictId,omitempty"`
}

func (s GetAsyncPredictResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncPredictResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncPredictResponseBody) SetStatus(v int32) *GetAsyncPredictResponseBody {
	s.Status = &v
	return s
}

func (s *GetAsyncPredictResponseBody) SetRequestId(v string) *GetAsyncPredictResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncPredictResponseBody) SetContent(v string) *GetAsyncPredictResponseBody {
	s.Content = &v
	return s
}

func (s *GetAsyncPredictResponseBody) SetAsyncPredictId(v int32) *GetAsyncPredictResponseBody {
	s.AsyncPredictId = &v
	return s
}

type GetAsyncPredictResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAsyncPredictResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncPredictResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncPredictResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncPredictResponse) SetHeaders(v map[string]*string) *GetAsyncPredictResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncPredictResponse) SetBody(v *GetAsyncPredictResponseBody) *GetAsyncPredictResponse {
	s.Body = v
	return s
}

type GetModelRequest struct {
	ModelId      *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ProjectId    *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s GetModelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetModelRequest) GoString() string {
	return s.String()
}

func (s *GetModelRequest) SetModelId(v int64) *GetModelRequest {
	s.ModelId = &v
	return s
}

func (s *GetModelRequest) SetProjectId(v int64) *GetModelRequest {
	s.ProjectId = &v
	return s
}

func (s *GetModelRequest) SetModelVersion(v string) *GetModelRequest {
	s.ModelVersion = &v
	return s
}

func (s *GetModelRequest) SetProduct(v string) *GetModelRequest {
	s.Product = &v
	return s
}

type GetModelResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelResponseBody) SetMessage(v string) *GetModelResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelResponseBody) SetRequestId(v string) *GetModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelResponseBody) SetData(v map[string]interface{}) *GetModelResponseBody {
	s.Data = v
	return s
}

func (s *GetModelResponseBody) SetCode(v int32) *GetModelResponseBody {
	s.Code = &v
	return s
}

func (s *GetModelResponseBody) SetSuccess(v bool) *GetModelResponseBody {
	s.Success = &v
	return s
}

type GetModelResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetModelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelResponse) GoString() string {
	return s.String()
}

func (s *GetModelResponse) SetHeaders(v map[string]*string) *GetModelResponse {
	s.Headers = v
	return s
}

func (s *GetModelResponse) SetBody(v *GetModelResponseBody) *GetModelResponse {
	s.Body = v
	return s
}

type GetPredictResultRequest struct {
	ModelId      *int32  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	DetailTag    *string `json:"DetailTag,omitempty" xml:"DetailTag,omitempty"`
	TopK         *int32  `json:"TopK,omitempty" xml:"TopK,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s GetPredictResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPredictResultRequest) GoString() string {
	return s.String()
}

func (s *GetPredictResultRequest) SetModelId(v int32) *GetPredictResultRequest {
	s.ModelId = &v
	return s
}

func (s *GetPredictResultRequest) SetContent(v string) *GetPredictResultRequest {
	s.Content = &v
	return s
}

func (s *GetPredictResultRequest) SetModelVersion(v string) *GetPredictResultRequest {
	s.ModelVersion = &v
	return s
}

func (s *GetPredictResultRequest) SetDetailTag(v string) *GetPredictResultRequest {
	s.DetailTag = &v
	return s
}

func (s *GetPredictResultRequest) SetTopK(v int32) *GetPredictResultRequest {
	s.TopK = &v
	return s
}

func (s *GetPredictResultRequest) SetProduct(v string) *GetPredictResultRequest {
	s.Product = &v
	return s
}

type GetPredictResultResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s GetPredictResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPredictResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetPredictResultResponseBody) SetRequestId(v string) *GetPredictResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPredictResultResponseBody) SetContent(v string) *GetPredictResultResponseBody {
	s.Content = &v
	return s
}

type GetPredictResultResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPredictResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPredictResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPredictResultResponse) GoString() string {
	return s.String()
}

func (s *GetPredictResultResponse) SetHeaders(v map[string]*string) *GetPredictResultResponse {
	s.Headers = v
	return s
}

func (s *GetPredictResultResponse) SetBody(v *GetPredictResultResponseBody) *GetPredictResultResponse {
	s.Body = v
	return s
}

type ListDatasetRequest struct {
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s ListDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetRequest) SetProjectId(v int64) *ListDatasetRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDatasetRequest) SetPageSize(v int32) *ListDatasetRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetRequest) SetPageNumber(v int32) *ListDatasetRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetRequest) SetProduct(v string) *ListDatasetRequest {
	s.Product = &v
	return s
}

type ListDatasetResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetResponseBody) SetMessage(v string) *ListDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *ListDatasetResponseBody) SetRequestId(v string) *ListDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetResponseBody) SetData(v map[string]interface{}) *ListDatasetResponseBody {
	s.Data = v
	return s
}

func (s *ListDatasetResponseBody) SetCode(v int32) *ListDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *ListDatasetResponseBody) SetSuccess(v bool) *ListDatasetResponseBody {
	s.Success = &v
	return s
}

type ListDatasetResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetResponse) SetHeaders(v map[string]*string) *ListDatasetResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetResponse) SetBody(v *ListDatasetResponseBody) *ListDatasetResponse {
	s.Body = v
	return s
}

type ListModelsRequest struct {
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s ListModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelsRequest) SetProjectId(v int64) *ListModelsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListModelsRequest) SetPageSize(v int32) *ListModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelsRequest) SetPageNumber(v int32) *ListModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelsRequest) SetProduct(v string) *ListModelsRequest {
	s.Product = &v
	return s
}

type ListModelsResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBody) SetMessage(v string) *ListModelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListModelsResponseBody) SetRequestId(v string) *ListModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelsResponseBody) SetData(v map[string]interface{}) *ListModelsResponseBody {
	s.Data = v
	return s
}

func (s *ListModelsResponseBody) SetCode(v int32) *ListModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelsResponseBody) SetSuccess(v bool) *ListModelsResponseBody {
	s.Success = &v
	return s
}

type ListModelsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponse) GoString() string {
	return s.String()
}

func (s *ListModelsResponse) SetHeaders(v map[string]*string) *ListModelsResponse {
	s.Headers = v
	return s
}

func (s *ListModelsResponse) SetBody(v *ListModelsResponseBody) *ListModelsResponse {
	s.Body = v
	return s
}

type RunContactReviewRequest struct {
	ContactScene *string `json:"ContactScene,omitempty" xml:"ContactScene,omitempty"`
	ContactPath  *string `json:"ContactPath,omitempty" xml:"ContactPath,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s RunContactReviewRequest) String() string {
	return tea.Prettify(s)
}

func (s RunContactReviewRequest) GoString() string {
	return s.String()
}

func (s *RunContactReviewRequest) SetContactScene(v string) *RunContactReviewRequest {
	s.ContactScene = &v
	return s
}

func (s *RunContactReviewRequest) SetContactPath(v string) *RunContactReviewRequest {
	s.ContactPath = &v
	return s
}

func (s *RunContactReviewRequest) SetProduct(v string) *RunContactReviewRequest {
	s.Product = &v
	return s
}

type RunContactReviewResponseBody struct {
	RequestId          *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RawContractContent *string                                     `json:"RawContractContent,omitempty" xml:"RawContractContent,omitempty"`
	ContactContent     *RunContactReviewResponseBodyContactContent `json:"ContactContent,omitempty" xml:"ContactContent,omitempty" type:"Struct"`
}

func (s RunContactReviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunContactReviewResponseBody) GoString() string {
	return s.String()
}

func (s *RunContactReviewResponseBody) SetRequestId(v string) *RunContactReviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunContactReviewResponseBody) SetRawContractContent(v string) *RunContactReviewResponseBody {
	s.RawContractContent = &v
	return s
}

func (s *RunContactReviewResponseBody) SetContactContent(v *RunContactReviewResponseBodyContactContent) *RunContactReviewResponseBody {
	s.ContactContent = v
	return s
}

type RunContactReviewResponseBodyContactContent struct {
	ReviewResults    []*RunContactReviewResponseBodyContactContentReviewResults    `json:"ReviewResults,omitempty" xml:"ReviewResults,omitempty" type:"Repeated"`
	StructureResults []*RunContactReviewResponseBodyContactContentStructureResults `json:"StructureResults,omitempty" xml:"StructureResults,omitempty" type:"Repeated"`
}

func (s RunContactReviewResponseBodyContactContent) String() string {
	return tea.Prettify(s)
}

func (s RunContactReviewResponseBodyContactContent) GoString() string {
	return s.String()
}

func (s *RunContactReviewResponseBodyContactContent) SetReviewResults(v []*RunContactReviewResponseBodyContactContentReviewResults) *RunContactReviewResponseBodyContactContent {
	s.ReviewResults = v
	return s
}

func (s *RunContactReviewResponseBodyContactContent) SetStructureResults(v []*RunContactReviewResponseBodyContactContentStructureResults) *RunContactReviewResponseBodyContactContent {
	s.StructureResults = v
	return s
}

type RunContactReviewResponseBodyContactContentReviewResults struct {
	RiskLevel              *string   `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Type                   *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Value                  []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	StartPosition          []*string `json:"StartPosition,omitempty" xml:"StartPosition,omitempty" type:"Repeated"`
	ModificationSuggestion *string   `json:"ModificationSuggestion,omitempty" xml:"ModificationSuggestion,omitempty"`
	Reason                 *string   `json:"Reason,omitempty" xml:"Reason,omitempty"`
	EndPosition            []*string `json:"EndPosition,omitempty" xml:"EndPosition,omitempty" type:"Repeated"`
}

func (s RunContactReviewResponseBodyContactContentReviewResults) String() string {
	return tea.Prettify(s)
}

func (s RunContactReviewResponseBodyContactContentReviewResults) GoString() string {
	return s.String()
}

func (s *RunContactReviewResponseBodyContactContentReviewResults) SetRiskLevel(v string) *RunContactReviewResponseBodyContactContentReviewResults {
	s.RiskLevel = &v
	return s
}

func (s *RunContactReviewResponseBodyContactContentReviewResults) SetType(v string) *RunContactReviewResponseBodyContactContentReviewResults {
	s.Type = &v
	return s
}

func (s *RunContactReviewResponseBodyContactContentReviewResults) SetValue(v []*string) *RunContactReviewResponseBodyContactContentReviewResults {
	s.Value = v
	return s
}

func (s *RunContactReviewResponseBodyContactContentReviewResults) SetStartPosition(v []*string) *RunContactReviewResponseBodyContactContentReviewResults {
	s.StartPosition = v
	return s
}

func (s *RunContactReviewResponseBodyContactContentReviewResults) SetModificationSuggestion(v string) *RunContactReviewResponseBodyContactContentReviewResults {
	s.ModificationSuggestion = &v
	return s
}

func (s *RunContactReviewResponseBodyContactContentReviewResults) SetReason(v string) *RunContactReviewResponseBodyContactContentReviewResults {
	s.Reason = &v
	return s
}

func (s *RunContactReviewResponseBodyContactContentReviewResults) SetEndPosition(v []*string) *RunContactReviewResponseBodyContactContentReviewResults {
	s.EndPosition = v
	return s
}

type RunContactReviewResponseBodyContactContentStructureResults struct {
	Type          *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Value         []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	StartPosition []*string `json:"StartPosition,omitempty" xml:"StartPosition,omitempty" type:"Repeated"`
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	EndPosition   []*string `json:"EndPosition,omitempty" xml:"EndPosition,omitempty" type:"Repeated"`
}

func (s RunContactReviewResponseBodyContactContentStructureResults) String() string {
	return tea.Prettify(s)
}

func (s RunContactReviewResponseBodyContactContentStructureResults) GoString() string {
	return s.String()
}

func (s *RunContactReviewResponseBodyContactContentStructureResults) SetType(v string) *RunContactReviewResponseBodyContactContentStructureResults {
	s.Type = &v
	return s
}

func (s *RunContactReviewResponseBodyContactContentStructureResults) SetValue(v []*string) *RunContactReviewResponseBodyContactContentStructureResults {
	s.Value = v
	return s
}

func (s *RunContactReviewResponseBodyContactContentStructureResults) SetStartPosition(v []*string) *RunContactReviewResponseBodyContactContentStructureResults {
	s.StartPosition = v
	return s
}

func (s *RunContactReviewResponseBodyContactContentStructureResults) SetName(v string) *RunContactReviewResponseBodyContactContentStructureResults {
	s.Name = &v
	return s
}

func (s *RunContactReviewResponseBodyContactContentStructureResults) SetEndPosition(v []*string) *RunContactReviewResponseBodyContactContentStructureResults {
	s.EndPosition = v
	return s
}

type RunContactReviewResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunContactReviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunContactReviewResponse) String() string {
	return tea.Prettify(s)
}

func (s RunContactReviewResponse) GoString() string {
	return s.String()
}

func (s *RunContactReviewResponse) SetHeaders(v map[string]*string) *RunContactReviewResponse {
	s.Headers = v
	return s
}

func (s *RunContactReviewResponse) SetBody(v *RunContactReviewResponseBody) *RunContactReviewResponse {
	s.Body = v
	return s
}

type RunPreTrainServiceRequest struct {
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	PredictContent *string `json:"PredictContent,omitempty" xml:"PredictContent,omitempty"`
	Product        *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s RunPreTrainServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RunPreTrainServiceRequest) GoString() string {
	return s.String()
}

func (s *RunPreTrainServiceRequest) SetServiceName(v string) *RunPreTrainServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *RunPreTrainServiceRequest) SetServiceVersion(v string) *RunPreTrainServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *RunPreTrainServiceRequest) SetPredictContent(v string) *RunPreTrainServiceRequest {
	s.PredictContent = &v
	return s
}

func (s *RunPreTrainServiceRequest) SetProduct(v string) *RunPreTrainServiceRequest {
	s.Product = &v
	return s
}

type RunPreTrainServiceResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PredictResult *string `json:"PredictResult,omitempty" xml:"PredictResult,omitempty"`
}

func (s RunPreTrainServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunPreTrainServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RunPreTrainServiceResponseBody) SetRequestId(v string) *RunPreTrainServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunPreTrainServiceResponseBody) SetPredictResult(v string) *RunPreTrainServiceResponseBody {
	s.PredictResult = &v
	return s
}

type RunPreTrainServiceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunPreTrainServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunPreTrainServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RunPreTrainServiceResponse) GoString() string {
	return s.String()
}

func (s *RunPreTrainServiceResponse) SetHeaders(v map[string]*string) *RunPreTrainServiceResponse {
	s.Headers = v
	return s
}

func (s *RunPreTrainServiceResponse) SetBody(v *RunPreTrainServiceResponseBody) *RunPreTrainServiceResponse {
	s.Body = v
	return s
}

type RunSmartCallServiceRequest struct {
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ParamContent *string `json:"ParamContent,omitempty" xml:"ParamContent,omitempty"`
	RobotId      *int64  `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s RunSmartCallServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RunSmartCallServiceRequest) GoString() string {
	return s.String()
}

func (s *RunSmartCallServiceRequest) SetServiceName(v string) *RunSmartCallServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *RunSmartCallServiceRequest) SetParamContent(v string) *RunSmartCallServiceRequest {
	s.ParamContent = &v
	return s
}

func (s *RunSmartCallServiceRequest) SetRobotId(v int64) *RunSmartCallServiceRequest {
	s.RobotId = &v
	return s
}

func (s *RunSmartCallServiceRequest) SetSessionId(v string) *RunSmartCallServiceRequest {
	s.SessionId = &v
	return s
}

func (s *RunSmartCallServiceRequest) SetProduct(v string) *RunSmartCallServiceRequest {
	s.Product = &v
	return s
}

type RunSmartCallServiceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RunSmartCallServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunSmartCallServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RunSmartCallServiceResponseBody) SetMessage(v string) *RunSmartCallServiceResponseBody {
	s.Message = &v
	return s
}

func (s *RunSmartCallServiceResponseBody) SetRequestId(v string) *RunSmartCallServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSmartCallServiceResponseBody) SetData(v string) *RunSmartCallServiceResponseBody {
	s.Data = &v
	return s
}

func (s *RunSmartCallServiceResponseBody) SetCode(v int32) *RunSmartCallServiceResponseBody {
	s.Code = &v
	return s
}

type RunSmartCallServiceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunSmartCallServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunSmartCallServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RunSmartCallServiceResponse) GoString() string {
	return s.String()
}

func (s *RunSmartCallServiceResponse) SetHeaders(v map[string]*string) *RunSmartCallServiceResponse {
	s.Headers = v
	return s
}

func (s *RunSmartCallServiceResponse) SetBody(v *RunSmartCallServiceResponseBody) *RunSmartCallServiceResponse {
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
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("nlp-automl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateAsyncPredictWithOptions(request *CreateAsyncPredictRequest, runtime *util.RuntimeOptions) (_result *CreateAsyncPredictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAsyncPredictResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAsyncPredict"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAsyncPredict(request *CreateAsyncPredictRequest) (_result *CreateAsyncPredictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAsyncPredictResponse{}
	_body, _err := client.CreateAsyncPredictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDatasetWithOptions(request *CreateDatasetRequest, runtime *util.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDataset"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataset(request *CreateDatasetRequest) (_result *CreateDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CreateDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDatasetRecordWithOptions(request *CreateDatasetRecordRequest, runtime *util.RuntimeOptions) (_result *CreateDatasetRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDatasetRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDatasetRecord"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDatasetRecord(request *CreateDatasetRecordRequest) (_result *CreateDatasetRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDatasetRecordResponse{}
	_body, _err := client.CreateDatasetRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateModelWithOptions(tmpReq *CreateModelRequest, runtime *util.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateModelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DatasetIdList)) {
		request.DatasetIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DatasetIdList, tea.String("DatasetIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TestDatasetIdList)) {
		request.TestDatasetIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TestDatasetIdList, tea.String("TestDatasetIdList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateModel"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModel(request *CreateModelRequest) (_result *CreateModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateModelResponse{}
	_body, _err := client.CreateModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateProject"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteModelWithOptions(request *DeleteModelRequest, runtime *util.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteModel"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteModel(request *DeleteModelRequest) (_result *DeleteModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteModelResponse{}
	_body, _err := client.DeleteModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployModelWithOptions(request *DeployModelRequest, runtime *util.RuntimeOptions) (_result *DeployModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeployModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeployModel"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployModel(request *DeployModelRequest) (_result *DeployModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployModelResponse{}
	_body, _err := client.DeployModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncPredictWithOptions(request *GetAsyncPredictRequest, runtime *util.RuntimeOptions) (_result *GetAsyncPredictResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAsyncPredictResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAsyncPredict"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncPredict(request *GetAsyncPredictRequest) (_result *GetAsyncPredictResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncPredictResponse{}
	_body, _err := client.GetAsyncPredictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModelWithOptions(request *GetModelRequest, runtime *util.RuntimeOptions) (_result *GetModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetModelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetModel"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModel(request *GetModelRequest) (_result *GetModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetModelResponse{}
	_body, _err := client.GetModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPredictResultWithOptions(request *GetPredictResultRequest, runtime *util.RuntimeOptions) (_result *GetPredictResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPredictResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPredictResult"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPredictResult(request *GetPredictResultRequest) (_result *GetPredictResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPredictResultResponse{}
	_body, _err := client.GetPredictResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDatasetWithOptions(request *ListDatasetRequest, runtime *util.RuntimeOptions) (_result *ListDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDatasetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDataset"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataset(request *ListDatasetRequest) (_result *ListDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDatasetResponse{}
	_body, _err := client.ListDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModelsWithOptions(request *ListModelsRequest, runtime *util.RuntimeOptions) (_result *ListModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListModelsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListModels"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModels(request *ListModelsRequest) (_result *ListModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListModelsResponse{}
	_body, _err := client.ListModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunContactReviewWithOptions(request *RunContactReviewRequest, runtime *util.RuntimeOptions) (_result *RunContactReviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RunContactReviewResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RunContactReview"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunContactReview(request *RunContactReviewRequest) (_result *RunContactReviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunContactReviewResponse{}
	_body, _err := client.RunContactReviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunPreTrainServiceWithOptions(request *RunPreTrainServiceRequest, runtime *util.RuntimeOptions) (_result *RunPreTrainServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RunPreTrainServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RunPreTrainService"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunPreTrainService(request *RunPreTrainServiceRequest) (_result *RunPreTrainServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunPreTrainServiceResponse{}
	_body, _err := client.RunPreTrainServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunSmartCallServiceWithOptions(request *RunSmartCallServiceRequest, runtime *util.RuntimeOptions) (_result *RunSmartCallServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RunSmartCallServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RunSmartCallService"), tea.String("2019-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunSmartCallService(request *RunSmartCallServiceRequest) (_result *RunSmartCallServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunSmartCallServiceResponse{}
	_body, _err := client.RunSmartCallServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
