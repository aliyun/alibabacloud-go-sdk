// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMigrationsResponseBodyData) *ListMigrationsResponseBody
	GetData() *ListMigrationsResponseBodyData
	SetRequestId(v string) *ListMigrationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMigrationsResponseBody
	GetSuccess() *bool
}

type ListMigrationsResponseBody struct {
	// The returned data.
	Data *ListMigrationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F9198AA3-9010-53D5-9714-DC4461427D3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMigrationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMigrationsResponseBody) GetData() *ListMigrationsResponseBodyData {
	return s.Data
}

func (s *ListMigrationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMigrationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMigrationsResponseBody) SetData(v *ListMigrationsResponseBodyData) *ListMigrationsResponseBody {
	s.Data = v
	return s
}

func (s *ListMigrationsResponseBody) SetRequestId(v string) *ListMigrationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMigrationsResponseBody) SetSuccess(v bool) *ListMigrationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMigrationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMigrationsResponseBodyData struct {
	// The list of migration tasks.
	Migrations []*ListMigrationsResponseBodyDataMigrations `json:"Migrations,omitempty" xml:"Migrations,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMigrationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMigrationsResponseBodyData) GetMigrations() []*ListMigrationsResponseBodyDataMigrations {
	return s.Migrations
}

func (s *ListMigrationsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMigrationsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMigrationsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMigrationsResponseBodyData) SetMigrations(v []*ListMigrationsResponseBodyDataMigrations) *ListMigrationsResponseBodyData {
	s.Migrations = v
	return s
}

func (s *ListMigrationsResponseBodyData) SetPageNumber(v int32) *ListMigrationsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMigrationsResponseBodyData) SetPageSize(v int32) *ListMigrationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMigrationsResponseBodyData) SetTotalCount(v int32) *ListMigrationsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListMigrationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListMigrationsResponseBodyDataMigrations struct {
	// The time when the migration task was created.
	//
	// example:
	//
	// 123124123123123
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the migration task.
	//
	// example:
	//
	// 123123****
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The name of the user who created the migration task.
	//
	// example:
	//
	// 3h1_test
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// The description of the export task.
	//
	// example:
	//
	// Automated Test creation
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL that is used to download the package of the export task. You can use the URL to download the package of the export task.
	//
	// example:
	//
	// http://geoip-sdk-user.oss-cn-zhangjiakou.aliyuncs.com/product/v1/ipv4/trace/v1.20220424123842.dex?Expires=1650780849&OSSAccessKeyId=XXXXXXeF4Lv5j&Signature=qxxxxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// 436064
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// An internal system error occurred. datasource kafka region is not cn-chengdu, can\\"t open network for it
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The type of the migration task. Valid values:
	//
	// 	- IMPORT
	//
	// 	- EXPORT
	//
	// example:
	//
	// EXPORT
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The name of the migration task.
	//
	// example:
	//
	// test_task_1638498642279
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the import or export package. Valid values:
	//
	// 	- DWMA (standard format)
	//
	// 	- DATAWORKS_MODEL (standard format)
	//
	// 	- DATAWORKS_V2 (Apsara Stack DataWorks V3.6.1 to V3.11)
	//
	// 	- DATAWORKS_V3 (Apsara Stack DataWorks V3.12 and later)
	//
	// The DWMA and DATAWORKS_MODEL types are interchangeable.
	//
	// example:
	//
	// DATAWORKS_MODEL
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The ID of the DataWorks workspace to which the task belongs.
	//
	// example:
	//
	// 72132
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the migration task. Valid values:
	//
	// 	- INIT: The migration task is in the initial state.
	//
	// 	- EDITING: The migration task is being edited.
	//
	// 	- RUNNING: The migration task is running.
	//
	// 	- FAILURE: The migration task fails to run.
	//
	// 	- SUCCESS: The migration task is successfully run.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 16307
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The time when the migration task was last updated.
	//
	// example:
	//
	// 123123123123123
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user who last updated the migration task.
	//
	// example:
	//
	// 1231****
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	// The name of the user who last updated the migration task.
	//
	// example:
	//
	// 3h1_test
	UpdateUserName *string `json:"UpdateUserName,omitempty" xml:"UpdateUserName,omitempty"`
}

func (s ListMigrationsResponseBodyDataMigrations) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsResponseBodyDataMigrations) GoString() string {
	return s.String()
}

func (s *ListMigrationsResponseBodyDataMigrations) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListMigrationsResponseBodyDataMigrations) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListMigrationsResponseBodyDataMigrations) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListMigrationsResponseBodyDataMigrations) GetDescription() *string {
	return s.Description
}

func (s *ListMigrationsResponseBodyDataMigrations) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListMigrationsResponseBodyDataMigrations) GetId() *int64 {
	return s.Id
}

func (s *ListMigrationsResponseBodyDataMigrations) GetMessage() *string {
	return s.Message
}

func (s *ListMigrationsResponseBodyDataMigrations) GetMigrationType() *string {
	return s.MigrationType
}

func (s *ListMigrationsResponseBodyDataMigrations) GetName() *string {
	return s.Name
}

func (s *ListMigrationsResponseBodyDataMigrations) GetPackageType() *string {
	return s.PackageType
}

func (s *ListMigrationsResponseBodyDataMigrations) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListMigrationsResponseBodyDataMigrations) GetStatus() *string {
	return s.Status
}

func (s *ListMigrationsResponseBodyDataMigrations) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListMigrationsResponseBodyDataMigrations) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListMigrationsResponseBodyDataMigrations) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *ListMigrationsResponseBodyDataMigrations) GetUpdateUserName() *string {
	return s.UpdateUserName
}

func (s *ListMigrationsResponseBodyDataMigrations) SetCreateTime(v int64) *ListMigrationsResponseBodyDataMigrations {
	s.CreateTime = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetCreateUser(v string) *ListMigrationsResponseBodyDataMigrations {
	s.CreateUser = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetCreateUserName(v string) *ListMigrationsResponseBodyDataMigrations {
	s.CreateUserName = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetDescription(v string) *ListMigrationsResponseBodyDataMigrations {
	s.Description = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetDownloadUrl(v string) *ListMigrationsResponseBodyDataMigrations {
	s.DownloadUrl = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetId(v int64) *ListMigrationsResponseBodyDataMigrations {
	s.Id = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetMessage(v string) *ListMigrationsResponseBodyDataMigrations {
	s.Message = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetMigrationType(v string) *ListMigrationsResponseBodyDataMigrations {
	s.MigrationType = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetName(v string) *ListMigrationsResponseBodyDataMigrations {
	s.Name = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetPackageType(v string) *ListMigrationsResponseBodyDataMigrations {
	s.PackageType = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetProjectId(v int64) *ListMigrationsResponseBodyDataMigrations {
	s.ProjectId = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetStatus(v string) *ListMigrationsResponseBodyDataMigrations {
	s.Status = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetTenantId(v int64) *ListMigrationsResponseBodyDataMigrations {
	s.TenantId = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetUpdateTime(v int64) *ListMigrationsResponseBodyDataMigrations {
	s.UpdateTime = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetUpdateUser(v string) *ListMigrationsResponseBodyDataMigrations {
	s.UpdateUser = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) SetUpdateUserName(v string) *ListMigrationsResponseBodyDataMigrations {
	s.UpdateUserName = &v
	return s
}

func (s *ListMigrationsResponseBodyDataMigrations) Validate() error {
	return dara.Validate(s)
}
