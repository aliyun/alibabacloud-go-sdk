// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouteDistributionStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestCidrBlock(v string) *ModifyRouteDistributionStrategyRequest
	GetDestCidrBlock() *string
	SetHcInstanceId(v string) *ModifyRouteDistributionStrategyRequest
	GetHcInstanceId() *string
	SetOwnerAccount(v string) *ModifyRouteDistributionStrategyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyRouteDistributionStrategyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyRouteDistributionStrategyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyRouteDistributionStrategyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyRouteDistributionStrategyRequest
	GetResourceOwnerId() *int64
	SetRouteDistribution(v string) *ModifyRouteDistributionStrategyRequest
	GetRouteDistribution() *string
	SetRouteSource(v string) *ModifyRouteDistributionStrategyRequest
	GetRouteSource() *string
	SetSmartAGId(v string) *ModifyRouteDistributionStrategyRequest
	GetSmartAGId() *string
	SetSourceType(v string) *ModifyRouteDistributionStrategyRequest
	GetSourceType() *string
}

type ModifyRouteDistributionStrategyRequest struct {
	// The destination CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.XX.XX.0/24
	DestCidrBlock *string `json:"DestCidrBlock,omitempty" xml:"DestCidrBlock,omitempty"`
	// The ID of the health check instance.
	//
	// example:
	//
	// hc-sztovuprqzgm50****
	HcInstanceId *string `json:"HcInstanceId,omitempty" xml:"HcInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Smart Access Gateway (SAG) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route advertisement policy. Valid values:
	//
	// 	- **publish**: advertises routes.
	//
	// 	- **no_publish**: does not advertise routes.
	//
	// 	- **no_publish_and_publish_on_health_success**: routes are advertised only when they pass the health check.
	//
	// 	- **no_publish_and_publish_on_health_fail**: routes are advertised only when they fail the health check.
	//
	// 	- **publish_and_revoke_on_health_success**: advertised routes are withdrawn only when they pass the health check.
	//
	// 	- **publish_and_revoke_on_health_fail**: advertised routes are withdrawn only when they fail the health check.
	//
	// For more information, see [Configure health checks](https://help.aliyun.com/document_detail/163971.html) and [Advertise routes](https://help.aliyun.com/document_detail/163973.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// publish
	RouteDistribution *string `json:"RouteDistribution,omitempty" xml:"RouteDistribution,omitempty"`
	// The source of routes. Valid values:
	//
	// 	- **Alibaba Cloud-facing routes**
	//
	//     	- **The ID of the Virtual Private Cloud (VPC)**: learns routes from the VPC network.
	//
	//     	- **The ID of the virtual border router (VBR)**: learns routes from the VBR.
	//
	//     	- **The ID of the SAG instance**: learns routes from the SAG instance.
	//
	// 	- **Private network-facing routes**
	//
	//     	- **STATIC**: static routes specified in the SAG console.
	//
	//     	- **OSPF**: learns routes through the Open Shortest Path First (OSPF) protocol.
	//
	//     	- **BGP**: learns routes through Border Gateway Protocol (BGP).
	//
	// This parameter is required.
	//
	// example:
	//
	// OSPF
	RouteSource *string `json:"RouteSource,omitempty" xml:"RouteSource,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-erx3qta5xg5zyq****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The type of the route source. Valid values:
	//
	// 	- **cloud**: Alibaba Cloud-facing routes.
	//
	// 	- **local**: private network-facing routes.
	//
	// This parameter is required.
	//
	// example:
	//
	// local
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ModifyRouteDistributionStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouteDistributionStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyRouteDistributionStrategyRequest) GetDestCidrBlock() *string {
	return s.DestCidrBlock
}

func (s *ModifyRouteDistributionStrategyRequest) GetHcInstanceId() *string {
	return s.HcInstanceId
}

func (s *ModifyRouteDistributionStrategyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyRouteDistributionStrategyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyRouteDistributionStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRouteDistributionStrategyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyRouteDistributionStrategyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyRouteDistributionStrategyRequest) GetRouteDistribution() *string {
	return s.RouteDistribution
}

func (s *ModifyRouteDistributionStrategyRequest) GetRouteSource() *string {
	return s.RouteSource
}

func (s *ModifyRouteDistributionStrategyRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifyRouteDistributionStrategyRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ModifyRouteDistributionStrategyRequest) SetDestCidrBlock(v string) *ModifyRouteDistributionStrategyRequest {
	s.DestCidrBlock = &v
	return s
}

func (s *ModifyRouteDistributionStrategyRequest) SetHcInstanceId(v string) *ModifyRouteDistributionStrategyRequest {
	s.HcInstanceId = &v
	return s
}

func (s *ModifyRouteDistributionStrategyRequest) SetOwnerAccount(v string) *ModifyRouteDistributionStrategyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyRouteDistributionStrategyRequest) SetOwnerId(v int64) *ModifyRouteDistributionStrategyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRouteDistributionStrategyRequest) SetRegionId(v string) *ModifyRouteDistributionStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRouteDistributionStrategyRequest) SetResourceOwnerAccount(v string) *ModifyRouteDistributionStrategyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRouteDistributionStrategyRequest) SetResourceOwnerId(v int64) *ModifyRouteDistributionStrategyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRouteDistributionStrategyRequest) SetRouteDistribution(v string) *ModifyRouteDistributionStrategyRequest {
	s.RouteDistribution = &v
	return s
}

func (s *ModifyRouteDistributionStrategyRequest) SetRouteSource(v string) *ModifyRouteDistributionStrategyRequest {
	s.RouteSource = &v
	return s
}

func (s *ModifyRouteDistributionStrategyRequest) SetSmartAGId(v string) *ModifyRouteDistributionStrategyRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifyRouteDistributionStrategyRequest) SetSourceType(v string) *ModifyRouteDistributionStrategyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyRouteDistributionStrategyRequest) Validate() error {
	return dara.Validate(s)
}
