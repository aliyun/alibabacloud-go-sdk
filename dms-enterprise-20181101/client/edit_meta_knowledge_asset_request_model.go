// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditMetaKnowledgeAssetRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAssetDescription(v string) *EditMetaKnowledgeAssetRequest
  GetAssetDescription() *string 
  SetColumnName(v string) *EditMetaKnowledgeAssetRequest
  GetColumnName() *string 
  SetDbId(v int32) *EditMetaKnowledgeAssetRequest
  GetDbId() *int32 
  SetTableName(v string) *EditMetaKnowledgeAssetRequest
  GetTableName() *string 
  SetTableSchemaName(v string) *EditMetaKnowledgeAssetRequest
  GetTableSchemaName() *string 
}

type EditMetaKnowledgeAssetRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // test
  AssetDescription *string `json:"AssetDescription,omitempty" xml:"AssetDescription,omitempty"`
  // example:
  // 
  // test_column
  ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1930****
  DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // table_name
  TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
  // example:
  // 
  // dbo
  TableSchemaName *string `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
}

func (s EditMetaKnowledgeAssetRequest) String() string {
  return dara.Prettify(s)
}

func (s EditMetaKnowledgeAssetRequest) GoString() string {
  return s.String()
}

func (s *EditMetaKnowledgeAssetRequest) GetAssetDescription() *string  {
  return s.AssetDescription
}

func (s *EditMetaKnowledgeAssetRequest) GetColumnName() *string  {
  return s.ColumnName
}

func (s *EditMetaKnowledgeAssetRequest) GetDbId() *int32  {
  return s.DbId
}

func (s *EditMetaKnowledgeAssetRequest) GetTableName() *string  {
  return s.TableName
}

func (s *EditMetaKnowledgeAssetRequest) GetTableSchemaName() *string  {
  return s.TableSchemaName
}

func (s *EditMetaKnowledgeAssetRequest) SetAssetDescription(v string) *EditMetaKnowledgeAssetRequest {
  s.AssetDescription = &v
  return s
}

func (s *EditMetaKnowledgeAssetRequest) SetColumnName(v string) *EditMetaKnowledgeAssetRequest {
  s.ColumnName = &v
  return s
}

func (s *EditMetaKnowledgeAssetRequest) SetDbId(v int32) *EditMetaKnowledgeAssetRequest {
  s.DbId = &v
  return s
}

func (s *EditMetaKnowledgeAssetRequest) SetTableName(v string) *EditMetaKnowledgeAssetRequest {
  s.TableName = &v
  return s
}

func (s *EditMetaKnowledgeAssetRequest) SetTableSchemaName(v string) *EditMetaKnowledgeAssetRequest {
  s.TableSchemaName = &v
  return s
}

func (s *EditMetaKnowledgeAssetRequest) Validate() error {
  return dara.Validate(s)
}

