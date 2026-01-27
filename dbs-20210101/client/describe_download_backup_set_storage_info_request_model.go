// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadBackupSetStorageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *DescribeDownloadBackupSetStorageInfoRequest
	GetBackupSetId() *string
	SetClusterName(v string) *DescribeDownloadBackupSetStorageInfoRequest
	GetClusterName() *string
	SetDuration(v string) *DescribeDownloadBackupSetStorageInfoRequest
	GetDuration() *string
	SetInstanceName(v string) *DescribeDownloadBackupSetStorageInfoRequest
	GetInstanceName() *string
	SetRegionCode(v string) *DescribeDownloadBackupSetStorageInfoRequest
	GetRegionCode() *string
	SetTaskId(v string) *DescribeDownloadBackupSetStorageInfoRequest
	GetTaskId() *string
}

type DescribeDownloadBackupSetStorageInfoRequest struct {
	// The ID of the backup set.
	//
	// example:
	//
	// 30****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// example:
	//
	// dds-example
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The validity period of the URL that is used as the download destination. Take note of the following items:
	//
	// 	- Default value: 7200. This means that the URL is valid for 2 hours by default.
	//
	// 	- Valid values: 300 to 86400. Unit: seconds. This means that you can specify a validity period in the range of 5 minutes to 1 day.
	//
	// 	- Before you specify this parameter, convert the validity period to seconds. For example, if you want to set the validity period of the URL to 5 minutes, enter 300.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance.
	//
	// > The **BackupSetId*	- parameter is required if you specify the **InstanceName*	- parameter.
	//
	// example:
	//
	// rm-uf6qqf569n435****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The download task ID.
	//
	// 	- The **BackupSetId*	- and **InstanceName*	- parameters are required if you do not specify the **TaskId*	- parameter.
	//
	// 	- To view the download task ID, go to the instance details page in the console and click **Backup and Restoration*	- in the left-side navigation pane. On the **Backup Download*	- tab, view the task ID.
	//
	// example:
	//
	// dt-s0ugzak9****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDownloadBackupSetStorageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadBackupSetStorageInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) GetDuration() *string {
	return s.Duration
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) SetBackupSetId(v string) *DescribeDownloadBackupSetStorageInfoRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) SetClusterName(v string) *DescribeDownloadBackupSetStorageInfoRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) SetDuration(v string) *DescribeDownloadBackupSetStorageInfoRequest {
	s.Duration = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) SetInstanceName(v string) *DescribeDownloadBackupSetStorageInfoRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) SetRegionCode(v string) *DescribeDownloadBackupSetStorageInfoRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) SetTaskId(v string) *DescribeDownloadBackupSetStorageInfoRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) Validate() error {
	return dara.Validate(s)
}
