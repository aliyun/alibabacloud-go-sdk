// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceKeywordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *DescribeInstanceKeywordsRequest
	GetKey() *string
	SetOwnerAccount(v string) *DescribeInstanceKeywordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceKeywordsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeInstanceKeywordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceKeywordsRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceKeywordsRequest struct {
	// The type of reserved keyword to query. Valid values:
	//
	// 	- **account**
	//
	// 	- **database**
	//
	// >  This parameter is required.
	//
	// example:
	//
	// account
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeInstanceKeywordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceKeywordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceKeywordsRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeInstanceKeywordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceKeywordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceKeywordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceKeywordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceKeywordsRequest) SetKey(v string) *DescribeInstanceKeywordsRequest {
	s.Key = &v
	return s
}

func (s *DescribeInstanceKeywordsRequest) SetOwnerAccount(v string) *DescribeInstanceKeywordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceKeywordsRequest) SetOwnerId(v int64) *DescribeInstanceKeywordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceKeywordsRequest) SetResourceOwnerAccount(v string) *DescribeInstanceKeywordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceKeywordsRequest) SetResourceOwnerId(v int64) *DescribeInstanceKeywordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceKeywordsRequest) Validate() error {
	return dara.Validate(s)
}
