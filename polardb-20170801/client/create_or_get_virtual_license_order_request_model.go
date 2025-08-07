// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrGetVirtualLicenseOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *CreateOrGetVirtualLicenseOrderRequest
	GetEngine() *string
	SetOwnerAccount(v string) *CreateOrGetVirtualLicenseOrderRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateOrGetVirtualLicenseOrderRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateOrGetVirtualLicenseOrderRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateOrGetVirtualLicenseOrderRequest
	GetResourceOwnerId() *int64
}

type CreateOrGetVirtualLicenseOrderRequest struct {
	// The type of the engine. Valid values: PG, Oracle, and MySQL.
	//
	// This parameter is required.
	//
	// example:
	//
	// PG
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateOrGetVirtualLicenseOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrGetVirtualLicenseOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrGetVirtualLicenseOrderRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateOrGetVirtualLicenseOrderRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateOrGetVirtualLicenseOrderRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateOrGetVirtualLicenseOrderRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateOrGetVirtualLicenseOrderRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateOrGetVirtualLicenseOrderRequest) SetEngine(v string) *CreateOrGetVirtualLicenseOrderRequest {
	s.Engine = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderRequest) SetOwnerAccount(v string) *CreateOrGetVirtualLicenseOrderRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderRequest) SetOwnerId(v int64) *CreateOrGetVirtualLicenseOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderRequest) SetResourceOwnerAccount(v string) *CreateOrGetVirtualLicenseOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderRequest) SetResourceOwnerId(v int64) *CreateOrGetVirtualLicenseOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderRequest) Validate() error {
	return dara.Validate(s)
}
