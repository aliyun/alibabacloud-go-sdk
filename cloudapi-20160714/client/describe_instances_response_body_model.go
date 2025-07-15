// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody
	GetInstances() *DescribeInstancesResponseBodyInstances
	SetPageNumber(v int32) *DescribeInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeInstancesResponseBody struct {
	// The information about the instances.
	Instances *DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEB6EC62-B6C7-5082-A45A-45A204724AC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetInstances() *DescribeInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBody) SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstances struct {
	InstanceAttribute []*DescribeInstancesResponseBodyInstancesInstanceAttribute `json:"InstanceAttribute,omitempty" xml:"InstanceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetInstanceAttribute() []*DescribeInstancesResponseBodyInstancesInstanceAttribute {
	return s.InstanceAttribute
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceAttribute(v []*DescribeInstancesResponseBodyInstancesInstanceAttribute) *DescribeInstancesResponseBodyInstances {
	s.InstanceAttribute = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceAttribute struct {
	// The ACL ID.
	//
	// example:
	//
	// acl-uf6f9zfxfxtp5j9ng3yv4
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The name of the access control list (ACL).
	//
	// example:
	//
	// test
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// Indicates whether the ACL is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The ACL type. Valid values:
	//
	// 	- black: blacklist
	//
	// 	- white: whitelist
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The egress IP address.
	//
	// example:
	//
	// 10.0.0.1
	ClassicEgressAddress *string `json:"ClassicEgressAddress,omitempty" xml:"ClassicEgressAddress,omitempty"`
	// The internal CIDR block of the user\\"s VPC that can be accessed if the instance is a VPC integration instance.
	//
	// example:
	//
	// [\\"172.16.0.0/24\\",\\"172.16.1.0/24\\"]
	ConnectCidrBlocks *string `json:"ConnectCidrBlocks,omitempty" xml:"ConnectCidrBlocks,omitempty"`
	// The ID of the user\\"s VPC if the instance is a VPC integration instance.
	//
	// example:
	//
	// vpc-m5eo7khlb4h4f8y9egsdg
	ConnectVpcId *string `json:"ConnectVpcId,omitempty" xml:"ConnectVpcId,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2021-10-22 15:36:53.0
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The type of the dedicated instance. Valid values:
	//
	// 	- vpc_connect: VPC integration instance
	//
	// 	- normal: conventional dedicated instance
	//
	// example:
	//
	// vpc_connect
	DedicatedInstanceType *string `json:"DedicatedInstanceType,omitempty" xml:"DedicatedInstanceType,omitempty"`
	// Indicates whether outbound IPv6 traffic is supported.
	//
	// example:
	//
	// true
	EgressIpv6Enable *bool `json:"EgressIpv6Enable,omitempty" xml:"EgressIpv6Enable,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 1659801600000
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The HTTPS security policy.
	//
	// example:
	//
	// HTTPS2_TLS1_2
	HttpsPolicies *string `json:"HttpsPolicies,omitempty" xml:"HttpsPolicies,omitempty"`
	// The ID of the IPv6 ACL.
	//
	// example:
	//
	// acl-124resFfs235
	IPV6AclId *string `json:"IPV6AclId,omitempty" xml:"IPV6AclId,omitempty"`
	// The name of the IPv6 ACL.
	//
	// example:
	//
	// testIPV6
	IPV6AclName *string `json:"IPV6AclName,omitempty" xml:"IPV6AclName,omitempty"`
	// Indicates whether the IPv6 ACL is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	IPV6AclStatus *string `json:"IPV6AclStatus,omitempty" xml:"IPV6AclStatus,omitempty"`
	// The type of the IPv6 ACL. Valid values:
	//
	// 	- black: blacklist
	//
	// 	- white: whitelist
	//
	// example:
	//
	// black
	IPV6AclType *string `json:"IPV6AclType,omitempty" xml:"IPV6AclType,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription
	//
	// 	- PayAsYouGo: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The CIDR block of the dedicated instance.
	//
	// 	- 172.16.0.0/12
	//
	// 	- 192.168.0.0/16
	//
	// example:
	//
	// 192.168.0.0/16
	InstanceCidrBlock *string `json:"InstanceCidrBlock,omitempty" xml:"InstanceCidrBlock,omitempty"`
	// The ID of the cluster to which the dedicated instance cluster belongs.
	//
	// example:
	//
	// apigateway-cluster-sh-1523cafbgffd
	InstanceClusterId *string `json:"InstanceClusterId,omitempty" xml:"InstanceClusterId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// api-shared-vpc-020
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The requests per second (RPS) limit on the instance.
	//
	// example:
	//
	// 500
	InstanceRpsLimit *int32 `json:"InstanceRpsLimit,omitempty" xml:"InstanceRpsLimit,omitempty"`
	// The instance specification.
	//
	// example:
	//
	// api.s1.small
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The instance specification details.
	InstanceSpecAttributes *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributes `json:"InstanceSpecAttributes,omitempty" xml:"InstanceSpecAttributes,omitempty" type:"Struct"`
	// The instance type. Valid values:
	//
	// 	- VPC_SHARED: shared instance (VPC)
	//
	// 	- VPC_DEDICATED: dedicated instance (VPC)
	//
	// example:
	//
	// VPC_SHARED
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The outbound public IP address.
	//
	// example:
	//
	// 47.241.89.244
	InternetEgressAddress *string `json:"InternetEgressAddress,omitempty" xml:"InternetEgressAddress,omitempty"`
	// The internal CIDR block that is allowed to access the API Gateway instance.
	//
	// example:
	//
	// [\\"172.36.0.0/16\\",\\"172.31.16.0/20\\"]
	IntranetSegments *string `json:"IntranetSegments,omitempty" xml:"IntranetSegments,omitempty"`
	// The end time of the maintenance window. The time is in the *HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 01:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. The time is in the *HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 22:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The network information of the user\\"s VPC if the instance is a VPC integration instance.
	NetworkInterfaceAttributes *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributes `json:"NetworkInterfaceAttributes,omitempty" xml:"NetworkInterfaceAttributes,omitempty" type:"Struct"`
	// The new VPC egress CIDR block.
	//
	// example:
	//
	// 100.104.253.0/26
	NewVpcEgressAddress *string `json:"NewVpcEgressAddress,omitempty" xml:"NewVpcEgressAddress,omitempty"`
	// The private DNS list.
	PrivateDnsList *DescribeInstancesResponseBodyInstancesInstanceAttributePrivateDnsList `json:"PrivateDnsList,omitempty" xml:"PrivateDnsList,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether IPv6 traffic is supported.
	//
	// example:
	//
	// true
	SupportIpv6 *bool `json:"SupportIpv6,omitempty" xml:"SupportIpv6,omitempty"`
	// The tags of the instance.
	Tags *DescribeInstancesResponseBodyInstancesInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The user VPC ID.
	//
	// example:
	//
	// vpc-t***hx****yu9****t0g4
	UserVpcId *string `json:"UserVpcId,omitempty" xml:"UserVpcId,omitempty"`
	// The user vSwitch ID.
	//
	// example:
	//
	// vsw-t4***eh****d7q****i2f
	UserVswitchId *string `json:"UserVswitchId,omitempty" xml:"UserVswitchId,omitempty"`
	// The VPC egress CIDR block.
	//
	// example:
	//
	// 100.104.254.0/26
	VpcEgressAddress *string `json:"VpcEgressAddress,omitempty" xml:"VpcEgressAddress,omitempty"`
	// Indicates whether VPC access is enabled.
	//
	// example:
	//
	// true
	VpcIntranetEnable *bool `json:"VpcIntranetEnable,omitempty" xml:"VpcIntranetEnable,omitempty"`
	// The ID of the account to which the VPC-based instance belongs.
	//
	// example:
	//
	// 1408453217640291****
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
	// Indicates whether virtual private cloud (VPC) Server Load Balancer (SLB) is enabled.
	//
	// example:
	//
	// true
	VpcSlbIntranetEnable *bool `json:"VpcSlbIntranetEnable,omitempty" xml:"VpcSlbIntranetEnable,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-MAZ5(g,h)
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone.
	//
	// example:
	//
	// Multi-Availability Zone 3(b,c,a)
	ZoneLocalName *string `json:"ZoneLocalName,omitempty" xml:"ZoneLocalName,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetAclId() *string {
	return s.AclId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetAclName() *string {
	return s.AclName
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetAclType() *string {
	return s.AclType
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetClassicEgressAddress() *string {
	return s.ClassicEgressAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetConnectCidrBlocks() *string {
	return s.ConnectCidrBlocks
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetConnectVpcId() *string {
	return s.ConnectVpcId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetDedicatedInstanceType() *string {
	return s.DedicatedInstanceType
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetEgressIpv6Enable() *bool {
	return s.EgressIpv6Enable
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetHttpsPolicies() *string {
	return s.HttpsPolicies
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetIPV6AclId() *string {
	return s.IPV6AclId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetIPV6AclName() *string {
	return s.IPV6AclName
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetIPV6AclStatus() *string {
	return s.IPV6AclStatus
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetIPV6AclType() *string {
	return s.IPV6AclType
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetInstanceCidrBlock() *string {
	return s.InstanceCidrBlock
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetInstanceClusterId() *string {
	return s.InstanceClusterId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetInstanceRpsLimit() *int32 {
	return s.InstanceRpsLimit
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetInstanceSpecAttributes() *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributes {
	return s.InstanceSpecAttributes
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetInternetEgressAddress() *string {
	return s.InternetEgressAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetIntranetSegments() *string {
	return s.IntranetSegments
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetNetworkInterfaceAttributes() *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributes {
	return s.NetworkInterfaceAttributes
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetNewVpcEgressAddress() *string {
	return s.NewVpcEgressAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetPrivateDnsList() *DescribeInstancesResponseBodyInstancesInstanceAttributePrivateDnsList {
	return s.PrivateDnsList
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetSupportIpv6() *bool {
	return s.SupportIpv6
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetTags() *DescribeInstancesResponseBodyInstancesInstanceAttributeTags {
	return s.Tags
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetUserVpcId() *string {
	return s.UserVpcId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetUserVswitchId() *string {
	return s.UserVswitchId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetVpcEgressAddress() *string {
	return s.VpcEgressAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetVpcIntranetEnable() *bool {
	return s.VpcIntranetEnable
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetVpcOwnerId() *int64 {
	return s.VpcOwnerId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetVpcSlbIntranetEnable() *bool {
	return s.VpcSlbIntranetEnable
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) GetZoneLocalName() *string {
	return s.ZoneLocalName
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetAclId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.AclId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetAclName(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.AclName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetAclStatus(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.AclStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetAclType(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.AclType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetClassicEgressAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.ClassicEgressAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetConnectCidrBlocks(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.ConnectCidrBlocks = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetConnectVpcId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.ConnectVpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetCreatedTime(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetDedicatedInstanceType(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.DedicatedInstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetEgressIpv6Enable(v bool) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.EgressIpv6Enable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetExpiredTime(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetHttpsPolicies(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.HttpsPolicies = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetIPV6AclId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.IPV6AclId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetIPV6AclName(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.IPV6AclName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetIPV6AclStatus(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.IPV6AclStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetIPV6AclType(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.IPV6AclType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetInstanceChargeType(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetInstanceCidrBlock(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.InstanceCidrBlock = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetInstanceClusterId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.InstanceClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetInstanceId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetInstanceName(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetInstanceRpsLimit(v int32) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.InstanceRpsLimit = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetInstanceSpec(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetInstanceSpecAttributes(v *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributes) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.InstanceSpecAttributes = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetInstanceType(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetInternetEgressAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.InternetEgressAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetIntranetSegments(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.IntranetSegments = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetMaintainEndTime(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetMaintainStartTime(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetNetworkInterfaceAttributes(v *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributes) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.NetworkInterfaceAttributes = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetNewVpcEgressAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.NewVpcEgressAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetPrivateDnsList(v *DescribeInstancesResponseBodyInstancesInstanceAttributePrivateDnsList) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.PrivateDnsList = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetStatus(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetSupportIpv6(v bool) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.SupportIpv6 = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetTags(v *DescribeInstancesResponseBodyInstancesInstanceAttributeTags) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetUserVpcId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.UserVpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetUserVswitchId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.UserVswitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetVpcEgressAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.VpcEgressAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetVpcIntranetEnable(v bool) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.VpcIntranetEnable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetVpcOwnerId(v int64) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.VpcOwnerId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetVpcSlbIntranetEnable(v bool) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.VpcSlbIntranetEnable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) SetZoneLocalName(v string) *DescribeInstancesResponseBodyInstancesInstanceAttribute {
	s.ZoneLocalName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributes struct {
	SpecAttribute []*DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute `json:"SpecAttribute,omitempty" xml:"SpecAttribute,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributes) GetSpecAttribute() []*DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute {
	return s.SpecAttribute
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributes) SetSpecAttribute(v []*DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute) *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributes {
	s.SpecAttribute = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute struct {
	// The variable name.
	//
	// example:
	//
	// SLA
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The variable value.
	//
	// example:
	//
	// 99.95%
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute) SetLocalName(v string) *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute {
	s.LocalName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute) SetValue(v string) *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute {
	s.Value = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributes struct {
	NetworkInterfaceAttribute []*DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute `json:"NetworkInterfaceAttribute,omitempty" xml:"NetworkInterfaceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributes) GetNetworkInterfaceAttribute() []*DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute {
	return s.NetworkInterfaceAttribute
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributes) SetNetworkInterfaceAttribute(v []*DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributes {
	s.NetworkInterfaceAttribute = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute struct {
	// The CIDR block of the vSwitch.
	//
	// example:
	//
	// 192.168.17.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The security group ID. Services in the same security group can access each other.
	//
	// example:
	//
	// sg-2zeehz13zcyj1kfk3o85
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2zeqals6rbj51bhjn8b89
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-shenzhen-d
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) SetCidrBlock(v string) *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute {
	s.CidrBlock = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) SetSecurityGroupId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) SetVswitchId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceAttributePrivateDnsList struct {
	PrivateDns []*string `json:"PrivateDns,omitempty" xml:"PrivateDns,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributePrivateDnsList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributePrivateDnsList) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributePrivateDnsList) GetPrivateDns() []*string {
	return s.PrivateDns
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributePrivateDnsList) SetPrivateDns(v []*string) *DescribeInstancesResponseBodyInstancesInstanceAttributePrivateDnsList {
	s.PrivateDns = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributePrivateDnsList) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceAttributeTags struct {
	TagInfo []*DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeTags) GetTagInfo() []*DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo {
	return s.TagInfo
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeTags) SetTagInfo(v []*DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo) *DescribeInstancesResponseBodyInstancesInstanceAttributeTags {
	s.TagInfo = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeTags) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo struct {
	// The tag key of the instance.
	//
	// example:
	//
	// Cookie
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance.
	//
	// example:
	//
	// 240
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo) SetKey(v string) *DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo {
	s.Key = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo) SetValue(v string) *DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo {
	s.Value = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo) Validate() error {
	return dara.Validate(s)
}
