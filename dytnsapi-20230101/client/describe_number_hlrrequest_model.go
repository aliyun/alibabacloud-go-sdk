// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNumberHLRRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribeNumberHLRRequest
	GetAuthCode() *string
	SetOwnerId(v int64) *DescribeNumberHLRRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *DescribeNumberHLRRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *DescribeNumberHLRRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNumberHLRRequest
	GetResourceOwnerId() *int64
}

type DescribeNumberHLRRequest struct {
	// example:
	//
	// 示例值示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeNumberHLRRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNumberHLRRequest) GoString() string {
	return s.String()
}

func (s *DescribeNumberHLRRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribeNumberHLRRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNumberHLRRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *DescribeNumberHLRRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNumberHLRRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNumberHLRRequest) SetAuthCode(v string) *DescribeNumberHLRRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribeNumberHLRRequest) SetOwnerId(v int64) *DescribeNumberHLRRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNumberHLRRequest) SetPhoneNumber(v string) *DescribeNumberHLRRequest {
	s.PhoneNumber = &v
	return s
}

func (s *DescribeNumberHLRRequest) SetResourceOwnerAccount(v string) *DescribeNumberHLRRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNumberHLRRequest) SetResourceOwnerId(v int64) *DescribeNumberHLRRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNumberHLRRequest) Validate() error {
	return dara.Validate(s)
}
