// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateModelItemDTO interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *BatchCreateModelItemDTO
	GetDescription() *string
	SetExtensions(v string) *BatchCreateModelItemDTO
	GetExtensions() *string
	SetInOut(v string) *BatchCreateModelItemDTO
	GetInOut() *string
	SetMaxInputLength(v string) *BatchCreateModelItemDTO
	GetMaxInputLength() *string
	SetMaxOutputLength(v string) *BatchCreateModelItemDTO
	GetMaxOutputLength() *string
	SetModelId(v string) *BatchCreateModelItemDTO
	GetModelId() *string
	SetModelType(v string) *BatchCreateModelItemDTO
	GetModelType() *string
	SetName(v string) *BatchCreateModelItemDTO
	GetName() *string
	SetTags(v string) *BatchCreateModelItemDTO
	GetTags() *string
}

type BatchCreateModelItemDTO struct {
	// example:
	//
	// 通义千问 Max
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// {}
	Extensions *string `json:"extensions,omitempty" xml:"extensions,omitempty"`
	// example:
	//
	// text
	InOut *string `json:"inOut,omitempty" xml:"inOut,omitempty"`
	// example:
	//
	// 128000
	MaxInputLength *string `json:"maxInputLength,omitempty" xml:"maxInputLength,omitempty"`
	// example:
	//
	// 8192
	MaxOutputLength *string `json:"maxOutputLength,omitempty" xml:"maxOutputLength,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Chat
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// chat,NLP
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s BatchCreateModelItemDTO) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateModelItemDTO) GoString() string {
	return s.String()
}

func (s *BatchCreateModelItemDTO) GetDescription() *string {
	return s.Description
}

func (s *BatchCreateModelItemDTO) GetExtensions() *string {
	return s.Extensions
}

func (s *BatchCreateModelItemDTO) GetInOut() *string {
	return s.InOut
}

func (s *BatchCreateModelItemDTO) GetMaxInputLength() *string {
	return s.MaxInputLength
}

func (s *BatchCreateModelItemDTO) GetMaxOutputLength() *string {
	return s.MaxOutputLength
}

func (s *BatchCreateModelItemDTO) GetModelId() *string {
	return s.ModelId
}

func (s *BatchCreateModelItemDTO) GetModelType() *string {
	return s.ModelType
}

func (s *BatchCreateModelItemDTO) GetName() *string {
	return s.Name
}

func (s *BatchCreateModelItemDTO) GetTags() *string {
	return s.Tags
}

func (s *BatchCreateModelItemDTO) SetDescription(v string) *BatchCreateModelItemDTO {
	s.Description = &v
	return s
}

func (s *BatchCreateModelItemDTO) SetExtensions(v string) *BatchCreateModelItemDTO {
	s.Extensions = &v
	return s
}

func (s *BatchCreateModelItemDTO) SetInOut(v string) *BatchCreateModelItemDTO {
	s.InOut = &v
	return s
}

func (s *BatchCreateModelItemDTO) SetMaxInputLength(v string) *BatchCreateModelItemDTO {
	s.MaxInputLength = &v
	return s
}

func (s *BatchCreateModelItemDTO) SetMaxOutputLength(v string) *BatchCreateModelItemDTO {
	s.MaxOutputLength = &v
	return s
}

func (s *BatchCreateModelItemDTO) SetModelId(v string) *BatchCreateModelItemDTO {
	s.ModelId = &v
	return s
}

func (s *BatchCreateModelItemDTO) SetModelType(v string) *BatchCreateModelItemDTO {
	s.ModelType = &v
	return s
}

func (s *BatchCreateModelItemDTO) SetName(v string) *BatchCreateModelItemDTO {
	s.Name = &v
	return s
}

func (s *BatchCreateModelItemDTO) SetTags(v string) *BatchCreateModelItemDTO {
	s.Tags = &v
	return s
}

func (s *BatchCreateModelItemDTO) Validate() error {
	return dara.Validate(s)
}
