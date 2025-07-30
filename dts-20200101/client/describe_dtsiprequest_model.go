// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDTSIPRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationEndpointRegion(v string) *DescribeDTSIPRequest
	GetDestinationEndpointRegion() *string
	SetRegionId(v string) *DescribeDTSIPRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDTSIPRequest
	GetResourceGroupId() *string
	SetSourceEndpointRegion(v string) *DescribeDTSIPRequest
	GetSourceEndpointRegion() *string
}

type DescribeDTSIPRequest struct {
	// The ID of the region where the destination instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// >  If the destination instance is a self-managed database with a public IP address, you can set the parameter to **cn-hangzhou*	- or the ID of the closest region.
	//
	// example:
	//
	// cn-hangzhou
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	// The ID of the region where the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the region where the source instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// >  If the source instance is a self-managed database with a public IP address, you can set the parameter to **cn-hangzhou*	- or the ID of the closest region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
}

func (s DescribeDTSIPRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDTSIPRequest) GoString() string {
	return s.String()
}

func (s *DescribeDTSIPRequest) GetDestinationEndpointRegion() *string {
	return s.DestinationEndpointRegion
}

func (s *DescribeDTSIPRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDTSIPRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDTSIPRequest) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *DescribeDTSIPRequest) SetDestinationEndpointRegion(v string) *DescribeDTSIPRequest {
	s.DestinationEndpointRegion = &v
	return s
}

func (s *DescribeDTSIPRequest) SetRegionId(v string) *DescribeDTSIPRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDTSIPRequest) SetResourceGroupId(v string) *DescribeDTSIPRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDTSIPRequest) SetSourceEndpointRegion(v string) *DescribeDTSIPRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *DescribeDTSIPRequest) Validate() error {
	return dara.Validate(s)
}
