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

type CreateDocumentTagRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateDocumentTagRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDocumentTagRequest) GoString() string {
	return s.String()
}

func (s *CreateDocumentTagRequest) SetAgentKey(v string) *CreateDocumentTagRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDocumentTagRequest) SetName(v string) *CreateDocumentTagRequest {
	s.Name = &v
	return s
}

type CreateDocumentTagResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagId     *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s CreateDocumentTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDocumentTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocumentTagResponseBody) SetRequestId(v string) *CreateDocumentTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDocumentTagResponseBody) SetTagId(v string) *CreateDocumentTagResponseBody {
	s.TagId = &v
	return s
}

type CreateDocumentTagResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDocumentTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDocumentTagResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDocumentTagResponse) GoString() string {
	return s.String()
}

func (s *CreateDocumentTagResponse) SetHeaders(v map[string]*string) *CreateDocumentTagResponse {
	s.Headers = v
	return s
}

func (s *CreateDocumentTagResponse) SetStatusCode(v int32) *CreateDocumentTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDocumentTagResponse) SetBody(v *CreateDocumentTagResponseBody) *CreateDocumentTagResponse {
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

type DeleteDocRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DocId    *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
}

func (s DeleteDocRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocRequest) SetAgentKey(v string) *DeleteDocRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteDocRequest) SetDocId(v string) *DeleteDocRequest {
	s.DocId = &v
	return s
}

type DeleteDocResponseBody struct {
	DocId     *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocResponseBody) SetDocId(v string) *DeleteDocResponseBody {
	s.DocId = &v
	return s
}

func (s *DeleteDocResponseBody) SetRequestId(v string) *DeleteDocResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDocResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDocResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDocResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocResponse) GoString() string {
	return s.String()
}

func (s *DeleteDocResponse) SetHeaders(v map[string]*string) *DeleteDocResponse {
	s.Headers = v
	return s
}

func (s *DeleteDocResponse) SetStatusCode(v int32) *DeleteDocResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDocResponse) SetBody(v *DeleteDocResponseBody) *DeleteDocResponse {
	s.Body = v
	return s
}

type DeleteDocumentTagRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	TagId    *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DeleteDocumentTagRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocumentTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentTagRequest) SetAgentKey(v string) *DeleteDocumentTagRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteDocumentTagRequest) SetTagId(v string) *DeleteDocumentTagRequest {
	s.TagId = &v
	return s
}

type DeleteDocumentTagResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagId     *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DeleteDocumentTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocumentTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocumentTagResponseBody) SetRequestId(v string) *DeleteDocumentTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentTagResponseBody) SetTagId(v string) *DeleteDocumentTagResponseBody {
	s.TagId = &v
	return s
}

type DeleteDocumentTagResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDocumentTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDocumentTagResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocumentTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteDocumentTagResponse) SetHeaders(v map[string]*string) *DeleteDocumentTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteDocumentTagResponse) SetStatusCode(v int32) *DeleteDocumentTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDocumentTagResponse) SetBody(v *DeleteDocumentTagResponseBody) *DeleteDocumentTagResponse {
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

type DescribeDocRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DocId    *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
}

func (s DescribeDocRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocRequest) SetAgentKey(v string) *DescribeDocRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeDocRequest) SetDocId(v string) *DescribeDocRequest {
	s.DocId = &v
	return s
}

type DescribeDocResponseBody struct {
	DocId      *string                        `json:"DocId,omitempty" xml:"DocId,omitempty"`
	FailReason *string                        `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	Name       *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId    *string                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size       *string                        `json:"Size,omitempty" xml:"Size,omitempty"`
	Status     *string                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags       []*DescribeDocResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Type       *string                        `json:"Type,omitempty" xml:"Type,omitempty"`
	UploadTime *string                        `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s DescribeDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBody) SetDocId(v string) *DescribeDocResponseBody {
	s.DocId = &v
	return s
}

func (s *DescribeDocResponseBody) SetFailReason(v string) *DescribeDocResponseBody {
	s.FailReason = &v
	return s
}

func (s *DescribeDocResponseBody) SetName(v string) *DescribeDocResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeDocResponseBody) SetOwnerId(v string) *DescribeDocResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeDocResponseBody) SetRequestId(v string) *DescribeDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDocResponseBody) SetSize(v string) *DescribeDocResponseBody {
	s.Size = &v
	return s
}

func (s *DescribeDocResponseBody) SetStatus(v string) *DescribeDocResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDocResponseBody) SetTags(v []*DescribeDocResponseBodyTags) *DescribeDocResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeDocResponseBody) SetType(v string) *DescribeDocResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeDocResponseBody) SetUploadTime(v string) *DescribeDocResponseBody {
	s.UploadTime = &v
	return s
}

type DescribeDocResponseBodyTags struct {
	TagId   *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeDocResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeDocResponseBodyTags) SetTagId(v string) *DescribeDocResponseBodyTags {
	s.TagId = &v
	return s
}

func (s *DescribeDocResponseBodyTags) SetTagName(v string) *DescribeDocResponseBodyTags {
	s.TagName = &v
	return s
}

type DescribeDocResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDocResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDocResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocResponse) GoString() string {
	return s.String()
}

func (s *DescribeDocResponse) SetHeaders(v map[string]*string) *DescribeDocResponse {
	s.Headers = v
	return s
}

func (s *DescribeDocResponse) SetStatusCode(v int32) *DescribeDocResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDocResponse) SetBody(v *DescribeDocResponseBody) *DescribeDocResponse {
	s.Body = v
	return s
}

type DescribeDocumentImportJobRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeDocumentImportJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocumentImportJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocumentImportJobRequest) SetAgentKey(v string) *DescribeDocumentImportJobRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeDocumentImportJobRequest) SetJobId(v string) *DescribeDocumentImportJobRequest {
	s.JobId = &v
	return s
}

type DescribeDocumentImportJobResponseBody struct {
	Docs      []*DescribeDocumentImportJobResponseBodyDocs `json:"Docs,omitempty" xml:"Docs,omitempty" type:"Repeated"`
	JobId     *string                                      `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDocumentImportJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocumentImportJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocumentImportJobResponseBody) SetDocs(v []*DescribeDocumentImportJobResponseBodyDocs) *DescribeDocumentImportJobResponseBody {
	s.Docs = v
	return s
}

func (s *DescribeDocumentImportJobResponseBody) SetJobId(v string) *DescribeDocumentImportJobResponseBody {
	s.JobId = &v
	return s
}

func (s *DescribeDocumentImportJobResponseBody) SetRequestId(v string) *DescribeDocumentImportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDocumentImportJobResponseBody) SetStatus(v string) *DescribeDocumentImportJobResponseBody {
	s.Status = &v
	return s
}

type DescribeDocumentImportJobResponseBodyDocs struct {
	DocId      *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDocumentImportJobResponseBodyDocs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocumentImportJobResponseBodyDocs) GoString() string {
	return s.String()
}

func (s *DescribeDocumentImportJobResponseBodyDocs) SetDocId(v string) *DescribeDocumentImportJobResponseBodyDocs {
	s.DocId = &v
	return s
}

func (s *DescribeDocumentImportJobResponseBodyDocs) SetFailReason(v string) *DescribeDocumentImportJobResponseBodyDocs {
	s.FailReason = &v
	return s
}

func (s *DescribeDocumentImportJobResponseBodyDocs) SetName(v string) *DescribeDocumentImportJobResponseBodyDocs {
	s.Name = &v
	return s
}

func (s *DescribeDocumentImportJobResponseBodyDocs) SetStatus(v string) *DescribeDocumentImportJobResponseBodyDocs {
	s.Status = &v
	return s
}

type DescribeDocumentImportJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDocumentImportJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDocumentImportJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDocumentImportJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeDocumentImportJobResponse) SetHeaders(v map[string]*string) *DescribeDocumentImportJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeDocumentImportJobResponse) SetStatusCode(v int32) *DescribeDocumentImportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDocumentImportJobResponse) SetBody(v *DescribeDocumentImportJobResponseBody) *DescribeDocumentImportJobResponse {
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

type GetFileStoreUploadPolicyRequest struct {
	AgentKey    *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileStoreId *int64  `json:"FileStoreId,omitempty" xml:"FileStoreId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetFileStoreUploadPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileStoreUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetFileStoreUploadPolicyRequest) SetAgentKey(v string) *GetFileStoreUploadPolicyRequest {
	s.AgentKey = &v
	return s
}

func (s *GetFileStoreUploadPolicyRequest) SetFileName(v string) *GetFileStoreUploadPolicyRequest {
	s.FileName = &v
	return s
}

func (s *GetFileStoreUploadPolicyRequest) SetFileStoreId(v int64) *GetFileStoreUploadPolicyRequest {
	s.FileStoreId = &v
	return s
}

func (s *GetFileStoreUploadPolicyRequest) SetUserId(v string) *GetFileStoreUploadPolicyRequest {
	s.UserId = &v
	return s
}

type GetFileStoreUploadPolicyResponseBody struct {
	Data      *GetFileStoreUploadPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                                   `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFileStoreUploadPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFileStoreUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileStoreUploadPolicyResponseBody) SetData(v *GetFileStoreUploadPolicyResponseBodyData) *GetFileStoreUploadPolicyResponseBody {
	s.Data = v
	return s
}

func (s *GetFileStoreUploadPolicyResponseBody) SetErrorCode(v string) *GetFileStoreUploadPolicyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetFileStoreUploadPolicyResponseBody) SetErrorMsg(v string) *GetFileStoreUploadPolicyResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetFileStoreUploadPolicyResponseBody) SetRequestId(v string) *GetFileStoreUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileStoreUploadPolicyResponseBody) SetSuccess(v bool) *GetFileStoreUploadPolicyResponseBody {
	s.Success = &v
	return s
}

type GetFileStoreUploadPolicyResponseBodyData struct {
	AccessId      *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Dir           *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire        *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host          *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Policy        *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Signature     *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetFileStoreUploadPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFileStoreUploadPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileStoreUploadPolicyResponseBodyData) SetAccessId(v string) *GetFileStoreUploadPolicyResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *GetFileStoreUploadPolicyResponseBodyData) SetDir(v string) *GetFileStoreUploadPolicyResponseBodyData {
	s.Dir = &v
	return s
}

func (s *GetFileStoreUploadPolicyResponseBodyData) SetExpire(v string) *GetFileStoreUploadPolicyResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetFileStoreUploadPolicyResponseBodyData) SetHost(v string) *GetFileStoreUploadPolicyResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetFileStoreUploadPolicyResponseBodyData) SetKey(v string) *GetFileStoreUploadPolicyResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetFileStoreUploadPolicyResponseBodyData) SetPolicy(v string) *GetFileStoreUploadPolicyResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetFileStoreUploadPolicyResponseBodyData) SetSecurityToken(v string) *GetFileStoreUploadPolicyResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetFileStoreUploadPolicyResponseBodyData) SetSignature(v string) *GetFileStoreUploadPolicyResponseBodyData {
	s.Signature = &v
	return s
}

type GetFileStoreUploadPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFileStoreUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFileStoreUploadPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileStoreUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetFileStoreUploadPolicyResponse) SetHeaders(v map[string]*string) *GetFileStoreUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetFileStoreUploadPolicyResponse) SetStatusCode(v int32) *GetFileStoreUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileStoreUploadPolicyResponse) SetBody(v *GetFileStoreUploadPolicyResponseBody) *GetFileStoreUploadPolicyResponse {
	s.Body = v
	return s
}

type GetPromptRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	PromptId *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
	Vars     *string `json:"Vars,omitempty" xml:"Vars,omitempty"`
}

func (s GetPromptRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPromptRequest) GoString() string {
	return s.String()
}

func (s *GetPromptRequest) SetAgentKey(v string) *GetPromptRequest {
	s.AgentKey = &v
	return s
}

func (s *GetPromptRequest) SetPromptId(v string) *GetPromptRequest {
	s.PromptId = &v
	return s
}

func (s *GetPromptRequest) SetVars(v string) *GetPromptRequest {
	s.Vars = &v
	return s
}

type GetPromptResponseBody struct {
	Code           *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetPromptResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *string                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPromptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPromptResponseBody) GoString() string {
	return s.String()
}

func (s *GetPromptResponseBody) SetCode(v string) *GetPromptResponseBody {
	s.Code = &v
	return s
}

func (s *GetPromptResponseBody) SetData(v *GetPromptResponseBodyData) *GetPromptResponseBody {
	s.Data = v
	return s
}

func (s *GetPromptResponseBody) SetHttpStatusCode(v string) *GetPromptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPromptResponseBody) SetMessage(v string) *GetPromptResponseBody {
	s.Message = &v
	return s
}

func (s *GetPromptResponseBody) SetRequestId(v string) *GetPromptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPromptResponseBody) SetSuccess(v bool) *GetPromptResponseBody {
	s.Success = &v
	return s
}

type GetPromptResponseBodyData struct {
	PromptContent *string `json:"PromptContent,omitempty" xml:"PromptContent,omitempty"`
	PromptId      *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
}

func (s GetPromptResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPromptResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPromptResponseBodyData) SetPromptContent(v string) *GetPromptResponseBodyData {
	s.PromptContent = &v
	return s
}

func (s *GetPromptResponseBodyData) SetPromptId(v string) *GetPromptResponseBodyData {
	s.PromptId = &v
	return s
}

type GetPromptResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPromptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPromptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPromptResponse) GoString() string {
	return s.String()
}

func (s *GetPromptResponse) SetHeaders(v map[string]*string) *GetPromptResponse {
	s.Headers = v
	return s
}

func (s *GetPromptResponse) SetStatusCode(v int32) *GetPromptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPromptResponse) SetBody(v *GetPromptResponseBody) *GetPromptResponse {
	s.Body = v
	return s
}

type GetText2ImageJobRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetText2ImageJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetText2ImageJobRequest) GoString() string {
	return s.String()
}

func (s *GetText2ImageJobRequest) SetAgentKey(v string) *GetText2ImageJobRequest {
	s.AgentKey = &v
	return s
}

func (s *GetText2ImageJobRequest) SetTaskId(v string) *GetText2ImageJobRequest {
	s.TaskId = &v
	return s
}

type GetText2ImageJobResponseBody struct {
	Images      []*GetText2ImageJobResponseBodyImages    `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId      *string                                  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskMetrics *GetText2ImageJobResponseBodyTaskMetrics `json:"TaskMetrics,omitempty" xml:"TaskMetrics,omitempty" type:"Struct"`
	TaskStatus  *string                                  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	Usage       []*GetText2ImageJobResponseBodyUsage     `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Repeated"`
}

func (s GetText2ImageJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetText2ImageJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetText2ImageJobResponseBody) SetImages(v []*GetText2ImageJobResponseBodyImages) *GetText2ImageJobResponseBody {
	s.Images = v
	return s
}

func (s *GetText2ImageJobResponseBody) SetRequestId(v string) *GetText2ImageJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetText2ImageJobResponseBody) SetTaskId(v string) *GetText2ImageJobResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetText2ImageJobResponseBody) SetTaskMetrics(v *GetText2ImageJobResponseBodyTaskMetrics) *GetText2ImageJobResponseBody {
	s.TaskMetrics = v
	return s
}

func (s *GetText2ImageJobResponseBody) SetTaskStatus(v string) *GetText2ImageJobResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetText2ImageJobResponseBody) SetUsage(v []*GetText2ImageJobResponseBodyUsage) *GetText2ImageJobResponseBody {
	s.Usage = v
	return s
}

type GetText2ImageJobResponseBodyImages struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	URL     *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetText2ImageJobResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s GetText2ImageJobResponseBodyImages) GoString() string {
	return s.String()
}

func (s *GetText2ImageJobResponseBodyImages) SetCode(v string) *GetText2ImageJobResponseBodyImages {
	s.Code = &v
	return s
}

func (s *GetText2ImageJobResponseBodyImages) SetMessage(v string) *GetText2ImageJobResponseBodyImages {
	s.Message = &v
	return s
}

func (s *GetText2ImageJobResponseBodyImages) SetURL(v string) *GetText2ImageJobResponseBodyImages {
	s.URL = &v
	return s
}

type GetText2ImageJobResponseBodyTaskMetrics struct {
	Failed    *int32 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	Succeeded *int32 `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
	Total     *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetText2ImageJobResponseBodyTaskMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetText2ImageJobResponseBodyTaskMetrics) GoString() string {
	return s.String()
}

func (s *GetText2ImageJobResponseBodyTaskMetrics) SetFailed(v int32) *GetText2ImageJobResponseBodyTaskMetrics {
	s.Failed = &v
	return s
}

func (s *GetText2ImageJobResponseBodyTaskMetrics) SetSucceeded(v int32) *GetText2ImageJobResponseBodyTaskMetrics {
	s.Succeeded = &v
	return s
}

func (s *GetText2ImageJobResponseBodyTaskMetrics) SetTotal(v int32) *GetText2ImageJobResponseBodyTaskMetrics {
	s.Total = &v
	return s
}

type GetText2ImageJobResponseBodyUsage struct {
	ImageCount *int32 `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
}

func (s GetText2ImageJobResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s GetText2ImageJobResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetText2ImageJobResponseBodyUsage) SetImageCount(v int32) *GetText2ImageJobResponseBodyUsage {
	s.ImageCount = &v
	return s
}

type GetText2ImageJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetText2ImageJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetText2ImageJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetText2ImageJobResponse) GoString() string {
	return s.String()
}

func (s *GetText2ImageJobResponse) SetHeaders(v map[string]*string) *GetText2ImageJobResponse {
	s.Headers = v
	return s
}

func (s *GetText2ImageJobResponse) SetStatusCode(v int32) *GetText2ImageJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetText2ImageJobResponse) SetBody(v *GetText2ImageJobResponseBody) *GetText2ImageJobResponse {
	s.Body = v
	return s
}

type ImportUserDocumentRequest struct {
	AgentKey    *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileStoreId *int64  `json:"FileStoreId,omitempty" xml:"FileStoreId,omitempty"`
	OssPath     *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	StoreId     *int64  `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ImportUserDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportUserDocumentRequest) GoString() string {
	return s.String()
}

func (s *ImportUserDocumentRequest) SetAgentKey(v string) *ImportUserDocumentRequest {
	s.AgentKey = &v
	return s
}

func (s *ImportUserDocumentRequest) SetFileName(v string) *ImportUserDocumentRequest {
	s.FileName = &v
	return s
}

func (s *ImportUserDocumentRequest) SetFileStoreId(v int64) *ImportUserDocumentRequest {
	s.FileStoreId = &v
	return s
}

func (s *ImportUserDocumentRequest) SetOssPath(v string) *ImportUserDocumentRequest {
	s.OssPath = &v
	return s
}

func (s *ImportUserDocumentRequest) SetStoreId(v int64) *ImportUserDocumentRequest {
	s.StoreId = &v
	return s
}

func (s *ImportUserDocumentRequest) SetUserId(v string) *ImportUserDocumentRequest {
	s.UserId = &v
	return s
}

type ImportUserDocumentResponseBody struct {
	Data      *ImportUserDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportUserDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportUserDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *ImportUserDocumentResponseBody) SetData(v *ImportUserDocumentResponseBodyData) *ImportUserDocumentResponseBody {
	s.Data = v
	return s
}

func (s *ImportUserDocumentResponseBody) SetErrorCode(v string) *ImportUserDocumentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ImportUserDocumentResponseBody) SetErrorMsg(v string) *ImportUserDocumentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ImportUserDocumentResponseBody) SetRequestId(v string) *ImportUserDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportUserDocumentResponseBody) SetSuccess(v bool) *ImportUserDocumentResponseBody {
	s.Success = &v
	return s
}

type ImportUserDocumentResponseBodyData struct {
	DataId     *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	DataStatus *int64  `json:"DataStatus,omitempty" xml:"DataStatus,omitempty"`
}

func (s ImportUserDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImportUserDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportUserDocumentResponseBodyData) SetDataId(v string) *ImportUserDocumentResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ImportUserDocumentResponseBodyData) SetDataStatus(v int64) *ImportUserDocumentResponseBodyData {
	s.DataStatus = &v
	return s
}

type ImportUserDocumentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImportUserDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportUserDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportUserDocumentResponse) GoString() string {
	return s.String()
}

func (s *ImportUserDocumentResponse) SetHeaders(v map[string]*string) *ImportUserDocumentResponse {
	s.Headers = v
	return s
}

func (s *ImportUserDocumentResponse) SetStatusCode(v int32) *ImportUserDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportUserDocumentResponse) SetBody(v *ImportUserDocumentResponseBody) *ImportUserDocumentResponse {
	s.Body = v
	return s
}

type ListDocsRequest struct {
	AgentKey *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Name     *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNo   *int32    `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StoreId  *string   `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	TagIds   []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListDocsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDocsRequest) GoString() string {
	return s.String()
}

func (s *ListDocsRequest) SetAgentKey(v string) *ListDocsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListDocsRequest) SetName(v string) *ListDocsRequest {
	s.Name = &v
	return s
}

func (s *ListDocsRequest) SetPageNo(v int32) *ListDocsRequest {
	s.PageNo = &v
	return s
}

func (s *ListDocsRequest) SetPageSize(v int32) *ListDocsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDocsRequest) SetStoreId(v string) *ListDocsRequest {
	s.StoreId = &v
	return s
}

func (s *ListDocsRequest) SetTagIds(v []*string) *ListDocsRequest {
	s.TagIds = v
	return s
}

type ListDocsShrinkRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNo       *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StoreId      *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	TagIdsShrink *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
}

func (s ListDocsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDocsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDocsShrinkRequest) SetAgentKey(v string) *ListDocsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListDocsShrinkRequest) SetName(v string) *ListDocsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListDocsShrinkRequest) SetPageNo(v int32) *ListDocsShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *ListDocsShrinkRequest) SetPageSize(v int32) *ListDocsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDocsShrinkRequest) SetStoreId(v string) *ListDocsShrinkRequest {
	s.StoreId = &v
	return s
}

func (s *ListDocsShrinkRequest) SetTagIdsShrink(v string) *ListDocsShrinkRequest {
	s.TagIdsShrink = &v
	return s
}

type ListDocsResponseBody struct {
	Docs      []*ListDocsResponseBodyDocs `json:"Docs,omitempty" xml:"Docs,omitempty" type:"Repeated"`
	PageNo    *int32                      `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDocsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDocsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDocsResponseBody) SetDocs(v []*ListDocsResponseBodyDocs) *ListDocsResponseBody {
	s.Docs = v
	return s
}

func (s *ListDocsResponseBody) SetPageNo(v int32) *ListDocsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListDocsResponseBody) SetPageSize(v int32) *ListDocsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDocsResponseBody) SetRequestId(v string) *ListDocsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDocsResponseBody) SetTotal(v int32) *ListDocsResponseBody {
	s.Total = &v
	return s
}

type ListDocsResponseBodyDocs struct {
	DocId      *string                         `json:"DocId,omitempty" xml:"DocId,omitempty"`
	FailReason *string                         `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	Name       *string                         `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId    *string                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Size       *string                         `json:"Size,omitempty" xml:"Size,omitempty"`
	Status     *string                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags       []*ListDocsResponseBodyDocsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Type       *string                         `json:"Type,omitempty" xml:"Type,omitempty"`
	UploadTime *string                         `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s ListDocsResponseBodyDocs) String() string {
	return tea.Prettify(s)
}

func (s ListDocsResponseBodyDocs) GoString() string {
	return s.String()
}

func (s *ListDocsResponseBodyDocs) SetDocId(v string) *ListDocsResponseBodyDocs {
	s.DocId = &v
	return s
}

func (s *ListDocsResponseBodyDocs) SetFailReason(v string) *ListDocsResponseBodyDocs {
	s.FailReason = &v
	return s
}

func (s *ListDocsResponseBodyDocs) SetName(v string) *ListDocsResponseBodyDocs {
	s.Name = &v
	return s
}

func (s *ListDocsResponseBodyDocs) SetOwnerId(v string) *ListDocsResponseBodyDocs {
	s.OwnerId = &v
	return s
}

func (s *ListDocsResponseBodyDocs) SetSize(v string) *ListDocsResponseBodyDocs {
	s.Size = &v
	return s
}

func (s *ListDocsResponseBodyDocs) SetStatus(v string) *ListDocsResponseBodyDocs {
	s.Status = &v
	return s
}

func (s *ListDocsResponseBodyDocs) SetTags(v []*ListDocsResponseBodyDocsTags) *ListDocsResponseBodyDocs {
	s.Tags = v
	return s
}

func (s *ListDocsResponseBodyDocs) SetType(v string) *ListDocsResponseBodyDocs {
	s.Type = &v
	return s
}

func (s *ListDocsResponseBodyDocs) SetUploadTime(v string) *ListDocsResponseBodyDocs {
	s.UploadTime = &v
	return s
}

type ListDocsResponseBodyDocsTags struct {
	TagId   *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListDocsResponseBodyDocsTags) String() string {
	return tea.Prettify(s)
}

func (s ListDocsResponseBodyDocsTags) GoString() string {
	return s.String()
}

func (s *ListDocsResponseBodyDocsTags) SetTagId(v string) *ListDocsResponseBodyDocsTags {
	s.TagId = &v
	return s
}

func (s *ListDocsResponseBodyDocsTags) SetTagName(v string) *ListDocsResponseBodyDocsTags {
	s.TagName = &v
	return s
}

type ListDocsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDocsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDocsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDocsResponse) GoString() string {
	return s.String()
}

func (s *ListDocsResponse) SetHeaders(v map[string]*string) *ListDocsResponse {
	s.Headers = v
	return s
}

func (s *ListDocsResponse) SetStatusCode(v int32) *ListDocsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDocsResponse) SetBody(v *ListDocsResponseBody) *ListDocsResponse {
	s.Body = v
	return s
}

type ListDocumentTagsRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNo   *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TagId    *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListDocumentTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDocumentTagsRequest) GoString() string {
	return s.String()
}

func (s *ListDocumentTagsRequest) SetAgentKey(v string) *ListDocumentTagsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListDocumentTagsRequest) SetName(v string) *ListDocumentTagsRequest {
	s.Name = &v
	return s
}

func (s *ListDocumentTagsRequest) SetPageNo(v int32) *ListDocumentTagsRequest {
	s.PageNo = &v
	return s
}

func (s *ListDocumentTagsRequest) SetPageSize(v int32) *ListDocumentTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDocumentTagsRequest) SetTagId(v string) *ListDocumentTagsRequest {
	s.TagId = &v
	return s
}

type ListDocumentTagsResponseBody struct {
	PageNo    *int32                                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagList   []*ListDocumentTagsResponseBodyTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	Total     *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDocumentTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDocumentTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDocumentTagsResponseBody) SetPageNo(v int32) *ListDocumentTagsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListDocumentTagsResponseBody) SetPageSize(v int32) *ListDocumentTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDocumentTagsResponseBody) SetRequestId(v string) *ListDocumentTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDocumentTagsResponseBody) SetTagList(v []*ListDocumentTagsResponseBodyTagList) *ListDocumentTagsResponseBody {
	s.TagList = v
	return s
}

func (s *ListDocumentTagsResponseBody) SetTotal(v int32) *ListDocumentTagsResponseBody {
	s.Total = &v
	return s
}

type ListDocumentTagsResponseBodyTagList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListDocumentTagsResponseBodyTagList) String() string {
	return tea.Prettify(s)
}

func (s ListDocumentTagsResponseBodyTagList) GoString() string {
	return s.String()
}

func (s *ListDocumentTagsResponseBodyTagList) SetName(v string) *ListDocumentTagsResponseBodyTagList {
	s.Name = &v
	return s
}

func (s *ListDocumentTagsResponseBodyTagList) SetTagId(v string) *ListDocumentTagsResponseBodyTagList {
	s.TagId = &v
	return s
}

type ListDocumentTagsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDocumentTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDocumentTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDocumentTagsResponse) GoString() string {
	return s.String()
}

func (s *ListDocumentTagsResponse) SetHeaders(v map[string]*string) *ListDocumentTagsResponse {
	s.Headers = v
	return s
}

func (s *ListDocumentTagsResponse) SetStatusCode(v int32) *ListDocumentTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDocumentTagsResponse) SetBody(v *ListDocumentTagsResponseBody) *ListDocumentTagsResponse {
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

type QueryUserDocumentRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DataId   *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryUserDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserDocumentRequest) GoString() string {
	return s.String()
}

func (s *QueryUserDocumentRequest) SetAgentKey(v string) *QueryUserDocumentRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryUserDocumentRequest) SetDataId(v string) *QueryUserDocumentRequest {
	s.DataId = &v
	return s
}

func (s *QueryUserDocumentRequest) SetUserId(v string) *QueryUserDocumentRequest {
	s.UserId = &v
	return s
}

type QueryUserDocumentResponseBody struct {
	Data      *QueryUserDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                            `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryUserDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserDocumentResponseBody) SetData(v *QueryUserDocumentResponseBodyData) *QueryUserDocumentResponseBody {
	s.Data = v
	return s
}

func (s *QueryUserDocumentResponseBody) SetErrorCode(v string) *QueryUserDocumentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryUserDocumentResponseBody) SetErrorMsg(v string) *QueryUserDocumentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryUserDocumentResponseBody) SetRequestId(v string) *QueryUserDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserDocumentResponseBody) SetSuccess(v bool) *QueryUserDocumentResponseBody {
	s.Success = &v
	return s
}

type QueryUserDocumentResponseBodyData struct {
	DataId     *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	DataStatus *int64  `json:"DataStatus,omitempty" xml:"DataStatus,omitempty"`
}

func (s QueryUserDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryUserDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryUserDocumentResponseBodyData) SetDataId(v string) *QueryUserDocumentResponseBodyData {
	s.DataId = &v
	return s
}

func (s *QueryUserDocumentResponseBodyData) SetDataStatus(v int64) *QueryUserDocumentResponseBodyData {
	s.DataStatus = &v
	return s
}

type QueryUserDocumentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserDocumentResponse) GoString() string {
	return s.String()
}

func (s *QueryUserDocumentResponse) SetHeaders(v map[string]*string) *QueryUserDocumentResponse {
	s.Headers = v
	return s
}

func (s *QueryUserDocumentResponse) SetStatusCode(v int32) *QueryUserDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserDocumentResponse) SetBody(v *QueryUserDocumentResponseBody) *QueryUserDocumentResponse {
	s.Body = v
	return s
}

type SubmitDocumentImportJobRequest struct {
	AgentKey *string                               `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Docs     []*SubmitDocumentImportJobRequestDocs `json:"Docs,omitempty" xml:"Docs,omitempty" type:"Repeated"`
}

func (s SubmitDocumentImportJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentImportJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocumentImportJobRequest) SetAgentKey(v string) *SubmitDocumentImportJobRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitDocumentImportJobRequest) SetDocs(v []*SubmitDocumentImportJobRequestDocs) *SubmitDocumentImportJobRequest {
	s.Docs = v
	return s
}

type SubmitDocumentImportJobRequestDocs struct {
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *string   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StoreId *string   `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	TagIds  []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	Type    *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	URL     *string   `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SubmitDocumentImportJobRequestDocs) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentImportJobRequestDocs) GoString() string {
	return s.String()
}

func (s *SubmitDocumentImportJobRequestDocs) SetName(v string) *SubmitDocumentImportJobRequestDocs {
	s.Name = &v
	return s
}

func (s *SubmitDocumentImportJobRequestDocs) SetOwnerId(v string) *SubmitDocumentImportJobRequestDocs {
	s.OwnerId = &v
	return s
}

func (s *SubmitDocumentImportJobRequestDocs) SetStoreId(v string) *SubmitDocumentImportJobRequestDocs {
	s.StoreId = &v
	return s
}

func (s *SubmitDocumentImportJobRequestDocs) SetTagIds(v []*string) *SubmitDocumentImportJobRequestDocs {
	s.TagIds = v
	return s
}

func (s *SubmitDocumentImportJobRequestDocs) SetType(v string) *SubmitDocumentImportJobRequestDocs {
	s.Type = &v
	return s
}

func (s *SubmitDocumentImportJobRequestDocs) SetURL(v string) *SubmitDocumentImportJobRequestDocs {
	s.URL = &v
	return s
}

type SubmitDocumentImportJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDocumentImportJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentImportJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocumentImportJobResponseBody) SetJobId(v string) *SubmitDocumentImportJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitDocumentImportJobResponseBody) SetRequestId(v string) *SubmitDocumentImportJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDocumentImportJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitDocumentImportJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitDocumentImportJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDocumentImportJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocumentImportJobResponse) SetHeaders(v map[string]*string) *SubmitDocumentImportJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocumentImportJobResponse) SetStatusCode(v int32) *SubmitDocumentImportJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocumentImportJobResponse) SetBody(v *SubmitDocumentImportJobResponseBody) *SubmitDocumentImportJobResponse {
	s.Body = v
	return s
}

type SubmitText2ImageJobRequest struct {
	AgentKey       *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	N              *int32  `json:"N,omitempty" xml:"N,omitempty"`
	NegativePrompt *string `json:"NegativePrompt,omitempty" xml:"NegativePrompt,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Seed           *int32  `json:"Seed,omitempty" xml:"Seed,omitempty"`
	Size           *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Style          *string `json:"Style,omitempty" xml:"Style,omitempty"`
}

func (s SubmitText2ImageJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitText2ImageJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitText2ImageJobRequest) SetAgentKey(v string) *SubmitText2ImageJobRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitText2ImageJobRequest) SetAppId(v string) *SubmitText2ImageJobRequest {
	s.AppId = &v
	return s
}

func (s *SubmitText2ImageJobRequest) SetN(v int32) *SubmitText2ImageJobRequest {
	s.N = &v
	return s
}

func (s *SubmitText2ImageJobRequest) SetNegativePrompt(v string) *SubmitText2ImageJobRequest {
	s.NegativePrompt = &v
	return s
}

func (s *SubmitText2ImageJobRequest) SetPrompt(v string) *SubmitText2ImageJobRequest {
	s.Prompt = &v
	return s
}

func (s *SubmitText2ImageJobRequest) SetSeed(v int32) *SubmitText2ImageJobRequest {
	s.Seed = &v
	return s
}

func (s *SubmitText2ImageJobRequest) SetSize(v string) *SubmitText2ImageJobRequest {
	s.Size = &v
	return s
}

func (s *SubmitText2ImageJobRequest) SetStyle(v string) *SubmitText2ImageJobRequest {
	s.Style = &v
	return s
}

type SubmitText2ImageJobResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s SubmitText2ImageJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitText2ImageJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitText2ImageJobResponseBody) SetRequestId(v string) *SubmitText2ImageJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitText2ImageJobResponseBody) SetTaskId(v string) *SubmitText2ImageJobResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitText2ImageJobResponseBody) SetTaskStatus(v string) *SubmitText2ImageJobResponseBody {
	s.TaskStatus = &v
	return s
}

type SubmitText2ImageJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitText2ImageJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitText2ImageJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitText2ImageJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitText2ImageJobResponse) SetHeaders(v map[string]*string) *SubmitText2ImageJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitText2ImageJobResponse) SetStatusCode(v int32) *SubmitText2ImageJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitText2ImageJobResponse) SetBody(v *SubmitText2ImageJobResponseBody) *SubmitText2ImageJobResponse {
	s.Body = v
	return s
}

type UpdateDocAttributeRequest struct {
	AgentKey   *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DelAllTags *bool     `json:"DelAllTags,omitempty" xml:"DelAllTags,omitempty"`
	DocId      *string   `json:"DocId,omitempty" xml:"DocId,omitempty"`
	Name       *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	TagIds     []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s UpdateDocAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocAttributeRequest) SetAgentKey(v string) *UpdateDocAttributeRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDocAttributeRequest) SetDelAllTags(v bool) *UpdateDocAttributeRequest {
	s.DelAllTags = &v
	return s
}

func (s *UpdateDocAttributeRequest) SetDocId(v string) *UpdateDocAttributeRequest {
	s.DocId = &v
	return s
}

func (s *UpdateDocAttributeRequest) SetName(v string) *UpdateDocAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateDocAttributeRequest) SetTagIds(v []*string) *UpdateDocAttributeRequest {
	s.TagIds = v
	return s
}

type UpdateDocAttributeShrinkRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DelAllTags   *bool   `json:"DelAllTags,omitempty" xml:"DelAllTags,omitempty"`
	DocId        *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TagIdsShrink *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
}

func (s UpdateDocAttributeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocAttributeShrinkRequest) SetAgentKey(v string) *UpdateDocAttributeShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDocAttributeShrinkRequest) SetDelAllTags(v bool) *UpdateDocAttributeShrinkRequest {
	s.DelAllTags = &v
	return s
}

func (s *UpdateDocAttributeShrinkRequest) SetDocId(v string) *UpdateDocAttributeShrinkRequest {
	s.DocId = &v
	return s
}

func (s *UpdateDocAttributeShrinkRequest) SetName(v string) *UpdateDocAttributeShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateDocAttributeShrinkRequest) SetTagIdsShrink(v string) *UpdateDocAttributeShrinkRequest {
	s.TagIdsShrink = &v
	return s
}

type UpdateDocAttributeResponseBody struct {
	DocId     *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDocAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDocAttributeResponseBody) SetDocId(v string) *UpdateDocAttributeResponseBody {
	s.DocId = &v
	return s
}

func (s *UpdateDocAttributeResponseBody) SetRequestId(v string) *UpdateDocAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDocAttributeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDocAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDocAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateDocAttributeResponse) SetHeaders(v map[string]*string) *UpdateDocAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateDocAttributeResponse) SetStatusCode(v int32) *UpdateDocAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDocAttributeResponse) SetBody(v *UpdateDocAttributeResponseBody) *UpdateDocAttributeResponse {
	s.Body = v
	return s
}

type UpdateDocumentTagRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TagId    *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s UpdateDocumentTagRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentTagRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocumentTagRequest) SetAgentKey(v string) *UpdateDocumentTagRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDocumentTagRequest) SetName(v string) *UpdateDocumentTagRequest {
	s.Name = &v
	return s
}

func (s *UpdateDocumentTagRequest) SetTagId(v string) *UpdateDocumentTagRequest {
	s.TagId = &v
	return s
}

type UpdateDocumentTagResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagId     *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s UpdateDocumentTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentTagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDocumentTagResponseBody) SetRequestId(v string) *UpdateDocumentTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDocumentTagResponseBody) SetTagId(v string) *UpdateDocumentTagResponseBody {
	s.TagId = &v
	return s
}

type UpdateDocumentTagResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDocumentTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDocumentTagResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentTagResponse) GoString() string {
	return s.String()
}

func (s *UpdateDocumentTagResponse) SetHeaders(v map[string]*string) *UpdateDocumentTagResponse {
	s.Headers = v
	return s
}

func (s *UpdateDocumentTagResponse) SetStatusCode(v int32) *UpdateDocumentTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDocumentTagResponse) SetBody(v *UpdateDocumentTagResponseBody) *UpdateDocumentTagResponse {
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

func (client *Client) CreateDocumentTagWithOptions(request *CreateDocumentTagRequest, runtime *util.RuntimeOptions) (_result *CreateDocumentTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDocumentTag"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDocumentTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDocumentTag(request *CreateDocumentTagRequest) (_result *CreateDocumentTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDocumentTagResponse{}
	_body, _err := client.CreateDocumentTagWithOptions(request, runtime)
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

func (client *Client) DeleteDocWithOptions(request *DeleteDocRequest, runtime *util.RuntimeOptions) (_result *DeleteDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		query["DocId"] = request.DocId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDoc"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDoc(request *DeleteDocRequest) (_result *DeleteDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDocResponse{}
	_body, _err := client.DeleteDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDocumentTagWithOptions(request *DeleteDocumentTagRequest, runtime *util.RuntimeOptions) (_result *DeleteDocumentTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDocumentTag"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDocumentTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDocumentTag(request *DeleteDocumentTagRequest) (_result *DeleteDocumentTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDocumentTagResponse{}
	_body, _err := client.DeleteDocumentTagWithOptions(request, runtime)
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

func (client *Client) DescribeDocWithOptions(request *DescribeDocRequest, runtime *util.RuntimeOptions) (_result *DescribeDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		query["DocId"] = request.DocId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDoc"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDoc(request *DescribeDocRequest) (_result *DescribeDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDocResponse{}
	_body, _err := client.DescribeDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDocumentImportJobWithOptions(request *DescribeDocumentImportJobRequest, runtime *util.RuntimeOptions) (_result *DescribeDocumentImportJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDocumentImportJob"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDocumentImportJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDocumentImportJob(request *DescribeDocumentImportJobRequest) (_result *DescribeDocumentImportJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDocumentImportJobResponse{}
	_body, _err := client.DescribeDocumentImportJobWithOptions(request, runtime)
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

func (client *Client) GetFileStoreUploadPolicyWithOptions(request *GetFileStoreUploadPolicyRequest, runtime *util.RuntimeOptions) (_result *GetFileStoreUploadPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileStoreId)) {
		query["FileStoreId"] = request.FileStoreId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFileStoreUploadPolicy"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileStoreUploadPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFileStoreUploadPolicy(request *GetFileStoreUploadPolicyRequest) (_result *GetFileStoreUploadPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFileStoreUploadPolicyResponse{}
	_body, _err := client.GetFileStoreUploadPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPromptWithOptions(request *GetPromptRequest, runtime *util.RuntimeOptions) (_result *GetPromptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.PromptId)) {
		query["PromptId"] = request.PromptId
	}

	if !tea.BoolValue(util.IsUnset(request.Vars)) {
		query["Vars"] = request.Vars
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPrompt"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPromptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPrompt(request *GetPromptRequest) (_result *GetPromptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPromptResponse{}
	_body, _err := client.GetPromptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetText2ImageJobWithOptions(request *GetText2ImageJobRequest, runtime *util.RuntimeOptions) (_result *GetText2ImageJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetText2ImageJob"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetText2ImageJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetText2ImageJob(request *GetText2ImageJobRequest) (_result *GetText2ImageJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetText2ImageJobResponse{}
	_body, _err := client.GetText2ImageJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportUserDocumentWithOptions(request *ImportUserDocumentRequest, runtime *util.RuntimeOptions) (_result *ImportUserDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileStoreId)) {
		query["FileStoreId"] = request.FileStoreId
	}

	if !tea.BoolValue(util.IsUnset(request.OssPath)) {
		query["OssPath"] = request.OssPath
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		query["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportUserDocument"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportUserDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportUserDocument(request *ImportUserDocumentRequest) (_result *ImportUserDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportUserDocumentResponse{}
	_body, _err := client.ImportUserDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDocsWithOptions(tmpReq *ListDocsRequest, runtime *util.RuntimeOptions) (_result *ListDocsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDocsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TagIds)) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, tea.String("TagIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		query["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdsShrink)) {
		query["TagIds"] = request.TagIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDocs"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDocsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDocs(request *ListDocsRequest) (_result *ListDocsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDocsResponse{}
	_body, _err := client.ListDocsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDocumentTagsWithOptions(request *ListDocumentTagsRequest, runtime *util.RuntimeOptions) (_result *ListDocumentTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDocumentTags"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDocumentTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDocumentTags(request *ListDocumentTagsRequest) (_result *ListDocumentTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDocumentTagsResponse{}
	_body, _err := client.ListDocumentTagsWithOptions(request, runtime)
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

func (client *Client) QueryUserDocumentWithOptions(request *QueryUserDocumentRequest, runtime *util.RuntimeOptions) (_result *QueryUserDocumentResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUserDocument"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserDocument(request *QueryUserDocumentRequest) (_result *QueryUserDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserDocumentResponse{}
	_body, _err := client.QueryUserDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocumentImportJobWithOptions(request *SubmitDocumentImportJobRequest, runtime *util.RuntimeOptions) (_result *SubmitDocumentImportJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Docs)) {
		query["Docs"] = request.Docs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDocumentImportJob"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitDocumentImportJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitDocumentImportJob(request *SubmitDocumentImportJobRequest) (_result *SubmitDocumentImportJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDocumentImportJobResponse{}
	_body, _err := client.SubmitDocumentImportJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitText2ImageJobWithOptions(request *SubmitText2ImageJobRequest, runtime *util.RuntimeOptions) (_result *SubmitText2ImageJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.N)) {
		query["N"] = request.N
	}

	if !tea.BoolValue(util.IsUnset(request.NegativePrompt)) {
		query["NegativePrompt"] = request.NegativePrompt
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		query["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.Seed)) {
		query["Seed"] = request.Seed
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Style)) {
		query["Style"] = request.Style
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitText2ImageJob"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitText2ImageJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitText2ImageJob(request *SubmitText2ImageJobRequest) (_result *SubmitText2ImageJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitText2ImageJobResponse{}
	_body, _err := client.SubmitText2ImageJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDocAttributeWithOptions(tmpReq *UpdateDocAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateDocAttributeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDocAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TagIds)) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, tea.String("TagIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.DelAllTags)) {
		query["DelAllTags"] = request.DelAllTags
	}

	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		query["DocId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdsShrink)) {
		query["TagIds"] = request.TagIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDocAttribute"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDocAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDocAttribute(request *UpdateDocAttributeRequest) (_result *UpdateDocAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDocAttributeResponse{}
	_body, _err := client.UpdateDocAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDocumentTagWithOptions(request *UpdateDocumentTagRequest, runtime *util.RuntimeOptions) (_result *UpdateDocumentTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		query["TagId"] = request.TagId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDocumentTag"),
		Version:     tea.String("2023-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDocumentTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDocumentTag(request *UpdateDocumentTagRequest) (_result *UpdateDocumentTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDocumentTagResponse{}
	_body, _err := client.UpdateDocumentTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
