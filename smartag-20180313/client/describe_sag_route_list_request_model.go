// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRouteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagRouteListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagRouteListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagRouteListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagRouteListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagRouteListRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagRouteListRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagRouteListRequest
	GetSmartAGSn() *string
}

type DescribeSagRouteListRequest struct {
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

func (s DescribeSagRouteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagRouteListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagRouteListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagRouteListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagRouteListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagRouteListRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagRouteListRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagRouteListRequest) SetOwnerAccount(v string) *DescribeSagRouteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagRouteListRequest) SetOwnerId(v int64) *DescribeSagRouteListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagRouteListRequest) SetRegionId(v string) *DescribeSagRouteListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagRouteListRequest) SetResourceOwnerAccount(v string) *DescribeSagRouteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagRouteListRequest) SetResourceOwnerId(v int64) *DescribeSagRouteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagRouteListRequest) SetSmartAGId(v string) *DescribeSagRouteListRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagRouteListRequest) SetSmartAGSn(v string) *DescribeSagRouteListRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagRouteListRequest) Validate() error {
	return dara.Validate(s)
}
