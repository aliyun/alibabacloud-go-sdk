// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMode(v string) *ModifyHybridCloudClusterRequest
	GetAccessMode() *string
	SetAccessRegion(v string) *ModifyHybridCloudClusterRequest
	GetAccessRegion() *string
	SetClusterName(v string) *ModifyHybridCloudClusterRequest
	GetClusterName() *string
	SetHttpPorts(v string) *ModifyHybridCloudClusterRequest
	GetHttpPorts() *string
	SetHttpsPorts(v string) *ModifyHybridCloudClusterRequest
	GetHttpsPorts() *string
	SetId(v int64) *ModifyHybridCloudClusterRequest
	GetId() *int64
	SetInstanceId(v string) *ModifyHybridCloudClusterRequest
	GetInstanceId() *string
	SetLogFieldsNotReturned(v string) *ModifyHybridCloudClusterRequest
	GetLogFieldsNotReturned() *string
	SetProtectionServerCount(v int32) *ModifyHybridCloudClusterRequest
	GetProtectionServerCount() *int32
	SetProxyStatus(v string) *ModifyHybridCloudClusterRequest
	GetProxyStatus() *string
	SetProxyType(v string) *ModifyHybridCloudClusterRequest
	GetProxyType() *string
	SetRegionId(v string) *ModifyHybridCloudClusterRequest
	GetRegionId() *string
	SetRemark(v string) *ModifyHybridCloudClusterRequest
	GetRemark() *string
	SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudClusterRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleConfig(v string) *ModifyHybridCloudClusterRequest
	GetRuleConfig() *string
	SetRuleStatus(v string) *ModifyHybridCloudClusterRequest
	GetRuleStatus() *string
	SetRuleType(v string) *ModifyHybridCloudClusterRequest
	GetRuleType() *string
}

type ModifyHybridCloudClusterRequest struct {
	// The network access mode of the hybrid cloud cluster. Valid values:
	//
	// - **internet**: access over the Internet.
	//
	// - **vpc**: access over a leased line through a virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// The region in which the leased line resides. This parameter is required when AccessMode is set to vpc. Valid values:
	//
	// - **cn-hangzhou**: Hangzhou.
	//
	// - **cn-beijing**: Beijing.
	//
	// - **cn-shanghai**: Shanghai.
	//
	// example:
	//
	// cn-beijing
	AccessRegion *string `json:"AccessRegion,omitempty" xml:"AccessRegion,omitempty"`
	// The name of the hybrid cloud cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-example-***
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The list of HTTP ports supported by the hybrid cloud cluster. Separate multiple ports with commas (,). Format: **port1,port2,port3**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80,8080
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The list of HTTPS ports supported by the hybrid cloud cluster. Separate multiple ports with commas (,). Format: **port1,port2,port3**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 443,8443
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// The ID of the hybrid cloud cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-mp9153****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The log fields that are excluded from the response.
	//
	// example:
	//
	// log_example
	LogFieldsNotReturned *string `json:"LogFieldsNotReturned,omitempty" xml:"LogFieldsNotReturned,omitempty"`
	// The maximum number of protection nodes that can be added to the hybrid cloud cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProtectionServerCount *int32 `json:"ProtectionServerCount,omitempty" xml:"ProtectionServerCount,omitempty"`
	// Indicates whether the proxy gateway is enabled. Valid values:
	//
	// - **on**: The proxy gateway is enabled.
	//
	// - **off**: The proxy gateway is disabled.
	//
	// example:
	//
	// off
	ProxyStatus *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	// The type of the hybrid cloud cluster. Valid values:
	//
	// - **cname**: a reverse proxy cluster.
	//
	// - **service**: a service cluster.
	//
	// example:
	//
	// service
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the hybrid cloud cluster.
	//
	// example:
	//
	// remarkExample
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the Alibaba Cloud resource group to which the WAF instance belongs.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The rule configuration in JSON format. This includes settings such as the circuit breaker, request body length limit, and timeout.
	//
	// example:
	//
	// {
	//
	//       "enable": true,
	//
	//       "param": {
	//
	//             "breaker": {
	//
	//                   "duration": 1,
	//
	//                   "failed": 1,
	//
	//                   "recent_failed": 1
	//
	//             },
	//
	//             "disable_protect": false,
	//
	//             "max_request_body_len": 1,
	//
	//             "timeout": 1
	//
	//       }
	//
	// }
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// Indicates whether the rule is enabled. Valid values:
	//
	// - **on**: The rule is enabled.
	//
	// - **off**: The rule is disabled.
	//
	// example:
	//
	// off
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The rule type. Valid values:
	//
	// - **bypass**: WAF bypasses security checks.
	//
	// example:
	//
	// bypass
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ModifyHybridCloudClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterRequest) GetAccessMode() *string {
	return s.AccessMode
}

func (s *ModifyHybridCloudClusterRequest) GetAccessRegion() *string {
	return s.AccessRegion
}

func (s *ModifyHybridCloudClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ModifyHybridCloudClusterRequest) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *ModifyHybridCloudClusterRequest) GetHttpsPorts() *string {
	return s.HttpsPorts
}

func (s *ModifyHybridCloudClusterRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyHybridCloudClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHybridCloudClusterRequest) GetLogFieldsNotReturned() *string {
	return s.LogFieldsNotReturned
}

func (s *ModifyHybridCloudClusterRequest) GetProtectionServerCount() *int32 {
	return s.ProtectionServerCount
}

func (s *ModifyHybridCloudClusterRequest) GetProxyStatus() *string {
	return s.ProxyStatus
}

func (s *ModifyHybridCloudClusterRequest) GetProxyType() *string {
	return s.ProxyType
}

func (s *ModifyHybridCloudClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHybridCloudClusterRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyHybridCloudClusterRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyHybridCloudClusterRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *ModifyHybridCloudClusterRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *ModifyHybridCloudClusterRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *ModifyHybridCloudClusterRequest) SetAccessMode(v string) *ModifyHybridCloudClusterRequest {
	s.AccessMode = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetAccessRegion(v string) *ModifyHybridCloudClusterRequest {
	s.AccessRegion = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetClusterName(v string) *ModifyHybridCloudClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetHttpPorts(v string) *ModifyHybridCloudClusterRequest {
	s.HttpPorts = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetHttpsPorts(v string) *ModifyHybridCloudClusterRequest {
	s.HttpsPorts = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetId(v int64) *ModifyHybridCloudClusterRequest {
	s.Id = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetInstanceId(v string) *ModifyHybridCloudClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetLogFieldsNotReturned(v string) *ModifyHybridCloudClusterRequest {
	s.LogFieldsNotReturned = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetProtectionServerCount(v int32) *ModifyHybridCloudClusterRequest {
	s.ProtectionServerCount = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetProxyStatus(v string) *ModifyHybridCloudClusterRequest {
	s.ProxyStatus = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetProxyType(v string) *ModifyHybridCloudClusterRequest {
	s.ProxyType = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetRegionId(v string) *ModifyHybridCloudClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetRemark(v string) *ModifyHybridCloudClusterRequest {
	s.Remark = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudClusterRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetRuleConfig(v string) *ModifyHybridCloudClusterRequest {
	s.RuleConfig = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetRuleStatus(v string) *ModifyHybridCloudClusterRequest {
	s.RuleStatus = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) SetRuleType(v string) *ModifyHybridCloudClusterRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyHybridCloudClusterRequest) Validate() error {
	return dara.Validate(s)
}
