// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakePartitionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *CreateDataLakePartitionShrinkRequest
	GetCatalogName() *string
	SetDataRegion(v string) *CreateDataLakePartitionShrinkRequest
	GetDataRegion() *string
	SetDbName(v string) *CreateDataLakePartitionShrinkRequest
	GetDbName() *string
	SetIfNotExists(v bool) *CreateDataLakePartitionShrinkRequest
	GetIfNotExists() *bool
	SetNeedResult(v bool) *CreateDataLakePartitionShrinkRequest
	GetNeedResult() *bool
	SetPartitionInputShrink(v string) *CreateDataLakePartitionShrinkRequest
	GetPartitionInputShrink() *string
	SetTableName(v string) *CreateDataLakePartitionShrinkRequest
	GetTableName() *string
	SetTid(v int64) *CreateDataLakePartitionShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *CreateDataLakePartitionShrinkRequest
	GetWorkspaceId() *int64
}

type CreateDataLakePartitionShrinkRequest struct {
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
	PartitionInputShrink *string `json:"PartitionInput,omitempty" xml:"PartitionInput,omitempty"`
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

func (s CreateDataLakePartitionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakePartitionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataLakePartitionShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *CreateDataLakePartitionShrinkRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *CreateDataLakePartitionShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateDataLakePartitionShrinkRequest) GetIfNotExists() *bool {
	return s.IfNotExists
}

func (s *CreateDataLakePartitionShrinkRequest) GetNeedResult() *bool {
	return s.NeedResult
}

func (s *CreateDataLakePartitionShrinkRequest) GetPartitionInputShrink() *string {
	return s.PartitionInputShrink
}

func (s *CreateDataLakePartitionShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateDataLakePartitionShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataLakePartitionShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateDataLakePartitionShrinkRequest) SetCatalogName(v string) *CreateDataLakePartitionShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *CreateDataLakePartitionShrinkRequest) SetDataRegion(v string) *CreateDataLakePartitionShrinkRequest {
	s.DataRegion = &v
	return s
}

func (s *CreateDataLakePartitionShrinkRequest) SetDbName(v string) *CreateDataLakePartitionShrinkRequest {
	s.DbName = &v
	return s
}

func (s *CreateDataLakePartitionShrinkRequest) SetIfNotExists(v bool) *CreateDataLakePartitionShrinkRequest {
	s.IfNotExists = &v
	return s
}

func (s *CreateDataLakePartitionShrinkRequest) SetNeedResult(v bool) *CreateDataLakePartitionShrinkRequest {
	s.NeedResult = &v
	return s
}

func (s *CreateDataLakePartitionShrinkRequest) SetPartitionInputShrink(v string) *CreateDataLakePartitionShrinkRequest {
	s.PartitionInputShrink = &v
	return s
}

func (s *CreateDataLakePartitionShrinkRequest) SetTableName(v string) *CreateDataLakePartitionShrinkRequest {
	s.TableName = &v
	return s
}

func (s *CreateDataLakePartitionShrinkRequest) SetTid(v int64) *CreateDataLakePartitionShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataLakePartitionShrinkRequest) SetWorkspaceId(v int64) *CreateDataLakePartitionShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataLakePartitionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
