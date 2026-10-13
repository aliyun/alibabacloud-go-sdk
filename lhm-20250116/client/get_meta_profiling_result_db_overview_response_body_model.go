// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaProfilingResultDbOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaProfilingResultDbOverviewResponseBodyData) *GetMetaProfilingResultDbOverviewResponseBody
	GetData() *GetMetaProfilingResultDbOverviewResponseBodyData
	SetErrCode(v string) *GetMetaProfilingResultDbOverviewResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetMetaProfilingResultDbOverviewResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetMetaProfilingResultDbOverviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaProfilingResultDbOverviewResponseBody
	GetSuccess() *bool
}

type GetMetaProfilingResultDbOverviewResponseBody struct {
	// The data source-level metadata profiling overview object. For more information about each field, see the descriptions of the sub-fields.
	Data *GetMetaProfilingResultDbOverviewResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code. It is an empty string if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. It is an empty string if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues related to this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. If the call fails, troubleshoot the issue based on the errCode and errMessage parameters.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMetaProfilingResultDbOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaProfilingResultDbOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaProfilingResultDbOverviewResponseBody) GetData() *GetMetaProfilingResultDbOverviewResponseBodyData {
	return s.Data
}

func (s *GetMetaProfilingResultDbOverviewResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetMetaProfilingResultDbOverviewResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetMetaProfilingResultDbOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaProfilingResultDbOverviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaProfilingResultDbOverviewResponseBody) SetData(v *GetMetaProfilingResultDbOverviewResponseBodyData) *GetMetaProfilingResultDbOverviewResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBody) SetErrCode(v string) *GetMetaProfilingResultDbOverviewResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBody) SetErrMessage(v string) *GetMetaProfilingResultDbOverviewResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBody) SetRequestId(v string) *GetMetaProfilingResultDbOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBody) SetSuccess(v bool) *GetMetaProfilingResultDbOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMetaProfilingResultDbOverviewResponseBodyData struct {
	// The number of catalogs.
	//
	// example:
	//
	// 1
	CatalogCount *int32 `json:"catalogCount,omitempty" xml:"catalogCount,omitempty"`
	// The component ID, which is the primary key identifier of the data source component.
	//
	// example:
	//
	// 12345
	ComponentId *int64 `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// The identifier of the data source component type.
	//
	// example:
	//
	// 1
	ComponentType *int32 `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// ds_demo
	DsName *string `json:"dsName,omitempty" xml:"dsName,omitempty"`
	// The version of the data source.
	//
	// example:
	//
	// 3.2.0
	DsVersion *string `json:"dsVersion,omitempty" xml:"dsVersion,omitempty"`
	// The number of external tables.
	//
	// example:
	//
	// 0
	EtTableCount *int32 `json:"etTableCount,omitempty" xml:"etTableCount,omitempty"`
	// The number of functions.
	//
	// example:
	//
	// 10
	FuncCount *int32 `json:"funcCount,omitempty" xml:"funcCount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-01-16T10:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2024-07-19 22:00:02
	LastModified *string `json:"lastModified,omitempty" xml:"lastModified,omitempty"`
	// The number of non-partitioned tables.
	//
	// example:
	//
	// 10
	NoPtTableCount *int32 `json:"noPtTableCount,omitempty" xml:"noPtTableCount,omitempty"`
	// The number of partitions.
	//
	// example:
	//
	// 10
	PartitionCount *int32 `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	// The number of partitioned tables.
	//
	// example:
	//
	// 10
	PtTableCount *int32 `json:"ptTableCount,omitempty" xml:"ptTableCount,omitempty"`
	// The number of data rows.
	//
	// example:
	//
	// 1000
	RecordCount *int64 `json:"recordCount,omitempty" xml:"recordCount,omitempty"`
	// The number of schemas.
	//
	// example:
	//
	// 10
	SchemaCount *int32 `json:"schemaCount,omitempty" xml:"schemaCount,omitempty"`
	// The data storage size.
	//
	// example:
	//
	// 1.5
	StorageSize *string `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
	// The number of tables.
	//
	// example:
	//
	// 100
	TableCount *int32 `json:"tableCount,omitempty" xml:"tableCount,omitempty"`
	// The number of transactional tables.
	//
	// example:
	//
	// 1
	TransactionalCount *int32 `json:"transactionalCount,omitempty" xml:"transactionalCount,omitempty"`
	// The number of views.
	//
	// example:
	//
	// 10
	ViewCount *int32 `json:"viewCount,omitempty" xml:"viewCount,omitempty"`
}

func (s GetMetaProfilingResultDbOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaProfilingResultDbOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetCatalogCount() *int32 {
	return s.CatalogCount
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetComponentId() *int64 {
	return s.ComponentId
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetComponentType() *int32 {
	return s.ComponentType
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetDsName() *string {
	return s.DsName
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetDsVersion() *string {
	return s.DsVersion
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetEtTableCount() *int32 {
	return s.EtTableCount
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetFuncCount() *int32 {
	return s.FuncCount
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetLastModified() *string {
	return s.LastModified
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetNoPtTableCount() *int32 {
	return s.NoPtTableCount
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetPartitionCount() *int32 {
	return s.PartitionCount
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetPtTableCount() *int32 {
	return s.PtTableCount
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetSchemaCount() *int32 {
	return s.SchemaCount
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetStorageSize() *string {
	return s.StorageSize
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetTableCount() *int32 {
	return s.TableCount
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetTransactionalCount() *int32 {
	return s.TransactionalCount
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) GetViewCount() *int32 {
	return s.ViewCount
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetCatalogCount(v int32) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.CatalogCount = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetComponentId(v int64) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.ComponentId = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetComponentType(v int32) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.ComponentType = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetDsName(v string) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.DsName = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetDsVersion(v string) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.DsVersion = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetEtTableCount(v int32) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.EtTableCount = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetFuncCount(v int32) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.FuncCount = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetGmtCreate(v string) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetLastModified(v string) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.LastModified = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetNoPtTableCount(v int32) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.NoPtTableCount = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetPartitionCount(v int32) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.PartitionCount = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetPtTableCount(v int32) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.PtTableCount = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetRecordCount(v int64) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.RecordCount = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetSchemaCount(v int32) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.SchemaCount = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetStorageSize(v string) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.StorageSize = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetTableCount(v int32) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.TableCount = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetTransactionalCount(v int32) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.TransactionalCount = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) SetViewCount(v int32) *GetMetaProfilingResultDbOverviewResponseBodyData {
	s.ViewCount = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
