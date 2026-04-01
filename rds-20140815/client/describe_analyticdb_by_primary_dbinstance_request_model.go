// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalyticdbByPrimaryDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeAnalyticdbByPrimaryDBInstanceRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeAnalyticdbByPrimaryDBInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAnalyticdbByPrimaryDBInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAnalyticdbByPrimaryDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAnalyticdbByPrimaryDBInstanceRequest
	GetResourceOwnerId() *int64
}

type DescribeAnalyticdbByPrimaryDBInstanceRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*******
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

func (s DescribeAnalyticdbByPrimaryDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalyticdbByPrimaryDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceRequest) SetDBInstanceId(v string) *DescribeAnalyticdbByPrimaryDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceRequest) SetOwnerId(v int64) *DescribeAnalyticdbByPrimaryDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceRequest) SetRegionId(v string) *DescribeAnalyticdbByPrimaryDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceRequest) SetResourceOwnerAccount(v string) *DescribeAnalyticdbByPrimaryDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceRequest) SetResourceOwnerId(v int64) *DescribeAnalyticdbByPrimaryDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAnalyticdbByPrimaryDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
