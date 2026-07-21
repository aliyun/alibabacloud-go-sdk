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
	SetExtensions(v string) *ModelDTO
	GetExtensions() *string
	SetGmtCreate(v string) *ModelDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *ModelDTO
	GetGmtModified() *string
	SetHasBillingRule(v bool) *ModelDTO
	GetHasBillingRule() *bool
	SetId(v int64) *ModelDTO
	GetId() *int64
	SetInOut(v string) *ModelDTO
	GetInOut() *string
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
	SetVersion(v int32) *ModelDTO
	GetVersion() *int32
}

type ModelDTO struct {
	// A masked preview of the API key.
	//
	// example:
	//
	// sk-xxx****xxx
	ApiKeyPreview *string `json:"apiKeyPreview,omitempty" xml:"apiKeyPreview,omitempty"`
	// The base URL for API requests.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com
	BaseUrl *string `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	// Indicates the model\\"s status. A value of 0 means enabled, and a non-zero value means disabled.
	//
	// example:
	//
	// 0
	DeleteTag *int32 `json:"deleteTag,omitempty" xml:"deleteTag,omitempty"`
	// The model description.
	//
	// example:
	//
	// 通义千问大模型
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Extensions  *string `json:"extensions,omitempty" xml:"extensions,omitempty"`
	// The time when the model was created, in ISO 8601 format.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the model was last updated, in ISO 8601 format.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// false
	HasBillingRule *bool `json:"hasBillingRule,omitempty" xml:"hasBillingRule,omitempty"`
	// The unique ID of the model.
	//
	// example:
	//
	// 1
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	InOut *string `json:"inOut,omitempty" xml:"inOut,omitempty"`
	// Indicates whether the model is custom.
	//
	// example:
	//
	// false
	IsCustom *bool `json:"isCustom,omitempty" xml:"isCustom,omitempty"`
	// The maximum input length.
	//
	// example:
	//
	// 8192
	MaxInputLength *string `json:"maxInputLength,omitempty" xml:"maxInputLength,omitempty"`
	// The maximum output length.
	//
	// example:
	//
	// 2048
	MaxOutputLength *string `json:"maxOutputLength,omitempty" xml:"maxOutputLength,omitempty"`
	// The model code.
	//
	// example:
	//
	// qwen-turbo
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// The model type.
	//
	// example:
	//
	// Chat
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The model name.
	//
	// example:
	//
	// 通义千问
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The vendor symbol.
	//
	// example:
	//
	// alibaba
	Symbol *string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	// The display names for the tags, separated by commas.
	//
	// example:
	//
	// 对话,自然语言处理
	TagNames *string `json:"tagNames,omitempty" xml:"tagNames,omitempty"`
	// A comma-separated list of model tags.
	//
	// example:
	//
	// chat,NLP
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The version number.
	//
	// example:
	//
	// 0
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
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

func (s *ModelDTO) GetExtensions() *string {
	return s.Extensions
}

func (s *ModelDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ModelDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ModelDTO) GetHasBillingRule() *bool {
	return s.HasBillingRule
}

func (s *ModelDTO) GetId() *int64 {
	return s.Id
}

func (s *ModelDTO) GetInOut() *string {
	return s.InOut
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

func (s *ModelDTO) GetVersion() *int32 {
	return s.Version
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

func (s *ModelDTO) SetExtensions(v string) *ModelDTO {
	s.Extensions = &v
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

func (s *ModelDTO) SetHasBillingRule(v bool) *ModelDTO {
	s.HasBillingRule = &v
	return s
}

func (s *ModelDTO) SetId(v int64) *ModelDTO {
	s.Id = &v
	return s
}

func (s *ModelDTO) SetInOut(v string) *ModelDTO {
	s.InOut = &v
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

func (s *ModelDTO) SetVersion(v int32) *ModelDTO {
	s.Version = &v
	return s
}

func (s *ModelDTO) Validate() error {
	return dara.Validate(s)
}
