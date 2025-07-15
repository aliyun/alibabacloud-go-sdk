// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectTrafficQosRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetClientToken() *string
	SetDstCidr(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetDstCidr() *string
	SetDstIPv6Cidr(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetDstIPv6Cidr() *string
	SetDstPortRange(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetDstPortRange() *string
	SetMatchDscp(v int32) *CreateExpressConnectTrafficQosRuleRequest
	GetMatchDscp() *int32
	SetOwnerAccount(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateExpressConnectTrafficQosRuleRequest
	GetOwnerId() *int64
	SetPriority(v int32) *CreateExpressConnectTrafficQosRuleRequest
	GetPriority() *int32
	SetProtocol(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetProtocol() *string
	SetQosId(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetQosId() *string
	SetQueueId(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetQueueId() *string
	SetRegionId(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetRegionId() *string
	SetRemarkingDscp(v int32) *CreateExpressConnectTrafficQosRuleRequest
	GetRemarkingDscp() *int32
	SetResourceOwnerAccount(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetResourceOwnerAccount() *string
	SetRuleDescription(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetRuleDescription() *string
	SetRuleName(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetRuleName() *string
	SetSrcCidr(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetSrcCidr() *string
	SetSrcIPv6Cidr(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetSrcIPv6Cidr() *string
	SetSrcPortRange(v string) *CreateExpressConnectTrafficQosRuleRequest
	GetSrcPortRange() *string
}

type CreateExpressConnectTrafficQosRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The destination IPv4 CIDR block that matches the QoS rule traffic.
	//
	// > When this parameter is unavailable, specify **SrcIPv6Cidr*	- or **DstIPv6Cidr**.
	//
	// example:
	//
	// 1.1.1.0/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination IPv6 CIDR block that matches the QoS rule traffic.
	//
	// > When this parameter is unavailable, specify **SrcCidr*	- or **DstCidr**.
	//
	// example:
	//
	// 2001:0db8:1234:5678::/64
	DstIPv6Cidr *string `json:"DstIPv6Cidr,omitempty" xml:"DstIPv6Cidr,omitempty"`
	// The range of destination ports that match the QoS rule traffic. Valid values: **0*	- to **65535**. If the traffic does not match, the value is -1. You can specify only one port. The start port number must be the same as the end port number. Different protocols correspond to different ports. Valid values:
	//
	// 	- **ALL*	- (uneditable): -1/-1.
	//
	// 	- **ICMP(IPv4)*	- (uneditable): -1/-1.
	//
	// 	- **ICMPv6(IPv6)*	- (uneditable): -1/-1.
	//
	// 	- **TCP*	- (editable): -1/-1.
	//
	// 	- **UDP*	- (editable): -1/-1.
	//
	// 	- **GRE*	- (uneditable): -1/-1.
	//
	// 	- **SSH*	- (uneditable): 22/22.
	//
	// 	- **Telnet*	- (uneditable): 23/23.
	//
	// 	- **HTTP*	- (uneditable): 80/80.
	//
	// 	- **HTTPS*	- (uneditable): 443/443.
	//
	// 	- **MS SQL*	- (uneditable): 1443/1443.
	//
	// 	- **Oracle*	- (uneditable): 1521/1521.
	//
	// 	- **MySql*	- (uneditable): 3306/3306.
	//
	// 	- **RDP*	- (uneditable): 3389/3389.
	//
	// 	- **PostgreSQL*	- (uneditable): 5432/5432.
	//
	// 	- **Redis*	- (uneditable): 6379/6379.
	//
	// example:
	//
	// -1/-1
	DstPortRange *string `json:"DstPortRange,omitempty" xml:"DstPortRange,omitempty"`
	// The DSCP value that matches the QoS rule traffic. Valid values: **0*	- to **63**. If no value is matched, the value is -1.
	//
	// example:
	//
	// 1
	MatchDscp    *int32  `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority of the QoS rule. Valid values: **1*	- to **9000**. A larger value indicates a higher priority. The priority of each QoS rule must be unique in the same QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The protocol of the QoS rule. Valid values:
	//
	// 	- **ALL**
	//
	// 	- **ICMP(IPv4)**
	//
	// 	- **ICMPv6(IPv6)**
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// 	- **GRE**
	//
	// 	- **SSH**
	//
	// 	- **Telnet**
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// 	- **MS SQL**
	//
	// 	- **Oracle**
	//
	// 	- **MySql**
	//
	// 	- **RDP**
	//
	// 	- **PostgreSQL**
	//
	// 	- **Redis**
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-2giu0a6vd5x0mv4700
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the QoS queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-queue-9nyx2u7n71s2rcy4n5
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The region ID of the QoS policy.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The new DSCP value. Valid values: **0*	- to **63**. If you do not change the value, set the value to -1.
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
	// The name of the QoS rule.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-rule-test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The source IPv4 CIDR block that matches the QoS rule traffic.
	//
	// > When this parameter is unavailable, specify **SrcIPv6Cidr*	- or **DstIPv6Cidr**.
	//
	// example:
	//
	// 1.1.1.0/24
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source IPv6 CIDR block that matches the QoS rule traffic.
	//
	// > When this parameter is unavailable, specify **SrcCidr*	- or **DstCidr**.
	//
	// example:
	//
	// 2001:0db8:1234:5678::/64
	SrcIPv6Cidr *string `json:"SrcIPv6Cidr,omitempty" xml:"SrcIPv6Cidr,omitempty"`
	// The range of source ports that match the QoS rule traffic. Valid values: **0*	- to **65535**. If the traffic does not match, the value is -1. You can specify only one port. The start port number must be the same as the end port number.
	//
	// example:
	//
	// -1/-1
	SrcPortRange *string `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty"`
}

func (s CreateExpressConnectTrafficQosRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectTrafficQosRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetDstCidr() *string {
	return s.DstCidr
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetDstIPv6Cidr() *string {
	return s.DstIPv6Cidr
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetDstPortRange() *string {
	return s.DstPortRange
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetMatchDscp() *int32 {
	return s.MatchDscp
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetQosId() *string {
	return s.QosId
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetQueueId() *string {
	return s.QueueId
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetRemarkingDscp() *int32 {
	return s.RemarkingDscp
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetRuleDescription() *string {
	return s.RuleDescription
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetSrcCidr() *string {
	return s.SrcCidr
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetSrcIPv6Cidr() *string {
	return s.SrcIPv6Cidr
}

func (s *CreateExpressConnectTrafficQosRuleRequest) GetSrcPortRange() *string {
	return s.SrcPortRange
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetClientToken(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetDstCidr(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.DstCidr = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetDstIPv6Cidr(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.DstIPv6Cidr = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetDstPortRange(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.DstPortRange = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetMatchDscp(v int32) *CreateExpressConnectTrafficQosRuleRequest {
	s.MatchDscp = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetOwnerAccount(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetOwnerId(v int64) *CreateExpressConnectTrafficQosRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetPriority(v int32) *CreateExpressConnectTrafficQosRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetProtocol(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.Protocol = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetQosId(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.QosId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetQueueId(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.QueueId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetRegionId(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetRemarkingDscp(v int32) *CreateExpressConnectTrafficQosRuleRequest {
	s.RemarkingDscp = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetResourceOwnerAccount(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetRuleDescription(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.RuleDescription = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetRuleName(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetSrcCidr(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.SrcCidr = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetSrcIPv6Cidr(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.SrcIPv6Cidr = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) SetSrcPortRange(v string) *CreateExpressConnectTrafficQosRuleRequest {
	s.SrcPortRange = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRuleRequest) Validate() error {
	return dara.Validate(s)
}
