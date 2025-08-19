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
	SetDryRun(v bool) *UpdateAccountRequest
	GetDryRun() *bool
	SetNewAccountType(v string) *UpdateAccountRequest
	GetNewAccountType() *string
	SetNewDisplayName(v string) *UpdateAccountRequest
	GetNewDisplayName() *string
}

type UpdateAccountRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks items such as whether the member status can be modified and whether security information is configured for the member. If the request does not pass the dry run, an error code is returned.
	//
	// 	- false (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new type of the member. Valid values:
	//
	// 	- ResourceAccount: resource account
	//
	// 	- CloudAccount: cloud account
	//
	// > You can specify either `NewDisplayName` or `NewAccountType`.
	//
	// example:
	//
	// ResourceAccount
	NewAccountType *string `json:"NewAccountType,omitempty" xml:"NewAccountType,omitempty"`
	// The new display name of the member.
	//
	// > You can specify either `NewDisplayName` or `NewAccountType`.
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

func (s *UpdateAccountRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *UpdateAccountRequest) SetDryRun(v bool) *UpdateAccountRequest {
	s.DryRun = &v
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
