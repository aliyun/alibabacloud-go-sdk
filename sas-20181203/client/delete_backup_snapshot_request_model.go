// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRegionIdList(v []*string) *DeleteBackupSnapshotRequest
	GetBackupRegionIdList() []*string
	SetBackupSnapshotList(v []*DeleteBackupSnapshotRequestBackupSnapshotList) *DeleteBackupSnapshotRequest
	GetBackupSnapshotList() []*DeleteBackupSnapshotRequestBackupSnapshotList
	SetRetainLatestSnapshot(v bool) *DeleteBackupSnapshotRequest
	GetRetainLatestSnapshot() *bool
}

type DeleteBackupSnapshotRequest struct {
	// The regions for backup.
	BackupRegionIdList []*string `json:"BackupRegionIdList,omitempty" xml:"BackupRegionIdList,omitempty" type:"Repeated"`
	// The backup snapshots.
	BackupSnapshotList []*DeleteBackupSnapshotRequestBackupSnapshotList `json:"BackupSnapshotList,omitempty" xml:"BackupSnapshotList,omitempty" type:"Repeated"`
	// Specifies whether to retain the latest snapshot. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	RetainLatestSnapshot *bool `json:"RetainLatestSnapshot,omitempty" xml:"RetainLatestSnapshot,omitempty"`
}

func (s DeleteBackupSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupSnapshotRequest) GetBackupRegionIdList() []*string {
	return s.BackupRegionIdList
}

func (s *DeleteBackupSnapshotRequest) GetBackupSnapshotList() []*DeleteBackupSnapshotRequestBackupSnapshotList {
	return s.BackupSnapshotList
}

func (s *DeleteBackupSnapshotRequest) GetRetainLatestSnapshot() *bool {
	return s.RetainLatestSnapshot
}

func (s *DeleteBackupSnapshotRequest) SetBackupRegionIdList(v []*string) *DeleteBackupSnapshotRequest {
	s.BackupRegionIdList = v
	return s
}

func (s *DeleteBackupSnapshotRequest) SetBackupSnapshotList(v []*DeleteBackupSnapshotRequestBackupSnapshotList) *DeleteBackupSnapshotRequest {
	s.BackupSnapshotList = v
	return s
}

func (s *DeleteBackupSnapshotRequest) SetRetainLatestSnapshot(v bool) *DeleteBackupSnapshotRequest {
	s.RetainLatestSnapshot = &v
	return s
}

func (s *DeleteBackupSnapshotRequest) Validate() error {
	if s.BackupSnapshotList != nil {
		for _, item := range s.BackupSnapshotList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteBackupSnapshotRequestBackupSnapshotList struct {
	// The ID of the Cloud Backup client.
	//
	// >  You can call the [DescribeSnapshots](~~DescribeSnapshots~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-000a4h33w14ka8xagb2s
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-j6cj8vyajp1fo4atxkae
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which Security Center is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: China (Hangzhou).
	//
	// 	- **ap-southeast-1**: Singapore.
	//
	// 	- **cn-beijing**: China (Beijing).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the snapshot that you want to delete.
	//
	// >  You can call the [DescribeSnapshots](~~DescribeSnapshots~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-000f9p6r5trm6u4dc1iq
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: Elastic Compute Service (ECS) files.
	//
	// 	- **OSS**: Object Storage Service (OSS) buckets.
	//
	// 	- **NAS**: File Storage NAS (NAS) file systems.
	//
	// 	- **OTS_TABLE**: Tablestore instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the backup vault that is used in the restoration task.
	//
	// >  You can call the [DescribeSnapshots](~~DescribeSnapshots~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-0004vhwcs2pmacfzrzt5
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteBackupSnapshotRequestBackupSnapshotList) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupSnapshotRequestBackupSnapshotList) GoString() string {
	return s.String()
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) GetClientId() *string {
	return s.ClientId
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) GetSourceType() *string {
	return s.SourceType
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) GetVaultId() *string {
	return s.VaultId
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) SetClientId(v string) *DeleteBackupSnapshotRequestBackupSnapshotList {
	s.ClientId = &v
	return s
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) SetInstanceId(v string) *DeleteBackupSnapshotRequestBackupSnapshotList {
	s.InstanceId = &v
	return s
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) SetRegionId(v string) *DeleteBackupSnapshotRequestBackupSnapshotList {
	s.RegionId = &v
	return s
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) SetSnapshotId(v string) *DeleteBackupSnapshotRequestBackupSnapshotList {
	s.SnapshotId = &v
	return s
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) SetSourceType(v string) *DeleteBackupSnapshotRequestBackupSnapshotList {
	s.SourceType = &v
	return s
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) SetVaultId(v string) *DeleteBackupSnapshotRequestBackupSnapshotList {
	s.VaultId = &v
	return s
}

func (s *DeleteBackupSnapshotRequestBackupSnapshotList) Validate() error {
	return dara.Validate(s)
}
