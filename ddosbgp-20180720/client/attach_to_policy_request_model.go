// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachToPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpPortProtocolList(v []*AttachToPolicyRequestIpPortProtocolList) *AttachToPolicyRequest
	GetIpPortProtocolList() []*AttachToPolicyRequestIpPortProtocolList
	SetPolicyId(v string) *AttachToPolicyRequest
	GetPolicyId() *string
	SetPortVersion(v string) *AttachToPolicyRequest
	GetPortVersion() *string
}

type AttachToPolicyRequest struct {
	// The protected objects.
	//
	// This parameter is required.
	IpPortProtocolList []*AttachToPolicyRequestIpPortProtocolList `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty" type:"Repeated"`
	// The policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cd8b4d70-e4e0-413a-b390-e71d********
	PolicyId    *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PortVersion *string `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
}

func (s AttachToPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachToPolicyRequest) GoString() string {
	return s.String()
}

func (s *AttachToPolicyRequest) GetIpPortProtocolList() []*AttachToPolicyRequestIpPortProtocolList {
	return s.IpPortProtocolList
}

func (s *AttachToPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *AttachToPolicyRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *AttachToPolicyRequest) SetIpPortProtocolList(v []*AttachToPolicyRequestIpPortProtocolList) *AttachToPolicyRequest {
	s.IpPortProtocolList = v
	return s
}

func (s *AttachToPolicyRequest) SetPolicyId(v string) *AttachToPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *AttachToPolicyRequest) SetPortVersion(v string) *AttachToPolicyRequest {
	s.PortVersion = &v
	return s
}

func (s *AttachToPolicyRequest) Validate() error {
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

type AttachToPolicyRequestIpPortProtocolList struct {
	// The IP address of the protected object.
	//
	// This parameter is required.
	//
	// example:
	//
	// 112.124.241.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port number of the protected object.
	//
	// >  This parameter is available for only port-specific mitigation policies.
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
	// >  This parameter is available for only port-specific mitigation policies.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s AttachToPolicyRequestIpPortProtocolList) String() string {
	return dara.Prettify(s)
}

func (s AttachToPolicyRequestIpPortProtocolList) GoString() string {
	return s.String()
}

func (s *AttachToPolicyRequestIpPortProtocolList) GetIp() *string {
	return s.Ip
}

func (s *AttachToPolicyRequestIpPortProtocolList) GetPort() *int32 {
	return s.Port
}

func (s *AttachToPolicyRequestIpPortProtocolList) GetPortRange() *string {
	return s.PortRange
}

func (s *AttachToPolicyRequestIpPortProtocolList) GetProtocol() *string {
	return s.Protocol
}

func (s *AttachToPolicyRequestIpPortProtocolList) SetIp(v string) *AttachToPolicyRequestIpPortProtocolList {
	s.Ip = &v
	return s
}

func (s *AttachToPolicyRequestIpPortProtocolList) SetPort(v int32) *AttachToPolicyRequestIpPortProtocolList {
	s.Port = &v
	return s
}

func (s *AttachToPolicyRequestIpPortProtocolList) SetPortRange(v string) *AttachToPolicyRequestIpPortProtocolList {
	s.PortRange = &v
	return s
}

func (s *AttachToPolicyRequestIpPortProtocolList) SetProtocol(v string) *AttachToPolicyRequestIpPortProtocolList {
	s.Protocol = &v
	return s
}

func (s *AttachToPolicyRequestIpPortProtocolList) Validate() error {
	return dara.Validate(s)
}
