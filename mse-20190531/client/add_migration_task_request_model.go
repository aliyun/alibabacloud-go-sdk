// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMigrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddMigrationTaskRequest
	GetAcceptLanguage() *string
	SetClusterType(v string) *AddMigrationTaskRequest
	GetClusterType() *string
	SetOriginInstanceAddress(v string) *AddMigrationTaskRequest
	GetOriginInstanceAddress() *string
	SetOriginInstanceName(v string) *AddMigrationTaskRequest
	GetOriginInstanceName() *string
	SetOriginInstanceNamespace(v string) *AddMigrationTaskRequest
	GetOriginInstanceNamespace() *string
	SetProjectDesc(v string) *AddMigrationTaskRequest
	GetProjectDesc() *string
	SetRequestPars(v string) *AddMigrationTaskRequest
	GetRequestPars() *string
	SetSyncType(v string) *AddMigrationTaskRequest
	GetSyncType() *string
	SetTargetClusterName(v string) *AddMigrationTaskRequest
	GetTargetClusterName() *string
	SetTargetClusterUrl(v string) *AddMigrationTaskRequest
	GetTargetClusterUrl() *string
	SetTargetInstanceId(v string) *AddMigrationTaskRequest
	GetTargetInstanceId() *string
}

type AddMigrationTaskRequest struct {
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The cluster type. Valid values:
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
	// Source cluster
	OriginInstanceName *string `json:"OriginInstanceName,omitempty" xml:"OriginInstanceName,omitempty"`
	// The list of namespaces. This parameter is optional when the source cluster is Nacos.
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
	// The extended request parameters in JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The synchronization type. Valid values: configuration synchronization, service synchronization, or full synchronization.
	//
	// example:
	//
	// Service
	SyncType *string `json:"SyncType,omitempty" xml:"SyncType,omitempty"`
	// The name of the target instance.
	//
	// example:
	//
	// Target cluster
	TargetClusterName *string `json:"TargetClusterName,omitempty" xml:"TargetClusterName,omitempty"`
	// The URL of the target instance.
	//
	// example:
	//
	// mse-66*****-nacos-ans.mse.aliyuncs.com:8848
	TargetClusterUrl *string `json:"TargetClusterUrl,omitempty" xml:"TargetClusterUrl,omitempty"`
	// The instance ID of the target instance.
	//
	// example:
	//
	// mse-cn-ud82*****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
}

func (s AddMigrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMigrationTaskRequest) GoString() string {
	return s.String()
}

func (s *AddMigrationTaskRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddMigrationTaskRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *AddMigrationTaskRequest) GetOriginInstanceAddress() *string {
	return s.OriginInstanceAddress
}

func (s *AddMigrationTaskRequest) GetOriginInstanceName() *string {
	return s.OriginInstanceName
}

func (s *AddMigrationTaskRequest) GetOriginInstanceNamespace() *string {
	return s.OriginInstanceNamespace
}

func (s *AddMigrationTaskRequest) GetProjectDesc() *string {
	return s.ProjectDesc
}

func (s *AddMigrationTaskRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *AddMigrationTaskRequest) GetSyncType() *string {
	return s.SyncType
}

func (s *AddMigrationTaskRequest) GetTargetClusterName() *string {
	return s.TargetClusterName
}

func (s *AddMigrationTaskRequest) GetTargetClusterUrl() *string {
	return s.TargetClusterUrl
}

func (s *AddMigrationTaskRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *AddMigrationTaskRequest) SetAcceptLanguage(v string) *AddMigrationTaskRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddMigrationTaskRequest) SetClusterType(v string) *AddMigrationTaskRequest {
	s.ClusterType = &v
	return s
}

func (s *AddMigrationTaskRequest) SetOriginInstanceAddress(v string) *AddMigrationTaskRequest {
	s.OriginInstanceAddress = &v
	return s
}

func (s *AddMigrationTaskRequest) SetOriginInstanceName(v string) *AddMigrationTaskRequest {
	s.OriginInstanceName = &v
	return s
}

func (s *AddMigrationTaskRequest) SetOriginInstanceNamespace(v string) *AddMigrationTaskRequest {
	s.OriginInstanceNamespace = &v
	return s
}

func (s *AddMigrationTaskRequest) SetProjectDesc(v string) *AddMigrationTaskRequest {
	s.ProjectDesc = &v
	return s
}

func (s *AddMigrationTaskRequest) SetRequestPars(v string) *AddMigrationTaskRequest {
	s.RequestPars = &v
	return s
}

func (s *AddMigrationTaskRequest) SetSyncType(v string) *AddMigrationTaskRequest {
	s.SyncType = &v
	return s
}

func (s *AddMigrationTaskRequest) SetTargetClusterName(v string) *AddMigrationTaskRequest {
	s.TargetClusterName = &v
	return s
}

func (s *AddMigrationTaskRequest) SetTargetClusterUrl(v string) *AddMigrationTaskRequest {
	s.TargetClusterUrl = &v
	return s
}

func (s *AddMigrationTaskRequest) SetTargetInstanceId(v string) *AddMigrationTaskRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *AddMigrationTaskRequest) Validate() error {
	return dara.Validate(s)
}
