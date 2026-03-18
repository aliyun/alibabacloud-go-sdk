// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *ModelRouterUpdateModelRequest
	GetApiKey() *string
	SetBaseUrl(v string) *ModelRouterUpdateModelRequest
	GetBaseUrl() *string
	SetDescription(v string) *ModelRouterUpdateModelRequest
	GetDescription() *string
	SetMaxInputLength(v string) *ModelRouterUpdateModelRequest
	GetMaxInputLength() *string
	SetMaxOutputLength(v string) *ModelRouterUpdateModelRequest
	GetMaxOutputLength() *string
	SetModelId(v string) *ModelRouterUpdateModelRequest
	GetModelId() *string
	SetModelType(v string) *ModelRouterUpdateModelRequest
	GetModelType() *string
	SetName(v string) *ModelRouterUpdateModelRequest
	GetName() *string
	SetStatus(v int32) *ModelRouterUpdateModelRequest
	GetStatus() *int32
	SetSymbol(v string) *ModelRouterUpdateModelRequest
	GetSymbol() *string
	SetTags(v string) *ModelRouterUpdateModelRequest
	GetTags() *string
}

type ModelRouterUpdateModelRequest struct {
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
	// 8192
	MaxInputLength *string `json:"maxInputLength,omitempty" xml:"maxInputLength,omitempty"`
	// example:
	//
	// 2048
	MaxOutputLength *string `json:"maxOutputLength,omitempty" xml:"maxOutputLength,omitempty"`
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
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// alibaba
	Symbol *string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	// example:
	//
	// chat,NLP
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ModelRouterUpdateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateModelRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateModelRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *ModelRouterUpdateModelRequest) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ModelRouterUpdateModelRequest) GetDescription() *string {
	return s.Description
}

func (s *ModelRouterUpdateModelRequest) GetMaxInputLength() *string {
	return s.MaxInputLength
}

func (s *ModelRouterUpdateModelRequest) GetMaxOutputLength() *string {
	return s.MaxOutputLength
}

func (s *ModelRouterUpdateModelRequest) GetModelId() *string {
	return s.ModelId
}

func (s *ModelRouterUpdateModelRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ModelRouterUpdateModelRequest) GetName() *string {
	return s.Name
}

func (s *ModelRouterUpdateModelRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModelRouterUpdateModelRequest) GetSymbol() *string {
	return s.Symbol
}

func (s *ModelRouterUpdateModelRequest) GetTags() *string {
	return s.Tags
}

func (s *ModelRouterUpdateModelRequest) SetApiKey(v string) *ModelRouterUpdateModelRequest {
	s.ApiKey = &v
	return s
}

func (s *ModelRouterUpdateModelRequest) SetBaseUrl(v string) *ModelRouterUpdateModelRequest {
	s.BaseUrl = &v
	return s
}

func (s *ModelRouterUpdateModelRequest) SetDescription(v string) *ModelRouterUpdateModelRequest {
	s.Description = &v
	return s
}

func (s *ModelRouterUpdateModelRequest) SetMaxInputLength(v string) *ModelRouterUpdateModelRequest {
	s.MaxInputLength = &v
	return s
}

func (s *ModelRouterUpdateModelRequest) SetMaxOutputLength(v string) *ModelRouterUpdateModelRequest {
	s.MaxOutputLength = &v
	return s
}

func (s *ModelRouterUpdateModelRequest) SetModelId(v string) *ModelRouterUpdateModelRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterUpdateModelRequest) SetModelType(v string) *ModelRouterUpdateModelRequest {
	s.ModelType = &v
	return s
}

func (s *ModelRouterUpdateModelRequest) SetName(v string) *ModelRouterUpdateModelRequest {
	s.Name = &v
	return s
}

func (s *ModelRouterUpdateModelRequest) SetStatus(v int32) *ModelRouterUpdateModelRequest {
	s.Status = &v
	return s
}

func (s *ModelRouterUpdateModelRequest) SetSymbol(v string) *ModelRouterUpdateModelRequest {
	s.Symbol = &v
	return s
}

func (s *ModelRouterUpdateModelRequest) SetTags(v string) *ModelRouterUpdateModelRequest {
	s.Tags = &v
	return s
}

func (s *ModelRouterUpdateModelRequest) Validate() error {
	return dara.Validate(s)
}
