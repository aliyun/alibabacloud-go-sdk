// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVpcsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpcsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpcsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcsResponseBody
	GetTotalCount() *int32
	SetVpcs(v *DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody
	GetVpcs() *DescribeVpcsResponseBodyVpcs
}

type DescribeVpcsResponseBody struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C6532AA8-D0F7-497F-A8EE-094126D441F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Vpcs       *DescribeVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
}

func (s DescribeVpcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpcsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcsResponseBody) GetVpcs() *DescribeVpcsResponseBodyVpcs {
	return s.Vpcs
}

func (s *DescribeVpcsResponseBody) SetPageNumber(v int32) *DescribeVpcsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetPageSize(v int32) *DescribeVpcsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetRequestId(v string) *DescribeVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetTotalCount(v int32) *DescribeVpcsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetVpcs(v *DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody {
	s.Vpcs = v
	return s
}

func (s *DescribeVpcsResponseBody) Validate() error {
	if s.Vpcs != nil {
		if err := s.Vpcs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpcsResponseBodyVpcs struct {
	Vpc []*DescribeVpcsResponseBodyVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcs) GetVpc() []*DescribeVpcsResponseBodyVpcsVpc {
	return s.Vpc
}

func (s *DescribeVpcsResponseBodyVpcs) SetVpc(v []*DescribeVpcsResponseBodyVpcsVpc) *DescribeVpcsResponseBodyVpcs {
	s.Vpc = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) Validate() error {
	if s.Vpc != nil {
		for _, item := range s.Vpc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcsResponseBodyVpcsVpc struct {
	CenStatus            *string                                             `json:"CenStatus,omitempty" xml:"CenStatus,omitempty"`
	CidrBlock            *string                                             `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	CreationTime         *string                                             `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description          *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DhcpOptionsSetId     *string                                             `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	DhcpOptionsSetStatus *string                                             `json:"DhcpOptionsSetStatus,omitempty" xml:"DhcpOptionsSetStatus,omitempty"`
	DnsHostnameStatus    *string                                             `json:"DnsHostnameStatus,omitempty" xml:"DnsHostnameStatus,omitempty"`
	EnabledIpv6          *bool                                               `json:"EnabledIpv6,omitempty" xml:"EnabledIpv6,omitempty"`
	Ipv6CidrBlock        *string                                             `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	Ipv6CidrBlocks       *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocks      `json:"Ipv6CidrBlocks,omitempty" xml:"Ipv6CidrBlocks,omitempty" type:"Struct"`
	IsDefault            *bool                                               `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	NatGatewayIds        *DescribeVpcsResponseBodyVpcsVpcNatGatewayIds       `json:"NatGatewayIds,omitempty" xml:"NatGatewayIds,omitempty" type:"Struct"`
	OwnerId              *int64                                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RouterTableIds       *DescribeVpcsResponseBodyVpcsVpcRouterTableIds      `json:"RouterTableIds,omitempty" xml:"RouterTableIds,omitempty" type:"Struct"`
	SecondaryCidrBlocks  *DescribeVpcsResponseBodyVpcsVpcSecondaryCidrBlocks `json:"SecondaryCidrBlocks,omitempty" xml:"SecondaryCidrBlocks,omitempty" type:"Struct"`
	Status               *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                 *DescribeVpcsResponseBodyVpcsVpcTags                `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UserCidrs            *DescribeVpcsResponseBodyVpcsVpcUserCidrs           `json:"UserCidrs,omitempty" xml:"UserCidrs,omitempty" type:"Struct"`
	VRouterId            *string                                             `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
	VSwitchIds           *DescribeVpcsResponseBodyVpcsVpcVSwitchIds          `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	VpcId                *string                                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName              *string                                             `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcsResponseBodyVpcsVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetCenStatus() *string {
	return s.CenStatus
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetDescription() *string {
	return s.Description
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetDhcpOptionsSetStatus() *string {
	return s.DhcpOptionsSetStatus
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetDnsHostnameStatus() *string {
	return s.DnsHostnameStatus
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetEnabledIpv6() *bool {
	return s.EnabledIpv6
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetIpv6CidrBlocks() *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocks {
	return s.Ipv6CidrBlocks
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetNatGatewayIds() *DescribeVpcsResponseBodyVpcsVpcNatGatewayIds {
	return s.NatGatewayIds
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetRouterTableIds() *DescribeVpcsResponseBodyVpcsVpcRouterTableIds {
	return s.RouterTableIds
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetSecondaryCidrBlocks() *DescribeVpcsResponseBodyVpcsVpcSecondaryCidrBlocks {
	return s.SecondaryCidrBlocks
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetTags() *DescribeVpcsResponseBodyVpcsVpcTags {
	return s.Tags
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetUserCidrs() *DescribeVpcsResponseBodyVpcsVpcUserCidrs {
	return s.UserCidrs
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetVRouterId() *string {
	return s.VRouterId
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetVSwitchIds() *DescribeVpcsResponseBodyVpcsVpcVSwitchIds {
	return s.VSwitchIds
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetCenStatus(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.CenStatus = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetCidrBlock(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetCreationTime(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.CreationTime = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetDescription(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.Description = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetDhcpOptionsSetId(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetDhcpOptionsSetStatus(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.DhcpOptionsSetStatus = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetDnsHostnameStatus(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.DnsHostnameStatus = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetEnabledIpv6(v bool) *DescribeVpcsResponseBodyVpcsVpc {
	s.EnabledIpv6 = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetIpv6CidrBlock(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetIpv6CidrBlocks(v *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocks) *DescribeVpcsResponseBodyVpcsVpc {
	s.Ipv6CidrBlocks = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetIsDefault(v bool) *DescribeVpcsResponseBodyVpcsVpc {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetNatGatewayIds(v *DescribeVpcsResponseBodyVpcsVpcNatGatewayIds) *DescribeVpcsResponseBodyVpcsVpc {
	s.NatGatewayIds = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetOwnerId(v int64) *DescribeVpcsResponseBodyVpcsVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetRegionId(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetResourceGroupId(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetRouterTableIds(v *DescribeVpcsResponseBodyVpcsVpcRouterTableIds) *DescribeVpcsResponseBodyVpcsVpc {
	s.RouterTableIds = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetSecondaryCidrBlocks(v *DescribeVpcsResponseBodyVpcsVpcSecondaryCidrBlocks) *DescribeVpcsResponseBodyVpcsVpc {
	s.SecondaryCidrBlocks = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetStatus(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.Status = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetTags(v *DescribeVpcsResponseBodyVpcsVpcTags) *DescribeVpcsResponseBodyVpcsVpc {
	s.Tags = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetUserCidrs(v *DescribeVpcsResponseBodyVpcsVpcUserCidrs) *DescribeVpcsResponseBodyVpcsVpc {
	s.UserCidrs = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetVRouterId(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.VRouterId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetVSwitchIds(v *DescribeVpcsResponseBodyVpcsVpcVSwitchIds) *DescribeVpcsResponseBodyVpcsVpc {
	s.VSwitchIds = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetVpcId(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetVpcName(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) Validate() error {
	if s.Ipv6CidrBlocks != nil {
		if err := s.Ipv6CidrBlocks.Validate(); err != nil {
			return err
		}
	}
	if s.NatGatewayIds != nil {
		if err := s.NatGatewayIds.Validate(); err != nil {
			return err
		}
	}
	if s.RouterTableIds != nil {
		if err := s.RouterTableIds.Validate(); err != nil {
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

type DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocks struct {
	Ipv6CidrBlock []*DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocks) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocks) GetIpv6CidrBlock() []*DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock {
	return s.Ipv6CidrBlock
}

func (s *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocks) SetIpv6CidrBlock(v []*DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock) *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocks {
	s.Ipv6CidrBlock = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocks) Validate() error {
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

type DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock struct {
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	Ipv6Isp       *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
}

func (s DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock) SetIpv6CidrBlock(v string) *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock) SetIpv6Isp(v string) *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock {
	s.Ipv6Isp = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcIpv6CidrBlocksIpv6CidrBlock) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcsResponseBodyVpcsVpcNatGatewayIds struct {
	NatGatewayIds []*string `json:"NatGatewayIds,omitempty" xml:"NatGatewayIds,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcsVpcNatGatewayIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcNatGatewayIds) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcNatGatewayIds) GetNatGatewayIds() []*string {
	return s.NatGatewayIds
}

func (s *DescribeVpcsResponseBodyVpcsVpcNatGatewayIds) SetNatGatewayIds(v []*string) *DescribeVpcsResponseBodyVpcsVpcNatGatewayIds {
	s.NatGatewayIds = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcNatGatewayIds) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcsResponseBodyVpcsVpcRouterTableIds struct {
	RouterTableIds []*string `json:"RouterTableIds,omitempty" xml:"RouterTableIds,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcsVpcRouterTableIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcRouterTableIds) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcRouterTableIds) GetRouterTableIds() []*string {
	return s.RouterTableIds
}

func (s *DescribeVpcsResponseBodyVpcsVpcRouterTableIds) SetRouterTableIds(v []*string) *DescribeVpcsResponseBodyVpcsVpcRouterTableIds {
	s.RouterTableIds = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcRouterTableIds) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcsResponseBodyVpcsVpcSecondaryCidrBlocks struct {
	SecondaryCidrBlock []*string `json:"SecondaryCidrBlock,omitempty" xml:"SecondaryCidrBlock,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcsVpcSecondaryCidrBlocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcSecondaryCidrBlocks) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcSecondaryCidrBlocks) GetSecondaryCidrBlock() []*string {
	return s.SecondaryCidrBlock
}

func (s *DescribeVpcsResponseBodyVpcsVpcSecondaryCidrBlocks) SetSecondaryCidrBlock(v []*string) *DescribeVpcsResponseBodyVpcsVpcSecondaryCidrBlocks {
	s.SecondaryCidrBlock = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcSecondaryCidrBlocks) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcsResponseBodyVpcsVpcTags struct {
	Tag []*DescribeVpcsResponseBodyVpcsVpcTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcsVpcTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcTags) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcTags) GetTag() []*DescribeVpcsResponseBodyVpcsVpcTagsTag {
	return s.Tag
}

func (s *DescribeVpcsResponseBodyVpcsVpcTags) SetTag(v []*DescribeVpcsResponseBodyVpcsVpcTagsTag) *DescribeVpcsResponseBodyVpcsVpcTags {
	s.Tag = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcTags) Validate() error {
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

type DescribeVpcsResponseBodyVpcsVpcTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpcsResponseBodyVpcsVpcTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVpcsResponseBodyVpcsVpcTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVpcsResponseBodyVpcsVpcTagsTag) SetKey(v string) *DescribeVpcsResponseBodyVpcsVpcTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcTagsTag) SetValue(v string) *DescribeVpcsResponseBodyVpcsVpcTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcsResponseBodyVpcsVpcUserCidrs struct {
	UserCidr []*string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcsVpcUserCidrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcUserCidrs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcUserCidrs) GetUserCidr() []*string {
	return s.UserCidr
}

func (s *DescribeVpcsResponseBodyVpcsVpcUserCidrs) SetUserCidr(v []*string) *DescribeVpcsResponseBodyVpcsVpcUserCidrs {
	s.UserCidr = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcUserCidrs) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcsResponseBodyVpcsVpcVSwitchIds struct {
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcsVpcVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcVSwitchIds) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchIds) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchIds) SetVSwitchId(v []*string) *DescribeVpcsResponseBodyVpcsVpcVSwitchIds {
	s.VSwitchId = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchIds) Validate() error {
	return dara.Validate(s)
}
