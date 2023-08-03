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
