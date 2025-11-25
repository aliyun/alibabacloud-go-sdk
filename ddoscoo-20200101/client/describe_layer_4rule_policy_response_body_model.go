// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer4RulePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackendPort(v int32) *DescribeLayer4RulePolicyResponseBody
	GetBackendPort() *int32
	SetBakMode(v string) *DescribeLayer4RulePolicyResponseBody
	GetBakMode() *string
	SetCurrentIndex(v int32) *DescribeLayer4RulePolicyResponseBody
	GetCurrentIndex() *int32
	SetForwardProtocol(v string) *DescribeLayer4RulePolicyResponseBody
	GetForwardProtocol() *string
	SetFrontendPort(v int32) *DescribeLayer4RulePolicyResponseBody
	GetFrontendPort() *int32
	SetInstanceId(v string) *DescribeLayer4RulePolicyResponseBody
	GetInstanceId() *string
	SetPriRealServers(v []*DescribeLayer4RulePolicyResponseBodyPriRealServers) *DescribeLayer4RulePolicyResponseBody
	GetPriRealServers() []*DescribeLayer4RulePolicyResponseBodyPriRealServers
	SetRequestId(v string) *DescribeLayer4RulePolicyResponseBody
	GetRequestId() *string
	SetSecRealServers(v []*DescribeLayer4RulePolicyResponseBodySecRealServers) *DescribeLayer4RulePolicyResponseBody
	GetSecRealServers() []*DescribeLayer4RulePolicyResponseBodySecRealServers
}

type DescribeLayer4RulePolicyResponseBody struct {
	// The port of the origin server.
	//
	// example:
	//
	// 2022
	BackendPort *int32 `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	// The mode that is used to forward service traffic. Valid values:
	//
	// 	- 0: the default mode. In this mode, Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the origin IP address that you specified when you created the port forwarding rule. You can call the [CreateNetworkRules](https://help.aliyun.com/document_detail/157482.html) operation to create a port forwarding rule.
	//
	// 	- 1: the origin redundancy mode. In this mode, Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the IP addresses of the primary or secondary origin servers. You can call the [ConfigLayer4RulePolicy](https://help.aliyun.com/document_detail/312684.html) operation to configure IP addresses.
	//
	// example:
	//
	// 1
	BakMode *string `json:"BakMode,omitempty" xml:"BakMode,omitempty"`
	// The origin server that is used to receive service traffic. Valid values:
	//
	// 	- **1**: the primary origin server, which indicates that Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the IP addresses of the primary origin server.
	//
	// 	- **2**: the secondary origin server, which indicates that Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the IP addresses of the secondary origin server.
	//
	// example:
	//
	// 1
	CurrentIndex *int32 `json:"CurrentIndex,omitempty" xml:"CurrentIndex,omitempty"`
	// The type of the protocol.
	//
	// example:
	//
	// udp
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	// The forwarding port.
	//
	// example:
	//
	// 2020
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddosDip-sg-4hr2b3l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// An array that consists of the information about the primary origin server, including the IP addresses, forwarding protocol, and forwarding port.
	PriRealServers []*DescribeLayer4RulePolicyResponseBodyPriRealServers `json:"PriRealServers,omitempty" xml:"PriRealServers,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6E46CC51-36BE-1100-B14C-DAF8381B8F73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the information about the secondary origin server, including the IP addresses, forwarding protocol, and forwarding port.
	SecRealServers []*DescribeLayer4RulePolicyResponseBodySecRealServers `json:"SecRealServers,omitempty" xml:"SecRealServers,omitempty" type:"Repeated"`
}

func (s DescribeLayer4RulePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RulePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulePolicyResponseBody) GetBackendPort() *int32 {
	return s.BackendPort
}

func (s *DescribeLayer4RulePolicyResponseBody) GetBakMode() *string {
	return s.BakMode
}

func (s *DescribeLayer4RulePolicyResponseBody) GetCurrentIndex() *int32 {
	return s.CurrentIndex
}

func (s *DescribeLayer4RulePolicyResponseBody) GetForwardProtocol() *string {
	return s.ForwardProtocol
}

func (s *DescribeLayer4RulePolicyResponseBody) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeLayer4RulePolicyResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLayer4RulePolicyResponseBody) GetPriRealServers() []*DescribeLayer4RulePolicyResponseBodyPriRealServers {
	return s.PriRealServers
}

func (s *DescribeLayer4RulePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLayer4RulePolicyResponseBody) GetSecRealServers() []*DescribeLayer4RulePolicyResponseBodySecRealServers {
	return s.SecRealServers
}

func (s *DescribeLayer4RulePolicyResponseBody) SetBackendPort(v int32) *DescribeLayer4RulePolicyResponseBody {
	s.BackendPort = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBody) SetBakMode(v string) *DescribeLayer4RulePolicyResponseBody {
	s.BakMode = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBody) SetCurrentIndex(v int32) *DescribeLayer4RulePolicyResponseBody {
	s.CurrentIndex = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBody) SetForwardProtocol(v string) *DescribeLayer4RulePolicyResponseBody {
	s.ForwardProtocol = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBody) SetFrontendPort(v int32) *DescribeLayer4RulePolicyResponseBody {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBody) SetInstanceId(v string) *DescribeLayer4RulePolicyResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBody) SetPriRealServers(v []*DescribeLayer4RulePolicyResponseBodyPriRealServers) *DescribeLayer4RulePolicyResponseBody {
	s.PriRealServers = v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBody) SetRequestId(v string) *DescribeLayer4RulePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBody) SetSecRealServers(v []*DescribeLayer4RulePolicyResponseBodySecRealServers) *DescribeLayer4RulePolicyResponseBody {
	s.SecRealServers = v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBody) Validate() error {
	if s.PriRealServers != nil {
		for _, item := range s.PriRealServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecRealServers != nil {
		for _, item := range s.SecRealServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLayer4RulePolicyResponseBodyPriRealServers struct {
	// The origin server that is used to receive service traffic. Valid values:
	//
	// 	- **1**: the primary origin server, which indicates that Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the IP addresses of the primary origin server.
	//
	// 	- **2**: the secondary origin server, which indicates that Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the IP addresses of the secondary origin server.
	//
	// example:
	//
	// 1
	CurrentIndex *int32 `json:"CurrentIndex,omitempty" xml:"CurrentIndex,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 203.107.XX.XX
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// The forwarding port.
	//
	// example:
	//
	// 2020
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddosDip-sg-4hr2b3l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the protocol.
	//
	// example:
	//
	// udp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The IP address of the primary origin server.
	//
	// example:
	//
	// 192.0.2.1
	RealServer *string `json:"RealServer,omitempty" xml:"RealServer,omitempty"`
}

func (s DescribeLayer4RulePolicyResponseBodyPriRealServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RulePolicyResponseBodyPriRealServers) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) GetCurrentIndex() *int32 {
	return s.CurrentIndex
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) GetEip() *string {
	return s.Eip
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) GetRealServer() *string {
	return s.RealServer
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) SetCurrentIndex(v int32) *DescribeLayer4RulePolicyResponseBodyPriRealServers {
	s.CurrentIndex = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) SetEip(v string) *DescribeLayer4RulePolicyResponseBodyPriRealServers {
	s.Eip = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) SetFrontendPort(v int32) *DescribeLayer4RulePolicyResponseBodyPriRealServers {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) SetInstanceId(v string) *DescribeLayer4RulePolicyResponseBodyPriRealServers {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) SetProtocol(v string) *DescribeLayer4RulePolicyResponseBodyPriRealServers {
	s.Protocol = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) SetRealServer(v string) *DescribeLayer4RulePolicyResponseBodyPriRealServers {
	s.RealServer = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodyPriRealServers) Validate() error {
	return dara.Validate(s)
}

type DescribeLayer4RulePolicyResponseBodySecRealServers struct {
	// The origin server that is used to receive service traffic. Valid values:
	//
	// 	- **1**: the primary origin server, which indicates that Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the IP addresses of the primary origin server.
	//
	// 	- **2**: the secondary origin server, which indicates that Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the IP addresses of the secondary origin server.
	//
	// example:
	//
	// 1
	CurrentIndex *int32 `json:"CurrentIndex,omitempty" xml:"CurrentIndex,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 203.107.XX.XX
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// The forwarding port.
	//
	// example:
	//
	// 2020
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddosDip-sg-4hr2b3l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the protocol.
	//
	// example:
	//
	// udp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The IP address of the secondary origin server.
	//
	// example:
	//
	// 192.0.2.3
	RealServer *string `json:"RealServer,omitempty" xml:"RealServer,omitempty"`
}

func (s DescribeLayer4RulePolicyResponseBodySecRealServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RulePolicyResponseBodySecRealServers) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) GetCurrentIndex() *int32 {
	return s.CurrentIndex
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) GetEip() *string {
	return s.Eip
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) GetRealServer() *string {
	return s.RealServer
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) SetCurrentIndex(v int32) *DescribeLayer4RulePolicyResponseBodySecRealServers {
	s.CurrentIndex = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) SetEip(v string) *DescribeLayer4RulePolicyResponseBodySecRealServers {
	s.Eip = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) SetFrontendPort(v int32) *DescribeLayer4RulePolicyResponseBodySecRealServers {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) SetInstanceId(v string) *DescribeLayer4RulePolicyResponseBodySecRealServers {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) SetProtocol(v string) *DescribeLayer4RulePolicyResponseBodySecRealServers {
	s.Protocol = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) SetRealServer(v string) *DescribeLayer4RulePolicyResponseBodySecRealServers {
	s.RealServer = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponseBodySecRealServers) Validate() error {
	return dara.Validate(s)
}
