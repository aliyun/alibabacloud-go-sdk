// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToTransitRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *GrantInstanceToTransitRouterRequest
	GetCenId() *string
	SetCenOwnerId(v int64) *GrantInstanceToTransitRouterRequest
	GetCenOwnerId() *int64
	SetInstanceId(v string) *GrantInstanceToTransitRouterRequest
	GetInstanceId() *string
	SetInstanceType(v string) *GrantInstanceToTransitRouterRequest
	GetInstanceType() *string
	SetOrderType(v string) *GrantInstanceToTransitRouterRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *GrantInstanceToTransitRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GrantInstanceToTransitRouterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GrantInstanceToTransitRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GrantInstanceToTransitRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GrantInstanceToTransitRouterRequest
	GetResourceOwnerId() *int64
}

type GrantInstanceToTransitRouterRequest struct {
	// Enter the ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-44m0p68spvlrqq****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1250123456123456
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of network instance. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **ExpressConnect**: VBR
	//
	// 	- **VPN**: IPsec-VPN connection
	//
	// 	- **ECR**: ECR
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The entity that pays the fees of the network instance. Valid values:
	//
	// 	- **PayByCenOwner**: the Alibaba Cloud account that owns the CEN instance.
	//
	// 	- **PayByResourceOwner**: the Alibaba Cloud account that owns the network instance.
	//
	// example:
	//
	// PayByCenOwner
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the network instance is deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GrantInstanceToTransitRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToTransitRouterRequest) GoString() string {
	return s.String()
}

func (s *GrantInstanceToTransitRouterRequest) GetCenId() *string {
	return s.CenId
}

func (s *GrantInstanceToTransitRouterRequest) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *GrantInstanceToTransitRouterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GrantInstanceToTransitRouterRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GrantInstanceToTransitRouterRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *GrantInstanceToTransitRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GrantInstanceToTransitRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GrantInstanceToTransitRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantInstanceToTransitRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GrantInstanceToTransitRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GrantInstanceToTransitRouterRequest) SetCenId(v string) *GrantInstanceToTransitRouterRequest {
	s.CenId = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetCenOwnerId(v int64) *GrantInstanceToTransitRouterRequest {
	s.CenOwnerId = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetInstanceId(v string) *GrantInstanceToTransitRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetInstanceType(v string) *GrantInstanceToTransitRouterRequest {
	s.InstanceType = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetOrderType(v string) *GrantInstanceToTransitRouterRequest {
	s.OrderType = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetOwnerAccount(v string) *GrantInstanceToTransitRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetOwnerId(v int64) *GrantInstanceToTransitRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetRegionId(v string) *GrantInstanceToTransitRouterRequest {
	s.RegionId = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetResourceOwnerAccount(v string) *GrantInstanceToTransitRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) SetResourceOwnerId(v int64) *GrantInstanceToTransitRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantInstanceToTransitRouterRequest) Validate() error {
	return dara.Validate(s)
}
