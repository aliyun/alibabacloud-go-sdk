// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateToolData interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateToolData
	GetId() *string
	SetName(v string) *UpdateToolData
	GetName() *string
	SetSourceType(v string) *UpdateToolData
	GetSourceType() *string
	SetToolType(v string) *UpdateToolData
	GetToolType() *string
	SetUpdatedAt(v string) *UpdateToolData
	GetUpdatedAt() *string
}

type UpdateToolData struct {
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	ToolType   *string `json:"toolType,omitempty" xml:"toolType,omitempty"`
	UpdatedAt  *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s UpdateToolData) String() string {
	return dara.Prettify(s)
}

func (s UpdateToolData) GoString() string {
	return s.String()
}

func (s *UpdateToolData) GetId() *string {
	return s.Id
}

func (s *UpdateToolData) GetName() *string {
	return s.Name
}

func (s *UpdateToolData) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateToolData) GetToolType() *string {
	return s.ToolType
}

func (s *UpdateToolData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateToolData) SetId(v string) *UpdateToolData {
	s.Id = &v
	return s
}

func (s *UpdateToolData) SetName(v string) *UpdateToolData {
	s.Name = &v
	return s
}

func (s *UpdateToolData) SetSourceType(v string) *UpdateToolData {
	s.SourceType = &v
	return s
}

func (s *UpdateToolData) SetToolType(v string) *UpdateToolData {
	s.ToolType = &v
	return s
}

func (s *UpdateToolData) SetUpdatedAt(v string) *UpdateToolData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateToolData) Validate() error {
	return dara.Validate(s)
}
