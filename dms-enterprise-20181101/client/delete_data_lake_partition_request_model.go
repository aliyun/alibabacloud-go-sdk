// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakePartitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *DeleteDataLakePartitionRequest
	GetCatalogName() *string
	SetDataRegion(v string) *DeleteDataLakePartitionRequest
	GetDataRegion() *string
	SetDbName(v string) *DeleteDataLakePartitionRequest
	GetDbName() *string
	SetIfExists(v bool) *DeleteDataLakePartitionRequest
	GetIfExists() *bool
	SetPartitionValues(v []*string) *DeleteDataLakePartitionRequest
	GetPartitionValues() []*string
	SetTableName(v string) *DeleteDataLakePartitionRequest
	GetTableName() *string
	SetTid(v int64) *DeleteDataLakePartitionRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *DeleteDataLakePartitionRequest
	GetWorkspaceId() *int64
}

type DeleteDataLakePartitionRequest struct {
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
	// The name of the database that you want to query.
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
	PartitionValues []*string `json:"PartitionValues,omitempty" xml:"PartitionValues,omitempty" type:"Repeated"`
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
	// > To view the ID of the tenant, go to the DMS console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDataLakePartitionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakePartitionRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLakePartitionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *DeleteDataLakePartitionRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *DeleteDataLakePartitionRequest) GetDbName() *string {
	return s.DbName
}

func (s *DeleteDataLakePartitionRequest) GetIfExists() *bool {
	return s.IfExists
}

func (s *DeleteDataLakePartitionRequest) GetPartitionValues() []*string {
	return s.PartitionValues
}

func (s *DeleteDataLakePartitionRequest) GetTableName() *string {
	return s.TableName
}

func (s *DeleteDataLakePartitionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteDataLakePartitionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *DeleteDataLakePartitionRequest) SetCatalogName(v string) *DeleteDataLakePartitionRequest {
	s.CatalogName = &v
	return s
}

func (s *DeleteDataLakePartitionRequest) SetDataRegion(v string) *DeleteDataLakePartitionRequest {
	s.DataRegion = &v
	return s
}

func (s *DeleteDataLakePartitionRequest) SetDbName(v string) *DeleteDataLakePartitionRequest {
	s.DbName = &v
	return s
}

func (s *DeleteDataLakePartitionRequest) SetIfExists(v bool) *DeleteDataLakePartitionRequest {
	s.IfExists = &v
	return s
}

func (s *DeleteDataLakePartitionRequest) SetPartitionValues(v []*string) *DeleteDataLakePartitionRequest {
	s.PartitionValues = v
	return s
}

func (s *DeleteDataLakePartitionRequest) SetTableName(v string) *DeleteDataLakePartitionRequest {
	s.TableName = &v
	return s
}

func (s *DeleteDataLakePartitionRequest) SetTid(v int64) *DeleteDataLakePartitionRequest {
	s.Tid = &v
	return s
}

func (s *DeleteDataLakePartitionRequest) SetWorkspaceId(v int64) *DeleteDataLakePartitionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataLakePartitionRequest) Validate() error {
	return dara.Validate(s)
}
