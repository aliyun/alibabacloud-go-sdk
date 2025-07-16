// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenBackupSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeOpenBackupSetRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeOpenBackupSetRequest
	GetRegionId() *string
	SetRestoreTime(v string) *DescribeOpenBackupSetRequest
	GetRestoreTime() *string
}

type DescribeOpenBackupSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2024-10-14T00:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
}

func (s DescribeOpenBackupSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenBackupSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenBackupSetRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeOpenBackupSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpenBackupSetRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *DescribeOpenBackupSetRequest) SetDBInstanceName(v string) *DescribeOpenBackupSetRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeOpenBackupSetRequest) SetRegionId(v string) *DescribeOpenBackupSetRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenBackupSetRequest) SetRestoreTime(v string) *DescribeOpenBackupSetRequest {
	s.RestoreTime = &v
	return s
}

func (s *DescribeOpenBackupSetRequest) Validate() error {
	return dara.Validate(s)
}
