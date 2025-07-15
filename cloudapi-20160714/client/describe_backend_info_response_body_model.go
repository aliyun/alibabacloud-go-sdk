// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackendInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendInfo(v *DescribeBackendInfoResponseBodyBackendInfo) *DescribeBackendInfoResponseBody
	GetBackendInfo() *DescribeBackendInfoResponseBodyBackendInfo
	SetRequestId(v string) *DescribeBackendInfoResponseBody
	GetRequestId() *string
}

type DescribeBackendInfoResponseBody struct {
	// The information about the backend service.
	BackendInfo *DescribeBackendInfoResponseBodyBackendInfo `json:"BackendInfo,omitempty" xml:"BackendInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0AA90E87-3506-5AA6-AFFB-A4D53B4F6231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackendInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBody) GetBackendInfo() *DescribeBackendInfoResponseBodyBackendInfo {
	return s.BackendInfo
}

func (s *DescribeBackendInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackendInfoResponseBody) SetBackendInfo(v *DescribeBackendInfoResponseBodyBackendInfo) *DescribeBackendInfoResponseBody {
	s.BackendInfo = v
	return s
}

func (s *DescribeBackendInfoResponseBody) SetRequestId(v string) *DescribeBackendInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackendInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfo struct {
	// The ID of the backend service.
	//
	// example:
	//
	// 6fc978bb63574146b766863dd7bdf661
	BackendId *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	// The configurations of the backend service in the environment.
	BackendModels []*DescribeBackendInfoResponseBodyBackendInfoBackendModels `json:"BackendModels,omitempty" xml:"BackendModels,omitempty" type:"Repeated"`
	// The name of the backend service.
	//
	// example:
	//
	// testoss2
	BackendName *string `json:"BackendName,omitempty" xml:"BackendName,omitempty"`
	// The type of the backend service.
	//
	// example:
	//
	// HTTP
	BackendType *string `json:"BackendType,omitempty" xml:"BackendType,omitempty"`
	// The time when the backend service was created.
	//
	// example:
	//
	// 2021-11-22T11:10:46+08:00
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the backend service.
	//
	// example:
	//
	// add
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the backend service was modified.
	//
	// example:
	//
	// 2017-12-11T15:18:09+08:00
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s DescribeBackendInfoResponseBodyBackendInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfo) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) GetBackendId() *string {
	return s.BackendId
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) GetBackendModels() []*DescribeBackendInfoResponseBodyBackendInfoBackendModels {
	return s.BackendModels
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) GetBackendName() *string {
	return s.BackendName
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) GetBackendType() *string {
	return s.BackendType
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) SetBackendId(v string) *DescribeBackendInfoResponseBodyBackendInfo {
	s.BackendId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) SetBackendModels(v []*DescribeBackendInfoResponseBodyBackendInfoBackendModels) *DescribeBackendInfoResponseBodyBackendInfo {
	s.BackendModels = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) SetBackendName(v string) *DescribeBackendInfoResponseBodyBackendInfo {
	s.BackendName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) SetBackendType(v string) *DescribeBackendInfoResponseBodyBackendInfo {
	s.BackendType = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) SetCreatedTime(v string) *DescribeBackendInfoResponseBodyBackendInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) SetDescription(v string) *DescribeBackendInfoResponseBodyBackendInfo {
	s.Description = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) SetModifiedTime(v string) *DescribeBackendInfoResponseBodyBackendInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModels struct {
	// The backend service configurations.
	BackendConfig *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig `json:"BackendConfig,omitempty" xml:"BackendConfig,omitempty" type:"Struct"`
	// The ID of the backend service in the environment.
	//
	// example:
	//
	// 5c4995d08e8b4954b0f326e8e4f2b97d
	BackendModelId *string `json:"BackendModelId,omitempty" xml:"BackendModelId,omitempty"`
	// The description of the backend service.
	//
	// example:
	//
	// testDvs 1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the backend service was created.
	//
	// example:
	//
	// 2021-12-20T03:22:03.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the backend service was modified.
	//
	// example:
	//
	// 2021-12-20T03:22:03.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the environment.
	//
	// example:
	//
	// 6fc978bb63574146b766863dd7bdf661
	StageModeId *string `json:"StageModeId,omitempty" xml:"StageModeId,omitempty"`
	// The environment name.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModels) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) GetBackendConfig() *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	return s.BackendConfig
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) GetBackendModelId() *string {
	return s.BackendModelId
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) GetDescription() *string {
	return s.Description
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) GetStageModeId() *string {
	return s.StageModeId
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) GetStageName() *string {
	return s.StageName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) SetBackendConfig(v *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) *DescribeBackendInfoResponseBodyBackendInfoBackendModels {
	s.BackendConfig = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) SetBackendModelId(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModels {
	s.BackendModelId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) SetDescription(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModels {
	s.Description = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) SetGmtCreate(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModels {
	s.GmtCreate = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) SetGmtModified(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModels {
	s.GmtModified = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) SetStageModeId(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModels {
	s.StageModeId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) SetStageName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModels {
	s.StageName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModels) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig struct {
	// The information about the backend service when the backend service is of the Service Discovery type.
	DiscoveryConfig *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig `json:"DiscoveryConfig,omitempty" xml:"DiscoveryConfig,omitempty" type:"Struct"`
	// The EDAS configuration.
	EdasConfig *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig `json:"EdasConfig,omitempty" xml:"EdasConfig,omitempty" type:"Struct"`
	// The information about the backend service whose type is EventBridge.
	EventBridgeConfig *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig `json:"EventBridgeConfig,omitempty" xml:"EventBridgeConfig,omitempty" type:"Struct"`
	// The information about the backend service whose type is Function Compute.
	FunctionComputeConfig *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty" type:"Struct"`
	// The host of the HTTP backend service.
	//
	// example:
	//
	// www.host.com
	HttpTargetHostName *string `json:"HttpTargetHostName,omitempty" xml:"HttpTargetHostName,omitempty"`
	// The information about the backend service when the backend service is of the Mock type.
	MockConfig *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig `json:"MockConfig,omitempty" xml:"MockConfig,omitempty" type:"Struct"`
	// The information about the backend service whose type is Object Storage Service (OSS).
	OssConfig *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig `json:"OssConfig,omitempty" xml:"OssConfig,omitempty" type:"Struct"`
	// The URL of the backend service.
	//
	// example:
	//
	// 10.0.0.1
	ServiceAddress *string `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	// The timeout period of the backend service, in millisecond.
	//
	// example:
	//
	// 10000
	ServiceTimeout *int32 `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	// The type of the backend service.
	//
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The information about the backend service when the backend service is of the VPC type.
	VpcConfig *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig `json:"VpcConfig,omitempty" xml:"VpcConfig,omitempty" type:"Struct"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GetDiscoveryConfig() *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig {
	return s.DiscoveryConfig
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GetEdasConfig() *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig {
	return s.EdasConfig
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GetEventBridgeConfig() *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig {
	return s.EventBridgeConfig
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GetFunctionComputeConfig() *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig {
	return s.FunctionComputeConfig
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GetHttpTargetHostName() *string {
	return s.HttpTargetHostName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GetMockConfig() *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig {
	return s.MockConfig
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GetOssConfig() *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig {
	return s.OssConfig
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GetServiceAddress() *string {
	return s.ServiceAddress
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GetServiceTimeout() *int32 {
	return s.ServiceTimeout
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GetType() *string {
	return s.Type
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) GetVpcConfig() *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig {
	return s.VpcConfig
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) SetDiscoveryConfig(v *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	s.DiscoveryConfig = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) SetEdasConfig(v *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	s.EdasConfig = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) SetEventBridgeConfig(v *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	s.EventBridgeConfig = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) SetFunctionComputeConfig(v *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	s.FunctionComputeConfig = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) SetHttpTargetHostName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	s.HttpTargetHostName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) SetMockConfig(v *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	s.MockConfig = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) SetOssConfig(v *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	s.OssConfig = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) SetServiceAddress(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	s.ServiceAddress = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) SetServiceTimeout(v int32) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	s.ServiceTimeout = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) SetType(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	s.Type = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) SetVpcConfig(v *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig {
	s.VpcConfig = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig struct {
	// The Nacos configurations.
	NacosConfig *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig `json:"NacosConfig,omitempty" xml:"NacosConfig,omitempty" type:"Struct"`
	// The registry type.
	//
	// example:
	//
	// NACOS
	RcType *string `json:"RcType,omitempty" xml:"RcType,omitempty"`
	// The ZooKeeper configuration.
	ZookeeperConfig *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig `json:"ZookeeperConfig,omitempty" xml:"ZookeeperConfig,omitempty" type:"Struct"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig) GetNacosConfig() *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig {
	return s.NacosConfig
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig) GetRcType() *string {
	return s.RcType
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig) GetZookeeperConfig() *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig {
	return s.ZookeeperConfig
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig) SetNacosConfig(v *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig {
	s.NacosConfig = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig) SetRcType(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig {
	s.RcType = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig) SetZookeeperConfig(v *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig {
	s.ZookeeperConfig = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig struct {
	// The AccessKey of the RAM user that has the resource management permissions on Microservices Engine (MSE).
	//
	// example:
	//
	// A5FIDxxxxxx
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The authentication method.
	//
	// example:
	//
	// PASSWORD
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The name of the cluster to which the microservice belongs.
	//
	// example:
	//
	// cluster1
	Clusters *string `json:"Clusters,omitempty" xml:"Clusters,omitempty"`
	// The name of the group to which the microservice that is registered with Nacos belongs.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the namespace where the microservice that is registered with Nacos resides.
	//
	// example:
	//
	// public
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password.
	//
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The SecretKey of the RAM user that has the resource management permissions on MSE.
	//
	// example:
	//
	// dl5loxxxxxx
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// The Nacos service address.
	//
	// example:
	//
	// http://1xx.2xx.3xx.4xx:8848
	ServerAddress *string `json:"ServerAddress,omitempty" xml:"ServerAddress,omitempty"`
	// The microservice name.
	//
	// example:
	//
	// service-provider
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The username.
	//
	// example:
	//
	// username
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) GetClusters() *string {
	return s.Clusters
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) GetPassword() *string {
	return s.Password
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) GetSecretKey() *string {
	return s.SecretKey
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) GetServerAddress() *string {
	return s.ServerAddress
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) GetUserName() *string {
	return s.UserName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) SetAccessKey(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig {
	s.AccessKey = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) SetAuthType(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig {
	s.AuthType = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) SetClusters(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig {
	s.Clusters = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) SetGroupName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig {
	s.GroupName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) SetNamespace(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig {
	s.Namespace = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) SetPassword(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig {
	s.Password = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) SetSecretKey(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig {
	s.SecretKey = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) SetServerAddress(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig {
	s.ServerAddress = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) SetServiceName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig {
	s.ServiceName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) SetUserName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig {
	s.UserName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigNacosConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig struct {
	// The connection URL of the ZooKeeper server.
	//
	// example:
	//
	// http://192.168.1.xxx:2181
	ConnectString *string `json:"ConnectString,omitempty" xml:"ConnectString,omitempty"`
	// The namespace.
	//
	// example:
	//
	// provider
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Service name
	//
	// example:
	//
	// service
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig) GetConnectString() *string {
	return s.ConnectString
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig) SetConnectString(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig {
	s.ConnectString = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig) SetNamespace(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig {
	s.Namespace = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig) SetServiceName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig {
	s.ServiceName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigDiscoveryConfigZookeeperConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig struct {
	// The EDAS application ID.
	//
	// example:
	//
	// 6cd0c599-dxxx-496d-b3d5-6a71c657xxxxx
	EdasAppId *string `json:"EdasAppId,omitempty" xml:"EdasAppId,omitempty"`
	// The ID of the microservices namespace in EDAS.
	//
	// example:
	//
	// cn-hangzhou:edasNacos
	MicroserviceNamespace *string `json:"MicroserviceNamespace,omitempty" xml:"MicroserviceNamespace,omitempty"`
	// The ID of the microservices namespace in EDAS.
	//
	// example:
	//
	// cn-hangzhou:edasNacos
	MicroserviceNamespaceId *string `json:"MicroserviceNamespaceId,omitempty" xml:"MicroserviceNamespaceId,omitempty"`
	// The name of the microservices namespace in EDAS.
	//
	// example:
	//
	// Edas-Nacos
	MicroserviceNamespaceName *string `json:"MicroserviceNamespaceName,omitempty" xml:"MicroserviceNamespaceName,omitempty"`
	// The MSE instance ID.
	//
	// example:
	//
	// mse-cn-jia3n1rxxxx
	MseInstanceId *string `json:"MseInstanceId,omitempty" xml:"MseInstanceId,omitempty"`
	// The registration type.
	//
	// example:
	//
	// EDAS
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The service name.
	//
	// example:
	//
	// service
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) GetEdasAppId() *string {
	return s.EdasAppId
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) GetMicroserviceNamespace() *string {
	return s.MicroserviceNamespace
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) GetMicroserviceNamespaceId() *string {
	return s.MicroserviceNamespaceId
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) GetMicroserviceNamespaceName() *string {
	return s.MicroserviceNamespaceName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) GetMseInstanceId() *string {
	return s.MseInstanceId
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) GetRegistryType() *string {
	return s.RegistryType
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) SetEdasAppId(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig {
	s.EdasAppId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) SetMicroserviceNamespace(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig {
	s.MicroserviceNamespace = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) SetMicroserviceNamespaceId(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig {
	s.MicroserviceNamespaceId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) SetMicroserviceNamespaceName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig {
	s.MicroserviceNamespaceName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) SetMseInstanceId(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig {
	s.MseInstanceId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) SetRegistryType(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig {
	s.RegistryType = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) SetServiceName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig {
	s.ServiceName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEdasConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig struct {
	// The region ID of the event bus in EventBridge.
	//
	// example:
	//
	// cn-hangzhou
	EventBridgeRegionId *string `json:"EventBridgeRegionId,omitempty" xml:"EventBridgeRegionId,omitempty"`
	// The event bus.
	//
	// example:
	//
	// testBus
	EventBus *string `json:"EventBus,omitempty" xml:"EventBus,omitempty"`
	// The event source.
	//
	// example:
	//
	// dds_driver
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The ARN of the RAM role to be assumed by API Gateway to access EventBridge.
	//
	// example:
	//
	// acs:ram::1975133748561***:role/aliyunserviceroleforiotlogexport
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) GetEventBridgeRegionId() *string {
	return s.EventBridgeRegionId
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) GetEventBus() *string {
	return s.EventBus
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) GetEventSource() *string {
	return s.EventSource
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) SetEventBridgeRegionId(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig {
	s.EventBridgeRegionId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) SetEventBus(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig {
	s.EventBus = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) SetEventSource(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig {
	s.EventSource = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) SetRoleArn(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig {
	s.RoleArn = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigEventBridgeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig struct {
	// The root path of the Function Compute service.
	//
	// example:
	//
	// https://t*******.ap-*****.fcapp.run/
	FcBaseUrl *string `json:"FcBaseUrl,omitempty" xml:"FcBaseUrl,omitempty"`
	// The region ID of the Function Compute service.
	//
	// example:
	//
	// cn-hangzhou
	FcRegionId *string `json:"FcRegionId,omitempty" xml:"FcRegionId,omitempty"`
	// The type of the service in Function Compute.
	//
	// example:
	//
	// HttpTrigger
	FcType *string `json:"FcType,omitempty" xml:"FcType,omitempty"`
	// The function name that is defined in Function Compute.
	//
	// example:
	//
	// edge_function
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// Indicates whether the backend service receives only the service path.
	//
	// example:
	//
	// false
	OnlyBusinessPath *bool `json:"OnlyBusinessPath,omitempty" xml:"OnlyBusinessPath,omitempty"`
	// The alias of the function.
	//
	// example:
	//
	// testQualifier
	Qualifier *string `json:"Qualifier,omitempty" xml:"Qualifier,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is assumed by API Gateway to access Function Compute.
	//
	// example:
	//
	// acs:ram::31985*:role/aliyunserviceroleforbastionhostpam
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The service name that is defined in Function Compute.
	//
	// example:
	//
	// myservice
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The name of the trigger.
	//
	// You can specify the TriggerName or TriggerUrl parameter. The TriggerName parameter is optional.
	//
	// example:
	//
	// test1
	TriggerName *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) GetFcBaseUrl() *string {
	return s.FcBaseUrl
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) GetFcRegionId() *string {
	return s.FcRegionId
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) GetFcType() *string {
	return s.FcType
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) GetOnlyBusinessPath() *bool {
	return s.OnlyBusinessPath
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) GetQualifier() *string {
	return s.Qualifier
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) GetTriggerName() *string {
	return s.TriggerName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) SetFcBaseUrl(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig {
	s.FcBaseUrl = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) SetFcRegionId(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig {
	s.FcRegionId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) SetFcType(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig {
	s.FcType = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) SetFunctionName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) SetOnlyBusinessPath(v bool) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig {
	s.OnlyBusinessPath = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) SetQualifier(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig {
	s.Qualifier = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) SetRoleArn(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig {
	s.RoleArn = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) SetServiceName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig {
	s.ServiceName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) SetTriggerName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig {
	s.TriggerName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigFunctionComputeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig struct {
	// The header in the mocked response.
	MockHeaders []*DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders `json:"MockHeaders,omitempty" xml:"MockHeaders,omitempty" type:"Repeated"`
	// The mocked response.
	//
	// example:
	//
	// test
	MockResult *string `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	// The status code in the mocked response.
	//
	// example:
	//
	// 200
	MockStatusCode *string `json:"MockStatusCode,omitempty" xml:"MockStatusCode,omitempty"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig) GetMockHeaders() []*DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders {
	return s.MockHeaders
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig) GetMockResult() *string {
	return s.MockResult
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig) GetMockStatusCode() *string {
	return s.MockStatusCode
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig) SetMockHeaders(v []*DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig {
	s.MockHeaders = v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig) SetMockResult(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig {
	s.MockResult = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig) SetMockStatusCode(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig {
	s.MockStatusCode = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders struct {
	// The header name.
	//
	// example:
	//
	// test
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	// The header value.
	//
	// example:
	//
	// 123
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders) GetHeaderName() *string {
	return s.HeaderName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders) GetHeaderValue() *string {
	return s.HeaderValue
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders) SetHeaderName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders {
	s.HeaderName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders) SetHeaderValue(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders {
	s.HeaderValue = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigMockConfigMockHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// my_bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The region ID of the OSS bucket.
	//
	// example:
	//
	// cn-hangzhou
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig) GetOssRegionId() *string {
	return s.OssRegionId
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig) SetBucketName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig {
	s.BucketName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig) SetOssRegionId(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig {
	s.OssRegionId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigOssConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig struct {
	// The ID of the Elastic Compute Service (ECS) or Server Load Balancer (SLB) instance in the VPC.
	//
	// example:
	//
	// i-uf6iaale3gfef9t9cb41
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the VPC configuration.
	//
	// example:
	//
	// dypls-cn-beijing-slb-pre
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port number that corresponds to the instance.
	//
	// example:
	//
	// 8080
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the VPC access authorization.
	//
	// example:
	//
	// 2zej3ehuzg9m77kvwnfpn
	VpcAccessId *string `json:"VpcAccessId,omitempty" xml:"VpcAccessId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-wz9v96hqi6d14744sxqmx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates whether HTTP or HTTPS is used.
	//
	// example:
	//
	// http
	VpcScheme *string `json:"VpcScheme,omitempty" xml:"VpcScheme,omitempty"`
	// The host of the VPC backend service.
	//
	// example:
	//
	// openapi.alipan.com
	VpcTargetHostName *string `json:"VpcTargetHostName,omitempty" xml:"VpcTargetHostName,omitempty"`
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) GoString() string {
	return s.String()
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) GetName() *string {
	return s.Name
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) GetPort() *int64 {
	return s.Port
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) GetVpcAccessId() *string {
	return s.VpcAccessId
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) GetVpcScheme() *string {
	return s.VpcScheme
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) GetVpcTargetHostName() *string {
	return s.VpcTargetHostName
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) SetInstanceId(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) SetName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig {
	s.Name = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) SetPort(v int64) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig {
	s.Port = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) SetVpcAccessId(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig {
	s.VpcAccessId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) SetVpcId(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig {
	s.VpcId = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) SetVpcScheme(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig {
	s.VpcScheme = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) SetVpcTargetHostName(v string) *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig {
	s.VpcTargetHostName = &v
	return s
}

func (s *DescribeBackendInfoResponseBodyBackendInfoBackendModelsBackendConfigVpcConfig) Validate() error {
	return dara.Validate(s)
}
