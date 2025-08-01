// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMigrationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateMigrationTaskResponseBodyData) *UpdateMigrationTaskResponseBody
	GetData() *UpdateMigrationTaskResponseBodyData
	SetErrorCode(v string) *UpdateMigrationTaskResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *UpdateMigrationTaskResponseBody
	GetHttpCode() *string
	SetMessage(v string) *UpdateMigrationTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMigrationTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMigrationTaskResponseBody
	GetSuccess() *bool
}

type UpdateMigrationTaskResponseBody struct {
	// The data structure.
	Data *UpdateMigrationTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The ID of the request.
	//
	// example:
	//
	// AF21683A-29C7-4853-AC0F-B5ADEE4****
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
}

func (s UpdateMigrationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMigrationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMigrationTaskResponseBody) GetData() *UpdateMigrationTaskResponseBodyData {
	return s.Data
}

func (s *UpdateMigrationTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateMigrationTaskResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *UpdateMigrationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMigrationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMigrationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMigrationTaskResponseBody) SetData(v *UpdateMigrationTaskResponseBodyData) *UpdateMigrationTaskResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMigrationTaskResponseBody) SetErrorCode(v string) *UpdateMigrationTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMigrationTaskResponseBody) SetHttpCode(v string) *UpdateMigrationTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateMigrationTaskResponseBody) SetMessage(v string) *UpdateMigrationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMigrationTaskResponseBody) SetRequestId(v string) *UpdateMigrationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMigrationTaskResponseBody) SetSuccess(v bool) *UpdateMigrationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMigrationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateMigrationTaskResponseBodyData struct {
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
	// The modification time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The address of the source instance node.
	//
	// example:
	//
	// 192.168.100.2:2181
	OriginInstanceAddress *string `json:"OriginInstanceAddress,omitempty" xml:"OriginInstanceAddress,omitempty"`
	// The name of the source instance.
	//
	// example:
	//
	// src
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
	// 1232345
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
	// mse-0b*****-nacos-ans.mse.aliyuncs.com:8848
	TargetClusterUrl *string `json:"TargetClusterUrl,omitempty" xml:"TargetClusterUrl,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// mse-cn-zvp2u*****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 183876217*****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateMigrationTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMigrationTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMigrationTaskResponseBodyData) GetClusterType() *string {
	return s.ClusterType
}

func (s *UpdateMigrationTaskResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateMigrationTaskResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateMigrationTaskResponseBodyData) GetId() *string {
	return s.Id
}

func (s *UpdateMigrationTaskResponseBodyData) GetOriginInstanceAddress() *string {
	return s.OriginInstanceAddress
}

func (s *UpdateMigrationTaskResponseBodyData) GetOriginInstanceName() *string {
	return s.OriginInstanceName
}

func (s *UpdateMigrationTaskResponseBodyData) GetOriginInstanceNamespace() *string {
	return s.OriginInstanceNamespace
}

func (s *UpdateMigrationTaskResponseBodyData) GetProjectDesc() *string {
	return s.ProjectDesc
}

func (s *UpdateMigrationTaskResponseBodyData) GetSyncType() *string {
	return s.SyncType
}

func (s *UpdateMigrationTaskResponseBodyData) GetTargetClusterName() *string {
	return s.TargetClusterName
}

func (s *UpdateMigrationTaskResponseBodyData) GetTargetClusterUrl() *string {
	return s.TargetClusterUrl
}

func (s *UpdateMigrationTaskResponseBodyData) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *UpdateMigrationTaskResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *UpdateMigrationTaskResponseBodyData) SetClusterType(v string) *UpdateMigrationTaskResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetGmtCreate(v string) *UpdateMigrationTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetGmtModified(v string) *UpdateMigrationTaskResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetId(v string) *UpdateMigrationTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetOriginInstanceAddress(v string) *UpdateMigrationTaskResponseBodyData {
	s.OriginInstanceAddress = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetOriginInstanceName(v string) *UpdateMigrationTaskResponseBodyData {
	s.OriginInstanceName = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetOriginInstanceNamespace(v string) *UpdateMigrationTaskResponseBodyData {
	s.OriginInstanceNamespace = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetProjectDesc(v string) *UpdateMigrationTaskResponseBodyData {
	s.ProjectDesc = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetSyncType(v string) *UpdateMigrationTaskResponseBodyData {
	s.SyncType = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetTargetClusterName(v string) *UpdateMigrationTaskResponseBodyData {
	s.TargetClusterName = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetTargetClusterUrl(v string) *UpdateMigrationTaskResponseBodyData {
	s.TargetClusterUrl = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetTargetInstanceId(v string) *UpdateMigrationTaskResponseBodyData {
	s.TargetInstanceId = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) SetUserId(v string) *UpdateMigrationTaskResponseBodyData {
	s.UserId = &v
	return s
}

func (s *UpdateMigrationTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
