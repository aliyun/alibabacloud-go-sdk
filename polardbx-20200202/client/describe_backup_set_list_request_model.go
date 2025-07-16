// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeBackupSetListRequest
	GetDBInstanceName() *string
	SetDestCrossRegion(v string) *DescribeBackupSetListRequest
	GetDestCrossRegion() *string
	SetEndTime(v int64) *DescribeBackupSetListRequest
	GetEndTime() *int64
	SetPageNumber(v int32) *DescribeBackupSetListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupSetListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBackupSetListRequest
	GetRegionId() *string
	SetStartTime(v int64) *DescribeBackupSetListRequest
	GetStartTime() *int64
}

type DescribeBackupSetListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxx
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DestCrossRegion *string `json:"DestCrossRegion,omitempty" xml:"DestCrossRegion,omitempty"`
	// example:
	//
	// 1635707845000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1635707845000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupSetListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeBackupSetListRequest) GetDestCrossRegion() *string {
	return s.DestCrossRegion
}

func (s *DescribeBackupSetListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeBackupSetListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupSetListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupSetListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackupSetListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeBackupSetListRequest) SetDBInstanceName(v string) *DescribeBackupSetListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetDestCrossRegion(v string) *DescribeBackupSetListRequest {
	s.DestCrossRegion = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetEndTime(v int64) *DescribeBackupSetListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetPageNumber(v int32) *DescribeBackupSetListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetPageSize(v int32) *DescribeBackupSetListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetRegionId(v string) *DescribeBackupSetListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupSetListRequest) SetStartTime(v int64) *DescribeBackupSetListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupSetListRequest) Validate() error {
	return dara.Validate(s)
}
