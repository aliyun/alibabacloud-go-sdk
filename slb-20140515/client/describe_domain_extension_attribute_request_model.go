// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainExtensionAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainExtensionId(v string) *DescribeDomainExtensionAttributeRequest
	GetDomainExtensionId() *string
	SetOwnerAccount(v string) *DescribeDomainExtensionAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDomainExtensionAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDomainExtensionAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDomainExtensionAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDomainExtensionAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeDomainExtensionAttributeRequest struct {
	// The ID of the additional certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// de-bp1rp7ta1****
	DomainExtensionId *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Server Load Balancer (SLB) instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDomainExtensionAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainExtensionAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionAttributeRequest) GetDomainExtensionId() *string {
	return s.DomainExtensionId
}

func (s *DescribeDomainExtensionAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDomainExtensionAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDomainExtensionAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDomainExtensionAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDomainExtensionAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDomainExtensionAttributeRequest) SetDomainExtensionId(v string) *DescribeDomainExtensionAttributeRequest {
	s.DomainExtensionId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeRequest) SetOwnerAccount(v string) *DescribeDomainExtensionAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDomainExtensionAttributeRequest) SetOwnerId(v int64) *DescribeDomainExtensionAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeRequest) SetRegionId(v string) *DescribeDomainExtensionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeRequest) SetResourceOwnerAccount(v string) *DescribeDomainExtensionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDomainExtensionAttributeRequest) SetResourceOwnerId(v int64) *DescribeDomainExtensionAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeRequest) Validate() error {
	return dara.Validate(s)
}
