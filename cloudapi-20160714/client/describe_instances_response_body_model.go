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
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.InstanceAttribute != nil {
		for _, item := range s.InstanceAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceAttribute struct {
	AclId                      *string                                                                            `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclName                    *string                                                                            `json:"AclName,omitempty" xml:"AclName,omitempty"`
	AclStatus                  *string                                                                            `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	AclType                    *string                                                                            `json:"AclType,omitempty" xml:"AclType,omitempty"`
	ClassicEgressAddress       *string                                                                            `json:"ClassicEgressAddress,omitempty" xml:"ClassicEgressAddress,omitempty"`
	ConnectCidrBlocks          *string                                                                            `json:"ConnectCidrBlocks,omitempty" xml:"ConnectCidrBlocks,omitempty"`
	ConnectVpcId               *string                                                                            `json:"ConnectVpcId,omitempty" xml:"ConnectVpcId,omitempty"`
	CreatedTime                *string                                                                            `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DedicatedInstanceType      *string                                                                            `json:"DedicatedInstanceType,omitempty" xml:"DedicatedInstanceType,omitempty"`
	EgressIpv6Enable           *bool                                                                              `json:"EgressIpv6Enable,omitempty" xml:"EgressIpv6Enable,omitempty"`
	ExpiredTime                *string                                                                            `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	HttpsPolicies              *string                                                                            `json:"HttpsPolicies,omitempty" xml:"HttpsPolicies,omitempty"`
	IPV6AclId                  *string                                                                            `json:"IPV6AclId,omitempty" xml:"IPV6AclId,omitempty"`
	IPV6AclName                *string                                                                            `json:"IPV6AclName,omitempty" xml:"IPV6AclName,omitempty"`
	IPV6AclStatus              *string                                                                            `json:"IPV6AclStatus,omitempty" xml:"IPV6AclStatus,omitempty"`
	IPV6AclType                *string                                                                            `json:"IPV6AclType,omitempty" xml:"IPV6AclType,omitempty"`
	InstanceChargeType         *string                                                                            `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceCidrBlock          *string                                                                            `json:"InstanceCidrBlock,omitempty" xml:"InstanceCidrBlock,omitempty"`
	InstanceClusterId          *string                                                                            `json:"InstanceClusterId,omitempty" xml:"InstanceClusterId,omitempty"`
	InstanceId                 *string                                                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName               *string                                                                            `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceRpsLimit           *int32                                                                             `json:"InstanceRpsLimit,omitempty" xml:"InstanceRpsLimit,omitempty"`
	InstanceSpec               *string                                                                            `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	InstanceSpecAttributes     *DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributes     `json:"InstanceSpecAttributes,omitempty" xml:"InstanceSpecAttributes,omitempty" type:"Struct"`
	InstanceType               *string                                                                            `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetEgressAddress      *string                                                                            `json:"InternetEgressAddress,omitempty" xml:"InternetEgressAddress,omitempty"`
	IntranetSegments           *string                                                                            `json:"IntranetSegments,omitempty" xml:"IntranetSegments,omitempty"`
	MaintainEndTime            *string                                                                            `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime          *string                                                                            `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	NetworkInterfaceAttributes *DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributes `json:"NetworkInterfaceAttributes,omitempty" xml:"NetworkInterfaceAttributes,omitempty" type:"Struct"`
	NewVpcEgressAddress        *string                                                                            `json:"NewVpcEgressAddress,omitempty" xml:"NewVpcEgressAddress,omitempty"`
	PrivateDnsList             *DescribeInstancesResponseBodyInstancesInstanceAttributePrivateDnsList             `json:"PrivateDnsList,omitempty" xml:"PrivateDnsList,omitempty" type:"Struct"`
	RegionId                   *string                                                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status                     *string                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportIpv6                *bool                                                                              `json:"SupportIpv6,omitempty" xml:"SupportIpv6,omitempty"`
	Tags                       *DescribeInstancesResponseBodyInstancesInstanceAttributeTags                       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UserVpcId                  *string                                                                            `json:"UserVpcId,omitempty" xml:"UserVpcId,omitempty"`
	UserVswitchId              *string                                                                            `json:"UserVswitchId,omitempty" xml:"UserVswitchId,omitempty"`
	VpcEgressAddress           *string                                                                            `json:"VpcEgressAddress,omitempty" xml:"VpcEgressAddress,omitempty"`
	VpcIntranetEnable          *bool                                                                              `json:"VpcIntranetEnable,omitempty" xml:"VpcIntranetEnable,omitempty"`
	VpcOwnerId                 *int64                                                                             `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
	VpcSlbIntranetEnable       *bool                                                                              `json:"VpcSlbIntranetEnable,omitempty" xml:"VpcSlbIntranetEnable,omitempty"`
	ZoneId                     *string                                                                            `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneLocalName              *string                                                                            `json:"ZoneLocalName,omitempty" xml:"ZoneLocalName,omitempty"`
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
	if s.InstanceSpecAttributes != nil {
		if err := s.InstanceSpecAttributes.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInterfaceAttributes != nil {
		if err := s.NetworkInterfaceAttributes.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateDnsList != nil {
		if err := s.PrivateDnsList.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.SpecAttribute != nil {
		for _, item := range s.SpecAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceAttributeInstanceSpecAttributesSpecAttribute struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	if s.NetworkInterfaceAttribute != nil {
		for _, item := range s.NetworkInterfaceAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceAttributeNetworkInterfaceAttributesNetworkInterfaceAttribute struct {
	CidrBlock       *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VswitchId       *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	if s.TagInfo != nil {
		for _, item := range s.TagInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceAttributeTagsTagInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
