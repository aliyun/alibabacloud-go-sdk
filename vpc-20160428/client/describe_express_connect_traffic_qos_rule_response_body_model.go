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
	// The QoS rules.
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
	// The destination IP address IPv4 CIDR block for QoS rule traffic matching.
	//
	// > You cannot specify this parameter together with **SrcIPv6Cidr*	- or **DstIPv6Cidr**.
	//
	// example:
	//
	// ``1.1.**.**``/24
	DstCidr *string `json:"DstCidr,omitempty" xml:"DstCidr,omitempty"`
	// The destination IP address IPv6 CIDR block for QoS rule traffic matching.
	//
	// > You cannot specify this parameter together with **SrcCidr*	- or **DstCidr**.
	//
	// example:
	//
	// 2001:0db8:1234:****::/64
	DstIPv6Cidr *string `json:"DstIPv6Cidr,omitempty" xml:"DstIPv6Cidr,omitempty"`
	// The destination port range for QoS rule traffic matching. Valid values: **0*	- to **65535**. A value of -1 indicates no match. Currently, only a single port number can be specified, and the start and end port numbers must be the same. The destination port range is fixed for each protocol type. Valid values:
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
	// The DSCP value for QoS rule traffic matching. Valid values: **0*	- to **63**. A value of -1 indicates no match.
	//
	// example:
	//
	// 1
	MatchDscp *int32 `json:"MatchDscp,omitempty" xml:"MatchDscp,omitempty"`
	// The priority of the QoS rule. Valid values: **1*	- to **9000**. A larger value indicates a higher priority. QoS rule priorities must be unique within the same QoS policy.
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
	// example:
	//
	// qos-pksbqfmotl5hzq****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The QoS queue ID.
	//
	// example:
	//
	// qos-queue-9nyx2u7n71s2rc****
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The remarked DSCP value in the traffic. Valid values: **0*	- to **63**. A value of -1 indicates no remarking.
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
	// The QoS rule ID.
	//
	// example:
	//
	// qos-rule-iugg0l9x27f2noc****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the QoS rule.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-rule-test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The source IPv4 CIDR block for QoS rule traffic matching.
	//
	// > You cannot specify this parameter together with **SrcIPv6Cidr*	- or **DstIPv6Cidr**.
	//
	// example:
	//
	// ``1.1.**.**``/24
	SrcCidr *string `json:"SrcCidr,omitempty" xml:"SrcCidr,omitempty"`
	// The source IPv6 CIDR block for QoS rule traffic matching.
	//
	// > You cannot specify this parameter together with **SrcCidr*	- or **DstCidr**.
	//
	// example:
	//
	// 2001:0db8:1234:****::/64
	SrcIPv6Cidr *string `json:"SrcIPv6Cidr,omitempty" xml:"SrcIPv6Cidr,omitempty"`
	// The source port range for QoS rule traffic matching. Valid values: **0*	- to **65535**. A value of -1 indicates no match. Currently, only a single port number can be specified, and the start and end port numbers must be the same.
	//
	// example:
	//
	// -1/-1
	SrcPortRange *string `json:"SrcPortRange,omitempty" xml:"SrcPortRange,omitempty"`
	// The status of the QoS rule. Valid values:
	//
	// - **Normal**: active.
	//
	// - **Configuring**: being configured.
	//
	// - **Deleting**: being deleted.
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
