// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRowDTO interface {
	dara.Model
	String() string
	GoString() string
	SetModelCode(v string) *ModelRowDTO
	GetModelCode() *string
	SetModelId(v int64) *ModelRowDTO
	GetModelId() *int64
	SetModelName(v string) *ModelRowDTO
	GetModelName() *string
	SetValues(v string) *ModelRowDTO
	GetValues() *string
}

type ModelRowDTO struct {
	// example:
	//
	// qwen-plus
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 通义千问-Plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	Values    *string `json:"values,omitempty" xml:"values,omitempty"`
}

func (s ModelRowDTO) String() string {
	return dara.Prettify(s)
}

func (s ModelRowDTO) GoString() string {
	return s.String()
}

func (s *ModelRowDTO) GetModelCode() *string {
	return s.ModelCode
}

func (s *ModelRowDTO) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRowDTO) GetModelName() *string {
	return s.ModelName
}

func (s *ModelRowDTO) GetValues() *string {
	return s.Values
}

func (s *ModelRowDTO) SetModelCode(v string) *ModelRowDTO {
	s.ModelCode = &v
	return s
}

func (s *ModelRowDTO) SetModelId(v int64) *ModelRowDTO {
	s.ModelId = &v
	return s
}

func (s *ModelRowDTO) SetModelName(v string) *ModelRowDTO {
	s.ModelName = &v
	return s
}

func (s *ModelRowDTO) SetValues(v string) *ModelRowDTO {
	s.Values = &v
	return s
}

func (s *ModelRowDTO) Validate() error {
	return dara.Validate(s)
}
