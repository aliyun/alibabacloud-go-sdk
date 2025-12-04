// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountAuthorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeAccountAuthorityRequest
	GetAccountName() *string
	SetDBClusterId(v string) *DescribeAccountAuthorityRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeAccountAuthorityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAccountAuthorityRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAccountAuthorityRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAccountAuthorityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccountAuthorityRequest
	GetResourceOwnerId() *int64
}

type DescribeAccountAuthorityRequest struct {
	// The name of the database account.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1p816075e21****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAccountAuthorityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAuthorityRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountAuthorityRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAccountAuthorityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAccountAuthorityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAccountAuthorityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccountAuthorityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAccountAuthorityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAccountAuthorityRequest) SetAccountName(v string) *DescribeAccountAuthorityRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetDBClusterId(v string) *DescribeAccountAuthorityRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetOwnerAccount(v string) *DescribeAccountAuthorityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetOwnerId(v int64) *DescribeAccountAuthorityRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetRegionId(v string) *DescribeAccountAuthorityRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetResourceOwnerAccount(v string) *DescribeAccountAuthorityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetResourceOwnerId(v int64) *DescribeAccountAuthorityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) Validate() error {
	return dara.Validate(s)
}
