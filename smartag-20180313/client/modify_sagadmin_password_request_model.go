// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySAGAdminPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ModifySAGAdminPasswordRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySAGAdminPasswordRequest
	GetOwnerId() *int64
	SetPassword(v string) *ModifySAGAdminPasswordRequest
	GetPassword() *string
	SetRegionId(v string) *ModifySAGAdminPasswordRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySAGAdminPasswordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySAGAdminPasswordRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ModifySAGAdminPasswordRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySAGAdminPasswordRequest
	GetSmartAGSn() *string
}

type ModifySAGAdminPasswordRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The new logon password for the SAG device.
	//
	// The password must be 8 to 30 characters in length and can contain letters, digits, and underscores (_).
	//
	// >  In the example, asterisks (\\*) are used to conceal the real password. This does not mean that the password supports asterisks (\\*). The actual format requirement prevails.
	//
	// This parameter is required.
	//
	// example:
	//
	// ********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
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
	// sag-tq3sazs17smldn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number (SN) of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s ModifySAGAdminPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySAGAdminPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifySAGAdminPasswordRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySAGAdminPasswordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySAGAdminPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifySAGAdminPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySAGAdminPasswordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySAGAdminPasswordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySAGAdminPasswordRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySAGAdminPasswordRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySAGAdminPasswordRequest) SetOwnerAccount(v string) *ModifySAGAdminPasswordRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySAGAdminPasswordRequest) SetOwnerId(v int64) *ModifySAGAdminPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySAGAdminPasswordRequest) SetPassword(v string) *ModifySAGAdminPasswordRequest {
	s.Password = &v
	return s
}

func (s *ModifySAGAdminPasswordRequest) SetRegionId(v string) *ModifySAGAdminPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySAGAdminPasswordRequest) SetResourceOwnerAccount(v string) *ModifySAGAdminPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySAGAdminPasswordRequest) SetResourceOwnerId(v int64) *ModifySAGAdminPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySAGAdminPasswordRequest) SetSmartAGId(v string) *ModifySAGAdminPasswordRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySAGAdminPasswordRequest) SetSmartAGSn(v string) *ModifySAGAdminPasswordRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySAGAdminPasswordRequest) Validate() error {
	return dara.Validate(s)
}
