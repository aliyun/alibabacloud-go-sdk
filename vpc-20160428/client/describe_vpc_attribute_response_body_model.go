// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssociatedCens(v *DescribeVpcAttributeResponseBodyAssociatedCens) *DescribeVpcAttributeResponseBody
	GetAssociatedCens() *DescribeVpcAttributeResponseBodyAssociatedCens
	SetAssociatedPropagationSources(v *DescribeVpcAttributeResponseBodyAssociatedPropagationSources) *DescribeVpcAttributeResponseBody
	GetAssociatedPropagationSources() *DescribeVpcAttributeResponseBodyAssociatedPropagationSources
	SetCidrBlock(v string) *DescribeVpcAttributeResponseBody
	GetCidrBlock() *string
	SetClassicLinkEnabled(v bool) *DescribeVpcAttributeResponseBody
	GetClassicLinkEnabled() *bool
	SetCloudResources(v *DescribeVpcAttributeResponseBodyCloudResources) *DescribeVpcAttributeResponseBody
	GetCloudResources() *DescribeVpcAttributeResponseBodyCloudResources
	SetCreationTime(v string) *DescribeVpcAttributeResponseBody
	GetCreationTime() *string
	SetDescription(v string) *DescribeVpcAttributeResponseBody
	GetDescription() *string
	SetDhcpOptionsSetId(v string) *DescribeVpcAttributeResponseBody
	GetDhcpOptionsSetId() *string
	SetDhcpOptionsSetStatus(v string) *DescribeVpcAttributeResponseBody
	GetDhcpOptionsSetStatus() *string
	SetDnsHostnameStatus(v string) *DescribeVpcAttributeResponseBody
	GetDnsHostnameStatus() *string
	SetEnabledIpv6(v bool) *DescribeVpcAttributeResponseBody
	GetEnabledIpv6() *bool
	SetIpv4GatewayId(v string) *DescribeVpcAttributeResponseBody
	GetIpv4GatewayId() *string
	SetIpv6CidrBlock(v string) *DescribeVpcAttributeResponseBody
	GetIpv6CidrBlock() *string
	SetIpv6CidrBlocks(v *DescribeVpcAttributeResponseBodyIpv6CidrBlocks) *DescribeVpcAttributeResponseBody
	GetIpv6CidrBlocks() *DescribeVpcAttributeResponseBodyIpv6CidrBlocks
	SetIsDefault(v bool) *DescribeVpcAttributeResponseBody
	GetIsDefault() *bool
	SetOwnerId(v int64) *DescribeVpcAttributeResponseBody
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVpcAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeVpcAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeVpcAttributeResponseBody
	GetResourceGroupId() *string
	SetSecondaryCidrBlocks(v *DescribeVpcAttributeResponseBodySecondaryCidrBlocks) *DescribeVpcAttributeResponseBody
	GetSecondaryCidrBlocks() *DescribeVpcAttributeResponseBodySecondaryCidrBlocks
	SetStatus(v string) *DescribeVpcAttributeResponseBody
	GetStatus() *string
	SetSupportIpv4Gateway(v bool) *DescribeVpcAttributeResponseBody
	GetSupportIpv4Gateway() *bool
	SetTags(v *DescribeVpcAttributeResponseBodyTags) *DescribeVpcAttributeResponseBody
	GetTags() *DescribeVpcAttributeResponseBodyTags
	SetUserCidrs(v *DescribeVpcAttributeResponseBodyUserCidrs) *DescribeVpcAttributeResponseBody
	GetUserCidrs() *DescribeVpcAttributeResponseBodyUserCidrs
	SetVRouterId(v string) *DescribeVpcAttributeResponseBody
	GetVRouterId() *string
	SetVSwitchIds(v *DescribeVpcAttributeResponseBodyVSwitchIds) *DescribeVpcAttributeResponseBody
	GetVSwitchIds() *DescribeVpcAttributeResponseBodyVSwitchIds
	SetVpcId(v string) *DescribeVpcAttributeResponseBody
	GetVpcId() *string
	SetVpcName(v string) *DescribeVpcAttributeResponseBody
	GetVpcName() *string
}

type DescribeVpcAttributeResponseBody struct {
	AssociatedCens               *DescribeVpcAttributeResponseBodyAssociatedCens               `json:"AssociatedCens,omitempty" xml:"AssociatedCens,omitempty" type:"Struct"`
	AssociatedPropagationSources *DescribeVpcAttributeResponseBodyAssociatedPropagationSources `json:"AssociatedPropagationSources,omitempty" xml:"AssociatedPropagationSources,omitempty" type:"Struct"`
	// The IPv4 CIDR block of the VPC.
	//
	// example:
	//
	// 192.168.0.0/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// Deprecated
	//
	// Indicates whether the ClassicLink feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ClassicLinkEnabled *bool                                           `json:"ClassicLinkEnabled,omitempty" xml:"ClassicLinkEnabled,omitempty"`
	CloudResources     *DescribeVpcAttributeResponseBodyCloudResources `json:"CloudResources,omitempty" xml:"CloudResources,omitempty" type:"Struct"`
	// The time when the VPC was created.
	//
	// example:
	//
	// 2021-10-16T07:31:09Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the VPC.
	//
	// example:
	//
	// VPC
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the DHCP options set.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	// The status of the DHCP options set. Valid values:
	//
	// 	- **Available**
	//
	// 	- **InUse**
	//
	// 	- **Deleted**
	//
	// 	- **Pending**
	//
	// example:
	//
	// Available
	DhcpOptionsSetStatus *string `json:"DhcpOptionsSetStatus,omitempty" xml:"DhcpOptionsSetStatus,omitempty"`
	// Indicates whether DNS hostname is enabled.
	//
	// example:
	//
	// DISABLED
	DnsHostnameStatus *string `json:"DnsHostnameStatus,omitempty" xml:"DnsHostnameStatus,omitempty"`
	// Indicates whether the VPC enables IPv6 .
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	EnabledIpv6 *bool `json:"EnabledIpv6,omitempty" xml:"EnabledIpv6,omitempty"`
	// The ID of the IPv4 gateway.
	//
	// example:
	//
	// ipv4gw-5tsnc6s4ogsedtp3k****
	Ipv4GatewayId *string `json:"Ipv4GatewayId,omitempty" xml:"Ipv4GatewayId,omitempty"`
	// The IPv6 CIDR block of the VPC.
	//
	// example:
	//
	// 2408:XXXX:0:a600::/56
	Ipv6CidrBlock  *string                                         `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	Ipv6CidrBlocks *DescribeVpcAttributeResponseBodyIpv6CidrBlocks `json:"Ipv6CidrBlocks,omitempty" xml:"Ipv6CidrBlocks,omitempty" type:"Struct"`
	// Indicates whether the VPC is the default VPC. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 28311773240248****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the VPC belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7486AE4A-129D-43DB-A714-2432C074BA04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazbvgb4ph****
	ResourceGroupId     *string                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecondaryCidrBlocks *DescribeVpcAttributeResponseBodySecondaryCidrBlocks `json:"SecondaryCidrBlocks,omitempty" xml:"SecondaryCidrBlocks,omitempty" type:"Struct"`
	// The status of the VPC. Valid values:
	//
	// 	- **Available**
	//
	// 	- **Pending**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the VPC supports IPv4 gateways.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SupportIpv4Gateway *bool                                      `json:"SupportIpv4Gateway,omitempty" xml:"SupportIpv4Gateway,omitempty"`
	Tags               *DescribeVpcAttributeResponseBodyTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UserCidrs          *DescribeVpcAttributeResponseBodyUserCidrs `json:"UserCidrs,omitempty" xml:"UserCidrs,omitempty" type:"Struct"`
	// The ID of the vRouter that belongs to the VPC.
	//
	// example:
	//
	// vrt-bp1jso6ng1at0ajsc****
	VRouterId  *string                                     `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
	VSwitchIds *DescribeVpcAttributeResponseBodyVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp18sth14qii3pnvo****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// doctest2
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBody) GetAssociatedCens() *DescribeVpcAttributeResponseBodyAssociatedCens {
	return s.AssociatedCens
}

func (s *DescribeVpcAttributeResponseBody) GetAssociatedPropagationSources() *DescribeVpcAttributeResponseBodyAssociatedPropagationSources {
	return s.AssociatedPropagationSources
}

func (s *DescribeVpcAttributeResponseBody) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVpcAttributeResponseBody) GetClassicLinkEnabled() *bool {
	return s.ClassicLinkEnabled
}

func (s *DescribeVpcAttributeResponseBody) GetCloudResources() *DescribeVpcAttributeResponseBodyCloudResources {
	return s.CloudResources
}

func (s *DescribeVpcAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVpcAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeVpcAttributeResponseBody) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *DescribeVpcAttributeResponseBody) GetDhcpOptionsSetStatus() *string {
	return s.DhcpOptionsSetStatus
}

func (s *DescribeVpcAttributeResponseBody) GetDnsHostnameStatus() *string {
	return s.DnsHostnameStatus
}

func (s *DescribeVpcAttributeResponseBody) GetEnabledIpv6() *bool {
	return s.EnabledIpv6
}

func (s *DescribeVpcAttributeResponseBody) GetIpv4GatewayId() *string {
	return s.Ipv4GatewayId
}

func (s *DescribeVpcAttributeResponseBody) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *DescribeVpcAttributeResponseBody) GetIpv6CidrBlocks() *DescribeVpcAttributeResponseBodyIpv6CidrBlocks {
	return s.Ipv6CidrBlocks
}

func (s *DescribeVpcAttributeResponseBody) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcAttributeResponseBody) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpcAttributeResponseBody) GetSecondaryCidrBlocks() *DescribeVpcAttributeResponseBodySecondaryCidrBlocks {
	return s.SecondaryCidrBlocks
}

func (s *DescribeVpcAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcAttributeResponseBody) GetSupportIpv4Gateway() *bool {
	return s.SupportIpv4Gateway
}

func (s *DescribeVpcAttributeResponseBody) GetTags() *DescribeVpcAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeVpcAttributeResponseBody) GetUserCidrs() *DescribeVpcAttributeResponseBodyUserCidrs {
	return s.UserCidrs
}

func (s *DescribeVpcAttributeResponseBody) GetVRouterId() *string {
	return s.VRouterId
}

func (s *DescribeVpcAttributeResponseBody) GetVSwitchIds() *DescribeVpcAttributeResponseBodyVSwitchIds {
	return s.VSwitchIds
}

func (s *DescribeVpcAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcAttributeResponseBody) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcAttributeResponseBody) SetAssociatedCens(v *DescribeVpcAttributeResponseBodyAssociatedCens) *DescribeVpcAttributeResponseBody {
	s.AssociatedCens = v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetAssociatedPropagationSources(v *DescribeVpcAttributeResponseBodyAssociatedPropagationSources) *DescribeVpcAttributeResponseBody {
	s.AssociatedPropagationSources = v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetCidrBlock(v string) *DescribeVpcAttributeResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetClassicLinkEnabled(v bool) *DescribeVpcAttributeResponseBody {
	s.ClassicLinkEnabled = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetCloudResources(v *DescribeVpcAttributeResponseBodyCloudResources) *DescribeVpcAttributeResponseBody {
	s.CloudResources = v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetCreationTime(v string) *DescribeVpcAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetDescription(v string) *DescribeVpcAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetDhcpOptionsSetId(v string) *DescribeVpcAttributeResponseBody {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetDhcpOptionsSetStatus(v string) *DescribeVpcAttributeResponseBody {
	s.DhcpOptionsSetStatus = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetDnsHostnameStatus(v string) *DescribeVpcAttributeResponseBody {
	s.DnsHostnameStatus = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetEnabledIpv6(v bool) *DescribeVpcAttributeResponseBody {
	s.EnabledIpv6 = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetIpv4GatewayId(v string) *DescribeVpcAttributeResponseBody {
	s.Ipv4GatewayId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetIpv6CidrBlock(v string) *DescribeVpcAttributeResponseBody {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetIpv6CidrBlocks(v *DescribeVpcAttributeResponseBodyIpv6CidrBlocks) *DescribeVpcAttributeResponseBody {
	s.Ipv6CidrBlocks = v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetIsDefault(v bool) *DescribeVpcAttributeResponseBody {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetOwnerId(v int64) *DescribeVpcAttributeResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetRegionId(v string) *DescribeVpcAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetRequestId(v string) *DescribeVpcAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetResourceGroupId(v string) *DescribeVpcAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetSecondaryCidrBlocks(v *DescribeVpcAttributeResponseBodySecondaryCidrBlocks) *DescribeVpcAttributeResponseBody {
	s.SecondaryCidrBlocks = v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetStatus(v string) *DescribeVpcAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetSupportIpv4Gateway(v bool) *DescribeVpcAttributeResponseBody {
	s.SupportIpv4Gateway = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetTags(v *DescribeVpcAttributeResponseBodyTags) *DescribeVpcAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetUserCidrs(v *DescribeVpcAttributeResponseBodyUserCidrs) *DescribeVpcAttributeResponseBody {
	s.UserCidrs = v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetVRouterId(v string) *DescribeVpcAttributeResponseBody {
	s.VRouterId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetVSwitchIds(v *DescribeVpcAttributeResponseBodyVSwitchIds) *DescribeVpcAttributeResponseBody {
	s.VSwitchIds = v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetVpcId(v string) *DescribeVpcAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) SetVpcName(v string) *DescribeVpcAttributeResponseBody {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcAttributeResponseBody) Validate() error {
	if s.AssociatedCens != nil {
		if err := s.AssociatedCens.Validate(); err != nil {
			return err
		}
	}
	if s.AssociatedPropagationSources != nil {
		if err := s.AssociatedPropagationSources.Validate(); err != nil {
			return err
		}
	}
	if s.CloudResources != nil {
		if err := s.CloudResources.Validate(); err != nil {
			return err
		}
	}
	if s.Ipv6CidrBlocks != nil {
		if err := s.Ipv6CidrBlocks.Validate(); err != nil {
			return err
		}
	}
	if s.SecondaryCidrBlocks != nil {
		if err := s.SecondaryCidrBlocks.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.UserCidrs != nil {
		if err := s.UserCidrs.Validate(); err != nil {
			return err
		}
	}
	if s.VSwitchIds != nil {
		if err := s.VSwitchIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpcAttributeResponseBodyAssociatedCens struct {
	AssociatedCen []*DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen `json:"AssociatedCen,omitempty" xml:"AssociatedCen,omitempty" type:"Repeated"`
}

func (s DescribeVpcAttributeResponseBodyAssociatedCens) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyAssociatedCens) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyAssociatedCens) GetAssociatedCen() []*DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen {
	return s.AssociatedCen
}

func (s *DescribeVpcAttributeResponseBodyAssociatedCens) SetAssociatedCen(v []*DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen) *DescribeVpcAttributeResponseBodyAssociatedCens {
	s.AssociatedCen = v
	return s
}

func (s *DescribeVpcAttributeResponseBodyAssociatedCens) Validate() error {
	if s.AssociatedCen != nil {
		for _, item := range s.AssociatedCen {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen struct {
	CenId      *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenOwnerId *int64  `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	CenStatus  *string `json:"CenStatus,omitempty" xml:"CenStatus,omitempty"`
}

func (s DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen) GetCenId() *string {
	return s.CenId
}

func (s *DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen) GetCenStatus() *string {
	return s.CenStatus
}

func (s *DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen) SetCenId(v string) *DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen {
	s.CenId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen) SetCenOwnerId(v int64) *DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen {
	s.CenOwnerId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen) SetCenStatus(v string) *DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen {
	s.CenStatus = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyAssociatedCensAssociatedCen) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAttributeResponseBodyAssociatedPropagationSources struct {
	AssociatedPropagationSources []*DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources `json:"AssociatedPropagationSources,omitempty" xml:"AssociatedPropagationSources,omitempty" type:"Repeated"`
}

func (s DescribeVpcAttributeResponseBodyAssociatedPropagationSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyAssociatedPropagationSources) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSources) GetAssociatedPropagationSources() []*DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources {
	return s.AssociatedPropagationSources
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSources) SetAssociatedPropagationSources(v []*DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) *DescribeVpcAttributeResponseBodyAssociatedPropagationSources {
	s.AssociatedPropagationSources = v
	return s
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSources) Validate() error {
	if s.AssociatedPropagationSources != nil {
		for _, item := range s.AssociatedPropagationSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources struct {
	RoutePropagated  *bool   `json:"RoutePropagated,omitempty" xml:"RoutePropagated,omitempty"`
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	SourceOwnerId    *int64  `json:"SourceOwnerId,omitempty" xml:"SourceOwnerId,omitempty"`
	SourceType       *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) GetRoutePropagated() *bool {
	return s.RoutePropagated
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) GetSourceOwnerId() *int64 {
	return s.SourceOwnerId
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) SetRoutePropagated(v bool) *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources {
	s.RoutePropagated = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) SetSourceInstanceId(v string) *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources {
	s.SourceInstanceId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) SetSourceOwnerId(v int64) *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources {
	s.SourceOwnerId = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) SetSourceType(v string) *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources {
	s.SourceType = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) SetStatus(v string) *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources {
	s.Status = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyAssociatedPropagationSourcesAssociatedPropagationSources) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAttributeResponseBodyCloudResources struct {
	CloudResourceSetType []*DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType `json:"CloudResourceSetType,omitempty" xml:"CloudResourceSetType,omitempty" type:"Repeated"`
}

func (s DescribeVpcAttributeResponseBodyCloudResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyCloudResources) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyCloudResources) GetCloudResourceSetType() []*DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType {
	return s.CloudResourceSetType
}

func (s *DescribeVpcAttributeResponseBodyCloudResources) SetCloudResourceSetType(v []*DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType) *DescribeVpcAttributeResponseBodyCloudResources {
	s.CloudResourceSetType = v
	return s
}

func (s *DescribeVpcAttributeResponseBodyCloudResources) Validate() error {
	if s.CloudResourceSetType != nil {
		for _, item := range s.CloudResourceSetType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType struct {
	ResourceCount *int32  `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType) GetResourceCount() *int32 {
	return s.ResourceCount
}

func (s *DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType) SetResourceCount(v int32) *DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType {
	s.ResourceCount = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType) SetResourceType(v string) *DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType {
	s.ResourceType = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyCloudResourcesCloudResourceSetType) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAttributeResponseBodyIpv6CidrBlocks struct {
	Ipv6CidrBlock []*DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty" type:"Repeated"`
}

func (s DescribeVpcAttributeResponseBodyIpv6CidrBlocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyIpv6CidrBlocks) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyIpv6CidrBlocks) GetIpv6CidrBlock() []*DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock {
	return s.Ipv6CidrBlock
}

func (s *DescribeVpcAttributeResponseBodyIpv6CidrBlocks) SetIpv6CidrBlock(v []*DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock) *DescribeVpcAttributeResponseBodyIpv6CidrBlocks {
	s.Ipv6CidrBlock = v
	return s
}

func (s *DescribeVpcAttributeResponseBodyIpv6CidrBlocks) Validate() error {
	if s.Ipv6CidrBlock != nil {
		for _, item := range s.Ipv6CidrBlock {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock struct {
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	Ipv6Isp       *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
}

func (s DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock) SetIpv6CidrBlock(v string) *DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock) SetIpv6Isp(v string) *DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock {
	s.Ipv6Isp = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyIpv6CidrBlocksIpv6CidrBlock) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAttributeResponseBodySecondaryCidrBlocks struct {
	SecondaryCidrBlock []*string `json:"SecondaryCidrBlock,omitempty" xml:"SecondaryCidrBlock,omitempty" type:"Repeated"`
}

func (s DescribeVpcAttributeResponseBodySecondaryCidrBlocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodySecondaryCidrBlocks) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodySecondaryCidrBlocks) GetSecondaryCidrBlock() []*string {
	return s.SecondaryCidrBlock
}

func (s *DescribeVpcAttributeResponseBodySecondaryCidrBlocks) SetSecondaryCidrBlock(v []*string) *DescribeVpcAttributeResponseBodySecondaryCidrBlocks {
	s.SecondaryCidrBlock = v
	return s
}

func (s *DescribeVpcAttributeResponseBodySecondaryCidrBlocks) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAttributeResponseBodyTags struct {
	Tag []*DescribeVpcAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVpcAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyTags) GetTag() []*DescribeVpcAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeVpcAttributeResponseBodyTags) SetTag(v []*DescribeVpcAttributeResponseBodyTagsTag) *DescribeVpcAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeVpcAttributeResponseBodyTags) Validate() error {
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

type DescribeVpcAttributeResponseBodyTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpcAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVpcAttributeResponseBodyTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVpcAttributeResponseBodyTagsTag) SetKey(v string) *DescribeVpcAttributeResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyTagsTag) SetValue(v string) *DescribeVpcAttributeResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeVpcAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAttributeResponseBodyUserCidrs struct {
	UserCidr []*string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty" type:"Repeated"`
}

func (s DescribeVpcAttributeResponseBodyUserCidrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyUserCidrs) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyUserCidrs) GetUserCidr() []*string {
	return s.UserCidr
}

func (s *DescribeVpcAttributeResponseBodyUserCidrs) SetUserCidr(v []*string) *DescribeVpcAttributeResponseBodyUserCidrs {
	s.UserCidr = v
	return s
}

func (s *DescribeVpcAttributeResponseBodyUserCidrs) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAttributeResponseBodyVSwitchIds struct {
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s DescribeVpcAttributeResponseBodyVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAttributeResponseBodyVSwitchIds) GoString() string {
	return s.String()
}

func (s *DescribeVpcAttributeResponseBodyVSwitchIds) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *DescribeVpcAttributeResponseBodyVSwitchIds) SetVSwitchId(v []*string) *DescribeVpcAttributeResponseBodyVSwitchIds {
	s.VSwitchId = v
	return s
}

func (s *DescribeVpcAttributeResponseBodyVSwitchIds) Validate() error {
	return dara.Validate(s)
}
