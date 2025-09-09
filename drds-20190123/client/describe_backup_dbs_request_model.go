// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDbsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeBackupDbsRequest
	GetBackupId() *string
	SetDrdsInstanceId(v string) *DescribeBackupDbsRequest
	GetDrdsInstanceId() *string
	SetPreferredRestoreTime(v string) *DescribeBackupDbsRequest
	GetPreferredRestoreTime() *string
}

type DescribeBackupDbsRequest struct {
	// Query by backup set ID
	//
	// example:
	//
	// 201908367
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of a DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// Query by restoration time.
	//
	// example:
	//
	// 1568632541236
	PreferredRestoreTime *string `json:"PreferredRestoreTime,omitempty" xml:"PreferredRestoreTime,omitempty"`
}

func (s DescribeBackupDbsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupDbsRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeBackupDbsRequest) GetPreferredRestoreTime() *string {
	return s.PreferredRestoreTime
}

func (s *DescribeBackupDbsRequest) SetBackupId(v string) *DescribeBackupDbsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupDbsRequest) SetDrdsInstanceId(v string) *DescribeBackupDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBackupDbsRequest) SetPreferredRestoreTime(v string) *DescribeBackupDbsRequest {
	s.PreferredRestoreTime = &v
	return s
}

func (s *DescribeBackupDbsRequest) Validate() error {
	return dara.Validate(s)
}
