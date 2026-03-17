// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearSagRouteableAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ClearSagRouteableAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ClearSagRouteableAddressRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ClearSagRouteableAddressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ClearSagRouteableAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClearSagRouteableAddressRequest
	GetResourceOwnerId() *int64
	SetSagId(v string) *ClearSagRouteableAddressRequest
	GetSagId() *string
}

type ClearSagRouteableAddressRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the SAG instance belongs.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-0phdojgu5tqr1p****
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
}

func (s ClearSagRouteableAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearSagRouteableAddressRequest) GoString() string {
	return s.String()
}

func (s *ClearSagRouteableAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ClearSagRouteableAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClearSagRouteableAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ClearSagRouteableAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClearSagRouteableAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClearSagRouteableAddressRequest) GetSagId() *string {
	return s.SagId
}

func (s *ClearSagRouteableAddressRequest) SetOwnerAccount(v string) *ClearSagRouteableAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ClearSagRouteableAddressRequest) SetOwnerId(v int64) *ClearSagRouteableAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ClearSagRouteableAddressRequest) SetRegionId(v string) *ClearSagRouteableAddressRequest {
	s.RegionId = &v
	return s
}

func (s *ClearSagRouteableAddressRequest) SetResourceOwnerAccount(v string) *ClearSagRouteableAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClearSagRouteableAddressRequest) SetResourceOwnerId(v int64) *ClearSagRouteableAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClearSagRouteableAddressRequest) SetSagId(v string) *ClearSagRouteableAddressRequest {
	s.SagId = &v
	return s
}

func (s *ClearSagRouteableAddressRequest) Validate() error {
	return dara.Validate(s)
}
