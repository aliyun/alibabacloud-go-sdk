// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLlmModelProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLlmModelProvidersResponseBody
	GetCode() *string
	SetData(v []*GetLlmModelProvidersResponseBodyData) *GetLlmModelProvidersResponseBody
	GetData() []*GetLlmModelProvidersResponseBodyData
	SetHttpStatusCode(v int32) *GetLlmModelProvidersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetLlmModelProvidersResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLlmModelProvidersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLlmModelProvidersResponseBody
	GetSuccess() *bool
}

type GetLlmModelProvidersResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetLlmModelProvidersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLlmModelProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLlmModelProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *GetLlmModelProvidersResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLlmModelProvidersResponseBody) GetData() []*GetLlmModelProvidersResponseBodyData {
	return s.Data
}

func (s *GetLlmModelProvidersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetLlmModelProvidersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLlmModelProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLlmModelProvidersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLlmModelProvidersResponseBody) SetCode(v string) *GetLlmModelProvidersResponseBody {
	s.Code = &v
	return s
}

func (s *GetLlmModelProvidersResponseBody) SetData(v []*GetLlmModelProvidersResponseBodyData) *GetLlmModelProvidersResponseBody {
	s.Data = v
	return s
}

func (s *GetLlmModelProvidersResponseBody) SetHttpStatusCode(v int32) *GetLlmModelProvidersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLlmModelProvidersResponseBody) SetMessage(v string) *GetLlmModelProvidersResponseBody {
	s.Message = &v
	return s
}

func (s *GetLlmModelProvidersResponseBody) SetRequestId(v string) *GetLlmModelProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLlmModelProvidersResponseBody) SetSuccess(v bool) *GetLlmModelProvidersResponseBody {
	s.Success = &v
	return s
}

func (s *GetLlmModelProvidersResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLlmModelProvidersResponseBodyData struct {
	// example:
	//
	// https://dashscope.aliyuncs.com
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 1
	Id        *int64                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	LlmModels []*GetLlmModelProvidersResponseBodyDataLlmModels `json:"LlmModels,omitempty" xml:"LlmModels,omitempty" type:"Repeated"`
	// example:
	//
	// BUILTIN_MODEL
	ProviderSource *string `json:"ProviderSource,omitempty" xml:"ProviderSource,omitempty"`
	// example:
	//
	// BAILIAN
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// example:
	//
	// bailian
	ServiceProvider *string `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
}

func (s GetLlmModelProvidersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLlmModelProvidersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLlmModelProvidersResponseBodyData) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *GetLlmModelProvidersResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetLlmModelProvidersResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetLlmModelProvidersResponseBodyData) GetLlmModels() []*GetLlmModelProvidersResponseBodyDataLlmModels {
	return s.LlmModels
}

func (s *GetLlmModelProvidersResponseBodyData) GetProviderSource() *string {
	return s.ProviderSource
}

func (s *GetLlmModelProvidersResponseBodyData) GetProviderType() *string {
	return s.ProviderType
}

func (s *GetLlmModelProvidersResponseBodyData) GetServiceProvider() *string {
	return s.ServiceProvider
}

func (s *GetLlmModelProvidersResponseBodyData) SetBaseUrl(v string) *GetLlmModelProvidersResponseBodyData {
	s.BaseUrl = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyData) SetEnabled(v bool) *GetLlmModelProvidersResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyData) SetId(v int64) *GetLlmModelProvidersResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyData) SetLlmModels(v []*GetLlmModelProvidersResponseBodyDataLlmModels) *GetLlmModelProvidersResponseBodyData {
	s.LlmModels = v
	return s
}

func (s *GetLlmModelProvidersResponseBodyData) SetProviderSource(v string) *GetLlmModelProvidersResponseBodyData {
	s.ProviderSource = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyData) SetProviderType(v string) *GetLlmModelProvidersResponseBodyData {
	s.ProviderType = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyData) SetServiceProvider(v string) *GetLlmModelProvidersResponseBodyData {
	s.ServiceProvider = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyData) Validate() error {
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

type GetLlmModelProvidersResponseBodyDataLlmModels struct {
	// example:
	//
	// 通义千问-Max
	CnName *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	// example:
	//
	// 通用文本生成模型
	Description         *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	EmbeddingDimensions []*int32 `json:"EmbeddingDimensions,omitempty" xml:"EmbeddingDimensions,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// DASH_SCOPE
	InvokeType *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	// example:
	//
	// 1001
	ModelId    *int64    `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelTypes []*string `json:"ModelTypes,omitempty" xml:"ModelTypes,omitempty" type:"Repeated"`
	// example:
	//
	// qwen-max
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// bailian
	ServiceProvider *string   `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
	Tasks           []*string `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s GetLlmModelProvidersResponseBodyDataLlmModels) String() string {
	return dara.Prettify(s)
}

func (s GetLlmModelProvidersResponseBodyDataLlmModels) GoString() string {
	return s.String()
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) GetCnName() *string {
	return s.CnName
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) GetDescription() *string {
	return s.Description
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) GetEmbeddingDimensions() []*int32 {
	return s.EmbeddingDimensions
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) GetInvokeType() *string {
	return s.InvokeType
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) GetModelId() *int64 {
	return s.ModelId
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) GetModelTypes() []*string {
	return s.ModelTypes
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) GetName() *string {
	return s.Name
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) GetServiceProvider() *string {
	return s.ServiceProvider
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) GetTasks() []*string {
	return s.Tasks
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) SetCnName(v string) *GetLlmModelProvidersResponseBodyDataLlmModels {
	s.CnName = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) SetDescription(v string) *GetLlmModelProvidersResponseBodyDataLlmModels {
	s.Description = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) SetEmbeddingDimensions(v []*int32) *GetLlmModelProvidersResponseBodyDataLlmModels {
	s.EmbeddingDimensions = v
	return s
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) SetEnabled(v bool) *GetLlmModelProvidersResponseBodyDataLlmModels {
	s.Enabled = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) SetInvokeType(v string) *GetLlmModelProvidersResponseBodyDataLlmModels {
	s.InvokeType = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) SetModelId(v int64) *GetLlmModelProvidersResponseBodyDataLlmModels {
	s.ModelId = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) SetModelTypes(v []*string) *GetLlmModelProvidersResponseBodyDataLlmModels {
	s.ModelTypes = v
	return s
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) SetName(v string) *GetLlmModelProvidersResponseBodyDataLlmModels {
	s.Name = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) SetServiceProvider(v string) *GetLlmModelProvidersResponseBodyDataLlmModels {
	s.ServiceProvider = &v
	return s
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) SetTasks(v []*string) *GetLlmModelProvidersResponseBodyDataLlmModels {
	s.Tasks = v
	return s
}

func (s *GetLlmModelProvidersResponseBodyDataLlmModels) Validate() error {
	return dara.Validate(s)
}
