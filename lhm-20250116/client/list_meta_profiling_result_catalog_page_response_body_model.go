// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingResultCatalogPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListMetaProfilingResultCatalogPageResponseBodyData) *ListMetaProfilingResultCatalogPageResponseBody
	GetData() []*ListMetaProfilingResultCatalogPageResponseBodyData
	SetErrCode(v string) *ListMetaProfilingResultCatalogPageResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMetaProfilingResultCatalogPageResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *ListMetaProfilingResultCatalogPageResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMetaProfilingResultCatalogPageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListMetaProfilingResultCatalogPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMetaProfilingResultCatalogPageResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListMetaProfilingResultCatalogPageResponseBody
	GetTotalCount() *int32
}

type ListMetaProfilingResultCatalogPageResponseBody struct {
	Data []*ListMetaProfilingResultCatalogPageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// InvalidParameter
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// The parameter is invalid or missing.
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 3F45B8C1-2A6E-4C9B-9E1F-7D2A1B0C5E66
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 128
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMetaProfilingResultCatalogPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultCatalogPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) GetData() []*ListMetaProfilingResultCatalogPageResponseBodyData {
	return s.Data
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) SetData(v []*ListMetaProfilingResultCatalogPageResponseBodyData) *ListMetaProfilingResultCatalogPageResponseBody {
	s.Data = v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) SetErrCode(v string) *ListMetaProfilingResultCatalogPageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) SetErrMessage(v string) *ListMetaProfilingResultCatalogPageResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) SetPageIndex(v int32) *ListMetaProfilingResultCatalogPageResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) SetPageSize(v int32) *ListMetaProfilingResultCatalogPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) SetRequestId(v string) *ListMetaProfilingResultCatalogPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) SetSuccess(v bool) *ListMetaProfilingResultCatalogPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) SetTotalCount(v int32) *ListMetaProfilingResultCatalogPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMetaProfilingResultCatalogPageResponseBodyData struct {
	// example:
	//
	// catalog_hive_prod
	CatalogName *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	// example:
	//
	// 2024-07-19 22:00:02
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// 100256
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// 2024-07-19 22:00:02
	LastModified *string `json:"lastModified,omitempty" xml:"lastModified,omitempty"`
	// example:
	//
	// 3001
	MetaDatasourceId *int64 `json:"metaDatasourceId,omitempty" xml:"metaDatasourceId,omitempty"`
	// example:
	//
	// HIVE
	MetaDatasourceType *string `json:"metaDatasourceType,omitempty" xml:"metaDatasourceType,omitempty"`
	// example:
	//
	// 80
	MigrationProgress *float64 `json:"migrationProgress,omitempty" xml:"migrationProgress,omitempty"`
	// example:
	//
	// DONE
	MigrationStatus *string `json:"migrationStatus,omitempty" xml:"migrationStatus,omitempty"`
	// example:
	//
	// 2048
	PartitionCount *int32 `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	// example:
	//
	// 96
	PtTableCount *int32 `json:"ptTableCount,omitempty" xml:"ptTableCount,omitempty"`
	// example:
	//
	// 1000
	RecordCount *int64 `json:"recordCount,omitempty" xml:"recordCount,omitempty"`
	// example:
	//
	// 5
	SchemaCount *int32 `json:"schemaCount,omitempty" xml:"schemaCount,omitempty"`
	// example:
	//
	// 1.5
	StorageSize *string `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
	// example:
	//
	// 128
	TableCount *int32 `json:"tableCount,omitempty" xml:"tableCount,omitempty"`
	// example:
	//
	// 12
	ViewCount *int32 `json:"viewCount,omitempty" xml:"viewCount,omitempty"`
}

func (s ListMetaProfilingResultCatalogPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultCatalogPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetCreated() *string {
	return s.Created
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetJobId() *int64 {
	return s.JobId
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetLastModified() *string {
	return s.LastModified
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetMetaDatasourceId() *int64 {
	return s.MetaDatasourceId
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetMetaDatasourceType() *string {
	return s.MetaDatasourceType
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetMigrationProgress() *float64 {
	return s.MigrationProgress
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetMigrationStatus() *string {
	return s.MigrationStatus
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetPartitionCount() *int32 {
	return s.PartitionCount
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetPtTableCount() *int32 {
	return s.PtTableCount
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetSchemaCount() *int32 {
	return s.SchemaCount
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetStorageSize() *string {
	return s.StorageSize
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetTableCount() *int32 {
	return s.TableCount
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) GetViewCount() *int32 {
	return s.ViewCount
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetCatalogName(v string) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.CatalogName = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetCreated(v string) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.Created = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetJobId(v int64) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetLastModified(v string) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.LastModified = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetMetaDatasourceId(v int64) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.MetaDatasourceId = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetMetaDatasourceType(v string) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.MetaDatasourceType = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetMigrationProgress(v float64) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.MigrationProgress = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetMigrationStatus(v string) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.MigrationStatus = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetPartitionCount(v int32) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.PartitionCount = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetPtTableCount(v int32) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.PtTableCount = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetRecordCount(v int64) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.RecordCount = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetSchemaCount(v int32) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.SchemaCount = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetStorageSize(v string) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.StorageSize = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetTableCount(v int32) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.TableCount = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) SetViewCount(v int32) *ListMetaProfilingResultCatalogPageResponseBodyData {
	s.ViewCount = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
