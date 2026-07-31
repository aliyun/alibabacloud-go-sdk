// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelGroupModelDTO interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ModelGroupModelDTO
	GetId() *int64
	SetModelCode(v string) *ModelGroupModelDTO
	GetModelCode() *string
	SetModelType(v string) *ModelGroupModelDTO
	GetModelType() *string
	SetName(v string) *ModelGroupModelDTO
	GetName() *string
	SetPlatform(v string) *ModelGroupModelDTO
	GetPlatform() *string
	SetVersion(v string) *ModelGroupModelDTO
	GetVersion() *string
}

type ModelGroupModelDTO struct {
	// example:
	//
	// 101
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// qwen-max
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// example:
	//
	// Chat
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// 通义千问-Max
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// qwen
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// example:
	//
	// v0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModelGroupModelDTO) String() string {
	return dara.Prettify(s)
}

func (s ModelGroupModelDTO) GoString() string {
	return s.String()
}

func (s *ModelGroupModelDTO) GetId() *int64 {
	return s.Id
}

func (s *ModelGroupModelDTO) GetModelCode() *string {
	return s.ModelCode
}

func (s *ModelGroupModelDTO) GetModelType() *string {
	return s.ModelType
}

func (s *ModelGroupModelDTO) GetName() *string {
	return s.Name
}

func (s *ModelGroupModelDTO) GetPlatform() *string {
	return s.Platform
}

func (s *ModelGroupModelDTO) GetVersion() *string {
	return s.Version
}

func (s *ModelGroupModelDTO) SetId(v int64) *ModelGroupModelDTO {
	s.Id = &v
	return s
}

func (s *ModelGroupModelDTO) SetModelCode(v string) *ModelGroupModelDTO {
	s.ModelCode = &v
	return s
}

func (s *ModelGroupModelDTO) SetModelType(v string) *ModelGroupModelDTO {
	s.ModelType = &v
	return s
}

func (s *ModelGroupModelDTO) SetName(v string) *ModelGroupModelDTO {
	s.Name = &v
	return s
}

func (s *ModelGroupModelDTO) SetPlatform(v string) *ModelGroupModelDTO {
	s.Platform = &v
	return s
}

func (s *ModelGroupModelDTO) SetVersion(v string) *ModelGroupModelDTO {
	s.Version = &v
	return s
}

func (s *ModelGroupModelDTO) Validate() error {
	return dara.Validate(s)
}
