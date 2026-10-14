// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigUdpReflectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ConfigUdpReflectRequest
	GetConfig() *string
	SetInstanceId(v string) *ConfigUdpReflectRequest
	GetInstanceId() *string
	SetRegionId(v string) *ConfigUdpReflectRequest
	GetRegionId() *string
}

type ConfigUdpReflectRequest struct {
	// The configuration of the UDP reflection attack attack prevention protection policy to add.
	//
	// This parameter is a string that represents a JSON structure transformed into a character string. The JSON structure contains the following field:
	//
	// - **UdpSports**: Array type. Required. The UDP reflection source ports to block. Example: `[17,19]`.
	//
	//     Block the following source ports that are commonly used in UDP reflection attacks:
	//
	//     - UDP 17: QOTD reflection attacks
	//
	//     - UDP 19: CharGEN reflection attacks
	//
	//     - UDP 69: TFTP reflection attacks
	//
	//     - UDP 111: Portmap reflection attacks
	//
	//     - UDP 123: NTP reflection attacks
	//
	//     - UDP 137: NetBIOS reflection attacks
	//
	//     - UDP 161: SNMPv2 reflection attacks
	//
	//     - UDP 389: CLDAP reflection attacks
	//
	//     - UDP 1194: OpenVPN reflection attacks
	//
	//     - UDP 1900: SSDP reflection attacks
	//
	//     - UDP 3389: RDP reflection attacks
	//
	//     - UDP 11211: Memcached reflection attacks
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"UdpSports\\":[17,19]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance to manage.
	//
	// > You can invoke [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) to query the IDs of all Anti-DDoS Pro or Anti-DDoS Premium instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-i7m25564****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Anti-DDoS Pro or Anti-DDoS Premium instance. Valid values:
	//
	// - **cn-hangzhou*	- (default): the Chinese mainland, which indicates an Anti-DDoS Pro instance.
	//
	// - **ap-southeast-1**: outside the Chinese mainland, which indicates an Anti-DDoS Premium instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigUdpReflectRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigUdpReflectRequest) GoString() string {
	return s.String()
}

func (s *ConfigUdpReflectRequest) GetConfig() *string {
	return s.Config
}

func (s *ConfigUdpReflectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigUdpReflectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigUdpReflectRequest) SetConfig(v string) *ConfigUdpReflectRequest {
	s.Config = &v
	return s
}

func (s *ConfigUdpReflectRequest) SetInstanceId(v string) *ConfigUdpReflectRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigUdpReflectRequest) SetRegionId(v string) *ConfigUdpReflectRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigUdpReflectRequest) Validate() error {
	return dara.Validate(s)
}
