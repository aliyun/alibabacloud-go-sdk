// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *InitModelRequest
	GetApiKey() *string
	SetHost(v string) *InitModelRequest
	GetHost() *string
	SetHttpSchema(v string) *InitModelRequest
	GetHttpSchema() *string
	SetModels(v []*InitModelRequestModels) *InitModelRequest
	GetModels() []*InitModelRequestModels
	SetWorkspace(v string) *InitModelRequest
	GetWorkspace() *string
}

type InitModelRequest struct {
	// API key for the AI Search Open Platform
	//
	// This parameter is required.
	//
	// example:
	//
	// os-xfdddf*
	ApiKey *string `json:"api_key,omitempty" xml:"api_key,omitempty"`
	// Service registration endpoint of the AI Search Open Platform
	//
	// This parameter is required.
	//
	// example:
	//
	// ****.platform-cn-hangzhou-vpc.opensearch.aliyuncs.com
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// HTTP protocol type
	//
	// This parameter is required.
	//
	// example:
	//
	// https
	HttpSchema *string `json:"http_schema,omitempty" xml:"http_schema,omitempty"`
	// Specify the models to initialize. If empty, all built-in models will be initialized.
	Models []*InitModelRequestModels `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	// Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s InitModelRequest) String() string {
	return dara.Prettify(s)
}

func (s InitModelRequest) GoString() string {
	return s.String()
}

func (s *InitModelRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *InitModelRequest) GetHost() *string {
	return s.Host
}

func (s *InitModelRequest) GetHttpSchema() *string {
	return s.HttpSchema
}

func (s *InitModelRequest) GetModels() []*InitModelRequestModels {
	return s.Models
}

func (s *InitModelRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *InitModelRequest) SetApiKey(v string) *InitModelRequest {
	s.ApiKey = &v
	return s
}

func (s *InitModelRequest) SetHost(v string) *InitModelRequest {
	s.Host = &v
	return s
}

func (s *InitModelRequest) SetHttpSchema(v string) *InitModelRequest {
	s.HttpSchema = &v
	return s
}

func (s *InitModelRequest) SetModels(v []*InitModelRequestModels) *InitModelRequest {
	s.Models = v
	return s
}

func (s *InitModelRequest) SetWorkspace(v string) *InitModelRequest {
	s.Workspace = &v
	return s
}

func (s *InitModelRequest) Validate() error {
	if s.Models != nil {
		for _, item := range s.Models {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InitModelRequestModels struct {
	// Model category
	//
	// - a. Text embedding: text_embedding
	//
	// - b. Reranking: rerank
	//
	// - c. Document segmentation: doc_split
	//
	// - d. Large Language Model (LLM) service: completion
	//
	// - e. Query analysis: query_analyze
	//
	// - f. Document content parsing: doc_analyze
	//
	// - g. Image content parsing: img_analyze
	//
	// example:
	//
	// text_embedding
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// Service ID of the Search Open Platform
	//
	// example:
	//
	// ops-text-embedding-**
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
}

func (s InitModelRequestModels) String() string {
	return dara.Prettify(s)
}

func (s InitModelRequestModels) GoString() string {
	return s.String()
}

func (s *InitModelRequestModels) GetModelType() *string {
	return s.ModelType
}

func (s *InitModelRequestModels) GetServiceId() *string {
	return s.ServiceId
}

func (s *InitModelRequestModels) SetModelType(v string) *InitModelRequestModels {
	s.ModelType = &v
	return s
}

func (s *InitModelRequestModels) SetServiceId(v string) *InitModelRequestModels {
	s.ServiceId = &v
	return s
}

func (s *InitModelRequestModels) Validate() error {
	return dara.Validate(s)
}
