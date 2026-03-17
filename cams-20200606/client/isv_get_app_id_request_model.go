// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvGetAppIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIntlVersion(v string) *IsvGetAppIdRequest
	GetIntlVersion() *string
	SetOwnerId(v int64) *IsvGetAppIdRequest
	GetOwnerId() *int64
	SetPermissions(v string) *IsvGetAppIdRequest
	GetPermissions() *string
	SetResourceOwnerAccount(v string) *IsvGetAppIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *IsvGetAppIdRequest
	GetResourceOwnerId() *int64
	SetType(v string) *IsvGetAppIdRequest
	GetType() *string
}

type IsvGetAppIdRequest struct {
	// example:
	//
	// 示例值示例值示例值
	IntlVersion *string `json:"IntlVersion,omitempty" xml:"IntlVersion,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The permission.
	//
	// Valid values:
	//
	// 	- whatsapp_business_messaging: sending permission on WhatsApp messages
	//
	// 	- ads_management: management permission on advertisements
	//
	// 	- catalog_management: management permission on catalogs
	//
	// example:
	//
	// catalog_management
	Permissions          *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the app. Valid value: WHATSAPP.
	//
	// This parameter is required.
	//
	// example:
	//
	// WHATSAPP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s IsvGetAppIdRequest) String() string {
	return dara.Prettify(s)
}

func (s IsvGetAppIdRequest) GoString() string {
	return s.String()
}

func (s *IsvGetAppIdRequest) GetIntlVersion() *string {
	return s.IntlVersion
}

func (s *IsvGetAppIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *IsvGetAppIdRequest) GetPermissions() *string {
	return s.Permissions
}

func (s *IsvGetAppIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *IsvGetAppIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *IsvGetAppIdRequest) GetType() *string {
	return s.Type
}

func (s *IsvGetAppIdRequest) SetIntlVersion(v string) *IsvGetAppIdRequest {
	s.IntlVersion = &v
	return s
}

func (s *IsvGetAppIdRequest) SetOwnerId(v int64) *IsvGetAppIdRequest {
	s.OwnerId = &v
	return s
}

func (s *IsvGetAppIdRequest) SetPermissions(v string) *IsvGetAppIdRequest {
	s.Permissions = &v
	return s
}

func (s *IsvGetAppIdRequest) SetResourceOwnerAccount(v string) *IsvGetAppIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *IsvGetAppIdRequest) SetResourceOwnerId(v int64) *IsvGetAppIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *IsvGetAppIdRequest) SetType(v string) *IsvGetAppIdRequest {
	s.Type = &v
	return s
}

func (s *IsvGetAppIdRequest) Validate() error {
	return dara.Validate(s)
}
