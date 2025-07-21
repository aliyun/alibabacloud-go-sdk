// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPWByDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ModifyPWByDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *ModifyPWByDomainRequest
	GetOwnerId() *int64
	SetPassword(v string) *ModifyPWByDomainRequest
	GetPassword() *string
	SetResourceOwnerAccount(v string) *ModifyPWByDomainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPWByDomainRequest
	GetResourceOwnerId() *int64
}

type ModifyPWByDomainRequest struct {
	// Domain name, length 1-50, can include numbers, uppercase letters, lowercase letters, ., and -.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// - Length should be between 10 to 20 characters, and must contain numbers, uppercase letters, and lowercase letters.
	//
	// - At least 2 digits, 2 uppercase letters, and 2 lowercase letters are required, and neither digits nor letters can consist of a single character repeated.
	//
	// - Cannot be the same as the last set password.
	//
	// This parameter is required.
	//
	// example:
	//
	// DM1mail1234
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyPWByDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPWByDomainRequest) GoString() string {
	return s.String()
}

func (s *ModifyPWByDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyPWByDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPWByDomainRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyPWByDomainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPWByDomainRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPWByDomainRequest) SetDomainName(v string) *ModifyPWByDomainRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyPWByDomainRequest) SetOwnerId(v int64) *ModifyPWByDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPWByDomainRequest) SetPassword(v string) *ModifyPWByDomainRequest {
	s.Password = &v
	return s
}

func (s *ModifyPWByDomainRequest) SetResourceOwnerAccount(v string) *ModifyPWByDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPWByDomainRequest) SetResourceOwnerId(v int64) *ModifyPWByDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPWByDomainRequest) Validate() error {
	return dara.Validate(s)
}
