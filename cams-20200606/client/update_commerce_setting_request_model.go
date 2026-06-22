// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommerceSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCartEnable(v bool) *UpdateCommerceSettingRequest
	GetCartEnable() *bool
	SetCatalogVisible(v bool) *UpdateCommerceSettingRequest
	GetCatalogVisible() *bool
	SetCustSpaceId(v string) *UpdateCommerceSettingRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *UpdateCommerceSettingRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *UpdateCommerceSettingRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *UpdateCommerceSettingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateCommerceSettingRequest
	GetResourceOwnerId() *int64
}

type UpdateCommerceSettingRequest struct {
	// Specifies whether to enable the shopping cart. Valid values:
	//
	// - true: Enable the shopping cart.
	//
	// - false: Disable the shopping cart.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	CartEnable *bool `json:"CartEnable,omitempty" xml:"CartEnable,omitempty"`
	// Specifies whether to enable the product catalog. Valid values:
	//
	// - true: Enable the product catalog.
	//
	// - false: Disable the product catalog.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	CatalogVisible *bool `json:"CatalogVisible,omitempty" xml:"CatalogVisible,omitempty"`
	// The Space ID of the Independent Software Vendor (ISV) sub-customer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 29348393884****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1380000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateCommerceSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommerceSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateCommerceSettingRequest) GetCartEnable() *bool {
	return s.CartEnable
}

func (s *UpdateCommerceSettingRequest) GetCatalogVisible() *bool {
	return s.CatalogVisible
}

func (s *UpdateCommerceSettingRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdateCommerceSettingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateCommerceSettingRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdateCommerceSettingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateCommerceSettingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateCommerceSettingRequest) SetCartEnable(v bool) *UpdateCommerceSettingRequest {
	s.CartEnable = &v
	return s
}

func (s *UpdateCommerceSettingRequest) SetCatalogVisible(v bool) *UpdateCommerceSettingRequest {
	s.CatalogVisible = &v
	return s
}

func (s *UpdateCommerceSettingRequest) SetCustSpaceId(v string) *UpdateCommerceSettingRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateCommerceSettingRequest) SetOwnerId(v int64) *UpdateCommerceSettingRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCommerceSettingRequest) SetPhoneNumber(v string) *UpdateCommerceSettingRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateCommerceSettingRequest) SetResourceOwnerAccount(v string) *UpdateCommerceSettingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateCommerceSettingRequest) SetResourceOwnerId(v int64) *UpdateCommerceSettingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateCommerceSettingRequest) Validate() error {
	return dara.Validate(s)
}
