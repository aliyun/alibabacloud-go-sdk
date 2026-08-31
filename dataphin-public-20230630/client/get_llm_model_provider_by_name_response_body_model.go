// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLlmModelProviderByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLlmModelProviderByNameResponseBody
	GetCode() *string
	SetData(v *GetLlmModelProviderByNameResponseBodyData) *GetLlmModelProviderByNameResponseBody
	GetData() *GetLlmModelProviderByNameResponseBodyData
	SetHttpStatusCode(v int32) *GetLlmModelProviderByNameResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetLlmModelProviderByNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLlmModelProviderByNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLlmModelProviderByNameResponseBody
	GetSuccess() *bool
}

type GetLlmModelProviderByNameResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The large language model service provider and available models.
	Data *GetLlmModelProviderByNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLlmModelProviderByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLlmModelProviderByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetLlmModelProviderByNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLlmModelProviderByNameResponseBody) GetData() *GetLlmModelProviderByNameResponseBodyData {
	return s.Data
}

func (s *GetLlmModelProviderByNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetLlmModelProviderByNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLlmModelProviderByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLlmModelProviderByNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLlmModelProviderByNameResponseBody) SetCode(v string) *GetLlmModelProviderByNameResponseBody {
	s.Code = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBody) SetData(v *GetLlmModelProviderByNameResponseBodyData) *GetLlmModelProviderByNameResponseBody {
	s.Data = v
	return s
}

func (s *GetLlmModelProviderByNameResponseBody) SetHttpStatusCode(v int32) *GetLlmModelProviderByNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBody) SetMessage(v string) *GetLlmModelProviderByNameResponseBody {
	s.Message = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBody) SetRequestId(v string) *GetLlmModelProviderByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBody) SetSuccess(v bool) *GetLlmModelProviderByNameResponseBody {
	s.Success = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLlmModelProviderByNameResponseBodyData struct {
	// The base URL of the model.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// The enabling status.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the model service provider.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The list of models provided by the model service provider.
	LlmModels []*GetLlmModelProviderByNameResponseBodyDataLlmModels `json:"LlmModels,omitempty" xml:"LlmModels,omitempty" type:"Repeated"`
	// The source of the model service provider. Valid values:
	//
	// - BUILTIN_MODEL: built-in
	//
	// - BYOM: user-provided
	//
	// example:
	//
	// BUILTIN_MODEL
	ProviderSource *string `json:"ProviderSource,omitempty" xml:"ProviderSource,omitempty"`
	// The service provider type. Valid values:
	//
	// - BAILIAN
	//
	// - DEEPSEEK
	//
	// - AI_STACK
	//
	// - VLLM
	//
	// - AGENT_ONE
	//
	// - DATAPHIN
	//
	// example:
	//
	// BAILIAN
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// The name of the model service provider.
	//
	// example:
	//
	// bailian
	ServiceProvider *string `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
}

func (s GetLlmModelProviderByNameResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLlmModelProviderByNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLlmModelProviderByNameResponseBodyData) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *GetLlmModelProviderByNameResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetLlmModelProviderByNameResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetLlmModelProviderByNameResponseBodyData) GetLlmModels() []*GetLlmModelProviderByNameResponseBodyDataLlmModels {
	return s.LlmModels
}

func (s *GetLlmModelProviderByNameResponseBodyData) GetProviderSource() *string {
	return s.ProviderSource
}

func (s *GetLlmModelProviderByNameResponseBodyData) GetProviderType() *string {
	return s.ProviderType
}

func (s *GetLlmModelProviderByNameResponseBodyData) GetServiceProvider() *string {
	return s.ServiceProvider
}

func (s *GetLlmModelProviderByNameResponseBodyData) SetBaseUrl(v string) *GetLlmModelProviderByNameResponseBodyData {
	s.BaseUrl = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyData) SetEnabled(v bool) *GetLlmModelProviderByNameResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyData) SetId(v int64) *GetLlmModelProviderByNameResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyData) SetLlmModels(v []*GetLlmModelProviderByNameResponseBodyDataLlmModels) *GetLlmModelProviderByNameResponseBodyData {
	s.LlmModels = v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyData) SetProviderSource(v string) *GetLlmModelProviderByNameResponseBodyData {
	s.ProviderSource = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyData) SetProviderType(v string) *GetLlmModelProviderByNameResponseBodyData {
	s.ProviderType = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyData) SetServiceProvider(v string) *GetLlmModelProviderByNameResponseBodyData {
	s.ServiceProvider = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyData) Validate() error {
	if s.LlmModels != nil {
		for _, item := range s.LlmModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLlmModelProviderByNameResponseBodyDataLlmModels struct {
	// The Chinese name of the model.
	//
	// example:
	//
	// 通义千问-Max
	CnName *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	// The model description.
	//
	// example:
	//
	// General text generation model
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of available dimension values for the embedding model. Only embedding models have this value.
	EmbeddingDimensions []*int32 `json:"EmbeddingDimensions,omitempty" xml:"EmbeddingDimensions,omitempty" type:"Repeated"`
	// The enabling status.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The model invocation method. Valid values:
	//
	// - OPEN_AI
	//
	// - DASH_SCOPE
	//
	// example:
	//
	// DASH_SCOPE
	InvokeType *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// The model ID.
	//
	// example:
	//
	// 1001
	ModelId *int64 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The list of model modality types. Valid values:
	//
	// - TEXT: text
	//
	// - IMAGE: image
	//
	// - AUDIO: audio
	//
	// - VIDEO: video
	//
	// - EMBEDDING: embedding
	ModelTypes []*string `json:"ModelTypes,omitempty" xml:"ModelTypes,omitempty" type:"Repeated"`
	// The model name.
	//
	// example:
	//
	// qwen-max
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The service provider.
	//
	// example:
	//
	// bailian
	ServiceProvider *string `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
	// The list of model task capabilities. Valid values:
	//
	// - ASR: speech recognition
	//
	// - TTS: speech synthesis
	//
	// - TRANSLATION: speech translation
	Tasks []*string `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s GetLlmModelProviderByNameResponseBodyDataLlmModels) String() string {
	return dara.Prettify(s)
}

func (s GetLlmModelProviderByNameResponseBodyDataLlmModels) GoString() string {
	return s.String()
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) GetCnName() *string {
	return s.CnName
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) GetDescription() *string {
	return s.Description
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) GetEmbeddingDimensions() []*int32 {
	return s.EmbeddingDimensions
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) GetInvokeType() *string {
	return s.InvokeType
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) GetModelId() *int64 {
	return s.ModelId
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) GetModelTypes() []*string {
	return s.ModelTypes
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) GetName() *string {
	return s.Name
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) GetServiceProvider() *string {
	return s.ServiceProvider
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) GetTasks() []*string {
	return s.Tasks
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) SetCnName(v string) *GetLlmModelProviderByNameResponseBodyDataLlmModels {
	s.CnName = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) SetDescription(v string) *GetLlmModelProviderByNameResponseBodyDataLlmModels {
	s.Description = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) SetEmbeddingDimensions(v []*int32) *GetLlmModelProviderByNameResponseBodyDataLlmModels {
	s.EmbeddingDimensions = v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) SetEnabled(v bool) *GetLlmModelProviderByNameResponseBodyDataLlmModels {
	s.Enabled = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) SetInvokeType(v string) *GetLlmModelProviderByNameResponseBodyDataLlmModels {
	s.InvokeType = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) SetModelId(v int64) *GetLlmModelProviderByNameResponseBodyDataLlmModels {
	s.ModelId = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) SetModelTypes(v []*string) *GetLlmModelProviderByNameResponseBodyDataLlmModels {
	s.ModelTypes = v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) SetName(v string) *GetLlmModelProviderByNameResponseBodyDataLlmModels {
	s.Name = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) SetServiceProvider(v string) *GetLlmModelProviderByNameResponseBodyDataLlmModels {
	s.ServiceProvider = &v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) SetTasks(v []*string) *GetLlmModelProviderByNameResponseBodyDataLlmModels {
	s.Tasks = v
	return s
}

func (s *GetLlmModelProviderByNameResponseBodyDataLlmModels) Validate() error {
	return dara.Validate(s)
}
