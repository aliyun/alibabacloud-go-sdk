// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *ModelUpdateCmd
	GetApiKey() *string
	SetBaseUrl(v string) *ModelUpdateCmd
	GetBaseUrl() *string
	SetDescription(v string) *ModelUpdateCmd
	GetDescription() *string
	SetMaxInputLength(v string) *ModelUpdateCmd
	GetMaxInputLength() *string
	SetMaxOutputLength(v string) *ModelUpdateCmd
	GetMaxOutputLength() *string
	SetModelId(v string) *ModelUpdateCmd
	GetModelId() *string
	SetModelType(v string) *ModelUpdateCmd
	GetModelType() *string
	SetName(v string) *ModelUpdateCmd
	GetName() *string
	SetStatus(v int32) *ModelUpdateCmd
	GetStatus() *int32
	SetSymbol(v string) *ModelUpdateCmd
	GetSymbol() *string
	SetTags(v string) *ModelUpdateCmd
	GetTags() *string
}

type ModelUpdateCmd struct {
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

func (s ModelUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s ModelUpdateCmd) GoString() string {
	return s.String()
}

func (s *ModelUpdateCmd) GetApiKey() *string {
	return s.ApiKey
}

func (s *ModelUpdateCmd) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ModelUpdateCmd) GetDescription() *string {
	return s.Description
}

func (s *ModelUpdateCmd) GetMaxInputLength() *string {
	return s.MaxInputLength
}

func (s *ModelUpdateCmd) GetMaxOutputLength() *string {
	return s.MaxOutputLength
}

func (s *ModelUpdateCmd) GetModelId() *string {
	return s.ModelId
}

func (s *ModelUpdateCmd) GetModelType() *string {
	return s.ModelType
}

func (s *ModelUpdateCmd) GetName() *string {
	return s.Name
}

func (s *ModelUpdateCmd) GetStatus() *int32 {
	return s.Status
}

func (s *ModelUpdateCmd) GetSymbol() *string {
	return s.Symbol
}

func (s *ModelUpdateCmd) GetTags() *string {
	return s.Tags
}

func (s *ModelUpdateCmd) SetApiKey(v string) *ModelUpdateCmd {
	s.ApiKey = &v
	return s
}

func (s *ModelUpdateCmd) SetBaseUrl(v string) *ModelUpdateCmd {
	s.BaseUrl = &v
	return s
}

func (s *ModelUpdateCmd) SetDescription(v string) *ModelUpdateCmd {
	s.Description = &v
	return s
}

func (s *ModelUpdateCmd) SetMaxInputLength(v string) *ModelUpdateCmd {
	s.MaxInputLength = &v
	return s
}

func (s *ModelUpdateCmd) SetMaxOutputLength(v string) *ModelUpdateCmd {
	s.MaxOutputLength = &v
	return s
}

func (s *ModelUpdateCmd) SetModelId(v string) *ModelUpdateCmd {
	s.ModelId = &v
	return s
}

func (s *ModelUpdateCmd) SetModelType(v string) *ModelUpdateCmd {
	s.ModelType = &v
	return s
}

func (s *ModelUpdateCmd) SetName(v string) *ModelUpdateCmd {
	s.Name = &v
	return s
}

func (s *ModelUpdateCmd) SetStatus(v int32) *ModelUpdateCmd {
	s.Status = &v
	return s
}

func (s *ModelUpdateCmd) SetSymbol(v string) *ModelUpdateCmd {
	s.Symbol = &v
	return s
}

func (s *ModelUpdateCmd) SetTags(v string) *ModelUpdateCmd {
	s.Tags = &v
	return s
}

func (s *ModelUpdateCmd) Validate() error {
	return dara.Validate(s)
}
