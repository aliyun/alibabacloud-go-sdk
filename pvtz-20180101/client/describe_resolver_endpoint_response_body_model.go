// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeResolverEndpointResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeResolverEndpointResponseBody
	GetCreateTimestamp() *int64
	SetId(v string) *DescribeResolverEndpointResponseBody
	GetId() *string
	SetIpConfigs(v []*DescribeResolverEndpointResponseBodyIpConfigs) *DescribeResolverEndpointResponseBody
	GetIpConfigs() []*DescribeResolverEndpointResponseBodyIpConfigs
	SetName(v string) *DescribeResolverEndpointResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeResolverEndpointResponseBody
	GetRequestId() *string
	SetSecurityGroupId(v string) *DescribeResolverEndpointResponseBody
	GetSecurityGroupId() *string
	SetStatus(v string) *DescribeResolverEndpointResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *DescribeResolverEndpointResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeResolverEndpointResponseBody
	GetUpdateTimestamp() *int64
	SetVpcId(v string) *DescribeResolverEndpointResponseBody
	GetVpcId() *string
	SetVpcName(v string) *DescribeResolverEndpointResponseBody
	GetVpcName() *string
	SetVpcRegionId(v string) *DescribeResolverEndpointResponseBody
	GetVpcRegionId() *string
	SetVpcRegionName(v string) *DescribeResolverEndpointResponseBody
	GetVpcRegionName() *string
}

type DescribeResolverEndpointResponseBody struct {
	// The time when the endpoint was created.
	//
	// example:
	//
	// 2020-07-13 10:45:56
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the endpoint was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594608356000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The endpoint ID. This ID uniquely identifies the endpoint.
	//
	// example:
	//
	// hr****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The configurations of the source IP addresses for outbound traffic.
	IpConfigs []*DescribeResolverEndpointResponseBodyIpConfigs `json:"IpConfigs,omitempty" xml:"IpConfigs,omitempty" type:"Repeated"`
	// The name of the endpoint.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 45020ED9-6319-4CA7-9475-6E8D6446E84F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the security group. The security group rules are applied to the outbound virtual private cloud (VPC).
	//
	// example:
	//
	// sg-8vb3sigz86xc-group-****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The state of the endpoint. Valid values:
	//
	// 	- SUCCESS: The endpoint works as expected.
	//
	// 	- INIT: The endpoint is being created.
	//
	// 	- FAILED: The endpoint failed to be created.
	//
	// 	- CHANGE_INIT: The endpoint is being modified.
	//
	// 	- CHANGE_FAILED: The endpoint failed to be modified.
	//
	// 	- EXCEPTION: The endpoint encountered an exception.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the endpoint was updated.
	//
	// example:
	//
	// 2020-07-13 10:48:39
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the endpoint was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594608519000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The ID of the outbound VPC. All outbound Domain Name System (DNS) requests of the resolver are forwarded by this VPC.
	//
	// example:
	//
	// vpc-0jl96awrjt75ezglc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the outbound VPC.
	//
	// example:
	//
	// vpc-name-test
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The region ID of the outbound VPC.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	// The name of the region where the outbound VPC resides.
	//
	// example:
	//
	// HuaBei
	VpcRegionName *string `json:"VpcRegionName,omitempty" xml:"VpcRegionName,omitempty"`
}

func (s DescribeResolverEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeResolverEndpointResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeResolverEndpointResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeResolverEndpointResponseBody) GetIpConfigs() []*DescribeResolverEndpointResponseBodyIpConfigs {
	return s.IpConfigs
}

func (s *DescribeResolverEndpointResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeResolverEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResolverEndpointResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeResolverEndpointResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeResolverEndpointResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeResolverEndpointResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeResolverEndpointResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeResolverEndpointResponseBody) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeResolverEndpointResponseBody) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *DescribeResolverEndpointResponseBody) GetVpcRegionName() *string {
	return s.VpcRegionName
}

func (s *DescribeResolverEndpointResponseBody) SetCreateTime(v string) *DescribeResolverEndpointResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetCreateTimestamp(v int64) *DescribeResolverEndpointResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetId(v string) *DescribeResolverEndpointResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetIpConfigs(v []*DescribeResolverEndpointResponseBodyIpConfigs) *DescribeResolverEndpointResponseBody {
	s.IpConfigs = v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetName(v string) *DescribeResolverEndpointResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetRequestId(v string) *DescribeResolverEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetSecurityGroupId(v string) *DescribeResolverEndpointResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetStatus(v string) *DescribeResolverEndpointResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetUpdateTime(v string) *DescribeResolverEndpointResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetUpdateTimestamp(v int64) *DescribeResolverEndpointResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetVpcId(v string) *DescribeResolverEndpointResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetVpcName(v string) *DescribeResolverEndpointResponseBody {
	s.VpcName = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetVpcRegionId(v string) *DescribeResolverEndpointResponseBody {
	s.VpcRegionId = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetVpcRegionName(v string) *DescribeResolverEndpointResponseBody {
	s.VpcRegionName = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) Validate() error {
	if s.IpConfigs != nil {
		for _, item := range s.IpConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResolverEndpointResponseBodyIpConfigs struct {
	// The ID of the zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-hangzhou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 172.16.XX.XX/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The source IP address of outbound traffic. The IP address must be within the specified CIDR block. If this parameter is left empty, the system automatically allocates an IP address.
	//
	// example:
	//
	// 172.16.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-0jlgeyq4oazkh5xue****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeResolverEndpointResponseBodyIpConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverEndpointResponseBodyIpConfigs) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) GetAzId() *string {
	return s.AzId
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) GetIp() *string {
	return s.Ip
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) SetAzId(v string) *DescribeResolverEndpointResponseBodyIpConfigs {
	s.AzId = &v
	return s
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) SetCidrBlock(v string) *DescribeResolverEndpointResponseBodyIpConfigs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) SetIp(v string) *DescribeResolverEndpointResponseBodyIpConfigs {
	s.Ip = &v
	return s
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) SetVSwitchId(v string) *DescribeResolverEndpointResponseBodyIpConfigs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) Validate() error {
	return dara.Validate(s)
}
