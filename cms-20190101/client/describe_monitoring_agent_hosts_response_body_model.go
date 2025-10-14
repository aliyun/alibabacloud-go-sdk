// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentHostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMonitoringAgentHostsResponseBody
	GetCode() *string
	SetHosts(v *DescribeMonitoringAgentHostsResponseBodyHosts) *DescribeMonitoringAgentHostsResponseBody
	GetHosts() *DescribeMonitoringAgentHostsResponseBodyHosts
	SetMessage(v string) *DescribeMonitoringAgentHostsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeMonitoringAgentHostsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMonitoringAgentHostsResponseBody
	GetPageSize() *int32
	SetPageTotal(v int32) *DescribeMonitoringAgentHostsResponseBody
	GetPageTotal() *int32
	SetRequestId(v string) *DescribeMonitoringAgentHostsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMonitoringAgentHostsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeMonitoringAgentHostsResponseBody
	GetTotal() *int32
}

type DescribeMonitoringAgentHostsResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the hosts.
	Hosts *DescribeMonitoringAgentHostsResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 50
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 63EEBB2A-9E51-41E4-9E83-5DE7F3B292E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMonitoringAgentHostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentHostsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentHostsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMonitoringAgentHostsResponseBody) GetHosts() *DescribeMonitoringAgentHostsResponseBodyHosts {
	return s.Hosts
}

func (s *DescribeMonitoringAgentHostsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitoringAgentHostsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMonitoringAgentHostsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMonitoringAgentHostsResponseBody) GetPageTotal() *int32 {
	return s.PageTotal
}

func (s *DescribeMonitoringAgentHostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitoringAgentHostsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitoringAgentHostsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetCode(v string) *DescribeMonitoringAgentHostsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetHosts(v *DescribeMonitoringAgentHostsResponseBodyHosts) *DescribeMonitoringAgentHostsResponseBody {
	s.Hosts = v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetMessage(v string) *DescribeMonitoringAgentHostsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetPageNumber(v int32) *DescribeMonitoringAgentHostsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetPageSize(v int32) *DescribeMonitoringAgentHostsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetPageTotal(v int32) *DescribeMonitoringAgentHostsResponseBody {
	s.PageTotal = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetRequestId(v string) *DescribeMonitoringAgentHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetSuccess(v bool) *DescribeMonitoringAgentHostsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetTotal(v int32) *DescribeMonitoringAgentHostsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) Validate() error {
	if s.Hosts != nil {
		if err := s.Hosts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitoringAgentHostsResponseBodyHosts struct {
	Host []*DescribeMonitoringAgentHostsResponseBodyHostsHost `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
}

func (s DescribeMonitoringAgentHostsResponseBodyHosts) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentHostsResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentHostsResponseBodyHosts) GetHost() []*DescribeMonitoringAgentHostsResponseBodyHostsHost {
	return s.Host
}

func (s *DescribeMonitoringAgentHostsResponseBodyHosts) SetHost(v []*DescribeMonitoringAgentHostsResponseBodyHostsHost) *DescribeMonitoringAgentHostsResponseBodyHosts {
	s.Host = v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHosts) Validate() error {
	if s.Host != nil {
		for _, item := range s.Host {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitoringAgentHostsResponseBodyHostsHost struct {
	// The version of the CloudMonitor agent.
	//
	// example:
	//
	// 3.4.6
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 103201326074****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The elastic IP address (EIP) of the host.
	//
	// example:
	//
	// 192.168.XX.XX
	EipAddress *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	// The ID of the EIP.
	//
	// example:
	//
	// eip-bp16i16k9gcezyfrp****
	EipId *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	// The name of the host.
	//
	// example:
	//
	// hostIP
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-a2d5q7pm3f9yr212****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the ECS instance.
	//
	// example:
	//
	// ecs.n4
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The IP address of the host.
	//
	// > Multiple IP addresses are separated with commas (,).
	//
	// example:
	//
	// 192.168.XX.XX
	IpGroup *string `json:"IpGroup,omitempty" xml:"IpGroup,omitempty"`
	// The IP address of the Network Address Translation (NAT) gateway.
	//
	// example:
	//
	// 192.168.XX.XX
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The network type.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The operating system.
	//
	// example:
	//
	// Linux
	OperatingSystem *string `json:"OperatingSystem,omitempty" xml:"OperatingSystem,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The serial number of the host. A host that is not provided by Alibaba Cloud has a serial number instead of an instance ID.
	//
	// > This parameter can be used to accurately search for a monitored host.
	//
	// example:
	//
	// x12335-6cc8-4a22-9f21-1a00a719****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// Indicates whether the host is provided by Alibaba Cloud. Valid values:
	//
	// 	- true: The host is provided by Alibaba Cloud.
	//
	// 	- false: The host is not provided by Alibaba Cloud.
	//
	// example:
	//
	// true
	IsAliyunHost *bool `json:"isAliyunHost,omitempty" xml:"isAliyunHost,omitempty"`
}

func (s DescribeMonitoringAgentHostsResponseBodyHostsHost) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentHostsResponseBodyHostsHost) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetEipAddress() *string {
	return s.EipAddress
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetEipId() *string {
	return s.EipId
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetHostName() *string {
	return s.HostName
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetIpGroup() *string {
	return s.IpGroup
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetNatIp() *string {
	return s.NatIp
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetOperatingSystem() *string {
	return s.OperatingSystem
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetRegion() *string {
	return s.Region
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) GetIsAliyunHost() *bool {
	return s.IsAliyunHost
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetAgentVersion(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.AgentVersion = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetAliUid(v int64) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.AliUid = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetEipAddress(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.EipAddress = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetEipId(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.EipId = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetHostName(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.HostName = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetInstanceId(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetInstanceTypeFamily(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetIpGroup(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.IpGroup = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetNatIp(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.NatIp = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetNetworkType(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.NetworkType = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetOperatingSystem(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.OperatingSystem = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetRegion(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.Region = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetSerialNumber(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.SerialNumber = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetIsAliyunHost(v bool) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.IsAliyunHost = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) Validate() error {
	return dara.Validate(s)
}
