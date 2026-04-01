// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLocalAvailableRecoveryTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeLocalAvailableRecoveryTimeRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeLocalAvailableRecoveryTimeRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeLocalAvailableRecoveryTimeRequest
	GetRegion() *string
	SetResourceGroupId(v string) *DescribeLocalAvailableRecoveryTimeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeLocalAvailableRecoveryTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLocalAvailableRecoveryTimeRequest
	GetResourceOwnerId() *int64
}

type DescribeLocalAvailableRecoveryTimeRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLocalAvailableRecoveryTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocalAvailableRecoveryTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetDBInstanceId(v string) *DescribeLocalAvailableRecoveryTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetOwnerId(v int64) *DescribeLocalAvailableRecoveryTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetRegion(v string) *DescribeLocalAvailableRecoveryTimeRequest {
	s.Region = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetResourceGroupId(v string) *DescribeLocalAvailableRecoveryTimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetResourceOwnerAccount(v string) *DescribeLocalAvailableRecoveryTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) SetResourceOwnerId(v int64) *DescribeLocalAvailableRecoveryTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeRequest) Validate() error {
	return dara.Validate(s)
}
