// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackIpListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribePackIpListRequest
	GetInstanceId() *string
	SetIp(v string) *DescribePackIpListRequest
	GetIp() *string
	SetMemberUid(v string) *DescribePackIpListRequest
	GetMemberUid() *string
	SetPageNo(v int32) *DescribePackIpListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribePackIpListRequest
	GetPageSize() *int32
	SetProductName(v string) *DescribePackIpListRequest
	GetProductName() *string
	SetRegionId(v string) *DescribePackIpListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribePackIpListRequest
	GetResourceGroupId() *string
}

type DescribePackIpListRequest struct {
	// The ID of the Anti-DDoS Origin instance to query.
	//
	// >  You can call the [DescribeInstanceList](https://help.aliyun.com/document_detail/118698.html) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddosbgp-cn-n6w1r7nz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The protected IP address to query.
	//
	// example:
	//
	// 47.98.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the member.
	//
	// example:
	//
	// 170858869679****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the cloud asset to which the protected IP address to query belongs. Valid values:
	//
	// 	- **ECS**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **SLB**: a Classic Load Balancer (CLB) instance, originally called a Server Load Balancer (SLB) instance.
	//
	// 	- **EIP**: an elastic IP address (EIP). An Internet-facing Application Load Balancer (ALB) instance uses an EIP. If the IP address belongs to the Internet-facing ALB instance, set this parameter to EIP.
	//
	// 	- **WAF**: a Web Application Firewall (WAF) instance.
	//
	// example:
	//
	// ECS
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Deprecated
	//
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribePackIpListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePackIpListRequest) GoString() string {
	return s.String()
}

func (s *DescribePackIpListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePackIpListRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribePackIpListRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribePackIpListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribePackIpListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePackIpListRequest) GetProductName() *string {
	return s.ProductName
}

func (s *DescribePackIpListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePackIpListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePackIpListRequest) SetInstanceId(v string) *DescribePackIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackIpListRequest) SetIp(v string) *DescribePackIpListRequest {
	s.Ip = &v
	return s
}

func (s *DescribePackIpListRequest) SetMemberUid(v string) *DescribePackIpListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribePackIpListRequest) SetPageNo(v int32) *DescribePackIpListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePackIpListRequest) SetPageSize(v int32) *DescribePackIpListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackIpListRequest) SetProductName(v string) *DescribePackIpListRequest {
	s.ProductName = &v
	return s
}

func (s *DescribePackIpListRequest) SetRegionId(v string) *DescribePackIpListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePackIpListRequest) SetResourceGroupId(v string) *DescribePackIpListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePackIpListRequest) Validate() error {
	return dara.Validate(s)
}
