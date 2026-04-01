// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossRegionBackupDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeCrossRegionBackupDBInstanceRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeCrossRegionBackupDBInstanceRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCrossRegionBackupDBInstanceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCrossRegionBackupDBInstanceRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCrossRegionBackupDBInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeCrossRegionBackupDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCrossRegionBackupDBInstanceRequest
	GetResourceOwnerId() *int64
}

type DescribeCrossRegionBackupDBInstanceRequest struct {
	// The instance ID. Up to 30 instance IDs are allowed in a single request. If you enter more than one instance ID, separate them with commas (,).
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Valid values: any non-zero positive integer.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
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

func (s DescribeCrossRegionBackupDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) SetDBInstanceId(v string) *DescribeCrossRegionBackupDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) SetOwnerId(v int64) *DescribeCrossRegionBackupDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) SetPageNumber(v int32) *DescribeCrossRegionBackupDBInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) SetPageSize(v int32) *DescribeCrossRegionBackupDBInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) SetRegionId(v string) *DescribeCrossRegionBackupDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) SetResourceOwnerAccount(v string) *DescribeCrossRegionBackupDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) SetResourceOwnerId(v int64) *DescribeCrossRegionBackupDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
