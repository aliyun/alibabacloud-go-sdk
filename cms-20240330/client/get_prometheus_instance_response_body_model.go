// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusInstance(v *GetPrometheusInstanceResponseBodyPrometheusInstance) *GetPrometheusInstanceResponseBody
	GetPrometheusInstance() *GetPrometheusInstanceResponseBodyPrometheusInstance
	SetRequestId(v string) *GetPrometheusInstanceResponseBody
	GetRequestId() *string
}

type GetPrometheusInstanceResponseBody struct {
	// Details of the Prometheus instance.
	PrometheusInstance *GetPrometheusInstanceResponseBodyPrometheusInstance `json:"prometheusInstance,omitempty" xml:"prometheusInstance,omitempty" type:"Struct"`
	// Unique identifier for the request.
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPrometheusInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrometheusInstanceResponseBody) GetPrometheusInstance() *GetPrometheusInstanceResponseBodyPrometheusInstance {
	return s.PrometheusInstance
}

func (s *GetPrometheusInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPrometheusInstanceResponseBody) SetPrometheusInstance(v *GetPrometheusInstanceResponseBodyPrometheusInstance) *GetPrometheusInstanceResponseBody {
	s.PrometheusInstance = v
	return s
}

func (s *GetPrometheusInstanceResponseBody) SetRequestId(v string) *GetPrometheusInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBody) Validate() error {
	if s.PrometheusInstance != nil {
		if err := s.PrometheusInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPrometheusInstanceResponseBodyPrometheusInstance struct {
	// Access type:
	//
	// readWrite, readOnly, httpReadOnly
	//
	// example:
	//
	// readOnly
	AccessType *string `json:"accessType,omitempty" xml:"accessType,omitempty"`
	// Number of days to automatically archive and save after storage expiration. 0 means no archiving, 3650 means permanent saving.
	//
	// example:
	//
	// 90
	ArchiveDuration *int32 `json:"archiveDuration,omitempty" xml:"archiveDuration,omitempty"`
	// Password-free read policy (supports IP segments and VpcId).
	//
	// example:
	//
	// {
	//
	//   "SourceIp": [
	//
	//     "192.168.1.0/24",
	//
	//     "172.168.2.22"
	//
	//   ],
	//
	//   "SourceVpc": [
	//
	//     "vpc-xx1",
	//
	//     "vpc-xx2"
	//
	//   ]
	//
	// }
	AuthFreeReadPolicy *string `json:"authFreeReadPolicy,omitempty" xml:"authFreeReadPolicy,omitempty"`
	// Password-free write policy (supports IP segments and VpcId).
	//
	// example:
	//
	// {
	//
	//   "SourceIp": [
	//
	//     "192.168.1.0/24",
	//
	//     "172.168.2.22"
	//
	//   ],
	//
	//   "SourceVpc": [
	//
	//     "vpc-xx1",
	//
	//     "vpc-xx2"
	//
	//   ]
	//
	// }
	AuthFreeWritePolicy *string `json:"authFreeWritePolicy,omitempty" xml:"authFreeWritePolicy,omitempty"`
	// authToken string.
	//
	// example:
	//
	// eJwixxxxx
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// Instance creation time, using UTC+0, formatted as yyyy-MM-ddTHH:mmZ.
	//
	// example:
	//
	// 2025-08-10T02:07:53Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Whether to enable password-free reading.
	//
	// example:
	//
	// true
	EnableAuthFreeRead *bool `json:"enableAuthFreeRead,omitempty" xml:"enableAuthFreeRead,omitempty"`
	// Whether to enable password-free writing.
	//
	// example:
	//
	// true
	EnableAuthFreeWrite *bool `json:"enableAuthFreeWrite,omitempty" xml:"enableAuthFreeWrite,omitempty"`
	// Whether to enable authentication token.
	//
	// example:
	//
	// true
	EnableAuthToken *bool `json:"enableAuthToken,omitempty" xml:"enableAuthToken,omitempty"`
	// Additional information.
	ExtraInfo map[string]*string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// URL of the visualization dashboard directory.
	//
	// example:
	//
	// https://gnew.console.aliyun.com/dashboards/f/c49a80d2a551d4a20a8c4b996b0be4e52/xxxxxxx
	FolderUrl *string `json:"folderUrl,omitempty" xml:"folderUrl,omitempty"`
	// ID of the managed Grafana instance that is bound.
	//
	// example:
	//
	// SHARED
	GrafanaInstanceId *string `json:"grafanaInstanceId,omitempty" xml:"grafanaInstanceId,omitempty"`
	// Name of the managed Grafana instance that is bound.
	//
	// example:
	//
	// 共享版
	GrafanaInstanceName *string `json:"grafanaInstanceName,omitempty" xml:"grafanaInstanceName,omitempty"`
	// HTTP public network address.
	//
	// example:
	//
	// http://workspace-default-cms-xxxxxxx
	HttpApiInterUrl *string `json:"httpApiInterUrl,omitempty" xml:"httpApiInterUrl,omitempty"`
	// HTTP intranet address.
	//
	// example:
	//
	// http://workspace-default-cms-xxxxxxx
	HttpApiIntraUrl *string `json:"httpApiIntraUrl,omitempty" xml:"httpApiIntraUrl,omitempty"`
	// Prometheus instance type.
	//
	// example:
	//
	// remote-write
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// Billing method:
	//
	// POSTPAY: Pay-as-you-go based on metric reporting volume.
	//
	// POSTPAY_GB: Pay-as-you-go based on metric write volume.
	//
	// example:
	//
	// POSTPAY
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// Time when the billing method of the instance was updated.
	//
	// example:
	//
	// 2025-08-10T02:07:53Z
	PaymentTypeUpdateTime *string `json:"paymentTypeUpdateTime,omitempty" xml:"paymentTypeUpdateTime,omitempty"`
	// The product to which the Prometheus instance belongs (arms or cms).
	//
	// example:
	//
	// cms
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// rw-524ada714221af267c73122af2e1
	PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	// Instance name.
	//
	// example:
	//
	// test-prom-name
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	// Public network address of PushGateway.
	//
	// example:
	//
	// http://workspace-default-cms-xxxxxxx
	PushGatewayInterUrl *string `json:"pushGatewayInterUrl,omitempty" xml:"pushGatewayInterUrl,omitempty"`
	// Intranet address of PushGateway.
	//
	// example:
	//
	// http://workspace-default-cms-xxxxxxx
	PushGatewayIntraUrl *string `json:"pushGatewayIntraUrl,omitempty" xml:"pushGatewayIntraUrl,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Public network read address.
	//
	// example:
	//
	// http://workspace-default-cms-xxxxxxx
	RemoteReadInterUrl *string `json:"remoteReadInterUrl,omitempty" xml:"remoteReadInterUrl,omitempty"`
	// Intranet read address.
	//
	// example:
	//
	// https://workspace-default-cms-1xxxxxxxxxx
	RemoteReadIntraUrl *string `json:"remoteReadIntraUrl,omitempty" xml:"remoteReadIntraUrl,omitempty"`
	// Public network write address.
	//
	// example:
	//
	// https://workspace-default-cms-xxxxxxxxxx
	RemoteWriteInterUrl *string `json:"remoteWriteInterUrl,omitempty" xml:"remoteWriteInterUrl,omitempty"`
	// Intranet write address.
	//
	// example:
	//
	// https://workspace-default-cms-xxxxxxxxxx
	RemoteWriteIntraUrl *string `json:"remoteWriteIntraUrl,omitempty" xml:"remoteWriteIntraUrl,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfm3gn5i6bigbi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Fixed value: PrometheusInstance.
	//
	// example:
	//
	// Prometheus
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// Instance status.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Storage duration (in days).
	//
	// example:
	//
	// 90
	StorageDuration *int32 `json:"storageDuration,omitempty" xml:"storageDuration,omitempty"`
	// Supported authentication types.
	SupportAuthTypes []*string `json:"supportAuthTypes,omitempty" xml:"supportAuthTypes,omitempty" type:"Repeated"`
	// List of tags.
	Tags []*GetPrometheusInstanceResponseBodyPrometheusInstanceTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// User ID.
	//
	// example:
	//
	// 170731234567
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// Version.
	//
	// example:
	//
	// V1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The workspace to which the Prometheus instance belongs.
	//
	// example:
	//
	// ws1
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetPrometheusInstanceResponseBodyPrometheusInstance) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusInstanceResponseBodyPrometheusInstance) GoString() string {
	return s.String()
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetAccessType() *string {
	return s.AccessType
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetArchiveDuration() *int32 {
	return s.ArchiveDuration
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetAuthFreeReadPolicy() *string {
	return s.AuthFreeReadPolicy
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetAuthFreeWritePolicy() *string {
	return s.AuthFreeWritePolicy
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetAuthToken() *string {
	return s.AuthToken
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetEnableAuthFreeRead() *bool {
	return s.EnableAuthFreeRead
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetEnableAuthFreeWrite() *bool {
	return s.EnableAuthFreeWrite
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetEnableAuthToken() *bool {
	return s.EnableAuthToken
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetExtraInfo() map[string]*string {
	return s.ExtraInfo
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetFolderUrl() *string {
	return s.FolderUrl
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetGrafanaInstanceId() *string {
	return s.GrafanaInstanceId
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetGrafanaInstanceName() *string {
	return s.GrafanaInstanceName
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetHttpApiInterUrl() *string {
	return s.HttpApiInterUrl
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetHttpApiIntraUrl() *string {
	return s.HttpApiIntraUrl
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetPaymentTypeUpdateTime() *string {
	return s.PaymentTypeUpdateTime
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetProduct() *string {
	return s.Product
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetPrometheusInstanceName() *string {
	return s.PrometheusInstanceName
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetPushGatewayInterUrl() *string {
	return s.PushGatewayInterUrl
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetPushGatewayIntraUrl() *string {
	return s.PushGatewayIntraUrl
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetRemoteReadInterUrl() *string {
	return s.RemoteReadInterUrl
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetRemoteReadIntraUrl() *string {
	return s.RemoteReadIntraUrl
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetRemoteWriteInterUrl() *string {
	return s.RemoteWriteInterUrl
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetRemoteWriteIntraUrl() *string {
	return s.RemoteWriteIntraUrl
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetStatus() *string {
	return s.Status
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetStorageDuration() *int32 {
	return s.StorageDuration
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetSupportAuthTypes() []*string {
	return s.SupportAuthTypes
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetTags() []*GetPrometheusInstanceResponseBodyPrometheusInstanceTags {
	return s.Tags
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetUserId() *string {
	return s.UserId
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetVersion() *string {
	return s.Version
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetAccessType(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.AccessType = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetArchiveDuration(v int32) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.ArchiveDuration = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetAuthFreeReadPolicy(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.AuthFreeReadPolicy = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetAuthFreeWritePolicy(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.AuthFreeWritePolicy = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetAuthToken(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.AuthToken = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetCreateTime(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.CreateTime = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetEnableAuthFreeRead(v bool) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.EnableAuthFreeRead = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetEnableAuthFreeWrite(v bool) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.EnableAuthFreeWrite = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetEnableAuthToken(v bool) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.EnableAuthToken = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetExtraInfo(v map[string]*string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.ExtraInfo = v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetFolderUrl(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.FolderUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetGrafanaInstanceId(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.GrafanaInstanceId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetGrafanaInstanceName(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.GrafanaInstanceName = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetHttpApiInterUrl(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.HttpApiInterUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetHttpApiIntraUrl(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.HttpApiIntraUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetInstanceType(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.InstanceType = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetPaymentType(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.PaymentType = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetPaymentTypeUpdateTime(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.PaymentTypeUpdateTime = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetProduct(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.Product = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetPrometheusInstanceId(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.PrometheusInstanceId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetPrometheusInstanceName(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.PrometheusInstanceName = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetPushGatewayInterUrl(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.PushGatewayInterUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetPushGatewayIntraUrl(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.PushGatewayIntraUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetRegionId(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.RegionId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetRemoteReadInterUrl(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.RemoteReadInterUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetRemoteReadIntraUrl(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.RemoteReadIntraUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetRemoteWriteInterUrl(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.RemoteWriteInterUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetRemoteWriteIntraUrl(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.RemoteWriteIntraUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetResourceGroupId(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetResourceType(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.ResourceType = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetStatus(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.Status = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetStorageDuration(v int32) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.StorageDuration = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetSupportAuthTypes(v []*string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.SupportAuthTypes = v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetTags(v []*GetPrometheusInstanceResponseBodyPrometheusInstanceTags) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.Tags = v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetUserId(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.UserId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetVersion(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.Version = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) SetWorkspace(v string) *GetPrometheusInstanceResponseBodyPrometheusInstance {
	s.Workspace = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstance) Validate() error {
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

type GetPrometheusInstanceResponseBodyPrometheusInstanceTags struct {
	// Tag key.
	//
	// example:
	//
	// openStorage
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Matched value.
	//
	// example:
	//
	// 130303196111114281
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetPrometheusInstanceResponseBodyPrometheusInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusInstanceResponseBodyPrometheusInstanceTags) GoString() string {
	return s.String()
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstanceTags) GetKey() *string {
	return s.Key
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstanceTags) GetValue() *string {
	return s.Value
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstanceTags) SetKey(v string) *GetPrometheusInstanceResponseBodyPrometheusInstanceTags {
	s.Key = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstanceTags) SetValue(v string) *GetPrometheusInstanceResponseBodyPrometheusInstanceTags {
	s.Value = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyPrometheusInstanceTags) Validate() error {
	return dara.Validate(s)
}
