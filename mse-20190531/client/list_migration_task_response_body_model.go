// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListMigrationTaskResponseBodyData) *ListMigrationTaskResponseBody
	GetData() []*ListMigrationTaskResponseBodyData
	SetErrorCode(v string) *ListMigrationTaskResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListMigrationTaskResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListMigrationTaskResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *ListMigrationTaskResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListMigrationTaskResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListMigrationTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMigrationTaskResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListMigrationTaskResponseBody
	GetTotalCount() *int64
}

type ListMigrationTaskResponseBody struct {
	// The array structure.
	Data []*ListMigrationTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 451EBE59-5F33-5B15-83C1-78593B9*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMigrationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListMigrationTaskResponseBody) GetData() []*ListMigrationTaskResponseBodyData {
	return s.Data
}

func (s *ListMigrationTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMigrationTaskResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListMigrationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMigrationTaskResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListMigrationTaskResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListMigrationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMigrationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMigrationTaskResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMigrationTaskResponseBody) SetData(v []*ListMigrationTaskResponseBodyData) *ListMigrationTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListMigrationTaskResponseBody) SetErrorCode(v string) *ListMigrationTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMigrationTaskResponseBody) SetHttpCode(v string) *ListMigrationTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListMigrationTaskResponseBody) SetMessage(v string) *ListMigrationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListMigrationTaskResponseBody) SetPageNumber(v int64) *ListMigrationTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListMigrationTaskResponseBody) SetPageSize(v int64) *ListMigrationTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMigrationTaskResponseBody) SetRequestId(v string) *ListMigrationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMigrationTaskResponseBody) SetSuccess(v bool) *ListMigrationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ListMigrationTaskResponseBody) SetTotalCount(v int64) *ListMigrationTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMigrationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMigrationTaskResponseBodyData struct {
	// The type of the instance.
	//
	// 	- Nacos-Ans
	//
	// 	- ZooKeeper
	//
	// 	- Eureka
	//
	// example:
	//
	// Nacos-Ans
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2021-12-30T06:41:52.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the job.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The address of the source instance node.
	//
	// example:
	//
	// 192.168.1.1:8848
	OriginInstanceAddress *string `json:"OriginInstanceAddress,omitempty" xml:"OriginInstanceAddress,omitempty"`
	// The name of the source instance.
	//
	// example:
	//
	// test
	OriginInstanceName *string `json:"OriginInstanceName,omitempty" xml:"OriginInstanceName,omitempty"`
	// The list of namespaces. This parameter is optional if applications are migrated from a Nacos instance.
	//
	// example:
	//
	// fsdfsdfdsf
	OriginInstanceNamespace *string `json:"OriginInstanceNamespace,omitempty" xml:"OriginInstanceNamespace,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	ProjectDesc *string `json:"ProjectDesc,omitempty" xml:"ProjectDesc,omitempty"`
	SyncType    *string `json:"SyncType,omitempty" xml:"SyncType,omitempty"`
	// The name of the destination instance.
	//
	// example:
	//
	// multiple-nacos
	TargetClusterName *string `json:"TargetClusterName,omitempty" xml:"TargetClusterName,omitempty"`
	// The URL of the destination instance.
	//
	// example:
	//
	// mse-f1******-nacos-ans.mse.aliyuncs.com:8848
	TargetClusterUrl *string `json:"TargetClusterUrl,omitempty" xml:"TargetClusterUrl,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// mse-cn-zv*****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 2
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMigrationTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMigrationTaskResponseBodyData) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListMigrationTaskResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMigrationTaskResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMigrationTaskResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListMigrationTaskResponseBodyData) GetOriginInstanceAddress() *string {
	return s.OriginInstanceAddress
}

func (s *ListMigrationTaskResponseBodyData) GetOriginInstanceName() *string {
	return s.OriginInstanceName
}

func (s *ListMigrationTaskResponseBodyData) GetOriginInstanceNamespace() *string {
	return s.OriginInstanceNamespace
}

func (s *ListMigrationTaskResponseBodyData) GetProjectDesc() *string {
	return s.ProjectDesc
}

func (s *ListMigrationTaskResponseBodyData) GetSyncType() *string {
	return s.SyncType
}

func (s *ListMigrationTaskResponseBodyData) GetTargetClusterName() *string {
	return s.TargetClusterName
}

func (s *ListMigrationTaskResponseBodyData) GetTargetClusterUrl() *string {
	return s.TargetClusterUrl
}

func (s *ListMigrationTaskResponseBodyData) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *ListMigrationTaskResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListMigrationTaskResponseBodyData) SetClusterType(v string) *ListMigrationTaskResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetGmtCreate(v string) *ListMigrationTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetGmtModified(v string) *ListMigrationTaskResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetId(v int64) *ListMigrationTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetOriginInstanceAddress(v string) *ListMigrationTaskResponseBodyData {
	s.OriginInstanceAddress = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetOriginInstanceName(v string) *ListMigrationTaskResponseBodyData {
	s.OriginInstanceName = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetOriginInstanceNamespace(v string) *ListMigrationTaskResponseBodyData {
	s.OriginInstanceNamespace = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetProjectDesc(v string) *ListMigrationTaskResponseBodyData {
	s.ProjectDesc = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetSyncType(v string) *ListMigrationTaskResponseBodyData {
	s.SyncType = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetTargetClusterName(v string) *ListMigrationTaskResponseBodyData {
	s.TargetClusterName = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetTargetClusterUrl(v string) *ListMigrationTaskResponseBodyData {
	s.TargetClusterUrl = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetTargetInstanceId(v string) *ListMigrationTaskResponseBodyData {
	s.TargetInstanceId = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) SetUserId(v string) *ListMigrationTaskResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListMigrationTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
