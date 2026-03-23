// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableRecoveryTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossBackupId(v int32) *DescribeAvailableRecoveryTimeRequest
	GetCrossBackupId() *int32
	SetDBInstanceId(v string) *DescribeAvailableRecoveryTimeRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeAvailableRecoveryTimeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAvailableRecoveryTimeRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeAvailableRecoveryTimeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAvailableRecoveryTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAvailableRecoveryTimeRequest
	GetResourceOwnerId() *int64
}

type DescribeAvailableRecoveryTimeRequest struct {
	// This parameter is required.
	CrossBackupId        *int32  `json:"CrossBackupId,omitempty" xml:"CrossBackupId,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAvailableRecoveryTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableRecoveryTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableRecoveryTimeRequest) GetCrossBackupId() *int32 {
	return s.CrossBackupId
}

func (s *DescribeAvailableRecoveryTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAvailableRecoveryTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAvailableRecoveryTimeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableRecoveryTimeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAvailableRecoveryTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAvailableRecoveryTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableRecoveryTimeRequest) SetCrossBackupId(v int32) *DescribeAvailableRecoveryTimeRequest {
	s.CrossBackupId = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeRequest) SetDBInstanceId(v string) *DescribeAvailableRecoveryTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeRequest) SetOwnerId(v int64) *DescribeAvailableRecoveryTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeRequest) SetRegionId(v string) *DescribeAvailableRecoveryTimeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeRequest) SetResourceGroupId(v string) *DescribeAvailableRecoveryTimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeRequest) SetResourceOwnerAccount(v string) *DescribeAvailableRecoveryTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeRequest) SetResourceOwnerId(v int64) *DescribeAvailableRecoveryTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeRequest) Validate() error {
	return dara.Validate(s)
}
