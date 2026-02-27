// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatGateways(v *DescribeNatGatewaysResponseBodyNatGateways) *DescribeNatGatewaysResponseBody
	GetNatGateways() *DescribeNatGatewaysResponseBodyNatGateways
	SetPageNumber(v int32) *DescribeNatGatewaysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNatGatewaysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeNatGatewaysResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNatGatewaysResponseBody
	GetTotalCount() *int32
}

type DescribeNatGatewaysResponseBody struct {
	NatGateways *DescribeNatGatewaysResponseBodyNatGateways `json:"NatGateways,omitempty" xml:"NatGateways,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of NAT gateway entries that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNatGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBody) GetNatGateways() *DescribeNatGatewaysResponseBodyNatGateways {
	return s.NatGateways
}

func (s *DescribeNatGatewaysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNatGatewaysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNatGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatGatewaysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNatGatewaysResponseBody) SetNatGateways(v *DescribeNatGatewaysResponseBodyNatGateways) *DescribeNatGatewaysResponseBody {
	s.NatGateways = v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetPageNumber(v int32) *DescribeNatGatewaysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetPageSize(v int32) *DescribeNatGatewaysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetRequestId(v string) *DescribeNatGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetTotalCount(v int32) *DescribeNatGatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) Validate() error {
	if s.NatGateways != nil {
		if err := s.NatGateways.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNatGatewaysResponseBodyNatGateways struct {
	NatGateway []*DescribeNatGatewaysResponseBodyNatGatewaysNatGateway `json:"NatGateway,omitempty" xml:"NatGateway,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewaysResponseBodyNatGateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGateways) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetNatGateway() []*DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	return s.NatGateway
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetNatGateway(v []*DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) *DescribeNatGatewaysResponseBodyNatGateways {
	s.NatGateway = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) Validate() error {
	if s.NatGateway != nil {
		for _, item := range s.NatGateway {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGateway struct {
	AccessMode         *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode `json:"AccessMode,omitempty" xml:"AccessMode,omitempty" type:"Struct"`
	AutoPay            *bool                                                           `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	BusinessStatus     *string                                                         `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	CreationTime       *string                                                         `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DeletionProtection *bool                                                           `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	Description        *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsMetricEnabled   *bool                                                           `json:"EcsMetricEnabled,omitempty" xml:"EcsMetricEnabled,omitempty"`
	// example:
	//
	// MULTI_BINDED
	EipBindMode               *string                                                                    `json:"EipBindMode,omitempty" xml:"EipBindMode,omitempty"`
	EnableSessionLog          *string                                                                    `json:"EnableSessionLog,omitempty" xml:"EnableSessionLog,omitempty"`
	ExpiredTime               *string                                                                    `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ForwardTableIds           *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds       `json:"ForwardTableIds,omitempty" xml:"ForwardTableIds,omitempty" type:"Struct"`
	FullNatTableIds           *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayFullNatTableIds       `json:"FullNatTableIds,omitempty" xml:"FullNatTableIds,omitempty" type:"Struct"`
	IcmpReplyEnabled          *bool                                                                      `json:"IcmpReplyEnabled,omitempty" xml:"IcmpReplyEnabled,omitempty"`
	InstanceChargeType        *string                                                                    `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType        *string                                                                    `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	IpLists                   *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpLists               `json:"IpLists,omitempty" xml:"IpLists,omitempty" type:"Struct"`
	IpPrefixList              *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixList          `json:"IpPrefixList,omitempty" xml:"IpPrefixList,omitempty" type:"Struct"`
	Name                      *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	NatGatewayId              *string                                                                    `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	NatGatewayPrivateInfo     *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo `json:"NatGatewayPrivateInfo,omitempty" xml:"NatGatewayPrivateInfo,omitempty" type:"Struct"`
	NatType                   *string                                                                    `json:"NatType,omitempty" xml:"NatType,omitempty"`
	NetworkType               *string                                                                    `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PrivateLinkEnabled        *bool                                                                      `json:"PrivateLinkEnabled,omitempty" xml:"PrivateLinkEnabled,omitempty"`
	PrivateLinkMode           *string                                                                    `json:"PrivateLinkMode,omitempty" xml:"PrivateLinkMode,omitempty"`
	RegionId                  *string                                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId           *string                                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityProtectionEnabled *bool                                                                      `json:"SecurityProtectionEnabled,omitempty" xml:"SecurityProtectionEnabled,omitempty"`
	SnatTableIds              *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewaySnatTableIds          `json:"SnatTableIds,omitempty" xml:"SnatTableIds,omitempty" type:"Struct"`
	Spec                      *string                                                                    `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status                    *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                      *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTags                  `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcId                     *string                                                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetAccessMode() *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode {
	return s.AccessMode
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetDescription() *string {
	return s.Description
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetEcsMetricEnabled() *bool {
	return s.EcsMetricEnabled
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetEipBindMode() *string {
	return s.EipBindMode
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetEnableSessionLog() *string {
	return s.EnableSessionLog
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetForwardTableIds() *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds {
	return s.ForwardTableIds
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetFullNatTableIds() *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayFullNatTableIds {
	return s.FullNatTableIds
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetIcmpReplyEnabled() *bool {
	return s.IcmpReplyEnabled
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetIpLists() *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpLists {
	return s.IpLists
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetIpPrefixList() *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixList {
	return s.IpPrefixList
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetName() *string {
	return s.Name
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetNatGatewayPrivateInfo() *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo {
	return s.NatGatewayPrivateInfo
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetNatType() *string {
	return s.NatType
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetPrivateLinkEnabled() *bool {
	return s.PrivateLinkEnabled
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetPrivateLinkMode() *string {
	return s.PrivateLinkMode
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetSecurityProtectionEnabled() *bool {
	return s.SecurityProtectionEnabled
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetSnatTableIds() *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewaySnatTableIds {
	return s.SnatTableIds
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetSpec() *string {
	return s.Spec
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetStatus() *string {
	return s.Status
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetTags() *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTags {
	return s.Tags
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetAccessMode(v *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.AccessMode = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetAutoPay(v bool) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.AutoPay = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetBusinessStatus(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetCreationTime(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.CreationTime = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetDeletionProtection(v bool) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetDescription(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.Description = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetEcsMetricEnabled(v bool) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.EcsMetricEnabled = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetEipBindMode(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.EipBindMode = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetEnableSessionLog(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.EnableSessionLog = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetExpiredTime(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetForwardTableIds(v *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.ForwardTableIds = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetFullNatTableIds(v *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayFullNatTableIds) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.FullNatTableIds = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetIcmpReplyEnabled(v bool) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.IcmpReplyEnabled = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetInstanceChargeType(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetInternetChargeType(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetIpLists(v *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpLists) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.IpLists = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetIpPrefixList(v *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixList) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.IpPrefixList = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetName(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.Name = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetNatGatewayId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetNatGatewayPrivateInfo(v *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.NatGatewayPrivateInfo = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetNatType(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.NatType = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetNetworkType(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.NetworkType = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetPrivateLinkEnabled(v bool) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.PrivateLinkEnabled = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetPrivateLinkMode(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.PrivateLinkMode = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetRegionId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.RegionId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetResourceGroupId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetSecurityProtectionEnabled(v bool) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.SecurityProtectionEnabled = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetSnatTableIds(v *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewaySnatTableIds) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.SnatTableIds = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetSpec(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.Spec = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetStatus(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.Status = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetTags(v *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTags) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.Tags = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetVpcId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.VpcId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) Validate() error {
	if s.AccessMode != nil {
		if err := s.AccessMode.Validate(); err != nil {
			return err
		}
	}
	if s.ForwardTableIds != nil {
		if err := s.ForwardTableIds.Validate(); err != nil {
			return err
		}
	}
	if s.FullNatTableIds != nil {
		if err := s.FullNatTableIds.Validate(); err != nil {
			return err
		}
	}
	if s.IpLists != nil {
		if err := s.IpLists.Validate(); err != nil {
			return err
		}
	}
	if s.IpPrefixList != nil {
		if err := s.IpPrefixList.Validate(); err != nil {
			return err
		}
	}
	if s.NatGatewayPrivateInfo != nil {
		if err := s.NatGatewayPrivateInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SnatTableIds != nil {
		if err := s.SnatTableIds.Validate(); err != nil {
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

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode struct {
	ModeValue  *string `json:"ModeValue,omitempty" xml:"ModeValue,omitempty"`
	TunnelType *string `json:"TunnelType,omitempty" xml:"TunnelType,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode) GetModeValue() *string {
	return s.ModeValue
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode) GetTunnelType() *string {
	return s.TunnelType
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode) SetModeValue(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode {
	s.ModeValue = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode) SetTunnelType(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode {
	s.TunnelType = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayAccessMode) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds struct {
	ForwardTableId []*string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) GetForwardTableId() []*string {
	return s.ForwardTableId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) SetForwardTableId(v []*string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds {
	s.ForwardTableId = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayFullNatTableIds struct {
	FullNatTableId []*string `json:"FullNatTableId,omitempty" xml:"FullNatTableId,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayFullNatTableIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayFullNatTableIds) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayFullNatTableIds) GetFullNatTableId() []*string {
	return s.FullNatTableId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayFullNatTableIds) SetFullNatTableId(v []*string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayFullNatTableIds {
	s.FullNatTableId = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayFullNatTableIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpLists struct {
	IpList []*DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpLists) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpLists) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpLists) GetIpList() []*DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList {
	return s.IpList
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpLists) SetIpList(v []*DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpLists {
	s.IpList = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpLists) Validate() error {
	if s.IpList != nil {
		for _, item := range s.IpList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList struct {
	AllocationId     *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	IpAddress        *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	SnatEntryEnabled *bool   `json:"SnatEntryEnabled,omitempty" xml:"SnatEntryEnabled,omitempty"`
	UsingStatus      *string `json:"UsingStatus,omitempty" xml:"UsingStatus,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) GetSnatEntryEnabled() *bool {
	return s.SnatEntryEnabled
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) GetUsingStatus() *string {
	return s.UsingStatus
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) SetAllocationId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList {
	s.AllocationId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) SetIpAddress(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList {
	s.IpAddress = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) SetPrivateIpAddress(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) SetSnatEntryEnabled(v bool) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList {
	s.SnatEntryEnabled = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) SetUsingStatus(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList {
	s.UsingStatus = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpListsIpList) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixList struct {
	IpPrefixList []*DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixListIpPrefixList `json:"IpPrefixList,omitempty" xml:"IpPrefixList,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixList) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixList) GetIpPrefixList() []*DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixListIpPrefixList {
	return s.IpPrefixList
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixList) SetIpPrefixList(v []*DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixListIpPrefixList) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixList {
	s.IpPrefixList = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixList) Validate() error {
	if s.IpPrefixList != nil {
		for _, item := range s.IpPrefixList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixListIpPrefixList struct {
	IpPrefix *string `json:"IpPrefix,omitempty" xml:"IpPrefix,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixListIpPrefixList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixListIpPrefixList) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixListIpPrefixList) GetIpPrefix() *string {
	return s.IpPrefix
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixListIpPrefixList) SetIpPrefix(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixListIpPrefixList {
	s.IpPrefix = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayIpPrefixListIpPrefixList) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo struct {
	EniInstanceId           *string `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty"`
	EniType                 *string `json:"EniType,omitempty" xml:"EniType,omitempty"`
	IzNo                    *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	MaxBandwidth            *int32  `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	MaxSessionEstablishRate *int32  `json:"MaxSessionEstablishRate,omitempty" xml:"MaxSessionEstablishRate,omitempty"`
	MaxSessionQuota         *int32  `json:"MaxSessionQuota,omitempty" xml:"MaxSessionQuota,omitempty"`
	PrivateIpAddress        *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	VswitchId               *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) GetEniInstanceId() *string {
	return s.EniInstanceId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) GetEniType() *string {
	return s.EniType
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) GetMaxBandwidth() *int32 {
	return s.MaxBandwidth
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) GetMaxSessionEstablishRate() *int32 {
	return s.MaxSessionEstablishRate
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) GetMaxSessionQuota() *int32 {
	return s.MaxSessionQuota
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) SetEniInstanceId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo {
	s.EniInstanceId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) SetEniType(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo {
	s.EniType = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) SetIzNo(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo {
	s.IzNo = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) SetMaxBandwidth(v int32) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo {
	s.MaxBandwidth = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) SetMaxSessionEstablishRate(v int32) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo {
	s.MaxSessionEstablishRate = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) SetMaxSessionQuota(v int32) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo {
	s.MaxSessionQuota = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) SetPrivateIpAddress(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) SetVswitchId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo {
	s.VswitchId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayNatGatewayPrivateInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewaySnatTableIds struct {
	SnatTableId []*string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewaySnatTableIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewaySnatTableIds) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewaySnatTableIds) GetSnatTableId() []*string {
	return s.SnatTableId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewaySnatTableIds) SetSnatTableId(v []*string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewaySnatTableIds {
	s.SnatTableId = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewaySnatTableIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTags struct {
	Tag []*DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTags) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTags) GetTag() []*DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag {
	return s.Tag
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTags) SetTag(v []*DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTags {
	s.Tag = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag) SetTagKey(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag) SetTagValue(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayTagsTag) Validate() error {
	return dara.Validate(s)
}
