// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeArchiveTableListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeArchiveTableListRequest
	GetDBInstanceName() *string
	SetPageIndex(v int64) *DescribeArchiveTableListRequest
	GetPageIndex() *int64
	SetPageSize(v int64) *DescribeArchiveTableListRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeArchiveTableListRequest
	GetRegionId() *string
	SetSchemaName(v string) *DescribeArchiveTableListRequest
	GetSchemaName() *string
	SetStatus(v string) *DescribeArchiveTableListRequest
	GetStatus() *string
	SetTableName(v string) *DescribeArchiveTableListRequest
	GetTableName() *string
}

type DescribeArchiveTableListRequest struct {
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// This parameter is required.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName  *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeArchiveTableListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeArchiveTableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeArchiveTableListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeArchiveTableListRequest) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *DescribeArchiveTableListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeArchiveTableListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeArchiveTableListRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeArchiveTableListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeArchiveTableListRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeArchiveTableListRequest) SetDBInstanceName(v string) *DescribeArchiveTableListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetPageIndex(v int64) *DescribeArchiveTableListRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetPageSize(v int64) *DescribeArchiveTableListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetRegionId(v string) *DescribeArchiveTableListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetSchemaName(v string) *DescribeArchiveTableListRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetStatus(v string) *DescribeArchiveTableListRequest {
	s.Status = &v
	return s
}

func (s *DescribeArchiveTableListRequest) SetTableName(v string) *DescribeArchiveTableListRequest {
	s.TableName = &v
	return s
}

func (s *DescribeArchiveTableListRequest) Validate() error {
	return dara.Validate(s)
}
