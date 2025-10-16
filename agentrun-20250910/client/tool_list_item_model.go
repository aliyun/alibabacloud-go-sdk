// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToolListItem interface {
	dara.Model
	String() string
	GoString() string
	SetCAPConfig(v *CAPConfig) *ToolListItem
	GetCAPConfig() *CAPConfig
	SetCreatedAt(v string) *ToolListItem
	GetCreatedAt() *string
	SetDescription(v string) *ToolListItem
	GetDescription() *string
	SetId(v string) *ToolListItem
	GetId() *string
	SetName(v string) *ToolListItem
	GetName() *string
	SetSchema(v string) *ToolListItem
	GetSchema() *string
	SetSourceType(v string) *ToolListItem
	GetSourceType() *string
	SetToolType(v string) *ToolListItem
	GetToolType() *string
	SetUpdatedAt(v string) *ToolListItem
	GetUpdatedAt() *string
}

type ToolListItem struct {
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

func (s ToolListItem) String() string {
	return dara.Prettify(s)
}

func (s ToolListItem) GoString() string {
	return s.String()
}

func (s *ToolListItem) GetCAPConfig() *CAPConfig {
	return s.CAPConfig
}

func (s *ToolListItem) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ToolListItem) GetDescription() *string {
	return s.Description
}

func (s *ToolListItem) GetId() *string {
	return s.Id
}

func (s *ToolListItem) GetName() *string {
	return s.Name
}

func (s *ToolListItem) GetSchema() *string {
	return s.Schema
}

func (s *ToolListItem) GetSourceType() *string {
	return s.SourceType
}

func (s *ToolListItem) GetToolType() *string {
	return s.ToolType
}

func (s *ToolListItem) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ToolListItem) SetCAPConfig(v *CAPConfig) *ToolListItem {
	s.CAPConfig = v
	return s
}

func (s *ToolListItem) SetCreatedAt(v string) *ToolListItem {
	s.CreatedAt = &v
	return s
}

func (s *ToolListItem) SetDescription(v string) *ToolListItem {
	s.Description = &v
	return s
}

func (s *ToolListItem) SetId(v string) *ToolListItem {
	s.Id = &v
	return s
}

func (s *ToolListItem) SetName(v string) *ToolListItem {
	s.Name = &v
	return s
}

func (s *ToolListItem) SetSchema(v string) *ToolListItem {
	s.Schema = &v
	return s
}

func (s *ToolListItem) SetSourceType(v string) *ToolListItem {
	s.SourceType = &v
	return s
}

func (s *ToolListItem) SetToolType(v string) *ToolListItem {
	s.ToolType = &v
	return s
}

func (s *ToolListItem) SetUpdatedAt(v string) *ToolListItem {
	s.UpdatedAt = &v
	return s
}

func (s *ToolListItem) Validate() error {
	if s.CAPConfig != nil {
		if err := s.CAPConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
