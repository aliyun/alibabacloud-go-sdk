// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableColumnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetMetaTableColumnRequest
	GetClusterId() *string
	SetDataSourceType(v string) *GetMetaTableColumnRequest
	GetDataSourceType() *string
	SetDatabaseName(v string) *GetMetaTableColumnRequest
	GetDatabaseName() *string
	SetPageNum(v int32) *GetMetaTableColumnRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetMetaTableColumnRequest
	GetPageSize() *int32
	SetTableGuid(v string) *GetMetaTableColumnRequest
	GetTableGuid() *string
	SetTableName(v string) *GetMetaTableColumnRequest
	GetTableName() *string
}

type GetMetaTableColumnRequest struct {
	// The ID of the E-MapReduce (EMR) cluster. You can log on to the EMR console to obtain the ID.
	//
	// example:
	//
	// C-010A704DA760****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the data source. Set the value to emr.
	//
	// example:
	//
	// emr
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The name of the metadatabase of the EMR cluster. You can call the [ListMetaDB](https://help.aliyun.com/document_detail/2780105.html) operation to query the name.
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
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The GUID of the metatable. You can call the [GetMetaDBTableList](https://help.aliyun.com/document_detail/2780086.html) operation to query the GUID.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the metatable in the EMR cluster. You can call the [GetMetaDBTableList](https://help.aliyun.com/document_detail/2780086.html) operation to query the name.
	//
	// example:
	//
	// abc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetMetaTableColumnRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableColumnRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableColumnRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaTableColumnRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetMetaTableColumnRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaTableColumnRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetMetaTableColumnRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTableColumnRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableColumnRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaTableColumnRequest) SetClusterId(v string) *GetMetaTableColumnRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMetaTableColumnRequest) SetDataSourceType(v string) *GetMetaTableColumnRequest {
	s.DataSourceType = &v
	return s
}

func (s *GetMetaTableColumnRequest) SetDatabaseName(v string) *GetMetaTableColumnRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaTableColumnRequest) SetPageNum(v int32) *GetMetaTableColumnRequest {
	s.PageNum = &v
	return s
}

func (s *GetMetaTableColumnRequest) SetPageSize(v int32) *GetMetaTableColumnRequest {
	s.PageSize = &v
	return s
}

func (s *GetMetaTableColumnRequest) SetTableGuid(v string) *GetMetaTableColumnRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableColumnRequest) SetTableName(v string) *GetMetaTableColumnRequest {
	s.TableName = &v
	return s
}

func (s *GetMetaTableColumnRequest) Validate() error {
	return dara.Validate(s)
}
