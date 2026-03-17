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
	// The page number of the returned page.
	//
	// example:
	//
	// 1
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
	// 3F2A0B80-D6D1-4764-8D77-38067DBBA345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the CCN instances.
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
	AssociatedCenId         *string                                                                              `json:"AssociatedCenId,omitempty" xml:"AssociatedCenId,omitempty"`
	AssociatedCenOwnerId    *string                                                                              `json:"AssociatedCenOwnerId,omitempty" xml:"AssociatedCenOwnerId,omitempty"`
	AssociatedCloudBoxCount *string                                                                              `json:"AssociatedCloudBoxCount,omitempty" xml:"AssociatedCloudBoxCount,omitempty"`
	AvailableCloudBoxCount  *string                                                                              `json:"AvailableCloudBoxCount,omitempty" xml:"AvailableCloudBoxCount,omitempty"`
	CcnId                   *string                                                                              `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	CidrBlock               *string                                                                              `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	CreateTime              *int64                                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description             *string                                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	InterworkingStatus      *string                                                                              `json:"InterworkingStatus,omitempty" xml:"InterworkingStatus,omitempty"`
	Name                    *string                                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	NewAgw                  *bool                                                                                `json:"NewAgw,omitempty" xml:"NewAgw,omitempty"`
	ResourceGroupId         *string                                                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SnatCidrBlock           *string                                                                              `json:"SnatCidrBlock,omitempty" xml:"SnatCidrBlock,omitempty"`
	Subnet                  *string                                                                              `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
	Tags                    *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetworkTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
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

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetAssociatedCloudBoxCount() *string {
	return s.AssociatedCloudBoxCount
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetAvailableCloudBoxCount() *string {
	return s.AvailableCloudBoxCount
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) GetCcnId() *string {
	return s.CcnId
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

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetAssociatedCenId(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.AssociatedCenId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetAssociatedCenOwnerId(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.AssociatedCenOwnerId = &v
	return s
}

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) SetAssociatedCloudBoxCount(v string) *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork {
	s.AssociatedCloudBoxCount = &v
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

func (s *DescribeCloudConnectNetworksResponseBodyCloudConnectNetworksCloudConnectNetwork) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
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
