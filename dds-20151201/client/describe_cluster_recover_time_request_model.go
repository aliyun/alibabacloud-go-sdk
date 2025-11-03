// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterRecoverTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeClusterRecoverTimeRequest
	GetDBInstanceId() *string
	SetDestRegion(v string) *DescribeClusterRecoverTimeRequest
	GetDestRegion() *string
	SetOwnerAccount(v string) *DescribeClusterRecoverTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeClusterRecoverTimeRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeClusterRecoverTimeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeClusterRecoverTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClusterRecoverTimeRequest
	GetResourceOwnerId() *int64
	SetSrcRegion(v string) *DescribeClusterRecoverTimeRequest
	GetSrcRegion() *string
}

type DescribeClusterRecoverTimeRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp18f7d6b6a7****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DestRegion   *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SrcRegion            *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
}

func (s DescribeClusterRecoverTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterRecoverTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterRecoverTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeClusterRecoverTimeRequest) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeClusterRecoverTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeClusterRecoverTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClusterRecoverTimeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClusterRecoverTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClusterRecoverTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClusterRecoverTimeRequest) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeClusterRecoverTimeRequest) SetDBInstanceId(v string) *DescribeClusterRecoverTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeClusterRecoverTimeRequest) SetDestRegion(v string) *DescribeClusterRecoverTimeRequest {
	s.DestRegion = &v
	return s
}

func (s *DescribeClusterRecoverTimeRequest) SetOwnerAccount(v string) *DescribeClusterRecoverTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeClusterRecoverTimeRequest) SetOwnerId(v int64) *DescribeClusterRecoverTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClusterRecoverTimeRequest) SetResourceGroupId(v string) *DescribeClusterRecoverTimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterRecoverTimeRequest) SetResourceOwnerAccount(v string) *DescribeClusterRecoverTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClusterRecoverTimeRequest) SetResourceOwnerId(v int64) *DescribeClusterRecoverTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterRecoverTimeRequest) SetSrcRegion(v string) *DescribeClusterRecoverTimeRequest {
	s.SrcRegion = &v
	return s
}

func (s *DescribeClusterRecoverTimeRequest) Validate() error {
	return dara.Validate(s)
}
