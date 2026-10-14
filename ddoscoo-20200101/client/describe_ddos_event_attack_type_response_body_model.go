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
	// The attack type information.
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
	if s.AttackTypes != nil {
		for _, item := range s.AttackTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDDosEventAttackTypeResponseBodyAttackTypes struct {
	// The attack type. Valid values:
	//
	// - **QOTD-Reflect-Flood**: QOTD reflection attack
	//
	// - **CharGEN-Reflect-Flood**: CharGEN reflection attack
	//
	// - **DNS-Reflect-Flood**: DNS reflection attack
	//
	// - **TFTP-Reflect-Flood**: TFTP reflection attack
	//
	// - **Portmap-Reflect-Flood**: Portmap reflection attack
	//
	// - **NTP-Reflect-Flood**: NTP reflection attack
	//
	// - **NetBIOS-Reflect-Flood**: NetBIOS reflection attack
	//
	// - **SNMPv2-Reflect-Flood**: SNMPv2 reflection attack
	//
	// - **CLDAP-Reflect-Flood**: CLDAP reflection attack
	//
	// - **Ripv1-Reflect-Flood**: Ripv1 reflection attack
	//
	// - **OpenVPN-Reflect-Flood**: OpenVPN reflection attack
	//
	// - **SSDP-Reflect-Flood**: SSDP reflection attack
	//
	// - **NetAssistant-Reflect-Flood**: NetAssistant reflection attack
	//
	// - **WSDiscovery-Reflect-Flood**: WSDiscovery reflection attack
	//
	// - **Kad-Reflect-Flood**: Kad reflection attack
	//
	// - **mDNS-Reflect-Flood**: mDNS reflection attack
	//
	// - **10001-Reflect-Flood**: 10001 reflection attack
	//
	// - **Memcached-Reflect-Flood**: Memcached reflection attack
	//
	// - **QNP-Reflect-Flood**: QNP reflection attack
	//
	// - **DVR-Reflect-Flood**: DVR reflection attack
	//
	// - **CoAP-Reflect-Flood**: CoAP reflection attack
	//
	// - **ADDP-Reflect-Flood**: ADDP reflection attack
	//
	// - **Tcp-Syn**: TCP SYN flood attacks
	//
	// - **Tcp-Fin**: TCP FIN flood attack
	//
	// - **Tcp-Ack**: TCP ACK Flood attacks
	//
	// - **Tcp-Rst**: TCP RST flood attack
	//
	// - **Tcp-Pushack**: TCP PSH+ACK flood attack
	//
	// - **Tcp-Synack**: TCP SYN+ACK flood attack
	//
	// - **Udp-None**: UDP attack
	//
	// - **Udp-Ssh**: UDP SSH protocol attack
	//
	// - **Udp-Dns**: UDP DNS protocol attack
	//
	// - **Udp-Http**: UDP HTTP protocol attack
	//
	// - **Udp-Https**: UDP HTTPS protocol attack
	//
	// - **Udp-Ntp**: UDP NTP protocol attack
	//
	// - **Udp-Ldap**: UDP LDAP protocol attack
	//
	// - **Udp-Ssdp**: UDP SSDP protocol attack
	//
	// - **Udp-Memcached**: Memcache UDP reflection amplification attack
	//
	// - **Tcp-Other**: TCP other type attack
	//
	// - **Icmp**: ICMP flood attack
	//
	// - **Igmp**: IGMP flood attack
	//
	// - **Ipv6**: IPv6 attack
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
