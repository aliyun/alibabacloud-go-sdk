// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridCloudClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMode(v string) *CreateHybridCloudClusterRequest
	GetAccessMode() *string
	SetAccessRegion(v string) *CreateHybridCloudClusterRequest
	GetAccessRegion() *string
	SetClusterName(v string) *CreateHybridCloudClusterRequest
	GetClusterName() *string
	SetHttpPorts(v string) *CreateHybridCloudClusterRequest
	GetHttpPorts() *string
	SetHttpsPorts(v string) *CreateHybridCloudClusterRequest
	GetHttpsPorts() *string
	SetInstanceId(v string) *CreateHybridCloudClusterRequest
	GetInstanceId() *string
	SetLogFieldsNotReturned(v string) *CreateHybridCloudClusterRequest
	GetLogFieldsNotReturned() *string
	SetProtectionServerCount(v int32) *CreateHybridCloudClusterRequest
	GetProtectionServerCount() *int32
	SetProxyStatus(v string) *CreateHybridCloudClusterRequest
	GetProxyStatus() *string
	SetProxyType(v string) *CreateHybridCloudClusterRequest
	GetProxyType() *string
	SetRegionId(v string) *CreateHybridCloudClusterRequest
	GetRegionId() *string
	SetRemark(v string) *CreateHybridCloudClusterRequest
	GetRemark() *string
	SetResourceManagerResourceGroupId(v string) *CreateHybridCloudClusterRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleConfig(v string) *CreateHybridCloudClusterRequest
	GetRuleConfig() *string
	SetRuleStatus(v string) *CreateHybridCloudClusterRequest
	GetRuleStatus() *string
	SetRuleType(v string) *CreateHybridCloudClusterRequest
	GetRuleType() *string
}

type CreateHybridCloudClusterRequest struct {
	// The network access mode of the cluster. Valid values:
	//
	// - **internet**: access over the Internet.
	//
	// - **vpc**: access over an Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// The region for Express Connect circuit access. Valid values:
	//
	// - **cn-hangzhou**: Hangzhou.
	//
	// - **cn-beijing**: Beijing.
	//
	// - **cn-shanghai**: Shanghai.
	//
	// example:
	//
	// cn-hangzhou
	AccessRegion *string `json:"AccessRegion,omitempty" xml:"AccessRegion,omitempty"`
	// The name of the hybrid cloud cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The listening ports for the HTTP protocol. Separate multiple ports with commas (,), such as **port1,port2,port3**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80,8080
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The listening ports for the HTTPS protocol. Separate multiple ports with commas (,), such as **port1,port2,port3**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 443,8443
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Deprecated
	//
	// > This parameter is deprecated. It no longer returns meaningful data.
	//
	// example:
	//
	// deprecated
	LogFieldsNotReturned *string `json:"LogFieldsNotReturned,omitempty" xml:"LogFieldsNotReturned,omitempty"`
	// The maximum number of protection nodes that can be added to the hybrid cloud cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProtectionServerCount *int32 `json:"ProtectionServerCount,omitempty" xml:"ProtectionServerCount,omitempty"`
	// Indicates whether the proxy gateway is enabled for the cluster. Valid values:
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
	// - **cname**: reverse proxy cluster. Traffic is forwarded through CNAME resolution.
	//
	// - **service**: transparent proxy cluster. Traffic is forwarded at the service level.
	//
	// example:
	//
	// cname
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
	// The description of the hybrid cloud cluster.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The configuration of the bypass rule, in JSON format. This includes settings such as circuit breaker thresholds, request body size limits, and timeout values.
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
	// Indicates whether manual bypass is enabled for the cluster. Valid values:
	//
	// - **on**: Manual bypass is enabled.
	//
	// - **off**: Manual bypass is disabled.
	//
	// example:
	//
	// on
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The type of the bypass rule. Valid values:
	//
	// - **bypass**: skips WAF security checks and allows traffic to pass through directly.
	//
	// example:
	//
	// bypass
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s CreateHybridCloudClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridCloudClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateHybridCloudClusterRequest) GetAccessMode() *string {
	return s.AccessMode
}

func (s *CreateHybridCloudClusterRequest) GetAccessRegion() *string {
	return s.AccessRegion
}

func (s *CreateHybridCloudClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateHybridCloudClusterRequest) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *CreateHybridCloudClusterRequest) GetHttpsPorts() *string {
	return s.HttpsPorts
}

func (s *CreateHybridCloudClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateHybridCloudClusterRequest) GetLogFieldsNotReturned() *string {
	return s.LogFieldsNotReturned
}

func (s *CreateHybridCloudClusterRequest) GetProtectionServerCount() *int32 {
	return s.ProtectionServerCount
}

func (s *CreateHybridCloudClusterRequest) GetProxyStatus() *string {
	return s.ProxyStatus
}

func (s *CreateHybridCloudClusterRequest) GetProxyType() *string {
	return s.ProxyType
}

func (s *CreateHybridCloudClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHybridCloudClusterRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateHybridCloudClusterRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateHybridCloudClusterRequest) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *CreateHybridCloudClusterRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *CreateHybridCloudClusterRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *CreateHybridCloudClusterRequest) SetAccessMode(v string) *CreateHybridCloudClusterRequest {
	s.AccessMode = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetAccessRegion(v string) *CreateHybridCloudClusterRequest {
	s.AccessRegion = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetClusterName(v string) *CreateHybridCloudClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetHttpPorts(v string) *CreateHybridCloudClusterRequest {
	s.HttpPorts = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetHttpsPorts(v string) *CreateHybridCloudClusterRequest {
	s.HttpsPorts = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetInstanceId(v string) *CreateHybridCloudClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetLogFieldsNotReturned(v string) *CreateHybridCloudClusterRequest {
	s.LogFieldsNotReturned = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetProtectionServerCount(v int32) *CreateHybridCloudClusterRequest {
	s.ProtectionServerCount = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetProxyStatus(v string) *CreateHybridCloudClusterRequest {
	s.ProxyStatus = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetProxyType(v string) *CreateHybridCloudClusterRequest {
	s.ProxyType = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetRegionId(v string) *CreateHybridCloudClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetRemark(v string) *CreateHybridCloudClusterRequest {
	s.Remark = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetResourceManagerResourceGroupId(v string) *CreateHybridCloudClusterRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetRuleConfig(v string) *CreateHybridCloudClusterRequest {
	s.RuleConfig = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetRuleStatus(v string) *CreateHybridCloudClusterRequest {
	s.RuleStatus = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) SetRuleType(v string) *CreateHybridCloudClusterRequest {
	s.RuleType = &v
	return s
}

func (s *CreateHybridCloudClusterRequest) Validate() error {
	return dara.Validate(s)
}
