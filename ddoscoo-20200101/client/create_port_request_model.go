// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendPort(v string) *CreatePortRequest
	GetBackendPort() *string
	SetFrontendPort(v string) *CreatePortRequest
	GetFrontendPort() *string
	SetFrontendProtocol(v string) *CreatePortRequest
	GetFrontendProtocol() *string
	SetInstanceId(v string) *CreatePortRequest
	GetInstanceId() *string
	SetIpMode(v string) *CreatePortRequest
	GetIpMode() *string
	SetProxyEnable(v int64) *CreatePortRequest
	GetProxyEnable() *int64
	SetRealServers(v []*string) *CreatePortRequest
	GetRealServers() []*string
}

type CreatePortRequest struct {
	// The origin server port. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 55
	BackendPort *string `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	// The forward port. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 55
	FrontendPort *string `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The forwarding protocol type. Valid values:
	//
	// - **tcp**: TCP protocol.
	//
	// - **udp**: UDP protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	FrontendProtocol *string `json:"FrontendProtocol,omitempty" xml:"FrontendProtocol,omitempty"`
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance to which the port forwarding rule belongs.
	//
	// > You can call [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) to query the IDs of all Anti-DDoS Pro or Anti-DDoS Premium instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-st21zbyq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpMode     *string `json:"IpMode,omitempty" xml:"IpMode,omitempty"`
	// The traffic redirection switch. Valid values:
	//
	// - **0**: disabled.
	//
	// - **1**: enabled.
	//
	// example:
	//
	// 0
	ProxyEnable *int64 `json:"ProxyEnable,omitempty" xml:"ProxyEnable,omitempty"`
	// The list of origin server IP addresses.
	//
	// This parameter is required.
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
}

func (s CreatePortRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePortRequest) GoString() string {
	return s.String()
}

func (s *CreatePortRequest) GetBackendPort() *string {
	return s.BackendPort
}

func (s *CreatePortRequest) GetFrontendPort() *string {
	return s.FrontendPort
}

func (s *CreatePortRequest) GetFrontendProtocol() *string {
	return s.FrontendProtocol
}

func (s *CreatePortRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreatePortRequest) GetIpMode() *string {
	return s.IpMode
}

func (s *CreatePortRequest) GetProxyEnable() *int64 {
	return s.ProxyEnable
}

func (s *CreatePortRequest) GetRealServers() []*string {
	return s.RealServers
}

func (s *CreatePortRequest) SetBackendPort(v string) *CreatePortRequest {
	s.BackendPort = &v
	return s
}

func (s *CreatePortRequest) SetFrontendPort(v string) *CreatePortRequest {
	s.FrontendPort = &v
	return s
}

func (s *CreatePortRequest) SetFrontendProtocol(v string) *CreatePortRequest {
	s.FrontendProtocol = &v
	return s
}

func (s *CreatePortRequest) SetInstanceId(v string) *CreatePortRequest {
	s.InstanceId = &v
	return s
}

func (s *CreatePortRequest) SetIpMode(v string) *CreatePortRequest {
	s.IpMode = &v
	return s
}

func (s *CreatePortRequest) SetProxyEnable(v int64) *CreatePortRequest {
	s.ProxyEnable = &v
	return s
}

func (s *CreatePortRequest) SetRealServers(v []*string) *CreatePortRequest {
	s.RealServers = v
	return s
}

func (s *CreatePortRequest) Validate() error {
	return dara.Validate(s)
}
