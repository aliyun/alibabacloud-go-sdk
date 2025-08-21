// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPortRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendPort(v string) *ModifyPortRequest
	GetBackendPort() *string
	SetFrontendPort(v string) *ModifyPortRequest
	GetFrontendPort() *string
	SetFrontendProtocol(v string) *ModifyPortRequest
	GetFrontendProtocol() *string
	SetInstanceId(v string) *ModifyPortRequest
	GetInstanceId() *string
	SetProxyEnable(v int64) *ModifyPortRequest
	GetProxyEnable() *int64
	SetRealServers(v []*string) *ModifyPortRequest
	GetRealServers() []*string
}

type ModifyPortRequest struct {
	// The port of the origin server. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
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

func (s ModifyPortRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPortRequest) GoString() string {
	return s.String()
}

func (s *ModifyPortRequest) GetBackendPort() *string {
	return s.BackendPort
}

func (s *ModifyPortRequest) GetFrontendPort() *string {
	return s.FrontendPort
}

func (s *ModifyPortRequest) GetFrontendProtocol() *string {
	return s.FrontendProtocol
}

func (s *ModifyPortRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPortRequest) GetProxyEnable() *int64 {
	return s.ProxyEnable
}

func (s *ModifyPortRequest) GetRealServers() []*string {
	return s.RealServers
}

func (s *ModifyPortRequest) SetBackendPort(v string) *ModifyPortRequest {
	s.BackendPort = &v
	return s
}

func (s *ModifyPortRequest) SetFrontendPort(v string) *ModifyPortRequest {
	s.FrontendPort = &v
	return s
}

func (s *ModifyPortRequest) SetFrontendProtocol(v string) *ModifyPortRequest {
	s.FrontendProtocol = &v
	return s
}

func (s *ModifyPortRequest) SetInstanceId(v string) *ModifyPortRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPortRequest) SetProxyEnable(v int64) *ModifyPortRequest {
	s.ProxyEnable = &v
	return s
}

func (s *ModifyPortRequest) SetRealServers(v []*string) *ModifyPortRequest {
	s.RealServers = v
	return s
}

func (s *ModifyPortRequest) Validate() error {
	return dara.Validate(s)
}
