// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUIProxyAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyUIProxyAccountPasswordRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ModifyUIProxyAccountPasswordRequest
	GetAccountPassword() *string
	SetClusterId(v string) *ModifyUIProxyAccountPasswordRequest
	GetClusterId() *string
	SetOwnerId(v int64) *ModifyUIProxyAccountPasswordRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyUIProxyAccountPasswordRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyUIProxyAccountPasswordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyUIProxyAccountPasswordRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *ModifyUIProxyAccountPasswordRequest
	GetZoneId() *string
}

type ModifyUIProxyAccountPasswordRequest struct {
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyUIProxyAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUIProxyAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyUIProxyAccountPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyUIProxyAccountPasswordRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ModifyUIProxyAccountPasswordRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyUIProxyAccountPasswordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyUIProxyAccountPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUIProxyAccountPasswordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyUIProxyAccountPasswordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyUIProxyAccountPasswordRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyUIProxyAccountPasswordRequest) SetAccountName(v string) *ModifyUIProxyAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetAccountPassword(v string) *ModifyUIProxyAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetClusterId(v string) *ModifyUIProxyAccountPasswordRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetOwnerId(v int64) *ModifyUIProxyAccountPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetRegionId(v string) *ModifyUIProxyAccountPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetResourceOwnerAccount(v string) *ModifyUIProxyAccountPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetResourceOwnerId(v int64) *ModifyUIProxyAccountPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) SetZoneId(v string) *ModifyUIProxyAccountPasswordRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
