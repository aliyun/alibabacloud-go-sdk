// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifySchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerId(v int64) *DescribeVerifySchemeRequest
	GetCustomerId() *int64
	SetOwnerId(v int64) *DescribeVerifySchemeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeVerifySchemeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVerifySchemeRequest
	GetResourceOwnerId() *int64
	SetSchemeCode(v string) *DescribeVerifySchemeRequest
	GetSchemeCode() *string
}

type DescribeVerifySchemeRequest struct {
	// The user ID.
	//
	// example:
	//
	// 1234****
	CustomerId           *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service code.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC10000010643****
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s DescribeVerifySchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeRequest) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *DescribeVerifySchemeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVerifySchemeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVerifySchemeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVerifySchemeRequest) GetSchemeCode() *string {
	return s.SchemeCode
}

func (s *DescribeVerifySchemeRequest) SetCustomerId(v int64) *DescribeVerifySchemeRequest {
	s.CustomerId = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetOwnerId(v int64) *DescribeVerifySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetResourceOwnerAccount(v string) *DescribeVerifySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetResourceOwnerId(v int64) *DescribeVerifySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetSchemeCode(v string) *DescribeVerifySchemeRequest {
	s.SchemeCode = &v
	return s
}

func (s *DescribeVerifySchemeRequest) Validate() error {
	return dara.Validate(s)
}
