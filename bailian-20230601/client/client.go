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

type CancelFineTuneJobRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CancelFineTuneJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelFineTuneJobRequest) GoString() string {
	return s.String()
}

func (s *CancelFineTuneJobRequest) SetAgentKey(v string) *CancelFineTuneJobRequest {
	s.AgentKey = &v
	return s
}

func (s *CancelFineTuneJobRequest) SetJobId(v string) *CancelFineTuneJobRequest {
	s.JobId = &v
	return s
}

type CancelFineTuneJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelFineTuneJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelFineTuneJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelFineTuneJobResponseBody) SetJobId(v string) *CancelFineTuneJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CancelFineTuneJobResponseBody) SetRequestId(v string) *CancelFineTuneJobResponseBody {
	s.RequestId = &v
	return s
}

type CancelFineTuneJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelFineTuneJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelFineTuneJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelFineTuneJobResponse) GoString() string {
	return s.String()
}

func (s *CancelFineTuneJobResponse) SetHeaders(v map[string]*string) *CancelFineTuneJobResponse {
	s.Headers = v
	return s
}

func (s *CancelFineTuneJobResponse) SetStatusCode(v int32) *CancelFineTuneJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelFineTuneJobResponse) SetBody(v *CancelFineTuneJobResponseBody) *CancelFineTuneJobResponse {
	s.Body = v
	return s
}

type CreateFineTuneJobRequest struct {
	AgentKey        *string                                  `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	BaseModel       *string                                  `json:"BaseModel,omitempty" xml:"BaseModel,omitempty"`
	HyperParameters *CreateFineTuneJobRequestHyperParameters `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Struct"`
	ModelName       *string                                  `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	TrainingFiles   []*string                                `json:"TrainingFiles,omitempty" xml:"TrainingFiles,omitempty" type:"Repeated"`
	TrainingType    *string                                  `json:"TrainingType,omitempty" xml:"TrainingType,omitempty"`
	ValidationFiles []*string                                `json:"ValidationFiles,omitempty" xml:"ValidationFiles,omitempty" type:"Repeated"`
}

func (s CreateFineTuneJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFineTuneJobRequest) GoString() string {
	return s.String()
}

func (s *CreateFineTuneJobRequest) SetAgentKey(v string) *CreateFineTuneJobRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateFineTuneJobRequest) SetBaseModel(v string) *CreateFineTuneJobRequest {
	s.BaseModel = &v
	return s
}

func (s *CreateFineTuneJobRequest) SetHyperParameters(v *CreateFineTuneJobRequestHyperParameters) *CreateFineTuneJobRequest {
	s.HyperParameters = v
	return s
}

func (s *CreateFineTuneJobRequest) SetModelName(v string) *CreateFineTuneJobRequest {
	s.ModelName = &v
	return s
}

func (s *CreateFineTuneJobRequest) SetTrainingFiles(v []*string) *CreateFineTuneJobRequest {
	s.TrainingFiles = v
	return s
}

func (s *CreateFineTuneJobRequest) SetTrainingType(v string) *CreateFineTuneJobRequest {
	s.TrainingType = &v
	return s
}

func (s *CreateFineTuneJobRequest) SetValidationFiles(v []*string) *CreateFineTuneJobRequest {
	s.ValidationFiles = v
	return s
}

type CreateFineTuneJobRequestHyperParameters struct {
	BatchSize        *int32   `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	Epochs           *int32   `json:"Epochs,omitempty" xml:"Epochs,omitempty"`
	LearningRate     *string  `json:"LearningRate,omitempty" xml:"LearningRate,omitempty"`
	PromptLossWeight *float64 `json:"PromptLossWeight,omitempty" xml:"PromptLossWeight,omitempty"`
}

func (s CreateFineTuneJobRequestHyperParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateFineTuneJobRequestHyperParameters) GoString() string {
	return s.String()
}

func (s *CreateFineTuneJobRequestHyperParameters) SetBatchSize(v int32) *CreateFineTuneJobRequestHyperParameters {
	s.BatchSize = &v
	return s
}

func (s *CreateFineTuneJobRequestHyperParameters) SetEpochs(v int32) *CreateFineTuneJobRequestHyperParameters {
	s.Epochs = &v
	return s
}

func (s *CreateFineTuneJobRequestHyperParameters) SetLearningRate(v string) *CreateFineTuneJobRequestHyperParameters {
	s.LearningRate = &v
	return s
}

func (s *CreateFineTuneJobRequestHyperParameters) SetPromptLossWeight(v float64) *CreateFineTuneJobRequestHyperParameters {
	s.PromptLossWeight = &v
	return s
}

type CreateFineTuneJobShrinkRequest struct {
	AgentKey              *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	BaseModel             *string `json:"BaseModel,omitempty" xml:"BaseModel,omitempty"`
	HyperParametersShrink *string `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty"`
	ModelName             *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	TrainingFilesShrink   *string `json:"TrainingFiles,omitempty" xml:"TrainingFiles,omitempty"`
	TrainingType          *string `json:"TrainingType,omitempty" xml:"TrainingType,omitempty"`
	ValidationFilesShrink *string `json:"ValidationFiles,omitempty" xml:"ValidationFiles,omitempty"`
}

func (s CreateFineTuneJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFineTuneJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFineTuneJobShrinkRequest) SetAgentKey(v string) *CreateFineTuneJobShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateFineTuneJobShrinkRequest) SetBaseModel(v string) *CreateFineTuneJobShrinkRequest {
	s.BaseModel = &v
	return s
}

func (s *CreateFineTuneJobShrinkRequest) SetHyperParametersShrink(v string) *CreateFineTuneJobShrinkRequest {
	s.HyperParametersShrink = &v
	return s
}

func (s *CreateFineTuneJobShrinkRequest) SetModelName(v string) *CreateFineTuneJobShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *CreateFineTuneJobShrinkRequest) SetTrainingFilesShrink(v string) *CreateFineTuneJobShrinkRequest {
	s.TrainingFilesShrink = &v
	return s
}

func (s *CreateFineTuneJobShrinkRequest) SetTrainingType(v string) *CreateFineTuneJobShrinkRequest {
	s.TrainingType = &v
	return s
}

func (s *CreateFineTuneJobShrinkRequest) SetValidationFilesShrink(v string) *CreateFineTuneJobShrinkRequest {
	s.ValidationFilesShrink = &v
	return s
}

type CreateFineTuneJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateFineTuneJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFineTuneJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFineTuneJobResponseBody) SetJobId(v string) *CreateFineTuneJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateFineTuneJobResponseBody) SetRequestId(v string) *CreateFineTuneJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFineTuneJobResponseBody) SetStatus(v string) *CreateFineTuneJobResponseBody {
	s.Status = &v
	return s
}

type CreateFineTuneJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFineTuneJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFineTuneJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFineTuneJobResponse) GoString() string {
	return s.String()
}

func (s *CreateFineTuneJobResponse) SetHeaders(v map[string]*string) *CreateFineTuneJobResponse {
	s.Headers = v
	return s
}

func (s *CreateFineTuneJobResponse) SetStatusCode(v int32) *CreateFineTuneJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFineTuneJobResponse) SetBody(v *CreateFineTuneJobResponseBody) *CreateFineTuneJobResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Model    *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetAgentKey(v string) *CreateServiceRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateServiceRequest) SetModel(v string) *CreateServiceRequest {
	s.Model = &v
	return s
}

type CreateServiceResponseBody struct {
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetModelServiceId(v string) *CreateServiceResponseBody {
	s.ModelServiceId = &v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetStatusCode(v int32) *CreateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type CreateTextEmbeddingsRequest struct {
	AgentKey *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Input    []*string `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	TextType *string   `json:"TextType,omitempty" xml:"TextType,omitempty"`
}

func (s CreateTextEmbeddingsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTextEmbeddingsRequest) GoString() string {
	return s.String()
}

func (s *CreateTextEmbeddingsRequest) SetAgentKey(v string) *CreateTextEmbeddingsRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateTextEmbeddingsRequest) SetInput(v []*string) *CreateTextEmbeddingsRequest {
	s.Input = v
	return s
}

func (s *CreateTextEmbeddingsRequest) SetTextType(v string) *CreateTextEmbeddingsRequest {
	s.TextType = &v
	return s
}

type CreateTextEmbeddingsShrinkRequest struct {
	AgentKey    *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	TextType    *string `json:"TextType,omitempty" xml:"TextType,omitempty"`
}

func (s CreateTextEmbeddingsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTextEmbeddingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTextEmbeddingsShrinkRequest) SetAgentKey(v string) *CreateTextEmbeddingsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateTextEmbeddingsShrinkRequest) SetInputShrink(v string) *CreateTextEmbeddingsShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *CreateTextEmbeddingsShrinkRequest) SetTextType(v string) *CreateTextEmbeddingsShrinkRequest {
	s.TextType = &v
	return s
}

type CreateTextEmbeddingsResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateTextEmbeddingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *string                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTextEmbeddingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTextEmbeddingsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTextEmbeddingsResponseBody) SetCode(v string) *CreateTextEmbeddingsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTextEmbeddingsResponseBody) SetData(v *CreateTextEmbeddingsResponseBodyData) *CreateTextEmbeddingsResponseBody {
	s.Data = v
	return s
}

func (s *CreateTextEmbeddingsResponseBody) SetHttpStatusCode(v string) *CreateTextEmbeddingsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTextEmbeddingsResponseBody) SetMessage(v string) *CreateTextEmbeddingsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTextEmbeddingsResponseBody) SetRequestId(v string) *CreateTextEmbeddingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTextEmbeddingsResponseBody) SetSuccess(v bool) *CreateTextEmbeddingsResponseBody {
	s.Success = &v
	return s
}

type CreateTextEmbeddingsResponseBodyData struct {
	Embeddings []*CreateTextEmbeddingsResponseBodyDataEmbeddings `json:"Embeddings,omitempty" xml:"Embeddings,omitempty" type:"Repeated"`
}

func (s CreateTextEmbeddingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTextEmbeddingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTextEmbeddingsResponseBodyData) SetEmbeddings(v []*CreateTextEmbeddingsResponseBodyDataEmbeddings) *CreateTextEmbeddingsResponseBodyData {
	s.Embeddings = v
	return s
}

type CreateTextEmbeddingsResponseBodyDataEmbeddings struct {
	Embedding []*float64 `json:"Embedding,omitempty" xml:"Embedding,omitempty" type:"Repeated"`
	TextIndex *int32     `json:"TextIndex,omitempty" xml:"TextIndex,omitempty"`
}

func (s CreateTextEmbeddingsResponseBodyDataEmbeddings) String() string {
	return tea.Prettify(s)
}

func (s CreateTextEmbeddingsResponseBodyDataEmbeddings) GoString() string {
	return s.String()
}

func (s *CreateTextEmbeddingsResponseBodyDataEmbeddings) SetEmbedding(v []*float64) *CreateTextEmbeddingsResponseBodyDataEmbeddings {
	s.Embedding = v
	return s
}

func (s *CreateTextEmbeddingsResponseBodyDataEmbeddings) SetTextIndex(v int32) *CreateTextEmbeddingsResponseBodyDataEmbeddings {
	s.TextIndex = &v
	return s
}

type CreateTextEmbeddingsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTextEmbeddingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTextEmbeddingsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTextEmbeddingsResponse) GoString() string {
	return s.String()
}

func (s *CreateTextEmbeddingsResponse) SetHeaders(v map[string]*string) *CreateTextEmbeddingsResponse {
	s.Headers = v
	return s
}

func (s *CreateTextEmbeddingsResponse) SetStatusCode(v int32) *CreateTextEmbeddingsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTextEmbeddingsResponse) SetBody(v *CreateTextEmbeddingsResponseBody) *CreateTextEmbeddingsResponse {
	s.Body = v
	return s
}

type CreateTokenRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s CreateTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenRequest) SetAgentKey(v string) *CreateTokenRequest {
	s.AgentKey = &v
	return s
}

type CreateTokenResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *string                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBody) SetCode(v string) *CreateTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTokenResponseBody) SetData(v *CreateTokenResponseBodyData) *CreateTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateTokenResponseBody) SetHttpStatusCode(v string) *CreateTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTokenResponseBody) SetMessage(v string) *CreateTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTokenResponseBody) SetRequestId(v string) *CreateTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTokenResponseBody) SetSuccess(v bool) *CreateTokenResponseBody {
	s.Success = &v
	return s
}

type CreateTokenResponseBodyData struct {
	ExpiredTime *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Token       *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBodyData) SetExpiredTime(v int64) *CreateTokenResponseBodyData {
	s.ExpiredTime = &v
	return s
}

func (s *CreateTokenResponseBodyData) SetToken(v string) *CreateTokenResponseBodyData {
	s.Token = &v
	return s
}

type CreateTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateTokenResponse) SetHeaders(v map[string]*string) *CreateTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateTokenResponse) SetStatusCode(v int32) *CreateTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTokenResponse) SetBody(v *CreateTokenResponseBody) *CreateTokenResponse {
	s.Body = v
	return s
}

type DeleteEnterpriseDataRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DataId   *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
}

func (s DeleteEnterpriseDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseDataRequest) SetAgentKey(v string) *DeleteEnterpriseDataRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteEnterpriseDataRequest) SetDataId(v string) *DeleteEnterpriseDataRequest {
	s.DataId = &v
	return s
}

type DeleteEnterpriseDataResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEnterpriseDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseDataResponseBody) SetData(v bool) *DeleteEnterpriseDataResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteEnterpriseDataResponseBody) SetErrorCode(v string) *DeleteEnterpriseDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteEnterpriseDataResponseBody) SetErrorMsg(v string) *DeleteEnterpriseDataResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteEnterpriseDataResponseBody) SetRequestId(v string) *DeleteEnterpriseDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnterpriseDataResponseBody) SetSuccess(v bool) *DeleteEnterpriseDataResponseBody {
	s.Success = &v
	return s
}

type DeleteEnterpriseDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEnterpriseDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEnterpriseDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnterpriseDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseDataResponse) SetHeaders(v map[string]*string) *DeleteEnterpriseDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnterpriseDataResponse) SetStatusCode(v int32) *DeleteEnterpriseDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnterpriseDataResponse) SetBody(v *DeleteEnterpriseDataResponseBody) *DeleteEnterpriseDataResponse {
	s.Body = v
	return s
}

type DeleteFineTuneJobRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteFineTuneJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFineTuneJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteFineTuneJobRequest) SetAgentKey(v string) *DeleteFineTuneJobRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteFineTuneJobRequest) SetJobId(v string) *DeleteFineTuneJobRequest {
	s.JobId = &v
	return s
}

type DeleteFineTuneJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFineTuneJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFineTuneJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFineTuneJobResponseBody) SetJobId(v string) *DeleteFineTuneJobResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteFineTuneJobResponseBody) SetRequestId(v string) *DeleteFineTuneJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFineTuneJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFineTuneJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFineTuneJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFineTuneJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteFineTuneJobResponse) SetHeaders(v map[string]*string) *DeleteFineTuneJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteFineTuneJobResponse) SetStatusCode(v int32) *DeleteFineTuneJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFineTuneJobResponse) SetBody(v *DeleteFineTuneJobResponseBody) *DeleteFineTuneJobResponse {
	s.Body = v
	return s
}

type DeleteServiceRequest struct {
	AgentKey       *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
}

func (s DeleteServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) SetAgentKey(v string) *DeleteServiceRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteServiceRequest) SetModelServiceId(v string) *DeleteServiceRequest {
	s.ModelServiceId = &v
	return s
}

type DeleteServiceResponseBody struct {
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) SetModelServiceId(v string) *DeleteServiceResponseBody {
	s.ModelServiceId = &v
	return s
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetStatusCode(v int32) *DeleteServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
	s.Body = v
	return s
}

type DescribeFineTuneJobRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeFineTuneJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFineTuneJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeFineTuneJobRequest) SetAgentKey(v string) *DescribeFineTuneJobRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeFineTuneJobRequest) SetJobId(v string) *DescribeFineTuneJobRequest {
	s.JobId = &v
	return s
}

type DescribeFineTuneJobResponseBody struct {
	BaseModel       *string                                         `json:"BaseModel,omitempty" xml:"BaseModel,omitempty"`
	FineTunedModel  *string                                         `json:"FineTunedModel,omitempty" xml:"FineTunedModel,omitempty"`
	HyperParameters *DescribeFineTuneJobResponseBodyHyperParameters `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Struct"`
	JobId           *string                                         `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message         *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	ModelName       *string                                         `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status          *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	TrainingFiles   []*string                                       `json:"TrainingFiles,omitempty" xml:"TrainingFiles,omitempty" type:"Repeated"`
	TrainingType    *string                                         `json:"TrainingType,omitempty" xml:"TrainingType,omitempty"`
	ValidationFiles []*string                                       `json:"ValidationFiles,omitempty" xml:"ValidationFiles,omitempty" type:"Repeated"`
}

func (s DescribeFineTuneJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFineTuneJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFineTuneJobResponseBody) SetBaseModel(v string) *DescribeFineTuneJobResponseBody {
	s.BaseModel = &v
	return s
}

func (s *DescribeFineTuneJobResponseBody) SetFineTunedModel(v string) *DescribeFineTuneJobResponseBody {
	s.FineTunedModel = &v
	return s
}

func (s *DescribeFineTuneJobResponseBody) SetHyperParameters(v *DescribeFineTuneJobResponseBodyHyperParameters) *DescribeFineTuneJobResponseBody {
	s.HyperParameters = v
	return s
}

func (s *DescribeFineTuneJobResponseBody) SetJobId(v string) *DescribeFineTuneJobResponseBody {
	s.JobId = &v
	return s
}

func (s *DescribeFineTuneJobResponseBody) SetMessage(v string) *DescribeFineTuneJobResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFineTuneJobResponseBody) SetModelName(v string) *DescribeFineTuneJobResponseBody {
	s.ModelName = &v
	return s
}

func (s *DescribeFineTuneJobResponseBody) SetRequestId(v string) *DescribeFineTuneJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFineTuneJobResponseBody) SetStatus(v string) *DescribeFineTuneJobResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFineTuneJobResponseBody) SetTrainingFiles(v []*string) *DescribeFineTuneJobResponseBody {
	s.TrainingFiles = v
	return s
}

func (s *DescribeFineTuneJobResponseBody) SetTrainingType(v string) *DescribeFineTuneJobResponseBody {
	s.TrainingType = &v
	return s
}

func (s *DescribeFineTuneJobResponseBody) SetValidationFiles(v []*string) *DescribeFineTuneJobResponseBody {
	s.ValidationFiles = v
	return s
}

type DescribeFineTuneJobResponseBodyHyperParameters struct {
	BatchSize        *int32   `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	Epochs           *int32   `json:"Epochs,omitempty" xml:"Epochs,omitempty"`
	LearningRate     *string  `json:"LearningRate,omitempty" xml:"LearningRate,omitempty"`
	PromptLossWeight *float64 `json:"PromptLossWeight,omitempty" xml:"PromptLossWeight,omitempty"`
}

func (s DescribeFineTuneJobResponseBodyHyperParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeFineTuneJobResponseBodyHyperParameters) GoString() string {
	return s.String()
}

func (s *DescribeFineTuneJobResponseBodyHyperParameters) SetBatchSize(v int32) *DescribeFineTuneJobResponseBodyHyperParameters {
	s.BatchSize = &v
	return s
}

func (s *DescribeFineTuneJobResponseBodyHyperParameters) SetEpochs(v int32) *DescribeFineTuneJobResponseBodyHyperParameters {
	s.Epochs = &v
	return s
}

func (s *DescribeFineTuneJobResponseBodyHyperParameters) SetLearningRate(v string) *DescribeFineTuneJobResponseBodyHyperParameters {
	s.LearningRate = &v
	return s
}

func (s *DescribeFineTuneJobResponseBodyHyperParameters) SetPromptLossWeight(v float64) *DescribeFineTuneJobResponseBodyHyperParameters {
	s.PromptLossWeight = &v
	return s
}

type DescribeFineTuneJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFineTuneJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFineTuneJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFineTuneJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeFineTuneJobResponse) SetHeaders(v map[string]*string) *DescribeFineTuneJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeFineTuneJobResponse) SetStatusCode(v int32) *DescribeFineTuneJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFineTuneJobResponse) SetBody(v *DescribeFineTuneJobResponseBody) *DescribeFineTuneJobResponse {
	s.Body = v
	return s
}

type DescribeServiceRequest struct {
	AgentKey       *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
}

func (s DescribeServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceRequest) SetAgentKey(v string) *DescribeServiceRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeServiceRequest) SetModelServiceId(v string) *DescribeServiceRequest {
	s.ModelServiceId = &v
	return s
}

type DescribeServiceResponseBody struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseBody) SetAppId(v string) *DescribeServiceResponseBody {
	s.AppId = &v
	return s
}

func (s *DescribeServiceResponseBody) SetModelServiceId(v string) *DescribeServiceResponseBody {
	s.ModelServiceId = &v
	return s
}

func (s *DescribeServiceResponseBody) SetRequestId(v string) *DescribeServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceResponseBody) SetStatus(v string) *DescribeServiceResponseBody {
	s.Status = &v
	return s
}

type DescribeServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponse) SetHeaders(v map[string]*string) *DescribeServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceResponse) SetStatusCode(v int32) *DescribeServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceResponse) SetBody(v *DescribeServiceResponseBody) *DescribeServiceResponse {
	s.Body = v
	return s
}

type GetEnterpriseDataByDataIdRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DataId   *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s GetEnterpriseDataByDataIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataByDataIdRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataByDataIdRequest) SetAgentKey(v string) *GetEnterpriseDataByDataIdRequest {
	s.AgentKey = &v
	return s
}

func (s *GetEnterpriseDataByDataIdRequest) SetDataId(v string) *GetEnterpriseDataByDataIdRequest {
	s.DataId = &v
	return s
}

func (s *GetEnterpriseDataByDataIdRequest) SetOwnerId(v int64) *GetEnterpriseDataByDataIdRequest {
	s.OwnerId = &v
	return s
}

type GetEnterpriseDataByDataIdResponseBody struct {
	Data      *GetEnterpriseDataByDataIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                                    `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEnterpriseDataByDataIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataByDataIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataByDataIdResponseBody) SetData(v *GetEnterpriseDataByDataIdResponseBodyData) *GetEnterpriseDataByDataIdResponseBody {
	s.Data = v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBody) SetErrorCode(v string) *GetEnterpriseDataByDataIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBody) SetErrorMsg(v string) *GetEnterpriseDataByDataIdResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBody) SetRequestId(v string) *GetEnterpriseDataByDataIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBody) SetSuccess(v bool) *GetEnterpriseDataByDataIdResponseBody {
	s.Success = &v
	return s
}

type GetEnterpriseDataByDataIdResponseBodyData struct {
	DataId         *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	DataName       *string `json:"DataName,omitempty" xml:"DataName,omitempty"`
	DataSize       *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	DataStatus     *string `json:"DataStatus,omitempty" xml:"DataStatus,omitempty"`
	DataStatusCode *int32  `json:"DataStatusCode,omitempty" xml:"DataStatusCode,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DataTypeCode   *int32  `json:"DataTypeCode,omitempty" xml:"DataTypeCode,omitempty"`
	DownloadUrl    *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	StatusDetail   *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	StoreType      *string `json:"StoreType,omitempty" xml:"StoreType,omitempty"`
	Tags           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UploadTime     *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s GetEnterpriseDataByDataIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataByDataIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetDataId(v string) *GetEnterpriseDataByDataIdResponseBodyData {
	s.DataId = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetDataName(v string) *GetEnterpriseDataByDataIdResponseBodyData {
	s.DataName = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetDataSize(v string) *GetEnterpriseDataByDataIdResponseBodyData {
	s.DataSize = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetDataStatus(v string) *GetEnterpriseDataByDataIdResponseBodyData {
	s.DataStatus = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetDataStatusCode(v int32) *GetEnterpriseDataByDataIdResponseBodyData {
	s.DataStatusCode = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetDataType(v string) *GetEnterpriseDataByDataIdResponseBodyData {
	s.DataType = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetDataTypeCode(v int32) *GetEnterpriseDataByDataIdResponseBodyData {
	s.DataTypeCode = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetDownloadUrl(v string) *GetEnterpriseDataByDataIdResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetStatusDetail(v string) *GetEnterpriseDataByDataIdResponseBodyData {
	s.StatusDetail = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetStoreType(v string) *GetEnterpriseDataByDataIdResponseBodyData {
	s.StoreType = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetTags(v string) *GetEnterpriseDataByDataIdResponseBodyData {
	s.Tags = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponseBodyData) SetUploadTime(v string) *GetEnterpriseDataByDataIdResponseBodyData {
	s.UploadTime = &v
	return s
}

type GetEnterpriseDataByDataIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEnterpriseDataByDataIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnterpriseDataByDataIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataByDataIdResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataByDataIdResponse) SetHeaders(v map[string]*string) *GetEnterpriseDataByDataIdResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseDataByDataIdResponse) SetStatusCode(v int32) *GetEnterpriseDataByDataIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnterpriseDataByDataIdResponse) SetBody(v *GetEnterpriseDataByDataIdResponseBody) *GetEnterpriseDataByDataIdResponse {
	s.Body = v
	return s
}

type GetEnterpriseDataChunkRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DataId   *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
}

func (s GetEnterpriseDataChunkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataChunkRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataChunkRequest) SetAgentKey(v string) *GetEnterpriseDataChunkRequest {
	s.AgentKey = &v
	return s
}

func (s *GetEnterpriseDataChunkRequest) SetDataId(v string) *GetEnterpriseDataChunkRequest {
	s.DataId = &v
	return s
}

type GetEnterpriseDataChunkResponseBody struct {
	Data      []*GetEnterpriseDataChunkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                                   `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEnterpriseDataChunkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataChunkResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataChunkResponseBody) SetData(v []*GetEnterpriseDataChunkResponseBodyData) *GetEnterpriseDataChunkResponseBody {
	s.Data = v
	return s
}

func (s *GetEnterpriseDataChunkResponseBody) SetErrorCode(v string) *GetEnterpriseDataChunkResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetEnterpriseDataChunkResponseBody) SetErrorMsg(v string) *GetEnterpriseDataChunkResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetEnterpriseDataChunkResponseBody) SetRequestId(v string) *GetEnterpriseDataChunkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnterpriseDataChunkResponseBody) SetSuccess(v bool) *GetEnterpriseDataChunkResponseBody {
	s.Success = &v
	return s
}

type GetEnterpriseDataChunkResponseBodyData struct {
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	TitlePath *string `json:"TitlePath,omitempty" xml:"TitlePath,omitempty"`
}

func (s GetEnterpriseDataChunkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataChunkResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataChunkResponseBodyData) SetText(v string) *GetEnterpriseDataChunkResponseBodyData {
	s.Text = &v
	return s
}

func (s *GetEnterpriseDataChunkResponseBodyData) SetTitle(v string) *GetEnterpriseDataChunkResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetEnterpriseDataChunkResponseBodyData) SetTitlePath(v string) *GetEnterpriseDataChunkResponseBodyData {
	s.TitlePath = &v
	return s
}

type GetEnterpriseDataChunkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEnterpriseDataChunkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnterpriseDataChunkResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataChunkResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataChunkResponse) SetHeaders(v map[string]*string) *GetEnterpriseDataChunkResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseDataChunkResponse) SetStatusCode(v int32) *GetEnterpriseDataChunkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnterpriseDataChunkResponse) SetBody(v *GetEnterpriseDataChunkResponseBody) *GetEnterpriseDataChunkResponse {
	s.Body = v
	return s
}

type GetEnterpriseDataPageImageRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DataId   *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
}

func (s GetEnterpriseDataPageImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataPageImageRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataPageImageRequest) SetAgentKey(v string) *GetEnterpriseDataPageImageRequest {
	s.AgentKey = &v
	return s
}

func (s *GetEnterpriseDataPageImageRequest) SetDataId(v string) *GetEnterpriseDataPageImageRequest {
	s.DataId = &v
	return s
}

type GetEnterpriseDataPageImageResponseBody struct {
	Data      []*GetEnterpriseDataPageImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                                       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEnterpriseDataPageImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataPageImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataPageImageResponseBody) SetData(v []*GetEnterpriseDataPageImageResponseBodyData) *GetEnterpriseDataPageImageResponseBody {
	s.Data = v
	return s
}

func (s *GetEnterpriseDataPageImageResponseBody) SetErrorCode(v string) *GetEnterpriseDataPageImageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetEnterpriseDataPageImageResponseBody) SetErrorMsg(v string) *GetEnterpriseDataPageImageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetEnterpriseDataPageImageResponseBody) SetRequestId(v string) *GetEnterpriseDataPageImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnterpriseDataPageImageResponseBody) SetSuccess(v bool) *GetEnterpriseDataPageImageResponseBody {
	s.Success = &v
	return s
}

type GetEnterpriseDataPageImageResponseBodyData struct {
	Height   *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	PageId   *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	Width    *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetEnterpriseDataPageImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataPageImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataPageImageResponseBodyData) SetHeight(v int32) *GetEnterpriseDataPageImageResponseBodyData {
	s.Height = &v
	return s
}

func (s *GetEnterpriseDataPageImageResponseBodyData) SetImageUrl(v string) *GetEnterpriseDataPageImageResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *GetEnterpriseDataPageImageResponseBodyData) SetPageId(v string) *GetEnterpriseDataPageImageResponseBodyData {
	s.PageId = &v
	return s
}

func (s *GetEnterpriseDataPageImageResponseBodyData) SetWidth(v int32) *GetEnterpriseDataPageImageResponseBodyData {
	s.Width = &v
	return s
}

type GetEnterpriseDataPageImageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEnterpriseDataPageImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnterpriseDataPageImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataPageImageResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataPageImageResponse) SetHeaders(v map[string]*string) *GetEnterpriseDataPageImageResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseDataPageImageResponse) SetStatusCode(v int32) *GetEnterpriseDataPageImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnterpriseDataPageImageResponse) SetBody(v *GetEnterpriseDataPageImageResponseBody) *GetEnterpriseDataPageImageResponse {
	s.Body = v
	return s
}

type GetEnterpriseDataParseResultRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DataId   *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
}

func (s GetEnterpriseDataParseResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataParseResultRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataParseResultRequest) SetAgentKey(v string) *GetEnterpriseDataParseResultRequest {
	s.AgentKey = &v
	return s
}

func (s *GetEnterpriseDataParseResultRequest) SetDataId(v string) *GetEnterpriseDataParseResultRequest {
	s.DataId = &v
	return s
}

type GetEnterpriseDataParseResultResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEnterpriseDataParseResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataParseResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataParseResultResponseBody) SetData(v string) *GetEnterpriseDataParseResultResponseBody {
	s.Data = &v
	return s
}

func (s *GetEnterpriseDataParseResultResponseBody) SetErrorCode(v string) *GetEnterpriseDataParseResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetEnterpriseDataParseResultResponseBody) SetErrorMsg(v string) *GetEnterpriseDataParseResultResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetEnterpriseDataParseResultResponseBody) SetRequestId(v string) *GetEnterpriseDataParseResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnterpriseDataParseResultResponseBody) SetSuccess(v bool) *GetEnterpriseDataParseResultResponseBody {
	s.Success = &v
	return s
}

type GetEnterpriseDataParseResultResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEnterpriseDataParseResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnterpriseDataParseResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnterpriseDataParseResultResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseDataParseResultResponse) SetHeaders(v map[string]*string) *GetEnterpriseDataParseResultResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseDataParseResultResponse) SetStatusCode(v int32) *GetEnterpriseDataParseResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnterpriseDataParseResultResponse) SetBody(v *GetEnterpriseDataParseResultResponseBody) *GetEnterpriseDataParseResultResponse {
	s.Body = v
	return s
}

type GetImportTaskResultRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetImportTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImportTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetImportTaskResultRequest) SetAgentKey(v string) *GetImportTaskResultRequest {
	s.AgentKey = &v
	return s
}

func (s *GetImportTaskResultRequest) SetTaskId(v string) *GetImportTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetImportTaskResultResponseBody struct {
	Data      *GetImportTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImportTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImportTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetImportTaskResultResponseBody) SetData(v *GetImportTaskResultResponseBodyData) *GetImportTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetImportTaskResultResponseBody) SetErrorCode(v string) *GetImportTaskResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetImportTaskResultResponseBody) SetErrorMsg(v string) *GetImportTaskResultResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetImportTaskResultResponseBody) SetRequestId(v string) *GetImportTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImportTaskResultResponseBody) SetSuccess(v bool) *GetImportTaskResultResponseBody {
	s.Success = &v
	return s
}

type GetImportTaskResultResponseBodyData struct {
	Details        []*GetImportTaskResultResponseBodyDataDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	TaskId         *string                                       `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus     *int32                                        `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusText *string                                       `json:"TaskStatusText,omitempty" xml:"TaskStatusText,omitempty"`
}

func (s GetImportTaskResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetImportTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImportTaskResultResponseBodyData) SetDetails(v []*GetImportTaskResultResponseBodyDataDetails) *GetImportTaskResultResponseBodyData {
	s.Details = v
	return s
}

func (s *GetImportTaskResultResponseBodyData) SetTaskId(v string) *GetImportTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetImportTaskResultResponseBodyData) SetTaskStatus(v int32) *GetImportTaskResultResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetImportTaskResultResponseBodyData) SetTaskStatusText(v string) *GetImportTaskResultResponseBodyData {
	s.TaskStatusText = &v
	return s
}

type GetImportTaskResultResponseBodyDataDetails struct {
	DataId   *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	DataName *string `json:"DataName,omitempty" xml:"DataName,omitempty"`
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Success  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImportTaskResultResponseBodyDataDetails) String() string {
	return tea.Prettify(s)
}

func (s GetImportTaskResultResponseBodyDataDetails) GoString() string {
	return s.String()
}

func (s *GetImportTaskResultResponseBodyDataDetails) SetDataId(v string) *GetImportTaskResultResponseBodyDataDetails {
	s.DataId = &v
	return s
}

func (s *GetImportTaskResultResponseBodyDataDetails) SetDataName(v string) *GetImportTaskResultResponseBodyDataDetails {
	s.DataName = &v
	return s
}

func (s *GetImportTaskResultResponseBodyDataDetails) SetErrorMsg(v string) *GetImportTaskResultResponseBodyDataDetails {
	s.ErrorMsg = &v
	return s
}

func (s *GetImportTaskResultResponseBodyDataDetails) SetSuccess(v bool) *GetImportTaskResultResponseBodyDataDetails {
	s.Success = &v
	return s
}

type GetImportTaskResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetImportTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImportTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImportTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetImportTaskResultResponse) SetHeaders(v map[string]*string) *GetImportTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetImportTaskResultResponse) SetStatusCode(v int32) *GetImportTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImportTaskResultResponse) SetBody(v *GetImportTaskResultResponseBody) *GetImportTaskResultResponse {
	s.Body = v
	return s
}

type ImportEnterpriseDocumentRequest struct {
	AgentKey     *string                                        `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DocumentList []*ImportEnterpriseDocumentRequestDocumentList `json:"DocumentList,omitempty" xml:"DocumentList,omitempty" type:"Repeated"`
	OwnerId      *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StoreId      *int64                                         `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	Tags         []*string                                      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ImportEnterpriseDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportEnterpriseDocumentRequest) GoString() string {
	return s.String()
}

func (s *ImportEnterpriseDocumentRequest) SetAgentKey(v string) *ImportEnterpriseDocumentRequest {
	s.AgentKey = &v
	return s
}

func (s *ImportEnterpriseDocumentRequest) SetDocumentList(v []*ImportEnterpriseDocumentRequestDocumentList) *ImportEnterpriseDocumentRequest {
	s.DocumentList = v
	return s
}

func (s *ImportEnterpriseDocumentRequest) SetOwnerId(v int64) *ImportEnterpriseDocumentRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportEnterpriseDocumentRequest) SetStoreId(v int64) *ImportEnterpriseDocumentRequest {
	s.StoreId = &v
	return s
}

func (s *ImportEnterpriseDocumentRequest) SetTags(v []*string) *ImportEnterpriseDocumentRequest {
	s.Tags = v
	return s
}

type ImportEnterpriseDocumentRequestDocumentList struct {
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	FileCanDownload *bool   `json:"FileCanDownload,omitempty" xml:"FileCanDownload,omitempty"`
	FileLink        *string `json:"FileLink,omitempty" xml:"FileLink,omitempty"`
	FileName        *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FilePreviewLink *string `json:"FilePreviewLink,omitempty" xml:"FilePreviewLink,omitempty"`
}

func (s ImportEnterpriseDocumentRequestDocumentList) String() string {
	return tea.Prettify(s)
}

func (s ImportEnterpriseDocumentRequestDocumentList) GoString() string {
	return s.String()
}

func (s *ImportEnterpriseDocumentRequestDocumentList) SetBizId(v string) *ImportEnterpriseDocumentRequestDocumentList {
	s.BizId = &v
	return s
}

func (s *ImportEnterpriseDocumentRequestDocumentList) SetFileCanDownload(v bool) *ImportEnterpriseDocumentRequestDocumentList {
	s.FileCanDownload = &v
	return s
}

func (s *ImportEnterpriseDocumentRequestDocumentList) SetFileLink(v string) *ImportEnterpriseDocumentRequestDocumentList {
	s.FileLink = &v
	return s
}

func (s *ImportEnterpriseDocumentRequestDocumentList) SetFileName(v string) *ImportEnterpriseDocumentRequestDocumentList {
	s.FileName = &v
	return s
}

func (s *ImportEnterpriseDocumentRequestDocumentList) SetFilePreviewLink(v string) *ImportEnterpriseDocumentRequestDocumentList {
	s.FilePreviewLink = &v
	return s
}

type ImportEnterpriseDocumentShrinkRequest struct {
	AgentKey           *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DocumentListShrink *string `json:"DocumentList,omitempty" xml:"DocumentList,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StoreId            *int64  `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	TagsShrink         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ImportEnterpriseDocumentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportEnterpriseDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportEnterpriseDocumentShrinkRequest) SetAgentKey(v string) *ImportEnterpriseDocumentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ImportEnterpriseDocumentShrinkRequest) SetDocumentListShrink(v string) *ImportEnterpriseDocumentShrinkRequest {
	s.DocumentListShrink = &v
	return s
}

func (s *ImportEnterpriseDocumentShrinkRequest) SetOwnerId(v int64) *ImportEnterpriseDocumentShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportEnterpriseDocumentShrinkRequest) SetStoreId(v int64) *ImportEnterpriseDocumentShrinkRequest {
	s.StoreId = &v
	return s
}

func (s *ImportEnterpriseDocumentShrinkRequest) SetTagsShrink(v string) *ImportEnterpriseDocumentShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ImportEnterpriseDocumentResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportEnterpriseDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportEnterpriseDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *ImportEnterpriseDocumentResponseBody) SetData(v string) *ImportEnterpriseDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *ImportEnterpriseDocumentResponseBody) SetErrorCode(v string) *ImportEnterpriseDocumentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ImportEnterpriseDocumentResponseBody) SetErrorMsg(v string) *ImportEnterpriseDocumentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ImportEnterpriseDocumentResponseBody) SetRequestId(v string) *ImportEnterpriseDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportEnterpriseDocumentResponseBody) SetSuccess(v bool) *ImportEnterpriseDocumentResponseBody {
	s.Success = &v
	return s
}

type ImportEnterpriseDocumentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImportEnterpriseDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportEnterpriseDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportEnterpriseDocumentResponse) GoString() string {
	return s.String()
}

func (s *ImportEnterpriseDocumentResponse) SetHeaders(v map[string]*string) *ImportEnterpriseDocumentResponse {
	s.Headers = v
	return s
}

func (s *ImportEnterpriseDocumentResponse) SetStatusCode(v int32) *ImportEnterpriseDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportEnterpriseDocumentResponse) SetBody(v *ImportEnterpriseDocumentResponseBody) *ImportEnterpriseDocumentResponse {
	s.Body = v
	return s
}

type ListFineTuneJobsRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	PageNo   *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFineTuneJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFineTuneJobsRequest) GoString() string {
	return s.String()
}

func (s *ListFineTuneJobsRequest) SetAgentKey(v string) *ListFineTuneJobsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListFineTuneJobsRequest) SetPageNo(v int32) *ListFineTuneJobsRequest {
	s.PageNo = &v
	return s
}

func (s *ListFineTuneJobsRequest) SetPageSize(v int32) *ListFineTuneJobsRequest {
	s.PageSize = &v
	return s
}

type ListFineTuneJobsResponseBody struct {
	Jobs      []*ListFineTuneJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	PageNo    *int32                              `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListFineTuneJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFineTuneJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFineTuneJobsResponseBody) SetJobs(v []*ListFineTuneJobsResponseBodyJobs) *ListFineTuneJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListFineTuneJobsResponseBody) SetPageNo(v int32) *ListFineTuneJobsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListFineTuneJobsResponseBody) SetPageSize(v int32) *ListFineTuneJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFineTuneJobsResponseBody) SetRequestId(v string) *ListFineTuneJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFineTuneJobsResponseBody) SetTotal(v int32) *ListFineTuneJobsResponseBody {
	s.Total = &v
	return s
}

type ListFineTuneJobsResponseBodyJobs struct {
	BaseModel       *string                                          `json:"BaseModel,omitempty" xml:"BaseModel,omitempty"`
	FineTunedModel  *string                                          `json:"FineTunedModel,omitempty" xml:"FineTunedModel,omitempty"`
	HyperParameters *ListFineTuneJobsResponseBodyJobsHyperParameters `json:"HyperParameters,omitempty" xml:"HyperParameters,omitempty" type:"Struct"`
	JobId           *string                                          `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message         *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	ModelName       *string                                          `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Status          *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	TrainingFiles   []*string                                        `json:"TrainingFiles,omitempty" xml:"TrainingFiles,omitempty" type:"Repeated"`
	TrainingType    *string                                          `json:"TrainingType,omitempty" xml:"TrainingType,omitempty"`
	ValidationFiles []*string                                        `json:"ValidationFiles,omitempty" xml:"ValidationFiles,omitempty" type:"Repeated"`
}

func (s ListFineTuneJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListFineTuneJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListFineTuneJobsResponseBodyJobs) SetBaseModel(v string) *ListFineTuneJobsResponseBodyJobs {
	s.BaseModel = &v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobs) SetFineTunedModel(v string) *ListFineTuneJobsResponseBodyJobs {
	s.FineTunedModel = &v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobs) SetHyperParameters(v *ListFineTuneJobsResponseBodyJobsHyperParameters) *ListFineTuneJobsResponseBodyJobs {
	s.HyperParameters = v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobs) SetJobId(v string) *ListFineTuneJobsResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobs) SetMessage(v string) *ListFineTuneJobsResponseBodyJobs {
	s.Message = &v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobs) SetModelName(v string) *ListFineTuneJobsResponseBodyJobs {
	s.ModelName = &v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobs) SetStatus(v string) *ListFineTuneJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobs) SetTrainingFiles(v []*string) *ListFineTuneJobsResponseBodyJobs {
	s.TrainingFiles = v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobs) SetTrainingType(v string) *ListFineTuneJobsResponseBodyJobs {
	s.TrainingType = &v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobs) SetValidationFiles(v []*string) *ListFineTuneJobsResponseBodyJobs {
	s.ValidationFiles = v
	return s
}

type ListFineTuneJobsResponseBodyJobsHyperParameters struct {
	BatchSize        *int32   `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
	Epochs           *int32   `json:"Epochs,omitempty" xml:"Epochs,omitempty"`
	LearningRate     *string  `json:"LearningRate,omitempty" xml:"LearningRate,omitempty"`
	PromptLossWeight *float64 `json:"PromptLossWeight,omitempty" xml:"PromptLossWeight,omitempty"`
}

func (s ListFineTuneJobsResponseBodyJobsHyperParameters) String() string {
	return tea.Prettify(s)
}

func (s ListFineTuneJobsResponseBodyJobsHyperParameters) GoString() string {
	return s.String()
}

func (s *ListFineTuneJobsResponseBodyJobsHyperParameters) SetBatchSize(v int32) *ListFineTuneJobsResponseBodyJobsHyperParameters {
	s.BatchSize = &v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobsHyperParameters) SetEpochs(v int32) *ListFineTuneJobsResponseBodyJobsHyperParameters {
	s.Epochs = &v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobsHyperParameters) SetLearningRate(v string) *ListFineTuneJobsResponseBodyJobsHyperParameters {
	s.LearningRate = &v
	return s
}

func (s *ListFineTuneJobsResponseBodyJobsHyperParameters) SetPromptLossWeight(v float64) *ListFineTuneJobsResponseBodyJobsHyperParameters {
	s.PromptLossWeight = &v
	return s
}

type ListFineTuneJobsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFineTuneJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFineTuneJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFineTuneJobsResponse) GoString() string {
	return s.String()
}

func (s *ListFineTuneJobsResponse) SetHeaders(v map[string]*string) *ListFineTuneJobsResponse {
	s.Headers = v
	return s
}

func (s *ListFineTuneJobsResponse) SetStatusCode(v int32) *ListFineTuneJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFineTuneJobsResponse) SetBody(v *ListFineTuneJobsResponseBody) *ListFineTuneJobsResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	PageNo   *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetAgentKey(v string) *ListServicesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListServicesRequest) SetPageNo(v int32) *ListServicesRequest {
	s.PageNo = &v
	return s
}

func (s *ListServicesRequest) SetPageSize(v int32) *ListServicesRequest {
	s.PageSize = &v
	return s
}

type ListServicesResponseBody struct {
	ModelServices []*ListServicesResponseBodyModelServices `json:"ModelServices,omitempty" xml:"ModelServices,omitempty" type:"Repeated"`
	PageNo        *int32                                   `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize      *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total         *int32                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetModelServices(v []*ListServicesResponseBodyModelServices) *ListServicesResponseBody {
	s.ModelServices = v
	return s
}

func (s *ListServicesResponseBody) SetPageNo(v int32) *ListServicesResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListServicesResponseBody) SetPageSize(v int32) *ListServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetTotal(v int32) *ListServicesResponseBody {
	s.Total = &v
	return s
}

type ListServicesResponseBodyModelServices struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListServicesResponseBodyModelServices) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyModelServices) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyModelServices) SetAppId(v string) *ListServicesResponseBodyModelServices {
	s.AppId = &v
	return s
}

func (s *ListServicesResponseBodyModelServices) SetModelServiceId(v string) *ListServicesResponseBodyModelServices {
	s.ModelServiceId = &v
	return s
}

func (s *ListServicesResponseBodyModelServices) SetStatus(v string) *ListServicesResponseBodyModelServices {
	s.Status = &v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type QueryEnterpriseDataListRequest struct {
	AgentKey  *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DataName  *string   `json:"DataName,omitempty" xml:"DataName,omitempty"`
	PageNo    *int32    `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StoreType *string   `json:"StoreType,omitempty" xml:"StoreType,omitempty"`
	Tags      []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s QueryEnterpriseDataListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseDataListRequest) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseDataListRequest) SetAgentKey(v string) *QueryEnterpriseDataListRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryEnterpriseDataListRequest) SetDataName(v string) *QueryEnterpriseDataListRequest {
	s.DataName = &v
	return s
}

func (s *QueryEnterpriseDataListRequest) SetPageNo(v int32) *QueryEnterpriseDataListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryEnterpriseDataListRequest) SetPageSize(v int32) *QueryEnterpriseDataListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryEnterpriseDataListRequest) SetStoreType(v string) *QueryEnterpriseDataListRequest {
	s.StoreType = &v
	return s
}

func (s *QueryEnterpriseDataListRequest) SetTags(v []*string) *QueryEnterpriseDataListRequest {
	s.Tags = v
	return s
}

type QueryEnterpriseDataListShrinkRequest struct {
	AgentKey   *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DataName   *string `json:"DataName,omitempty" xml:"DataName,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StoreType  *string `json:"StoreType,omitempty" xml:"StoreType,omitempty"`
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s QueryEnterpriseDataListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseDataListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseDataListShrinkRequest) SetAgentKey(v string) *QueryEnterpriseDataListShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryEnterpriseDataListShrinkRequest) SetDataName(v string) *QueryEnterpriseDataListShrinkRequest {
	s.DataName = &v
	return s
}

func (s *QueryEnterpriseDataListShrinkRequest) SetPageNo(v int32) *QueryEnterpriseDataListShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *QueryEnterpriseDataListShrinkRequest) SetPageSize(v int32) *QueryEnterpriseDataListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryEnterpriseDataListShrinkRequest) SetStoreType(v string) *QueryEnterpriseDataListShrinkRequest {
	s.StoreType = &v
	return s
}

func (s *QueryEnterpriseDataListShrinkRequest) SetTagsShrink(v string) *QueryEnterpriseDataListShrinkRequest {
	s.TagsShrink = &v
	return s
}

type QueryEnterpriseDataListResponseBody struct {
	Data      *QueryEnterpriseDataListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                                  `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEnterpriseDataListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseDataListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseDataListResponseBody) SetData(v *QueryEnterpriseDataListResponseBodyData) *QueryEnterpriseDataListResponseBody {
	s.Data = v
	return s
}

func (s *QueryEnterpriseDataListResponseBody) SetErrorCode(v string) *QueryEnterpriseDataListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBody) SetErrorMsg(v string) *QueryEnterpriseDataListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBody) SetRequestId(v string) *QueryEnterpriseDataListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBody) SetSuccess(v bool) *QueryEnterpriseDataListResponseBody {
	s.Success = &v
	return s
}

type QueryEnterpriseDataListResponseBodyData struct {
	List     []*QueryEnterpriseDataListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNo   *int32                                         `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryEnterpriseDataListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseDataListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseDataListResponseBodyData) SetList(v []*QueryEnterpriseDataListResponseBodyDataList) *QueryEnterpriseDataListResponseBodyData {
	s.List = v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyData) SetPageNo(v int32) *QueryEnterpriseDataListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyData) SetPageSize(v int32) *QueryEnterpriseDataListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyData) SetTotal(v int64) *QueryEnterpriseDataListResponseBodyData {
	s.Total = &v
	return s
}

type QueryEnterpriseDataListResponseBodyDataList struct {
	DataId         *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	DataName       *string `json:"DataName,omitempty" xml:"DataName,omitempty"`
	DataSize       *string `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	DataStatus     *string `json:"DataStatus,omitempty" xml:"DataStatus,omitempty"`
	DataStatusCode *int32  `json:"DataStatusCode,omitempty" xml:"DataStatusCode,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DataTypeCode   *int32  `json:"DataTypeCode,omitempty" xml:"DataTypeCode,omitempty"`
	DownloadUrl    *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	StatusDetail   *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	StoreType      *string `json:"StoreType,omitempty" xml:"StoreType,omitempty"`
	Tags           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UploadTime     *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s QueryEnterpriseDataListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseDataListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetDataId(v string) *QueryEnterpriseDataListResponseBodyDataList {
	s.DataId = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetDataName(v string) *QueryEnterpriseDataListResponseBodyDataList {
	s.DataName = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetDataSize(v string) *QueryEnterpriseDataListResponseBodyDataList {
	s.DataSize = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetDataStatus(v string) *QueryEnterpriseDataListResponseBodyDataList {
	s.DataStatus = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetDataStatusCode(v int32) *QueryEnterpriseDataListResponseBodyDataList {
	s.DataStatusCode = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetDataType(v string) *QueryEnterpriseDataListResponseBodyDataList {
	s.DataType = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetDataTypeCode(v int32) *QueryEnterpriseDataListResponseBodyDataList {
	s.DataTypeCode = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetDownloadUrl(v string) *QueryEnterpriseDataListResponseBodyDataList {
	s.DownloadUrl = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetStatusDetail(v string) *QueryEnterpriseDataListResponseBodyDataList {
	s.StatusDetail = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetStoreType(v string) *QueryEnterpriseDataListResponseBodyDataList {
	s.StoreType = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetTags(v string) *QueryEnterpriseDataListResponseBodyDataList {
	s.Tags = &v
	return s
}

func (s *QueryEnterpriseDataListResponseBodyDataList) SetUploadTime(v string) *QueryEnterpriseDataListResponseBodyDataList {
	s.UploadTime = &v
	return s
}

type QueryEnterpriseDataListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryEnterpriseDataListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEnterpriseDataListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEnterpriseDataListResponse) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseDataListResponse) SetHeaders(v map[string]*string) *QueryEnterpriseDataListResponse {
	s.Headers = v
	return s
}

func (s *QueryEnterpriseDataListResponse) SetStatusCode(v int32) *QueryEnterpriseDataListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEnterpriseDataListResponse) SetBody(v *QueryEnterpriseDataListResponseBody) *QueryEnterpriseDataListResponse {
	s.Body = v
	return s
}

type SearchEnterpriseDataRequest struct {
	AgentKey   *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DataIdList []*string `json:"DataIdList,omitempty" xml:"DataIdList,omitempty" type:"Repeated"`
	EnableRank *bool     `json:"EnableRank,omitempty" xml:"EnableRank,omitempty"`
	OwnerId    *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Query      *string   `json:"Query,omitempty" xml:"Query,omitempty"`
	StoreId    *int64    `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	TagIdList  []*int64  `json:"TagIdList,omitempty" xml:"TagIdList,omitempty" type:"Repeated"`
}

func (s SearchEnterpriseDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchEnterpriseDataRequest) GoString() string {
	return s.String()
}

func (s *SearchEnterpriseDataRequest) SetAgentKey(v string) *SearchEnterpriseDataRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchEnterpriseDataRequest) SetDataIdList(v []*string) *SearchEnterpriseDataRequest {
	s.DataIdList = v
	return s
}

func (s *SearchEnterpriseDataRequest) SetEnableRank(v bool) *SearchEnterpriseDataRequest {
	s.EnableRank = &v
	return s
}

func (s *SearchEnterpriseDataRequest) SetOwnerId(v int64) *SearchEnterpriseDataRequest {
	s.OwnerId = &v
	return s
}

func (s *SearchEnterpriseDataRequest) SetQuery(v string) *SearchEnterpriseDataRequest {
	s.Query = &v
	return s
}

func (s *SearchEnterpriseDataRequest) SetStoreId(v int64) *SearchEnterpriseDataRequest {
	s.StoreId = &v
	return s
}

func (s *SearchEnterpriseDataRequest) SetTagIdList(v []*int64) *SearchEnterpriseDataRequest {
	s.TagIdList = v
	return s
}

type SearchEnterpriseDataShrinkRequest struct {
	AgentKey         *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DataIdListShrink *string `json:"DataIdList,omitempty" xml:"DataIdList,omitempty"`
	EnableRank       *bool   `json:"EnableRank,omitempty" xml:"EnableRank,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Query            *string `json:"Query,omitempty" xml:"Query,omitempty"`
	StoreId          *int64  `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	TagIdListShrink  *string `json:"TagIdList,omitempty" xml:"TagIdList,omitempty"`
}

func (s SearchEnterpriseDataShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchEnterpriseDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchEnterpriseDataShrinkRequest) SetAgentKey(v string) *SearchEnterpriseDataShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchEnterpriseDataShrinkRequest) SetDataIdListShrink(v string) *SearchEnterpriseDataShrinkRequest {
	s.DataIdListShrink = &v
	return s
}

func (s *SearchEnterpriseDataShrinkRequest) SetEnableRank(v bool) *SearchEnterpriseDataShrinkRequest {
	s.EnableRank = &v
	return s
}

func (s *SearchEnterpriseDataShrinkRequest) SetOwnerId(v int64) *SearchEnterpriseDataShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SearchEnterpriseDataShrinkRequest) SetQuery(v string) *SearchEnterpriseDataShrinkRequest {
	s.Query = &v
	return s
}

func (s *SearchEnterpriseDataShrinkRequest) SetStoreId(v int64) *SearchEnterpriseDataShrinkRequest {
	s.StoreId = &v
	return s
}

func (s *SearchEnterpriseDataShrinkRequest) SetTagIdListShrink(v string) *SearchEnterpriseDataShrinkRequest {
	s.TagIdListShrink = &v
	return s
}

type SearchEnterpriseDataResponseBody struct {
	Data      []*SearchEnterpriseDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                                 `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchEnterpriseDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchEnterpriseDataResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEnterpriseDataResponseBody) SetData(v []*SearchEnterpriseDataResponseBodyData) *SearchEnterpriseDataResponseBody {
	s.Data = v
	return s
}

func (s *SearchEnterpriseDataResponseBody) SetErrorCode(v string) *SearchEnterpriseDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchEnterpriseDataResponseBody) SetErrorMsg(v string) *SearchEnterpriseDataResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SearchEnterpriseDataResponseBody) SetRequestId(v string) *SearchEnterpriseDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchEnterpriseDataResponseBody) SetSuccess(v bool) *SearchEnterpriseDataResponseBody {
	s.Success = &v
	return s
}

type SearchEnterpriseDataResponseBodyData struct {
	DataId    *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	DataName  *string `json:"DataName,omitempty" xml:"DataName,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	TitlePath *string `json:"TitlePath,omitempty" xml:"TitlePath,omitempty"`
}

func (s SearchEnterpriseDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchEnterpriseDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchEnterpriseDataResponseBodyData) SetDataId(v string) *SearchEnterpriseDataResponseBodyData {
	s.DataId = &v
	return s
}

func (s *SearchEnterpriseDataResponseBodyData) SetDataName(v string) *SearchEnterpriseDataResponseBodyData {
	s.DataName = &v
	return s
}

func (s *SearchEnterpriseDataResponseBodyData) SetScore(v string) *SearchEnterpriseDataResponseBodyData {
	s.Score = &v
	return s
}

func (s *SearchEnterpriseDataResponseBodyData) SetSource(v string) *SearchEnterpriseDataResponseBodyData {
	s.Source = &v
	return s
}

func (s *SearchEnterpriseDataResponseBodyData) SetText(v string) *SearchEnterpriseDataResponseBodyData {
	s.Text = &v
	return s
}

func (s *SearchEnterpriseDataResponseBodyData) SetTitle(v string) *SearchEnterpriseDataResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchEnterpriseDataResponseBodyData) SetTitlePath(v string) *SearchEnterpriseDataResponseBodyData {
	s.TitlePath = &v
	return s
}

type SearchEnterpriseDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchEnterpriseDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchEnterpriseDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchEnterpriseDataResponse) GoString() string {
	return s.String()
}

func (s *SearchEnterpriseDataResponse) SetHeaders(v map[string]*string) *SearchEnterpriseDataResponse {
	s.Headers = v
	return s
}

func (s *SearchEnterpriseDataResponse) SetStatusCode(v int32) *SearchEnterpriseDataResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchEnterpriseDataResponse) SetBody(v *SearchEnterpriseDataResponseBody) *SearchEnterpriseDataResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("bailian"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CancelFineTuneJobWithOptions(request *CancelFineTuneJobRequest, runtime *util.RuntimeOptions) (_result *CancelFineTuneJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelFineTuneJob"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelFineTuneJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelFineTuneJob(request *CancelFineTuneJobRequest) (_result *CancelFineTuneJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelFineTuneJobResponse{}
	_body, _err := client.CancelFineTuneJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFineTuneJobWithOptions(tmpReq *CreateFineTuneJobRequest, runtime *util.RuntimeOptions) (_result *CreateFineTuneJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateFineTuneJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HyperParameters)) {
		request.HyperParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HyperParameters, tea.String("HyperParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TrainingFiles)) {
		request.TrainingFilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TrainingFiles, tea.String("TrainingFiles"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ValidationFiles)) {
		request.ValidationFilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ValidationFiles, tea.String("ValidationFiles"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaseModel)) {
		body["BaseModel"] = request.BaseModel
	}

	if !tea.BoolValue(util.IsUnset(request.HyperParametersShrink)) {
		body["HyperParameters"] = request.HyperParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		body["ModelName"] = request.ModelName
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingFilesShrink)) {
		body["TrainingFiles"] = request.TrainingFilesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingType)) {
		body["TrainingType"] = request.TrainingType
	}

	if !tea.BoolValue(util.IsUnset(request.ValidationFilesShrink)) {
		body["ValidationFiles"] = request.ValidationFilesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFineTuneJob"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFineTuneJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFineTuneJob(request *CreateFineTuneJobRequest) (_result *CreateFineTuneJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFineTuneJobResponse{}
	_body, _err := client.CreateFineTuneJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Model)) {
		body["Model"] = request.Model
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateService"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTextEmbeddingsWithOptions(tmpReq *CreateTextEmbeddingsRequest, runtime *util.RuntimeOptions) (_result *CreateTextEmbeddingsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTextEmbeddingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Input)) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, tea.String("Input"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.InputShrink)) {
		query["Input"] = request.InputShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TextType)) {
		query["TextType"] = request.TextType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTextEmbeddings"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTextEmbeddingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTextEmbeddings(request *CreateTextEmbeddingsRequest) (_result *CreateTextEmbeddingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTextEmbeddingsResponse{}
	_body, _err := client.CreateTextEmbeddingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTokenWithOptions(request *CreateTokenRequest, runtime *util.RuntimeOptions) (_result *CreateTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateToken"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateToken(request *CreateTokenRequest) (_result *CreateTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTokenResponse{}
	_body, _err := client.CreateTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEnterpriseDataWithOptions(request *DeleteEnterpriseDataRequest, runtime *util.RuntimeOptions) (_result *DeleteEnterpriseDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		query["DataId"] = request.DataId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEnterpriseData"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEnterpriseDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEnterpriseData(request *DeleteEnterpriseDataRequest) (_result *DeleteEnterpriseDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEnterpriseDataResponse{}
	_body, _err := client.DeleteEnterpriseDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFineTuneJobWithOptions(request *DeleteFineTuneJobRequest, runtime *util.RuntimeOptions) (_result *DeleteFineTuneJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFineTuneJob"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFineTuneJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFineTuneJob(request *DeleteFineTuneJobRequest) (_result *DeleteFineTuneJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFineTuneJobResponse{}
	_body, _err := client.DeleteFineTuneJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceWithOptions(request *DeleteServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelServiceId)) {
		body["ModelServiceId"] = request.ModelServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteService"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteService(request *DeleteServiceRequest) (_result *DeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFineTuneJobWithOptions(request *DescribeFineTuneJobRequest, runtime *util.RuntimeOptions) (_result *DescribeFineTuneJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFineTuneJob"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFineTuneJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFineTuneJob(request *DescribeFineTuneJobRequest) (_result *DescribeFineTuneJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFineTuneJobResponse{}
	_body, _err := client.DescribeFineTuneJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceWithOptions(request *DescribeServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelServiceId)) {
		body["ModelServiceId"] = request.ModelServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeService"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeService(request *DescribeServiceRequest) (_result *DescribeServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceResponse{}
	_body, _err := client.DescribeServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnterpriseDataByDataIdWithOptions(request *GetEnterpriseDataByDataIdRequest, runtime *util.RuntimeOptions) (_result *GetEnterpriseDataByDataIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		query["DataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnterpriseDataByDataId"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnterpriseDataByDataIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnterpriseDataByDataId(request *GetEnterpriseDataByDataIdRequest) (_result *GetEnterpriseDataByDataIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEnterpriseDataByDataIdResponse{}
	_body, _err := client.GetEnterpriseDataByDataIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnterpriseDataChunkWithOptions(request *GetEnterpriseDataChunkRequest, runtime *util.RuntimeOptions) (_result *GetEnterpriseDataChunkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		query["DataId"] = request.DataId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnterpriseDataChunk"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnterpriseDataChunkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnterpriseDataChunk(request *GetEnterpriseDataChunkRequest) (_result *GetEnterpriseDataChunkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEnterpriseDataChunkResponse{}
	_body, _err := client.GetEnterpriseDataChunkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnterpriseDataPageImageWithOptions(request *GetEnterpriseDataPageImageRequest, runtime *util.RuntimeOptions) (_result *GetEnterpriseDataPageImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		query["DataId"] = request.DataId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnterpriseDataPageImage"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnterpriseDataPageImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnterpriseDataPageImage(request *GetEnterpriseDataPageImageRequest) (_result *GetEnterpriseDataPageImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEnterpriseDataPageImageResponse{}
	_body, _err := client.GetEnterpriseDataPageImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnterpriseDataParseResultWithOptions(request *GetEnterpriseDataParseResultRequest, runtime *util.RuntimeOptions) (_result *GetEnterpriseDataParseResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		query["DataId"] = request.DataId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnterpriseDataParseResult"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnterpriseDataParseResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnterpriseDataParseResult(request *GetEnterpriseDataParseResultRequest) (_result *GetEnterpriseDataParseResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEnterpriseDataParseResultResponse{}
	_body, _err := client.GetEnterpriseDataParseResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImportTaskResultWithOptions(request *GetImportTaskResultRequest, runtime *util.RuntimeOptions) (_result *GetImportTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImportTaskResult"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImportTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImportTaskResult(request *GetImportTaskResultRequest) (_result *GetImportTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImportTaskResultResponse{}
	_body, _err := client.GetImportTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportEnterpriseDocumentWithOptions(tmpReq *ImportEnterpriseDocumentRequest, runtime *util.RuntimeOptions) (_result *ImportEnterpriseDocumentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportEnterpriseDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DocumentList)) {
		request.DocumentListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentList, tea.String("DocumentList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentListShrink)) {
		query["DocumentList"] = request.DocumentListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		query["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportEnterpriseDocument"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportEnterpriseDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportEnterpriseDocument(request *ImportEnterpriseDocumentRequest) (_result *ImportEnterpriseDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportEnterpriseDocumentResponse{}
	_body, _err := client.ImportEnterpriseDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFineTuneJobsWithOptions(request *ListFineTuneJobsRequest, runtime *util.RuntimeOptions) (_result *ListFineTuneJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFineTuneJobs"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFineTuneJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFineTuneJobs(request *ListFineTuneJobsRequest) (_result *ListFineTuneJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFineTuneJobsResponse{}
	_body, _err := client.ListFineTuneJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServicesWithOptions(request *ListServicesRequest, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		body["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEnterpriseDataListWithOptions(tmpReq *QueryEnterpriseDataListRequest, runtime *util.RuntimeOptions) (_result *QueryEnterpriseDataListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryEnterpriseDataListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataName)) {
		query["DataName"] = request.DataName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StoreType)) {
		query["StoreType"] = request.StoreType
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEnterpriseDataList"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEnterpriseDataListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEnterpriseDataList(request *QueryEnterpriseDataListRequest) (_result *QueryEnterpriseDataListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEnterpriseDataListResponse{}
	_body, _err := client.QueryEnterpriseDataListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchEnterpriseDataWithOptions(tmpReq *SearchEnterpriseDataRequest, runtime *util.RuntimeOptions) (_result *SearchEnterpriseDataResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchEnterpriseDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DataIdList)) {
		request.DataIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataIdList, tea.String("DataIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagIdList)) {
		request.TagIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIdList, tea.String("TagIdList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DataIdListShrink)) {
		query["DataIdList"] = request.DataIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EnableRank)) {
		query["EnableRank"] = request.EnableRank
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		query["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdListShrink)) {
		query["TagIdList"] = request.TagIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchEnterpriseData"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchEnterpriseDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchEnterpriseData(request *SearchEnterpriseDataRequest) (_result *SearchEnterpriseDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchEnterpriseDataResponse{}
	_body, _err := client.SearchEnterpriseDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
