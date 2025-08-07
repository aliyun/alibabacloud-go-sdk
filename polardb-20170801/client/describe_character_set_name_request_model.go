// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCharacterSetNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeCharacterSetNameRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeCharacterSetNameRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCharacterSetNameRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCharacterSetNameRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeCharacterSetNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCharacterSetNameRequest
	GetResourceOwnerId() *int64
}

type DescribeCharacterSetNameRequest struct {
	// The ID of the cluster.
	//
	// > You can only query character sets that PolarDB for MySQL clusters support. If you enter the ID of a PolarDB for PostgreSQL or PolarDB for Oracle cluster, the returned value of the `CharacterSetNameItems` parameter is an empty string.
	//
	// example:
	//
	// pc-****************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCharacterSetNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCharacterSetNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetNameRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeCharacterSetNameRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCharacterSetNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCharacterSetNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCharacterSetNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCharacterSetNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCharacterSetNameRequest) SetDBClusterId(v string) *DescribeCharacterSetNameRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetOwnerAccount(v string) *DescribeCharacterSetNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetOwnerId(v int64) *DescribeCharacterSetNameRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetRegionId(v string) *DescribeCharacterSetNameRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetResourceOwnerAccount(v string) *DescribeCharacterSetNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetResourceOwnerId(v int64) *DescribeCharacterSetNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) Validate() error {
	return dara.Validate(s)
}
