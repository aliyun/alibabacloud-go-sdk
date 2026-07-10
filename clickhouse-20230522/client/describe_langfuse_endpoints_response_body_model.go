// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeLangfuseEndpointsResponseBodyData) *DescribeLangfuseEndpointsResponseBody
	GetData() *DescribeLangfuseEndpointsResponseBodyData
	SetRequestId(v string) *DescribeLangfuseEndpointsResponseBody
	GetRequestId() *string
}

type DescribeLangfuseEndpointsResponseBody struct {
	// The returned data.
	Data *DescribeLangfuseEndpointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLangfuseEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseEndpointsResponseBody) GetData() *DescribeLangfuseEndpointsResponseBodyData {
	return s.Data
}

func (s *DescribeLangfuseEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLangfuseEndpointsResponseBody) SetData(v *DescribeLangfuseEndpointsResponseBodyData) *DescribeLangfuseEndpointsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBody) SetRequestId(v string) *DescribeLangfuseEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLangfuseEndpointsResponseBodyData struct {
	// The list of endpoints.
	Endpoints []*DescribeLangfuseEndpointsResponseBodyDataEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC).
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
}

func (s DescribeLangfuseEndpointsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseEndpointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseEndpointsResponseBodyData) GetEndpoints() []*DescribeLangfuseEndpointsResponseBodyDataEndpoints {
	return s.Endpoints
}

func (s *DescribeLangfuseEndpointsResponseBodyData) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeLangfuseEndpointsResponseBodyData) SetEndpoints(v []*DescribeLangfuseEndpointsResponseBodyDataEndpoints) *DescribeLangfuseEndpointsResponseBodyData {
	s.Endpoints = v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBodyData) SetInstanceNetworkType(v string) *DescribeLangfuseEndpointsResponseBodyData {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBodyData) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLangfuseEndpointsResponseBodyDataEndpoints struct {
	// The endpoint of the instance.
	//
	// example:
	//
	// lfs-2zeejcvmzn1******.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The endpoint name.
	//
	// example:
	//
	// lfs-2zeejcvmzn1******
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 172.30.****.****
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// - VPC: VPC.
	//
	// - PUBLIC: Internet.
	//
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port details.
	Ports []*DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-0xi8829****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-uf61z****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeLangfuseEndpointsResponseBodyDataEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseEndpointsResponseBodyDataEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) GetEndpointName() *string {
	return s.EndpointName
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) GetNetType() *string {
	return s.NetType
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) GetPorts() []*DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts {
	return s.Ports
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) SetConnectionString(v string) *DescribeLangfuseEndpointsResponseBodyDataEndpoints {
	s.ConnectionString = &v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) SetEndpointName(v string) *DescribeLangfuseEndpointsResponseBodyDataEndpoints {
	s.EndpointName = &v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) SetIPAddress(v string) *DescribeLangfuseEndpointsResponseBodyDataEndpoints {
	s.IPAddress = &v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) SetNetType(v string) *DescribeLangfuseEndpointsResponseBodyDataEndpoints {
	s.NetType = &v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) SetPorts(v []*DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts) *DescribeLangfuseEndpointsResponseBodyDataEndpoints {
	s.Ports = v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) SetVSwitchId(v string) *DescribeLangfuseEndpointsResponseBodyDataEndpoints {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) SetVpcId(v string) *DescribeLangfuseEndpointsResponseBodyDataEndpoints {
	s.VpcId = &v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpoints) Validate() error {
	if s.Ports != nil {
		for _, item := range s.Ports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts struct {
	// The access port. Valid values:
	//
	// - http: 3000
	//
	// example:
	//
	// 3000
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// - http: HTTP port.
	//
	// example:
	//
	// http
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts) GetPort() *int32 {
	return s.Port
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts) SetPort(v int32) *DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts {
	s.Port = &v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts) SetProtocol(v string) *DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeLangfuseEndpointsResponseBodyDataEndpointsPorts) Validate() error {
	return dara.Validate(s)
}
