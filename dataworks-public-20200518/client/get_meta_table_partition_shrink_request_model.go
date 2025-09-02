// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTablePartitionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetMetaTablePartitionShrinkRequest
	GetClusterId() *string
	SetDataSourceType(v string) *GetMetaTablePartitionShrinkRequest
	GetDataSourceType() *string
	SetDatabaseName(v string) *GetMetaTablePartitionShrinkRequest
	GetDatabaseName() *string
	SetPageNumber(v int32) *GetMetaTablePartitionShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetMetaTablePartitionShrinkRequest
	GetPageSize() *int32
	SetSortCriterionShrink(v string) *GetMetaTablePartitionShrinkRequest
	GetSortCriterionShrink() *string
	SetTableGuid(v string) *GetMetaTablePartitionShrinkRequest
	GetTableGuid() *string
	SetTableName(v string) *GetMetaTablePartitionShrinkRequest
	GetTableName() *string
}

type GetMetaTablePartitionShrinkRequest struct {
	// The ID of the EMR cluster. This parameter is required only if you set the DataSourceType parameter to emr.
	//
	// You can log on to the [EMR console](https://emr.console.aliyun.com/?spm=a2c4g.11186623.0.0.965cc5c2GeiHet#/cn-hangzhou) to obtain the ID.
	//
	// example:
	//
	// abc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the data source. Valid values: odps and emr.
	//
	// example:
	//
	// emr
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The name of the database. This parameter is required only if you set the DataSourceType parameter to emr.
	//
	// You can call the [ListMetaDB](https://help.aliyun.com/document_detail/2780105.html) operation to query the name of the metadatabase.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The logic for sorting partitions in the metatable.
	SortCriterionShrink *string `json:"SortCriterion,omitempty" xml:"SortCriterion,omitempty"`
	// The unique identifier of the metatable.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the metatable in the EMR cluster. This parameter is required only if you set the DataSourceType parameter to emr.
	//
	// You can call the [GetMetaDBTableList](https://help.aliyun.com/document_detail/2780086.html) operation to query the name of the metatable.
	//
	// example:
	//
	// abc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetMetaTablePartitionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTablePartitionShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTablePartitionShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaTablePartitionShrinkRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetMetaTablePartitionShrinkRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaTablePartitionShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMetaTablePartitionShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTablePartitionShrinkRequest) GetSortCriterionShrink() *string {
	return s.SortCriterionShrink
}

func (s *GetMetaTablePartitionShrinkRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTablePartitionShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaTablePartitionShrinkRequest) SetClusterId(v string) *GetMetaTablePartitionShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMetaTablePartitionShrinkRequest) SetDataSourceType(v string) *GetMetaTablePartitionShrinkRequest {
	s.DataSourceType = &v
	return s
}

func (s *GetMetaTablePartitionShrinkRequest) SetDatabaseName(v string) *GetMetaTablePartitionShrinkRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaTablePartitionShrinkRequest) SetPageNumber(v int32) *GetMetaTablePartitionShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMetaTablePartitionShrinkRequest) SetPageSize(v int32) *GetMetaTablePartitionShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetMetaTablePartitionShrinkRequest) SetSortCriterionShrink(v string) *GetMetaTablePartitionShrinkRequest {
	s.SortCriterionShrink = &v
	return s
}

func (s *GetMetaTablePartitionShrinkRequest) SetTableGuid(v string) *GetMetaTablePartitionShrinkRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTablePartitionShrinkRequest) SetTableName(v string) *GetMetaTablePartitionShrinkRequest {
	s.TableName = &v
	return s
}

func (s *GetMetaTablePartitionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
