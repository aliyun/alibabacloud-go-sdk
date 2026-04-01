// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCloudResourceAuthorizedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CheckCloudResourceAuthorizedRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckCloudResourceAuthorizedRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CheckCloudResourceAuthorizedRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CheckCloudResourceAuthorizedRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckCloudResourceAuthorizedRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CheckCloudResourceAuthorizedRequest
	GetSecurityToken() *string
	SetTargetRegionId(v string) *CheckCloudResourceAuthorizedRequest
	GetTargetRegionId() *string
}

type CheckCloudResourceAuthorizedRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// example:
	//
	// rm-t4n7j9eb52y7c1960
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy**********
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The destination region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// us-east-1
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
}

func (s CheckCloudResourceAuthorizedRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCloudResourceAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CheckCloudResourceAuthorizedRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckCloudResourceAuthorizedRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckCloudResourceAuthorizedRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckCloudResourceAuthorizedRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CheckCloudResourceAuthorizedRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckCloudResourceAuthorizedRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckCloudResourceAuthorizedRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CheckCloudResourceAuthorizedRequest) GetTargetRegionId() *string {
	return s.TargetRegionId
}

func (s *CheckCloudResourceAuthorizedRequest) SetDBInstanceId(v string) *CheckCloudResourceAuthorizedRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetOwnerId(v int64) *CheckCloudResourceAuthorizedRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetRegionId(v string) *CheckCloudResourceAuthorizedRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetResourceGroupId(v string) *CheckCloudResourceAuthorizedRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetResourceOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetResourceOwnerId(v int64) *CheckCloudResourceAuthorizedRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetSecurityToken(v string) *CheckCloudResourceAuthorizedRequest {
	s.SecurityToken = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetTargetRegionId(v string) *CheckCloudResourceAuthorizedRequest {
	s.TargetRegionId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) Validate() error {
	return dara.Validate(s)
}
