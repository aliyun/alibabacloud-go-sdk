// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachFromPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpPortProtocolList(v []*DetachFromPolicyRequestIpPortProtocolList) *DetachFromPolicyRequest
	GetIpPortProtocolList() []*DetachFromPolicyRequestIpPortProtocolList
	SetPolicyType(v string) *DetachFromPolicyRequest
	GetPolicyType() *string
	SetPortVersion(v string) *DetachFromPolicyRequest
	GetPortVersion() *string
}

type DetachFromPolicyRequest struct {
	// The protected objects.
	//
	// This parameter is required.
	IpPortProtocolList []*DetachFromPolicyRequestIpPortProtocolList `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty" type:"Repeated"`
	// The type of the policy. Valid values:
	//
	// 	- **default**: the default mitigation policies.
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// l3
	PolicyType  *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PortVersion *string `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
}

func (s DetachFromPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachFromPolicyRequest) GoString() string {
	return s.String()
}

func (s *DetachFromPolicyRequest) GetIpPortProtocolList() []*DetachFromPolicyRequestIpPortProtocolList {
	return s.IpPortProtocolList
}

func (s *DetachFromPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DetachFromPolicyRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *DetachFromPolicyRequest) SetIpPortProtocolList(v []*DetachFromPolicyRequestIpPortProtocolList) *DetachFromPolicyRequest {
	s.IpPortProtocolList = v
	return s
}

func (s *DetachFromPolicyRequest) SetPolicyType(v string) *DetachFromPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *DetachFromPolicyRequest) SetPortVersion(v string) *DetachFromPolicyRequest {
	s.PortVersion = &v
	return s
}

func (s *DetachFromPolicyRequest) Validate() error {
	if s.IpPortProtocolList != nil {
		for _, item := range s.IpPortProtocolList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetachFromPolicyRequestIpPortProtocolList struct {
	// The IP address of the protected object.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.118.172.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port of the protected object.
	//
	// example:
	//
	// 8*
	Port      *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The protocol type of the protected object. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DetachFromPolicyRequestIpPortProtocolList) String() string {
	return dara.Prettify(s)
}

func (s DetachFromPolicyRequestIpPortProtocolList) GoString() string {
	return s.String()
}

func (s *DetachFromPolicyRequestIpPortProtocolList) GetIp() *string {
	return s.Ip
}

func (s *DetachFromPolicyRequestIpPortProtocolList) GetPort() *int32 {
	return s.Port
}

func (s *DetachFromPolicyRequestIpPortProtocolList) GetPortRange() *string {
	return s.PortRange
}

func (s *DetachFromPolicyRequestIpPortProtocolList) GetProtocol() *string {
	return s.Protocol
}

func (s *DetachFromPolicyRequestIpPortProtocolList) SetIp(v string) *DetachFromPolicyRequestIpPortProtocolList {
	s.Ip = &v
	return s
}

func (s *DetachFromPolicyRequestIpPortProtocolList) SetPort(v int32) *DetachFromPolicyRequestIpPortProtocolList {
	s.Port = &v
	return s
}

func (s *DetachFromPolicyRequestIpPortProtocolList) SetPortRange(v string) *DetachFromPolicyRequestIpPortProtocolList {
	s.PortRange = &v
	return s
}

func (s *DetachFromPolicyRequestIpPortProtocolList) SetProtocol(v string) *DetachFromPolicyRequestIpPortProtocolList {
	s.Protocol = &v
	return s
}

func (s *DetachFromPolicyRequestIpPortProtocolList) Validate() error {
	return dara.Validate(s)
}
