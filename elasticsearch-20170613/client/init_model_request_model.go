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
	// The API key for the AISearch open platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// os-xfdddf*
	ApiKey *string `json:"api_key,omitempty" xml:"api_key,omitempty"`
	// The service registration address of the AISearch open platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****.platform-cn-hangzhou-vpc.opensearch.aliyuncs.com
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The HTTP protocol type.
	//
	// This parameter is required.
	//
	// example:
	//
	// https
	HttpSchema *string `json:"http_schema,omitempty" xml:"http_schema,omitempty"`
	// The models to initialize. If this parameter is left empty, all built-in models are initialized.
	Models []*InitModelRequestModels `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	// The workspace.
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
	// The model type. Valid values:
	//
	// - text_embedding: text embedding
	//
	// - rerank: reranking
	//
	// - doc_split: document splitting
	//
	// - completion: large language model service
	//
	// - query_analyze: query analysis
	//
	// - doc_analyze: document content parsing
	//
	// - img_analyze: image content parsing
	//
	// example:
	//
	// text_embedding
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The service ID of the AISearch open platform.
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
