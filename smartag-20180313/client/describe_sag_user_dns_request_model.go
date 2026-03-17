// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagUserDnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagUserDnsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagUserDnsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagUserDnsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagUserDnsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagUserDnsRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagUserDnsRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagUserDnsRequest
	GetSmartAGSn() *string
}

type DescribeSagUserDnsRequest struct {
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

func (s DescribeSagUserDnsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagUserDnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagUserDnsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagUserDnsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagUserDnsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagUserDnsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagUserDnsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagUserDnsRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagUserDnsRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagUserDnsRequest) SetOwnerAccount(v string) *DescribeSagUserDnsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagUserDnsRequest) SetOwnerId(v int64) *DescribeSagUserDnsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagUserDnsRequest) SetRegionId(v string) *DescribeSagUserDnsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagUserDnsRequest) SetResourceOwnerAccount(v string) *DescribeSagUserDnsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagUserDnsRequest) SetResourceOwnerId(v int64) *DescribeSagUserDnsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagUserDnsRequest) SetSmartAGId(v string) *DescribeSagUserDnsRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagUserDnsRequest) SetSmartAGSn(v string) *DescribeSagUserDnsRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagUserDnsRequest) Validate() error {
	return dara.Validate(s)
}
