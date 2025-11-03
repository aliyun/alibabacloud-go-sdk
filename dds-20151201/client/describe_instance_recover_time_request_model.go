// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRecoverTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeInstanceRecoverTimeRequest
	GetDBInstanceId() *string
	SetDestRegion(v string) *DescribeInstanceRecoverTimeRequest
	GetDestRegion() *string
	SetOwnerAccount(v string) *DescribeInstanceRecoverTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceRecoverTimeRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeInstanceRecoverTimeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceRecoverTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceRecoverTimeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeInstanceRecoverTimeRequest
	GetSecurityToken() *string
	SetSrcRegion(v string) *DescribeInstanceRecoverTimeRequest
	GetSrcRegion() *string
}

type DescribeInstanceRecoverTimeRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
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
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SrcRegion            *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
}

func (s DescribeInstanceRecoverTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRecoverTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRecoverTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeInstanceRecoverTimeRequest) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeInstanceRecoverTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceRecoverTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceRecoverTimeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceRecoverTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceRecoverTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceRecoverTimeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceRecoverTimeRequest) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeInstanceRecoverTimeRequest) SetDBInstanceId(v string) *DescribeInstanceRecoverTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeInstanceRecoverTimeRequest) SetDestRegion(v string) *DescribeInstanceRecoverTimeRequest {
	s.DestRegion = &v
	return s
}

func (s *DescribeInstanceRecoverTimeRequest) SetOwnerAccount(v string) *DescribeInstanceRecoverTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceRecoverTimeRequest) SetOwnerId(v int64) *DescribeInstanceRecoverTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceRecoverTimeRequest) SetResourceGroupId(v string) *DescribeInstanceRecoverTimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceRecoverTimeRequest) SetResourceOwnerAccount(v string) *DescribeInstanceRecoverTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceRecoverTimeRequest) SetResourceOwnerId(v int64) *DescribeInstanceRecoverTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceRecoverTimeRequest) SetSecurityToken(v string) *DescribeInstanceRecoverTimeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceRecoverTimeRequest) SetSrcRegion(v string) *DescribeInstanceRecoverTimeRequest {
	s.SrcRegion = &v
	return s
}

func (s *DescribeInstanceRecoverTimeRequest) Validate() error {
	return dara.Validate(s)
}
