// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaDBTableListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGuid(v string) *GetMetaDBTableListRequest
	GetAppGuid() *string
	SetClusterId(v string) *GetMetaDBTableListRequest
	GetClusterId() *string
	SetDataSourceType(v string) *GetMetaDBTableListRequest
	GetDataSourceType() *string
	SetDatabaseName(v string) *GetMetaDBTableListRequest
	GetDatabaseName() *string
	SetPageNumber(v int32) *GetMetaDBTableListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetMetaDBTableListRequest
	GetPageSize() *int32
}

type GetMetaDBTableListRequest struct {
	// The unique identifier for the project, in the format `odps.{projectName}`. This parameter is required only if the `DataSourceType` is set to `odps`.
	//
	// example:
	//
	// odps.testProjectName
	AppGuid *string `json:"AppGuid,omitempty" xml:"AppGuid,omitempty"`
	// The ID of the E-MapReduce (EMR) cluster. This parameter is required only if the `DataSourceType` is set to `emr`.
	//
	// example:
	//
	// abc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the data source. Valid values: `odps` and `emr`.
	//
	// example:
	//
	// odps
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The page number to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 1,000.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetMetaDBTableListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaDBTableListRequest) GoString() string {
	return s.String()
}

func (s *GetMetaDBTableListRequest) GetAppGuid() *string {
	return s.AppGuid
}

func (s *GetMetaDBTableListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaDBTableListRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetMetaDBTableListRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaDBTableListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMetaDBTableListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaDBTableListRequest) SetAppGuid(v string) *GetMetaDBTableListRequest {
	s.AppGuid = &v
	return s
}

func (s *GetMetaDBTableListRequest) SetClusterId(v string) *GetMetaDBTableListRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMetaDBTableListRequest) SetDataSourceType(v string) *GetMetaDBTableListRequest {
	s.DataSourceType = &v
	return s
}

func (s *GetMetaDBTableListRequest) SetDatabaseName(v string) *GetMetaDBTableListRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaDBTableListRequest) SetPageNumber(v int32) *GetMetaDBTableListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMetaDBTableListRequest) SetPageSize(v int32) *GetMetaDBTableListRequest {
	s.PageSize = &v
	return s
}

func (s *GetMetaDBTableListRequest) Validate() error {
	return dara.Validate(s)
}
