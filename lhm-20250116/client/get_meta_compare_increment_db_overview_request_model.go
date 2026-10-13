// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCompareIncrementDbOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *GetMetaCompareIncrementDbOverviewRequest
	GetCatalogName() *string
	SetDbName(v string) *GetMetaCompareIncrementDbOverviewRequest
	GetDbName() *string
	SetIsSummary(v bool) *GetMetaCompareIncrementDbOverviewRequest
	GetIsSummary() *bool
	SetTaskId(v int64) *GetMetaCompareIncrementDbOverviewRequest
	GetTaskId() *int64
}

type GetMetaCompareIncrementDbOverviewRequest struct {
	// The name of the catalog. In a multi-catalog environment, it is used to locate the catalog where the destination database resides.
	//
	// example:
	//
	// fluss_e2e_dlf_catalog
	CatalogName *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	// The name of the database for which you want to query incremental changes.
	//
	// This parameter is required.
	//
	// example:
	//
	// ods_eaps
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// Specifies whether to return only summary statistics.
	IsSummary *bool `json:"isSummary,omitempty" xml:"isSummary,omitempty"`
	// The ID of the metadata comparison task, which is obtained from the return value of the create or save comparison task API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetMetaCompareIncrementDbOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCompareIncrementDbOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetMetaCompareIncrementDbOverviewRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetMetaCompareIncrementDbOverviewRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetMetaCompareIncrementDbOverviewRequest) GetIsSummary() *bool {
	return s.IsSummary
}

func (s *GetMetaCompareIncrementDbOverviewRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetMetaCompareIncrementDbOverviewRequest) SetCatalogName(v string) *GetMetaCompareIncrementDbOverviewRequest {
	s.CatalogName = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewRequest) SetDbName(v string) *GetMetaCompareIncrementDbOverviewRequest {
	s.DbName = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewRequest) SetIsSummary(v bool) *GetMetaCompareIncrementDbOverviewRequest {
	s.IsSummary = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewRequest) SetTaskId(v int64) *GetMetaCompareIncrementDbOverviewRequest {
	s.TaskId = &v
	return s
}

func (s *GetMetaCompareIncrementDbOverviewRequest) Validate() error {
	return dara.Validate(s)
}
