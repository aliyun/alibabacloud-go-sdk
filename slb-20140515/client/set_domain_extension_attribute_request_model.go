// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainExtensionAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainExtensionId(v string) *SetDomainExtensionAttributeRequest
	GetDomainExtensionId() *string
	SetOwnerAccount(v string) *SetDomainExtensionAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetDomainExtensionAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetDomainExtensionAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetDomainExtensionAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetDomainExtensionAttributeRequest
	GetResourceOwnerId() *int64
	SetServerCertificateId(v string) *SetDomainExtensionAttributeRequest
	GetServerCertificateId() *string
}

type SetDomainExtensionAttributeRequest struct {
	// The ID of the domain name that is associated with the additional certificate to be replaced.
	//
	// This parameter is required.
	//
	// example:
	//
	// de-bp1rp7ta*****
	DomainExtensionId *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Server Load Balancer (SLB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the new certificate.
	//
	// example:
	//
	// 1231579xxxxxxxx_166f8204689_1714763408_709981xxx
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s SetDomainExtensionAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDomainExtensionAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetDomainExtensionAttributeRequest) GetDomainExtensionId() *string {
	return s.DomainExtensionId
}

func (s *SetDomainExtensionAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetDomainExtensionAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetDomainExtensionAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDomainExtensionAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetDomainExtensionAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetDomainExtensionAttributeRequest) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *SetDomainExtensionAttributeRequest) SetDomainExtensionId(v string) *SetDomainExtensionAttributeRequest {
	s.DomainExtensionId = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetOwnerAccount(v string) *SetDomainExtensionAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetOwnerId(v int64) *SetDomainExtensionAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetRegionId(v string) *SetDomainExtensionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetResourceOwnerAccount(v string) *SetDomainExtensionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetResourceOwnerId(v int64) *SetDomainExtensionAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) SetServerCertificateId(v string) *SetDomainExtensionAttributeRequest {
	s.ServerCertificateId = &v
	return s
}

func (s *SetDomainExtensionAttributeRequest) Validate() error {
	return dara.Validate(s)
}
