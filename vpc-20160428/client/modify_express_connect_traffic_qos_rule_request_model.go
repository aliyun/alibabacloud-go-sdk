// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectTrafficQosRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetClientToken() *string
	SetDstCidr(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetDstCidr() *string
	SetDstIPv6Cidr(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetDstIPv6Cidr() *string
	SetDstPortRange(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetDstPortRange() *string
	SetMatchDscp(v int32) *ModifyExpressConnectTrafficQosRuleRequest
	GetMatchDscp() *int32
	SetOwnerAccount(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyExpressConnectTrafficQosRuleRequest
	GetOwnerId() *int64
	SetPriority(v int32) *ModifyExpressConnectTrafficQosRuleRequest
	GetPriority() *int32
	SetProtocol(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetProtocol() *string
	SetQosId(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetQosId() *string
	SetQueueId(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetQueueId() *string
	SetRegionId(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetRegionId() *string
	SetRemarkingDscp(v int32) *ModifyExpressConnectTrafficQosRuleRequest
	GetRemarkingDscp() *int32
	SetResourceOwnerAccount(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetResourceOwnerAccount() *string
	SetRuleDescription(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetRuleDescription() *string
	SetRuleId(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetRuleName() *string
	SetSrcCidr(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetSrcCidr() *string
	SetSrcIPv6Cidr(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetSrcIPv6Cidr() *string
	SetSrcPortRange(v string) *ModifyExpressConnectTrafficQosRuleRequest
	GetSrcPortRange() *string
}

type ModifyExpressConnectTrafficQosRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The destination IP address IPv4 CIDR block for traffic matching in the QoS rule.
	//
	// > This parameter cannot be specified together with **SrcIPv6Cidr*	- or **DstIPv6Cidr**.
	//
	// example:
	//
	// ``1.1.**.**``/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination IP address IPv6 CIDR block for traffic matching in the QoS rule.
	//
	// > This parameter cannot be specified together with **SrcCidr*	- or **DstCidr**.
	//
	// example:
	//
	// 2001:0db8:1234:****::/64
	DstIPv6Cidr *string `json:"DstIPv6Cidr,omitempty" xml:"DstIPv6Cidr,omitempty"`
	// The destination port range for traffic matching in the QoS rule. Valid values: **0*	- to **65535**. Set the value to -1 if no matching is required. Currently, only a single port number can be specified. The start and end port numbers must be the same. The destination port range is fixed for each protocol type. Valid values:
	//
	// - **ALL**: -1/-1. Not editable.
	//
	// - **ICMP(IPv4)**: -1/-1. Not editable.
	//
	// - **ICMPv6(IPv6)**: -1/-1. Not editable.
	//
	// - **TCP**: -1/-1. Editable.
	//
	// - **UDP**: -1/-1. Editable.
	//
	// - **GRE**: -1/-1. Not editable.
	//
	// - **SSH**: 22/22. Not editable.
	//
	// - **Telnet**: 23/23. Not editable.
	//
	// - **HTTP**: 80/80. Not editable.
	//
	// - **HTTPS**: 443/443. Not editable.
	//
	// - **MS SQL**: 1443/1443. Not editable.
	//
	// - **Oracle**: 1521/1521. Not editable.
	//
	// - **MySql**: 3306/3306. Not editable.
	//
	// - **RDP**: 3389/3389. Not editable.
	//
	// - **PostgreSQL**: 5432/5432. Not editable.
	//
	// - **Redis**: 6379/6379. Not editable.
	//
	// example:
	//
	// -1/-1
	DstPortRange *string `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty"`
	// The DSCP value for traffic matching in the QoS rule. Valid values: **0*	- to **63**. Set the value to -1 if no matching is required.
	//
	// example:
	//
	// 1
	MatchDscp    *int32  `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority of the QoS rule. Valid values: **1*	- to **9000**. A larger value indicates a higher priority. The priority of each QoS rule must be unique within the same QoS policy.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The protocol type of the QoS rule. Valid values:
	//
	// - **ALL**
	//
	// - **ICMP(IPv4)**
	//
	// - **ICMPv6(IPv6)**
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **GRE**
	//
	// - **SSH**
	//
	// - **Telnet**
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// - **MS SQL**
	//
	// - **Oracle**
	//
	// - **MySql**
	//
	// - **RDP**
	//
	// - **PostgreSQL**
	//
	// - **Redis**
	//
	// example:
	//
	// ALL
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The QoS policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-2giu0a6vd5x0mv****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The QoS queue ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-queue-9nyx2u7n71s2rc****
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The region ID of the QoS policy.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The new DSCP value to remark in the traffic. Valid values: **0*	- to **63**. Set the value to -1 if no remarking is required.
	//
	// example:
	//
	// 1
	RemarkingDscp        *int32  `json:"RemarkingDscp,omitempty" xml:"RemarkingDscp,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The description of the QoS rule.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-rule-test
	RuleDescription *string `json:"RuleDescription,omitempty" xml:"RuleDescription,omitempty"`
	// The QoS rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-rule-iugg0l9x27f2no****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the QoS rule.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-rule-test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The source IPv4 CIDR block for traffic matching in the QoS rule.
	//
	// > This parameter cannot be specified together with **SrcIPv6Cidr*	- or **DstIPv6Cidr**.
	//
	// example:
	//
	// ``1.1.**.**``/24
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source IPv6 CIDR block for traffic matching in the QoS rule.
	//
	// > This parameter cannot be specified together with **SrcCidr*	- or **DstCidr**.
	//
	// example:
	//
	// 2001:0db8:1234:****::/64
	SrcIPv6Cidr *string `json:"SrcIPv6Cidr,omitempty" xml:"SrcIPv6Cidr,omitempty"`
	// The source port range for traffic matching in the QoS rule. Valid values: **0*	- to **65535**. Set the value to -1 if no matching is required. Currently, only a single port number can be specified. The start and end port numbers must be the same.
	//
	// example:
	//
	// -1/-1
	SrcPortRange *string `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty"`
}

func (s ModifyExpressConnectTrafficQosRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectTrafficQosRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetDstCidr() *string {
	return s.DstCidr
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetDstIPv6Cidr() *string {
	return s.DstIPv6Cidr
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetDstPortRange() *string {
	return s.DstPortRange
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetMatchDscp() *int32 {
	return s.MatchDscp
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetQosId() *string {
	return s.QosId
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetQueueId() *string {
	return s.QueueId
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetRemarkingDscp() *int32 {
	return s.RemarkingDscp
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetRuleDescription() *string {
	return s.RuleDescription
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetSrcCidr() *string {
	return s.SrcCidr
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetSrcIPv6Cidr() *string {
	return s.SrcIPv6Cidr
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) GetSrcPortRange() *string {
	return s.SrcPortRange
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetClientToken(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetDstCidr(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.DstCidr = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetDstIPv6Cidr(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.DstIPv6Cidr = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetDstPortRange(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.DstPortRange = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetMatchDscp(v int32) *ModifyExpressConnectTrafficQosRuleRequest {
	s.MatchDscp = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetOwnerAccount(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetOwnerId(v int64) *ModifyExpressConnectTrafficQosRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetPriority(v int32) *ModifyExpressConnectTrafficQosRuleRequest {
	s.Priority = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetProtocol(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.Protocol = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetQosId(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.QosId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetQueueId(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.QueueId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetRegionId(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetRemarkingDscp(v int32) *ModifyExpressConnectTrafficQosRuleRequest {
	s.RemarkingDscp = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetResourceOwnerAccount(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetRuleDescription(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.RuleDescription = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetRuleId(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetRuleName(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetSrcCidr(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.SrcCidr = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetSrcIPv6Cidr(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.SrcIPv6Cidr = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) SetSrcPortRange(v string) *ModifyExpressConnectTrafficQosRuleRequest {
	s.SrcPortRange = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRuleRequest) Validate() error {
	return dara.Validate(s)
}
