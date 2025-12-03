// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceClusterInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *DescribeInstanceClusterInfoResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *DescribeInstanceClusterInfoResponseBody
	GetDescription() *string
	SetInstanceClusterAttribute(v *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) *DescribeInstanceClusterInfoResponseBody
	GetInstanceClusterAttribute() *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute
	SetInstanceClusterId(v string) *DescribeInstanceClusterInfoResponseBody
	GetInstanceClusterId() *string
	SetInstanceClusterName(v string) *DescribeInstanceClusterInfoResponseBody
	GetInstanceClusterName() *string
	SetInstanceClusterStatus(v string) *DescribeInstanceClusterInfoResponseBody
	GetInstanceClusterStatus() *string
	SetInstanceClusterType(v string) *DescribeInstanceClusterInfoResponseBody
	GetInstanceClusterType() *string
	SetInstanceClusterVersion(v string) *DescribeInstanceClusterInfoResponseBody
	GetInstanceClusterVersion() *string
	SetInstanceList(v *DescribeInstanceClusterInfoResponseBodyInstanceList) *DescribeInstanceClusterInfoResponseBody
	GetInstanceList() *DescribeInstanceClusterInfoResponseBodyInstanceList
	SetModifiedTime(v string) *DescribeInstanceClusterInfoResponseBody
	GetModifiedTime() *string
	SetRegionId(v string) *DescribeInstanceClusterInfoResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeInstanceClusterInfoResponseBody
	GetRequestId() *string
}

type DescribeInstanceClusterInfoResponseBody struct {
	// The time when the cluster was created.
	//
	// example:
	//
	// 2022-10-10T18:29:27
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cluster description, which can be up to 200 characters in length.
	//
	// example:
	//
	// v0.0.4
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The cluster details.
	InstanceClusterAttribute *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute `json:"InstanceClusterAttribute,omitempty" xml:"InstanceClusterAttribute,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// apigateway-ht-04e41d95e9c1
	InstanceClusterId *string `json:"InstanceClusterId,omitempty" xml:"InstanceClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test
	InstanceClusterName *string `json:"InstanceClusterName,omitempty" xml:"InstanceClusterName,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// RUNNING
	InstanceClusterStatus *string `json:"InstanceClusterStatus,omitempty" xml:"InstanceClusterStatus,omitempty"`
	// The cluster type.
	//
	// example:
	//
	// normal
	InstanceClusterType *string `json:"InstanceClusterType,omitempty" xml:"InstanceClusterType,omitempty"`
	// The cluster version.
	//
	// example:
	//
	// 3.5.3.856
	InstanceClusterVersion *string `json:"InstanceClusterVersion,omitempty" xml:"InstanceClusterVersion,omitempty"`
	// The dedicated instances contained in the cluster.
	InstanceList *DescribeInstanceClusterInfoResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	// The time when the cluster was last modified.
	//
	// example:
	//
	// 2023-06-19 10:40:29 +0800
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceClusterInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceClusterInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceClusterInfoResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeInstanceClusterInfoResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceClusterInfoResponseBody) GetInstanceClusterAttribute() *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	return s.InstanceClusterAttribute
}

func (s *DescribeInstanceClusterInfoResponseBody) GetInstanceClusterId() *string {
	return s.InstanceClusterId
}

func (s *DescribeInstanceClusterInfoResponseBody) GetInstanceClusterName() *string {
	return s.InstanceClusterName
}

func (s *DescribeInstanceClusterInfoResponseBody) GetInstanceClusterStatus() *string {
	return s.InstanceClusterStatus
}

func (s *DescribeInstanceClusterInfoResponseBody) GetInstanceClusterType() *string {
	return s.InstanceClusterType
}

func (s *DescribeInstanceClusterInfoResponseBody) GetInstanceClusterVersion() *string {
	return s.InstanceClusterVersion
}

func (s *DescribeInstanceClusterInfoResponseBody) GetInstanceList() *DescribeInstanceClusterInfoResponseBodyInstanceList {
	return s.InstanceList
}

func (s *DescribeInstanceClusterInfoResponseBody) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeInstanceClusterInfoResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceClusterInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceClusterInfoResponseBody) SetCreatedTime(v string) *DescribeInstanceClusterInfoResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) SetDescription(v string) *DescribeInstanceClusterInfoResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) SetInstanceClusterAttribute(v *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) *DescribeInstanceClusterInfoResponseBody {
	s.InstanceClusterAttribute = v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) SetInstanceClusterId(v string) *DescribeInstanceClusterInfoResponseBody {
	s.InstanceClusterId = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) SetInstanceClusterName(v string) *DescribeInstanceClusterInfoResponseBody {
	s.InstanceClusterName = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) SetInstanceClusterStatus(v string) *DescribeInstanceClusterInfoResponseBody {
	s.InstanceClusterStatus = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) SetInstanceClusterType(v string) *DescribeInstanceClusterInfoResponseBody {
	s.InstanceClusterType = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) SetInstanceClusterVersion(v string) *DescribeInstanceClusterInfoResponseBody {
	s.InstanceClusterVersion = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) SetInstanceList(v *DescribeInstanceClusterInfoResponseBodyInstanceList) *DescribeInstanceClusterInfoResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) SetModifiedTime(v string) *DescribeInstanceClusterInfoResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) SetRegionId(v string) *DescribeInstanceClusterInfoResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) SetRequestId(v string) *DescribeInstanceClusterInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBody) Validate() error {
	if s.InstanceClusterAttribute != nil {
		if err := s.InstanceClusterAttribute.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceList != nil {
		if err := s.InstanceList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute struct {
	// The internal CIDR block of the user VPC that can be accessed by the cluster if the cluster consists of VPC integration instances.
	//
	// example:
	//
	// ["192.168.1.0/24","192.168.0.0/24"]
	ConnectCidrBlocks *string `json:"ConnectCidrBlocks,omitempty" xml:"ConnectCidrBlocks,omitempty"`
	// The ID of the user VPC that is connected to the cluster if the cluster consists of VPC integration instances.
	//
	// example:
	//
	// vpc-p0w3kxxxxxxxxxxxxxxxx
	ConnectVpcId *string `json:"ConnectVpcId,omitempty" xml:"ConnectVpcId,omitempty"`
	// Indicates whether outbound IPv6 traffic is supported.
	//
	// example:
	//
	// true
	EgressIpv6Enable *bool `json:"EgressIpv6Enable,omitempty" xml:"EgressIpv6Enable,omitempty"`
	// The HTTPS security policy.
	//
	// example:
	//
	// HTTPS2_TLS1_0
	HttpsPolicies *string `json:"HttpsPolicies,omitempty" xml:"HttpsPolicies,omitempty"`
	// The ID of the IPv4 access control list (ACL).
	//
	// example:
	//
	// acl-t4n8i4rvvp70kcyuoXXXX
	IPV4AclId *string `json:"IPV4AclId,omitempty" xml:"IPV4AclId,omitempty"`
	// The name of the IPv4 ACL.
	//
	// example:
	//
	// test-black
	IPV4AclName *string `json:"IPV4AclName,omitempty" xml:"IPV4AclName,omitempty"`
	// Indicates whether IPv4 access control is enabled. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	IPV4AclStatus *string `json:"IPV4AclStatus,omitempty" xml:"IPV4AclStatus,omitempty"`
	// The type of the IPv4 ACL.
	//
	// 	- black: blacklist
	//
	// 	- white: whitelist
	//
	// example:
	//
	// black
	IPV4AclType *string `json:"IPV4AclType,omitempty" xml:"IPV4AclType,omitempty"`
	// The ID of the IPv6 ACL.
	//
	// example:
	//
	// acl-t4nevzhwbpe7cup18XXXX
	IPV6AclId *string `json:"IPV6AclId,omitempty" xml:"IPV6AclId,omitempty"`
	// The name of the IPv6 ACL.
	//
	// example:
	//
	// test
	IPV6AclName *string `json:"IPV6AclName,omitempty" xml:"IPV6AclName,omitempty"`
	// Indicates whether IPv6 access control is enabled. Valid values:
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
	// white
	IPV6AclType *string `json:"IPV6AclType,omitempty" xml:"IPV6AclType,omitempty"`
	// The outbound public IP address.
	//
	// example:
	//
	// 39.106.XX.XX
	InternetEgressAddress *string `json:"InternetEgressAddress,omitempty" xml:"InternetEgressAddress,omitempty"`
	// The outbound private IP address.
	//
	// example:
	//
	// 100.104.XX.XX/26
	IntranetEgressAddress *string `json:"IntranetEgressAddress,omitempty" xml:"IntranetEgressAddress,omitempty"`
	// The custom CIDR block. The configured CIDR block is considered as a private block.
	//
	// example:
	//
	// 123.0.0.1
	IntranetSegments *string `json:"IntranetSegments,omitempty" xml:"IntranetSegments,omitempty"`
	// Indicates whether IPv6 traffic is supported.
	//
	// example:
	//
	// true
	SupportIpv6 *bool `json:"SupportIpv6,omitempty" xml:"SupportIpv6,omitempty"`
	// The ID of the client VPC.
	//
	// example:
	//
	// vpc-2zew2v4vcg78mXXXX
	UserVpcId *string `json:"UserVpcId,omitempty" xml:"UserVpcId,omitempty"`
	// The vSwitch of the client VPC.
	//
	// example:
	//
	// vsw-2zecr5r7ao44tslsXXXX
	UserVswitchId *string `json:"UserVswitchId,omitempty" xml:"UserVswitchId,omitempty"`
	// The VIPs of the cluster.
	//
	// example:
	//
	// VPC_INTERNET_IPV6
	VipTypeList *string `json:"VipTypeList,omitempty" xml:"VipTypeList,omitempty"`
	// Indicates whether a virtual private cloud (VPC) domain name is enabled.
	VpcIntranetEnable *bool `json:"VpcIntranetEnable,omitempty" xml:"VpcIntranetEnable,omitempty"`
	// The ID of the account to which the VPC belongs.
	//
	// example:
	//
	// 165438596694XXXX
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
	// Indicates whether self-calling is enabled.
	//
	// example:
	//
	// false
	VpcSlbIntranetEnable *bool `json:"VpcSlbIntranetEnable,omitempty" xml:"VpcSlbIntranetEnable,omitempty"`
}

func (s DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetConnectCidrBlocks() *string {
	return s.ConnectCidrBlocks
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetConnectVpcId() *string {
	return s.ConnectVpcId
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetEgressIpv6Enable() *bool {
	return s.EgressIpv6Enable
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetHttpsPolicies() *string {
	return s.HttpsPolicies
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetIPV4AclId() *string {
	return s.IPV4AclId
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetIPV4AclName() *string {
	return s.IPV4AclName
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetIPV4AclStatus() *string {
	return s.IPV4AclStatus
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetIPV4AclType() *string {
	return s.IPV4AclType
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetIPV6AclId() *string {
	return s.IPV6AclId
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetIPV6AclName() *string {
	return s.IPV6AclName
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetIPV6AclStatus() *string {
	return s.IPV6AclStatus
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetIPV6AclType() *string {
	return s.IPV6AclType
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetInternetEgressAddress() *string {
	return s.InternetEgressAddress
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetIntranetEgressAddress() *string {
	return s.IntranetEgressAddress
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetIntranetSegments() *string {
	return s.IntranetSegments
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetSupportIpv6() *bool {
	return s.SupportIpv6
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetUserVpcId() *string {
	return s.UserVpcId
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetUserVswitchId() *string {
	return s.UserVswitchId
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetVipTypeList() *string {
	return s.VipTypeList
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetVpcIntranetEnable() *bool {
	return s.VpcIntranetEnable
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetVpcOwnerId() *int64 {
	return s.VpcOwnerId
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) GetVpcSlbIntranetEnable() *bool {
	return s.VpcSlbIntranetEnable
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetConnectCidrBlocks(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.ConnectCidrBlocks = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetConnectVpcId(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.ConnectVpcId = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetEgressIpv6Enable(v bool) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.EgressIpv6Enable = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetHttpsPolicies(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.HttpsPolicies = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetIPV4AclId(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.IPV4AclId = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetIPV4AclName(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.IPV4AclName = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetIPV4AclStatus(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.IPV4AclStatus = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetIPV4AclType(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.IPV4AclType = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetIPV6AclId(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.IPV6AclId = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetIPV6AclName(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.IPV6AclName = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetIPV6AclStatus(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.IPV6AclStatus = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetIPV6AclType(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.IPV6AclType = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetInternetEgressAddress(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.InternetEgressAddress = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetIntranetEgressAddress(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.IntranetEgressAddress = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetIntranetSegments(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.IntranetSegments = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetSupportIpv6(v bool) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.SupportIpv6 = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetUserVpcId(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.UserVpcId = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetUserVswitchId(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.UserVswitchId = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetVipTypeList(v string) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.VipTypeList = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetVpcIntranetEnable(v bool) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.VpcIntranetEnable = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetVpcOwnerId(v int64) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.VpcOwnerId = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) SetVpcSlbIntranetEnable(v bool) *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute {
	s.VpcSlbIntranetEnable = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceClusterAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceClusterInfoResponseBodyInstanceList struct {
	Instance []*DescribeInstanceClusterInfoResponseBodyInstanceListInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstanceClusterInfoResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceClusterInfoResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceList) GetInstance() []*DescribeInstanceClusterInfoResponseBodyInstanceListInstance {
	return s.Instance
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceList) SetInstance(v []*DescribeInstanceClusterInfoResponseBodyInstanceListInstance) *DescribeInstanceClusterInfoResponseBodyInstanceList {
	s.Instance = v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceList) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceClusterInfoResponseBodyInstanceListInstance struct {
	// The error message returned if the call fails.
	//
	// example:
	//
	// Instance not found.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// apigateway-ht-04e41d95e9c1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// tf-testacceu-central-1apigatewayinstance8752
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceClusterInfoResponseBodyInstanceListInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceClusterInfoResponseBodyInstanceListInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceListInstance) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceListInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceListInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceListInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceListInstance) SetErrorMessage(v string) *DescribeInstanceClusterInfoResponseBodyInstanceListInstance {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceListInstance) SetInstanceId(v string) *DescribeInstanceClusterInfoResponseBodyInstanceListInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceListInstance) SetInstanceName(v string) *DescribeInstanceClusterInfoResponseBodyInstanceListInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceListInstance) SetStatus(v string) *DescribeInstanceClusterInfoResponseBodyInstanceListInstance {
	s.Status = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponseBodyInstanceListInstance) Validate() error {
	return dara.Validate(s)
}
