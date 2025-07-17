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
	SetDataRegion(v string) *GetDataLakePartitionRequest
	GetDataRegion() *string
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
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	PartitionValues []*string `json:"PartitionValues,omitempty" xml:"PartitionValues,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 3***
	Tid         *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
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

func (s *GetDataLakePartitionRequest) GetDataRegion() *string {
	return s.DataRegion
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

func (s *GetDataLakePartitionRequest) SetDataRegion(v string) *GetDataLakePartitionRequest {
	s.DataRegion = &v
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
