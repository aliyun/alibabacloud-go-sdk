// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *ModelCreateCmd
	GetApiKey() *string
	SetBaseUrl(v string) *ModelCreateCmd
	GetBaseUrl() *string
	SetDescription(v string) *ModelCreateCmd
	GetDescription() *string
	SetMaxInputLength(v string) *ModelCreateCmd
	GetMaxInputLength() *string
	SetMaxOutputLength(v string) *ModelCreateCmd
	GetMaxOutputLength() *string
	SetModelId(v string) *ModelCreateCmd
	GetModelId() *string
	SetModelType(v string) *ModelCreateCmd
	GetModelType() *string
	SetName(v string) *ModelCreateCmd
	GetName() *string
	SetSymbol(v string) *ModelCreateCmd
	GetSymbol() *string
	SetTags(v string) *ModelCreateCmd
	GetTags() *string
}

type ModelCreateCmd struct {
	// example:
	//
	// sk-xxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
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
	// alibaba
	Symbol *string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	// example:
	//
	// chat,NLP
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ModelCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s ModelCreateCmd) GoString() string {
	return s.String()
}

func (s *ModelCreateCmd) GetApiKey() *string {
	return s.ApiKey
}

func (s *ModelCreateCmd) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ModelCreateCmd) GetDescription() *string {
	return s.Description
}

func (s *ModelCreateCmd) GetMaxInputLength() *string {
	return s.MaxInputLength
}

func (s *ModelCreateCmd) GetMaxOutputLength() *string {
	return s.MaxOutputLength
}

func (s *ModelCreateCmd) GetModelId() *string {
	return s.ModelId
}

func (s *ModelCreateCmd) GetModelType() *string {
	return s.ModelType
}

func (s *ModelCreateCmd) GetName() *string {
	return s.Name
}

func (s *ModelCreateCmd) GetSymbol() *string {
	return s.Symbol
}

func (s *ModelCreateCmd) GetTags() *string {
	return s.Tags
}

func (s *ModelCreateCmd) SetApiKey(v string) *ModelCreateCmd {
	s.ApiKey = &v
	return s
}

func (s *ModelCreateCmd) SetBaseUrl(v string) *ModelCreateCmd {
	s.BaseUrl = &v
	return s
}

func (s *ModelCreateCmd) SetDescription(v string) *ModelCreateCmd {
	s.Description = &v
	return s
}

func (s *ModelCreateCmd) SetMaxInputLength(v string) *ModelCreateCmd {
	s.MaxInputLength = &v
	return s
}

func (s *ModelCreateCmd) SetMaxOutputLength(v string) *ModelCreateCmd {
	s.MaxOutputLength = &v
	return s
}

func (s *ModelCreateCmd) SetModelId(v string) *ModelCreateCmd {
	s.ModelId = &v
	return s
}

func (s *ModelCreateCmd) SetModelType(v string) *ModelCreateCmd {
	s.ModelType = &v
	return s
}

func (s *ModelCreateCmd) SetName(v string) *ModelCreateCmd {
	s.Name = &v
	return s
}

func (s *ModelCreateCmd) SetSymbol(v string) *ModelCreateCmd {
	s.Symbol = &v
	return s
}

func (s *ModelCreateCmd) SetTags(v string) *ModelCreateCmd {
	s.Tags = &v
	return s
}

func (s *ModelCreateCmd) Validate() error {
	return dara.Validate(s)
}
