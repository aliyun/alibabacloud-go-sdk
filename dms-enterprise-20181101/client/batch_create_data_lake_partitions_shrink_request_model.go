// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateDataLakePartitionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *BatchCreateDataLakePartitionsShrinkRequest
	GetCatalogName() *string
	SetDataRegion(v string) *BatchCreateDataLakePartitionsShrinkRequest
	GetDataRegion() *string
	SetDbName(v string) *BatchCreateDataLakePartitionsShrinkRequest
	GetDbName() *string
	SetIfNotExists(v bool) *BatchCreateDataLakePartitionsShrinkRequest
	GetIfNotExists() *bool
	SetNeedResult(v bool) *BatchCreateDataLakePartitionsShrinkRequest
	GetNeedResult() *bool
	SetPartitionInputsShrink(v string) *BatchCreateDataLakePartitionsShrinkRequest
	GetPartitionInputsShrink() *string
	SetTableName(v string) *BatchCreateDataLakePartitionsShrinkRequest
	GetTableName() *string
	SetTid(v int64) *BatchCreateDataLakePartitionsShrinkRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *BatchCreateDataLakePartitionsShrinkRequest
	GetWorkspaceId() *int64
}

type BatchCreateDataLakePartitionsShrinkRequest struct {
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
	// example:
	//
	// true
	IfNotExists *bool `json:"IfNotExists,omitempty" xml:"IfNotExists,omitempty"`
	// example:
	//
	// true
	NeedResult *bool `json:"NeedResult,omitempty" xml:"NeedResult,omitempty"`
	// This parameter is required.
	PartitionInputsShrink *string `json:"PartitionInputs,omitempty" xml:"PartitionInputs,omitempty"`
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

func (s BatchCreateDataLakePartitionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateDataLakePartitionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) GetIfNotExists() *bool {
	return s.IfNotExists
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) GetNeedResult() *bool {
	return s.NeedResult
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) GetPartitionInputsShrink() *string {
	return s.PartitionInputsShrink
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) SetCatalogName(v string) *BatchCreateDataLakePartitionsShrinkRequest {
	s.CatalogName = &v
	return s
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) SetDataRegion(v string) *BatchCreateDataLakePartitionsShrinkRequest {
	s.DataRegion = &v
	return s
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) SetDbName(v string) *BatchCreateDataLakePartitionsShrinkRequest {
	s.DbName = &v
	return s
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) SetIfNotExists(v bool) *BatchCreateDataLakePartitionsShrinkRequest {
	s.IfNotExists = &v
	return s
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) SetNeedResult(v bool) *BatchCreateDataLakePartitionsShrinkRequest {
	s.NeedResult = &v
	return s
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) SetPartitionInputsShrink(v string) *BatchCreateDataLakePartitionsShrinkRequest {
	s.PartitionInputsShrink = &v
	return s
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) SetTableName(v string) *BatchCreateDataLakePartitionsShrinkRequest {
	s.TableName = &v
	return s
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) SetTid(v int64) *BatchCreateDataLakePartitionsShrinkRequest {
	s.Tid = &v
	return s
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) SetWorkspaceId(v int64) *BatchCreateDataLakePartitionsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchCreateDataLakePartitionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
