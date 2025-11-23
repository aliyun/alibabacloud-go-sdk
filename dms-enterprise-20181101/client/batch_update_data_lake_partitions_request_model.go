// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateDataLakePartitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *BatchUpdateDataLakePartitionsRequest
	GetCatalogName() *string
	SetDataRegion(v string) *BatchUpdateDataLakePartitionsRequest
	GetDataRegion() *string
	SetDbName(v string) *BatchUpdateDataLakePartitionsRequest
	GetDbName() *string
	SetPartitionInputs(v []*DLPartitionInput) *BatchUpdateDataLakePartitionsRequest
	GetPartitionInputs() []*DLPartitionInput
	SetTableName(v string) *BatchUpdateDataLakePartitionsRequest
	GetTableName() *string
	SetTid(v int64) *BatchUpdateDataLakePartitionsRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *BatchUpdateDataLakePartitionsRequest
	GetWorkspaceId() *int64
}

type BatchUpdateDataLakePartitionsRequest struct {
	// The name of the data directory.
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
	// The information about the created partition.
	//
	// This parameter is required.
	PartitionInputs []*DLPartitionInput `json:"PartitionInputs,omitempty" xml:"PartitionInputs,omitempty" type:"Repeated"`
	// The name of the table
	//
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
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

func (s BatchUpdateDataLakePartitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateDataLakePartitionsRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateDataLakePartitionsRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *BatchUpdateDataLakePartitionsRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *BatchUpdateDataLakePartitionsRequest) GetDbName() *string {
	return s.DbName
}

func (s *BatchUpdateDataLakePartitionsRequest) GetPartitionInputs() []*DLPartitionInput {
	return s.PartitionInputs
}

func (s *BatchUpdateDataLakePartitionsRequest) GetTableName() *string {
	return s.TableName
}

func (s *BatchUpdateDataLakePartitionsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *BatchUpdateDataLakePartitionsRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *BatchUpdateDataLakePartitionsRequest) SetCatalogName(v string) *BatchUpdateDataLakePartitionsRequest {
	s.CatalogName = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsRequest) SetDataRegion(v string) *BatchUpdateDataLakePartitionsRequest {
	s.DataRegion = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsRequest) SetDbName(v string) *BatchUpdateDataLakePartitionsRequest {
	s.DbName = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsRequest) SetPartitionInputs(v []*DLPartitionInput) *BatchUpdateDataLakePartitionsRequest {
	s.PartitionInputs = v
	return s
}

func (s *BatchUpdateDataLakePartitionsRequest) SetTableName(v string) *BatchUpdateDataLakePartitionsRequest {
	s.TableName = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsRequest) SetTid(v int64) *BatchUpdateDataLakePartitionsRequest {
	s.Tid = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsRequest) SetWorkspaceId(v int64) *BatchUpdateDataLakePartitionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsRequest) Validate() error {
	if s.PartitionInputs != nil {
		for _, item := range s.PartitionInputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
