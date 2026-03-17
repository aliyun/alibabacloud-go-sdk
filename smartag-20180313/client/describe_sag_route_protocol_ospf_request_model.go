// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRouteProtocolOspfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagRouteProtocolOspfRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagRouteProtocolOspfRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagRouteProtocolOspfRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagRouteProtocolOspfRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagRouteProtocolOspfRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagRouteProtocolOspfRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagRouteProtocolOspfRequest
	GetSmartAGSn() *string
}

type DescribeSagRouteProtocolOspfRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
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
	// sag-whfn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s DescribeSagRouteProtocolOspfRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteProtocolOspfRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteProtocolOspfRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagRouteProtocolOspfRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagRouteProtocolOspfRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagRouteProtocolOspfRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagRouteProtocolOspfRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagRouteProtocolOspfRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagRouteProtocolOspfRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagRouteProtocolOspfRequest) SetOwnerAccount(v string) *DescribeSagRouteProtocolOspfRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfRequest) SetOwnerId(v int64) *DescribeSagRouteProtocolOspfRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfRequest) SetRegionId(v string) *DescribeSagRouteProtocolOspfRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfRequest) SetResourceOwnerAccount(v string) *DescribeSagRouteProtocolOspfRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfRequest) SetResourceOwnerId(v int64) *DescribeSagRouteProtocolOspfRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfRequest) SetSmartAGId(v string) *DescribeSagRouteProtocolOspfRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfRequest) SetSmartAGSn(v string) *DescribeSagRouteProtocolOspfRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagRouteProtocolOspfRequest) Validate() error {
	return dara.Validate(s)
}
