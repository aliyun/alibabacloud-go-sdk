// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoints(v []*DescribeResolverEndpointsResponseBodyEndpoints) *DescribeResolverEndpointsResponseBody
	GetEndpoints() []*DescribeResolverEndpointsResponseBodyEndpoints
	SetPageNumber(v int32) *DescribeResolverEndpointsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeResolverEndpointsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeResolverEndpointsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeResolverEndpointsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeResolverEndpointsResponseBody
	GetTotalPages() *int32
}

type DescribeResolverEndpointsResponseBody struct {
	// The endpoints.
	Endpoints []*DescribeResolverEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83D1682B-B69A-4060-9FA8-2907C2A35600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of endpoints.
	//
	// example:
	//
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeResolverEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointsResponseBody) GetEndpoints() []*DescribeResolverEndpointsResponseBodyEndpoints {
	return s.Endpoints
}

func (s *DescribeResolverEndpointsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeResolverEndpointsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeResolverEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResolverEndpointsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeResolverEndpointsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeResolverEndpointsResponseBody) SetEndpoints(v []*DescribeResolverEndpointsResponseBodyEndpoints) *DescribeResolverEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *DescribeResolverEndpointsResponseBody) SetPageNumber(v int32) *DescribeResolverEndpointsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBody) SetPageSize(v int32) *DescribeResolverEndpointsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBody) SetRequestId(v string) *DescribeResolverEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBody) SetTotalItems(v int32) *DescribeResolverEndpointsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBody) SetTotalPages(v int32) *DescribeResolverEndpointsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeResolverEndpointsResponseBodyEndpoints struct {
	// The time when the endpoint was created.
	//
	// example:
	//
	// 2020-07-13 10:36:26
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the endpoint was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594607786000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// hr****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The source IP addresses of outbound traffic.
	IpConfigs []*DescribeResolverEndpointsResponseBodyEndpointsIpConfigs `json:"IpConfigs,omitempty" xml:"IpConfigs,omitempty" type:"Repeated"`
	// The name of the endpoint.
	//
	// example:
	//
	// endpoint-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-0jld3m9yq7l2cw12****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The state of the endpoint that you queried. Valid values:
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
	// 2020-07-13 10:38:24
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the endpoint was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594607904000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The ID of the outbound VPC. All outbound Domain Name System (DNS) requests of the resolver are forwarded by this VPC.
	//
	// example:
	//
	// vpc-0jlxhpfnj5bfu0bsd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the outbound VPC.
	//
	// example:
	//
	// vpc-test-name
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The region ID of the outbound VPC.
	//
	// example:
	//
	// cn-zhangjiakou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	// The name of the region where the VPC resides.
	//
	// example:
	//
	// China East 1 (Hangzhou)
	VpcRegionName *string `json:"VpcRegionName,omitempty" xml:"VpcRegionName,omitempty"`
}

func (s DescribeResolverEndpointsResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetId() *string {
	return s.Id
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetIpConfigs() []*DescribeResolverEndpointsResponseBodyEndpointsIpConfigs {
	return s.IpConfigs
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetName() *string {
	return s.Name
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetStatus() *string {
	return s.Status
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetVpcRegionId() *string {
	return s.VpcRegionId
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) GetVpcRegionName() *string {
	return s.VpcRegionName
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetCreateTime(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.CreateTime = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetCreateTimestamp(v int64) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetId(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.Id = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetIpConfigs(v []*DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.IpConfigs = v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetName(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.Name = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetSecurityGroupId(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetStatus(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.Status = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetUpdateTime(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.UpdateTime = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetUpdateTimestamp(v int64) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetVpcId(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.VpcId = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetVpcName(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.VpcName = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetVpcRegionId(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.VpcRegionId = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetVpcRegionName(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.VpcRegionName = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) Validate() error {
	return dara.Validate(s)
}

type DescribeResolverEndpointsResponseBodyEndpointsIpConfigs struct {
	// The ID of the zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-zhangjiakou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 172.16.XX.XX/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The source IP address of outbound traffic. The IP address must be within the specified CIDR block.
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

func (s DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) GetAzId() *string {
	return s.AzId
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) GetIp() *string {
	return s.Ip
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) SetAzId(v string) *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs {
	s.AzId = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) SetCidrBlock(v string) *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) SetIp(v string) *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs {
	s.Ip = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) SetVSwitchId(v string) *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) Validate() error {
	return dara.Validate(s)
}
