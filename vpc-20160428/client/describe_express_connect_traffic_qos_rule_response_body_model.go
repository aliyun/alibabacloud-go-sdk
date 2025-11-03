// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectTrafficQosRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeExpressConnectTrafficQosRuleResponseBody
	GetRequestId() *string
	SetRuleList(v []*DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) *DescribeExpressConnectTrafficQosRuleResponseBody
	GetRuleList() []*DescribeExpressConnectTrafficQosRuleResponseBodyRuleList
}

type DescribeExpressConnectTrafficQosRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9C7FA9D6-72E0-48A9-A9C3-2DA8569CD5EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of QoS rules.
	RuleList []*DescribeExpressConnectTrafficQosRuleResponseBodyRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeExpressConnectTrafficQosRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBody) GetRuleList() []*DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	return s.RuleList
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBody) SetRequestId(v string) *DescribeExpressConnectTrafficQosRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBody) SetRuleList(v []*DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) *DescribeExpressConnectTrafficQosRuleResponseBody {
	s.RuleList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBody) Validate() error {
	if s.RuleList != nil {
		for _, item := range s.RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExpressConnectTrafficQosRuleResponseBodyRuleList struct {
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
	// qos-pksbqfmotl5hzqmhf8
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the QoS queue.
	//
	// example:
	//
	// qos-queue-9nyx2u7n71s2rcy4n5
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The new DSCP value. Valid values: **0*	- to **63**. If you do not change the value, set the value to -1.
	//
	// example:
	//
	// 1
	RemarkingDscp *int32 `json:"RemarkingDscp,omitempty" xml:"RemarkingDscp,omitempty"`
	// The description of the QoS rule.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
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
	// The status of the QoS rule. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **Configuring**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetDstCidr() *string {
	return s.DstCidr
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetDstIPv6Cidr() *string {
	return s.DstIPv6Cidr
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetDstPortRange() *string {
	return s.DstPortRange
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetMatchDscp() *int32 {
	return s.MatchDscp
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetQosId() *string {
	return s.QosId
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetQueueId() *string {
	return s.QueueId
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetRemarkingDscp() *int32 {
	return s.RemarkingDscp
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetRuleDescription() *string {
	return s.RuleDescription
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetSrcCidr() *string {
	return s.SrcCidr
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetSrcIPv6Cidr() *string {
	return s.SrcIPv6Cidr
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetSrcPortRange() *string {
	return s.SrcPortRange
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetDstCidr(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.DstCidr = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetDstIPv6Cidr(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.DstIPv6Cidr = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetDstPortRange(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.DstPortRange = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetMatchDscp(v int32) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.MatchDscp = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetPriority(v int32) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.Priority = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetProtocol(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.Protocol = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetQosId(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.QosId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetQueueId(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.QueueId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetRemarkingDscp(v int32) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.RemarkingDscp = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetRuleDescription(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.RuleDescription = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetRuleId(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.RuleId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetRuleName(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.RuleName = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetSrcCidr(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.SrcCidr = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetSrcIPv6Cidr(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.SrcIPv6Cidr = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetSrcPortRange(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.SrcPortRange = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) SetStatus(v string) *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleResponseBodyRuleList) Validate() error {
	return dara.Validate(s)
}
