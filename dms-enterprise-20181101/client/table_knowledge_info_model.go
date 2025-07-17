// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableKnowledgeInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAssetDescription(v string) *TableKnowledgeInfo
	GetAssetDescription() *string
	SetAssetModifiedGmt(v string) *TableKnowledgeInfo
	GetAssetModifiedGmt() *string
	SetColumnList(v []*ColumnKnowledgeInfo) *TableKnowledgeInfo
	GetColumnList() []*ColumnKnowledgeInfo
	SetDescription(v string) *TableKnowledgeInfo
	GetDescription() *string
	SetSummary(v string) *TableKnowledgeInfo
	GetSummary() *string
	SetTableName(v string) *TableKnowledgeInfo
	GetTableName() *string
}

type TableKnowledgeInfo struct {
	AssetDescription *string                `json:"AssetDescription,omitempty" xml:"AssetDescription,omitempty"`
	AssetModifiedGmt *string                `json:"AssetModifiedGmt,omitempty" xml:"AssetModifiedGmt,omitempty"`
	ColumnList       []*ColumnKnowledgeInfo `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
	Description      *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Summary          *string                `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TableName        *string                `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s TableKnowledgeInfo) String() string {
	return dara.Prettify(s)
}

func (s TableKnowledgeInfo) GoString() string {
	return s.String()
}

func (s *TableKnowledgeInfo) GetAssetDescription() *string {
	return s.AssetDescription
}

func (s *TableKnowledgeInfo) GetAssetModifiedGmt() *string {
	return s.AssetModifiedGmt
}

func (s *TableKnowledgeInfo) GetColumnList() []*ColumnKnowledgeInfo {
	return s.ColumnList
}

func (s *TableKnowledgeInfo) GetDescription() *string {
	return s.Description
}

func (s *TableKnowledgeInfo) GetSummary() *string {
	return s.Summary
}

func (s *TableKnowledgeInfo) GetTableName() *string {
	return s.TableName
}

func (s *TableKnowledgeInfo) SetAssetDescription(v string) *TableKnowledgeInfo {
	s.AssetDescription = &v
	return s
}

func (s *TableKnowledgeInfo) SetAssetModifiedGmt(v string) *TableKnowledgeInfo {
	s.AssetModifiedGmt = &v
	return s
}

func (s *TableKnowledgeInfo) SetColumnList(v []*ColumnKnowledgeInfo) *TableKnowledgeInfo {
	s.ColumnList = v
	return s
}

func (s *TableKnowledgeInfo) SetDescription(v string) *TableKnowledgeInfo {
	s.Description = &v
	return s
}

func (s *TableKnowledgeInfo) SetSummary(v string) *TableKnowledgeInfo {
	s.Summary = &v
	return s
}

func (s *TableKnowledgeInfo) SetTableName(v string) *TableKnowledgeInfo {
	s.TableName = &v
	return s
}

func (s *TableKnowledgeInfo) Validate() error {
	return dara.Validate(s)
}
