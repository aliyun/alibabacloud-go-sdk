// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteDistributionStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestCidrBlock(v string) *DeleteRouteDistributionStrategyRequest
	GetDestCidrBlock() *string
	SetOwnerAccount(v string) *DeleteRouteDistributionStrategyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteRouteDistributionStrategyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteRouteDistributionStrategyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteRouteDistributionStrategyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRouteDistributionStrategyRequest
	GetResourceOwnerId() *int64
	SetRouteSource(v string) *DeleteRouteDistributionStrategyRequest
	GetRouteSource() *string
	SetSmartAGId(v string) *DeleteRouteDistributionStrategyRequest
	GetSmartAGId() *string
	SetSourceType(v string) *DeleteRouteDistributionStrategyRequest
	GetSourceType() *string
}

type DeleteRouteDistributionStrategyRequest struct {
	// The destination CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.0/24
	DestCidrBlock *string `json:"DestCidrBlock,omitempty" xml:"DestCidrBlock,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Smart Access Gateway (SAG) instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source of routes. Valid values:
	//
	// 	- **Alibaba Cloud**
	//
	//     	- **Virtual private cloud (VPC) IDs**: Routes that are learned from VPCs.
	//
	//     	- **Virtual border router (VBR) IDs**: Routes that are learned from VBRs.
	//
	//     	- **SAG instance IDs**: Routes that are learned from SAG instances.
	//
	// 	- **On-premises network**
	//
	//     	- **STATIC**: Static routes that are specified in the SAG console.
	//
	//     	- **OSPF**: Routes that are learned through the Open Shortest Path First (OSPF) protocol.
	//
	//     	- **BGP**: Routes that are learned through Border Gateway Protocol (BGP).
	//
	// This parameter is required.
	//
	// example:
	//
	// STATIC
	RouteSource *string `json:"RouteSource,omitempty" xml:"RouteSource,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-gyat6giidkvyk****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The type of the route source. Valid values:
	//
	// 	- **cloud**: Alibaba Cloud
	//
	// 	- **local**: on-premises network
	//
	// This parameter is required.
	//
	// example:
	//
	// local
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DeleteRouteDistributionStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteDistributionStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteDistributionStrategyRequest) GetDestCidrBlock() *string {
	return s.DestCidrBlock
}

func (s *DeleteRouteDistributionStrategyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteRouteDistributionStrategyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRouteDistributionStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRouteDistributionStrategyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRouteDistributionStrategyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRouteDistributionStrategyRequest) GetRouteSource() *string {
	return s.RouteSource
}

func (s *DeleteRouteDistributionStrategyRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DeleteRouteDistributionStrategyRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DeleteRouteDistributionStrategyRequest) SetDestCidrBlock(v string) *DeleteRouteDistributionStrategyRequest {
	s.DestCidrBlock = &v
	return s
}

func (s *DeleteRouteDistributionStrategyRequest) SetOwnerAccount(v string) *DeleteRouteDistributionStrategyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRouteDistributionStrategyRequest) SetOwnerId(v int64) *DeleteRouteDistributionStrategyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRouteDistributionStrategyRequest) SetRegionId(v string) *DeleteRouteDistributionStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRouteDistributionStrategyRequest) SetResourceOwnerAccount(v string) *DeleteRouteDistributionStrategyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRouteDistributionStrategyRequest) SetResourceOwnerId(v int64) *DeleteRouteDistributionStrategyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRouteDistributionStrategyRequest) SetRouteSource(v string) *DeleteRouteDistributionStrategyRequest {
	s.RouteSource = &v
	return s
}

func (s *DeleteRouteDistributionStrategyRequest) SetSmartAGId(v string) *DeleteRouteDistributionStrategyRequest {
	s.SmartAGId = &v
	return s
}

func (s *DeleteRouteDistributionStrategyRequest) SetSourceType(v string) *DeleteRouteDistributionStrategyRequest {
	s.SourceType = &v
	return s
}

func (s *DeleteRouteDistributionStrategyRequest) Validate() error {
	return dara.Validate(s)
}
