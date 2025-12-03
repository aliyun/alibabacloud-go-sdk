// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerUDPListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetAclId() *string
	SetAclIds(v *DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetAclIds() *DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds
	SetAclStatus(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetAclStatus() *string
	SetAclType(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetAclType() *string
	SetBackendServerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetBandwidth() *int32
	SetDescription(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetDescription() *string
	SetHealthCheck(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckExp(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheckExp() *string
	SetHealthCheckInterval(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheckInterval() *int32
	SetHealthCheckReq(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthCheckReq() *string
	SetHealthyThreshold(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetLoadBalancerId() *string
	SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetMasterSlaveServerGroupId() *string
	SetProxyProtocolV2Enabled(v bool) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetProxyProtocolV2Enabled() *bool
	SetRequestId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetRequestId() *string
	SetScheduler(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetScheduler() *string
	SetStatus(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetStatus() *string
	SetTags(v *DescribeLoadBalancerUDPListenerAttributeResponseBodyTags) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetTags() *DescribeLoadBalancerUDPListenerAttributeResponseBodyTags
	SetUnhealthyThreshold(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody
	GetVServerGroupId() *string
}

type DescribeLoadBalancerUDPListenerAttributeResponseBody struct {
	// The ID of the network ACL.
	//
	// example:
	//
	// 123943****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The ID of the access control list (ACL).
	AclIds *DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Struct"`
	// Indicates whether access control is enabled. Valid values: **on*	- and **off**. Default value: off.
	//
	// example:
	//
	// off
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The type of the ACL. Valid values:
	//
	// 	- **white**: a whitelist. Only requests from the IP addresses or CIDR blocks in the network ACL are forwarded. Whitelists apply to scenarios in which you want to allow only specified IP addresses to access an application.
	//
	//     Your service may be adversely affected if the whitelist is not properly configured. After a whitelist is configured, only requests from IP addresses that are added to the whitelist are forwarded by the listener. If you enable a whitelist but do not add an IP address to the ACL, the listener forwards all requests.
	//
	// 	- **black**: a blacklist. All requests from the IP addresses or CIDR blocks in the network ACL are blocked. Blacklists apply to scenarios in which you want to deny access from specific IP addresses or CIDR blocks to an application.
	//
	//     If a blacklist is configured for a listener but no IP address is added to the blacklist, the listener forwards all requests.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The backend port used by the CLB instance.
	//
	// >  If the listener is associated with a vServer group, this parameter is not returned.
	//
	// example:
	//
	// 90
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s. Valid values:
	//
	// 	- **-1**: For a pay-by-data-transfer Internet-facing CLB instance, this parameter is set to -1. This indicates that the bandwidth of the listener is unlimited.
	//
	// 	- **1*	- to **5120**: For a pay-by-bandwidth Internet-facing CLB instance, you can specify the maximum bandwidth of each listener. The sum of maximum bandwidth of all listeners cannot exceed the maximum bandwidth of the CLB instance.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The description of the listener.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port that is used for health checks. Valid values: **1*	- to **65535**. If this parameter is not set, the port specified by BackendServerPort is used for health checks.
	//
	// >  This parameter takes effect only when the **HealthCheck*	- parameter is set to **on**.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period of a health check. If a backend Elastic Compute Service (ECS) instance does not return a health check response within the specified timeout period, the server fails the health check. Valid values: **1*	- to **300**. Unit: seconds.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The response string for UDP listener health checks. The string is up to 64 characters in length, and can contain letters and digits.
	//
	// example:
	//
	// ok
	HealthCheckExp *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	// The interval between two consecutive health checks. Valid values: **1*	- to **50**. Unit: seconds.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The request string for UDP listener health checks. The string is up to 64 characters in length, and can contain letters and digits.
	//
	// example:
	//
	// hello
	HealthCheckReq *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	// The healthy threshold. The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail*	- to **success**. Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The frontend port used by the CLB instance.
	//
	// example:
	//
	// 53
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-bp1rtfnodmywb43e*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the primary/secondary server group that is associated with the listener.
	//
	// example:
	//
	// rsp-0bfucw****
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	// Indicates whether the Proxy protocol is used to pass client IP addresses to backend servers. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	ProxyProtocolV2Enabled *bool `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr*	- (default): Backend servers with higher weights receive more requests than backend servers with lower weights.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The status of the listener. Valid values:
	//
	// 	- **running**
	//
	// 	- **stopped**
	//
	// example:
	//
	// stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags *DescribeLoadBalancerUDPListenerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The unhealthy threshold. The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success*	- to **fail**. Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// The ID of the vServer group that is associated with the listener.
	//
	// example:
	//
	// rsp-cige6j****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetAclIds() *DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds {
	return s.AclIds
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetAclType() *string {
	return s.AclType
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheckExp() *string {
	return s.HealthCheckExp
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetProxyProtocolV2Enabled() *bool {
	return s.ProxyProtocolV2Enabled
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetTags() *DescribeLoadBalancerUDPListenerAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetAclId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetAclIds(v *DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.AclIds = v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetAclStatus(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.AclStatus = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetAclType(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.AclType = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckExp(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckExp = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthCheckReq(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthCheckReq = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetLoadBalancerId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetProxyProtocolV2Enabled(v bool) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetTags(v *DescribeLoadBalancerUDPListenerAttributeResponseBodyTags) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) SetVServerGroupId(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBody) Validate() error {
	if s.AclIds != nil {
		if err := s.AclIds.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds struct {
	AclId []*string `json:"AclId,omitempty" xml:"AclId,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds) GetAclId() []*string {
	return s.AclId
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds) SetAclId(v []*string) *DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds {
	s.AclId = v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBodyAclIds) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerUDPListenerAttributeResponseBodyTags struct {
	Tag []*DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBodyTags) GetTag() []*DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBodyTags) SetTag(v []*DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag) *DescribeLoadBalancerUDPListenerAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBodyTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag struct {
	// The key of tag N. Valid values of N: **1*	- to **20**. The tag value cannot be an empty string. The tag key can be up to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of tag N. Valid values of N: **1*	- to **20**. The tag value can be an empty string. The tag value can be up to 128 characters in length, and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag) SetTagKey(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag) SetTagValue(v string) *DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
