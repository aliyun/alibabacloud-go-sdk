// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagPortRouteProtocolListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagPortRouteProtocolListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagPortRouteProtocolListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagPortRouteProtocolListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagPortRouteProtocolListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagPortRouteProtocolListRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagPortRouteProtocolListRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagPortRouteProtocolListRequest
	GetSmartAGSn() *string
}

type DescribeSagPortRouteProtocolListRequest struct {
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

func (s DescribeSagPortRouteProtocolListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagPortRouteProtocolListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagPortRouteProtocolListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagPortRouteProtocolListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagPortRouteProtocolListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagPortRouteProtocolListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagPortRouteProtocolListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagPortRouteProtocolListRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagPortRouteProtocolListRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagPortRouteProtocolListRequest) SetOwnerAccount(v string) *DescribeSagPortRouteProtocolListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListRequest) SetOwnerId(v int64) *DescribeSagPortRouteProtocolListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListRequest) SetRegionId(v string) *DescribeSagPortRouteProtocolListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListRequest) SetResourceOwnerAccount(v string) *DescribeSagPortRouteProtocolListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListRequest) SetResourceOwnerId(v int64) *DescribeSagPortRouteProtocolListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListRequest) SetSmartAGId(v string) *DescribeSagPortRouteProtocolListRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListRequest) SetSmartAGSn(v string) *DescribeSagPortRouteProtocolListRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListRequest) Validate() error {
	return dara.Validate(s)
}
