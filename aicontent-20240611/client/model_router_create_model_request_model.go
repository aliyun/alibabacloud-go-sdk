// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *ModelRouterCreateModelRequest
	GetApiKey() *string
	SetBaseUrl(v string) *ModelRouterCreateModelRequest
	GetBaseUrl() *string
	SetDescription(v string) *ModelRouterCreateModelRequest
	GetDescription() *string
	SetModelId(v string) *ModelRouterCreateModelRequest
	GetModelId() *string
	SetModelType(v string) *ModelRouterCreateModelRequest
	GetModelType() *string
	SetName(v string) *ModelRouterCreateModelRequest
	GetName() *string
	SetSymbol(v string) *ModelRouterCreateModelRequest
	GetSymbol() *string
	SetTags(v string) *ModelRouterCreateModelRequest
	GetTags() *string
}

type ModelRouterCreateModelRequest struct {
	// API Key
	//
	// example:
	//
	// sk-xxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// Base URL
	//
	// example:
	//
	// https://dashscope.aliyuncs.com
	BaseUrl *string `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	// example:
	//
	// 通义千问大模型
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// qwen-turbo
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// Chat
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// 通义千问
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// alibaba
	Symbol *string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	// example:
	//
	// TXT_GEN,DEEP_THINK
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ModelRouterCreateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateModelRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateModelRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *ModelRouterCreateModelRequest) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ModelRouterCreateModelRequest) GetDescription() *string {
	return s.Description
}

func (s *ModelRouterCreateModelRequest) GetModelId() *string {
	return s.ModelId
}

func (s *ModelRouterCreateModelRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ModelRouterCreateModelRequest) GetName() *string {
	return s.Name
}

func (s *ModelRouterCreateModelRequest) GetSymbol() *string {
	return s.Symbol
}

func (s *ModelRouterCreateModelRequest) GetTags() *string {
	return s.Tags
}

func (s *ModelRouterCreateModelRequest) SetApiKey(v string) *ModelRouterCreateModelRequest {
	s.ApiKey = &v
	return s
}

func (s *ModelRouterCreateModelRequest) SetBaseUrl(v string) *ModelRouterCreateModelRequest {
	s.BaseUrl = &v
	return s
}

func (s *ModelRouterCreateModelRequest) SetDescription(v string) *ModelRouterCreateModelRequest {
	s.Description = &v
	return s
}

func (s *ModelRouterCreateModelRequest) SetModelId(v string) *ModelRouterCreateModelRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterCreateModelRequest) SetModelType(v string) *ModelRouterCreateModelRequest {
	s.ModelType = &v
	return s
}

func (s *ModelRouterCreateModelRequest) SetName(v string) *ModelRouterCreateModelRequest {
	s.Name = &v
	return s
}

func (s *ModelRouterCreateModelRequest) SetSymbol(v string) *ModelRouterCreateModelRequest {
	s.Symbol = &v
	return s
}

func (s *ModelRouterCreateModelRequest) SetTags(v string) *ModelRouterCreateModelRequest {
	s.Tags = &v
	return s
}

func (s *ModelRouterCreateModelRequest) Validate() error {
	return dara.Validate(s)
}
