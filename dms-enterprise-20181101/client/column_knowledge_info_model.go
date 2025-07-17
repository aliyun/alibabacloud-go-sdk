// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iColumnKnowledgeInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAssetDescription(v string) *ColumnKnowledgeInfo
	GetAssetDescription() *string
	SetAssetModifiedGmt(v string) *ColumnKnowledgeInfo
	GetAssetModifiedGmt() *string
	SetColumnName(v string) *ColumnKnowledgeInfo
	GetColumnName() *string
	SetColumnType(v string) *ColumnKnowledgeInfo
	GetColumnType() *string
	SetDescription(v string) *ColumnKnowledgeInfo
	GetDescription() *string
	SetPosition(v int32) *ColumnKnowledgeInfo
	GetPosition() *int32
}

type ColumnKnowledgeInfo struct {
	AssetDescription *string `json:"AssetDescription,omitempty" xml:"AssetDescription,omitempty"`
	AssetModifiedGmt *string `json:"AssetModifiedGmt,omitempty" xml:"AssetModifiedGmt,omitempty"`
	ColumnName       *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	ColumnType       *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Position         *int32  `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ColumnKnowledgeInfo) String() string {
	return dara.Prettify(s)
}

func (s ColumnKnowledgeInfo) GoString() string {
	return s.String()
}

func (s *ColumnKnowledgeInfo) GetAssetDescription() *string {
	return s.AssetDescription
}

func (s *ColumnKnowledgeInfo) GetAssetModifiedGmt() *string {
	return s.AssetModifiedGmt
}

func (s *ColumnKnowledgeInfo) GetColumnName() *string {
	return s.ColumnName
}

func (s *ColumnKnowledgeInfo) GetColumnType() *string {
	return s.ColumnType
}

func (s *ColumnKnowledgeInfo) GetDescription() *string {
	return s.Description
}

func (s *ColumnKnowledgeInfo) GetPosition() *int32 {
	return s.Position
}

func (s *ColumnKnowledgeInfo) SetAssetDescription(v string) *ColumnKnowledgeInfo {
	s.AssetDescription = &v
	return s
}

func (s *ColumnKnowledgeInfo) SetAssetModifiedGmt(v string) *ColumnKnowledgeInfo {
	s.AssetModifiedGmt = &v
	return s
}

func (s *ColumnKnowledgeInfo) SetColumnName(v string) *ColumnKnowledgeInfo {
	s.ColumnName = &v
	return s
}

func (s *ColumnKnowledgeInfo) SetColumnType(v string) *ColumnKnowledgeInfo {
	s.ColumnType = &v
	return s
}

func (s *ColumnKnowledgeInfo) SetDescription(v string) *ColumnKnowledgeInfo {
	s.Description = &v
	return s
}

func (s *ColumnKnowledgeInfo) SetPosition(v int32) *ColumnKnowledgeInfo {
	s.Position = &v
	return s
}

func (s *ColumnKnowledgeInfo) Validate() error {
	return dara.Validate(s)
}
