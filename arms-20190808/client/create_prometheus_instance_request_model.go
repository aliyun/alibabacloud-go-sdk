// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllSubClustersSuccess(v bool) *CreatePrometheusInstanceRequest
	GetAllSubClustersSuccess() *bool
	SetArchiveDuration(v int32) *CreatePrometheusInstanceRequest
	GetArchiveDuration() *int32
	SetClusterId(v string) *CreatePrometheusInstanceRequest
	GetClusterId() *string
	SetClusterName(v string) *CreatePrometheusInstanceRequest
	GetClusterName() *string
	SetClusterType(v string) *CreatePrometheusInstanceRequest
	GetClusterType() *string
	SetDuration(v int32) *CreatePrometheusInstanceRequest
	GetDuration() *int32
	SetGrafanaInstanceId(v string) *CreatePrometheusInstanceRequest
	GetGrafanaInstanceId() *string
	SetPaymentType(v string) *CreatePrometheusInstanceRequest
	GetPaymentType() *string
	SetRegionId(v string) *CreatePrometheusInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePrometheusInstanceRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *CreatePrometheusInstanceRequest
	GetSecurityGroupId() *string
	SetSubClustersJson(v string) *CreatePrometheusInstanceRequest
	GetSubClustersJson() *string
	SetTags(v []*CreatePrometheusInstanceRequestTags) *CreatePrometheusInstanceRequest
	GetTags() []*CreatePrometheusInstanceRequestTags
	SetVSwitchId(v string) *CreatePrometheusInstanceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreatePrometheusInstanceRequest
	GetVpcId() *string
}

type CreatePrometheusInstanceRequest struct {
	// Specifies whether all sub-instances must pass validation before the GlobalView instance is created. Default value: false, which indicates that partial success is allowed.
	//
	// example:
	//
	// true
	AllSubClustersSuccess *bool `json:"AllSubClustersSuccess,omitempty" xml:"AllSubClustersSuccess,omitempty"`
	// The number of days to automatically archive data after the storage period expires. Valid values: 60, 90, 180, and 365. A value of 0 indicates that data is not archived.
	//
	// example:
	//
	// 90
	ArchiveDuration *int32 `json:"ArchiveDuration,omitempty" xml:"ArchiveDuration,omitempty"`
	// The Container Service cluster ID. This parameter is required when ClusterType is set to aliyun-cs.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster to create. This parameter is required when ClusterType is set to remote-write, ecs, or global-view.
	//
	// For ecs instances, the ClusterName must follow the format "name-vpc-id", and the name part cannot exceed 24 characters. Example: "mytest1-vpc-xxxxxxxxxxx".
	//
	// example:
	//
	// clusterNameOfTest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The instance type. Valid values:
	//
	// -  remote-write: Prometheus for Remote Write.
	//
	// -  ecs (no longer supported): Prometheus for ECS.
	//
	// -  global-view: Prometheus for GlobalView.
	//
	// -  aliyun-cs (no longer supported): Prometheus for Container Service.
	//
	// - cloud-product (no longer supported): Prometheus for Cloud Service.
	//
	// - cloud-monitor (no longer supported): Prometheus for Hybrid Cloud Monitoring.
	//
	// - flink (no longer supported): Prometheus for Flink.
	//
	// This parameter is required.
	//
	// example:
	//
	// remote-write
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The data storage duration, in days.
	//
	// example:
	//
	// 90
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the bound Grafana workspace. Set this parameter to "free" when you use the shared Grafana edition.
	//
	// example:
	//
	// grafana-bp1*****
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitempty" xml:"GrafanaInstanceId,omitempty"`
	// The Billable methods. Valid values:
	//
	// POSTPAY: pay-as-you-go based on the number of reported metrics.
	//
	// POSTPAY_GB: pay-as-you-go based on the volume of written metrics.
	//
	// Empty: uses the default billing method configured by the user. If no default is configured, the system defaults to billing based on the number of reported metrics.
	//
	// example:
	//
	// POSTPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The actual region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Network Security group ID. This parameter is required when ClusterType is set to ecs or aliyun-cs for a managed ASK cluster.
	//
	// example:
	//
	// sg-bp1********
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The JSON string of sub-instances for the GlobalView instance.
	//
	// example:
	//
	// 当clusterType为global-view时，需要传此参数：需要聚合的集群的信息列表；示例：
	//
	// [
	//
	//     {
	//
	//         "headers":{
	//
	//         },
	//
	//         "regionId":"cn-hangzhou",
	//
	//         "sourceType":"AlibabaPrometheus",
	//
	//         "extras":{
	//
	//         },
	//
	//         "clusterId":"c39a1048921e04f***********",
	//
	//         "sourceName":"arms-luyao-test",
	//
	//         "dataSource":"",
	//
	//         "userId":"1672753***********"
	//
	//     },
	//
	//     {
	//
	//         "headers":{
	//
	//         },
	//
	//         "regionId":"cn-beijing",
	//
	//         "sourceType":"AlibabaPrometheus",
	//
	//         "extras":{
	//
	//         },
	//
	//         "clusterId":"c6b6485496d5b40***********",
	//
	//         "sourceName":"agent-321-测试",
	//
	//         "dataSource":"",
	//
	//         "userId":"1672753***********"
	//
	//     },
	//
	//     {
	//
	//         "headers":{
	//
	//         },
	//
	//         "regionId":"cn-zhangjiakou",
	//
	//         "sourceType":"AlibabaPrometheus",
	//
	//         "extras":{
	//
	//         },
	//
	//         "clusterId":"c261a4f3200c446***********",
	//
	//         "sourceName":"zaifeng-cardinality-01",
	//
	//         "dataSource":"",
	//
	//         "userId":"1672753***********"
	//
	//     }
	//
	// ]
	SubClustersJson *string `json:"SubClustersJson,omitempty" xml:"SubClustersJson,omitempty"`
	// The custom tags.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "labelName":"labelValue"
	//
	//     },
	//
	//     {
	//
	//         "testName":"clusterA"
	//
	//     }
	//
	// ]
	Tags []*CreatePrometheusInstanceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The vSwitch ID. This parameter is required when ClusterType is set to ecs or aliyun-cs for a managed ASK cluster.
	//
	// example:
	//
	// vsw-bp1*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID. This parameter is required when ClusterType is set to ecs or aliyun-cs for a managed ASK cluster.
	//
	// example:
	//
	// vpc-rpn**********
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreatePrometheusInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePrometheusInstanceRequest) GetAllSubClustersSuccess() *bool {
	return s.AllSubClustersSuccess
}

func (s *CreatePrometheusInstanceRequest) GetArchiveDuration() *int32 {
	return s.ArchiveDuration
}

func (s *CreatePrometheusInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreatePrometheusInstanceRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreatePrometheusInstanceRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *CreatePrometheusInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreatePrometheusInstanceRequest) GetGrafanaInstanceId() *string {
	return s.GrafanaInstanceId
}

func (s *CreatePrometheusInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreatePrometheusInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePrometheusInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePrometheusInstanceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreatePrometheusInstanceRequest) GetSubClustersJson() *string {
	return s.SubClustersJson
}

func (s *CreatePrometheusInstanceRequest) GetTags() []*CreatePrometheusInstanceRequestTags {
	return s.Tags
}

func (s *CreatePrometheusInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreatePrometheusInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreatePrometheusInstanceRequest) SetAllSubClustersSuccess(v bool) *CreatePrometheusInstanceRequest {
	s.AllSubClustersSuccess = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetArchiveDuration(v int32) *CreatePrometheusInstanceRequest {
	s.ArchiveDuration = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetClusterId(v string) *CreatePrometheusInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetClusterName(v string) *CreatePrometheusInstanceRequest {
	s.ClusterName = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetClusterType(v string) *CreatePrometheusInstanceRequest {
	s.ClusterType = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetDuration(v int32) *CreatePrometheusInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetGrafanaInstanceId(v string) *CreatePrometheusInstanceRequest {
	s.GrafanaInstanceId = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetPaymentType(v string) *CreatePrometheusInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetRegionId(v string) *CreatePrometheusInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetResourceGroupId(v string) *CreatePrometheusInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetSecurityGroupId(v string) *CreatePrometheusInstanceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetSubClustersJson(v string) *CreatePrometheusInstanceRequest {
	s.SubClustersJson = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetTags(v []*CreatePrometheusInstanceRequestTags) *CreatePrometheusInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetVSwitchId(v string) *CreatePrometheusInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) SetVpcId(v string) *CreatePrometheusInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreatePrometheusInstanceRequest) Validate() error {
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

type CreatePrometheusInstanceRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrometheusInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreatePrometheusInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreatePrometheusInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreatePrometheusInstanceRequestTags) SetKey(v string) *CreatePrometheusInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreatePrometheusInstanceRequestTags) SetValue(v string) *CreatePrometheusInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *CreatePrometheusInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
