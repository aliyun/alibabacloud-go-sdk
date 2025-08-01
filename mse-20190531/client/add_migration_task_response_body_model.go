// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMigrationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddMigrationTaskResponseBodyData) *AddMigrationTaskResponseBody
	GetData() *AddMigrationTaskResponseBodyData
	SetErrorCode(v string) *AddMigrationTaskResponseBody
	GetErrorCode() *string
	SetMessage(v string) *AddMigrationTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddMigrationTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddMigrationTaskResponseBody
	GetSuccess() *bool
}

type AddMigrationTaskResponseBody struct {
	// Data structure.
	Data *AddMigrationTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Message.
	//
	// example:
	//
	// The request is processed successfully.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 7466566F-F30F-4A29-965D-3E0AF21D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求结果，取值如下：
	//
	// - `true`：请求成功。
	//
	// - `false`：请求失败。
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMigrationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMigrationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AddMigrationTaskResponseBody) GetData() *AddMigrationTaskResponseBodyData {
	return s.Data
}

func (s *AddMigrationTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddMigrationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddMigrationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMigrationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddMigrationTaskResponseBody) SetData(v *AddMigrationTaskResponseBodyData) *AddMigrationTaskResponseBody {
	s.Data = v
	return s
}

func (s *AddMigrationTaskResponseBody) SetErrorCode(v string) *AddMigrationTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddMigrationTaskResponseBody) SetMessage(v string) *AddMigrationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *AddMigrationTaskResponseBody) SetRequestId(v string) *AddMigrationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMigrationTaskResponseBody) SetSuccess(v bool) *AddMigrationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *AddMigrationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddMigrationTaskResponseBodyData struct {
	// Cluster type.
	//
	// - Nacos-Ans
	//
	// - ZooKeeper
	//
	// - Eureka
	//
	// example:
	//
	// Nacos-Ans
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 12
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Source instance node address.
	//
	// example:
	//
	// 192.168.1.1:8848
	OriginInstanceAddress *string `json:"OriginInstanceAddress,omitempty" xml:"OriginInstanceAddress,omitempty"`
	// Source instance name.
	//
	// example:
	//
	// Source instance
	OriginInstanceName *string `json:"OriginInstanceName,omitempty" xml:"OriginInstanceName,omitempty"`
	// Namespace list, required when the source cluster is Nacos.
	//
	// example:
	//
	// namesapceId1,namesapceId2
	OriginInstanceNamespace *string `json:"OriginInstanceNamespace,omitempty" xml:"OriginInstanceNamespace,omitempty"`
	// Description.
	//
	// example:
	//
	// testsdfsdfsd
	ProjectDesc *string `json:"ProjectDesc,omitempty" xml:"ProjectDesc,omitempty"`
	// SyncType
	//
	// example:
	//
	// Service
	SyncType *string `json:"SyncType,omitempty" xml:"SyncType,omitempty"`
	// Target instance name.
	//
	// example:
	//
	// Destination instance
	TargetClusterName *string `json:"TargetClusterName,omitempty" xml:"TargetClusterName,omitempty"`
	// Target instance URL.
	//
	// example:
	//
	// mse-94d****-nacos-ans.mse.aliyuncs.com:8848
	TargetClusterUrl *string `json:"TargetClusterUrl,omitempty" xml:"TargetClusterUrl,omitempty"`
	// Target instance ID.
	//
	// example:
	//
	// mse-cn-7pp2w*****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// User ID.
	//
	// example:
	//
	// 183876217*****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddMigrationTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddMigrationTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddMigrationTaskResponseBodyData) GetClusterType() *string {
	return s.ClusterType
}

func (s *AddMigrationTaskResponseBodyData) GetId() *string {
	return s.Id
}

func (s *AddMigrationTaskResponseBodyData) GetOriginInstanceAddress() *string {
	return s.OriginInstanceAddress
}

func (s *AddMigrationTaskResponseBodyData) GetOriginInstanceName() *string {
	return s.OriginInstanceName
}

func (s *AddMigrationTaskResponseBodyData) GetOriginInstanceNamespace() *string {
	return s.OriginInstanceNamespace
}

func (s *AddMigrationTaskResponseBodyData) GetProjectDesc() *string {
	return s.ProjectDesc
}

func (s *AddMigrationTaskResponseBodyData) GetSyncType() *string {
	return s.SyncType
}

func (s *AddMigrationTaskResponseBodyData) GetTargetClusterName() *string {
	return s.TargetClusterName
}

func (s *AddMigrationTaskResponseBodyData) GetTargetClusterUrl() *string {
	return s.TargetClusterUrl
}

func (s *AddMigrationTaskResponseBodyData) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *AddMigrationTaskResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *AddMigrationTaskResponseBodyData) SetClusterType(v string) *AddMigrationTaskResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *AddMigrationTaskResponseBodyData) SetId(v string) *AddMigrationTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *AddMigrationTaskResponseBodyData) SetOriginInstanceAddress(v string) *AddMigrationTaskResponseBodyData {
	s.OriginInstanceAddress = &v
	return s
}

func (s *AddMigrationTaskResponseBodyData) SetOriginInstanceName(v string) *AddMigrationTaskResponseBodyData {
	s.OriginInstanceName = &v
	return s
}

func (s *AddMigrationTaskResponseBodyData) SetOriginInstanceNamespace(v string) *AddMigrationTaskResponseBodyData {
	s.OriginInstanceNamespace = &v
	return s
}

func (s *AddMigrationTaskResponseBodyData) SetProjectDesc(v string) *AddMigrationTaskResponseBodyData {
	s.ProjectDesc = &v
	return s
}

func (s *AddMigrationTaskResponseBodyData) SetSyncType(v string) *AddMigrationTaskResponseBodyData {
	s.SyncType = &v
	return s
}

func (s *AddMigrationTaskResponseBodyData) SetTargetClusterName(v string) *AddMigrationTaskResponseBodyData {
	s.TargetClusterName = &v
	return s
}

func (s *AddMigrationTaskResponseBodyData) SetTargetClusterUrl(v string) *AddMigrationTaskResponseBodyData {
	s.TargetClusterUrl = &v
	return s
}

func (s *AddMigrationTaskResponseBodyData) SetTargetInstanceId(v string) *AddMigrationTaskResponseBodyData {
	s.TargetInstanceId = &v
	return s
}

func (s *AddMigrationTaskResponseBodyData) SetUserId(v string) *AddMigrationTaskResponseBodyData {
	s.UserId = &v
	return s
}

func (s *AddMigrationTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
