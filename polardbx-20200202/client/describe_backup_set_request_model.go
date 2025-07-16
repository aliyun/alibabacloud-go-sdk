// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *DescribeBackupSetRequest
	GetBackupSetId() *string
	SetDBInstanceName(v string) *DescribeBackupSetRequest
	GetDBInstanceName() *string
	SetDestCrossRegion(v string) *DescribeBackupSetRequest
	GetDestCrossRegion() *string
	SetRegionId(v string) *DescribeBackupSetRequest
	GetRegionId() *string
}

type DescribeBackupSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0ori2r****
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DestCrossRegion *string `json:"DestCrossRegion,omitempty" xml:"DestCrossRegion,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBackupSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeBackupSetRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeBackupSetRequest) GetDestCrossRegion() *string {
	return s.DestCrossRegion
}

func (s *DescribeBackupSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackupSetRequest) SetBackupSetId(v string) *DescribeBackupSetRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeBackupSetRequest) SetDBInstanceName(v string) *DescribeBackupSetRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBackupSetRequest) SetDestCrossRegion(v string) *DescribeBackupSetRequest {
	s.DestCrossRegion = &v
	return s
}

func (s *DescribeBackupSetRequest) SetRegionId(v string) *DescribeBackupSetRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupSetRequest) Validate() error {
	return dara.Validate(s)
}
