// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMetaTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SearchMetaTablesResponseBodyData) *SearchMetaTablesResponseBody
	GetData() *SearchMetaTablesResponseBodyData
	SetErrorCode(v string) *SearchMetaTablesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SearchMetaTablesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *SearchMetaTablesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SearchMetaTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchMetaTablesResponseBody
	GetSuccess() *bool
}

type SearchMetaTablesResponseBody struct {
	// The search results.
	Data *SearchMetaTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchMetaTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMetaTablesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMetaTablesResponseBody) GetData() *SearchMetaTablesResponseBodyData {
	return s.Data
}

func (s *SearchMetaTablesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchMetaTablesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchMetaTablesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SearchMetaTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMetaTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchMetaTablesResponseBody) SetData(v *SearchMetaTablesResponseBodyData) *SearchMetaTablesResponseBody {
	s.Data = v
	return s
}

func (s *SearchMetaTablesResponseBody) SetErrorCode(v string) *SearchMetaTablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchMetaTablesResponseBody) SetErrorMessage(v string) *SearchMetaTablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchMetaTablesResponseBody) SetHttpStatusCode(v int32) *SearchMetaTablesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchMetaTablesResponseBody) SetRequestId(v string) *SearchMetaTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMetaTablesResponseBody) SetSuccess(v bool) *SearchMetaTablesResponseBody {
	s.Success = &v
	return s
}

func (s *SearchMetaTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchMetaTablesResponseBodyData struct {
	// The list of metatables.
	DataEntityList []*SearchMetaTablesResponseBodyDataDataEntityList `json:"DataEntityList,omitempty" xml:"DataEntityList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of metatables.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchMetaTablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchMetaTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchMetaTablesResponseBodyData) GetDataEntityList() []*SearchMetaTablesResponseBodyDataDataEntityList {
	return s.DataEntityList
}

func (s *SearchMetaTablesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchMetaTablesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchMetaTablesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchMetaTablesResponseBodyData) SetDataEntityList(v []*SearchMetaTablesResponseBodyDataDataEntityList) *SearchMetaTablesResponseBodyData {
	s.DataEntityList = v
	return s
}

func (s *SearchMetaTablesResponseBodyData) SetPageNumber(v int32) *SearchMetaTablesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *SearchMetaTablesResponseBodyData) SetPageSize(v int32) *SearchMetaTablesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *SearchMetaTablesResponseBodyData) SetTotalCount(v int64) *SearchMetaTablesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *SearchMetaTablesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type SearchMetaTablesResponseBodyDataDataEntityList struct {
	// The ID of the EMR cluster.
	//
	// example:
	//
	// abc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the metadatabase.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The type of the metatable. Valid values:
	//
	// 	- 0: table
	//
	// 	- 1: view
	//
	// example:
	//
	// 0
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The type of the environment. Valid values:
	//
	// 	- 1: production environment
	//
	// 	- 0: development environment
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The ID of the Alibaba Cloud account used by the workspace owner.
	//
	// example:
	//
	// 123
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 323
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The schema information of the table. You must configure this parameter if you enable the three-layer model of MaxCompute.
	//
	// example:
	//
	// default
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The GUID of the metatable.
	//
	// example:
	//
	// odps.engine_name.test_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the metatable.
	//
	// example:
	//
	// test_name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 12345
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s SearchMetaTablesResponseBodyDataDataEntityList) String() string {
	return dara.Prettify(s)
}

func (s SearchMetaTablesResponseBodyDataDataEntityList) GoString() string {
	return s.String()
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) GetClusterId() *string {
	return s.ClusterId
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) GetEntityType() *int32 {
	return s.EntityType
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) GetEnvType() *int32 {
	return s.EnvType
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) GetProjectName() *string {
	return s.ProjectName
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) GetSchema() *string {
	return s.Schema
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) GetTableGuid() *string {
	return s.TableGuid
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) GetTableName() *string {
	return s.TableName
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) GetTenantId() *int64 {
	return s.TenantId
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) SetClusterId(v string) *SearchMetaTablesResponseBodyDataDataEntityList {
	s.ClusterId = &v
	return s
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) SetDatabaseName(v string) *SearchMetaTablesResponseBodyDataDataEntityList {
	s.DatabaseName = &v
	return s
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) SetEntityType(v int32) *SearchMetaTablesResponseBodyDataDataEntityList {
	s.EntityType = &v
	return s
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) SetEnvType(v int32) *SearchMetaTablesResponseBodyDataDataEntityList {
	s.EnvType = &v
	return s
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) SetOwnerId(v string) *SearchMetaTablesResponseBodyDataDataEntityList {
	s.OwnerId = &v
	return s
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) SetProjectId(v int64) *SearchMetaTablesResponseBodyDataDataEntityList {
	s.ProjectId = &v
	return s
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) SetProjectName(v string) *SearchMetaTablesResponseBodyDataDataEntityList {
	s.ProjectName = &v
	return s
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) SetSchema(v string) *SearchMetaTablesResponseBodyDataDataEntityList {
	s.Schema = &v
	return s
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) SetTableGuid(v string) *SearchMetaTablesResponseBodyDataDataEntityList {
	s.TableGuid = &v
	return s
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) SetTableName(v string) *SearchMetaTablesResponseBodyDataDataEntityList {
	s.TableName = &v
	return s
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) SetTenantId(v int64) *SearchMetaTablesResponseBodyDataDataEntityList {
	s.TenantId = &v
	return s
}

func (s *SearchMetaTablesResponseBodyDataDataEntityList) Validate() error {
	return dara.Validate(s)
}
