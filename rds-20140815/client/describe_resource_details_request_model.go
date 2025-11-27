// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeResourceDetailsRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeResourceDetailsRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeResourceDetailsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeResourceDetailsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeResourceDetailsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeResourceDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeResourceDetailsRequest
	GetResourceOwnerId() *int64
}

type DescribeResourceDetailsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1ul2y10grt91m68
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm3kyoa2wqhyy
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeResourceDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceDetailsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeResourceDetailsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeResourceDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeResourceDetailsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceDetailsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeResourceDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeResourceDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeResourceDetailsRequest) SetClientToken(v string) *DescribeResourceDetailsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeResourceDetailsRequest) SetDBInstanceId(v string) *DescribeResourceDetailsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeResourceDetailsRequest) SetOwnerId(v int64) *DescribeResourceDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeResourceDetailsRequest) SetRegionId(v string) *DescribeResourceDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceDetailsRequest) SetResourceGroupId(v string) *DescribeResourceDetailsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeResourceDetailsRequest) SetResourceOwnerAccount(v string) *DescribeResourceDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeResourceDetailsRequest) SetResourceOwnerId(v int64) *DescribeResourceDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeResourceDetailsRequest) Validate() error {
	return dara.Validate(s)
}
