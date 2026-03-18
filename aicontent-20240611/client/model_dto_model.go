// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelDTO interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyPreview(v string) *ModelDTO
	GetApiKeyPreview() *string
	SetBaseUrl(v string) *ModelDTO
	GetBaseUrl() *string
	SetDeleteTag(v int32) *ModelDTO
	GetDeleteTag() *int32
	SetDescription(v string) *ModelDTO
	GetDescription() *string
	SetGmtCreate(v string) *ModelDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *ModelDTO
	GetGmtModified() *string
	SetId(v int64) *ModelDTO
	GetId() *int64
	SetIsCustom(v bool) *ModelDTO
	GetIsCustom() *bool
	SetMaxInputLength(v string) *ModelDTO
	GetMaxInputLength() *string
	SetMaxOutputLength(v string) *ModelDTO
	GetMaxOutputLength() *string
	SetModelCode(v string) *ModelDTO
	GetModelCode() *string
	SetModelType(v string) *ModelDTO
	GetModelType() *string
	SetName(v string) *ModelDTO
	GetName() *string
	SetSymbol(v string) *ModelDTO
	GetSymbol() *string
	SetTagNames(v string) *ModelDTO
	GetTagNames() *string
	SetTags(v string) *ModelDTO
	GetTags() *string
}

type ModelDTO struct {
	// example:
	//
	// sk-xxx****xxx
	ApiKeyPreview *string `json:"apiKeyPreview,omitempty" xml:"apiKeyPreview,omitempty"`
	// example:
	//
	// https://dashscope.aliyuncs.com
	BaseUrl *string `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	// example:
	//
	// 0
	DeleteTag *int32 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// example:
	//
	// 通义千问大模型
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// false
	IsCustom *bool `json:"isCustom,omitempty" xml:"isCustom,omitempty"`
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
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
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
	// 对话,自然语言处理
	TagNames *string `json:"tagNames,omitempty" xml:"tagNames,omitempty"`
	// example:
	//
	// chat,NLP
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ModelDTO) String() string {
	return dara.Prettify(s)
}

func (s ModelDTO) GoString() string {
	return s.String()
}

func (s *ModelDTO) GetApiKeyPreview() *string {
	return s.ApiKeyPreview
}

func (s *ModelDTO) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ModelDTO) GetDeleteTag() *int32 {
	return s.DeleteTag
}

func (s *ModelDTO) GetDescription() *string {
	return s.Description
}

func (s *ModelDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ModelDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ModelDTO) GetId() *int64 {
	return s.Id
}

func (s *ModelDTO) GetIsCustom() *bool {
	return s.IsCustom
}

func (s *ModelDTO) GetMaxInputLength() *string {
	return s.MaxInputLength
}

func (s *ModelDTO) GetMaxOutputLength() *string {
	return s.MaxOutputLength
}

func (s *ModelDTO) GetModelCode() *string {
	return s.ModelCode
}

func (s *ModelDTO) GetModelType() *string {
	return s.ModelType
}

func (s *ModelDTO) GetName() *string {
	return s.Name
}

func (s *ModelDTO) GetSymbol() *string {
	return s.Symbol
}

func (s *ModelDTO) GetTagNames() *string {
	return s.TagNames
}

func (s *ModelDTO) GetTags() *string {
	return s.Tags
}

func (s *ModelDTO) SetApiKeyPreview(v string) *ModelDTO {
	s.ApiKeyPreview = &v
	return s
}

func (s *ModelDTO) SetBaseUrl(v string) *ModelDTO {
	s.BaseUrl = &v
	return s
}

func (s *ModelDTO) SetDeleteTag(v int32) *ModelDTO {
	s.DeleteTag = &v
	return s
}

func (s *ModelDTO) SetDescription(v string) *ModelDTO {
	s.Description = &v
	return s
}

func (s *ModelDTO) SetGmtCreate(v string) *ModelDTO {
	s.GmtCreate = &v
	return s
}

func (s *ModelDTO) SetGmtModified(v string) *ModelDTO {
	s.GmtModified = &v
	return s
}

func (s *ModelDTO) SetId(v int64) *ModelDTO {
	s.Id = &v
	return s
}

func (s *ModelDTO) SetIsCustom(v bool) *ModelDTO {
	s.IsCustom = &v
	return s
}

func (s *ModelDTO) SetMaxInputLength(v string) *ModelDTO {
	s.MaxInputLength = &v
	return s
}

func (s *ModelDTO) SetMaxOutputLength(v string) *ModelDTO {
	s.MaxOutputLength = &v
	return s
}

func (s *ModelDTO) SetModelCode(v string) *ModelDTO {
	s.ModelCode = &v
	return s
}

func (s *ModelDTO) SetModelType(v string) *ModelDTO {
	s.ModelType = &v
	return s
}

func (s *ModelDTO) SetName(v string) *ModelDTO {
	s.Name = &v
	return s
}

func (s *ModelDTO) SetSymbol(v string) *ModelDTO {
	s.Symbol = &v
	return s
}

func (s *ModelDTO) SetTagNames(v string) *ModelDTO {
	s.TagNames = &v
	return s
}

func (s *ModelDTO) SetTags(v string) *ModelDTO {
	s.Tags = &v
	return s
}

func (s *ModelDTO) Validate() error {
	return dara.Validate(s)
}
