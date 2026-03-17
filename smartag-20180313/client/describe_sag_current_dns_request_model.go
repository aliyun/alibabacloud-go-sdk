// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagCurrentDnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagCurrentDnsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagCurrentDnsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagCurrentDnsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagCurrentDnsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagCurrentDnsRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagCurrentDnsRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagCurrentDnsRequest
	GetSmartAGSn() *string
}

type DescribeSagCurrentDnsRequest struct {
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

func (s DescribeSagCurrentDnsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagCurrentDnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagCurrentDnsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagCurrentDnsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagCurrentDnsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagCurrentDnsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagCurrentDnsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagCurrentDnsRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagCurrentDnsRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagCurrentDnsRequest) SetOwnerAccount(v string) *DescribeSagCurrentDnsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagCurrentDnsRequest) SetOwnerId(v int64) *DescribeSagCurrentDnsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagCurrentDnsRequest) SetRegionId(v string) *DescribeSagCurrentDnsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagCurrentDnsRequest) SetResourceOwnerAccount(v string) *DescribeSagCurrentDnsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagCurrentDnsRequest) SetResourceOwnerId(v int64) *DescribeSagCurrentDnsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagCurrentDnsRequest) SetSmartAGId(v string) *DescribeSagCurrentDnsRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagCurrentDnsRequest) SetSmartAGSn(v string) *DescribeSagCurrentDnsRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagCurrentDnsRequest) Validate() error {
	return dara.Validate(s)
}
