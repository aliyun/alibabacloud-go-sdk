// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *GetTableColumnsRequest
	GetColumnName() *string
	SetDBClusterId(v string) *GetTableColumnsRequest
	GetDBClusterId() *string
	SetPageNumber(v int64) *GetTableColumnsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetTableColumnsRequest
	GetPageSize() *int64
	SetRegionId(v string) *GetTableColumnsRequest
	GetRegionId() *string
	SetSchemaName(v string) *GetTableColumnsRequest
	GetSchemaName() *string
	SetTableName(v string) *GetTableColumnsRequest
	GetTableName() *string
}

type GetTableColumnsRequest struct {
	// The name of the column.
	//
	// example:
	//
	// assist_user_phone
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnsRequest) GoString() string {
	return s.String()
}

func (s *GetTableColumnsRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetTableColumnsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetTableColumnsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetTableColumnsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetTableColumnsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTableColumnsRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetTableColumnsRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetTableColumnsRequest) SetColumnName(v string) *GetTableColumnsRequest {
	s.ColumnName = &v
	return s
}

func (s *GetTableColumnsRequest) SetDBClusterId(v string) *GetTableColumnsRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetTableColumnsRequest) SetPageNumber(v int64) *GetTableColumnsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTableColumnsRequest) SetPageSize(v int64) *GetTableColumnsRequest {
	s.PageSize = &v
	return s
}

func (s *GetTableColumnsRequest) SetRegionId(v string) *GetTableColumnsRequest {
	s.RegionId = &v
	return s
}

func (s *GetTableColumnsRequest) SetSchemaName(v string) *GetTableColumnsRequest {
	s.SchemaName = &v
	return s
}

func (s *GetTableColumnsRequest) SetTableName(v string) *GetTableColumnsRequest {
	s.TableName = &v
	return s
}

func (s *GetTableColumnsRequest) Validate() error {
	return dara.Validate(s)
}
