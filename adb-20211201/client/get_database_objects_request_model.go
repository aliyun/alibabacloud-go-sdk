// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetDatabaseObjectsRequest
	GetDBClusterId() *string
	SetFilterOwner(v string) *GetDatabaseObjectsRequest
	GetFilterOwner() *string
	SetFilterSchemaName(v string) *GetDatabaseObjectsRequest
	GetFilterSchemaName() *string
	SetOrderBy(v string) *GetDatabaseObjectsRequest
	GetOrderBy() *string
	SetPageNumber(v int64) *GetDatabaseObjectsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetDatabaseObjectsRequest
	GetPageSize() *int64
	SetRegionId(v string) *GetDatabaseObjectsRequest
	GetRegionId() *string
}

type GetDatabaseObjectsRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// am-bp1565u55p32****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The owner of the database.
	//
	// example:
	//
	// admin
	FilterOwner *string `json:"FilterOwner,omitempty" xml:"FilterOwner,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test_db
	FilterSchemaName *string `json:"FilterSchemaName,omitempty" xml:"FilterSchemaName,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// 	- Asc
	//
	// 	- Desc
	//
	// Valid values for Field: DatabaseName, CreateTime, and UpdateTime. -CreateTime; -UpdateTime;
	//
	// Default value: {"Type": "Desc","Field": "DatabaseName"}.
	//
	// example:
	//
	// {"Type": "Desc","Field": "DbName"}
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- 30
	//
	// 	- 50
	//
	// 	- 100
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDatabaseObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseObjectsRequest) GoString() string {
	return s.String()
}

func (s *GetDatabaseObjectsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetDatabaseObjectsRequest) GetFilterOwner() *string {
	return s.FilterOwner
}

func (s *GetDatabaseObjectsRequest) GetFilterSchemaName() *string {
	return s.FilterSchemaName
}

func (s *GetDatabaseObjectsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetDatabaseObjectsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetDatabaseObjectsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetDatabaseObjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDatabaseObjectsRequest) SetDBClusterId(v string) *GetDatabaseObjectsRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetFilterOwner(v string) *GetDatabaseObjectsRequest {
	s.FilterOwner = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetFilterSchemaName(v string) *GetDatabaseObjectsRequest {
	s.FilterSchemaName = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetOrderBy(v string) *GetDatabaseObjectsRequest {
	s.OrderBy = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetPageNumber(v int64) *GetDatabaseObjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetPageSize(v int64) *GetDatabaseObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *GetDatabaseObjectsRequest) SetRegionId(v string) *GetDatabaseObjectsRequest {
	s.RegionId = &v
	return s
}

func (s *GetDatabaseObjectsRequest) Validate() error {
	return dara.Validate(s)
}
