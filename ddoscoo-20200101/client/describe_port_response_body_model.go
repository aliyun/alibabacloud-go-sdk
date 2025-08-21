// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkRules(v []*DescribePortResponseBodyNetworkRules) *DescribePortResponseBody
	GetNetworkRules() []*DescribePortResponseBodyNetworkRules
	SetRequestId(v string) *DescribePortResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePortResponseBody
	GetTotalCount() *int64
}

type DescribePortResponseBody struct {
	// An array that consists of port forwarding rules.
	NetworkRules []*DescribePortResponseBodyNetworkRules `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 39499F01-19D9-4EA4-A0E9-C6014BA5CDBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of port forwarding rules returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePortResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortResponseBody) GetNetworkRules() []*DescribePortResponseBodyNetworkRules {
	return s.NetworkRules
}

func (s *DescribePortResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePortResponseBody) SetNetworkRules(v []*DescribePortResponseBodyNetworkRules) *DescribePortResponseBody {
	s.NetworkRules = v
	return s
}

func (s *DescribePortResponseBody) SetRequestId(v string) *DescribePortResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortResponseBody) SetTotalCount(v int64) *DescribePortResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePortResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePortResponseBodyNetworkRules struct {
	// The port of the origin server.
	//
	// example:
	//
	// 55
	BackendPort *int32 `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	// The forwarding port.
	//
	// example:
	//
	// 55
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// tcp
	FrontendProtocol *string `json:"FrontendProtocol,omitempty" xml:"FrontendProtocol,omitempty"`
	// The ID of the instance to which the port forwarding rule is applied.
	//
	// example:
	//
	// ddoscoo-cn-7e225i41****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the port forwarding rule is automatically created by the instance. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	IsAutoCreate *bool `json:"IsAutoCreate,omitempty" xml:"IsAutoCreate,omitempty"`
	// An array that consists of IP addresses of origin servers.
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
}

func (s DescribePortResponseBodyNetworkRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePortResponseBodyNetworkRules) GoString() string {
	return s.String()
}

func (s *DescribePortResponseBodyNetworkRules) GetBackendPort() *int32 {
	return s.BackendPort
}

func (s *DescribePortResponseBodyNetworkRules) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribePortResponseBodyNetworkRules) GetFrontendProtocol() *string {
	return s.FrontendProtocol
}

func (s *DescribePortResponseBodyNetworkRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePortResponseBodyNetworkRules) GetIsAutoCreate() *bool {
	return s.IsAutoCreate
}

func (s *DescribePortResponseBodyNetworkRules) GetRealServers() []*string {
	return s.RealServers
}

func (s *DescribePortResponseBodyNetworkRules) SetBackendPort(v int32) *DescribePortResponseBodyNetworkRules {
	s.BackendPort = &v
	return s
}

func (s *DescribePortResponseBodyNetworkRules) SetFrontendPort(v int32) *DescribePortResponseBodyNetworkRules {
	s.FrontendPort = &v
	return s
}

func (s *DescribePortResponseBodyNetworkRules) SetFrontendProtocol(v string) *DescribePortResponseBodyNetworkRules {
	s.FrontendProtocol = &v
	return s
}

func (s *DescribePortResponseBodyNetworkRules) SetInstanceId(v string) *DescribePortResponseBodyNetworkRules {
	s.InstanceId = &v
	return s
}

func (s *DescribePortResponseBodyNetworkRules) SetIsAutoCreate(v bool) *DescribePortResponseBodyNetworkRules {
	s.IsAutoCreate = &v
	return s
}

func (s *DescribePortResponseBodyNetworkRules) SetRealServers(v []*string) *DescribePortResponseBodyNetworkRules {
	s.RealServers = v
	return s
}

func (s *DescribePortResponseBodyNetworkRules) Validate() error {
	return dara.Validate(s)
}
