// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateDataLakePartitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *BatchCreateDataLakePartitionsRequest
	GetCatalogName() *string
	SetDbName(v string) *BatchCreateDataLakePartitionsRequest
	GetDbName() *string
	SetIfNotExists(v bool) *BatchCreateDataLakePartitionsRequest
	GetIfNotExists() *bool
	SetNeedResult(v bool) *BatchCreateDataLakePartitionsRequest
	GetNeedResult() *bool
	SetPartitionInputs(v []*DLPartitionInput) *BatchCreateDataLakePartitionsRequest
	GetPartitionInputs() []*DLPartitionInput
	SetTableName(v string) *BatchCreateDataLakePartitionsRequest
	GetTableName() *string
	SetTid(v int64) *BatchCreateDataLakePartitionsRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *BatchCreateDataLakePartitionsRequest
	GetWorkspaceId() *int64
}

type BatchCreateDataLakePartitionsRequest struct {
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
	PartitionInputs []*DLPartitionInput `json:"PartitionInputs,omitempty" xml:"PartitionInputs,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s BatchCreateDataLakePartitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateDataLakePartitionsRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateDataLakePartitionsRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *BatchCreateDataLakePartitionsRequest) GetDbName() *string {
	return s.DbName
}

func (s *BatchCreateDataLakePartitionsRequest) GetIfNotExists() *bool {
	return s.IfNotExists
}

func (s *BatchCreateDataLakePartitionsRequest) GetNeedResult() *bool {
	return s.NeedResult
}

func (s *BatchCreateDataLakePartitionsRequest) GetPartitionInputs() []*DLPartitionInput {
	return s.PartitionInputs
}

func (s *BatchCreateDataLakePartitionsRequest) GetTableName() *string {
	return s.TableName
}

func (s *BatchCreateDataLakePartitionsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *BatchCreateDataLakePartitionsRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *BatchCreateDataLakePartitionsRequest) SetCatalogName(v string) *BatchCreateDataLakePartitionsRequest {
	s.CatalogName = &v
	return s
}

func (s *BatchCreateDataLakePartitionsRequest) SetDbName(v string) *BatchCreateDataLakePartitionsRequest {
	s.DbName = &v
	return s
}

func (s *BatchCreateDataLakePartitionsRequest) SetIfNotExists(v bool) *BatchCreateDataLakePartitionsRequest {
	s.IfNotExists = &v
	return s
}

func (s *BatchCreateDataLakePartitionsRequest) SetNeedResult(v bool) *BatchCreateDataLakePartitionsRequest {
	s.NeedResult = &v
	return s
}

func (s *BatchCreateDataLakePartitionsRequest) SetPartitionInputs(v []*DLPartitionInput) *BatchCreateDataLakePartitionsRequest {
	s.PartitionInputs = v
	return s
}

func (s *BatchCreateDataLakePartitionsRequest) SetTableName(v string) *BatchCreateDataLakePartitionsRequest {
	s.TableName = &v
	return s
}

func (s *BatchCreateDataLakePartitionsRequest) SetTid(v int64) *BatchCreateDataLakePartitionsRequest {
	s.Tid = &v
	return s
}

func (s *BatchCreateDataLakePartitionsRequest) SetWorkspaceId(v int64) *BatchCreateDataLakePartitionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BatchCreateDataLakePartitionsRequest) Validate() error {
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
