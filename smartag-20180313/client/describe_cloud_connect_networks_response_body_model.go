// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudConnectNetworksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudConnectNetworks(v *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks) *DescribeCloudConnectNetworksResponseBody
	GetCloudConnectNetworks() *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks
	SetPageNumber(v int32) *DescribeCloudConnectNetworksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCloudConnectNetworksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCloudConnectNetworksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCloudConnectNetworksResponseBody
	GetTotalCount() *int32
}

type DescribeCloudConnectNetworksResponseBody struct {
	CloudConnectNetworks *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks `json:"CloudConnectNetworks,omitempty" xml:"CloudConnectNetworks,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for paging in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F2A0B80-D6D1-4764-8D77-38067DBBA345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudConnectNetworksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponseBody) GetCloudConnectNetworks() *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks {
	return s.CloudConnectNetworks
}

func (s *DescribeCloudConnectNetworksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCloudConnectNetworksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudConnectNetworksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudConnectNetworksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCloudConnectNetworksResponseBody) SetCloudConnectNetworks(v *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks) *DescribeCloudConnectNetworksResponseBody {
	s.CloudConnectNetworks = v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBody) SetPageNumber(v int32) *DescribeCloudConnectNetworksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBody) SetPageSize(v int32) *DescribeCloudConnectNetworksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBody) SetRequestId(v string) *DescribeCloudConnectNetworksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBody) SetTotalCount(v int32) *DescribeCloudConnectNetworksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBody) Validate() error {
	if s.CloudConnectNetworks != nil {
		if err := s.CloudConnectNetworks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks struct {
	CloudConnectNetwork []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork `json:"CloudConnectNetwork,omitempty" xml:"CloudConnectNetwork,omitempty" type:"Repeated"`
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks) GetCloudConnectNetwork() []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	return s.CloudConnectNetwork
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks) SetCloudConnectNetwork(v []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks {
	s.CloudConnectNetwork = v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworks) Validate() error {
	if s.CloudConnectNetwork != nil {
		for _, item := range s.CloudConnectNetwork {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork struct {
	AssociatedCenId         *string                                                                                          `json:"AssociatedCenId,omitempty" xml:"AssociatedCenId,omitempty"`
	AssociatedCenOwnerId    *string                                                                                          `json:"AssociatedCenOwnerId,omitempty" xml:"AssociatedCenOwnerId,omitempty"`
	AssociatedCenStatus     *string                                                                                          `json:"AssociatedCenStatus,omitempty" xml:"AssociatedCenStatus,omitempty"`
	AssociatedCloudBoxCount *string                                                                                          `json:"AssociatedCloudBoxCount,omitempty" xml:"AssociatedCloudBoxCount,omitempty"`
	AssociatedVpcStatus     *string                                                                                          `json:"AssociatedVpcStatus,omitempty" xml:"AssociatedVpcStatus,omitempty"`
	AvailableCloudBoxCount  *string                                                                                          `json:"AvailableCloudBoxCount,omitempty" xml:"AvailableCloudBoxCount,omitempty"`
	CcnId                   *string                                                                                          `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	CenStatusErrorCode      *string                                                                                          `json:"CenStatusErrorCode,omitempty" xml:"CenStatusErrorCode,omitempty"`
	CenStatusErrorMessage   *string                                                                                          `json:"CenStatusErrorMessage,omitempty" xml:"CenStatusErrorMessage,omitempty"`
	CidrBlock               *string                                                                                          `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	CreateTime              *int64                                                                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description             *string                                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	InterworkingStatus      *string                                                                                          `json:"InterworkingStatus,omitempty" xml:"InterworkingStatus,omitempty"`
	Name                    *string                                                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	NewAgw                  *bool                                                                                            `json:"NewAgw,omitempty" xml:"NewAgw,omitempty"`
	ProcessingCenId         *string                                                                                          `json:"ProcessingCenId,omitempty" xml:"ProcessingCenId,omitempty"`
	ResourceGroupId         *string                                                                                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SnatCidrBlock           *string                                                                                          `json:"SnatCidrBlock,omitempty" xml:"SnatCidrBlock,omitempty"`
	Subnet                  *string                                                                                          `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
	Tags                    *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTags             `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TrNetworkConfig         *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig  `json:"TrNetworkConfig,omitempty" xml:"TrNetworkConfig,omitempty" type:"Struct"`
	VpcNetworkConfig        *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig `json:"VpcNetworkConfig,omitempty" xml:"VpcNetworkConfig,omitempty" type:"Struct"`
	VpcStatusErrorCode      *string                                                                                          `json:"VpcStatusErrorCode,omitempty" xml:"VpcStatusErrorCode,omitempty"`
	VpcStatusErrorMessage   *string                                                                                          `json:"VpcStatusErrorMessage,omitempty" xml:"VpcStatusErrorMessage,omitempty"`
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetAssociatedCenId() *string {
	return s.AssociatedCenId
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetAssociatedCenOwnerId() *string {
	return s.AssociatedCenOwnerId
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetAssociatedCenStatus() *string {
	return s.AssociatedCenStatus
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetAssociatedCloudBoxCount() *string {
	return s.AssociatedCloudBoxCount
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetAssociatedVpcStatus() *string {
	return s.AssociatedVpcStatus
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetAvailableCloudBoxCount() *string {
	return s.AvailableCloudBoxCount
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetCcnId() *string {
	return s.CcnId
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetCenStatusErrorCode() *string {
	return s.CenStatusErrorCode
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetCenStatusErrorMessage() *string {
	return s.CenStatusErrorMessage
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetDescription() *string {
	return s.Description
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetInterworkingStatus() *string {
	return s.InterworkingStatus
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetName() *string {
	return s.Name
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetNewAgw() *bool {
	return s.NewAgw
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetProcessingCenId() *string {
	return s.ProcessingCenId
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetSnatCidrBlock() *string {
	return s.SnatCidrBlock
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetSubnet() *string {
	return s.Subnet
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetTags() *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTags {
	return s.Tags
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetTrNetworkConfig() *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig {
	return s.TrNetworkConfig
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetVpcNetworkConfig() *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig {
	return s.VpcNetworkConfig
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetVpcStatusErrorCode() *string {
	return s.VpcStatusErrorCode
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetVpcStatusErrorMessage() *string {
	return s.VpcStatusErrorMessage
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetAssociatedCenId(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.AssociatedCenId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetAssociatedCenOwnerId(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.AssociatedCenOwnerId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetAssociatedCenStatus(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.AssociatedCenStatus = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetAssociatedCloudBoxCount(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.AssociatedCloudBoxCount = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetAssociatedVpcStatus(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.AssociatedVpcStatus = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetAvailableCloudBoxCount(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.AvailableCloudBoxCount = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetCcnId(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.CcnId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetCenStatusErrorCode(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.CenStatusErrorCode = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetCenStatusErrorMessage(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.CenStatusErrorMessage = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetCidrBlock(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.CidrBlock = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetCreateTime(v int64) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetDescription(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.Description = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetInterworkingStatus(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.InterworkingStatus = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetName(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.Name = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetNewAgw(v bool) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.NewAgw = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetProcessingCenId(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.ProcessingCenId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetResourceGroupId(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetSnatCidrBlock(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.SnatCidrBlock = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetSubnet(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.Subnet = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetTags(v *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTags) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.Tags = v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetTrNetworkConfig(v *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.TrNetworkConfig = v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetVpcNetworkConfig(v *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.VpcNetworkConfig = v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetVpcStatusErrorCode(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.VpcStatusErrorCode = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetVpcStatusErrorMessage(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.VpcStatusErrorMessage = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.TrNetworkConfig != nil {
		if err := s.TrNetworkConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VpcNetworkConfig != nil {
		if err := s.VpcNetworkConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTags struct {
	Tag []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTags) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTags) GetTag() []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag {
	return s.Tag
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTags) SetTag(v []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTags {
	s.Tag = v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTags) Validate() error {
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

type DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag) SetKey(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag) SetValue(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig struct {
	CidrAllocationMode *string                                                                                                          `json:"CidrAllocationMode,omitempty" xml:"CidrAllocationMode,omitempty"`
	RegionBindingList  *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingList `json:"RegionBindingList,omitempty" xml:"RegionBindingList,omitempty" type:"Struct"`
	Subnet             *string                                                                                                          `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig) GetCidrAllocationMode() *string {
	return s.CidrAllocationMode
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig) GetRegionBindingList() *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingList {
	return s.RegionBindingList
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig) GetSubnet() *string {
	return s.Subnet
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig) SetCidrAllocationMode(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig {
	s.CidrAllocationMode = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig) SetRegionBindingList(v *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingList) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig {
	s.RegionBindingList = v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig) SetSubnet(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig {
	s.Subnet = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfig) Validate() error {
	if s.RegionBindingList != nil {
		if err := s.RegionBindingList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingList struct {
	RegionBinding []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding `json:"RegionBinding,omitempty" xml:"RegionBinding,omitempty" type:"Repeated"`
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingList) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingList) GetRegionBinding() []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding {
	return s.RegionBinding
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingList) SetRegionBinding(v []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingList {
	s.RegionBinding = v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingList) Validate() error {
	if s.RegionBinding != nil {
		for _, item := range s.RegionBinding {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding struct {
	CidrBlock       *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RelayCidrBlock  *string `json:"RelayCidrBlock,omitempty" xml:"RelayCidrBlock,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) GetRelayCidrBlock() *string {
	return s.RelayCidrBlock
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) SetCidrBlock(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding {
	s.CidrBlock = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) SetRegionId(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) SetRelayCidrBlock(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding {
	s.RelayCidrBlock = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) SetStatus(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding {
	s.Status = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) SetTransitRouterId(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTrNetworkConfigRegionBindingListRegionBinding) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig struct {
	CidrAllocationMode *string                                                                                                           `json:"CidrAllocationMode,omitempty" xml:"CidrAllocationMode,omitempty"`
	RegionBindingList  *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingList `json:"RegionBindingList,omitempty" xml:"RegionBindingList,omitempty" type:"Struct"`
	Subnet             *string                                                                                                           `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig) GetCidrAllocationMode() *string {
	return s.CidrAllocationMode
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig) GetRegionBindingList() *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingList {
	return s.RegionBindingList
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig) GetSubnet() *string {
	return s.Subnet
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig) SetCidrAllocationMode(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig {
	s.CidrAllocationMode = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig) SetRegionBindingList(v *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingList) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig {
	s.RegionBindingList = v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig) SetSubnet(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig {
	s.Subnet = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfig) Validate() error {
	if s.RegionBindingList != nil {
		if err := s.RegionBindingList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingList struct {
	RegionBinding []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding `json:"RegionBinding,omitempty" xml:"RegionBinding,omitempty" type:"Repeated"`
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingList) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingList) GetRegionBinding() []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding {
	return s.RegionBinding
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingList) SetRegionBinding(v []*DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingList {
	s.RegionBinding = v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingList) Validate() error {
	if s.RegionBinding != nil {
		for _, item := range s.RegionBinding {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RelayCidrBlock *string `json:"RelayCidrBlock,omitempty" xml:"RelayCidrBlock,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcCidrBlock   *string `json:"VpcCidrBlock,omitempty" xml:"VpcCidrBlock,omitempty"`
	VpcId          *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) GoString() string {
	return s.String()
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) GetRelayCidrBlock() *string {
	return s.RelayCidrBlock
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) GetVpcCidrBlock() *string {
	return s.VpcCidrBlock
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) SetRegionId(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) SetRelayCidrBlock(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding {
	s.RelayCidrBlock = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) SetStatus(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding {
	s.Status = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) SetVpcCidrBlock(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding {
	s.VpcCidrBlock = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) SetVpcId(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding {
	s.VpcId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkVpcNetworkConfigRegionBindingListRegionBinding) Validate() error {
	return dara.Validate(s)
}
