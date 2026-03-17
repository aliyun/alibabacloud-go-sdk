// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSmartAccessGatewayClientUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ResetSmartAccessGatewayClientUserPasswordRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResetSmartAccessGatewayClientUserPasswordRequest
	GetOwnerId() *int64
	SetPassword(v string) *ResetSmartAccessGatewayClientUserPasswordRequest
	GetPassword() *string
	SetRegionId(v string) *ResetSmartAccessGatewayClientUserPasswordRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ResetSmartAccessGatewayClientUserPasswordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResetSmartAccessGatewayClientUserPasswordRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ResetSmartAccessGatewayClientUserPasswordRequest
	GetSmartAGId() *string
	SetUserName(v string) *ResetSmartAccessGatewayClientUserPasswordRequest
	GetUserName() *string
}

type ResetSmartAccessGatewayClientUserPasswordRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The new password.
	//
	// 	- If you do not specify a new password, the system generates a random password.
	//
	// 	- If you specify a new password, the password is reset to the specified new password.
	//
	//     The password must be 8 to 32 characters in length, and can contain letters, digits, underscores (_), at signs (@), periods (.), and hyphens (-). It must start with a letter or a digit.
	//
	// After the password is reset, an email that contains the new password is sent to the email address of the client account.
	//
	// example:
	//
	// Password****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the region where the SAG app instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-kzo5dvms3dqii3****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The username of the client account for which you want to reset the password.
	//
	// This parameter is required.
	//
	// example:
	//
	// username
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ResetSmartAccessGatewayClientUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetSmartAccessGatewayClientUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) GetUserName() *string {
	return s.UserName
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) SetOwnerAccount(v string) *ResetSmartAccessGatewayClientUserPasswordRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) SetOwnerId(v int64) *ResetSmartAccessGatewayClientUserPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) SetPassword(v string) *ResetSmartAccessGatewayClientUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) SetRegionId(v string) *ResetSmartAccessGatewayClientUserPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) SetResourceOwnerAccount(v string) *ResetSmartAccessGatewayClientUserPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) SetResourceOwnerId(v int64) *ResetSmartAccessGatewayClientUserPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) SetSmartAGId(v string) *ResetSmartAccessGatewayClientUserPasswordRequest {
	s.SmartAGId = &v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) SetUserName(v string) *ResetSmartAccessGatewayClientUserPasswordRequest {
	s.UserName = &v
	return s
}

func (s *ResetSmartAccessGatewayClientUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
