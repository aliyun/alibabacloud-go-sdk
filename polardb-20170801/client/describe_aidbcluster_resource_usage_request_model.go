// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterResourceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAIDBClusterResourceUsageRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeAIDBClusterResourceUsageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAIDBClusterResourceUsageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAIDBClusterResourceUsageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAIDBClusterResourceUsageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAIDBClusterResourceUsageRequest
	GetResourceOwnerId() *int64
}

type DescribeAIDBClusterResourceUsageRequest struct {
	// The ID of the AI database cluster.
	//
	// example:
	//
	// pc-2zeu0t5v0p98*****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the AI database cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAIDBClusterResourceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterResourceUsageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterResourceUsageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAIDBClusterResourceUsageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAIDBClusterResourceUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAIDBClusterResourceUsageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAIDBClusterResourceUsageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAIDBClusterResourceUsageRequest) SetDBClusterId(v string) *DescribeAIDBClusterResourceUsageRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageRequest) SetOwnerAccount(v string) *DescribeAIDBClusterResourceUsageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageRequest) SetOwnerId(v int64) *DescribeAIDBClusterResourceUsageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageRequest) SetRegionId(v string) *DescribeAIDBClusterResourceUsageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageRequest) SetResourceOwnerAccount(v string) *DescribeAIDBClusterResourceUsageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageRequest) SetResourceOwnerId(v int64) *DescribeAIDBClusterResourceUsageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAIDBClusterResourceUsageRequest) Validate() error {
	return dara.Validate(s)
}
