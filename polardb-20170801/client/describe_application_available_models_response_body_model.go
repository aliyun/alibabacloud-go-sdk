// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAvailableModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsrModels(v []*DescribeApplicationAvailableModelsResponseBodyAsrModels) *DescribeApplicationAvailableModelsResponseBody
	GetAsrModels() []*DescribeApplicationAvailableModelsResponseBodyAsrModels
	SetDefaultEmbeddingModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel) *DescribeApplicationAvailableModelsResponseBody
	GetDefaultEmbeddingModel() *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel
	SetDefaultGraphLLMModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel) *DescribeApplicationAvailableModelsResponseBody
	GetDefaultGraphLLMModel() *DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel
	SetDefaultLLMModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel) *DescribeApplicationAvailableModelsResponseBody
	GetDefaultLLMModel() *DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel
	SetDefaultMemoryType(v string) *DescribeApplicationAvailableModelsResponseBody
	GetDefaultMemoryType() *string
	SetDefaultRerankModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel) *DescribeApplicationAvailableModelsResponseBody
	GetDefaultRerankModel() *DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel
	SetDefaultVLModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultVLModel) *DescribeApplicationAvailableModelsResponseBody
	GetDefaultVLModel() *DescribeApplicationAvailableModelsResponseBodyDefaultVLModel
	SetDefaultVisualEmbedderModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel) *DescribeApplicationAvailableModelsResponseBody
	GetDefaultVisualEmbedderModel() *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel
	SetDefaultVisualLlmModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel) *DescribeApplicationAvailableModelsResponseBody
	GetDefaultVisualLlmModel() *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel
	SetEmbeddingModels(v []*DescribeApplicationAvailableModelsResponseBodyEmbeddingModels) *DescribeApplicationAvailableModelsResponseBody
	GetEmbeddingModels() []*DescribeApplicationAvailableModelsResponseBodyEmbeddingModels
	SetGraphLlmModels(v []*DescribeApplicationAvailableModelsResponseBodyGraphLlmModels) *DescribeApplicationAvailableModelsResponseBody
	GetGraphLlmModels() []*DescribeApplicationAvailableModelsResponseBodyGraphLlmModels
	SetLlmModels(v []*DescribeApplicationAvailableModelsResponseBodyLlmModels) *DescribeApplicationAvailableModelsResponseBody
	GetLlmModels() []*DescribeApplicationAvailableModelsResponseBodyLlmModels
	SetMemoryTypes(v []*string) *DescribeApplicationAvailableModelsResponseBody
	GetMemoryTypes() []*string
	SetMultimodalSupported(v bool) *DescribeApplicationAvailableModelsResponseBody
	GetMultimodalSupported() *bool
	SetRequestId(v string) *DescribeApplicationAvailableModelsResponseBody
	GetRequestId() *string
	SetRerankModels(v []*DescribeApplicationAvailableModelsResponseBodyRerankModels) *DescribeApplicationAvailableModelsResponseBody
	GetRerankModels() []*DescribeApplicationAvailableModelsResponseBodyRerankModels
	SetVLModels(v []*DescribeApplicationAvailableModelsResponseBodyVLModels) *DescribeApplicationAvailableModelsResponseBody
	GetVLModels() []*DescribeApplicationAvailableModelsResponseBodyVLModels
	SetVisualEmbedderModels(v []*DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels) *DescribeApplicationAvailableModelsResponseBody
	GetVisualEmbedderModels() []*DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels
	SetVisualLlmModels(v []*DescribeApplicationAvailableModelsResponseBodyVisualLlmModels) *DescribeApplicationAvailableModelsResponseBody
	GetVisualLlmModels() []*DescribeApplicationAvailableModelsResponseBodyVisualLlmModels
}

type DescribeApplicationAvailableModelsResponseBody struct {
	// The list of speech recognition models.
	AsrModels []*DescribeApplicationAvailableModelsResponseBodyAsrModels `json:"AsrModels,omitempty" xml:"AsrModels,omitempty" type:"Repeated"`
	// The default text embedding model.
	DefaultEmbeddingModel *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel `json:"DefaultEmbeddingModel,omitempty" xml:"DefaultEmbeddingModel,omitempty" type:"Struct"`
	// The default graph large language model.
	DefaultGraphLLMModel *DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel `json:"DefaultGraphLLMModel,omitempty" xml:"DefaultGraphLLMModel,omitempty" type:"Struct"`
	// The default large language model.
	DefaultLLMModel *DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel `json:"DefaultLLMModel,omitempty" xml:"DefaultLLMModel,omitempty" type:"Struct"`
	// The default memory type.
	//
	// example:
	//
	// TEXT
	DefaultMemoryType *string `json:"DefaultMemoryType,omitempty" xml:"DefaultMemoryType,omitempty"`
	// The default reranking model.
	DefaultRerankModel *DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel `json:"DefaultRerankModel,omitempty" xml:"DefaultRerankModel,omitempty" type:"Struct"`
	// The default visual language model.
	DefaultVLModel *DescribeApplicationAvailableModelsResponseBodyDefaultVLModel `json:"DefaultVLModel,omitempty" xml:"DefaultVLModel,omitempty" type:"Struct"`
	// The default VisualEmbedder model.
	DefaultVisualEmbedderModel *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel `json:"DefaultVisualEmbedderModel,omitempty" xml:"DefaultVisualEmbedderModel,omitempty" type:"Struct"`
	// The default VisualLLM model.
	DefaultVisualLlmModel *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel `json:"DefaultVisualLlmModel,omitempty" xml:"DefaultVisualLlmModel,omitempty" type:"Struct"`
	// The list of available text embedding models.
	EmbeddingModels []*DescribeApplicationAvailableModelsResponseBodyEmbeddingModels `json:"EmbeddingModels,omitempty" xml:"EmbeddingModels,omitempty" type:"Repeated"`
	// The list of available graph large language models.
	GraphLlmModels []*DescribeApplicationAvailableModelsResponseBodyGraphLlmModels `json:"GraphLlmModels,omitempty" xml:"GraphLlmModels,omitempty" type:"Repeated"`
	// The list of available large language models.
	LlmModels []*DescribeApplicationAvailableModelsResponseBodyLlmModels `json:"LlmModels,omitempty" xml:"LlmModels,omitempty" type:"Repeated"`
	// The available memory types.
	MemoryTypes []*string `json:"MemoryTypes,omitempty" xml:"MemoryTypes,omitempty" type:"Repeated"`
	// Indicates whether the current account supports Mem0 multimodal.
	//
	// example:
	//
	// true
	MultimodalSupported *bool `json:"MultimodalSupported,omitempty" xml:"MultimodalSupported,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of available reranking models.
	RerankModels []*DescribeApplicationAvailableModelsResponseBodyRerankModels `json:"RerankModels,omitempty" xml:"RerankModels,omitempty" type:"Repeated"`
	// The list of available visual language models.
	VLModels []*DescribeApplicationAvailableModelsResponseBodyVLModels `json:"VLModels,omitempty" xml:"VLModels,omitempty" type:"Repeated"`
	// The list of available VisualEmbedder models.
	VisualEmbedderModels []*DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels `json:"VisualEmbedderModels,omitempty" xml:"VisualEmbedderModels,omitempty" type:"Repeated"`
	// The list of available VisualLLM models.
	VisualLlmModels []*DescribeApplicationAvailableModelsResponseBodyVisualLlmModels `json:"VisualLlmModels,omitempty" xml:"VisualLlmModels,omitempty" type:"Repeated"`
}

func (s DescribeApplicationAvailableModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetAsrModels() []*DescribeApplicationAvailableModelsResponseBodyAsrModels {
	return s.AsrModels
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetDefaultEmbeddingModel() *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel {
	return s.DefaultEmbeddingModel
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetDefaultGraphLLMModel() *DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel {
	return s.DefaultGraphLLMModel
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetDefaultLLMModel() *DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel {
	return s.DefaultLLMModel
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetDefaultMemoryType() *string {
	return s.DefaultMemoryType
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetDefaultRerankModel() *DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel {
	return s.DefaultRerankModel
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetDefaultVLModel() *DescribeApplicationAvailableModelsResponseBodyDefaultVLModel {
	return s.DefaultVLModel
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetDefaultVisualEmbedderModel() *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel {
	return s.DefaultVisualEmbedderModel
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetDefaultVisualLlmModel() *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel {
	return s.DefaultVisualLlmModel
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetEmbeddingModels() []*DescribeApplicationAvailableModelsResponseBodyEmbeddingModels {
	return s.EmbeddingModels
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetGraphLlmModels() []*DescribeApplicationAvailableModelsResponseBodyGraphLlmModels {
	return s.GraphLlmModels
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetLlmModels() []*DescribeApplicationAvailableModelsResponseBodyLlmModels {
	return s.LlmModels
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetMemoryTypes() []*string {
	return s.MemoryTypes
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetMultimodalSupported() *bool {
	return s.MultimodalSupported
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetRerankModels() []*DescribeApplicationAvailableModelsResponseBodyRerankModels {
	return s.RerankModels
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetVLModels() []*DescribeApplicationAvailableModelsResponseBodyVLModels {
	return s.VLModels
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetVisualEmbedderModels() []*DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels {
	return s.VisualEmbedderModels
}

func (s *DescribeApplicationAvailableModelsResponseBody) GetVisualLlmModels() []*DescribeApplicationAvailableModelsResponseBodyVisualLlmModels {
	return s.VisualLlmModels
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetAsrModels(v []*DescribeApplicationAvailableModelsResponseBodyAsrModels) *DescribeApplicationAvailableModelsResponseBody {
	s.AsrModels = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetDefaultEmbeddingModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel) *DescribeApplicationAvailableModelsResponseBody {
	s.DefaultEmbeddingModel = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetDefaultGraphLLMModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel) *DescribeApplicationAvailableModelsResponseBody {
	s.DefaultGraphLLMModel = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetDefaultLLMModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel) *DescribeApplicationAvailableModelsResponseBody {
	s.DefaultLLMModel = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetDefaultMemoryType(v string) *DescribeApplicationAvailableModelsResponseBody {
	s.DefaultMemoryType = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetDefaultRerankModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel) *DescribeApplicationAvailableModelsResponseBody {
	s.DefaultRerankModel = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetDefaultVLModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultVLModel) *DescribeApplicationAvailableModelsResponseBody {
	s.DefaultVLModel = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetDefaultVisualEmbedderModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel) *DescribeApplicationAvailableModelsResponseBody {
	s.DefaultVisualEmbedderModel = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetDefaultVisualLlmModel(v *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel) *DescribeApplicationAvailableModelsResponseBody {
	s.DefaultVisualLlmModel = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetEmbeddingModels(v []*DescribeApplicationAvailableModelsResponseBodyEmbeddingModels) *DescribeApplicationAvailableModelsResponseBody {
	s.EmbeddingModels = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetGraphLlmModels(v []*DescribeApplicationAvailableModelsResponseBodyGraphLlmModels) *DescribeApplicationAvailableModelsResponseBody {
	s.GraphLlmModels = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetLlmModels(v []*DescribeApplicationAvailableModelsResponseBodyLlmModels) *DescribeApplicationAvailableModelsResponseBody {
	s.LlmModels = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetMemoryTypes(v []*string) *DescribeApplicationAvailableModelsResponseBody {
	s.MemoryTypes = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetMultimodalSupported(v bool) *DescribeApplicationAvailableModelsResponseBody {
	s.MultimodalSupported = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetRequestId(v string) *DescribeApplicationAvailableModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetRerankModels(v []*DescribeApplicationAvailableModelsResponseBodyRerankModels) *DescribeApplicationAvailableModelsResponseBody {
	s.RerankModels = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetVLModels(v []*DescribeApplicationAvailableModelsResponseBodyVLModels) *DescribeApplicationAvailableModelsResponseBody {
	s.VLModels = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetVisualEmbedderModels(v []*DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels) *DescribeApplicationAvailableModelsResponseBody {
	s.VisualEmbedderModels = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) SetVisualLlmModels(v []*DescribeApplicationAvailableModelsResponseBodyVisualLlmModels) *DescribeApplicationAvailableModelsResponseBody {
	s.VisualLlmModels = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBody) Validate() error {
	if s.AsrModels != nil {
		for _, item := range s.AsrModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DefaultEmbeddingModel != nil {
		if err := s.DefaultEmbeddingModel.Validate(); err != nil {
			return err
		}
	}
	if s.DefaultGraphLLMModel != nil {
		if err := s.DefaultGraphLLMModel.Validate(); err != nil {
			return err
		}
	}
	if s.DefaultLLMModel != nil {
		if err := s.DefaultLLMModel.Validate(); err != nil {
			return err
		}
	}
	if s.DefaultRerankModel != nil {
		if err := s.DefaultRerankModel.Validate(); err != nil {
			return err
		}
	}
	if s.DefaultVLModel != nil {
		if err := s.DefaultVLModel.Validate(); err != nil {
			return err
		}
	}
	if s.DefaultVisualEmbedderModel != nil {
		if err := s.DefaultVisualEmbedderModel.Validate(); err != nil {
			return err
		}
	}
	if s.DefaultVisualLlmModel != nil {
		if err := s.DefaultVisualLlmModel.Validate(); err != nil {
			return err
		}
	}
	if s.EmbeddingModels != nil {
		for _, item := range s.EmbeddingModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GraphLlmModels != nil {
		for _, item := range s.GraphLlmModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LlmModels != nil {
		for _, item := range s.LlmModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RerankModels != nil {
		for _, item := range s.RerankModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VLModels != nil {
		for _, item := range s.VLModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VisualEmbedderModels != nil {
		for _, item := range s.VisualEmbedderModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VisualLlmModels != nil {
		for _, item := range s.VisualLlmModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationAvailableModelsResponseBodyAsrModels struct {
	// The model name.
	//
	// example:
	//
	// paraformer-realtime-v2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The protocol used by the speech recognition model.
	//
	// example:
	//
	// websocket
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyAsrModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyAsrModels) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyAsrModels) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyAsrModels) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeApplicationAvailableModelsResponseBodyAsrModels) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyAsrModels {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyAsrModels) SetProtocol(v string) *DescribeApplicationAvailableModelsResponseBodyAsrModels {
	s.Protocol = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyAsrModels) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel struct {
	// The default vector dimensions of the model.
	//
	// example:
	//
	// 1024
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The list of vector dimensions supported by the model.
	Dimensions []*int32 `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The model name.
	//
	// example:
	//
	// text-embedding-v3
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel) GetDimension() *int32 {
	return s.Dimension
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel) GetDimensions() []*int32 {
	return s.Dimensions
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel) SetDimension(v int32) *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel {
	s.Dimension = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel) SetDimensions(v []*int32) *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel {
	s.Dimensions = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultEmbeddingModel) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel struct {
	// The model name.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultGraphLLMModel) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel struct {
	// The model name.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultLLMModel) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel struct {
	// The model name.
	//
	// example:
	//
	// gte-rerank-v2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultRerankModel) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyDefaultVLModel struct {
	// The model name.
	//
	// example:
	//
	// qwen-vl-max
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultVLModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultVLModel) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVLModel) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVLModel) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyDefaultVLModel {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVLModel) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel struct {
	// The default dimension.
	//
	// example:
	//
	// 2560
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The available dimensions.
	Dimensions []*int32 `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The model name.
	//
	// example:
	//
	// qwen3.7-text-embedding
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel) GetDimension() *int32 {
	return s.Dimension
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel) GetDimensions() []*int32 {
	return s.Dimensions
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel) SetDimension(v int32) *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel {
	s.Dimension = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel) SetDimensions(v []*int32) *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel {
	s.Dimensions = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualEmbedderModel) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel struct {
	// The model name.
	//
	// example:
	//
	// qwen3.6-flash
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The supported modalities.
	SupportedModalities []*string `json:"SupportedModalities,omitempty" xml:"SupportedModalities,omitempty" type:"Repeated"`
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel) GetSupportedModalities() []*string {
	return s.SupportedModalities
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel) SetSupportedModalities(v []*string) *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel {
	s.SupportedModalities = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyDefaultVisualLlmModel) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyEmbeddingModels struct {
	// The default vector dimensions of the model.
	//
	// example:
	//
	// 1024
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The list of vector dimensions supported by the model.
	Dimensions []*int32 `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The model name.
	//
	// example:
	//
	// text-embedding-v3
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyEmbeddingModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyEmbeddingModels) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyEmbeddingModels) GetDimension() *int32 {
	return s.Dimension
}

func (s *DescribeApplicationAvailableModelsResponseBodyEmbeddingModels) GetDimensions() []*int32 {
	return s.Dimensions
}

func (s *DescribeApplicationAvailableModelsResponseBodyEmbeddingModels) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyEmbeddingModels) SetDimension(v int32) *DescribeApplicationAvailableModelsResponseBodyEmbeddingModels {
	s.Dimension = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyEmbeddingModels) SetDimensions(v []*int32) *DescribeApplicationAvailableModelsResponseBodyEmbeddingModels {
	s.Dimensions = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyEmbeddingModels) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyEmbeddingModels {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyEmbeddingModels) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyGraphLlmModels struct {
	// The model name.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyGraphLlmModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyGraphLlmModels) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyGraphLlmModels) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyGraphLlmModels) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyGraphLlmModels {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyGraphLlmModels) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyLlmModels struct {
	// The model name.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyLlmModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyLlmModels) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyLlmModels) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyLlmModels) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyLlmModels {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyLlmModels) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyRerankModels struct {
	// The model name.
	//
	// example:
	//
	// gte-rerank-v2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyRerankModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyRerankModels) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyRerankModels) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyRerankModels) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyRerankModels {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyRerankModels) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyVLModels struct {
	// The model name.
	//
	// example:
	//
	// qwen-vl-max
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyVLModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyVLModels) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyVLModels) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyVLModels) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyVLModels {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyVLModels) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels struct {
	// The default dimension.
	//
	// example:
	//
	// 2560
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The available dimensions.
	Dimensions []*int32 `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The model name.
	//
	// example:
	//
	// qwen3.7-text-embedding
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels) GetDimension() *int32 {
	return s.Dimension
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels) GetDimensions() []*int32 {
	return s.Dimensions
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels) SetDimension(v int32) *DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels {
	s.Dimension = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels) SetDimensions(v []*int32) *DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels {
	s.Dimensions = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualEmbedderModels) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationAvailableModelsResponseBodyVisualLlmModels struct {
	// The model name.
	//
	// example:
	//
	// qwen3.6-flash
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The supported modalities.
	SupportedModalities []*string `json:"SupportedModalities,omitempty" xml:"SupportedModalities,omitempty" type:"Repeated"`
}

func (s DescribeApplicationAvailableModelsResponseBodyVisualLlmModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAvailableModelsResponseBodyVisualLlmModels) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualLlmModels) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualLlmModels) GetSupportedModalities() []*string {
	return s.SupportedModalities
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualLlmModels) SetName(v string) *DescribeApplicationAvailableModelsResponseBodyVisualLlmModels {
	s.Name = &v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualLlmModels) SetSupportedModalities(v []*string) *DescribeApplicationAvailableModelsResponseBodyVisualLlmModels {
	s.SupportedModalities = v
	return s
}

func (s *DescribeApplicationAvailableModelsResponseBodyVisualLlmModels) Validate() error {
	return dara.Validate(s)
}
