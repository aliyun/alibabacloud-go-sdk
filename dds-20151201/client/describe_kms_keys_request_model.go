// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKmsKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeKmsKeysRequest
	GetOwnerAccount() *string
	SetRegionId(v string) *DescribeKmsKeysRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeKmsKeysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeKmsKeysRequest
	GetResourceOwnerId() *int64
}

type DescribeKmsKeysRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeKmsKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeKmsKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKmsKeysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeKmsKeysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeKmsKeysRequest) SetOwnerAccount(v string) *DescribeKmsKeysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeKmsKeysRequest) SetRegionId(v string) *DescribeKmsKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKmsKeysRequest) SetResourceOwnerAccount(v string) *DescribeKmsKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeKmsKeysRequest) SetResourceOwnerId(v int64) *DescribeKmsKeysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeKmsKeysRequest) Validate() error {
	return dara.Validate(s)
}
