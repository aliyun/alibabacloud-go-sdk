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
	// Does it require all child instances to be verified successfully before creating a GlobalView instance. The default is false, which means partial success is possible.
	//
	// example:
	//
	// true
	AllSubClustersSuccess *bool `json:"AllSubClustersSuccess,omitempty" xml:"AllSubClustersSuccess,omitempty"`
	// The number of days for which data is automatically archived after the storage expires. Valid values: 60, 90, 180, and 365. 0 indicates that the data is not archived.
	//
	// example:
	//
	// 90
	ArchiveDuration *int32 `json:"ArchiveDuration,omitempty" xml:"ArchiveDuration,omitempty"`
	// The ID of the ACK cluster. This parameter is required if you set the ClusterType parameter to aliyun-cs.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the created cluster. This parameter is required if you set the ClusterType parameter to remote-write or ecs.
	//
	// example:
	//
	// clusterNameOfTest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the Prometheus instance. Valid values:
	//
	// 	- remote-write: Prometheus instance for Remote Write
	//
	// 	- ecs (unavailable): Prometheus instance for ECS
	//
	// 	- global-view: Prometheus instance for GlobalView
	//
	// 	- aliyun-cs: Prometheus instance for Container Service
	//
	// 	- cloud-product (unavailable): Prometheus instance for Alibaba Cloud services
	//
	// 	- cloud-monitor (unavailable): Prometheus instance for Hybrid Cloud Monitoring
	//
	// 	- flink (unavailable): Prometheus instance for Flink
	//
	// This parameter is required.
	//
	// example:
	//
	// remote-write
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The data storage duration. Unit: days.
	//
	// example:
	//
	// 90
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the Grafana dedicated instance. This parameter is available if you set the ClusterType parameter to ecs.
	//
	// example:
	//
	// grafana-bp1*****
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitempty" xml:"GrafanaInstanceId,omitempty"`
	// The billing mode. Valid values: POSTPAY: charges fees based on the amount of reported metric data. POSTPAY_GB: charges fees based on the amount of written metric data. Empty: The user-defined default billing mode is used. If you do not specify a default value, you are charged based on the amount of reported metric data.
	//
	// example:
	//
	// POSTPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The ID of the region. If you use a Prometheus instance to monitor an Alibaba Cloud service in China, this parameter must be set to cn-shanghai.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the custom resource group. You can configure this parameter to bind the instance to the resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group. This parameter is required if you set the ClusterType parameter to ecs.
	//
	// example:
	//
	// sg-bp1********
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// JSON string for child instances of the globalView instance.
	//
	// example:
	//
	// When the clusterType is global view, this parameter needs to be passed: a list of information about the clusters that need to be aggregated.
	//
	// Example:
	//
	// [
	//
	//   {
	//
	//     "Headers":{
	//
	//     },
	//
	//     "RegionId": "cn hangzhou",
	//
	//     "SourceType": "Alibaba Prometheus",
	//
	//     "Extras":{
	//
	//     },
	//
	//     "ClusterId": "c39a1048921e04f ****************",
	//
	//     "SourceName": "test1",
	//
	//     "DataSource": "",
	//
	//     "UserId": "1672753 ******************"
	//
	//   },
	//
	//   {
	//
	//     "Headers":{
	//
	//     },
	//
	//     "RegionId": "cn beijing",
	//
	//     "SourceType": "Alibaba Prometheus",
	//
	//     "Extras":{
	//
	//     },
	//
	//     "ClusterId": "c6b6485496d5b40 ****************",
	//
	//     "SourceName": "test2",
	//
	//     "DataSource": "",
	//
	//     "UserId": "1672753 ******************"
	//
	//   },
	//
	//   {
	//
	//     "Headers":{
	//
	//     },
	//
	//     "RegionId": "cn zhangjiakou",
	//
	//     "SourceType": "Alibaba Prometheus",
	//
	//     "Extras":{
	//
	//     },
	//
	//     "ClusterId": "c261a4f3200c446 ****************",
	//
	//     "SourceName": "test3",
	//
	//     "DataSource": "",
	//
	//     "UserId": "1672753 ******************"
	//
	//   }
	//
	// ]
	SubClustersJson *string `json:"SubClustersJson,omitempty" xml:"SubClustersJson,omitempty"`
	// The tags of the instance. You can configure this parameter to manage tags for the instance.
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
	// The ID of the vSwitch. This parameter is required if you set the ClusterType parameter to ecs.
	//
	// example:
	//
	// vsw-bp1*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of virtual private cloud (VPC). This parameter is required if you set the ClusterType parameter to ecs.
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
	return dara.Validate(s)
}

type CreatePrometheusInstanceRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
