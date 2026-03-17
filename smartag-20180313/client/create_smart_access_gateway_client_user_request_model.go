// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartAccessGatewayClientUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateSmartAccessGatewayClientUserRequest
	GetBandwidth() *int64
	SetClientIp(v string) *CreateSmartAccessGatewayClientUserRequest
	GetClientIp() *string
	SetOwnerAccount(v string) *CreateSmartAccessGatewayClientUserRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSmartAccessGatewayClientUserRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateSmartAccessGatewayClientUserRequest
	GetPassword() *string
	SetRegionId(v string) *CreateSmartAccessGatewayClientUserRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateSmartAccessGatewayClientUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmartAccessGatewayClientUserRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *CreateSmartAccessGatewayClientUserRequest
	GetSmartAGId() *string
	SetUserMail(v string) *CreateSmartAccessGatewayClientUserRequest
	GetUserMail() *string
	SetUserName(v string) *CreateSmartAccessGatewayClientUserRequest
	GetUserName() *string
}

type CreateSmartAccessGatewayClientUserRequest struct {
	// The maximum bandwidth value. Unit: Kbit/s. Valid values: **1 to 20000**. Default value: **2000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// 	- If you enable the client app service, you must set the IP address of the client app. The current client account uses the specified IP address to connect to Alibaba Cloud.
	//
	// >  The IP address must fall within a private CIDR block.
	//
	// 	- If you disable the client app service, an IP address within a private CIDR block is assigned to the client account. Each connection to Alibaba Cloud uses a different IP address.
	//
	// example:
	//
	// 10.0.XX.XX
	ClientIp     *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password that is used to log on to the SAG app.
	//
	// The password must be 8 to 32 characters in length. It can contain letters, digits, underscores (_), at signs (@), and hyphens (-). It must start with a letter or a digit.
	//
	// example:
	//
	// duuf****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the Smart Access Gateway (SAG) app instance.
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
	// sag-gnhe6sywtare5****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The email address of the client account. The username and password are sent to the specified email address by the administrator.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	UserMail *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	// The username of the client account. The usernames of client accounts added to the same SAG app instance must be unique.
	//
	// The username must be 7 to 33 characters in length, and can contain letters, digits, underscores (_), at signs (@), periods (.), and hyphens (-). It must start with a letter or a digit.
	//
	// >  For a client account, if you specify the username, you must also specify the password. If you specify the password, you must specify the username.
	//
	// example:
	//
	// doctest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateSmartAccessGatewayClientUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartAccessGatewayClientUserRequest) GoString() string {
	return s.String()
}

func (s *CreateSmartAccessGatewayClientUserRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateSmartAccessGatewayClientUserRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *CreateSmartAccessGatewayClientUserRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSmartAccessGatewayClientUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmartAccessGatewayClientUserRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateSmartAccessGatewayClientUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSmartAccessGatewayClientUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmartAccessGatewayClientUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmartAccessGatewayClientUserRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *CreateSmartAccessGatewayClientUserRequest) GetUserMail() *string {
	return s.UserMail
}

func (s *CreateSmartAccessGatewayClientUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateSmartAccessGatewayClientUserRequest) SetBandwidth(v int64) *CreateSmartAccessGatewayClientUserRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserRequest) SetClientIp(v string) *CreateSmartAccessGatewayClientUserRequest {
	s.ClientIp = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserRequest) SetOwnerAccount(v string) *CreateSmartAccessGatewayClientUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserRequest) SetOwnerId(v int64) *CreateSmartAccessGatewayClientUserRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserRequest) SetPassword(v string) *CreateSmartAccessGatewayClientUserRequest {
	s.Password = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserRequest) SetRegionId(v string) *CreateSmartAccessGatewayClientUserRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserRequest) SetResourceOwnerAccount(v string) *CreateSmartAccessGatewayClientUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserRequest) SetResourceOwnerId(v int64) *CreateSmartAccessGatewayClientUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserRequest) SetSmartAGId(v string) *CreateSmartAccessGatewayClientUserRequest {
	s.SmartAGId = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserRequest) SetUserMail(v string) *CreateSmartAccessGatewayClientUserRequest {
	s.UserMail = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserRequest) SetUserName(v string) *CreateSmartAccessGatewayClientUserRequest {
	s.UserName = &v
	return s
}

func (s *CreateSmartAccessGatewayClientUserRequest) Validate() error {
	return dara.Validate(s)
}
