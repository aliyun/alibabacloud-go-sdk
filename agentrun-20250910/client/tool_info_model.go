// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToolInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCAPConfig(v *CAPConfig) *ToolInfo
	GetCAPConfig() *CAPConfig
	SetCreatedAt(v string) *ToolInfo
	GetCreatedAt() *string
	SetDescription(v string) *ToolInfo
	GetDescription() *string
	SetId(v string) *ToolInfo
	GetId() *string
	SetName(v string) *ToolInfo
	GetName() *string
	SetSchema(v string) *ToolInfo
	GetSchema() *string
	SetSourceType(v string) *ToolInfo
	GetSourceType() *string
	SetToolType(v string) *ToolInfo
	GetToolType() *string
	SetUpdatedAt(v string) *ToolInfo
	GetUpdatedAt() *string
}

type ToolInfo struct {
	CAPConfig   *CAPConfig `json:"CAPConfig,omitempty" xml:"CAPConfig,omitempty"`
	CreatedAt   *string    `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description *string    `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string    `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string    `json:"name,omitempty" xml:"name,omitempty"`
	Schema      *string    `json:"schema,omitempty" xml:"schema,omitempty"`
	SourceType  *string    `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	ToolType    *string    `json:"toolType,omitempty" xml:"toolType,omitempty"`
	UpdatedAt   *string    `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ToolInfo) String() string {
	return dara.Prettify(s)
}

func (s ToolInfo) GoString() string {
	return s.String()
}

func (s *ToolInfo) GetCAPConfig() *CAPConfig {
	return s.CAPConfig
}

func (s *ToolInfo) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ToolInfo) GetDescription() *string {
	return s.Description
}

func (s *ToolInfo) GetId() *string {
	return s.Id
}

func (s *ToolInfo) GetName() *string {
	return s.Name
}

func (s *ToolInfo) GetSchema() *string {
	return s.Schema
}

func (s *ToolInfo) GetSourceType() *string {
	return s.SourceType
}

func (s *ToolInfo) GetToolType() *string {
	return s.ToolType
}

func (s *ToolInfo) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ToolInfo) SetCAPConfig(v *CAPConfig) *ToolInfo {
	s.CAPConfig = v
	return s
}

func (s *ToolInfo) SetCreatedAt(v string) *ToolInfo {
	s.CreatedAt = &v
	return s
}

func (s *ToolInfo) SetDescription(v string) *ToolInfo {
	s.Description = &v
	return s
}

func (s *ToolInfo) SetId(v string) *ToolInfo {
	s.Id = &v
	return s
}

func (s *ToolInfo) SetName(v string) *ToolInfo {
	s.Name = &v
	return s
}

func (s *ToolInfo) SetSchema(v string) *ToolInfo {
	s.Schema = &v
	return s
}

func (s *ToolInfo) SetSourceType(v string) *ToolInfo {
	s.SourceType = &v
	return s
}

func (s *ToolInfo) SetToolType(v string) *ToolInfo {
	s.ToolType = &v
	return s
}

func (s *ToolInfo) SetUpdatedAt(v string) *ToolInfo {
	s.UpdatedAt = &v
	return s
}

func (s *ToolInfo) Validate() error {
	if s.CAPConfig != nil {
		if err := s.CAPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
