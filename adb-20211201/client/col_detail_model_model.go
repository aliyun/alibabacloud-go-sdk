// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iColDetailModel interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *ColDetailModel
	GetColumnName() *string
	SetCreateTime(v string) *ColDetailModel
	GetCreateTime() *string
	SetDescription(v string) *ColDetailModel
	GetDescription() *string
	SetDistributeKey(v bool) *ColDetailModel
	GetDistributeKey() *bool
	SetNullable(v bool) *ColDetailModel
	GetNullable() *bool
	SetPartitionKey(v bool) *ColDetailModel
	GetPartitionKey() *bool
	SetPrimaryKey(v bool) *ColDetailModel
	GetPrimaryKey() *bool
	SetSchemaName(v string) *ColDetailModel
	GetSchemaName() *string
	SetTableName(v string) *ColDetailModel
	GetTableName() *string
	SetType(v string) *ColDetailModel
	GetType() *string
	SetUpdateTime(v string) *ColDetailModel
	GetUpdateTime() *string
}

type ColDetailModel struct {
	ColumnName    *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DistributeKey *bool   `json:"DistributeKey,omitempty" xml:"DistributeKey,omitempty"`
	Nullable      *bool   `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	PartitionKey  *bool   `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty"`
	PrimaryKey    *bool   `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	SchemaName    *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableName     *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ColDetailModel) String() string {
	return dara.Prettify(s)
}

func (s ColDetailModel) GoString() string {
	return s.String()
}

func (s *ColDetailModel) GetColumnName() *string {
	return s.ColumnName
}

func (s *ColDetailModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ColDetailModel) GetDescription() *string {
	return s.Description
}

func (s *ColDetailModel) GetDistributeKey() *bool {
	return s.DistributeKey
}

func (s *ColDetailModel) GetNullable() *bool {
	return s.Nullable
}

func (s *ColDetailModel) GetPartitionKey() *bool {
	return s.PartitionKey
}

func (s *ColDetailModel) GetPrimaryKey() *bool {
	return s.PrimaryKey
}

func (s *ColDetailModel) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ColDetailModel) GetTableName() *string {
	return s.TableName
}

func (s *ColDetailModel) GetType() *string {
	return s.Type
}

func (s *ColDetailModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ColDetailModel) SetColumnName(v string) *ColDetailModel {
	s.ColumnName = &v
	return s
}

func (s *ColDetailModel) SetCreateTime(v string) *ColDetailModel {
	s.CreateTime = &v
	return s
}

func (s *ColDetailModel) SetDescription(v string) *ColDetailModel {
	s.Description = &v
	return s
}

func (s *ColDetailModel) SetDistributeKey(v bool) *ColDetailModel {
	s.DistributeKey = &v
	return s
}

func (s *ColDetailModel) SetNullable(v bool) *ColDetailModel {
	s.Nullable = &v
	return s
}

func (s *ColDetailModel) SetPartitionKey(v bool) *ColDetailModel {
	s.PartitionKey = &v
	return s
}

func (s *ColDetailModel) SetPrimaryKey(v bool) *ColDetailModel {
	s.PrimaryKey = &v
	return s
}

func (s *ColDetailModel) SetSchemaName(v string) *ColDetailModel {
	s.SchemaName = &v
	return s
}

func (s *ColDetailModel) SetTableName(v string) *ColDetailModel {
	s.TableName = &v
	return s
}

func (s *ColDetailModel) SetType(v string) *ColDetailModel {
	s.Type = &v
	return s
}

func (s *ColDetailModel) SetUpdateTime(v string) *ColDetailModel {
	s.UpdateTime = &v
	return s
}

func (s *ColDetailModel) Validate() error {
	return dara.Validate(s)
}
