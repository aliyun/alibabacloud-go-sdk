// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableInstructionsVO interface {
	dara.Model
	String() string
	GoString() string
	SetAssetCreatedGmt(v string) *TableInstructionsVO
	GetAssetCreatedGmt() *string
	SetAssetDescription(v string) *TableInstructionsVO
	GetAssetDescription() *string
	SetAssetModifiedGmt(v string) *TableInstructionsVO
	GetAssetModifiedGmt() *string
	SetDbId(v int32) *TableInstructionsVO
	GetDbId() *int32
	SetDbType(v string) *TableInstructionsVO
	GetDbType() *string
	SetSummary(v string) *TableInstructionsVO
	GetSummary() *string
	SetTableName(v string) *TableInstructionsVO
	GetTableName() *string
}

type TableInstructionsVO struct {
	AssetCreatedGmt  *string `json:"AssetCreatedGmt,omitempty" xml:"AssetCreatedGmt,omitempty"`
	AssetDescription *string `json:"AssetDescription,omitempty" xml:"AssetDescription,omitempty"`
	AssetModifiedGmt *string `json:"AssetModifiedGmt,omitempty" xml:"AssetModifiedGmt,omitempty"`
	DbId             *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbType           *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	Summary          *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TableName        *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s TableInstructionsVO) String() string {
	return dara.Prettify(s)
}

func (s TableInstructionsVO) GoString() string {
	return s.String()
}

func (s *TableInstructionsVO) GetAssetCreatedGmt() *string {
	return s.AssetCreatedGmt
}

func (s *TableInstructionsVO) GetAssetDescription() *string {
	return s.AssetDescription
}

func (s *TableInstructionsVO) GetAssetModifiedGmt() *string {
	return s.AssetModifiedGmt
}

func (s *TableInstructionsVO) GetDbId() *int32 {
	return s.DbId
}

func (s *TableInstructionsVO) GetDbType() *string {
	return s.DbType
}

func (s *TableInstructionsVO) GetSummary() *string {
	return s.Summary
}

func (s *TableInstructionsVO) GetTableName() *string {
	return s.TableName
}

func (s *TableInstructionsVO) SetAssetCreatedGmt(v string) *TableInstructionsVO {
	s.AssetCreatedGmt = &v
	return s
}

func (s *TableInstructionsVO) SetAssetDescription(v string) *TableInstructionsVO {
	s.AssetDescription = &v
	return s
}

func (s *TableInstructionsVO) SetAssetModifiedGmt(v string) *TableInstructionsVO {
	s.AssetModifiedGmt = &v
	return s
}

func (s *TableInstructionsVO) SetDbId(v int32) *TableInstructionsVO {
	s.DbId = &v
	return s
}

func (s *TableInstructionsVO) SetDbType(v string) *TableInstructionsVO {
	s.DbType = &v
	return s
}

func (s *TableInstructionsVO) SetSummary(v string) *TableInstructionsVO {
	s.Summary = &v
	return s
}

func (s *TableInstructionsVO) SetTableName(v string) *TableInstructionsVO {
	s.TableName = &v
	return s
}

func (s *TableInstructionsVO) Validate() error {
	return dara.Validate(s)
}
