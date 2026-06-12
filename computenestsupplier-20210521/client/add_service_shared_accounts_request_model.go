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
	// A unique identifier that you provide to ensure the idempotence of the request. The token can contain only ASCII characters and cannot be longer than 64 characters.
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
	// The shared accounts and their permissions.
	//
	// This parameter is required.
	SharedAccounts []*AddServiceSharedAccountsRequestSharedAccounts `json:"SharedAccounts,omitempty" xml:"SharedAccounts,omitempty" type:"Repeated"`
	// The service sharing type. The default value is SharedAccount. Valid values:
	//
	// - SharedAccount: The service is shared with a specified account.
	//
	// - Reseller: The service is shared with a reseller.
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
	if s.SharedAccounts != nil {
		for _, item := range s.SharedAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddServiceSharedAccountsRequestSharedAccounts struct {
	// The permission type. Valid values:
	//
	// - Deployable: The service can be deployed.
	//
	// - Accessible: The service can be accessed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Accessible
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The UID of the Alibaba Cloud account.
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
