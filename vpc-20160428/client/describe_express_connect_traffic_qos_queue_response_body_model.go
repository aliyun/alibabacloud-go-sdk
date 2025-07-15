// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectTrafficQosQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueueList(v []*DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) *DescribeExpressConnectTrafficQosQueueResponseBody
	GetQueueList() []*DescribeExpressConnectTrafficQosQueueResponseBodyQueueList
	SetRequestId(v string) *DescribeExpressConnectTrafficQosQueueResponseBody
	GetRequestId() *string
}

type DescribeExpressConnectTrafficQosQueueResponseBody struct {
	// The information about the QoS queues.
	QueueList []*DescribeExpressConnectTrafficQosQueueResponseBodyQueueList `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 606998F0-B94D-48FE-8316-ACA81BB230DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExpressConnectTrafficQosQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBody) GetQueueList() []*DescribeExpressConnectTrafficQosQueueResponseBodyQueueList {
	return s.QueueList
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBody) SetQueueList(v []*DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) *DescribeExpressConnectTrafficQosQueueResponseBody {
	s.QueueList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBody) SetRequestId(v string) *DescribeExpressConnectTrafficQosQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExpressConnectTrafficQosQueueResponseBodyQueueList struct {
	// The percentage of bandwidth allocated to a QoS queue.
	//
	// 	- If QueueType is set to **Medium**, this parameter is returned. Valid values: **1*	- to **100**.
	//
	// 	- If QueueType is set to **Default**, a value of - is returned.
	//
	// example:
	//
	// 100
	BandwidthPercent *string `json:"BandwidthPercent,omitempty" xml:"BandwidthPercent,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-ncfgzxg40zks5n0qze
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The description of the QoS queue.
	//
	// The name must be **0*	- to **256*	- characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-queue-test
	QueueDescription *string `json:"QueueDescription,omitempty" xml:"QueueDescription,omitempty"`
	// The ID of the QoS queue.
	//
	// example:
	//
	// qos-queue-9nyx2u7n71s2rcy4n5
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The name of the QoS queue.
	//
	// The name must be **0*	- to **128*	- characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-queue-test
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The priority of the QoS queue. Valid values:
	//
	// 	- **High**
	//
	// 	- **Medium**
	//
	// 	- **Default**
	//
	//
	//
	// >  You cannot create a QoS queue of the default priority.
	//
	// example:
	//
	// High
	QueueType *string `json:"QueueType,omitempty" xml:"QueueType,omitempty"`
	// The information about the QoS rules.
	RuleList []*DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// The state of the QoS queue. Valid values:
	//
	// 	- **Normal**: The QoS queue is available.
	//
	// 	- **Configuring**: The QoS queue is being configured.
	//
	// 	- **Deleting**: The QoS queue is being deleted.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) GetBandwidthPercent() *string {
	return s.BandwidthPercent
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) GetQosId() *string {
	return s.QosId
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) GetQueueDescription() *string {
	return s.QueueDescription
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) GetQueueId() *string {
	return s.QueueId
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) GetQueueName() *string {
	return s.QueueName
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) GetQueueType() *string {
	return s.QueueType
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) GetRuleList() []*DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	return s.RuleList
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) SetBandwidthPercent(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList {
	s.BandwidthPercent = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) SetQosId(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList {
	s.QosId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) SetQueueDescription(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList {
	s.QueueDescription = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) SetQueueId(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList {
	s.QueueId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) SetQueueName(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList {
	s.QueueName = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) SetQueueType(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList {
	s.QueueType = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) SetRuleList(v []*DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList {
	s.RuleList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) SetStatus(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueList) Validate() error {
	return dara.Validate(s)
}

type DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList struct {
	// The destination IPv4 CIDR block that matches the QoS rule traffic.
	//
	//
	//
	// > If the parameter is unavailable, specify **SrcIPv6Cidr*	- or **DstIPv6Cidr**.
	//
	// example:
	//
	// 1.1.1.0/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination IPv6 CIDR block that matches the QoS rule traffic.
	//
	//
	//
	// > If the parameter is unavailable, specify **SrcCidr*	- or **DstCidr**.
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
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The priority of the QoS rule. Valid values: **1*	- to **9000**. A larger value indicates a higher priority. The priority of each QoS rule must be unique in the same QoS policy.
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
	// example:
	//
	// ALL
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-91xz9f8zd7yj8xwknz
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the QoS queue.
	//
	// example:
	//
	// qos-queue-iugg0l9x27f2nocouj
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The new DSCP value. Valid values: **0*	- to **63**. If you do not change the value, the value is -1.
	//
	// example:
	//
	// 1
	RemarkingDscp *int32 `json:"RemarkingDscp,omitempty" xml:"RemarkingDscp,omitempty"`
	// The description of the QoS rule.
	//
	// The name must be **0*	- to **256*	- characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-rule-test
	RuleDescription *string `json:"RuleDescription,omitempty" xml:"RuleDescription,omitempty"`
	// The ID of the QoS rule.
	//
	// example:
	//
	// qos-rule-iugg0l9x27f2nocouj
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the QoS rule.
	//
	// The name must be **0*	- to **128*	- characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-rule-test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The source IPv4 CIDR block that matches the QoS rule traffic.
	//
	//
	//
	//  > If the parameter is unavailable, specify **SrcIPv6Cidr*	- or **DstIPv6Cidr**.
	//
	// example:
	//
	// 1.1.1.0/24
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source IPv6 CIDR block that matches the QoS rule traffic.
	//
	//
	//
	// > If the parameter is unavailable, specify **SrcCidr*	- or **DstCidr**.
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
	// The status of the QoS rule. Valid values:
	//
	// 	- **Normal**: The QoS queue is available.
	//
	// 	- **Configuring**: The QoS queue is being configured.
	//
	// 	- **Deleting**: The QoS queue is being deleted.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetDstCidr() *string {
	return s.DstCidr
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetDstIPv6Cidr() *string {
	return s.DstIPv6Cidr
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetDstPortRange() *string {
	return s.DstPortRange
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetMatchDscp() *int32 {
	return s.MatchDscp
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetQosId() *string {
	return s.QosId
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetQueueId() *string {
	return s.QueueId
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetRemarkingDscp() *int32 {
	return s.RemarkingDscp
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetRuleDescription() *string {
	return s.RuleDescription
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetSrcCidr() *string {
	return s.SrcCidr
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetSrcIPv6Cidr() *string {
	return s.SrcIPv6Cidr
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetSrcPortRange() *string {
	return s.SrcPortRange
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetDstCidr(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.DstCidr = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetDstIPv6Cidr(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.DstIPv6Cidr = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetDstPortRange(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.DstPortRange = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetMatchDscp(v int32) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.MatchDscp = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetPriority(v int32) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.Priority = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetProtocol(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.Protocol = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetQosId(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.QosId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetQueueId(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.QueueId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetRemarkingDscp(v int32) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.RemarkingDscp = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetRuleDescription(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.RuleDescription = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetRuleId(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.RuleId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetRuleName(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.RuleName = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetSrcCidr(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.SrcCidr = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetSrcIPv6Cidr(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.SrcIPv6Cidr = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetSrcPortRange(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.SrcPortRange = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) SetStatus(v string) *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueResponseBodyQueueListRuleList) Validate() error {
	return dara.Validate(s)
}
