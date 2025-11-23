// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakePartitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *CreateDataLakePartitionRequest
	GetCatalogName() *string
	SetDataRegion(v string) *CreateDataLakePartitionRequest
	GetDataRegion() *string
	SetDbName(v string) *CreateDataLakePartitionRequest
	GetDbName() *string
	SetIfNotExists(v bool) *CreateDataLakePartitionRequest
	GetIfNotExists() *bool
	SetNeedResult(v bool) *CreateDataLakePartitionRequest
	GetNeedResult() *bool
	SetPartitionInput(v *DLPartitionInput) *CreateDataLakePartitionRequest
	GetPartitionInput() *DLPartitionInput
	SetTableName(v string) *CreateDataLakePartitionRequest
	GetTableName() *string
	SetTid(v int64) *CreateDataLakePartitionRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *CreateDataLakePartitionRequest
	GetWorkspaceId() *int64
}

type CreateDataLakePartitionRequest struct {
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
	// Specifies whether to ignore the exception if the name of the created partition is the same as the name of an existing partition.
	//
	// example:
	//
	// true
	IfNotExists *bool `json:"IfNotExists,omitempty" xml:"IfNotExists,omitempty"`
	// Specifies whether to return information about the created partition. If the value is true, the Partition parameter is returned. Valid values:
	//
	// 	- true: returns information about the created partition.
	//
	// 	- false: does not return information about the created partition.
	//
	// example:
	//
	// true
	NeedResult *bool `json:"NeedResult,omitempty" xml:"NeedResult,omitempty"`
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
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
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

func (s CreateDataLakePartitionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakePartitionRequest) GoString() string {
	return s.String()
}

func (s *CreateDataLakePartitionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *CreateDataLakePartitionRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *CreateDataLakePartitionRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateDataLakePartitionRequest) GetIfNotExists() *bool {
	return s.IfNotExists
}

func (s *CreateDataLakePartitionRequest) GetNeedResult() *bool {
	return s.NeedResult
}

func (s *CreateDataLakePartitionRequest) GetPartitionInput() *DLPartitionInput {
	return s.PartitionInput
}

func (s *CreateDataLakePartitionRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateDataLakePartitionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataLakePartitionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateDataLakePartitionRequest) SetCatalogName(v string) *CreateDataLakePartitionRequest {
	s.CatalogName = &v
	return s
}

func (s *CreateDataLakePartitionRequest) SetDataRegion(v string) *CreateDataLakePartitionRequest {
	s.DataRegion = &v
	return s
}

func (s *CreateDataLakePartitionRequest) SetDbName(v string) *CreateDataLakePartitionRequest {
	s.DbName = &v
	return s
}

func (s *CreateDataLakePartitionRequest) SetIfNotExists(v bool) *CreateDataLakePartitionRequest {
	s.IfNotExists = &v
	return s
}

func (s *CreateDataLakePartitionRequest) SetNeedResult(v bool) *CreateDataLakePartitionRequest {
	s.NeedResult = &v
	return s
}

func (s *CreateDataLakePartitionRequest) SetPartitionInput(v *DLPartitionInput) *CreateDataLakePartitionRequest {
	s.PartitionInput = v
	return s
}

func (s *CreateDataLakePartitionRequest) SetTableName(v string) *CreateDataLakePartitionRequest {
	s.TableName = &v
	return s
}

func (s *CreateDataLakePartitionRequest) SetTid(v int64) *CreateDataLakePartitionRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataLakePartitionRequest) SetWorkspaceId(v int64) *CreateDataLakePartitionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataLakePartitionRequest) Validate() error {
	if s.PartitionInput != nil {
		if err := s.PartitionInput.Validate(); err != nil {
			return err
		}
	}
	return nil
}
