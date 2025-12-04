// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOSSStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeOSSStorageRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeOSSStorageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeOSSStorageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeOSSStorageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeOSSStorageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeOSSStorageRequest
	GetResourceOwnerId() *int64
}

type DescribeOSSStorageRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
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

func (s DescribeOSSStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOSSStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeOSSStorageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeOSSStorageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeOSSStorageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeOSSStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOSSStorageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeOSSStorageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeOSSStorageRequest) SetDBClusterId(v string) *DescribeOSSStorageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeOSSStorageRequest) SetOwnerAccount(v string) *DescribeOSSStorageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeOSSStorageRequest) SetOwnerId(v int64) *DescribeOSSStorageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeOSSStorageRequest) SetRegionId(v string) *DescribeOSSStorageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOSSStorageRequest) SetResourceOwnerAccount(v string) *DescribeOSSStorageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeOSSStorageRequest) SetResourceOwnerId(v int64) *DescribeOSSStorageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeOSSStorageRequest) Validate() error {
	return dara.Validate(s)
}
