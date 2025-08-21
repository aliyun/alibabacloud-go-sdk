// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendPort(v string) *DeletePortRequest
	GetBackendPort() *string
	SetFrontendPort(v string) *DeletePortRequest
	GetFrontendPort() *string
	SetFrontendProtocol(v string) *DeletePortRequest
	GetFrontendProtocol() *string
	SetInstanceId(v string) *DeletePortRequest
	GetInstanceId() *string
	SetRealServers(v []*string) *DeletePortRequest
	GetRealServers() []*string
}

type DeletePortRequest struct {
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// An array that consists of the IP addresses of origin servers.
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
}

func (s DeletePortRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePortRequest) GoString() string {
	return s.String()
}

func (s *DeletePortRequest) GetBackendPort() *string {
	return s.BackendPort
}

func (s *DeletePortRequest) GetFrontendPort() *string {
	return s.FrontendPort
}

func (s *DeletePortRequest) GetFrontendProtocol() *string {
	return s.FrontendProtocol
}

func (s *DeletePortRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeletePortRequest) GetRealServers() []*string {
	return s.RealServers
}

func (s *DeletePortRequest) SetBackendPort(v string) *DeletePortRequest {
	s.BackendPort = &v
	return s
}

func (s *DeletePortRequest) SetFrontendPort(v string) *DeletePortRequest {
	s.FrontendPort = &v
	return s
}

func (s *DeletePortRequest) SetFrontendProtocol(v string) *DeletePortRequest {
	s.FrontendProtocol = &v
	return s
}

func (s *DeletePortRequest) SetInstanceId(v string) *DeletePortRequest {
	s.InstanceId = &v
	return s
}

func (s *DeletePortRequest) SetRealServers(v []*string) *DeletePortRequest {
	s.RealServers = v
	return s
}

func (s *DeletePortRequest) Validate() error {
	return dara.Validate(s)
}
