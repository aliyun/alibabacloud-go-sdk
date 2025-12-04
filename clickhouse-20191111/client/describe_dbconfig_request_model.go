// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBConfigRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBConfigRequest
	GetResourceOwnerId() *int64
}

type DescribeDBConfigRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-t4nw17nh2e4t2****
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

func (s DescribeDBConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBConfigRequest) SetDBClusterId(v string) *DescribeDBConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBConfigRequest) SetOwnerAccount(v string) *DescribeDBConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBConfigRequest) SetOwnerId(v int64) *DescribeDBConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBConfigRequest) SetRegionId(v string) *DescribeDBConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBConfigRequest) SetResourceOwnerAccount(v string) *DescribeDBConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBConfigRequest) SetResourceOwnerId(v int64) *DescribeDBConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBConfigRequest) Validate() error {
	return dara.Validate(s)
}
