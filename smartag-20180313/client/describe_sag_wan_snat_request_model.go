// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWanSnatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagWanSnatRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagWanSnatRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagWanSnatRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagWanSnatRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagWanSnatRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagWanSnatRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagWanSnatRequest
	GetSmartAGSn() *string
}

type DescribeSagWanSnatRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SAG instance.
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

func (s DescribeSagWanSnatRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWanSnatRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagWanSnatRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagWanSnatRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagWanSnatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagWanSnatRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagWanSnatRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagWanSnatRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagWanSnatRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagWanSnatRequest) SetOwnerAccount(v string) *DescribeSagWanSnatRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagWanSnatRequest) SetOwnerId(v int64) *DescribeSagWanSnatRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagWanSnatRequest) SetRegionId(v string) *DescribeSagWanSnatRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagWanSnatRequest) SetResourceOwnerAccount(v string) *DescribeSagWanSnatRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagWanSnatRequest) SetResourceOwnerId(v int64) *DescribeSagWanSnatRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagWanSnatRequest) SetSmartAGId(v string) *DescribeSagWanSnatRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagWanSnatRequest) SetSmartAGSn(v string) *DescribeSagWanSnatRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagWanSnatRequest) Validate() error {
	return dara.Validate(s)
}
