// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventAttackTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttackTypes(v []*DescribeDDosEventAttackTypeResponseBodyAttackTypes) *DescribeDDosEventAttackTypeResponseBody
	GetAttackTypes() []*DescribeDDosEventAttackTypeResponseBodyAttackTypes
	SetRequestId(v string) *DescribeDDosEventAttackTypeResponseBody
	GetRequestId() *string
}

type DescribeDDosEventAttackTypeResponseBody struct {
	// The information about the attack types.
	AttackTypes []*DescribeDDosEventAttackTypeResponseBodyAttackTypes `json:"AttackTypes,omitempty" xml:"AttackTypes,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6F644A6E-40E7-483F-9DBB-CC27E16BB555
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDDosEventAttackTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventAttackTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAttackTypeResponseBody) GetAttackTypes() []*DescribeDDosEventAttackTypeResponseBodyAttackTypes {
	return s.AttackTypes
}

func (s *DescribeDDosEventAttackTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDosEventAttackTypeResponseBody) SetAttackTypes(v []*DescribeDDosEventAttackTypeResponseBodyAttackTypes) *DescribeDDosEventAttackTypeResponseBody {
	s.AttackTypes = v
	return s
}

func (s *DescribeDDosEventAttackTypeResponseBody) SetRequestId(v string) *DescribeDDosEventAttackTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosEventAttackTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDDosEventAttackTypeResponseBodyAttackTypes struct {
	// The type of the attack Valid values:
	//
	// 	- **QOTD-Reflect-Flood**: QOTD reflection attacks
	//
	// 	- **CharGEN-Reflect-Flood**: CHARGEN reflection attacks
	//
	// 	- **DNS-Reflect-Flood**: DNS reflection attacks
	//
	// 	- **TFTP-Reflect-Flood**: TFTP reflection attacks
	//
	// 	- **Portmap-Reflect-Flood**: Portmap reflection attacks
	//
	// 	- **NTP-Reflect-Flood**: NTP reflection attacks
	//
	// 	- **NetBIOS-Reflect-Flood**: NetBIOS reflection attacks
	//
	// 	- **SNMPv2-Reflect-Flood**: SNMPv2 reflection attacks
	//
	// 	- **CLDAP-Reflect-Flood**: CLDAP reflection attacks
	//
	// 	- **Ripv1-Reflect-Flood**: RIPv1 reflection attacks
	//
	// 	- **OpenVPN-Reflect-Flood**: OpenVPN reflection attacks
	//
	// 	- **SSDP-Reflect-Flood**: SSDP reflection attacks
	//
	// 	- **NetAssistant-Reflect-Flood**: NetAssistant reflection attacks
	//
	// 	- **WSDiscovery-Reflect-Flood**: WS-Discovery reflection attacks
	//
	// 	- **Kad-Reflect-Flood**: Kad reflection attacks
	//
	// 	- **mDNS-Reflect-Flood**: mDNS reflection attacks
	//
	// 	- **10001-Reflect-Flood**: reflection attacks over port 10001
	//
	// 	- **Memcached-Reflect-Flood**: Memcached reflection attacks
	//
	// 	- **QNP-Reflect-Flood**: QNP reflection attacks
	//
	// 	- **DVR-Reflect-Flood**: DVR reflection attacks
	//
	// 	- **CoAP-Reflect-Flood**: CoAP reflection attacks
	//
	// 	- **ADDP-Reflect-Flood**: ADDP reflection attacks
	//
	// 	- **Tcp-Syn**: TCP SYN flood attacks
	//
	// 	- **Tcp-Fin**: TCP FIN flood attacks
	//
	// 	- **Tcp-Ack**: TCP ACK flood attacks
	//
	// 	- **Tcp-Rst**: TCP RST flood attacks
	//
	// 	- **Tcp-Pushack**: TCP PSH-ACK flood attacks
	//
	// 	- **Tcp-Synack**: TCP SYN-ACK flood attacks
	//
	// 	- **Udp-None**: UDP attacks
	//
	// 	- **Udp-Ssh**: UDP-based SSH attacks
	//
	// 	- **Udp-Dns**: UDP-based DNS attacks
	//
	// 	- **Udp-Http**: UDP-based HTTP attacks
	//
	// 	- **Udp-Https**: UDP-based HTTPS attacks
	//
	// 	- **Udp-Ntp**: UDP-based NTP attacks
	//
	// 	- **Udp-Ldap**: UDP-based LDAP attacks
	//
	// 	- **Udp-Ssdp**: UDP-based SSDP attacks
	//
	// 	- **Udp-Memcached**: Memcached UDP reflection attacks
	//
	// 	- **Tcp-Other**: other TCP attacks
	//
	// 	- **Icmp**: ICMP flood attacks
	//
	// 	- **Igmp**: IGMP flood attacks
	//
	// 	- **Ipv6**: IPv6 attacks
	//
	// example:
	//
	// Tcp-Syn
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The number of request packets of the attack type.
	//
	// example:
	//
	// 145902
	InPkts *int64 `json:"InPkts,omitempty" xml:"InPkts,omitempty"`
}

func (s DescribeDDosEventAttackTypeResponseBodyAttackTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventAttackTypeResponseBodyAttackTypes) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAttackTypeResponseBodyAttackTypes) GetAttackType() *string {
	return s.AttackType
}

func (s *DescribeDDosEventAttackTypeResponseBodyAttackTypes) GetInPkts() *int64 {
	return s.InPkts
}

func (s *DescribeDDosEventAttackTypeResponseBodyAttackTypes) SetAttackType(v string) *DescribeDDosEventAttackTypeResponseBodyAttackTypes {
	s.AttackType = &v
	return s
}

func (s *DescribeDDosEventAttackTypeResponseBodyAttackTypes) SetInPkts(v int64) *DescribeDDosEventAttackTypeResponseBodyAttackTypes {
	s.InPkts = &v
	return s
}

func (s *DescribeDDosEventAttackTypeResponseBodyAttackTypes) Validate() error {
	return dara.Validate(s)
}
