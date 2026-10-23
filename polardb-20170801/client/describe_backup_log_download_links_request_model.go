// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupLogDownloadLinksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupLogIds(v string) *DescribeBackupLogDownloadLinksRequest
	GetBackupLogIds() *string
	SetDBClusterId(v string) *DescribeBackupLogDownloadLinksRequest
	GetDBClusterId() *string
	SetExpireHours(v int32) *DescribeBackupLogDownloadLinksRequest
	GetExpireHours() *int32
}

type DescribeBackupLogDownloadLinksRequest struct {
	// The IDs of the redo log backups.
	//
	// example:
	//
	// 78167430
	BackupLogIds *string `json:"BackupLogIds,omitempty" xml:"BackupLogIds,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The validity period of the download links in hours. Valid values: 1 to 596523.
	//
	// example:
	//
	// 3
	ExpireHours *int32 `json:"ExpireHours,omitempty" xml:"ExpireHours,omitempty"`
}

func (s DescribeBackupLogDownloadLinksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLogDownloadLinksRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogDownloadLinksRequest) GetBackupLogIds() *string {
	return s.BackupLogIds
}

func (s *DescribeBackupLogDownloadLinksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeBackupLogDownloadLinksRequest) GetExpireHours() *int32 {
	return s.ExpireHours
}

func (s *DescribeBackupLogDownloadLinksRequest) SetBackupLogIds(v string) *DescribeBackupLogDownloadLinksRequest {
	s.BackupLogIds = &v
	return s
}

func (s *DescribeBackupLogDownloadLinksRequest) SetDBClusterId(v string) *DescribeBackupLogDownloadLinksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupLogDownloadLinksRequest) SetExpireHours(v int32) *DescribeBackupLogDownloadLinksRequest {
	s.ExpireHours = &v
	return s
}

func (s *DescribeBackupLogDownloadLinksRequest) Validate() error {
	return dara.Validate(s)
}
