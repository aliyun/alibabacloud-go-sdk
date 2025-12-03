// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainExtensionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainExtensionId(v string) *DeleteDomainExtensionRequest
	GetDomainExtensionId() *string
	SetOwnerAccount(v string) *DeleteDomainExtensionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDomainExtensionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteDomainExtensionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteDomainExtensionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDomainExtensionRequest
	GetResourceOwnerId() *int64
}

type DeleteDomainExtensionRequest struct {
	// The ID of the additional domain name that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// de-bp1rp7ta1****
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
}

func (s DeleteDomainExtensionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainExtensionRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainExtensionRequest) GetDomainExtensionId() *string {
	return s.DomainExtensionId
}

func (s *DeleteDomainExtensionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDomainExtensionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDomainExtensionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDomainExtensionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDomainExtensionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDomainExtensionRequest) SetDomainExtensionId(v string) *DeleteDomainExtensionRequest {
	s.DomainExtensionId = &v
	return s
}

func (s *DeleteDomainExtensionRequest) SetOwnerAccount(v string) *DeleteDomainExtensionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDomainExtensionRequest) SetOwnerId(v int64) *DeleteDomainExtensionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDomainExtensionRequest) SetRegionId(v string) *DeleteDomainExtensionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDomainExtensionRequest) SetResourceOwnerAccount(v string) *DeleteDomainExtensionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDomainExtensionRequest) SetResourceOwnerId(v int64) *DeleteDomainExtensionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDomainExtensionRequest) Validate() error {
	return dara.Validate(s)
}
