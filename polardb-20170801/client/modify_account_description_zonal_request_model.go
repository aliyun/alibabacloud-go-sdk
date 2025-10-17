// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *ModifyAccountDescriptionZonalRequest
	GetAccountDescription() *string
	SetAccountName(v string) *ModifyAccountDescriptionZonalRequest
	GetAccountName() *string
	SetClientToken(v string) *ModifyAccountDescriptionZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *ModifyAccountDescriptionZonalRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyAccountDescriptionZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAccountDescriptionZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAccountDescriptionZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountDescriptionZonalRequest
	GetResourceOwnerId() *int64
}

type ModifyAccountDescriptionZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testdes
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 6000170000591aed949d0f5********************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAccountDescriptionZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionZonalRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionZonalRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *ModifyAccountDescriptionZonalRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountDescriptionZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAccountDescriptionZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAccountDescriptionZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAccountDescriptionZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAccountDescriptionZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountDescriptionZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountDescriptionZonalRequest) SetAccountDescription(v string) *ModifyAccountDescriptionZonalRequest {
	s.AccountDescription = &v
	return s
}

func (s *ModifyAccountDescriptionZonalRequest) SetAccountName(v string) *ModifyAccountDescriptionZonalRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionZonalRequest) SetClientToken(v string) *ModifyAccountDescriptionZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAccountDescriptionZonalRequest) SetDBClusterId(v string) *ModifyAccountDescriptionZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountDescriptionZonalRequest) SetOwnerAccount(v string) *ModifyAccountDescriptionZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionZonalRequest) SetOwnerId(v int64) *ModifyAccountDescriptionZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionZonalRequest) SetResourceOwnerAccount(v string) *ModifyAccountDescriptionZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionZonalRequest) SetResourceOwnerId(v int64) *ModifyAccountDescriptionZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionZonalRequest) Validate() error {
	return dara.Validate(s)
}
