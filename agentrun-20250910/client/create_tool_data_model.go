// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateToolData interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *CreateToolData
	GetCreatedAt() *string
	SetId(v string) *CreateToolData
	GetId() *string
	SetName(v string) *CreateToolData
	GetName() *string
	SetSourceType(v string) *CreateToolData
	GetSourceType() *string
	SetToolType(v string) *CreateToolData
	GetToolType() *string
}

type CreateToolData struct {
	CreatedAt  *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	ToolType   *string `json:"toolType,omitempty" xml:"toolType,omitempty"`
}

func (s CreateToolData) String() string {
	return dara.Prettify(s)
}

func (s CreateToolData) GoString() string {
	return s.String()
}

func (s *CreateToolData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateToolData) GetId() *string {
	return s.Id
}

func (s *CreateToolData) GetName() *string {
	return s.Name
}

func (s *CreateToolData) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateToolData) GetToolType() *string {
	return s.ToolType
}

func (s *CreateToolData) SetCreatedAt(v string) *CreateToolData {
	s.CreatedAt = &v
	return s
}

func (s *CreateToolData) SetId(v string) *CreateToolData {
	s.Id = &v
	return s
}

func (s *CreateToolData) SetName(v string) *CreateToolData {
	s.Name = &v
	return s
}

func (s *CreateToolData) SetSourceType(v string) *CreateToolData {
	s.SourceType = &v
	return s
}

func (s *CreateToolData) SetToolType(v string) *CreateToolData {
	s.ToolType = &v
	return s
}

func (s *CreateToolData) Validate() error {
	return dara.Validate(s)
}
