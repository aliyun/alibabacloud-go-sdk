// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagStaticRouteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagStaticRouteListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagStaticRouteListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagStaticRouteListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagStaticRouteListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagStaticRouteListRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagStaticRouteListRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagStaticRouteListRequest
	GetSmartAGSn() *string
}

type DescribeSagStaticRouteListRequest struct {
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

func (s DescribeSagStaticRouteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagStaticRouteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagStaticRouteListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagStaticRouteListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagStaticRouteListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagStaticRouteListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagStaticRouteListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagStaticRouteListRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagStaticRouteListRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagStaticRouteListRequest) SetOwnerAccount(v string) *DescribeSagStaticRouteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagStaticRouteListRequest) SetOwnerId(v int64) *DescribeSagStaticRouteListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagStaticRouteListRequest) SetRegionId(v string) *DescribeSagStaticRouteListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagStaticRouteListRequest) SetResourceOwnerAccount(v string) *DescribeSagStaticRouteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagStaticRouteListRequest) SetResourceOwnerId(v int64) *DescribeSagStaticRouteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagStaticRouteListRequest) SetSmartAGId(v string) *DescribeSagStaticRouteListRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagStaticRouteListRequest) SetSmartAGSn(v string) *DescribeSagStaticRouteListRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagStaticRouteListRequest) Validate() error {
	return dara.Validate(s)
}
