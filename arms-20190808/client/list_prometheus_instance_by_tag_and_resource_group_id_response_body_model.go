// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusInstanceByTagAndResourceGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody
	GetCode() *string
	SetData(v *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody
	GetData() *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData
	SetMessage(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody
	GetRequestId() *string
}

type ListPrometheusInstanceByTagAndResourceGroupIdResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4789C3E9-A85A-524B-B97B-9D2B14BA06BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) GetData() *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData {
	return s.Data
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) SetCode(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) SetData(v *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) SetMessage(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) SetRequestId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData struct {
	// The queried Prometheus instances.
	PrometheusInstances []*ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances `json:"PrometheusInstances,omitempty" xml:"PrometheusInstances,omitempty" type:"Repeated"`
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData) GetPrometheusInstances() []*ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	return s.PrometheusInstances
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData) SetPrometheusInstances(v []*ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData {
	s.PrometheusInstances = v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances struct {
	// The authorization token.
	//
	// example:
	//
	// ad32dxxxx
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// The ID of the Prometheus instance.
	//
	// example:
	//
	// c9d5dda1aeca64220853ace304baeb03d
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the Prometheus instance.
	//
	// example:
	//
	// prom1
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The instance type. Valid values:
	//
	// 	- remote-write: Prometheus instance for Remote Write
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
	// ecs
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The ID of the Grafana workspace.
	//
	// example:
	//
	// grafana-rnggfvhlcdl6m71l**
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitempty" xml:"GrafanaInstanceId,omitempty"`
	// The public URL for the HTTP API.
	//
	// example:
	//
	// http://cn-beijing.arms.aliyuncs.com:9090/api/v1/prometheus/xxx/cn-beijing
	HttpApiInterUrl *string `json:"HttpApiInterUrl,omitempty" xml:"HttpApiInterUrl,omitempty"`
	// The internal URL for the HTTP API.
	//
	// example:
	//
	// http://cn-beijing-intranet.arms.aliyuncs.com:9090/api/v1/prometheus/xxx/cn-beijing
	HttpApiIntraUrl *string `json:"HttpApiIntraUrl,omitempty" xml:"HttpApiIntraUrl,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PREPAY: subscription
	//
	// 	- POSTPAY: pay-as-you-go
	//
	// example:
	//
	// POSTPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The public URL for Pushgateway.
	//
	// example:
	//
	// http://cn-beijing.arms.aliyuncs.com/prometheus/xxx/api/v2
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
	// cn-shanghai
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
	// http://cn-beijing.arms.aliyuncs.com/prometheus/xxx/api/v3/write
	RemoteWriteInterUrl *string `json:"RemoteWriteInterUrl,omitempty" xml:"RemoteWriteInterUrl,omitempty"`
	// The internal URL for remote write.
	//
	// example:
	//
	// http://cn-beijing-intranet.arms.aliyuncs.com/prometheus/xxx/api/v3/write
	RemoteWriteIntraUrl *string `json:"RemoteWriteIntraUrl,omitempty" xml:"RemoteWriteIntraUrl,omitempty"`
	// The ID of the resource group to which the Prometheus instance belongs.
	//
	// example:
	//
	// rg-acfmz7nocpeidcy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// PROMETHEUS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-f8zd1toc10wmbi1v5rom
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The child instances of the global aggregation instance. The value is a JSON string.
	//
	// example:
	//
	// [ { "headers": {}, "regionId": "cn-hangzhou", "sourceType": "AlibabaPrometheus", "extras": {}, "clusterId": "c39a1048921e04fceb039db2fb\\*\\*\\*\\*", "sourceName": "arms-luyao-test", "dataSource": "", "userId": "167275301789\\*\\*\\*\\*" }, { "headers": {}, "regionId": "cn-beijing", "sourceType": "AlibabaPrometheus", "extras": {}, "clusterId": "c6b6485496d5b400abde22cb47b5\\*\\*\\*\\*", "sourceName": "agent-321-test", "dataSource": "", "userId": "167275301789\\*\\*\\*\\*" }, { "headers": {}, "regionId": "cn-zhangjiakou", "sourceType": "AlibabaPrometheus", "extras": {}, "clusterId": "c261a4f3200c446659133f1ade78\\*\\*\\*\\*", "sourceName": "zaifeng-cardinality-01", "dataSource": "", "userId": "167275301789\\*\\*\\*\\*" } ]
	SubClustersJson *string `json:"SubClustersJson,omitempty" xml:"SubClustersJson,omitempty"`
	// The list of tags.
	Tags []*ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the user.
	//
	// example:
	//
	// 1672753017899***
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-f8z73vcja1tqnw90aav5a
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-8vbtp1fsm8mir18l8rl0u
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetAuthToken() *string {
	return s.AuthToken
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetGrafanaInstanceId() *string {
	return s.GrafanaInstanceId
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetHttpApiInterUrl() *string {
	return s.HttpApiInterUrl
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetHttpApiIntraUrl() *string {
	return s.HttpApiIntraUrl
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetPushGatewayInterUrl() *string {
	return s.PushGatewayInterUrl
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetPushGatewayIntraUrl() *string {
	return s.PushGatewayIntraUrl
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetRemoteReadInterUrl() *string {
	return s.RemoteReadInterUrl
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetRemoteReadIntraUrl() *string {
	return s.RemoteReadIntraUrl
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetRemoteWriteInterUrl() *string {
	return s.RemoteWriteInterUrl
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetRemoteWriteIntraUrl() *string {
	return s.RemoteWriteIntraUrl
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetSubClustersJson() *string {
	return s.SubClustersJson
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetTags() []*ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags {
	return s.Tags
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetUserId() *string {
	return s.UserId
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetAuthToken(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.AuthToken = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetClusterId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetClusterName(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.ClusterName = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetClusterType(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.ClusterType = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetGrafanaInstanceId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.GrafanaInstanceId = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetHttpApiInterUrl(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.HttpApiInterUrl = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetHttpApiIntraUrl(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.HttpApiIntraUrl = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetPaymentType(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.PaymentType = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetPushGatewayInterUrl(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.PushGatewayInterUrl = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetPushGatewayIntraUrl(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.PushGatewayIntraUrl = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetRegionId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetRemoteReadInterUrl(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.RemoteReadInterUrl = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetRemoteReadIntraUrl(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.RemoteReadIntraUrl = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetRemoteWriteInterUrl(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.RemoteWriteInterUrl = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetRemoteWriteIntraUrl(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.RemoteWriteIntraUrl = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetResourceGroupId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetResourceType(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.ResourceType = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetSecurityGroupId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.SecurityGroupId = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetSubClustersJson(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.SubClustersJson = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetTags(v []*ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.Tags = v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetUserId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.UserId = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetVSwitchId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.VSwitchId = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) SetVpcId(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances {
	s.VpcId = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstances) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags struct {
	// The tag key.
	//
	// example:
	//
	// ac-cus-tag-3
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// asg-2vc8qq7x89o11rus9uvu
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags) SetTagKey(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags {
	s.TagKey = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags) SetTagValue(v string) *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags {
	s.TagValue = &v
	return s
}

func (s *ListPrometheusInstanceByTagAndResourceGroupIdResponseBodyDataPrometheusInstancesTags) Validate() error {
	return dara.Validate(s)
}
