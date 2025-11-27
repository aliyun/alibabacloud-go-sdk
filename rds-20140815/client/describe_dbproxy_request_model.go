// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBProxyRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *DescribeDBProxyRequest
	GetDBProxyEngineType() *string
	SetOwnerId(v int64) *DescribeDBProxyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBProxyRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBProxyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBProxyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBProxyRequest
	GetResourceOwnerId() *int64
}

type DescribeDBProxyRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1ja4f56s7us****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// A deprecated parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBProxyRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *DescribeDBProxyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBProxyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBProxyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBProxyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBProxyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBProxyRequest) SetDBInstanceId(v string) *DescribeDBProxyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBProxyRequest) SetDBProxyEngineType(v string) *DescribeDBProxyRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *DescribeDBProxyRequest) SetOwnerId(v int64) *DescribeDBProxyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBProxyRequest) SetRegionId(v string) *DescribeDBProxyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBProxyRequest) SetResourceGroupId(v string) *DescribeDBProxyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBProxyRequest) SetResourceOwnerAccount(v string) *DescribeDBProxyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBProxyRequest) SetResourceOwnerId(v int64) *DescribeDBProxyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBProxyRequest) Validate() error {
	return dara.Validate(s)
}
