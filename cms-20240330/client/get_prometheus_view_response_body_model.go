// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusView(v *GetPrometheusViewResponseBodyPrometheusView) *GetPrometheusViewResponseBody
	GetPrometheusView() *GetPrometheusViewResponseBodyPrometheusView
	SetRequestId(v string) *GetPrometheusViewResponseBody
	GetRequestId() *string
}

type GetPrometheusViewResponseBody struct {
	// View instance.
	PrometheusView *GetPrometheusViewResponseBodyPrometheusView `json:"prometheusView,omitempty" xml:"prometheusView,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPrometheusViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusViewResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrometheusViewResponseBody) GetPrometheusView() *GetPrometheusViewResponseBodyPrometheusView {
	return s.PrometheusView
}

func (s *GetPrometheusViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPrometheusViewResponseBody) SetPrometheusView(v *GetPrometheusViewResponseBodyPrometheusView) *GetPrometheusViewResponseBody {
	s.PrometheusView = v
	return s
}

func (s *GetPrometheusViewResponseBody) SetRequestId(v string) *GetPrometheusViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrometheusViewResponseBody) Validate() error {
	if s.PrometheusView != nil {
		if err := s.PrometheusView.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPrometheusViewResponseBodyPrometheusView struct {
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
	// authToken string.
	//
	// example:
	//
	// eJxxxxxx
	AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
	// Instance creation time, using UTC+0 time, format is yyyy-MM-ddTHH:mmZ.
	//
	// example:
	//
	// 2025-08-10T02:07:53Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Whether to enable password-free read.
	//
	// example:
	//
	// true
	EnableAuthFreeRead *bool `json:"enableAuthFreeRead,omitempty" xml:"enableAuthFreeRead,omitempty"`
	// Whether to enable authToken.
	//
	// example:
	//
	// true
	EnableAuthToken *bool `json:"enableAuthToken,omitempty" xml:"enableAuthToken,omitempty"`
	// Observability dashboard URL.
	//
	// example:
	//
	// https://xxxx
	FolderUrl *string `json:"folderUrl,omitempty" xml:"folderUrl,omitempty"`
	// Bound managed Grafana instance ID.
	//
	// example:
	//
	// g-xxx
	GrafanaInstanceId *string `json:"grafanaInstanceId,omitempty" xml:"grafanaInstanceId,omitempty"`
	// Bound managed Grafana instance name.
	//
	// example:
	//
	// gxxx
	GrafanaInstanceName *string `json:"grafanaInstanceName,omitempty" xml:"grafanaInstanceName,omitempty"`
	// Public HTTP address.
	//
	// example:
	//
	// http://xxxxxxxx
	HttpApiInterUrl *string `json:"httpApiInterUrl,omitempty" xml:"httpApiInterUrl,omitempty"`
	// Private HTTP address.
	//
	// example:
	//
	// http://xxxxxxxx
	HttpApiIntraUrl *string `json:"httpApiIntraUrl,omitempty" xml:"httpApiIntraUrl,omitempty"`
	// Instance type, fixed value prom-view.
	//
	// example:
	//
	// prom-view
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// Payment type. Currently, the fixed value is FREE (free).
	//
	// example:
	//
	// FREE
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// Product that the prom instance belongs to.
	//
	// example:
	//
	// cms
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// Prometheus instance list.
	PrometheusInstances []*GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances `json:"prometheusInstances,omitempty" xml:"prometheusInstances,omitempty" type:"Repeated"`
	// Prometheus view ID.
	//
	// example:
	//
	// view-xxx
	PrometheusViewId *string `json:"prometheusViewId,omitempty" xml:"prometheusViewId,omitempty"`
	// Prometheus view name.
	//
	// example:
	//
	// view1
	PrometheusViewName *string `json:"prometheusViewName,omitempty" xml:"prometheusViewName,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Remote read public URL.
	//
	// example:
	//
	// http://workspace-default-cms-xxx-cn-hangzhou.cn-hangzhou.log.aliyuncs.com/prometheus/workspace-default-cms-xxx-cn-hangzhou/xxx/api/v1/read
	RemoteReadInterUrl *string `json:"remoteReadInterUrl,omitempty" xml:"remoteReadInterUrl,omitempty"`
	// Remote read intranet URL.
	//
	// example:
	//
	// http://workspace-default-cms-xxx-cn-hangzhou.cn-hangzhou-intranet.log.aliyuncs.com/prometheus/workspace-default-cms-xxx-cn-hangzhou/xxx/api/v1/read
	RemoteReadIntraUrl *string `json:"remoteReadIntraUrl,omitempty" xml:"remoteReadIntraUrl,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfm3gn5i6bigbi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Fixed value: PrometheusView
	//
	// example:
	//
	// RegistryModule
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// Backend data storage status
	//
	// example:
	//
	// Pending2Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Supported authentication types.
	SupportAuthTypes []*string `json:"supportAuthTypes,omitempty" xml:"supportAuthTypes,omitempty" type:"Repeated"`
	// Instance tag keys.
	Tags []*GetPrometheusViewResponseBodyPrometheusViewTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// User ID.
	//
	// example:
	//
	// 11222
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// Version.
	//
	// example:
	//
	// V1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// Workspace to which the environment belongs
	//
	// example:
	//
	// cms-monitor-test-aysls-pub-cn-fuzhou-monitor
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetPrometheusViewResponseBodyPrometheusView) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusViewResponseBodyPrometheusView) GoString() string {
	return s.String()
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetAuthFreeReadPolicy() *string {
	return s.AuthFreeReadPolicy
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetAuthToken() *string {
	return s.AuthToken
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetEnableAuthFreeRead() *bool {
	return s.EnableAuthFreeRead
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetEnableAuthToken() *bool {
	return s.EnableAuthToken
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetFolderUrl() *string {
	return s.FolderUrl
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetGrafanaInstanceId() *string {
	return s.GrafanaInstanceId
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetGrafanaInstanceName() *string {
	return s.GrafanaInstanceName
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetHttpApiInterUrl() *string {
	return s.HttpApiInterUrl
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetHttpApiIntraUrl() *string {
	return s.HttpApiIntraUrl
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetProduct() *string {
	return s.Product
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetPrometheusInstances() []*GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances {
	return s.PrometheusInstances
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetPrometheusViewId() *string {
	return s.PrometheusViewId
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetPrometheusViewName() *string {
	return s.PrometheusViewName
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetRemoteReadInterUrl() *string {
	return s.RemoteReadInterUrl
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetRemoteReadIntraUrl() *string {
	return s.RemoteReadIntraUrl
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetStatus() *string {
	return s.Status
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetSupportAuthTypes() []*string {
	return s.SupportAuthTypes
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetTags() []*GetPrometheusViewResponseBodyPrometheusViewTags {
	return s.Tags
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetUserId() *string {
	return s.UserId
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetVersion() *string {
	return s.Version
}

func (s *GetPrometheusViewResponseBodyPrometheusView) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetAuthFreeReadPolicy(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.AuthFreeReadPolicy = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetAuthToken(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.AuthToken = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetCreateTime(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.CreateTime = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetEnableAuthFreeRead(v bool) *GetPrometheusViewResponseBodyPrometheusView {
	s.EnableAuthFreeRead = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetEnableAuthToken(v bool) *GetPrometheusViewResponseBodyPrometheusView {
	s.EnableAuthToken = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetFolderUrl(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.FolderUrl = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetGrafanaInstanceId(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.GrafanaInstanceId = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetGrafanaInstanceName(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.GrafanaInstanceName = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetHttpApiInterUrl(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.HttpApiInterUrl = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetHttpApiIntraUrl(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.HttpApiIntraUrl = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetInstanceType(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.InstanceType = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetPaymentType(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.PaymentType = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetProduct(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.Product = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetPrometheusInstances(v []*GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances) *GetPrometheusViewResponseBodyPrometheusView {
	s.PrometheusInstances = v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetPrometheusViewId(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.PrometheusViewId = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetPrometheusViewName(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.PrometheusViewName = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetRegionId(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.RegionId = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetRemoteReadInterUrl(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.RemoteReadInterUrl = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetRemoteReadIntraUrl(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.RemoteReadIntraUrl = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetResourceGroupId(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.ResourceGroupId = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetResourceType(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.ResourceType = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetStatus(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.Status = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetSupportAuthTypes(v []*string) *GetPrometheusViewResponseBodyPrometheusView {
	s.SupportAuthTypes = v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetTags(v []*GetPrometheusViewResponseBodyPrometheusViewTags) *GetPrometheusViewResponseBodyPrometheusView {
	s.Tags = v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetUserId(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.UserId = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetVersion(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.Version = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) SetWorkspace(v string) *GetPrometheusViewResponseBodyPrometheusView {
	s.Workspace = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusView) Validate() error {
	if s.PrometheusInstances != nil {
		for _, item := range s.PrometheusInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances struct {
	// Instance ID.
	//
	// example:
	//
	// rw-63549e054ff596a4149927961dff
	PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-north-2-gov-1
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// User ID.
	//
	// example:
	//
	// 122xxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances) GoString() string {
	return s.String()
}

func (s *GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances) GetUserId() *string {
	return s.UserId
}

func (s *GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances) SetPrometheusInstanceId(v string) *GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances {
	s.PrometheusInstanceId = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances) SetRegionId(v string) *GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances {
	s.RegionId = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances) SetUserId(v string) *GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances {
	s.UserId = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusViewPrometheusInstances) Validate() error {
	return dara.Validate(s)
}

type GetPrometheusViewResponseBodyPrometheusViewTags struct {
	// PagerDuty integration key.
	//
	// example:
	//
	// global_score_series
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Tag value.
	//
	// example:
	//
	// 371293199010092839
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetPrometheusViewResponseBodyPrometheusViewTags) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusViewResponseBodyPrometheusViewTags) GoString() string {
	return s.String()
}

func (s *GetPrometheusViewResponseBodyPrometheusViewTags) GetKey() *string {
	return s.Key
}

func (s *GetPrometheusViewResponseBodyPrometheusViewTags) GetValue() *string {
	return s.Value
}

func (s *GetPrometheusViewResponseBodyPrometheusViewTags) SetKey(v string) *GetPrometheusViewResponseBodyPrometheusViewTags {
	s.Key = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusViewTags) SetValue(v string) *GetPrometheusViewResponseBodyPrometheusViewTags {
	s.Value = &v
	return s
}

func (s *GetPrometheusViewResponseBodyPrometheusViewTags) Validate() error {
	return dara.Validate(s)
}
