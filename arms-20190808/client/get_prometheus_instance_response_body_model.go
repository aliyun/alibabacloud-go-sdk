// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetPrometheusInstanceResponseBody
	GetCode() *int32
	SetData(v *GetPrometheusInstanceResponseBodyData) *GetPrometheusInstanceResponseBody
	GetData() *GetPrometheusInstanceResponseBodyData
	SetMessage(v string) *GetPrometheusInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPrometheusInstanceResponseBody
	GetRequestId() *string
}

type GetPrometheusInstanceResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. If another status code is returned, the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetPrometheusInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 52C422FD-6B43-524D-B8A1-A4693294318C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPrometheusInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrometheusInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetPrometheusInstanceResponseBody) GetData() *GetPrometheusInstanceResponseBodyData {
	return s.Data
}

func (s *GetPrometheusInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPrometheusInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPrometheusInstanceResponseBody) SetCode(v int32) *GetPrometheusInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetPrometheusInstanceResponseBody) SetData(v *GetPrometheusInstanceResponseBodyData) *GetPrometheusInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetPrometheusInstanceResponseBody) SetMessage(v string) *GetPrometheusInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetPrometheusInstanceResponseBody) SetRequestId(v string) *GetPrometheusInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPrometheusInstanceResponseBodyData struct {
	// The permission type. Valid values: readWrite, readOnly, and httpReadOnly
	//
	// example:
	//
	// readWrite
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The number of days for which data is automatically archived after the storage duration expires. Valid values: 60, 90, 180, and 365. 0 indicates that the data is not archived.
	//
	// example:
	//
	// 60
	ArchiveDuration *int32 `json:"ArchiveDuration,omitempty" xml:"ArchiveDuration,omitempty"`
	// The whitelist of IP addresses for which password-free read is enabled.
	//
	// example:
	//
	// null
	AuthFreeReadPolicy *string `json:"AuthFreeReadPolicy,omitempty" xml:"AuthFreeReadPolicy,omitempty"`
	// The whitelist of IP addresses for which password-free write is enabled.
	//
	// example:
	//
	// null
	AuthFreeWritePolicy *string `json:"AuthFreeWritePolicy,omitempty" xml:"AuthFreeWritePolicy,omitempty"`
	// The authorization token.
	//
	// example:
	//
	// GciOiJIUzI1NiJ9***
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// The ID of the Prometheus instance.
	//
	// example:
	//
	// c589a1b8db05c4561aefbb898ca8fb1cf
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the monitoring object.
	//
	// example:
	//
	// prom1
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// 	- remote-write: general-purpose Prometheus instance
	//
	// 	- ecs: Prometheus instances for ECS
	//
	// 	- cloud-monitor: Prometheus instance for Alibaba Cloud services in the Chinese mainland
	//
	// 	- cloud-product: Prometheus instance for Alibaba Cloud services outside the Chinese mainland
	//
	// 	- global-view: global aggregation instance
	//
	// 	- aliyun-cs: Prometheus instance for Container Service
	//
	// example:
	//
	// remote-write
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The data storage status at the backend.
	//
	// example:
	//
	// RUNNING
	DbInstanceStatus *string `json:"DbInstanceStatus,omitempty" xml:"DbInstanceStatus,omitempty"`
	// Indicates whether password-free read is enabled.
	//
	// example:
	//
	// false
	EnableAuthFreeRead *bool `json:"EnableAuthFreeRead,omitempty" xml:"EnableAuthFreeRead,omitempty"`
	// Indicates whether password-free write is enabled.
	//
	// example:
	//
	// false
	EnableAuthFreeWrite *bool `json:"EnableAuthFreeWrite,omitempty" xml:"EnableAuthFreeWrite,omitempty"`
	// Indicates whether access token authentication is enabled.
	//
	// example:
	//
	// true
	EnableAuthToken *string `json:"EnableAuthToken,omitempty" xml:"EnableAuthToken,omitempty"`
	// The extra information. This parameter is returned only for console requests.
	ExtraInfo map[string]*string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The ID of the Grafana workspace.
	//
	// example:
	//
	// grafana-rnggfvhlcdl6m71***
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitempty" xml:"GrafanaInstanceId,omitempty"`
	// The public URL for the HTTP API.
	//
	// example:
	//
	// https://cn-beijing.arms.aliyuncs.com:9443/api/v1/prometheus/xxx
	HttpApiInterUrl *string `json:"HttpApiInterUrl,omitempty" xml:"HttpApiInterUrl,omitempty"`
	// The internal URL for the HTTP API.
	//
	// example:
	//
	// http://cn-beijing-intranet.arms.aliyuncs.com:9090/api/v1/prometheus/xxx
	HttpApiIntraUrl       *string `json:"HttpApiIntraUrl,omitempty" xml:"HttpApiIntraUrl,omitempty"`
	OpenTelemetryInterUrl *string `json:"OpenTelemetryInterUrl,omitempty" xml:"OpenTelemetryInterUrl,omitempty"`
	OpenTelemetryIntraUrl *string `json:"OpenTelemetryIntraUrl,omitempty" xml:"OpenTelemetryIntraUrl,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PREPAY: subscription
	//
	// 	- POSTPAY: pay-as-you-go
	//
	// example:
	//
	// PREPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The time when the billing method was modified.
	//
	// example:
	//
	// 2025-02-26T06:05:01Z
	PaymentTypeUpdateTime *string `json:"PaymentTypeUpdateTime,omitempty" xml:"PaymentTypeUpdateTime,omitempty"`
	// The product to which the Prometheus instance belongs. Valid values: arms and cms.
	//
	// example:
	//
	// arms
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The public URL for Pushgateway.
	//
	// example:
	//
	// https://cn-beijing.arms.aliyuncs.com/prometheus/xxx/api/v2
	PushGatewayInterUrl *string `json:"PushGatewayInterUrl,omitempty" xml:"PushGatewayInterUrl,omitempty"`
	// The internal URL for Pushgateway.
	//
	// example:
	//
	// http://cn-beijing-intranet.arms.aliyuncs.com/prometheus/xxx/api/v2
	PushGatewayIntraUrl *string `json:"PushGatewayIntraUrl,omitempty" xml:"PushGatewayIntraUrl,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The public URL for remote read.
	//
	// example:
	//
	// http://cn-beijing.arms.aliyuncs.com:9090/api/v1/prometheus/xxx/api/v1/read
	RemoteReadInterUrl *string `json:"RemoteReadInterUrl,omitempty" xml:"RemoteReadInterUrl,omitempty"`
	// The internal URL for remote read.
	//
	// example:
	//
	// http://cn-beijing-intranet.arms.aliyuncs.com:9090/api/v1/prometheus/xxx/api/v1/read
	RemoteReadIntraUrl *string `json:"RemoteReadIntraUrl,omitempty" xml:"RemoteReadIntraUrl,omitempty"`
	// The public URL for remote write.
	//
	// example:
	//
	// https://cn-beijing.arms.aliyuncs.com/prometheus/xxx/api/v3/write
	RemoteWriteInterUrl *string `json:"RemoteWriteInterUrl,omitempty" xml:"RemoteWriteInterUrl,omitempty"`
	// The internal URL for remote write.
	//
	// example:
	//
	// http://cn-beijing-intranet.arms.aliyuncs.com/prometheus/xxx/api/v3/write
	RemoteWriteIntraUrl *string `json:"RemoteWriteIntraUrl,omitempty" xml:"RemoteWriteIntraUrl,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aek2vezare****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the resource. Set the value to PROMETHEUS.
	//
	// example:
	//
	// PROMETHEUS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the security group. This parameter is returned only for Prometheus instances for ECS.
	//
	// example:
	//
	// sg-8vbdgmf4nraiqa9bx0jo
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The data storage duration. Unit: days.
	//
	// example:
	//
	// 90
	StorageDuration *int32 `json:"StorageDuration,omitempty" xml:"StorageDuration,omitempty"`
	// The child instances of the global aggregation instance. The value is a JSON string.
	//
	// example:
	//
	// [{"headers":{},"regionId":"cn-hangzhou","sourceType":"AlibabaPrometheus","extras":{},"clusterId":"c39a1048921e04fceb039db2fbb73\\*\\*\\*","sourceName":"arms-luyao-test","dataSource":"","userId":"167275301789\\*\\*\\*"},{"headers":{},"regionId":"cn-beijing","sourceType":"AlibabaPrometheus","extras":{},"clusterId":"c6b6485496d5b400abde22cb47b5\\*\\*\\*\\*","sourceName":"agent-321-test","dataSource":"","userId":"1672753017899\\*\\*\\*"},{"headers":{},"regionId":"cn-zhangjiakou","sourceType":"AlibabaPrometheus","extras":{},"clusterId":"c261a4f3200c446659133f1ade789b15e","sourceName":"zaifeng-cardinality-01","dataSource":"","userId":"167275301789\\*\\*\\*"}]
	SubClustersJson *string `json:"SubClustersJson,omitempty" xml:"SubClustersJson,omitempty"`
	// The supported authentication types.
	SupportAuthTypes []*string `json:"SupportAuthTypes,omitempty" xml:"SupportAuthTypes,omitempty" type:"Repeated"`
	// The tags of the instance.
	Tags []*GetPrometheusInstanceResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 167275301789****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The vSwitch ID. This parameter is returned only for Prometheus instances for ECS.
	//
	// example:
	//
	// vsw-f8z73vcja1tqnw90aav5a
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Version
	//
	// example:
	//
	// V1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The VPC ID. This parameter is returned only for Prometheus instances for ECS.
	//
	// example:
	//
	// vpc-8vb02uk57qbcktqcvqqqj
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetPrometheusInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPrometheusInstanceResponseBodyData) GetAccessType() *string {
	return s.AccessType
}

func (s *GetPrometheusInstanceResponseBodyData) GetArchiveDuration() *int32 {
	return s.ArchiveDuration
}

func (s *GetPrometheusInstanceResponseBodyData) GetAuthFreeReadPolicy() *string {
	return s.AuthFreeReadPolicy
}

func (s *GetPrometheusInstanceResponseBodyData) GetAuthFreeWritePolicy() *string {
	return s.AuthFreeWritePolicy
}

func (s *GetPrometheusInstanceResponseBodyData) GetAuthToken() *string {
	return s.AuthToken
}

func (s *GetPrometheusInstanceResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetPrometheusInstanceResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetPrometheusInstanceResponseBodyData) GetClusterType() *string {
	return s.ClusterType
}

func (s *GetPrometheusInstanceResponseBodyData) GetDbInstanceStatus() *string {
	return s.DbInstanceStatus
}

func (s *GetPrometheusInstanceResponseBodyData) GetEnableAuthFreeRead() *bool {
	return s.EnableAuthFreeRead
}

func (s *GetPrometheusInstanceResponseBodyData) GetEnableAuthFreeWrite() *bool {
	return s.EnableAuthFreeWrite
}

func (s *GetPrometheusInstanceResponseBodyData) GetEnableAuthToken() *string {
	return s.EnableAuthToken
}

func (s *GetPrometheusInstanceResponseBodyData) GetExtraInfo() map[string]*string {
	return s.ExtraInfo
}

func (s *GetPrometheusInstanceResponseBodyData) GetGrafanaInstanceId() *string {
	return s.GrafanaInstanceId
}

func (s *GetPrometheusInstanceResponseBodyData) GetHttpApiInterUrl() *string {
	return s.HttpApiInterUrl
}

func (s *GetPrometheusInstanceResponseBodyData) GetHttpApiIntraUrl() *string {
	return s.HttpApiIntraUrl
}

func (s *GetPrometheusInstanceResponseBodyData) GetOpenTelemetryInterUrl() *string {
	return s.OpenTelemetryInterUrl
}

func (s *GetPrometheusInstanceResponseBodyData) GetOpenTelemetryIntraUrl() *string {
	return s.OpenTelemetryIntraUrl
}

func (s *GetPrometheusInstanceResponseBodyData) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetPrometheusInstanceResponseBodyData) GetPaymentTypeUpdateTime() *string {
	return s.PaymentTypeUpdateTime
}

func (s *GetPrometheusInstanceResponseBodyData) GetProduct() *string {
	return s.Product
}

func (s *GetPrometheusInstanceResponseBodyData) GetPushGatewayInterUrl() *string {
	return s.PushGatewayInterUrl
}

func (s *GetPrometheusInstanceResponseBodyData) GetPushGatewayIntraUrl() *string {
	return s.PushGatewayIntraUrl
}

func (s *GetPrometheusInstanceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPrometheusInstanceResponseBodyData) GetRemoteReadInterUrl() *string {
	return s.RemoteReadInterUrl
}

func (s *GetPrometheusInstanceResponseBodyData) GetRemoteReadIntraUrl() *string {
	return s.RemoteReadIntraUrl
}

func (s *GetPrometheusInstanceResponseBodyData) GetRemoteWriteInterUrl() *string {
	return s.RemoteWriteInterUrl
}

func (s *GetPrometheusInstanceResponseBodyData) GetRemoteWriteIntraUrl() *string {
	return s.RemoteWriteIntraUrl
}

func (s *GetPrometheusInstanceResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetPrometheusInstanceResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetPrometheusInstanceResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetPrometheusInstanceResponseBodyData) GetStorageDuration() *int32 {
	return s.StorageDuration
}

func (s *GetPrometheusInstanceResponseBodyData) GetSubClustersJson() *string {
	return s.SubClustersJson
}

func (s *GetPrometheusInstanceResponseBodyData) GetSupportAuthTypes() []*string {
	return s.SupportAuthTypes
}

func (s *GetPrometheusInstanceResponseBodyData) GetTags() []*GetPrometheusInstanceResponseBodyDataTags {
	return s.Tags
}

func (s *GetPrometheusInstanceResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetPrometheusInstanceResponseBodyData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetPrometheusInstanceResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetPrometheusInstanceResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetPrometheusInstanceResponseBodyData) SetAccessType(v string) *GetPrometheusInstanceResponseBodyData {
	s.AccessType = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetArchiveDuration(v int32) *GetPrometheusInstanceResponseBodyData {
	s.ArchiveDuration = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetAuthFreeReadPolicy(v string) *GetPrometheusInstanceResponseBodyData {
	s.AuthFreeReadPolicy = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetAuthFreeWritePolicy(v string) *GetPrometheusInstanceResponseBodyData {
	s.AuthFreeWritePolicy = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetAuthToken(v string) *GetPrometheusInstanceResponseBodyData {
	s.AuthToken = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetClusterId(v string) *GetPrometheusInstanceResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetClusterName(v string) *GetPrometheusInstanceResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetClusterType(v string) *GetPrometheusInstanceResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetDbInstanceStatus(v string) *GetPrometheusInstanceResponseBodyData {
	s.DbInstanceStatus = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetEnableAuthFreeRead(v bool) *GetPrometheusInstanceResponseBodyData {
	s.EnableAuthFreeRead = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetEnableAuthFreeWrite(v bool) *GetPrometheusInstanceResponseBodyData {
	s.EnableAuthFreeWrite = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetEnableAuthToken(v string) *GetPrometheusInstanceResponseBodyData {
	s.EnableAuthToken = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetExtraInfo(v map[string]*string) *GetPrometheusInstanceResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetGrafanaInstanceId(v string) *GetPrometheusInstanceResponseBodyData {
	s.GrafanaInstanceId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetHttpApiInterUrl(v string) *GetPrometheusInstanceResponseBodyData {
	s.HttpApiInterUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetHttpApiIntraUrl(v string) *GetPrometheusInstanceResponseBodyData {
	s.HttpApiIntraUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetOpenTelemetryInterUrl(v string) *GetPrometheusInstanceResponseBodyData {
	s.OpenTelemetryInterUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetOpenTelemetryIntraUrl(v string) *GetPrometheusInstanceResponseBodyData {
	s.OpenTelemetryIntraUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetPaymentType(v string) *GetPrometheusInstanceResponseBodyData {
	s.PaymentType = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetPaymentTypeUpdateTime(v string) *GetPrometheusInstanceResponseBodyData {
	s.PaymentTypeUpdateTime = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetProduct(v string) *GetPrometheusInstanceResponseBodyData {
	s.Product = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetPushGatewayInterUrl(v string) *GetPrometheusInstanceResponseBodyData {
	s.PushGatewayInterUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetPushGatewayIntraUrl(v string) *GetPrometheusInstanceResponseBodyData {
	s.PushGatewayIntraUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetRegionId(v string) *GetPrometheusInstanceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetRemoteReadInterUrl(v string) *GetPrometheusInstanceResponseBodyData {
	s.RemoteReadInterUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetRemoteReadIntraUrl(v string) *GetPrometheusInstanceResponseBodyData {
	s.RemoteReadIntraUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetRemoteWriteInterUrl(v string) *GetPrometheusInstanceResponseBodyData {
	s.RemoteWriteInterUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetRemoteWriteIntraUrl(v string) *GetPrometheusInstanceResponseBodyData {
	s.RemoteWriteIntraUrl = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetResourceGroupId(v string) *GetPrometheusInstanceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetResourceType(v string) *GetPrometheusInstanceResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetSecurityGroupId(v string) *GetPrometheusInstanceResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetStorageDuration(v int32) *GetPrometheusInstanceResponseBodyData {
	s.StorageDuration = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetSubClustersJson(v string) *GetPrometheusInstanceResponseBodyData {
	s.SubClustersJson = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetSupportAuthTypes(v []*string) *GetPrometheusInstanceResponseBodyData {
	s.SupportAuthTypes = v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetTags(v []*GetPrometheusInstanceResponseBodyDataTags) *GetPrometheusInstanceResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetUserId(v string) *GetPrometheusInstanceResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetVSwitchId(v string) *GetPrometheusInstanceResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetVersion(v string) *GetPrometheusInstanceResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) SetVpcId(v string) *GetPrometheusInstanceResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyData) Validate() error {
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

type GetPrometheusInstanceResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tagValue1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetPrometheusInstanceResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusInstanceResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetPrometheusInstanceResponseBodyDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetPrometheusInstanceResponseBodyDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetPrometheusInstanceResponseBodyDataTags) SetTagKey(v string) *GetPrometheusInstanceResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyDataTags) SetTagValue(v string) *GetPrometheusInstanceResponseBodyDataTags {
	s.TagValue = &v
	return s
}

func (s *GetPrometheusInstanceResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
