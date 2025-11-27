// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceCrossBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeInstanceCrossBackupPolicyRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeInstanceCrossBackupPolicyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeInstanceCrossBackupPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceCrossBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceCrossBackupPolicyRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceCrossBackupPolicyRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
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

func (s DescribeInstanceCrossBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceCrossBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCrossBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeInstanceCrossBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceCrossBackupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceCrossBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceCrossBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceCrossBackupPolicyRequest) SetDBInstanceId(v string) *DescribeInstanceCrossBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyRequest) SetOwnerId(v int64) *DescribeInstanceCrossBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyRequest) SetRegionId(v string) *DescribeInstanceCrossBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeInstanceCrossBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeInstanceCrossBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
