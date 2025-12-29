// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribePhoneNumberAttributeRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *DescribePhoneNumberAttributeRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *DescribePhoneNumberAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhoneNumberAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribePhoneNumberAttributeRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhoneNumberAttributeRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *DescribePhoneNumberAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhoneNumberAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhoneNumberAttributeRequest) SetOwnerId(v int64) *DescribePhoneNumberAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetPhoneNumber(v string) *DescribePhoneNumberAttributeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberAttributeRequest) Validate() error {
	return dara.Validate(s)
}
