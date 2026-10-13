// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingResultUdfPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListMetaProfilingResultUdfPageResponseBodyData) *ListMetaProfilingResultUdfPageResponseBody
	GetData() []*ListMetaProfilingResultUdfPageResponseBodyData
	SetErrCode(v string) *ListMetaProfilingResultUdfPageResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMetaProfilingResultUdfPageResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *ListMetaProfilingResultUdfPageResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMetaProfilingResultUdfPageResponseBody
	GetPageSize() *int32
	SetSuccess(v bool) *ListMetaProfilingResultUdfPageResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListMetaProfilingResultUdfPageResponseBody
	GetTotalCount() *int32
}

type ListMetaProfilingResultUdfPageResponseBody struct {
	Data []*ListMetaProfilingResultUdfPageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	Success  *bool  `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 128
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMetaProfilingResultUdfPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultUdfPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultUdfPageResponseBody) GetData() []*ListMetaProfilingResultUdfPageResponseBodyData {
	return s.Data
}

func (s *ListMetaProfilingResultUdfPageResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMetaProfilingResultUdfPageResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMetaProfilingResultUdfPageResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMetaProfilingResultUdfPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaProfilingResultUdfPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMetaProfilingResultUdfPageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMetaProfilingResultUdfPageResponseBody) SetData(v []*ListMetaProfilingResultUdfPageResponseBodyData) *ListMetaProfilingResultUdfPageResponseBody {
	s.Data = v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBody) SetErrCode(v string) *ListMetaProfilingResultUdfPageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBody) SetErrMessage(v string) *ListMetaProfilingResultUdfPageResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBody) SetPageIndex(v int32) *ListMetaProfilingResultUdfPageResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBody) SetPageSize(v int32) *ListMetaProfilingResultUdfPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBody) SetSuccess(v bool) *ListMetaProfilingResultUdfPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBody) SetTotalCount(v int32) *ListMetaProfilingResultUdfPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBody) Validate() error {
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

type ListMetaProfilingResultUdfPageResponseBodyData struct {
	// example:
	//
	// catalog_hive_prod
	CatalogName *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	// example:
	//
	// com.example.udf.ParseJson
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// example:
	//
	// ods
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// MYSQL
	DsType *string `json:"dsType,omitempty" xml:"dsType,omitempty"`
	// example:
	//
	// 解析 JSON 字符串的 UDF
	FuncComment *string `json:"funcComment,omitempty" xml:"funcComment,omitempty"`
	// example:
	//
	// 2024-07-19 22:00:02
	FuncCrtTime *string `json:"funcCrtTime,omitempty" xml:"funcCrtTime,omitempty"`
	// example:
	//
	// 6001
	FuncId *string `json:"funcId,omitempty" xml:"funcId,omitempty"`
	// example:
	//
	// 2024-07-20 17:52:57
	FuncModTime *string `json:"funcModTime,omitempty" xml:"funcModTime,omitempty"`
	// example:
	//
	// udf_parse_json
	FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty"`
	// example:
	//
	// UDF
	FuncType *string `json:"funcType,omitempty" xml:"funcType,omitempty"`
	// example:
	//
	// 7001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// DONE
	MigrationStatus *string `json:"migrationStatus,omitempty" xml:"migrationStatus,omitempty"`
	// example:
	//
	// zhangsan
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// oss://lhm-bucket/udf/parse-json.jar
	ResourceLocation *string `json:"resourceLocation,omitempty" xml:"resourceLocation,omitempty"`
	// example:
	//
	// JAR
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// 1009318965420000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListMetaProfilingResultUdfPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultUdfPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetClassName() *string {
	return s.ClassName
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetDbName() *string {
	return s.DbName
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetDsType() *string {
	return s.DsType
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetFuncComment() *string {
	return s.FuncComment
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetFuncCrtTime() *string {
	return s.FuncCrtTime
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetFuncId() *string {
	return s.FuncId
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetFuncModTime() *string {
	return s.FuncModTime
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetFuncName() *string {
	return s.FuncName
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetFuncType() *string {
	return s.FuncType
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetMigrationStatus() *string {
	return s.MigrationStatus
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetResourceLocation() *string {
	return s.ResourceLocation
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetCatalogName(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.CatalogName = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetClassName(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.ClassName = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetDbName(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.DbName = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetDsType(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.DsType = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetFuncComment(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.FuncComment = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetFuncCrtTime(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.FuncCrtTime = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetFuncId(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.FuncId = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetFuncModTime(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.FuncModTime = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetFuncName(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.FuncName = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetFuncType(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.FuncType = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetId(v int64) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetMigrationStatus(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.MigrationStatus = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetOwner(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.Owner = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetResourceLocation(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.ResourceLocation = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetResourceType(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) SetTenantId(v string) *ListMetaProfilingResultUdfPageResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
