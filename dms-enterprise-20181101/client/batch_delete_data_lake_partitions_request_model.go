// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDataLakePartitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *BatchDeleteDataLakePartitionsRequest
	GetCatalogName() *string
	SetDataRegion(v string) *BatchDeleteDataLakePartitionsRequest
	GetDataRegion() *string
	SetDbName(v string) *BatchDeleteDataLakePartitionsRequest
	GetDbName() *string
	SetIfExists(v bool) *BatchDeleteDataLakePartitionsRequest
	GetIfExists() *bool
	SetPartitionValuesList(v [][]*string) *BatchDeleteDataLakePartitionsRequest
	GetPartitionValuesList() [][]*string
	SetTableName(v string) *BatchDeleteDataLakePartitionsRequest
	GetTableName() *string
	SetTid(v int64) *BatchDeleteDataLakePartitionsRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *BatchDeleteDataLakePartitionsRequest
	GetWorkspaceId() *int64
}

type BatchDeleteDataLakePartitionsRequest struct {
	// The name of the data catalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The region where the data lake resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Specifies whether to ignore the exception if the partition that you want to delete does not exist.
	//
	// example:
	//
	// true
	IfExists *bool `json:"IfExists,omitempty" xml:"IfExists,omitempty"`
	// The values in a partition key column.
	//
	// This parameter is required.
	PartitionValuesList [][]*string `json:"PartitionValuesList,omitempty" xml:"PartitionValuesList,omitempty" type:"Repeated"`
	// The table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// table_name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s BatchDeleteDataLakePartitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDataLakePartitionsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDataLakePartitionsRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *BatchDeleteDataLakePartitionsRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *BatchDeleteDataLakePartitionsRequest) GetDbName() *string {
	return s.DbName
}

func (s *BatchDeleteDataLakePartitionsRequest) GetIfExists() *bool {
	return s.IfExists
}

func (s *BatchDeleteDataLakePartitionsRequest) GetPartitionValuesList() [][]*string {
	return s.PartitionValuesList
}

func (s *BatchDeleteDataLakePartitionsRequest) GetTableName() *string {
	return s.TableName
}

func (s *BatchDeleteDataLakePartitionsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *BatchDeleteDataLakePartitionsRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *BatchDeleteDataLakePartitionsRequest) SetCatalogName(v string) *BatchDeleteDataLakePartitionsRequest {
	s.CatalogName = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetDataRegion(v string) *BatchDeleteDataLakePartitionsRequest {
	s.DataRegion = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetDbName(v string) *BatchDeleteDataLakePartitionsRequest {
	s.DbName = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetIfExists(v bool) *BatchDeleteDataLakePartitionsRequest {
	s.IfExists = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetPartitionValuesList(v [][]*string) *BatchDeleteDataLakePartitionsRequest {
	s.PartitionValuesList = v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetTableName(v string) *BatchDeleteDataLakePartitionsRequest {
	s.TableName = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetTid(v int64) *BatchDeleteDataLakePartitionsRequest {
	s.Tid = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) SetWorkspaceId(v int64) *BatchDeleteDataLakePartitionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsRequest) Validate() error {
	return dara.Validate(s)
}
