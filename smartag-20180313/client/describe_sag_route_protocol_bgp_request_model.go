// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRouteProtocolBgpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagRouteProtocolBgpRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagRouteProtocolBgpRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagRouteProtocolBgpRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagRouteProtocolBgpRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagRouteProtocolBgpRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagRouteProtocolBgpRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagRouteProtocolBgpRequest
	GetSmartAGSn() *string
}

type DescribeSagRouteProtocolBgpRequest struct {
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
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-whfn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device that is associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s DescribeSagRouteProtocolBgpRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteProtocolBgpRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteProtocolBgpRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagRouteProtocolBgpRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagRouteProtocolBgpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagRouteProtocolBgpRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagRouteProtocolBgpRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagRouteProtocolBgpRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagRouteProtocolBgpRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagRouteProtocolBgpRequest) SetOwnerAccount(v string) *DescribeSagRouteProtocolBgpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpRequest) SetOwnerId(v int64) *DescribeSagRouteProtocolBgpRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpRequest) SetRegionId(v string) *DescribeSagRouteProtocolBgpRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpRequest) SetResourceOwnerAccount(v string) *DescribeSagRouteProtocolBgpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpRequest) SetResourceOwnerId(v int64) *DescribeSagRouteProtocolBgpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpRequest) SetSmartAGId(v string) *DescribeSagRouteProtocolBgpRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpRequest) SetSmartAGSn(v string) *DescribeSagRouteProtocolBgpRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpRequest) Validate() error {
	return dara.Validate(s)
}
