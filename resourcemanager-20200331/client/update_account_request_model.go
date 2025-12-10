// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *UpdateAccountRequest
	GetAccountId() *string
	SetNewAccountType(v string) *UpdateAccountRequest
	GetNewAccountType() *string
	SetNewDisplayName(v string) *UpdateAccountRequest
	GetNewDisplayName() *string
}

type UpdateAccountRequest struct {
	// The ID of the Alibaba Cloud account that corresponds to the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The new type of the member. Valid values:
	//
	// 	- ResourceAccount: resource account
	//
	// 	- CloudAccount: cloud account
	//
	// >  You can configure either the `NewDisplayName` or `NewAccountType` parameter.
	//
	// example:
	//
	// ResourceAccount
	NewAccountType *string `json:"NewAccountType,omitempty" xml:"NewAccountType,omitempty"`
	// The new display name of the member.
	//
	// >  You can configure either the `NewDisplayName` or `NewAccountType` parameter.
	//
	// example:
	//
	// admin
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
}

func (s UpdateAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateAccountRequest) GetNewAccountType() *string {
	return s.NewAccountType
}

func (s *UpdateAccountRequest) GetNewDisplayName() *string {
	return s.NewDisplayName
}

func (s *UpdateAccountRequest) SetAccountId(v string) *UpdateAccountRequest {
	s.AccountId = &v
	return s
}

func (s *UpdateAccountRequest) SetNewAccountType(v string) *UpdateAccountRequest {
	s.NewAccountType = &v
	return s
}

func (s *UpdateAccountRequest) SetNewDisplayName(v string) *UpdateAccountRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateAccountRequest) Validate() error {
	return dara.Validate(s)
}
