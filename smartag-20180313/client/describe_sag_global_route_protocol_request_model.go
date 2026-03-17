// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagGlobalRouteProtocolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagGlobalRouteProtocolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagGlobalRouteProtocolRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagGlobalRouteProtocolRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagGlobalRouteProtocolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagGlobalRouteProtocolRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagGlobalRouteProtocolRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagGlobalRouteProtocolRequest
	GetSmartAGSn() *string
}

type DescribeSagGlobalRouteProtocolRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s DescribeSagGlobalRouteProtocolRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagGlobalRouteProtocolRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagGlobalRouteProtocolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagGlobalRouteProtocolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagGlobalRouteProtocolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagGlobalRouteProtocolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagGlobalRouteProtocolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagGlobalRouteProtocolRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagGlobalRouteProtocolRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagGlobalRouteProtocolRequest) SetOwnerAccount(v string) *DescribeSagGlobalRouteProtocolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolRequest) SetOwnerId(v int64) *DescribeSagGlobalRouteProtocolRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolRequest) SetRegionId(v string) *DescribeSagGlobalRouteProtocolRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolRequest) SetResourceOwnerAccount(v string) *DescribeSagGlobalRouteProtocolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolRequest) SetResourceOwnerId(v int64) *DescribeSagGlobalRouteProtocolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolRequest) SetSmartAGId(v string) *DescribeSagGlobalRouteProtocolRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolRequest) SetSmartAGSn(v string) *DescribeSagGlobalRouteProtocolRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolRequest) Validate() error {
	return dara.Validate(s)
}
