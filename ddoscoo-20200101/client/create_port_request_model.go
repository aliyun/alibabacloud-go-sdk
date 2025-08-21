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
	SetProxyEnable(v int64) *CreatePortRequest
	GetProxyEnable() *int64
	SetRealServers(v []*string) *CreatePortRequest
	GetRealServers() []*string
}

type CreatePortRequest struct {
	// The port of the origin server. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 55
	BackendPort *string `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	// The forwarding port. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 55
	FrontendPort *string `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The type of the forwarding protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	FrontendProtocol *string `json:"FrontendProtocol,omitempty" xml:"FrontendProtocol,omitempty"`
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance to which the port forwarding rule belongs.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-st21zbyq****
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProxyEnable *int64  `json:"ProxyEnable,omitempty" xml:"ProxyEnable,omitempty"`
	// An array that consists of the IP addresses of origin servers.
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
