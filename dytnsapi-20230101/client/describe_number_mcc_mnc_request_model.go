// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNumberMccMncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribeNumberMccMncRequest
	GetAuthCode() *string
	SetOwnerId(v int64) *DescribeNumberMccMncRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *DescribeNumberMccMncRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *DescribeNumberMccMncRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNumberMccMncRequest
	GetResourceOwnerId() *int64
}

type DescribeNumberMccMncRequest struct {
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 86123434345
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeNumberMccMncRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNumberMccMncRequest) GoString() string {
	return s.String()
}

func (s *DescribeNumberMccMncRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribeNumberMccMncRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNumberMccMncRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *DescribeNumberMccMncRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNumberMccMncRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNumberMccMncRequest) SetAuthCode(v string) *DescribeNumberMccMncRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribeNumberMccMncRequest) SetOwnerId(v int64) *DescribeNumberMccMncRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNumberMccMncRequest) SetPhoneNumber(v string) *DescribeNumberMccMncRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DescribeNumberMccMncRequest) SetResourceOwnerAccount(v string) *DescribeNumberMccMncRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNumberMccMncRequest) SetResourceOwnerId(v int64) *DescribeNumberMccMncRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNumberMccMncRequest) Validate() error {
	return dara.Validate(s)
}
