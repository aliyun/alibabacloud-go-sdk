// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTablePartitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetMetaTablePartitionRequest
	GetClusterId() *string
	SetDataSourceType(v string) *GetMetaTablePartitionRequest
	GetDataSourceType() *string
	SetDatabaseName(v string) *GetMetaTablePartitionRequest
	GetDatabaseName() *string
	SetPageNumber(v int32) *GetMetaTablePartitionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetMetaTablePartitionRequest
	GetPageSize() *int32
	SetSortCriterion(v *GetMetaTablePartitionRequestSortCriterion) *GetMetaTablePartitionRequest
	GetSortCriterion() *GetMetaTablePartitionRequestSortCriterion
	SetTableGuid(v string) *GetMetaTablePartitionRequest
	GetTableGuid() *string
	SetTableName(v string) *GetMetaTablePartitionRequest
	GetTableName() *string
}

type GetMetaTablePartitionRequest struct {
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
	SortCriterion *GetMetaTablePartitionRequestSortCriterion `json:"SortCriterion,omitempty" xml:"SortCriterion,omitempty" type:"Struct"`
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

func (s GetMetaTablePartitionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTablePartitionRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTablePartitionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaTablePartitionRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetMetaTablePartitionRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaTablePartitionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMetaTablePartitionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTablePartitionRequest) GetSortCriterion() *GetMetaTablePartitionRequestSortCriterion {
	return s.SortCriterion
}

func (s *GetMetaTablePartitionRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTablePartitionRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaTablePartitionRequest) SetClusterId(v string) *GetMetaTablePartitionRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMetaTablePartitionRequest) SetDataSourceType(v string) *GetMetaTablePartitionRequest {
	s.DataSourceType = &v
	return s
}

func (s *GetMetaTablePartitionRequest) SetDatabaseName(v string) *GetMetaTablePartitionRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaTablePartitionRequest) SetPageNumber(v int32) *GetMetaTablePartitionRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMetaTablePartitionRequest) SetPageSize(v int32) *GetMetaTablePartitionRequest {
	s.PageSize = &v
	return s
}

func (s *GetMetaTablePartitionRequest) SetSortCriterion(v *GetMetaTablePartitionRequestSortCriterion) *GetMetaTablePartitionRequest {
	s.SortCriterion = v
	return s
}

func (s *GetMetaTablePartitionRequest) SetTableGuid(v string) *GetMetaTablePartitionRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTablePartitionRequest) SetTableName(v string) *GetMetaTablePartitionRequest {
	s.TableName = &v
	return s
}

func (s *GetMetaTablePartitionRequest) Validate() error {
	return dara.Validate(s)
}

type GetMetaTablePartitionRequestSortCriterion struct {
	// The order in which partitions in the metatable are sorted. Valid values: asc and desc. Default value: desc.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field that is used to sort partitions in the metatable. Valid values: name and modify_time. By default, partitions in the metatable are sorted based on their creation time.
	//
	// example:
	//
	// name
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s GetMetaTablePartitionRequestSortCriterion) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTablePartitionRequestSortCriterion) GoString() string {
	return s.String()
}

func (s *GetMetaTablePartitionRequestSortCriterion) GetOrder() *string {
	return s.Order
}

func (s *GetMetaTablePartitionRequestSortCriterion) GetSortField() *string {
	return s.SortField
}

func (s *GetMetaTablePartitionRequestSortCriterion) SetOrder(v string) *GetMetaTablePartitionRequestSortCriterion {
	s.Order = &v
	return s
}

func (s *GetMetaTablePartitionRequestSortCriterion) SetSortField(v string) *GetMetaTablePartitionRequestSortCriterion {
	s.SortField = &v
	return s
}

func (s *GetMetaTablePartitionRequestSortCriterion) Validate() error {
	return dara.Validate(s)
}
