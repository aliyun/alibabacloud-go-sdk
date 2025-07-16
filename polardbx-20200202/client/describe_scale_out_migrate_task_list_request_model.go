// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScaleOutMigrateTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeScaleOutMigrateTaskListRequest
	GetDBInstanceName() *string
	SetOwnerAccount(v string) *DescribeScaleOutMigrateTaskListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeScaleOutMigrateTaskListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeScaleOutMigrateTaskListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeScaleOutMigrateTaskListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeScaleOutMigrateTaskListRequest
	GetResourceOwnerId() *int64
}

type DescribeScaleOutMigrateTaskListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeScaleOutMigrateTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScaleOutMigrateTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeScaleOutMigrateTaskListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeScaleOutMigrateTaskListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeScaleOutMigrateTaskListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeScaleOutMigrateTaskListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScaleOutMigrateTaskListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeScaleOutMigrateTaskListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetDBInstanceName(v string) *DescribeScaleOutMigrateTaskListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetOwnerAccount(v string) *DescribeScaleOutMigrateTaskListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetOwnerId(v int64) *DescribeScaleOutMigrateTaskListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetRegionId(v string) *DescribeScaleOutMigrateTaskListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetResourceOwnerAccount(v string) *DescribeScaleOutMigrateTaskListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListRequest) SetResourceOwnerId(v int64) *DescribeScaleOutMigrateTaskListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScaleOutMigrateTaskListRequest) Validate() error {
	return dara.Validate(s)
}
