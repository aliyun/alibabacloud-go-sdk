// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnsureSharedPrometheusInstanceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCreated(v bool) *EnsureSharedPrometheusInstanceResponseBody
  GetCreated() *bool 
  SetPrometheusInstance(v *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) *EnsureSharedPrometheusInstanceResponseBody
  GetPrometheusInstance() *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance 
  SetRequestId(v string) *EnsureSharedPrometheusInstanceResponseBody
  GetRequestId() *string 
}

type EnsureSharedPrometheusInstanceResponseBody struct {
  // Indicates whether a new instance is created by this call. The value is false if an existing instance is found.
  // 
  // example:
  // 
  // false
  Created *bool `json:"created,omitempty" xml:"created,omitempty"`
  // The details of the shared Prometheus instance. The fields are the same as the prometheusInstance field in GetPrometheusInstance.
  PrometheusInstance *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance `json:"prometheusInstance,omitempty" xml:"prometheusInstance,omitempty" type:"Struct"`
  // The request ID.
  // 
  // example:
  // 
  // A1B2C3D4-E5F6-7890-ABCD-EF1234567890
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EnsureSharedPrometheusInstanceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnsureSharedPrometheusInstanceResponseBody) GoString() string {
  return s.String()
}

func (s *EnsureSharedPrometheusInstanceResponseBody) GetCreated() *bool  {
  return s.Created
}

func (s *EnsureSharedPrometheusInstanceResponseBody) GetPrometheusInstance() *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance  {
  return s.PrometheusInstance
}

func (s *EnsureSharedPrometheusInstanceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnsureSharedPrometheusInstanceResponseBody) SetCreated(v bool) *EnsureSharedPrometheusInstanceResponseBody {
  s.Created = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBody) SetPrometheusInstance(v *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) *EnsureSharedPrometheusInstanceResponseBody {
  s.PrometheusInstance = v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBody) SetRequestId(v string) *EnsureSharedPrometheusInstanceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBody) Validate() error {
  if s.PrometheusInstance != nil {
    if err := s.PrometheusInstance.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance struct {
  // The access permissions of the instance. Valid values: readWrite, readOnly, and httpReadOnly.
  // 
  // example:
  // 
  // readWrite
  AccessType *string `json:"accessType,omitempty" xml:"accessType,omitempty"`
  // The number of days for automatic archive retention after the storage expires. A value of 0 indicates no archiving, and 3650 indicates permanent retention.
  // 
  // example:
  // 
  // 0
  ArchiveDuration *int32 `json:"archiveDuration,omitempty" xml:"archiveDuration,omitempty"`
  // The IP address whitelist policy that allows access to the authentication-free read endpoint.
  // 
  // example:
  // 
  // 127.0.0.1
  AuthFreeReadPolicy *string `json:"authFreeReadPolicy,omitempty" xml:"authFreeReadPolicy,omitempty"`
  // The IP address whitelist policy that allows access to the authentication-free write endpoint.
  // 
  // example:
  // 
  // 127.0.0.1
  AuthFreeWritePolicy *string `json:"authFreeWritePolicy,omitempty" xml:"authFreeWritePolicy,omitempty"`
  // The authentication token of the instance.
  // 
  // example:
  // 
  // example-auth-token
  AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
  // The time when the instance was created. The time is displayed in UTC+0 in the yyyy-MM-ddTHH:mmZ format.
  // 
  // example:
  // 
  // 2026-11-11T06:00Z
  CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
  // Indicates whether authentication-free read is enabled.
  // 
  // example:
  // 
  // false
  EnableAuthFreeRead *bool `json:"enableAuthFreeRead,omitempty" xml:"enableAuthFreeRead,omitempty"`
  // Indicates whether authentication-free write is enabled.
  // 
  // example:
  // 
  // false
  EnableAuthFreeWrite *bool `json:"enableAuthFreeWrite,omitempty" xml:"enableAuthFreeWrite,omitempty"`
  // Indicates whether authentication token-based authentication is enabled.
  // 
  // example:
  // 
  // false
  EnableAuthToken *bool `json:"enableAuthToken,omitempty" xml:"enableAuthToken,omitempty"`
  // The extended information. This field is returned only for console requests.
  // 
  // example:
  // 
  // {"slsProject":"proj-cms-cn-hangzhou"}
  ExtraInfo map[string]*string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
  // The URL of the dashboard associated with the instance.
  // 
  // example:
  // 
  // https://g.console.aliyun.com/folders/example
  FolderUrl *string `json:"folderUrl,omitempty" xml:"folderUrl,omitempty"`
  // The ID of the bound managed Grafana instance.
  // 
  // example:
  // 
  // grafana-cn-hangzhou-example
  GrafanaInstanceId *string `json:"grafanaInstanceId,omitempty" xml:"grafanaInstanceId,omitempty"`
  // The name of the bound managed Grafana instance.
  // 
  // example:
  // 
  // shared-grafana
  GrafanaInstanceName *string `json:"grafanaInstanceName,omitempty" xml:"grafanaInstanceName,omitempty"`
  // The public endpoint of the Prometheus HTTP API.
  // 
  // example:
  // 
  // https://cn-hangzhou.arms.aliyuncs.com/prometheus/rw-example/api/v1
  HttpApiInterUrl *string `json:"httpApiInterUrl,omitempty" xml:"httpApiInterUrl,omitempty"`
  // The internal endpoint of the Prometheus HTTP API.
  // 
  // example:
  // 
  // http://prometheus.cn-hangzhou.aliyuncs.com/prometheus/rw-example/api/v1
  HttpApiInternalUrl *string `json:"httpApiInternalUrl,omitempty" xml:"httpApiInternalUrl,omitempty"`
  // The internal network endpoint of the Prometheus HTTP API.
  // 
  // example:
  // 
  // https://cn-hangzhou-intranet.arms.aliyuncs.com/prometheus/rw-example/api/v1
  HttpApiIntraUrl *string `json:"httpApiIntraUrl,omitempty" xml:"httpApiIntraUrl,omitempty"`
  // The instance type. For a shared instance, the value is remote-write, which indicates Prometheus for Remote Write.
  // 
  // example:
  // 
  // remote-write
  InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
  // The billing method. POSTPAY indicates pay-as-you-go billing based on the number of reported metrics, and POSTPAY_GB indicates pay-as-you-go billing based on the volume of written metrics.
  // 
  // example:
  // 
  // POSTPAY_GB
  PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
  // The time when the billing method of the instance was last modified. The time is in UTC+0 and formatted as yyyy-MM-ddTHH:mmZ.
  // 
  // example:
  // 
  // 2026-11-11T06:00Z
  PaymentTypeUpdateTime *string `json:"paymentTypeUpdateTime,omitempty" xml:"paymentTypeUpdateTime,omitempty"`
  // The product to which the instance belongs. Valid values: arms and cms.
  // 
  // example:
  // 
  // cms
  Product *string `json:"product,omitempty" xml:"product,omitempty"`
  // The ID of the Prometheus instance.
  // 
  // example:
  // 
  // rw-8f3c1a2b4d5e6f70
  PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
  // The name of the Prometheus instance. For a workspace shared instance, the name is generated by the server in the RegionShare:{workspace}:{regionId} format.
  // 
  // example:
  // 
  // RegionShare:team-workspace:cn-hangzhou
  PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
  // The public endpoint of Pushgateway.
  // 
  // example:
  // 
  // https://cn-hangzhou.arms.aliyuncs.com/prometheus/rw-example/api/v1/pushgateway
  PushGatewayInterUrl *string `json:"pushGatewayInterUrl,omitempty" xml:"pushGatewayInterUrl,omitempty"`
  // The internal endpoint of Pushgateway.
  // 
  // example:
  // 
  // http://prometheus.cn-hangzhou.aliyuncs.com/prometheus/rw-example/api/v1/pushgateway
  PushGatewayInternalUrl *string `json:"pushGatewayInternalUrl,omitempty" xml:"pushGatewayInternalUrl,omitempty"`
  // The internal network endpoint of Pushgateway.
  // 
  // example:
  // 
  // https://cn-hangzhou-intranet.arms.aliyuncs.com/prometheus/rw-example/api/v1/pushgateway
  PushGatewayIntraUrl *string `json:"pushGatewayIntraUrl,omitempty" xml:"pushGatewayIntraUrl,omitempty"`
  // The ID of the region where the instance resides.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
  // The public endpoint of Remote Read.
  // 
  // example:
  // 
  // https://cn-hangzhou.arms.aliyuncs.com/prometheus/rw-example/api/v1/read
  RemoteReadInterUrl *string `json:"remoteReadInterUrl,omitempty" xml:"remoteReadInterUrl,omitempty"`
  // The internal endpoint of Remote Read.
  // 
  // example:
  // 
  // http://prometheus.cn-hangzhou.aliyuncs.com/prometheus/rw-example/api/v1/read
  RemoteReadInternalUrl *string `json:"remoteReadInternalUrl,omitempty" xml:"remoteReadInternalUrl,omitempty"`
  // The internal network endpoint of Remote Read.
  // 
  // example:
  // 
  // https://cn-hangzhou-intranet.arms.aliyuncs.com/prometheus/rw-example/api/v1/read
  RemoteReadIntraUrl *string `json:"remoteReadIntraUrl,omitempty" xml:"remoteReadIntraUrl,omitempty"`
  // The public endpoint of Remote Write.
  // 
  // example:
  // 
  // https://cn-hangzhou.arms.aliyuncs.com/prometheus/rw-example/api/v1/write
  RemoteWriteInterUrl *string `json:"remoteWriteInterUrl,omitempty" xml:"remoteWriteInterUrl,omitempty"`
  // The internal endpoint of Remote Write.
  // 
  // example:
  // 
  // http://prometheus.cn-hangzhou.aliyuncs.com/prometheus/rw-example/api/v1/write
  RemoteWriteInternalUrl *string `json:"remoteWriteInternalUrl,omitempty" xml:"remoteWriteInternalUrl,omitempty"`
  // The internal network endpoint of Remote Write.
  // 
  // example:
  // 
  // https://cn-hangzhou-intranet.arms.aliyuncs.com/prometheus/rw-example/api/v1/write
  RemoteWriteIntraUrl *string `json:"remoteWriteIntraUrl,omitempty" xml:"remoteWriteIntraUrl,omitempty"`
  // The ID of the resource group to which the instance belongs.
  // 
  // example:
  // 
  // rg-acfmyexample
  ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
  // The resource type. The value is fixed as PrometheusInstance.
  // 
  // example:
  // 
  // PrometheusInstance
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
  // The backend data storage status, such as Running.
  // 
  // example:
  // 
  // Running
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // The metric storage duration in days.
  // 
  // example:
  // 
  // 90
  StorageDuration *int32 `json:"storageDuration,omitempty" xml:"storageDuration,omitempty"`
  // The Prometheus query cache and parallel query configurations.
  StoreConfig *PrometheusInstanceStoreConfig `json:"storeConfig,omitempty" xml:"storeConfig,omitempty"`
  // The list of authentication and authorization types supported by the instance.
  // 
  // example:
  // 
  // BasicAuth
  SupportAuthTypes []*string `json:"supportAuthTypes,omitempty" xml:"supportAuthTypes,omitempty" type:"Repeated"`
  // The instance tags. A workspace shared instance includes the system tag acs:cms:shareScope.
  Tags []*EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
  // The ID of the Alibaba Cloud account to which the instance belongs.
  // 
  // example:
  // 
  // 1234567890123456
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
  // The version of the instance, such as V2.
  // 
  // example:
  // 
  // V2
  Version *string `json:"version,omitempty" xml:"version,omitempty"`
  // The name of the workspace to which the instance belongs.
  // 
  // example:
  // 
  // team-workspace
  Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) String() string {
  return dara.Prettify(s)
}

func (s EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GoString() string {
  return s.String()
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetAccessType() *string  {
  return s.AccessType
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetArchiveDuration() *int32  {
  return s.ArchiveDuration
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetAuthFreeReadPolicy() *string  {
  return s.AuthFreeReadPolicy
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetAuthFreeWritePolicy() *string  {
  return s.AuthFreeWritePolicy
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetAuthToken() *string  {
  return s.AuthToken
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetCreateTime() *string  {
  return s.CreateTime
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetEnableAuthFreeRead() *bool  {
  return s.EnableAuthFreeRead
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetEnableAuthFreeWrite() *bool  {
  return s.EnableAuthFreeWrite
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetEnableAuthToken() *bool  {
  return s.EnableAuthToken
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetExtraInfo() map[string]*string  {
  return s.ExtraInfo
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetFolderUrl() *string  {
  return s.FolderUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetGrafanaInstanceId() *string  {
  return s.GrafanaInstanceId
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetGrafanaInstanceName() *string  {
  return s.GrafanaInstanceName
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetHttpApiInterUrl() *string  {
  return s.HttpApiInterUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetHttpApiInternalUrl() *string  {
  return s.HttpApiInternalUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetHttpApiIntraUrl() *string  {
  return s.HttpApiIntraUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetInstanceType() *string  {
  return s.InstanceType
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetPaymentType() *string  {
  return s.PaymentType
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetPaymentTypeUpdateTime() *string  {
  return s.PaymentTypeUpdateTime
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetProduct() *string  {
  return s.Product
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetPrometheusInstanceId() *string  {
  return s.PrometheusInstanceId
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetPrometheusInstanceName() *string  {
  return s.PrometheusInstanceName
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetPushGatewayInterUrl() *string  {
  return s.PushGatewayInterUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetPushGatewayInternalUrl() *string  {
  return s.PushGatewayInternalUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetPushGatewayIntraUrl() *string  {
  return s.PushGatewayIntraUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetRemoteReadInterUrl() *string  {
  return s.RemoteReadInterUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetRemoteReadInternalUrl() *string  {
  return s.RemoteReadInternalUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetRemoteReadIntraUrl() *string  {
  return s.RemoteReadIntraUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetRemoteWriteInterUrl() *string  {
  return s.RemoteWriteInterUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetRemoteWriteInternalUrl() *string  {
  return s.RemoteWriteInternalUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetRemoteWriteIntraUrl() *string  {
  return s.RemoteWriteIntraUrl
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetResourceType() *string  {
  return s.ResourceType
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetStatus() *string  {
  return s.Status
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetStorageDuration() *int32  {
  return s.StorageDuration
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetStoreConfig() *PrometheusInstanceStoreConfig  {
  return s.StoreConfig
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetSupportAuthTypes() []*string  {
  return s.SupportAuthTypes
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetTags() []*EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags  {
  return s.Tags
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetUserId() *string  {
  return s.UserId
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetVersion() *string  {
  return s.Version
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) GetWorkspace() *string  {
  return s.Workspace
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetAccessType(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.AccessType = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetArchiveDuration(v int32) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.ArchiveDuration = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetAuthFreeReadPolicy(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.AuthFreeReadPolicy = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetAuthFreeWritePolicy(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.AuthFreeWritePolicy = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetAuthToken(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.AuthToken = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetCreateTime(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.CreateTime = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetEnableAuthFreeRead(v bool) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.EnableAuthFreeRead = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetEnableAuthFreeWrite(v bool) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.EnableAuthFreeWrite = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetEnableAuthToken(v bool) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.EnableAuthToken = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetExtraInfo(v map[string]*string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.ExtraInfo = v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetFolderUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.FolderUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetGrafanaInstanceId(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.GrafanaInstanceId = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetGrafanaInstanceName(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.GrafanaInstanceName = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetHttpApiInterUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.HttpApiInterUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetHttpApiInternalUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.HttpApiInternalUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetHttpApiIntraUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.HttpApiIntraUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetInstanceType(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.InstanceType = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetPaymentType(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.PaymentType = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetPaymentTypeUpdateTime(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.PaymentTypeUpdateTime = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetProduct(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.Product = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetPrometheusInstanceId(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.PrometheusInstanceId = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetPrometheusInstanceName(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.PrometheusInstanceName = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetPushGatewayInterUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.PushGatewayInterUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetPushGatewayInternalUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.PushGatewayInternalUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetPushGatewayIntraUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.PushGatewayIntraUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetRegionId(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.RegionId = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetRemoteReadInterUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.RemoteReadInterUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetRemoteReadInternalUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.RemoteReadInternalUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetRemoteReadIntraUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.RemoteReadIntraUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetRemoteWriteInterUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.RemoteWriteInterUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetRemoteWriteInternalUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.RemoteWriteInternalUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetRemoteWriteIntraUrl(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.RemoteWriteIntraUrl = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetResourceGroupId(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.ResourceGroupId = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetResourceType(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.ResourceType = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetStatus(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.Status = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetStorageDuration(v int32) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.StorageDuration = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetStoreConfig(v *PrometheusInstanceStoreConfig) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.StoreConfig = v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetSupportAuthTypes(v []*string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.SupportAuthTypes = v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetTags(v []*EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.Tags = v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetUserId(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.UserId = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetVersion(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.Version = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) SetWorkspace(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance {
  s.Workspace = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstance) Validate() error {
  if s.StoreConfig != nil {
    if err := s.StoreConfig.Validate(); err != nil {
      return err
    }
  }
  if s.Tags != nil {
    for _, item := range s.Tags {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags struct {
  // The key of the instance tag.
  // 
  // example:
  // 
  // acs:cms:shareScope
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // The value of the instance tag.
  // 
  // example:
  // 
  // region:cn-hangzhou:team-workspace
  Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags) String() string {
  return dara.Prettify(s)
}

func (s EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags) GoString() string {
  return s.String()
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags) GetKey() *string  {
  return s.Key
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags) GetValue() *string  {
  return s.Value
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags) SetKey(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags {
  s.Key = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags) SetValue(v string) *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags {
  s.Value = &v
  return s
}

func (s *EnsureSharedPrometheusInstanceResponseBodyPrometheusInstanceTags) Validate() error {
  return dara.Validate(s)
}

