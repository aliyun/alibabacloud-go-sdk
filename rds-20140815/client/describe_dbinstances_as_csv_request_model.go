// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesAsCsvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCachedAsync(v bool) *DescribeDBInstancesAsCsvRequest
	GetCachedAsync() *bool
	SetDBInstanceId(v string) *DescribeDBInstancesAsCsvRequest
	GetDBInstanceId() *string
	SetExportKey(v string) *DescribeDBInstancesAsCsvRequest
	GetExportKey() *string
	SetOwnerId(v int64) *DescribeDBInstancesAsCsvRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBInstancesAsCsvRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesAsCsvRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstancesAsCsvRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancesAsCsvRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstancesAsCsvRequest struct {
	CachedAsync  *bool   `json:"CachedAsync,omitempty" xml:"CachedAsync,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ExportKey    *string `json:"ExportKey,omitempty" xml:"ExportKey,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstancesAsCsvRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesAsCsvRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesAsCsvRequest) GetCachedAsync() *bool {
	return s.CachedAsync
}

func (s *DescribeDBInstancesAsCsvRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesAsCsvRequest) GetExportKey() *string {
	return s.ExportKey
}

func (s *DescribeDBInstancesAsCsvRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancesAsCsvRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesAsCsvRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesAsCsvRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstancesAsCsvRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancesAsCsvRequest) SetCachedAsync(v bool) *DescribeDBInstancesAsCsvRequest {
	s.CachedAsync = &v
	return s
}

func (s *DescribeDBInstancesAsCsvRequest) SetDBInstanceId(v string) *DescribeDBInstancesAsCsvRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvRequest) SetExportKey(v string) *DescribeDBInstancesAsCsvRequest {
	s.ExportKey = &v
	return s
}

func (s *DescribeDBInstancesAsCsvRequest) SetOwnerId(v int64) *DescribeDBInstancesAsCsvRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvRequest) SetRegionId(v string) *DescribeDBInstancesAsCsvRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvRequest) SetResourceGroupId(v string) *DescribeDBInstancesAsCsvRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancesAsCsvRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesAsCsvRequest) SetResourceOwnerId(v int64) *DescribeDBInstancesAsCsvRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvRequest) Validate() error {
	return dara.Validate(s)
}
