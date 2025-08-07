// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *ModifyAccountDescriptionRequest
	GetAccountDescription() *string
	SetAccountName(v string) *ModifyAccountDescriptionRequest
	GetAccountName() *string
	SetDBClusterId(v string) *ModifyAccountDescriptionRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyAccountDescriptionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAccountDescriptionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAccountDescriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountDescriptionRequest
	GetResourceOwnerId() *int64
}

type ModifyAccountDescriptionRequest struct {
	// The description of the account. The description must meet the following requirements:
	//
	// 	- The description cannot start with `http://` or `https://`.
	//
	// 	- The description must be 2 to 256 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID of the cluster.
	//
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

func (s ModifyAccountDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *ModifyAccountDescriptionRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountDescriptionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAccountDescriptionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAccountDescriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAccountDescriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountDescriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBClusterId(v string) *ModifyAccountDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
