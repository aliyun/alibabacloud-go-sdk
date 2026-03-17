// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWan4GRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagWan4GRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagWan4GRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagWan4GRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagWan4GRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagWan4GRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagWan4GRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagWan4GRequest
	GetSmartAGSn() *string
}

type DescribeSagWan4GRequest struct {
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
	// The SAG instance ID.
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

func (s DescribeSagWan4GRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWan4GRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagWan4GRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagWan4GRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagWan4GRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagWan4GRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagWan4GRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagWan4GRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagWan4GRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagWan4GRequest) SetOwnerAccount(v string) *DescribeSagWan4GRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagWan4GRequest) SetOwnerId(v int64) *DescribeSagWan4GRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagWan4GRequest) SetRegionId(v string) *DescribeSagWan4GRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagWan4GRequest) SetResourceOwnerAccount(v string) *DescribeSagWan4GRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagWan4GRequest) SetResourceOwnerId(v int64) *DescribeSagWan4GRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagWan4GRequest) SetSmartAGId(v string) *DescribeSagWan4GRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagWan4GRequest) SetSmartAGSn(v string) *DescribeSagWan4GRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagWan4GRequest) Validate() error {
	return dara.Validate(s)
}
