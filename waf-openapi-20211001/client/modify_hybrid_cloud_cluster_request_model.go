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
	// This parameter is required.
	AccessMode   *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	AccessRegion *string `json:"AccessRegion,omitempty" xml:"AccessRegion,omitempty"`
	// The name of the cluster.
	//
	// This parameter is required.
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The HTTP ports that are supported. Set this parameter to a string. Specify multiple ports in the **port1,port2,port3*	- format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80,8080
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The HTTPS ports that are supported. Set this parameter to a string. Specify multiple ports in the **port1,port2,port3*	- format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 443,8443
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-mp9153****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LogFieldsNotReturned *string `json:"LogFieldsNotReturned,omitempty" xml:"LogFieldsNotReturned,omitempty"`
	// The number of protection nodes that can be added to the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProtectionServerCount *int32  `json:"ProtectionServerCount,omitempty" xml:"ProtectionServerCount,omitempty"`
	ProxyStatus           *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	ProxyType             *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks about the cluster.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	RuleConfig                     *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	RuleStatus                     *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	RuleType                       *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
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
