// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossRegionBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v int32) *DescribeCrossRegionBackupsRequest
	GetBackupId() *int32
	SetCrossBackupId(v int32) *DescribeCrossRegionBackupsRequest
	GetCrossBackupId() *int32
	SetCrossBackupRegion(v string) *DescribeCrossRegionBackupsRequest
	GetCrossBackupRegion() *string
	SetDBInstanceId(v string) *DescribeCrossRegionBackupsRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeCrossRegionBackupsRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeCrossRegionBackupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCrossRegionBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCrossRegionBackupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCrossRegionBackupsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCrossRegionBackupsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCrossRegionBackupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCrossRegionBackupsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeCrossRegionBackupsRequest
	GetStartTime() *string
}

type DescribeCrossRegionBackupsRequest struct {
	BackupId          *int32  `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	CrossBackupId     *int32  `json:"CrossBackupId,omitempty" xml:"CrossBackupId,omitempty"`
	CrossBackupRegion *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCrossRegionBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupsRequest) GetBackupId() *int32 {
	return s.BackupId
}

func (s *DescribeCrossRegionBackupsRequest) GetCrossBackupId() *int32 {
	return s.CrossBackupId
}

func (s *DescribeCrossRegionBackupsRequest) GetCrossBackupRegion() *string {
	return s.CrossBackupRegion
}

func (s *DescribeCrossRegionBackupsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeCrossRegionBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCrossRegionBackupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCrossRegionBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCrossRegionBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCrossRegionBackupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCrossRegionBackupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCrossRegionBackupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCrossRegionBackupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCrossRegionBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCrossRegionBackupsRequest) SetBackupId(v int32) *DescribeCrossRegionBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetCrossBackupId(v int32) *DescribeCrossRegionBackupsRequest {
	s.CrossBackupId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetCrossBackupRegion(v string) *DescribeCrossRegionBackupsRequest {
	s.CrossBackupRegion = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetDBInstanceId(v string) *DescribeCrossRegionBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetEndTime(v string) *DescribeCrossRegionBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetOwnerId(v int64) *DescribeCrossRegionBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetPageNumber(v int32) *DescribeCrossRegionBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetPageSize(v int32) *DescribeCrossRegionBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetRegionId(v string) *DescribeCrossRegionBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetResourceGroupId(v string) *DescribeCrossRegionBackupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetResourceOwnerAccount(v string) *DescribeCrossRegionBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetResourceOwnerId(v int64) *DescribeCrossRegionBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetStartTime(v string) *DescribeCrossRegionBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) Validate() error {
	return dara.Validate(s)
}
