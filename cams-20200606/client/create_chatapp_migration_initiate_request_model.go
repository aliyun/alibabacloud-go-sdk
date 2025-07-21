// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatappMigrationInitiateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCountryCode(v string) *CreateChatappMigrationInitiateRequest
	GetCountryCode() *string
	SetCustSpaceId(v string) *CreateChatappMigrationInitiateRequest
	GetCustSpaceId() *string
	SetMobileNumber(v string) *CreateChatappMigrationInitiateRequest
	GetMobileNumber() *string
	SetOwnerId(v int64) *CreateChatappMigrationInitiateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateChatappMigrationInitiateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateChatappMigrationInitiateRequest
	GetResourceOwnerId() *int64
}

type CreateChatappMigrationInitiateRequest struct {
	// The code of the country or region.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The space ID of the user within the ISV account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The mobile number without the country code or region code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13900001234
	MobileNumber         *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateChatappMigrationInitiateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChatappMigrationInitiateRequest) GoString() string {
	return s.String()
}

func (s *CreateChatappMigrationInitiateRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *CreateChatappMigrationInitiateRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CreateChatappMigrationInitiateRequest) GetMobileNumber() *string {
	return s.MobileNumber
}

func (s *CreateChatappMigrationInitiateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateChatappMigrationInitiateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateChatappMigrationInitiateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateChatappMigrationInitiateRequest) SetCountryCode(v string) *CreateChatappMigrationInitiateRequest {
	s.CountryCode = &v
	return s
}

func (s *CreateChatappMigrationInitiateRequest) SetCustSpaceId(v string) *CreateChatappMigrationInitiateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateChatappMigrationInitiateRequest) SetMobileNumber(v string) *CreateChatappMigrationInitiateRequest {
	s.MobileNumber = &v
	return s
}

func (s *CreateChatappMigrationInitiateRequest) SetOwnerId(v int64) *CreateChatappMigrationInitiateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateChatappMigrationInitiateRequest) SetResourceOwnerAccount(v string) *CreateChatappMigrationInitiateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateChatappMigrationInitiateRequest) SetResourceOwnerId(v int64) *CreateChatappMigrationInitiateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateChatappMigrationInitiateRequest) Validate() error {
	return dara.Validate(s)
}
