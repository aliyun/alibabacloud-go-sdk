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
	PrometheusInstance *GetPrometheusInstanceResponseBodyPrometheusInstance `json:"prometheusInstance,omitempty" xml:"prometheusInstance,omitempty" type:"Struct"`
	// Id of the request
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
	return dara.Validate(s)
}

type GetPrometheusInstanceResponseBodyPrometheusInstance struct {
	// example:
	//
	// readOnly
	AccessType *string `json:"accessType,omitempty" xml:"accessType,omitempty"`
	// example:
	//
	// 165
	ArchiveDuration *int32 `json:"archiveDuration,omitempty" xml:"archiveDuration,omitempty"`
	// example:
	//
	// 0.0.0.0/0
	AuthFreeReadPolicy *string `json:"authFreeReadPolicy,omitempty" xml:"authFreeReadPolicy,omitempty"`
	// example:
	//
	// 0.0.0.0/0
	AuthFreeWritePolicy *string `json:"authFreeWritePolicy,omitempty" xml:"authFreeWritePolicy,omitempty"`
	// example:
	//
	// eJwixxxxx
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// example:
	//
	// 2025-08-10T02:07:53Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// true
	EnableAuthFreeRead *bool `json:"enableAuthFreeRead,omitempty" xml:"enableAuthFreeRead,omitempty"`
	// example:
	//
	// true
	EnableAuthFreeWrite *bool `json:"enableAuthFreeWrite,omitempty" xml:"enableAuthFreeWrite,omitempty"`
	// example:
	//
	// true
	EnableAuthToken *bool              `json:"enableAuthToken,omitempty" xml:"enableAuthToken,omitempty"`
	ExtraInfo       map[string]*string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// https://gnew.console.aliyun.com/dashboards/f/c49a80d2a551d4a20a8c4b996b0be4e52/xxxxxxx
	FolderUrl *string `json:"folderUrl,omitempty" xml:"folderUrl,omitempty"`
	// example:
	//
	// SHARED
	GrafanaInstanceId   *string `json:"grafanaInstanceId,omitempty" xml:"grafanaInstanceId,omitempty"`
	GrafanaInstanceName *string `json:"grafanaInstanceName,omitempty" xml:"grafanaInstanceName,omitempty"`
	// example:
	//
	// http://workspace-default-cms-xxxxxxx
	HttpApiInterUrl *string `json:"httpApiInterUrl,omitempty" xml:"httpApiInterUrl,omitempty"`
	// example:
	//
	// http://workspace-default-cms-xxxxxxx
	HttpApiIntraUrl *string `json:"httpApiIntraUrl,omitempty" xml:"httpApiIntraUrl,omitempty"`
	// remote-write（Prometheus for Remote Write）
	//
	// example:
	//
	// e1
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// example:
	//
	// prepaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// example:
	//
	// null
	PaymentTypeUpdateTime *string `json:"paymentTypeUpdateTime,omitempty" xml:"paymentTypeUpdateTime,omitempty"`
	// example:
	//
	// NAS
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// example:
	//
	// rw-524ada714221af267c73122af2e1
	PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	// example:
	//
	// test-prom-name
	PrometheusInstanceName *string `json:"prometheusInstanceName,omitempty" xml:"prometheusInstanceName,omitempty"`
	// example:
	//
	// http://workspace-default-cms-xxxxxxx
	PushGatewayInterUrl *string `json:"pushGatewayInterUrl,omitempty" xml:"pushGatewayInterUrl,omitempty"`
	// example:
	//
	// http://workspace-default-cms-xxxxxxx
	PushGatewayIntraUrl *string `json:"pushGatewayIntraUrl,omitempty" xml:"pushGatewayIntraUrl,omitempty"`
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// http://workspace-default-cms-xxxxxxx
	RemoteReadInterUrl *string `json:"remoteReadInterUrl,omitempty" xml:"remoteReadInterUrl,omitempty"`
	// example:
	//
	// https://workspace-default-cms-1xxxxxxxxxx
	RemoteReadIntraUrl *string `json:"remoteReadIntraUrl,omitempty" xml:"remoteReadIntraUrl,omitempty"`
	// example:
	//
	// https://workspace-default-cms-xxxxxxxxxx
	RemoteWriteInterUrl *string `json:"remoteWriteInterUrl,omitempty" xml:"remoteWriteInterUrl,omitempty"`
	// example:
	//
	// https://workspace-default-cms-xxxxxxxxxx
	RemoteWriteIntraUrl *string `json:"remoteWriteIntraUrl,omitempty" xml:"remoteWriteIntraUrl,omitempty"`
	// example:
	//
	// rg-acfm3gn5i6bigbi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// Prometheus
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 90
	StorageDuration  *int32                                                     `json:"storageDuration,omitempty" xml:"storageDuration,omitempty"`
	SupportAuthTypes []*string                                                  `json:"supportAuthTypes,omitempty" xml:"supportAuthTypes,omitempty" type:"Repeated"`
	Tags             []*GetPrometheusInstanceResponseBodyPrometheusInstanceTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 170731234567
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// *
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// rum-monitor-test-aysls-pub-cn-qingdao-monitor
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
	return dara.Validate(s)
}

type GetPrometheusInstanceResponseBodyPrometheusInstanceTags struct {
	// example:
	//
	// openStorage
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
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
