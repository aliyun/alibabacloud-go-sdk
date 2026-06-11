// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakePartitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *GetDataLakePartitionRequest
	GetCatalogName() *string
	SetDbName(v string) *GetDataLakePartitionRequest
	GetDbName() *string
	SetPartitionValues(v []*string) *GetDataLakePartitionRequest
	GetPartitionValues() []*string
	SetTableName(v string) *GetDataLakePartitionRequest
	GetTableName() *string
	SetTid(v int64) *GetDataLakePartitionRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *GetDataLakePartitionRequest
	GetWorkspaceId() *int64
}

type GetDataLakePartitionRequest struct {
	// The name of the data catalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// A list of partition values.
	//
	// This parameter is required.
	PartitionValues []*string `json:"PartitionValues,omitempty" xml:"PartitionValues,omitempty" type:"Repeated"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The tenant ID.
	//
	// > Hover over your profile picture in the upper-right corner of the DMS console to obtain the tenant ID. For details, see [View tenant information](https://help.aliyun.com/document_detail/181330.html).
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

func (s GetDataLakePartitionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakePartitionRequest) GoString() string {
	return s.String()
}

func (s *GetDataLakePartitionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetDataLakePartitionRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetDataLakePartitionRequest) GetPartitionValues() []*string {
	return s.PartitionValues
}

func (s *GetDataLakePartitionRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetDataLakePartitionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataLakePartitionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *GetDataLakePartitionRequest) SetCatalogName(v string) *GetDataLakePartitionRequest {
	s.CatalogName = &v
	return s
}

func (s *GetDataLakePartitionRequest) SetDbName(v string) *GetDataLakePartitionRequest {
	s.DbName = &v
	return s
}

func (s *GetDataLakePartitionRequest) SetPartitionValues(v []*string) *GetDataLakePartitionRequest {
	s.PartitionValues = v
	return s
}

func (s *GetDataLakePartitionRequest) SetTableName(v string) *GetDataLakePartitionRequest {
	s.TableName = &v
	return s
}

func (s *GetDataLakePartitionRequest) SetTid(v int64) *GetDataLakePartitionRequest {
	s.Tid = &v
	return s
}

func (s *GetDataLakePartitionRequest) SetWorkspaceId(v int64) *GetDataLakePartitionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDataLakePartitionRequest) Validate() error {
	return dara.Validate(s)
}
