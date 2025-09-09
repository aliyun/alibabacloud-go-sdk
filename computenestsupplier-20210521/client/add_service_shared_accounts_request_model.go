// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddServiceSharedAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddServiceSharedAccountsRequest
	GetClientToken() *string
	SetRegionId(v string) *AddServiceSharedAccountsRequest
	GetRegionId() *string
	SetServiceId(v string) *AddServiceSharedAccountsRequest
	GetServiceId() *string
	SetSharedAccounts(v []*AddServiceSharedAccountsRequestSharedAccounts) *AddServiceSharedAccountsRequest
	GetSharedAccounts() []*AddServiceSharedAccountsRequestSharedAccounts
	SetType(v string) *AddServiceSharedAccountsRequest
	GetType() *string
}

type AddServiceSharedAccountsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-63b8a060e9d54cxxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The shared account and permissions of the service.
	//
	// This parameter is required.
	SharedAccounts []*AddServiceSharedAccountsRequestSharedAccounts `json:"SharedAccounts,omitempty" xml:"SharedAccounts,omitempty" type:"Repeated"`
	// The share type of the service. Default value: SharedAccount. Valid values:
	//
	// 	- SharedAccount: The service is shared by multiple accounts.
	//
	// 	- Reseller: The service is distributed.
	//
	// example:
	//
	// SharedAccount
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddServiceSharedAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddServiceSharedAccountsRequest) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddServiceSharedAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddServiceSharedAccountsRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *AddServiceSharedAccountsRequest) GetSharedAccounts() []*AddServiceSharedAccountsRequestSharedAccounts {
	return s.SharedAccounts
}

func (s *AddServiceSharedAccountsRequest) GetType() *string {
	return s.Type
}

func (s *AddServiceSharedAccountsRequest) SetClientToken(v string) *AddServiceSharedAccountsRequest {
	s.ClientToken = &v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetRegionId(v string) *AddServiceSharedAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetServiceId(v string) *AddServiceSharedAccountsRequest {
	s.ServiceId = &v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetSharedAccounts(v []*AddServiceSharedAccountsRequestSharedAccounts) *AddServiceSharedAccountsRequest {
	s.SharedAccounts = v
	return s
}

func (s *AddServiceSharedAccountsRequest) SetType(v string) *AddServiceSharedAccountsRequest {
	s.Type = &v
	return s
}

func (s *AddServiceSharedAccountsRequest) Validate() error {
	return dara.Validate(s)
}

type AddServiceSharedAccountsRequestSharedAccounts struct {
	// The permissions on the service. Valid values:
	//
	// 	- Deployable: Permissions to deploy the service.
	//
	// 	- Accessible: Permissions to access the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// Accessible
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The Alibaba Cloud account ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	UserAliUid *string `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
}

func (s AddServiceSharedAccountsRequestSharedAccounts) String() string {
	return dara.Prettify(s)
}

func (s AddServiceSharedAccountsRequestSharedAccounts) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsRequestSharedAccounts) GetPermission() *string {
	return s.Permission
}

func (s *AddServiceSharedAccountsRequestSharedAccounts) GetUserAliUid() *string {
	return s.UserAliUid
}

func (s *AddServiceSharedAccountsRequestSharedAccounts) SetPermission(v string) *AddServiceSharedAccountsRequestSharedAccounts {
	s.Permission = &v
	return s
}

func (s *AddServiceSharedAccountsRequestSharedAccounts) SetUserAliUid(v string) *AddServiceSharedAccountsRequestSharedAccounts {
	s.UserAliUid = &v
	return s
}

func (s *AddServiceSharedAccountsRequestSharedAccounts) Validate() error {
	return dara.Validate(s)
}
