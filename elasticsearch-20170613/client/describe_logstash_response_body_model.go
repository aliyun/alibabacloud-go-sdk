// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogstashResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLogstashResponseBody
	GetRequestId() *string
	SetResult(v *DescribeLogstashResponseBodyResult) *DescribeLogstashResponseBody
	GetResult() *DescribeLogstashResponseBodyResult
}

type DescribeLogstashResponseBody struct {
	// Detailed information about the instance.
	//
	// example:
	//
	// C9334241-4837-46C2-B24B-9BDC517318DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the instance.
	Result *DescribeLogstashResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeLogstashResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogstashResponseBody) GetResult() *DescribeLogstashResponseBodyResult {
	return s.Result
}

func (s *DescribeLogstashResponseBody) SetRequestId(v string) *DescribeLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogstashResponseBody) SetResult(v *DescribeLogstashResponseBodyResult) *DescribeLogstashResponseBody {
	s.Result = v
	return s
}

func (s *DescribeLogstashResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogstashResponseBodyResult struct {
	// The configuration information of the node.
	ExtendConfigs []map[string]interface{} `json:"ExtendConfigs,omitempty" xml:"ExtendConfigs,omitempty" type:"Repeated"`
	// The number of data nodes.
	//
	// example:
	//
	// rg-aekzvowej3i****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The key of the tag.
	Tags []*DescribeLogstashResponseBodyResultTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The status of the zone. Valid values:
	//
	// 	- ISOLATION: offline
	//
	// 	- NORMAL
	ZoneInfos []*DescribeLogstashResponseBodyResultZoneInfos `json:"ZoneInfos,omitempty" xml:"ZoneInfos,omitempty" type:"Repeated"`
	// The billing method of the instance. Valid values:
	//
	// 	- prepaid: subscription
	//
	// 	- postpaid: pay-as-you-go
	//
	// example:
	//
	// {"slowlog.threshold.warn": "2s","slowlog.threshold.info": "1s","slowlog.threshold.debug": "500ms","slowlog.threshold.trace": "100ms" }
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The state of the instance. Four states are supported:
	//
	// 	- Normal: active
	//
	// 	- Active: activating
	//
	// 	- Freeze: inactive
	//
	// 	- Invalid: invalid
	//
	// example:
	//
	// 2020-02-06T14:12:03.672Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// ls-cn-abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the zone where the node resides.
	EndpointList []*DescribeLogstashResponseBodyResultEndpointList `json:"endpointList,omitempty" xml:"endpointList,omitempty" type:"Repeated"`
	// The access information of the node.
	//
	// example:
	//
	// ls-cn-abc
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	NetworkConfig *DescribeLogstashResponseBodyResultNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	// The name of the instance.
	//
	// example:
	//
	// 2
	NodeAmount *int32 `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	// The specifications of the node.
	NodeSpec *DescribeLogstashResponseBodyResultNodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// prepaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the elastic container instances belong.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The edition of the dedicated KMS instance.
	//
	// example:
	//
	// 2020-02-06T14:22:36.850Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 7.4.0_with_X-Pack
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The time when the instance was last updated.
	//
	// example:
	//
	// vpc-bp16k1dvzxtmagcva****
	VpcInstanceId *string `json:"vpcInstanceId,omitempty" xml:"vpcInstanceId,omitempty"`
}

func (s DescribeLogstashResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogstashResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResult) GetExtendConfigs() []map[string]interface{} {
	return s.ExtendConfigs
}

func (s *DescribeLogstashResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLogstashResponseBodyResult) GetTags() []*DescribeLogstashResponseBodyResultTags {
	return s.Tags
}

func (s *DescribeLogstashResponseBodyResult) GetZoneInfos() []*DescribeLogstashResponseBodyResultZoneInfos {
	return s.ZoneInfos
}

func (s *DescribeLogstashResponseBodyResult) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *DescribeLogstashResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeLogstashResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeLogstashResponseBodyResult) GetEndpointList() []*DescribeLogstashResponseBodyResultEndpointList {
	return s.EndpointList
}

func (s *DescribeLogstashResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLogstashResponseBodyResult) GetNetworkConfig() *DescribeLogstashResponseBodyResultNetworkConfig {
	return s.NetworkConfig
}

func (s *DescribeLogstashResponseBodyResult) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *DescribeLogstashResponseBodyResult) GetNodeSpec() *DescribeLogstashResponseBodyResultNodeSpec {
	return s.NodeSpec
}

func (s *DescribeLogstashResponseBodyResult) GetPaymentType() *string {
	return s.PaymentType
}

func (s *DescribeLogstashResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeLogstashResponseBodyResult) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DescribeLogstashResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *DescribeLogstashResponseBodyResult) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeLogstashResponseBodyResult) SetExtendConfigs(v []map[string]interface{}) *DescribeLogstashResponseBodyResult {
	s.ExtendConfigs = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetResourceGroupId(v string) *DescribeLogstashResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetTags(v []*DescribeLogstashResponseBodyResultTags) *DescribeLogstashResponseBodyResult {
	s.Tags = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetZoneInfos(v []*DescribeLogstashResponseBodyResultZoneInfos) *DescribeLogstashResponseBodyResult {
	s.ZoneInfos = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetConfig(v map[string]interface{}) *DescribeLogstashResponseBodyResult {
	s.Config = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetCreatedAt(v string) *DescribeLogstashResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetDescription(v string) *DescribeLogstashResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetEndpointList(v []*DescribeLogstashResponseBodyResultEndpointList) *DescribeLogstashResponseBodyResult {
	s.EndpointList = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetInstanceId(v string) *DescribeLogstashResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetNetworkConfig(v *DescribeLogstashResponseBodyResultNetworkConfig) *DescribeLogstashResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetNodeAmount(v int32) *DescribeLogstashResponseBodyResult {
	s.NodeAmount = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetNodeSpec(v *DescribeLogstashResponseBodyResultNodeSpec) *DescribeLogstashResponseBodyResult {
	s.NodeSpec = v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetPaymentType(v string) *DescribeLogstashResponseBodyResult {
	s.PaymentType = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetStatus(v string) *DescribeLogstashResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetUpdatedAt(v string) *DescribeLogstashResponseBodyResult {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetVersion(v string) *DescribeLogstashResponseBodyResult {
	s.Version = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) SetVpcInstanceId(v string) *DescribeLogstashResponseBodyResult {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeLogstashResponseBodyResult) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ZoneInfos != nil {
		for _, item := range s.ZoneInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EndpointList != nil {
		for _, item := range s.EndpointList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkConfig != nil {
		if err := s.NetworkConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NodeSpec != nil {
		if err := s.NodeSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogstashResponseBodyResultTags struct {
	// The value of the tag.
	//
	// example:
	//
	// env
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The information about the zones.
	//
	// example:
	//
	// dev
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s DescribeLogstashResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogstashResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResultTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeLogstashResponseBodyResultTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeLogstashResponseBodyResultTags) SetTagKey(v string) *DescribeLogstashResponseBodyResultTags {
	s.TagKey = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultTags) SetTagValue(v string) *DescribeLogstashResponseBodyResultTags {
	s.TagValue = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}

type DescribeLogstashResponseBodyResultZoneInfos struct {
	// The zone ID of the new instance.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The configuration of cluster extension parameters.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s DescribeLogstashResponseBodyResultZoneInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogstashResponseBodyResultZoneInfos) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResultZoneInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeLogstashResponseBodyResultZoneInfos) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeLogstashResponseBodyResultZoneInfos) SetStatus(v string) *DescribeLogstashResponseBodyResultZoneInfos {
	s.Status = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultZoneInfos) SetZoneId(v string) *DescribeLogstashResponseBodyResultZoneInfos {
	s.ZoneId = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultZoneInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeLogstashResponseBodyResultEndpointList struct {
	// The tags added to the ALB instance.
	//
	// example:
	//
	// ``172.16.**.**``
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 9600
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The port number.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s DescribeLogstashResponseBodyResultEndpointList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogstashResponseBodyResultEndpointList) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResultEndpointList) GetHost() *string {
	return s.Host
}

func (s *DescribeLogstashResponseBodyResultEndpointList) GetPort() *string {
	return s.Port
}

func (s *DescribeLogstashResponseBodyResultEndpointList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeLogstashResponseBodyResultEndpointList) SetHost(v string) *DescribeLogstashResponseBodyResultEndpointList {
	s.Host = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultEndpointList) SetPort(v string) *DescribeLogstashResponseBodyResultEndpointList {
	s.Port = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultEndpointList) SetZoneId(v string) *DescribeLogstashResponseBodyResultEndpointList {
	s.ZoneId = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultEndpointList) Validate() error {
	return dara.Validate(s)
}

type DescribeLogstashResponseBodyResultNetworkConfig struct {
	// The ID of the vSwitch to which the instance is connected.
	//
	// example:
	//
	// vpc
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The zone where the cluster resides.
	//
	// example:
	//
	// vpc-bp16k1dvzxtmagcva****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The network type of the instance. Valid values: Currently, only Virtual Private Cloud (VPC) are supported.
	//
	// example:
	//
	// cn-hangzhou-*
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// example:
	//
	// vsw-bp1k4ec6s7sjdbudw****
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s DescribeLogstashResponseBodyResultNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogstashResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) GetType() *string {
	return s.Type
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) GetVsArea() *string {
	return s.VsArea
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) SetType(v string) *DescribeLogstashResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) SetVpcId(v string) *DescribeLogstashResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) SetVsArea(v string) *DescribeLogstashResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) SetVswitchId(v string) *DescribeLogstashResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLogstashResponseBodyResultNodeSpec struct {
	// Whether to use disk encryption:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// 20
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The disk type of the node.
	//
	// example:
	//
	// true
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// The network configurations.
	//
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// The disk size of the node.
	//
	// example:
	//
	// elasticsearch.sn1ne.large
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeLogstashResponseBodyResultNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogstashResponseBodyResultNodeSpec) GoString() string {
	return s.String()
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) GetDisk() *int32 {
	return s.Disk
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) GetSpec() *string {
	return s.Spec
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) SetDisk(v int32) *DescribeLogstashResponseBodyResultNodeSpec {
	s.Disk = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) SetDiskEncryption(v bool) *DescribeLogstashResponseBodyResultNodeSpec {
	s.DiskEncryption = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) SetDiskType(v string) *DescribeLogstashResponseBodyResultNodeSpec {
	s.DiskType = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) SetSpec(v string) *DescribeLogstashResponseBodyResultNodeSpec {
	s.Spec = &v
	return s
}

func (s *DescribeLogstashResponseBodyResultNodeSpec) Validate() error {
	return dara.Validate(s)
}
