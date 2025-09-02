// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMetaTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGuid(v string) *SearchMetaTablesRequest
	GetAppGuid() *string
	SetClusterId(v string) *SearchMetaTablesRequest
	GetClusterId() *string
	SetDataSourceType(v string) *SearchMetaTablesRequest
	GetDataSourceType() *string
	SetEntityType(v int32) *SearchMetaTablesRequest
	GetEntityType() *int32
	SetKeyword(v string) *SearchMetaTablesRequest
	GetKeyword() *string
	SetPageNumber(v int32) *SearchMetaTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchMetaTablesRequest
	GetPageSize() *int32
	SetSchema(v string) *SearchMetaTablesRequest
	GetSchema() *string
}

type SearchMetaTablesRequest struct {
	// The GUID of the workspace where the metatables reside.
	//
	// example:
	//
	// odps.engine_name
	AppGuid *string `json:"AppGuid,omitempty" xml:"AppGuid,omitempty"`
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
	// The type of the metatables. Valid values: 0 and 1. The value 0 indicates that tables are queried. The value 1 indicates that views are queried. If you do not configure this parameter, all types of metatables are queried.
	//
	// example:
	//
	// 0
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The keyword based on which metatables are queried. During the query, the system tokenizes the names of metatables and matches the names with the keyword. If no name is matched, the value null is returned. By default, the system uses underscores (_) to tokenize the names.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
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
	// The schema information of the table. You must configure this parameter if you enable the three-layer model of MaxCompute.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// default
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s SearchMetaTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMetaTablesRequest) GoString() string {
	return s.String()
}

func (s *SearchMetaTablesRequest) GetAppGuid() *string {
	return s.AppGuid
}

func (s *SearchMetaTablesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *SearchMetaTablesRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *SearchMetaTablesRequest) GetEntityType() *int32 {
	return s.EntityType
}

func (s *SearchMetaTablesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchMetaTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchMetaTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchMetaTablesRequest) GetSchema() *string {
	return s.Schema
}

func (s *SearchMetaTablesRequest) SetAppGuid(v string) *SearchMetaTablesRequest {
	s.AppGuid = &v
	return s
}

func (s *SearchMetaTablesRequest) SetClusterId(v string) *SearchMetaTablesRequest {
	s.ClusterId = &v
	return s
}

func (s *SearchMetaTablesRequest) SetDataSourceType(v string) *SearchMetaTablesRequest {
	s.DataSourceType = &v
	return s
}

func (s *SearchMetaTablesRequest) SetEntityType(v int32) *SearchMetaTablesRequest {
	s.EntityType = &v
	return s
}

func (s *SearchMetaTablesRequest) SetKeyword(v string) *SearchMetaTablesRequest {
	s.Keyword = &v
	return s
}

func (s *SearchMetaTablesRequest) SetPageNumber(v int32) *SearchMetaTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchMetaTablesRequest) SetPageSize(v int32) *SearchMetaTablesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchMetaTablesRequest) SetSchema(v string) *SearchMetaTablesRequest {
	s.Schema = &v
	return s
}

func (s *SearchMetaTablesRequest) Validate() error {
	return dara.Validate(s)
}
