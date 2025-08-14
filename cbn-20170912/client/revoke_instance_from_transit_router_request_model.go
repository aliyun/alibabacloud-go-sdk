// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromTransitRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *RevokeInstanceFromTransitRouterRequest
	GetCenId() *string
	SetCenOwnerId(v int64) *RevokeInstanceFromTransitRouterRequest
	GetCenOwnerId() *int64
	SetInstanceId(v string) *RevokeInstanceFromTransitRouterRequest
	GetInstanceId() *string
	SetInstanceType(v string) *RevokeInstanceFromTransitRouterRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *RevokeInstanceFromTransitRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RevokeInstanceFromTransitRouterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RevokeInstanceFromTransitRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RevokeInstanceFromTransitRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RevokeInstanceFromTransitRouterRequest
	GetResourceOwnerId() *int64
}

type RevokeInstanceFromTransitRouterRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
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
	// The network instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the network instance. Valid values:
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
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the network instance.
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

func (s RevokeInstanceFromTransitRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromTransitRouterRequest) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromTransitRouterRequest) GetCenId() *string {
	return s.CenId
}

func (s *RevokeInstanceFromTransitRouterRequest) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *RevokeInstanceFromTransitRouterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeInstanceFromTransitRouterRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RevokeInstanceFromTransitRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RevokeInstanceFromTransitRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RevokeInstanceFromTransitRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeInstanceFromTransitRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RevokeInstanceFromTransitRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RevokeInstanceFromTransitRouterRequest) SetCenId(v string) *RevokeInstanceFromTransitRouterRequest {
	s.CenId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetCenOwnerId(v int64) *RevokeInstanceFromTransitRouterRequest {
	s.CenOwnerId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetInstanceId(v string) *RevokeInstanceFromTransitRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetInstanceType(v string) *RevokeInstanceFromTransitRouterRequest {
	s.InstanceType = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetOwnerAccount(v string) *RevokeInstanceFromTransitRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetOwnerId(v int64) *RevokeInstanceFromTransitRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetRegionId(v string) *RevokeInstanceFromTransitRouterRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetResourceOwnerAccount(v string) *RevokeInstanceFromTransitRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) SetResourceOwnerId(v int64) *RevokeInstanceFromTransitRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterRequest) Validate() error {
	return dara.Validate(s)
}
