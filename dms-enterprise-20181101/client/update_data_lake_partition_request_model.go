// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakePartitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *UpdateDataLakePartitionRequest
	GetCatalogName() *string
	SetDataRegion(v string) *UpdateDataLakePartitionRequest
	GetDataRegion() *string
	SetDbName(v string) *UpdateDataLakePartitionRequest
	GetDbName() *string
	SetPartitionInput(v *DLPartitionInput) *UpdateDataLakePartitionRequest
	GetPartitionInput() *DLPartitionInput
	SetTableName(v string) *UpdateDataLakePartitionRequest
	GetTableName() *string
	SetTid(v int64) *UpdateDataLakePartitionRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *UpdateDataLakePartitionRequest
	GetWorkspaceId() *int64
}

type UpdateDataLakePartitionRequest struct {
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
	// The information about the created partition.
	//
	// This parameter is required.
	PartitionInput *DLPartitionInput `json:"PartitionInput,omitempty" xml:"PartitionInput,omitempty"`
	// The name of the table
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
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDataLakePartitionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakePartitionRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataLakePartitionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *UpdateDataLakePartitionRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *UpdateDataLakePartitionRequest) GetDbName() *string {
	return s.DbName
}

func (s *UpdateDataLakePartitionRequest) GetPartitionInput() *DLPartitionInput {
	return s.PartitionInput
}

func (s *UpdateDataLakePartitionRequest) GetTableName() *string {
	return s.TableName
}

func (s *UpdateDataLakePartitionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateDataLakePartitionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdateDataLakePartitionRequest) SetCatalogName(v string) *UpdateDataLakePartitionRequest {
	s.CatalogName = &v
	return s
}

func (s *UpdateDataLakePartitionRequest) SetDataRegion(v string) *UpdateDataLakePartitionRequest {
	s.DataRegion = &v
	return s
}

func (s *UpdateDataLakePartitionRequest) SetDbName(v string) *UpdateDataLakePartitionRequest {
	s.DbName = &v
	return s
}

func (s *UpdateDataLakePartitionRequest) SetPartitionInput(v *DLPartitionInput) *UpdateDataLakePartitionRequest {
	s.PartitionInput = v
	return s
}

func (s *UpdateDataLakePartitionRequest) SetTableName(v string) *UpdateDataLakePartitionRequest {
	s.TableName = &v
	return s
}

func (s *UpdateDataLakePartitionRequest) SetTid(v int64) *UpdateDataLakePartitionRequest {
	s.Tid = &v
	return s
}

func (s *UpdateDataLakePartitionRequest) SetWorkspaceId(v int64) *UpdateDataLakePartitionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataLakePartitionRequest) Validate() error {
	if s.PartitionInput != nil {
		if err := s.PartitionInput.Validate(); err != nil {
			return err
		}
	}
	return nil
}
