// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossRegionLogBackupFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossBackupRegion(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetCrossBackupRegion() *string
	SetDBInstanceId(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeCrossRegionLogBackupFilesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCrossRegionLogBackupFilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCrossRegionLogBackupFilesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCrossRegionLogBackupFilesRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetStartTime() *string
}

type DescribeCrossRegionLogBackupFilesRequest struct {
	CrossBackupRegion *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCrossRegionLogBackupFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionLogBackupFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetCrossBackupRegion() *string {
	return s.CrossBackupRegion
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetCrossBackupRegion(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.CrossBackupRegion = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetDBInstanceId(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetEndTime(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetOwnerId(v int64) *DescribeCrossRegionLogBackupFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetPageNumber(v int32) *DescribeCrossRegionLogBackupFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetPageSize(v int32) *DescribeCrossRegionLogBackupFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetRegionId(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetResourceOwnerAccount(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetResourceOwnerId(v int64) *DescribeCrossRegionLogBackupFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetStartTime(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) Validate() error {
	return dara.Validate(s)
}
