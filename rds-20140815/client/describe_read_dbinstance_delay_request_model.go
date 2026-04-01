// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReadDBInstanceDelayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeReadDBInstanceDelayRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeReadDBInstanceDelayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeReadDBInstanceDelayRequest
	GetOwnerId() *int64
	SetReadInstanceId(v string) *DescribeReadDBInstanceDelayRequest
	GetReadInstanceId() *string
	SetRegionId(v string) *DescribeReadDBInstanceDelayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeReadDBInstanceDelayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeReadDBInstanceDelayRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeReadDBInstanceDelayRequest
	GetSecurityToken() *string
}

type DescribeReadDBInstanceDelayRequest struct {
	// The primary instance ID. You can call the DescribeDBInstances operation to query the primary instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The read-only instance ID. You can call the DescribeDBInstances operation to query the read-only instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rr-bp*****
	ReadInstanceId *string `json:"ReadInstanceId,omitempty" xml:"ReadInstanceId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeReadDBInstanceDelayRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReadDBInstanceDelayRequest) GoString() string {
	return s.String()
}

func (s *DescribeReadDBInstanceDelayRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeReadDBInstanceDelayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeReadDBInstanceDelayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeReadDBInstanceDelayRequest) GetReadInstanceId() *string {
	return s.ReadInstanceId
}

func (s *DescribeReadDBInstanceDelayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeReadDBInstanceDelayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeReadDBInstanceDelayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeReadDBInstanceDelayRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeReadDBInstanceDelayRequest) SetDBInstanceId(v string) *DescribeReadDBInstanceDelayRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeReadDBInstanceDelayRequest) SetOwnerAccount(v string) *DescribeReadDBInstanceDelayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeReadDBInstanceDelayRequest) SetOwnerId(v int64) *DescribeReadDBInstanceDelayRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeReadDBInstanceDelayRequest) SetReadInstanceId(v string) *DescribeReadDBInstanceDelayRequest {
	s.ReadInstanceId = &v
	return s
}

func (s *DescribeReadDBInstanceDelayRequest) SetRegionId(v string) *DescribeReadDBInstanceDelayRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeReadDBInstanceDelayRequest) SetResourceOwnerAccount(v string) *DescribeReadDBInstanceDelayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeReadDBInstanceDelayRequest) SetResourceOwnerId(v int64) *DescribeReadDBInstanceDelayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeReadDBInstanceDelayRequest) SetSecurityToken(v string) *DescribeReadDBInstanceDelayRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeReadDBInstanceDelayRequest) Validate() error {
	return dara.Validate(s)
}
