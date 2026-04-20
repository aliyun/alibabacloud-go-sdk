// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGrantInstanceToTransitRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ModifyGrantInstanceToTransitRouterRequest
	GetCenId() *string
	SetClientToken(v string) *ModifyGrantInstanceToTransitRouterRequest
	GetClientToken() *string
	SetInstanceId(v string) *ModifyGrantInstanceToTransitRouterRequest
	GetInstanceId() *string
	SetOrderType(v string) *ModifyGrantInstanceToTransitRouterRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *ModifyGrantInstanceToTransitRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyGrantInstanceToTransitRouterRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyGrantInstanceToTransitRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyGrantInstanceToTransitRouterRequest
	GetResourceOwnerId() *int64
}

type ModifyGrantInstanceToTransitRouterRequest struct {
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// cen-dc4vwznpwbobrl****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PayByCenOwner
	OrderType            *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyGrantInstanceToTransitRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGrantInstanceToTransitRouterRequest) GoString() string {
	return s.String()
}

func (s *ModifyGrantInstanceToTransitRouterRequest) GetCenId() *string {
	return s.CenId
}

func (s *ModifyGrantInstanceToTransitRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyGrantInstanceToTransitRouterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyGrantInstanceToTransitRouterRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ModifyGrantInstanceToTransitRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyGrantInstanceToTransitRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyGrantInstanceToTransitRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyGrantInstanceToTransitRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyGrantInstanceToTransitRouterRequest) SetCenId(v string) *ModifyGrantInstanceToTransitRouterRequest {
	s.CenId = &v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterRequest) SetClientToken(v string) *ModifyGrantInstanceToTransitRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterRequest) SetInstanceId(v string) *ModifyGrantInstanceToTransitRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterRequest) SetOrderType(v string) *ModifyGrantInstanceToTransitRouterRequest {
	s.OrderType = &v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterRequest) SetOwnerAccount(v string) *ModifyGrantInstanceToTransitRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterRequest) SetOwnerId(v int64) *ModifyGrantInstanceToTransitRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterRequest) SetResourceOwnerAccount(v string) *ModifyGrantInstanceToTransitRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterRequest) SetResourceOwnerId(v int64) *ModifyGrantInstanceToTransitRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyGrantInstanceToTransitRouterRequest) Validate() error {
	return dara.Validate(s)
}
