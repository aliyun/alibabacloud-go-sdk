// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMigrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateMigrationTaskRequest
	GetAcceptLanguage() *string
	SetClusterType(v string) *UpdateMigrationTaskRequest
	GetClusterType() *string
	SetId(v string) *UpdateMigrationTaskRequest
	GetId() *string
	SetOriginInstanceAddress(v string) *UpdateMigrationTaskRequest
	GetOriginInstanceAddress() *string
	SetOriginInstanceName(v string) *UpdateMigrationTaskRequest
	GetOriginInstanceName() *string
	SetOriginInstanceNamespace(v string) *UpdateMigrationTaskRequest
	GetOriginInstanceNamespace() *string
	SetProjectDesc(v string) *UpdateMigrationTaskRequest
	GetProjectDesc() *string
	SetRequestPars(v string) *UpdateMigrationTaskRequest
	GetRequestPars() *string
	SetSyncType(v string) *UpdateMigrationTaskRequest
	GetSyncType() *string
	SetTargetClusterName(v string) *UpdateMigrationTaskRequest
	GetTargetClusterName() *string
	SetTargetClusterUrl(v string) *UpdateMigrationTaskRequest
	GetTargetClusterUrl() *string
	SetTargetInstanceId(v string) *UpdateMigrationTaskRequest
	GetTargetInstanceId() *string
}

type UpdateMigrationTaskRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The type of the instance. Valid values:
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
	// 192.168.1.1:8848
	OriginInstanceAddress *string `json:"OriginInstanceAddress,omitempty" xml:"OriginInstanceAddress,omitempty"`
	// The name of the source instance.
	//
	// example:
	//
	// Source instance
	OriginInstanceName *string `json:"OriginInstanceName,omitempty" xml:"OriginInstanceName,omitempty"`
	// The list of namespaces. This parameter is optional if you want to migrate applications from a Nacos instance.
	//
	// example:
	//
	// namesapceId1,namesapceId2
	OriginInstanceNamespace *string `json:"OriginInstanceNamespace,omitempty" xml:"OriginInstanceNamespace,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a description.
	ProjectDesc *string `json:"ProjectDesc,omitempty" xml:"ProjectDesc,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	SyncType    *string `json:"SyncType,omitempty" xml:"SyncType,omitempty"`
	// The name of the destination instance.
	//
	// example:
	//
	// Destination instance
	TargetClusterName *string `json:"TargetClusterName,omitempty" xml:"TargetClusterName,omitempty"`
	// The URL of the destination instance.
	//
	// example:
	//
	// mse-66*****-nacos-ans.mse.aliyuncs.com:8848
	TargetClusterUrl *string `json:"TargetClusterUrl,omitempty" xml:"TargetClusterUrl,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// mse-cn-ud82*****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
}

func (s UpdateMigrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMigrationTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateMigrationTaskRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateMigrationTaskRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *UpdateMigrationTaskRequest) GetId() *string {
	return s.Id
}

func (s *UpdateMigrationTaskRequest) GetOriginInstanceAddress() *string {
	return s.OriginInstanceAddress
}

func (s *UpdateMigrationTaskRequest) GetOriginInstanceName() *string {
	return s.OriginInstanceName
}

func (s *UpdateMigrationTaskRequest) GetOriginInstanceNamespace() *string {
	return s.OriginInstanceNamespace
}

func (s *UpdateMigrationTaskRequest) GetProjectDesc() *string {
	return s.ProjectDesc
}

func (s *UpdateMigrationTaskRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *UpdateMigrationTaskRequest) GetSyncType() *string {
	return s.SyncType
}

func (s *UpdateMigrationTaskRequest) GetTargetClusterName() *string {
	return s.TargetClusterName
}

func (s *UpdateMigrationTaskRequest) GetTargetClusterUrl() *string {
	return s.TargetClusterUrl
}

func (s *UpdateMigrationTaskRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *UpdateMigrationTaskRequest) SetAcceptLanguage(v string) *UpdateMigrationTaskRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetClusterType(v string) *UpdateMigrationTaskRequest {
	s.ClusterType = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetId(v string) *UpdateMigrationTaskRequest {
	s.Id = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetOriginInstanceAddress(v string) *UpdateMigrationTaskRequest {
	s.OriginInstanceAddress = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetOriginInstanceName(v string) *UpdateMigrationTaskRequest {
	s.OriginInstanceName = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetOriginInstanceNamespace(v string) *UpdateMigrationTaskRequest {
	s.OriginInstanceNamespace = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetProjectDesc(v string) *UpdateMigrationTaskRequest {
	s.ProjectDesc = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetRequestPars(v string) *UpdateMigrationTaskRequest {
	s.RequestPars = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetSyncType(v string) *UpdateMigrationTaskRequest {
	s.SyncType = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetTargetClusterName(v string) *UpdateMigrationTaskRequest {
	s.TargetClusterName = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetTargetClusterUrl(v string) *UpdateMigrationTaskRequest {
	s.TargetClusterUrl = &v
	return s
}

func (s *UpdateMigrationTaskRequest) SetTargetInstanceId(v string) *UpdateMigrationTaskRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *UpdateMigrationTaskRequest) Validate() error {
	return dara.Validate(s)
}
