// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeEndpointsResponseBodyData) *DescribeEndpointsResponseBody
	GetData() *DescribeEndpointsResponseBodyData
	SetRequestId(v string) *DescribeEndpointsResponseBody
	GetRequestId() *string
}

type DescribeEndpointsResponseBody struct {
	// The returned data.
	Data *DescribeEndpointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBody) GetData() *DescribeEndpointsResponseBodyData {
	return s.Data
}

func (s *DescribeEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEndpointsResponseBody) SetData(v *DescribeEndpointsResponseBodyData) *DescribeEndpointsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEndpointsResponseBody) SetRequestId(v string) *DescribeEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEndpointsResponseBodyData struct {
	// The list of endpoint details.
	Endpoints []*DescribeEndpointsResponseBodyDataEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud.
	//
	// 	- **PUBLIC**: public network.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
}

func (s DescribeEndpointsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyData) GetEndpoints() []*DescribeEndpointsResponseBodyDataEndpoints {
	return s.Endpoints
}

func (s *DescribeEndpointsResponseBodyData) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeEndpointsResponseBodyData) SetEndpoints(v []*DescribeEndpointsResponseBodyDataEndpoints) *DescribeEndpointsResponseBodyData {
	s.Endpoints = v
	return s
}

func (s *DescribeEndpointsResponseBodyData) SetInstanceNetworkType(v string) *DescribeEndpointsResponseBodyData {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeEndpointsResponseBodyData) Validate() error {
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

type DescribeEndpointsResponseBodyDataEndpoints struct {
	// The compute group ID.
	//
	// example:
	//
	// cc-ad321**-clickhouse
	ComputingGroupId *string `json:"ComputingGroupId,omitempty" xml:"ComputingGroupId,omitempty"`
	// The endpoint of the instance.
	//
	// example:
	//
	// cc-****-clickhouse.clickhouseserver.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The endpoint name.
	//
	// example:
	//
	// cc-*****-clickhouse
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 172.30.XX.XX
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// - VPC: virtual private cloud.
	//
	// - PUBLIC: public network.
	//
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The list of port details.
	Ports []*DescribeEndpointsResponseBodyDataEndpointsPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The status.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// The VPC-connected instance ID.
	//
	// example:
	//
	// vpc-uf61z****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeEndpointsResponseBodyDataEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointsResponseBodyDataEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) GetComputingGroupId() *string {
	return s.ComputingGroupId
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) GetEndpointName() *string {
	return s.EndpointName
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) GetNetType() *string {
	return s.NetType
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) GetPorts() []*DescribeEndpointsResponseBodyDataEndpointsPorts {
	return s.Ports
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) GetStatus() *string {
	return s.Status
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetComputingGroupId(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.ComputingGroupId = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetConnectionString(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.ConnectionString = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetEndpointName(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.EndpointName = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetIPAddress(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.IPAddress = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetNetType(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.NetType = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetPorts(v []*DescribeEndpointsResponseBodyDataEndpointsPorts) *DescribeEndpointsResponseBodyDataEndpoints {
	s.Ports = v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetStatus(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.Status = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetVSwitchId(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetVpcId(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.VpcId = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) SetVpcInstanceId(v string) *DescribeEndpointsResponseBodyDataEndpoints {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpoints) Validate() error {
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

type DescribeEndpointsResponseBodyDataEndpointsPorts struct {
	// The access port. Valid values:
	//
	// - HttpPort: 8123
	//
	// - HttpsPort: 8443
	//
	// - TcpPort: 9000
	//
	// example:
	//
	// 8123
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// - HttpPort: HTTP port.
	//
	// - HttpsPort: HTTPS port.
	//
	// - TcpPort: TCP port.
	//
	// example:
	//
	// HttpPort
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeEndpointsResponseBodyDataEndpointsPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointsResponseBodyDataEndpointsPorts) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyDataEndpointsPorts) GetPort() *int32 {
	return s.Port
}

func (s *DescribeEndpointsResponseBodyDataEndpointsPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeEndpointsResponseBodyDataEndpointsPorts) SetPort(v int32) *DescribeEndpointsResponseBodyDataEndpointsPorts {
	s.Port = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpointsPorts) SetProtocol(v string) *DescribeEndpointsResponseBodyDataEndpointsPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeEndpointsResponseBodyDataEndpointsPorts) Validate() error {
	return dara.Validate(s)
}
